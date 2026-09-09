package rootly

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/reservedkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// upgradeScheduleRotationState converts the SDKv2 string map stored by v3 into
// the Plugin Framework's typed object. Current state is returned unchanged.
func upgradeScheduleRotationState(_ context.Context, state resource.PropertyMap) (resource.PropertyMap, error) {
	value, ok := state["scheduleRotationableAttributes"]
	if !ok {
		return state, nil
	}
	secret := value.IsSecret()
	if secret {
		value = value.SecretValue().Element
	}
	if !value.IsObject() {
		return state, nil
	}
	attrs := value.ObjectValue().Copy()
	changed := false
	for oldKey, newKey := range map[resource.PropertyKey]resource.PropertyKey{
		"handoff_time": "handoffTime", "handoff_day": "handoffDay",
		"shift_length": "shiftLength", "shift_length_unit": "shiftLengthUnit",
	} {
		if v, exists := attrs[oldKey]; exists {
			if _, exists := attrs[newKey]; !exists {
				attrs[newKey] = v
			}
			delete(attrs, oldKey)
			changed = true
		}
	}
	length := attrs["shiftLength"]
	lengthSecret := length.IsSecret()
	if lengthSecret {
		length = length.SecretValue().Element
	}
	if length.IsString() {
		n, err := strconv.ParseInt(length.StringValue(), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("saved schedule rotation shift_length must be an integer")
		}
		length = resource.NewNumberProperty(float64(n))
		if lengthSecret {
			length = resource.MakeSecret(length)
		}
		attrs["shiftLength"] = length
		changed = true
	}
	if !changed {
		return state, nil
	}
	state = state.Copy()
	value = resource.NewObjectProperty(attrs)
	if secret {
		value = resource.MakeSecret(value)
	}
	state["scheduleRotationableAttributes"] = value
	// The SDKv2 delta describes a string map and cannot recover the new object.
	delete(state, reservedkeys.RawStateDelta)
	return state, nil
}
