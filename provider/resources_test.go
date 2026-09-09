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

func TestMuxedProviderRoutesAPIRequests(t *testing.T) {
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/vnd.api+json")
		var body string
		switch r.URL.Path {
		case "/v1/services/service-1":
			body = `{"data":{"id":"service-1","type":"services","attributes":{"name":"API","slug":"api"}}}`
		case "/v1/services":
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
		{"rootly:index/getSeverity:getSeverity", map[string]any{"slug": "sev1"}, "SEV1"},
		{"rootly:index/getServices:getServices", map[string]any{}, ""},
	} {
		t.Run(tc.token, func(t *testing.T) {
			args, err := structpb.NewStruct(tc.args)
			require.NoError(t, err)
			result, err := server.Invoke(ctx, &pulumirpc.InvokeRequest{Tok: tc.token, Args: args})
			require.NoError(t, err)
			require.Empty(t, result.Failures)
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
