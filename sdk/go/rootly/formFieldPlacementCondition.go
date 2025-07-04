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

type FormFieldPlacementCondition struct {
	pulumi.CustomResourceState

	// The condition comparison.. Value must be one of `equal`.
	Comparison pulumi.StringPtrOutput `pulumi:"comparison"`
	// The resource or attribute the condition applies.. Value must be one of `placement`, `required`.
	Conditioned pulumi.StringPtrOutput `pulumi:"conditioned"`
	// The condition field.
	FormFieldId pulumi.StringOutput `pulumi:"formFieldId"`
	// The form field placement this condition applies.
	FormFieldPlacementId pulumi.StringOutput `pulumi:"formFieldPlacementId"`
	// The condition position.
	Position pulumi.IntOutput `pulumi:"position"`
	// The values for comparison.
	Values pulumi.StringArrayOutput `pulumi:"values"`
}

// NewFormFieldPlacementCondition registers a new resource with the given unique name, arguments, and options.
func NewFormFieldPlacementCondition(ctx *pulumi.Context,
	name string, args *FormFieldPlacementConditionArgs, opts ...pulumi.ResourceOption) (*FormFieldPlacementCondition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FormFieldId == nil {
		return nil, errors.New("invalid value for required argument 'FormFieldId'")
	}
	if args.Values == nil {
		return nil, errors.New("invalid value for required argument 'Values'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FormFieldPlacementCondition
	err := ctx.RegisterResource("rootly:index/formFieldPlacementCondition:FormFieldPlacementCondition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFormFieldPlacementCondition gets an existing FormFieldPlacementCondition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFormFieldPlacementCondition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FormFieldPlacementConditionState, opts ...pulumi.ResourceOption) (*FormFieldPlacementCondition, error) {
	var resource FormFieldPlacementCondition
	err := ctx.ReadResource("rootly:index/formFieldPlacementCondition:FormFieldPlacementCondition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FormFieldPlacementCondition resources.
type formFieldPlacementConditionState struct {
	// The condition comparison.. Value must be one of `equal`.
	Comparison *string `pulumi:"comparison"`
	// The resource or attribute the condition applies.. Value must be one of `placement`, `required`.
	Conditioned *string `pulumi:"conditioned"`
	// The condition field.
	FormFieldId *string `pulumi:"formFieldId"`
	// The form field placement this condition applies.
	FormFieldPlacementId *string `pulumi:"formFieldPlacementId"`
	// The condition position.
	Position *int `pulumi:"position"`
	// The values for comparison.
	Values []string `pulumi:"values"`
}

type FormFieldPlacementConditionState struct {
	// The condition comparison.. Value must be one of `equal`.
	Comparison pulumi.StringPtrInput
	// The resource or attribute the condition applies.. Value must be one of `placement`, `required`.
	Conditioned pulumi.StringPtrInput
	// The condition field.
	FormFieldId pulumi.StringPtrInput
	// The form field placement this condition applies.
	FormFieldPlacementId pulumi.StringPtrInput
	// The condition position.
	Position pulumi.IntPtrInput
	// The values for comparison.
	Values pulumi.StringArrayInput
}

func (FormFieldPlacementConditionState) ElementType() reflect.Type {
	return reflect.TypeOf((*formFieldPlacementConditionState)(nil)).Elem()
}

type formFieldPlacementConditionArgs struct {
	// The condition comparison.. Value must be one of `equal`.
	Comparison *string `pulumi:"comparison"`
	// The resource or attribute the condition applies.. Value must be one of `placement`, `required`.
	Conditioned *string `pulumi:"conditioned"`
	// The condition field.
	FormFieldId string `pulumi:"formFieldId"`
	// The form field placement this condition applies.
	FormFieldPlacementId *string `pulumi:"formFieldPlacementId"`
	// The condition position.
	Position *int `pulumi:"position"`
	// The values for comparison.
	Values []string `pulumi:"values"`
}

// The set of arguments for constructing a FormFieldPlacementCondition resource.
type FormFieldPlacementConditionArgs struct {
	// The condition comparison.. Value must be one of `equal`.
	Comparison pulumi.StringPtrInput
	// The resource or attribute the condition applies.. Value must be one of `placement`, `required`.
	Conditioned pulumi.StringPtrInput
	// The condition field.
	FormFieldId pulumi.StringInput
	// The form field placement this condition applies.
	FormFieldPlacementId pulumi.StringPtrInput
	// The condition position.
	Position pulumi.IntPtrInput
	// The values for comparison.
	Values pulumi.StringArrayInput
}

func (FormFieldPlacementConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*formFieldPlacementConditionArgs)(nil)).Elem()
}

type FormFieldPlacementConditionInput interface {
	pulumi.Input

	ToFormFieldPlacementConditionOutput() FormFieldPlacementConditionOutput
	ToFormFieldPlacementConditionOutputWithContext(ctx context.Context) FormFieldPlacementConditionOutput
}

func (*FormFieldPlacementCondition) ElementType() reflect.Type {
	return reflect.TypeOf((**FormFieldPlacementCondition)(nil)).Elem()
}

func (i *FormFieldPlacementCondition) ToFormFieldPlacementConditionOutput() FormFieldPlacementConditionOutput {
	return i.ToFormFieldPlacementConditionOutputWithContext(context.Background())
}

func (i *FormFieldPlacementCondition) ToFormFieldPlacementConditionOutputWithContext(ctx context.Context) FormFieldPlacementConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormFieldPlacementConditionOutput)
}

// FormFieldPlacementConditionArrayInput is an input type that accepts FormFieldPlacementConditionArray and FormFieldPlacementConditionArrayOutput values.
// You can construct a concrete instance of `FormFieldPlacementConditionArrayInput` via:
//
//	FormFieldPlacementConditionArray{ FormFieldPlacementConditionArgs{...} }
type FormFieldPlacementConditionArrayInput interface {
	pulumi.Input

	ToFormFieldPlacementConditionArrayOutput() FormFieldPlacementConditionArrayOutput
	ToFormFieldPlacementConditionArrayOutputWithContext(context.Context) FormFieldPlacementConditionArrayOutput
}

type FormFieldPlacementConditionArray []FormFieldPlacementConditionInput

func (FormFieldPlacementConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FormFieldPlacementCondition)(nil)).Elem()
}

func (i FormFieldPlacementConditionArray) ToFormFieldPlacementConditionArrayOutput() FormFieldPlacementConditionArrayOutput {
	return i.ToFormFieldPlacementConditionArrayOutputWithContext(context.Background())
}

func (i FormFieldPlacementConditionArray) ToFormFieldPlacementConditionArrayOutputWithContext(ctx context.Context) FormFieldPlacementConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormFieldPlacementConditionArrayOutput)
}

// FormFieldPlacementConditionMapInput is an input type that accepts FormFieldPlacementConditionMap and FormFieldPlacementConditionMapOutput values.
// You can construct a concrete instance of `FormFieldPlacementConditionMapInput` via:
//
//	FormFieldPlacementConditionMap{ "key": FormFieldPlacementConditionArgs{...} }
type FormFieldPlacementConditionMapInput interface {
	pulumi.Input

	ToFormFieldPlacementConditionMapOutput() FormFieldPlacementConditionMapOutput
	ToFormFieldPlacementConditionMapOutputWithContext(context.Context) FormFieldPlacementConditionMapOutput
}

type FormFieldPlacementConditionMap map[string]FormFieldPlacementConditionInput

func (FormFieldPlacementConditionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FormFieldPlacementCondition)(nil)).Elem()
}

func (i FormFieldPlacementConditionMap) ToFormFieldPlacementConditionMapOutput() FormFieldPlacementConditionMapOutput {
	return i.ToFormFieldPlacementConditionMapOutputWithContext(context.Background())
}

func (i FormFieldPlacementConditionMap) ToFormFieldPlacementConditionMapOutputWithContext(ctx context.Context) FormFieldPlacementConditionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormFieldPlacementConditionMapOutput)
}

type FormFieldPlacementConditionOutput struct{ *pulumi.OutputState }

func (FormFieldPlacementConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FormFieldPlacementCondition)(nil)).Elem()
}

func (o FormFieldPlacementConditionOutput) ToFormFieldPlacementConditionOutput() FormFieldPlacementConditionOutput {
	return o
}

func (o FormFieldPlacementConditionOutput) ToFormFieldPlacementConditionOutputWithContext(ctx context.Context) FormFieldPlacementConditionOutput {
	return o
}

// The condition comparison.. Value must be one of `equal`.
func (o FormFieldPlacementConditionOutput) Comparison() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FormFieldPlacementCondition) pulumi.StringPtrOutput { return v.Comparison }).(pulumi.StringPtrOutput)
}

// The resource or attribute the condition applies.. Value must be one of `placement`, `required`.
func (o FormFieldPlacementConditionOutput) Conditioned() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FormFieldPlacementCondition) pulumi.StringPtrOutput { return v.Conditioned }).(pulumi.StringPtrOutput)
}

// The condition field.
func (o FormFieldPlacementConditionOutput) FormFieldId() pulumi.StringOutput {
	return o.ApplyT(func(v *FormFieldPlacementCondition) pulumi.StringOutput { return v.FormFieldId }).(pulumi.StringOutput)
}

// The form field placement this condition applies.
func (o FormFieldPlacementConditionOutput) FormFieldPlacementId() pulumi.StringOutput {
	return o.ApplyT(func(v *FormFieldPlacementCondition) pulumi.StringOutput { return v.FormFieldPlacementId }).(pulumi.StringOutput)
}

// The condition position.
func (o FormFieldPlacementConditionOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *FormFieldPlacementCondition) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// The values for comparison.
func (o FormFieldPlacementConditionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FormFieldPlacementCondition) pulumi.StringArrayOutput { return v.Values }).(pulumi.StringArrayOutput)
}

type FormFieldPlacementConditionArrayOutput struct{ *pulumi.OutputState }

func (FormFieldPlacementConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FormFieldPlacementCondition)(nil)).Elem()
}

func (o FormFieldPlacementConditionArrayOutput) ToFormFieldPlacementConditionArrayOutput() FormFieldPlacementConditionArrayOutput {
	return o
}

func (o FormFieldPlacementConditionArrayOutput) ToFormFieldPlacementConditionArrayOutputWithContext(ctx context.Context) FormFieldPlacementConditionArrayOutput {
	return o
}

func (o FormFieldPlacementConditionArrayOutput) Index(i pulumi.IntInput) FormFieldPlacementConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FormFieldPlacementCondition {
		return vs[0].([]*FormFieldPlacementCondition)[vs[1].(int)]
	}).(FormFieldPlacementConditionOutput)
}

type FormFieldPlacementConditionMapOutput struct{ *pulumi.OutputState }

func (FormFieldPlacementConditionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FormFieldPlacementCondition)(nil)).Elem()
}

func (o FormFieldPlacementConditionMapOutput) ToFormFieldPlacementConditionMapOutput() FormFieldPlacementConditionMapOutput {
	return o
}

func (o FormFieldPlacementConditionMapOutput) ToFormFieldPlacementConditionMapOutputWithContext(ctx context.Context) FormFieldPlacementConditionMapOutput {
	return o
}

func (o FormFieldPlacementConditionMapOutput) MapIndex(k pulumi.StringInput) FormFieldPlacementConditionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FormFieldPlacementCondition {
		return vs[0].(map[string]*FormFieldPlacementCondition)[vs[1].(string)]
	}).(FormFieldPlacementConditionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FormFieldPlacementConditionInput)(nil)).Elem(), &FormFieldPlacementCondition{})
	pulumi.RegisterInputType(reflect.TypeOf((*FormFieldPlacementConditionArrayInput)(nil)).Elem(), FormFieldPlacementConditionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FormFieldPlacementConditionMapInput)(nil)).Elem(), FormFieldPlacementConditionMap{})
	pulumi.RegisterOutputType(FormFieldPlacementConditionOutput{})
	pulumi.RegisterOutputType(FormFieldPlacementConditionArrayOutput{})
	pulumi.RegisterOutputType(FormFieldPlacementConditionMapOutput{})
}
