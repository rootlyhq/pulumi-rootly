package rootly

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"

	rootlyshim "github.com/rootlyhq/terraform-provider-rootly/v5/pulumi-shim"
)

const rotationURN = "urn:pulumi:test::test::rootly:index/scheduleRotation:ScheduleRotation::primary"

var (
	rotationMapType    = tftypes.Map{ElementType: tftypes.String}
	rotationObjectType = tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"handoff_time": tftypes.String, "handoff_day": tftypes.String,
		"shift_length": tftypes.Number, "shift_length_unit": tftypes.String,
	}}
)

func rotationObject(handoffDay, shiftLength, unit any) tftypes.Value {
	return tftypes.NewValue(rotationObjectType, map[string]tftypes.Value{
		"handoff_time":      tftypes.NewValue(tftypes.String, "09:00"),
		"handoff_day":       tftypes.NewValue(tftypes.String, handoffDay),
		"shift_length":      tftypes.NewValue(tftypes.Number, shiftLength),
		"shift_length_unit": tftypes.NewValue(tftypes.String, unit),
	})
}

func rotationMap(values map[string]any) tftypes.Value {
	m := map[string]tftypes.Value{}
	for k, v := range values {
		m[k] = tftypes.NewValue(tftypes.String, v)
	}
	return tftypes.NewValue(rotationMapType, m)
}

func TestScheduleRotationAttributesConversion(t *testing.T) {
	for _, tc := range []struct {
		name   string
		object tftypes.Value
		m      tftypes.Value
	}{
		{
			name:   "all keys",
			object: rotationObject("monday", big.NewFloat(12), "hours"),
			m: rotationMap(map[string]any{
				"handoff_time": "09:00", "handoff_day": "monday", "shift_length": "12", "shift_length_unit": "hours",
			}),
		},
		{
			name:   "omitted optional keys",
			object: rotationObject(nil, nil, nil),
			m:      rotationMap(map[string]any{"handoff_time": "09:00"}),
		},
		{
			name:   "unknown values",
			object: rotationObject(tftypes.UnknownValue, tftypes.UnknownValue, nil),
			m: rotationMap(map[string]any{
				"handoff_time": "09:00", "handoff_day": tftypes.UnknownValue, "shift_length": tftypes.UnknownValue,
			}),
		},
		{"null", tftypes.NewValue(rotationObjectType, nil), tftypes.NewValue(rotationMapType, nil)},
		{
			"unknown", tftypes.NewValue(rotationObjectType, tftypes.UnknownValue),
			tftypes.NewValue(rotationMapType, tftypes.UnknownValue),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m, err := rotationObjectToMap(tc.object, rotationMapType)
			require.NoError(t, err)
			require.True(t, tc.m.Equal(m), "object to map: got %s", m)
			object, err := rotationMapToObject(tc.m, rotationObjectType)
			require.NoError(t, err)
			require.True(t, tc.object.Equal(object), "map to object: got %s", object)
		})
	}
}

func TestScheduleRotationAttributesConversionErrors(t *testing.T) {
	_, err := rotationObjectToMap(rotationObject(nil, big.NewFloat(1.5), nil), rotationMapType)
	require.ErrorContains(t, err, "shift_length must be a whole number")
	_, err = rotationMapToObject(rotationMap(map[string]any{"shift_length": "twelve"}), rotationObjectType)
	require.ErrorContains(t, err, "shift_length must be an integer")
	_, err = rotationMapToObject(rotationMap(map[string]any{"unexpected": "x"}), rotationObjectType)
	require.ErrorContains(t, err, `unsupported key "unexpected"`)
}

// The adapter forwards only the interfaces below. Fail if upstream starts
// implementing others, because embedding would silently hide them.
func TestScheduleRotationAdapterCoversUpstreamInterfaces(t *testing.T) {
	ctx := context.Background()
	upstream := rootlyshim.New("4.0.0")
	for name, implemented := range map[string]bool{
		"ConfigValidators":   is[provider.ProviderWithConfigValidators](upstream),
		"Functions":          is[provider.ProviderWithFunctions](upstream),
		"EphemeralResources": is[provider.ProviderWithEphemeralResources](upstream),
		"MetaSchema":         is[provider.ProviderWithMetaSchema](upstream),
		"ListResources":      is[provider.ProviderWithListResources](upstream),
		"Actions":            is[provider.ProviderWithActions](upstream),
		"ValidateConfig":     is[provider.ProviderWithValidateConfig](upstream),
		"StateStores":        is[provider.ProviderWithStateStores](upstream),
	} {
		require.False(t, implemented, "upstream provider now implements %s; forward it", name)
	}

	var adapted *scheduleRotationResource
	for _, factory := range adaptScheduleRotation(upstream).Resources(ctx) {
		if r, ok := factory().(*scheduleRotationResource); ok {
			adapted = r
		}
	}
	require.NotNil(t, adapted, "rootly_schedule_rotation must be adapted")
	inner := adapted.inner
	for name, implemented := range map[string]bool{
		"ConfigValidators": is[resource.ResourceWithConfigValidators](inner),
		"ModifyPlan":       is[resource.ResourceWithModifyPlan](inner),
		"MoveState":        is[resource.ResourceWithMoveState](inner),
		"UpgradeState":     is[resource.ResourceWithUpgradeState](inner),
		"ValidateConfig":   is[resource.ResourceWithValidateConfig](inner),
		"Identity":         is[resource.ResourceWithIdentity](inner),
		"UpgradeIdentity":  is[resource.ResourceWithUpgradeIdentity](inner),
	} {
		require.False(t, implemented, "upstream schedule rotation now implements %s; forward it", name)
	}

	var resp resource.SchemaResponse
	adapted.Schema(ctx, resource.SchemaRequest{}, &resp)
	require.False(t, resp.Diagnostics.HasError())
	require.IsType(t, schema.SingleNestedAttribute{}, resp.Schema.Attributes[rotationAttributes])
	require.True(t, rotationObjectType.Equal(resp.Schema.Attributes[rotationAttributes].GetType().TerraformType(ctx)))
}

func is[T any](v any) bool {
	_, ok := v.(T)
	return ok
}

// TestScheduleRotationLifecycle drives the muxed provider against a fake API:
// Pulumi sees the v4 typed object while upstream sends its string map.
func TestScheduleRotationLifecycle(t *testing.T) {
	sent := make(chan map[string]any, 2)
	rotation := func(attrs string) string {
		return `{"data":{"id":"rotation-1","type":"schedule_rotations","attributes":{` +
			`"schedule_id":"schedule-1","name":"primary","position":1,` +
			`"schedule_rotationable_type":"ScheduleCustomRotation","active_all_week":true,"active_days":[],` +
			`"active_time_type":"all_day","time_zone":"Etc/UTC",` +
			`"schedule_rotationable_attributes":` + attrs + `}}}`
	}
	current := `{"handoff_time":"09:00","shift_length":12,"shift_length_unit":"hours"}`
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/vnd.api+json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/v1/schedules/schedule-1/schedule_rotations",
			r.Method == http.MethodPut && r.URL.Path == "/v1/schedule_rotations/rotation-1",
			r.Method == http.MethodPatch && r.URL.Path == "/v1/schedule_rotations/rotation-1":
			body, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			var payload struct {
				Data struct {
					Attributes map[string]any `json:"attributes"`
				} `json:"data"`
			}
			require.NoError(t, json.Unmarshal(body, &payload))
			attrs := payload.Data.Attributes["schedule_rotationable_attributes"].(map[string]any)
			sent <- attrs
			encoded, err := json.Marshal(attrs)
			require.NoError(t, err)
			current = string(encoded)
			w.WriteHeader(http.StatusCreated)
		case r.Method == http.MethodGet && r.URL.Path == "/v1/schedule_rotations/rotation-1":
		default:
			t.Errorf("unexpected API request: %s %s", r.Method, r.URL)
			http.NotFound(w, r)
			return
		}
		_, err := fmt.Fprint(w, rotation(current))
		require.NoError(t, err)
	}))
	defer api.Close()

	ctx := context.Background()
	info, data := testProvider(t)
	server, err := pfbridge.MakeMuxedServer(ctx, "rootly", info, data)(nil)
	require.NoError(t, err)
	config, err := structpb.NewStruct(map[string]any{"apiHost": api.URL, "apiToken": "test-token"})
	require.NoError(t, err)
	_, err = server.Configure(ctx, &pulumirpc.ConfigureRequest{Args: config, AcceptSecrets: true, AcceptResources: true})
	require.NoError(t, err)

	news, err := structpb.NewStruct(map[string]any{
		"name": "primary", "scheduleId": "schedule-1", "scheduleRotationableType": "ScheduleCustomRotation",
		"scheduleRotationableAttributes": map[string]any{
			"handoffTime": "09:00", "shiftLength": 12, "shiftLengthUnit": "hours",
		},
	})
	require.NoError(t, err)
	checked, err := server.Check(ctx, &pulumirpc.CheckRequest{Urn: rotationURN, News: news})
	require.NoError(t, err)
	require.Empty(t, checked.Failures)

	created, err := server.Create(ctx, &pulumirpc.CreateRequest{Urn: rotationURN, Properties: checked.Inputs})
	require.NoError(t, err)
	require.Equal(t, "rotation-1", created.Id)
	require.Equal(t, map[string]any{
		"handoff_time": "09:00", "handoff_day": nil, "shift_length": float64(12), "shift_length_unit": "hours",
	}, <-sent, "upstream must send the snake_case attributes")
	want := map[string]any{"handoffTime": "09:00", "shiftLength": float64(12), "shiftLengthUnit": "hours"}
	require.Equal(t, want, created.Properties.AsMap()["scheduleRotationableAttributes"])

	// Refreshing v4 state round-trips the typed object without a diff.
	read, err := server.Read(ctx, &pulumirpc.ReadRequest{
		Urn: rotationURN, Id: "rotation-1", Properties: created.Properties, Inputs: checked.Inputs,
	})
	require.NoError(t, err)
	require.Equal(t, want, read.Properties.AsMap()["scheduleRotationableAttributes"])
	diff, err := server.Diff(ctx, &pulumirpc.DiffRequest{
		Urn: rotationURN, Id: "rotation-1", Olds: read.Properties, News: checked.Inputs,
	})
	require.NoError(t, err)
	require.Equal(t, pulumirpc.DiffResponse_DIFF_NONE, diff.Changes)

	// Refreshing v3 state upgrades snake_case strings to the typed object.
	v3 := created.Properties.AsMap()
	v3["scheduleRotationableAttributes"] = map[string]any{
		"handoff_time": "09:00", "shift_length": "12", "shift_length_unit": "hours",
	}
	v3State, err := structpb.NewStruct(v3)
	require.NoError(t, err)
	read, err = server.Read(ctx, &pulumirpc.ReadRequest{Urn: rotationURN, Id: "rotation-1", Properties: v3State})
	require.NoError(t, err)
	require.Equal(t, want, read.Properties.AsMap()["scheduleRotationableAttributes"])

	// Updating sends a changed shift length and drops removed optional keys.
	updatedNews, err := structpb.NewStruct(map[string]any{
		"name": "primary", "scheduleId": "schedule-1", "scheduleRotationableType": "ScheduleCustomRotation",
		"scheduleRotationableAttributes": map[string]any{"handoffTime": "10:00", "handoffDay": "monday"},
	})
	require.NoError(t, err)
	checked, err = server.Check(ctx, &pulumirpc.CheckRequest{
		Urn: rotationURN, News: updatedNews, Olds: created.Properties,
	})
	require.NoError(t, err)
	require.Empty(t, checked.Failures)
	updated, err := server.Update(ctx, &pulumirpc.UpdateRequest{
		Urn: rotationURN, Id: "rotation-1", Olds: created.Properties, News: checked.Inputs,
	})
	require.NoError(t, err)
	require.Equal(t, map[string]any{
		"handoff_time": "10:00", "handoff_day": "monday", "shift_length": nil, "shift_length_unit": nil,
	}, <-sent)
	require.Equal(t, map[string]any{"handoffTime": "10:00", "handoffDay": "monday"},
		updated.Properties.AsMap()["scheduleRotationableAttributes"])

	invalid, err := structpb.NewStruct(map[string]any{
		"name": "primary", "scheduleId": "schedule-1",
		"scheduleRotationableAttributes": map[string]any{"shiftLength": 12},
	})
	require.NoError(t, err)
	checked, err = server.Check(ctx, &pulumirpc.CheckRequest{Urn: rotationURN, News: invalid})
	require.NoError(t, err)
	require.NotEmpty(t, checked.Failures, "handoffTime stays required")
	require.True(t, strings.Contains(fmt.Sprint(checked.Failures), "handoff"))
}
