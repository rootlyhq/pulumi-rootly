package rootly

import (
	"context"
	"testing"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/reservedkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/require"
)

func TestUpgradeScheduleRotationState(t *testing.T) {
	for _, secret := range []bool{false, true} {
		attrs := resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]any{
			"handoff_time": "09:00", "handoff_day": "monday", "shift_length": "12", "shift_length_unit": "hours",
		}))
		if secret {
			attrs = resource.MakeSecret(attrs)
		}
		old := resource.PropertyMap{
			"scheduleRotationableAttributes": attrs,
			"id":                             resource.NewStringProperty("rotation-1"),
			reservedkeys.RawStateDelta:       resource.NewStringProperty("old delta"),
		}
		upgraded, err := upgradeScheduleRotationState(context.Background(), old)
		require.NoError(t, err)
		require.Equal(t, old["id"], upgraded["id"])
		require.NotContains(t, upgraded, reservedkeys.RawStateDelta)
		got := upgraded["scheduleRotationableAttributes"]
		require.Equal(t, secret, got.IsSecret())
		if secret {
			got = got.SecretValue().Element
		}
		require.Equal(t, resource.NewPropertyMapFromMap(map[string]any{
			"handoffTime": "09:00", "handoffDay": "monday", "shiftLength": float64(12), "shiftLengthUnit": "hours",
		}), got.ObjectValue())
		require.Equal(t, attrs, old["scheduleRotationableAttributes"], "must not mutate prior state")
		again, err := upgradeScheduleRotationState(context.Background(), upgraded)
		require.NoError(t, err)
		require.Equal(t, upgraded, again)
	}
}

func TestUpgradeScheduleRotationStateRejectsInvalidLength(t *testing.T) {
	_, err := upgradeScheduleRotationState(context.Background(), resource.NewPropertyMapFromMap(map[string]any{
		"scheduleRotationableAttributes": map[string]any{"shift_length": "invalid"},
	}))
	require.ErrorContains(t, err, "shift_length must be an integer")
}
