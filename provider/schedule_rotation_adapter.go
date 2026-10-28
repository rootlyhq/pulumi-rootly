package rootly

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Rootly 5.23.5 models schedule_rotationable_attributes as a string map. The
// adapter keeps the typed object published by v4 and converts at the boundary.
const (
	scheduleRotationType      = "_schedule_rotation"
	rotationAttributes        = "schedule_rotationable_attributes"
	rotationHandoffTime       = "handoff_time"
	rotationHandoffDay        = "handoff_day"
	rotationShiftLength       = "shift_length"
	rotationShiftLengthUnit   = "shift_length_unit"
	rotationAttributesSummary = "handoff_time and/or handoff_day may be required, depending on " +
		"schedule_rotationable_type. Please see API docs for options based on schedule_rotationable_type: " +
		"https://docs.rootly.com/api-reference/schedulerotations/creates-a-schedule-rotation" +
		"#response-data-attributes-schedule-rotationable-attributes"
)

var rotationAttributeKeys = []string{
	rotationHandoffTime, rotationHandoffDay, rotationShiftLength, rotationShiftLengthUnit,
}

// scheduleRotationProvider replaces the upstream schedule rotation with an
// adapter. Upstream implements only the base provider interface; a test fails
// if that changes, since embedding would hide the optional interfaces.
type scheduleRotationProvider struct {
	provider.Provider
}

func adaptScheduleRotation(p provider.Provider) provider.Provider {
	return scheduleRotationProvider{p}
}

func (p scheduleRotationProvider) Resources(ctx context.Context) []func() resource.Resource {
	var meta provider.MetadataResponse
	p.Metadata(ctx, provider.MetadataRequest{}, &meta)
	factories := p.Provider.Resources(ctx)
	adapted := make([]func() resource.Resource, len(factories))
	for i, factory := range factories {
		adapted[i] = func() resource.Resource {
			inner := factory()
			var resp resource.MetadataResponse
			inner.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: meta.TypeName}, &resp)
			if resp.TypeName != meta.TypeName+scheduleRotationType {
				return inner
			}
			var schemaResp resource.SchemaResponse
			inner.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
			if _, ok := schemaResp.Schema.Attributes[rotationAttributes].(schema.MapAttribute); !ok {
				return inner
			}
			return &scheduleRotationResource{inner: inner}
		}
	}
	return adapted
}

// scheduleRotationResource forwards upstream's resource interfaces with the
// typed attributes object while upstream receives its string map.
type scheduleRotationResource struct {
	inner resource.Resource
}

var (
	_ resource.ResourceWithConfigure   = &scheduleRotationResource{}
	_ resource.ResourceWithImportState = &scheduleRotationResource{}
)

func (r *scheduleRotationResource) Metadata(
	ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse,
) {
	r.inner.Metadata(ctx, req, resp)
}

func (r *scheduleRotationResource) innerSchema(ctx context.Context) (schema.Schema, diag.Diagnostics) {
	var resp resource.SchemaResponse
	r.inner.Schema(ctx, resource.SchemaRequest{}, &resp)
	return resp.Schema, resp.Diagnostics
}

func (r *scheduleRotationResource) Schema(
	ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse,
) {
	inner, diags := r.innerSchema(ctx)
	resp.Diagnostics.Append(diags...)
	attrs := make(map[string]schema.Attribute, len(inner.Attributes))
	for name, attr := range inner.Attributes {
		attrs[name] = attr
	}
	attrs[rotationAttributes] = schema.SingleNestedAttribute{
		MarkdownDescription: rotationAttributesSummary,
		Required:            true,
		Attributes: map[string]schema.Attribute{
			rotationHandoffTime: schema.StringAttribute{
				MarkdownDescription: "Hand off time. Only applicable for daily, weekly/biweekly, monthly, and custom rotations.",
				Required:            true,
			},
			rotationHandoffDay: schema.StringAttribute{
				MarkdownDescription: "Hand off day. Only applicable for weekly/biweekly, and monthly.",
				Optional:            true,
			},
			rotationShiftLength: schema.Int64Attribute{
				MarkdownDescription: "Shift length for custom rotation.",
				Optional:            true,
			},
			rotationShiftLengthUnit: schema.StringAttribute{
				MarkdownDescription: "Shift length unit for custom rotation. Value must be one of `hours`, `days`, `weeks`.",
				Optional:            true,
			},
		},
	}
	inner.Attributes = attrs
	resp.Schema = inner
}

func (r *scheduleRotationResource) Configure(
	ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse,
) {
	if inner, ok := r.inner.(resource.ResourceWithConfigure); ok {
		inner.Configure(ctx, req, resp)
	}
}

// ImportState only sets the ID, which both schemas share.
func (r *scheduleRotationResource) ImportState(
	ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse,
) {
	if inner, ok := r.inner.(resource.ResourceWithImportState); ok {
		inner.ImportState(ctx, req, resp)
		return
	}
	resp.Diagnostics.AddError("Resource Import Not Implemented", "rootly_schedule_rotation does not support import.")
}

// schemas returns the upstream and adapted schemas.
func (r *scheduleRotationResource) schemas(
	ctx context.Context, diags *diag.Diagnostics,
) (schema.Schema, schema.Schema) {
	inner, d := r.innerSchema(ctx)
	diags.Append(d...)
	var outer resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &outer)
	return inner, outer.Schema
}

func convertRotationValue(
	ctx context.Context, diags *diag.Diagnostics, raw tftypes.Value, target schema.Schema,
	convert func(tftypes.Value, tftypes.Type) (tftypes.Value, error),
) tftypes.Value {
	v, err := convertRotation(ctx, raw, target, convert)
	if err != nil {
		diags.AddError("Unable to convert schedule rotation attributes", err.Error())
		return tftypes.NewValue(target.Type().TerraformType(ctx), nil)
	}
	return v
}

func (r *scheduleRotationResource) Create(
	ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse,
) {
	inner, outer := r.schemas(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	innerReq := req
	innerReq.Config.Schema = inner
	innerReq.Config.Raw = convertRotationValue(ctx, &resp.Diagnostics, req.Config.Raw, inner, rotationObjectToMap)
	innerReq.Plan.Schema = inner
	innerReq.Plan.Raw = convertRotationValue(ctx, &resp.Diagnostics, req.Plan.Raw, inner, rotationObjectToMap)
	if resp.Diagnostics.HasError() {
		return
	}
	innerResp := *resp
	innerResp.State.Schema = inner
	innerResp.State.Raw = tftypes.NewValue(inner.Type().TerraformType(ctx), nil)
	r.inner.Create(ctx, innerReq, &innerResp)
	*resp = innerResp
	resp.State.Schema = outer
	resp.State.Raw = convertRotationValue(ctx, &resp.Diagnostics, innerResp.State.Raw, outer, rotationMapToObject)
}

func (r *scheduleRotationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	inner, outer := r.schemas(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	innerReq := req
	innerReq.State.Schema = inner
	innerReq.State.Raw = convertRotationValue(ctx, &resp.Diagnostics, req.State.Raw, inner, rotationObjectToMap)
	if resp.Diagnostics.HasError() {
		return
	}
	innerResp := *resp
	innerResp.State.Schema = inner
	innerResp.State.Raw = innerReq.State.Raw.Copy()
	r.inner.Read(ctx, innerReq, &innerResp)
	*resp = innerResp
	resp.State.Schema = outer
	resp.State.Raw = convertRotationValue(ctx, &resp.Diagnostics, innerResp.State.Raw, outer, rotationMapToObject)
}

func (r *scheduleRotationResource) Update(
	ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse,
) {
	inner, outer := r.schemas(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	innerReq := req
	innerReq.Config.Schema = inner
	innerReq.Config.Raw = convertRotationValue(ctx, &resp.Diagnostics, req.Config.Raw, inner, rotationObjectToMap)
	innerReq.Plan.Schema = inner
	innerReq.Plan.Raw = convertRotationValue(ctx, &resp.Diagnostics, req.Plan.Raw, inner, rotationObjectToMap)
	innerReq.State.Schema = inner
	innerReq.State.Raw = convertRotationValue(ctx, &resp.Diagnostics, req.State.Raw, inner, rotationObjectToMap)
	if resp.Diagnostics.HasError() {
		return
	}
	innerResp := *resp
	innerResp.State.Schema = inner
	innerResp.State.Raw = tftypes.NewValue(inner.Type().TerraformType(ctx), nil)
	r.inner.Update(ctx, innerReq, &innerResp)
	*resp = innerResp
	resp.State.Schema = outer
	resp.State.Raw = convertRotationValue(ctx, &resp.Diagnostics, innerResp.State.Raw, outer, rotationMapToObject)
}

func (r *scheduleRotationResource) Delete(
	ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse,
) {
	inner, outer := r.schemas(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	innerReq := req
	innerReq.State.Schema = inner
	innerReq.State.Raw = convertRotationValue(ctx, &resp.Diagnostics, req.State.Raw, inner, rotationObjectToMap)
	if resp.Diagnostics.HasError() {
		return
	}
	innerResp := *resp
	innerResp.State.Schema = inner
	innerResp.State.Raw = convertRotationValue(ctx, &resp.Diagnostics, resp.State.Raw, inner, rotationObjectToMap)
	r.inner.Delete(ctx, innerReq, &innerResp)
	*resp = innerResp
	resp.State.Schema = outer
	resp.State.Raw = convertRotationValue(ctx, &resp.Diagnostics, innerResp.State.Raw, outer, rotationMapToObject)
}

// convertRotation rewrites the attributes value of a resource object into the
// target schema's type, leaving every other attribute untouched.
func convertRotation(
	ctx context.Context, raw tftypes.Value, target schema.Schema,
	convert func(tftypes.Value, tftypes.Type) (tftypes.Value, error),
) (tftypes.Value, error) {
	typ, ok := target.Type().TerraformType(ctx).(tftypes.Object)
	if !ok {
		return tftypes.Value{}, fmt.Errorf("unexpected schema type %s", target.Type().TerraformType(ctx))
	}
	if raw.IsNull() {
		return tftypes.NewValue(typ, nil), nil
	}
	if !raw.IsKnown() {
		return tftypes.NewValue(typ, tftypes.UnknownValue), nil
	}
	var current map[string]tftypes.Value
	if err := raw.As(&current); err != nil {
		return tftypes.Value{}, err
	}
	// As shares the map backing raw; copy it so the caller's value is untouched.
	attrs := make(map[string]tftypes.Value, len(current))
	for name, value := range current {
		attrs[name] = value
	}
	converted, err := convert(attrs[rotationAttributes], typ.AttributeTypes[rotationAttributes])
	if err != nil {
		return tftypes.Value{}, fmt.Errorf("%s: %w", rotationAttributes, err)
	}
	attrs[rotationAttributes] = converted
	if err := tftypes.ValidateValue(typ, attrs); err != nil {
		return tftypes.Value{}, err
	}
	return tftypes.NewValue(typ, attrs), nil
}

// rotationObjectToMap converts the typed attributes into upstream's string
// map, dropping null attributes and formatting shift_length as an integer.
func rotationObjectToMap(v tftypes.Value, typ tftypes.Type) (tftypes.Value, error) {
	if v.IsNull() || !v.IsKnown() {
		return tftypes.NewValue(typ, valueOrUnknown(v)), nil
	}
	var attrs map[string]tftypes.Value
	if err := v.As(&attrs); err != nil {
		return tftypes.Value{}, err
	}
	out := map[string]tftypes.Value{}
	for _, key := range rotationAttributeKeys {
		attr, ok := attrs[key]
		if !ok || attr.IsNull() {
			continue
		}
		if key != rotationShiftLength || !attr.IsKnown() {
			out[key] = tftypes.NewValue(tftypes.String, valueOrUnknown(attr))
			continue
		}
		var n big.Float
		if err := attr.As(&n); err != nil {
			return tftypes.Value{}, err
		}
		i, accuracy := n.Int64()
		if accuracy != big.Exact {
			return tftypes.Value{}, fmt.Errorf("%s must be a whole number, got %s", rotationShiftLength, n.String())
		}
		out[key] = tftypes.NewValue(tftypes.String, strconv.FormatInt(i, 10))
	}
	return tftypes.NewValue(typ, out), nil
}

// rotationMapToObject converts upstream's string map into the typed
// attributes, using null for missing keys and parsing shift_length.
func rotationMapToObject(v tftypes.Value, typ tftypes.Type) (tftypes.Value, error) {
	objType, ok := typ.(tftypes.Object)
	if !ok {
		return tftypes.Value{}, fmt.Errorf("unexpected attributes type %s", typ)
	}
	if v.IsNull() || !v.IsKnown() {
		return tftypes.NewValue(typ, valueOrUnknown(v)), nil
	}
	var values map[string]tftypes.Value
	if err := v.As(&values); err != nil {
		return tftypes.Value{}, err
	}
	out := make(map[string]tftypes.Value, len(objType.AttributeTypes))
	for key, attrType := range objType.AttributeTypes {
		out[key] = tftypes.NewValue(attrType, nil)
	}
	for key, value := range values {
		attrType, ok := objType.AttributeTypes[key]
		if !ok {
			return tftypes.Value{}, fmt.Errorf("unsupported key %q", key)
		}
		if key != rotationShiftLength || value.IsNull() || !value.IsKnown() {
			out[key] = tftypes.NewValue(attrType, valueOrUnknown(value))
			continue
		}
		var s string
		if err := value.As(&s); err != nil {
			return tftypes.Value{}, err
		}
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return tftypes.Value{}, fmt.Errorf("%s must be an integer, got %q", rotationShiftLength, s)
		}
		out[key] = tftypes.NewValue(attrType, new(big.Float).SetInt64(n))
	}
	return tftypes.NewValue(typ, out), nil
}

// valueOrUnknown returns a primitive's Go value for re-typing it, keeping
// null and unknown markers.
func valueOrUnknown(v tftypes.Value) any {
	if !v.IsKnown() {
		return tftypes.UnknownValue
	}
	if v.IsNull() {
		return nil
	}
	var s string
	if err := v.As(&s); err == nil {
		return s
	}
	return nil
}
