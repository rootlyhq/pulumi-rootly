package rootly

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"
)

func testProvider(t *testing.T) (tfbridge.ProviderInfo, []byte) {
	t.Helper()
	info := Provider()
	info.Version = "4.0.0"
	data, err := os.ReadFile("cmd/pulumi-resource-rootly/schema.json")
	require.NoError(t, err)
	return info, data
}

func TestSchemaRetainsResourcesAndMuxDispatch(t *testing.T) {
	info, data := testProvider(t)
	var spec schema.PackageSpec
	require.NoError(t, json.Unmarshal(data, &spec))
	for name, resource := range info.Resources {
		require.Contains(t, spec.Resources, string(resource.Tok), "missing resource %s", name)
	}
	for name, source := range info.DataSources {
		require.Contains(t, spec.Functions, string(source.Tok), "missing data source %s", name)
	}
	var bridgeMetadata struct {
		Mux struct {
			Resources map[string]int `json:"resources"`
			Functions map[string]int `json:"functions"`
		} `json:"mux"`
	}
	require.NoError(t, json.Unmarshal(metadata, &bridgeMetadata))
	for token, index := range map[string]int{
		"rootly:index/scheduleRotation:ScheduleRotation": 1,
		"rootly:index/service:Service":                   0,
		"rootly:index/userOnCallRole:UserOnCallRole":     0,
	} {
		require.Contains(t, spec.Resources, token)
		require.Contains(t, bridgeMetadata.Mux.Resources, token)
		require.Equal(t, index, bridgeMetadata.Mux.Resources[token])
	}
	for token, index := range map[string]int{
		"rootly:index/getService:getService":   1,
		"rootly:index/getServices:getServices": 1,
		"rootly:index/getSeverity:getSeverity": 0,
		"rootly:index/getUsers:getUsers":       0,
	} {
		require.Contains(t, spec.Functions, token)
		require.Contains(t, bridgeMetadata.Mux.Functions, token)
		require.Equal(t, index, bridgeMetadata.Mux.Functions[token])
	}
}

func TestCredentialSchemaMarksSecrets(t *testing.T) {
	_, data := testProvider(t)
	var spec schema.PackageSpec
	require.NoError(t, json.Unmarshal(data, &spec))
	for token, properties := range map[string][]string{
		"rootly:index/heartbeat:Heartbeat":               {"pingUrl", "secret"},
		"rootly:index/secret:Secret":                     {"secret"},
		"rootly:index/statusPage:StatusPage":             {"authenticationPassword"},
		"rootly:index/webhooksEndpoint:WebhooksEndpoint": {"secret"},
	} {
		for _, property := range properties {
			require.True(t, spec.Resources[token].Properties[property].Secret, "%s.%s output", token, property)
			if input, ok := spec.Resources[token].InputProperties[property]; ok {
				require.True(t, input.Secret, "%s.%s input", token, property)
			}
		}
	}
	for token, property := range map[string]string{
		"rootly:index/WebhooksEndpointCustomHeader:WebhooksEndpointCustomHeader":                         "value",
		"rootly:index/WorkflowTaskHttpClientTaskParams:WorkflowTaskHttpClientTaskParams":                 "headers",
		"rootly:index/WorkflowTaskCreateWebexMeetingTaskParams:WorkflowTaskCreateWebexMeetingTaskParams": "password",
		"rootly:index/WorkflowTaskCreateZoomMeetingTaskParams:WorkflowTaskCreateZoomMeetingTaskParams":   "password",
	} {
		require.True(t, spec.Types[token].Properties[property].Secret, "%s.%s", token, property)
	}
}

//nolint:gosec // Synthetic credential fixtures verify secret propagation.
func TestCredentialInputsAreSecrets(t *testing.T) {
	credential := t.Name()
	ctx := context.Background()
	info, data := testProvider(t)
	server, err := pfbridge.MakeMuxedServer(ctx, "rootly", info, data)(nil)
	require.NoError(t, err)
	config, err := structpb.NewStruct(map[string]any{"apiHost": "http://127.0.0.1", "apiToken": "test-token"})
	require.NoError(t, err)
	_, err = server.Configure(ctx, &pulumirpc.ConfigureRequest{
		Args: config, AcceptSecrets: true, AcceptResources: true,
	})
	require.NoError(t, err)

	for _, tc := range []struct {
		token          string
		args           map[string]any
		secretPaths    [][]any
		publicProperty string
	}{
		{
			token:          "rootly:index/secret:Secret",
			args:           map[string]any{"name": "test-secret", "secret": credential},
			secretPaths:    [][]any{{"secret"}},
			publicProperty: "name",
		},
		{
			token:          "rootly:index/statusPage:StatusPage",
			args:           map[string]any{"title": "Status", "authenticationPassword": credential},
			secretPaths:    [][]any{{"authenticationPassword"}},
			publicProperty: "title",
		},
		{
			token: "rootly:index/webhooksEndpoint:WebhooksEndpoint",
			args: map[string]any{
				"url": "https://example.com/hook", "secret": credential,
				"customHeaders": []any{map[string]any{"name": "Authorization", "value": credential}},
			},
			secretPaths:    [][]any{{"secret"}, {"customHeaders", 0, "value"}},
			publicProperty: "url",
		},
		{
			token: "rootly:index/workflowTaskHttpClient:WorkflowTaskHttpClient",
			args: map[string]any{
				"workflowId": "workflow-1",
				"taskParams": map[string]any{"url": "https://example.com", "succeedOnStatus": "200", "headers": credential},
			},
			secretPaths:    [][]any{{"taskParams", "headers"}},
			publicProperty: "workflowId",
		},
		{
			token: "rootly:index/workflowTaskCreateWebexMeeting:WorkflowTaskCreateWebexMeeting",
			args: map[string]any{
				"workflowId": "workflow-1",
				"taskParams": map[string]any{"topic": "Incident", "password": credential},
			},
			secretPaths:    [][]any{{"taskParams", "password"}},
			publicProperty: "workflowId",
		},
		{
			token: "rootly:index/workflowTaskCreateZoomMeeting:WorkflowTaskCreateZoomMeeting",
			args: map[string]any{
				"workflowId": "workflow-1",
				"taskParams": map[string]any{"topic": "Incident", "password": credential},
			},
			secretPaths:    [][]any{{"taskParams", "password"}},
			publicProperty: "workflowId",
		},
	} {
		t.Run(tc.token, func(t *testing.T) {
			args, err := structpb.NewStruct(tc.args)
			require.NoError(t, err)
			checked, err := server.Check(ctx, &pulumirpc.CheckRequest{
				Urn: "urn:pulumi:test::test::" + tc.token + "::credential-test", News: args,
			})
			require.NoError(t, err)
			require.Empty(t, checked.Failures)
			properties, err := plugin.UnmarshalProperties(checked.Inputs, plugin.MarshalOptions{KeepSecrets: true})
			require.NoError(t, err)
			for _, path := range tc.secretPaths {
				value := resource.NewObjectProperty(properties)
				for _, segment := range path {
					switch key := segment.(type) {
					case string:
						value = value.ObjectValue()[resource.PropertyKey(key)]
					case int:
						value = value.ArrayValue()[key]
					}
				}
				require.True(t, value.IsSecret(), "%v must be secret", path)
				require.Equal(t, credential, value.SecretValue().Element.StringValue())
			}
			public := properties[resource.PropertyKey(tc.publicProperty)]
			require.False(t, public.IsSecret())
			require.Equal(t, tc.args[tc.publicProperty], public.StringValue())
		})
	}
}

func TestHeartbeatReadOutputsAreSecrets(t *testing.T) {
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/v1/heartbeats/heartbeat-1", r.URL.Path)
		require.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/vnd.api+json")
		_, err := fmt.Fprint(w, `{
			"data": {
				"id": "heartbeat-1", "type": "heartbeats",
				"attributes": {
					"name": "Heartbeat", "secret": "heartbeat-credential",
					"ping_url": "https://example.com/heartbeat-credential"
				}
			}
		}`)
		require.NoError(t, err)
	}))
	defer api.Close()

	ctx := context.Background()
	info, data := testProvider(t)
	server, err := pfbridge.MakeMuxedServer(ctx, "rootly", info, data)(nil)
	require.NoError(t, err)
	config, err := structpb.NewStruct(map[string]any{"apiHost": api.URL, "apiToken": "test-token"})
	require.NoError(t, err)
	_, err = server.Configure(ctx, &pulumirpc.ConfigureRequest{
		Args: config, AcceptSecrets: true, AcceptResources: true,
	})
	require.NoError(t, err)
	state, err := structpb.NewStruct(map[string]any{"name": "Heartbeat"})
	require.NoError(t, err)
	result, err := server.Read(ctx, &pulumirpc.ReadRequest{
		Urn: "urn:pulumi:test::test::rootly:index/heartbeat:Heartbeat::test",
		Id:  "heartbeat-1", Properties: state,
	})
	require.NoError(t, err)
	properties, err := plugin.UnmarshalProperties(result.Properties, plugin.MarshalOptions{KeepSecrets: true})
	require.NoError(t, err)
	for name, expected := range map[resource.PropertyKey]string{
		"secret":  "heartbeat-credential",
		"pingUrl": "https://example.com/heartbeat-credential",
	} {
		value := properties[name]
		require.True(t, value.IsSecret(), "%s must be secret", name)
		require.Equal(t, expected, value.SecretValue().Element.StringValue())
	}
	require.False(t, properties["name"].IsSecret())
	require.Equal(t, "Heartbeat", properties["name"].StringValue())
}

func TestMuxedProviderRoutesAPIRequests(t *testing.T) {
	serviceQueries := make(chan string, 4)
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/vnd.api+json")
		var body string
		switch r.URL.Path {
		case "/v1/services/service-1":
			body = `{"data":{"id":"service-1","type":"services","attributes":{"name":"API","slug":"api"}}}`
		case "/v1/services":
			serviceQueries <- r.URL.Query().Get("filter[slug]")
			body = `{"data":[{"id":"service-1","type":"services","attributes":{"name":"API","slug":"api"}}]}`
		case "/v1/severities":
			body = `{"data":[{"id":"severity-1","type":"severities","attributes":{"name":"SEV1","slug":"sev1"}}]}`
		default:
			t.Errorf("unexpected API request: %s", r.URL)
			http.NotFound(w, r)
			return
		}
		_, err := fmt.Fprint(w, body)
		require.NoError(t, err)
	}))
	defer api.Close()
	ctx := context.Background()
	info, data := testProvider(t)
	server, err := pfbridge.MakeMuxedServer(ctx, "rootly", info, data)(nil)
	require.NoError(t, err)
	args, err := structpb.NewStruct(map[string]any{"apiHost": api.URL, "apiToken": "test-token"})
	require.NoError(t, err)
	_, err = server.Configure(ctx, &pulumirpc.ConfigureRequest{
		Args: args, AcceptSecrets: true, AcceptResources: true,
	})
	require.NoError(t, err)
	for _, tc := range []struct {
		token string
		args  map[string]any
		name  string
	}{
		{"rootly:index/getService:getService", map[string]any{"id": "service-1"}, "API"},
		{"rootly:index/getService:getService", map[string]any{"slug": "api"}, "API"},
		{"rootly:index/getSeverity:getSeverity", map[string]any{"slug": "sev1"}, "SEV1"},
		{"rootly:index/getServices:getServices", map[string]any{}, ""},
		{"rootly:index/getServices:getServices", map[string]any{"slug": "api"}, ""},
	} {
		t.Run(tc.token, func(t *testing.T) {
			args, err := structpb.NewStruct(tc.args)
			require.NoError(t, err)
			result, err := server.Invoke(ctx, &pulumirpc.InvokeRequest{Tok: tc.token, Args: args})
			require.NoError(t, err)
			require.Empty(t, result.Failures)
			if tc.args["slug"] == "api" {
				require.Equal(t, "api", <-serviceQueries, "service filter must reach the upstream API")
			} else if tc.token == "rootly:index/getServices:getServices" {
				require.Empty(t, <-serviceQueries)
			}
			if tc.name != "" {
				require.Equal(t, tc.name, result.Return.AsMap()["name"])
			} else {
				require.Len(t, result.Return.AsMap()["services"], 1)
			}
		})
	}
	rotation, err := structpb.NewStruct(map[string]any{
		"name": "primary", "scheduleId": "schedule-1",
		"scheduleRotationableAttributes": map[string]any{"handoffTime": "09:00"},
	})
	require.NoError(t, err)
	checked, err := server.Check(ctx, &pulumirpc.CheckRequest{
		Urn: "urn:pulumi:test::test::rootly:index/scheduleRotation:ScheduleRotation::primary", News: rotation,
	})
	require.NoError(t, err)
	require.Empty(t, checked.Failures)
	require.Equal(t, "primary", checked.Inputs.AsMap()["name"])
	old := checked.Inputs.AsMap()
	old["id"] = "rotation-1"
	old["scheduleRotationableAttributes"] = map[string]any{"handoff_time": "09:00"}
	oldState, err := structpb.NewStruct(old)
	require.NoError(t, err)
	diff, err := server.Diff(ctx, &pulumirpc.DiffRequest{
		Urn: "urn:pulumi:test::test::rootly:index/scheduleRotation:ScheduleRotation::primary",
		Id:  "rotation-1", Olds: oldState, News: checked.Inputs,
	})
	require.NoError(t, err)
	require.Empty(t, diff.Replaces, "migrating a v3 rotation must not replace it")
}
