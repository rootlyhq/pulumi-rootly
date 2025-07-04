// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rootly

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rootlyhq/pulumi-rootly/sdk/go/rootly/internal"
)

// DEPRECATED: Please use `FormField` resource instead.
type CustomField struct {
	pulumi.CustomResourceState

	// The default value for text field kinds.
	Default pulumi.StringOutput `pulumi:"default"`
	// The description of the custom_field
	Description pulumi.StringOutput  `pulumi:"description"`
	Enabled     pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The kind of the custom_field
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the custom_field
	Label pulumi.StringOutput `pulumi:"label"`
	// The position of the custom_field
	Position  pulumi.IntOutput         `pulumi:"position"`
	Requireds pulumi.StringArrayOutput `pulumi:"requireds"`
	Showns    pulumi.StringArrayOutput `pulumi:"showns"`
	// The slug of the custom_field
	Slug pulumi.StringOutput `pulumi:"slug"`
}

// NewCustomField registers a new resource with the given unique name, arguments, and options.
func NewCustomField(ctx *pulumi.Context,
	name string, args *CustomFieldArgs, opts ...pulumi.ResourceOption) (*CustomField, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Label == nil {
		return nil, errors.New("invalid value for required argument 'Label'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomField
	err := ctx.RegisterResource("rootly:index/customField:CustomField", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomField gets an existing CustomField resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomField(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomFieldState, opts ...pulumi.ResourceOption) (*CustomField, error) {
	var resource CustomField
	err := ctx.ReadResource("rootly:index/customField:CustomField", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomField resources.
type customFieldState struct {
	// The default value for text field kinds.
	Default *string `pulumi:"default"`
	// The description of the custom_field
	Description *string `pulumi:"description"`
	Enabled     *bool   `pulumi:"enabled"`
	// The kind of the custom_field
	Kind *string `pulumi:"kind"`
	// The name of the custom_field
	Label *string `pulumi:"label"`
	// The position of the custom_field
	Position  *int     `pulumi:"position"`
	Requireds []string `pulumi:"requireds"`
	Showns    []string `pulumi:"showns"`
	// The slug of the custom_field
	Slug *string `pulumi:"slug"`
}

type CustomFieldState struct {
	// The default value for text field kinds.
	Default pulumi.StringPtrInput
	// The description of the custom_field
	Description pulumi.StringPtrInput
	Enabled     pulumi.BoolPtrInput
	// The kind of the custom_field
	Kind pulumi.StringPtrInput
	// The name of the custom_field
	Label pulumi.StringPtrInput
	// The position of the custom_field
	Position  pulumi.IntPtrInput
	Requireds pulumi.StringArrayInput
	Showns    pulumi.StringArrayInput
	// The slug of the custom_field
	Slug pulumi.StringPtrInput
}

func (CustomFieldState) ElementType() reflect.Type {
	return reflect.TypeOf((*customFieldState)(nil)).Elem()
}

type customFieldArgs struct {
	// The default value for text field kinds.
	Default *string `pulumi:"default"`
	// The description of the custom_field
	Description *string `pulumi:"description"`
	Enabled     *bool   `pulumi:"enabled"`
	// The kind of the custom_field
	Kind *string `pulumi:"kind"`
	// The name of the custom_field
	Label string `pulumi:"label"`
	// The position of the custom_field
	Position  *int     `pulumi:"position"`
	Requireds []string `pulumi:"requireds"`
	Showns    []string `pulumi:"showns"`
	// The slug of the custom_field
	Slug *string `pulumi:"slug"`
}

// The set of arguments for constructing a CustomField resource.
type CustomFieldArgs struct {
	// The default value for text field kinds.
	Default pulumi.StringPtrInput
	// The description of the custom_field
	Description pulumi.StringPtrInput
	Enabled     pulumi.BoolPtrInput
	// The kind of the custom_field
	Kind pulumi.StringPtrInput
	// The name of the custom_field
	Label pulumi.StringInput
	// The position of the custom_field
	Position  pulumi.IntPtrInput
	Requireds pulumi.StringArrayInput
	Showns    pulumi.StringArrayInput
	// The slug of the custom_field
	Slug pulumi.StringPtrInput
}

func (CustomFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customFieldArgs)(nil)).Elem()
}

type CustomFieldInput interface {
	pulumi.Input

	ToCustomFieldOutput() CustomFieldOutput
	ToCustomFieldOutputWithContext(ctx context.Context) CustomFieldOutput
}

func (*CustomField) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomField)(nil)).Elem()
}

func (i *CustomField) ToCustomFieldOutput() CustomFieldOutput {
	return i.ToCustomFieldOutputWithContext(context.Background())
}

func (i *CustomField) ToCustomFieldOutputWithContext(ctx context.Context) CustomFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomFieldOutput)
}

// CustomFieldArrayInput is an input type that accepts CustomFieldArray and CustomFieldArrayOutput values.
// You can construct a concrete instance of `CustomFieldArrayInput` via:
//
//	CustomFieldArray{ CustomFieldArgs{...} }
type CustomFieldArrayInput interface {
	pulumi.Input

	ToCustomFieldArrayOutput() CustomFieldArrayOutput
	ToCustomFieldArrayOutputWithContext(context.Context) CustomFieldArrayOutput
}

type CustomFieldArray []CustomFieldInput

func (CustomFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomField)(nil)).Elem()
}

func (i CustomFieldArray) ToCustomFieldArrayOutput() CustomFieldArrayOutput {
	return i.ToCustomFieldArrayOutputWithContext(context.Background())
}

func (i CustomFieldArray) ToCustomFieldArrayOutputWithContext(ctx context.Context) CustomFieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomFieldArrayOutput)
}

// CustomFieldMapInput is an input type that accepts CustomFieldMap and CustomFieldMapOutput values.
// You can construct a concrete instance of `CustomFieldMapInput` via:
//
//	CustomFieldMap{ "key": CustomFieldArgs{...} }
type CustomFieldMapInput interface {
	pulumi.Input

	ToCustomFieldMapOutput() CustomFieldMapOutput
	ToCustomFieldMapOutputWithContext(context.Context) CustomFieldMapOutput
}

type CustomFieldMap map[string]CustomFieldInput

func (CustomFieldMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomField)(nil)).Elem()
}

func (i CustomFieldMap) ToCustomFieldMapOutput() CustomFieldMapOutput {
	return i.ToCustomFieldMapOutputWithContext(context.Background())
}

func (i CustomFieldMap) ToCustomFieldMapOutputWithContext(ctx context.Context) CustomFieldMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomFieldMapOutput)
}

type CustomFieldOutput struct{ *pulumi.OutputState }

func (CustomFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomField)(nil)).Elem()
}

func (o CustomFieldOutput) ToCustomFieldOutput() CustomFieldOutput {
	return o
}

func (o CustomFieldOutput) ToCustomFieldOutputWithContext(ctx context.Context) CustomFieldOutput {
	return o
}

// The default value for text field kinds.
func (o CustomFieldOutput) Default() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomField) pulumi.StringOutput { return v.Default }).(pulumi.StringOutput)
}

// The description of the custom_field
func (o CustomFieldOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomField) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o CustomFieldOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomField) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The kind of the custom_field
func (o CustomFieldOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomField) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the custom_field
func (o CustomFieldOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomField) pulumi.StringOutput { return v.Label }).(pulumi.StringOutput)
}

// The position of the custom_field
func (o CustomFieldOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *CustomField) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

func (o CustomFieldOutput) Requireds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomField) pulumi.StringArrayOutput { return v.Requireds }).(pulumi.StringArrayOutput)
}

func (o CustomFieldOutput) Showns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomField) pulumi.StringArrayOutput { return v.Showns }).(pulumi.StringArrayOutput)
}

// The slug of the custom_field
func (o CustomFieldOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomField) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

type CustomFieldArrayOutput struct{ *pulumi.OutputState }

func (CustomFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomField)(nil)).Elem()
}

func (o CustomFieldArrayOutput) ToCustomFieldArrayOutput() CustomFieldArrayOutput {
	return o
}

func (o CustomFieldArrayOutput) ToCustomFieldArrayOutputWithContext(ctx context.Context) CustomFieldArrayOutput {
	return o
}

func (o CustomFieldArrayOutput) Index(i pulumi.IntInput) CustomFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomField {
		return vs[0].([]*CustomField)[vs[1].(int)]
	}).(CustomFieldOutput)
}

type CustomFieldMapOutput struct{ *pulumi.OutputState }

func (CustomFieldMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomField)(nil)).Elem()
}

func (o CustomFieldMapOutput) ToCustomFieldMapOutput() CustomFieldMapOutput {
	return o
}

func (o CustomFieldMapOutput) ToCustomFieldMapOutputWithContext(ctx context.Context) CustomFieldMapOutput {
	return o
}

func (o CustomFieldMapOutput) MapIndex(k pulumi.StringInput) CustomFieldOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomField {
		return vs[0].(map[string]*CustomField)[vs[1].(string)]
	}).(CustomFieldOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomFieldInput)(nil)).Elem(), &CustomField{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomFieldArrayInput)(nil)).Elem(), CustomFieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomFieldMapInput)(nil)).Elem(), CustomFieldMap{})
	pulumi.RegisterOutputType(CustomFieldOutput{})
	pulumi.RegisterOutputType(CustomFieldArrayOutput{})
	pulumi.RegisterOutputType(CustomFieldMapOutput{})
}
