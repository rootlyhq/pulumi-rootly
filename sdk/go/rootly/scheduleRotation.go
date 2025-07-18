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

type ScheduleRotation struct {
	pulumi.CustomResourceState

	// Schedule rotation active all week?. Value must be one of true or false
	ActiveAllWeek pulumi.BoolOutput `pulumi:"activeAllWeek"`
	// Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
	ActiveDays pulumi.StringArrayOutput `pulumi:"activeDays"`
	// Schedule rotation's active times
	ActiveTimeAttributes ScheduleRotationActiveTimeAttributeArrayOutput `pulumi:"activeTimeAttributes"`
	ActiveTimeType       pulumi.StringOutput                            `pulumi:"activeTimeType"`
	// The name of the schedule rotation
	Name pulumi.StringOutput `pulumi:"name"`
	// Position of the schedule rotation
	Position pulumi.IntOutput `pulumi:"position"`
	// The ID of parent schedule
	ScheduleId                     pulumi.StringOutput    `pulumi:"scheduleId"`
	ScheduleRotationableAttributes pulumi.StringMapOutput `pulumi:"scheduleRotationableAttributes"`
	// Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
	ScheduleRotationableType pulumi.StringPtrOutput `pulumi:"scheduleRotationableType"`
	// A valid IANA time zone name.
	TimeZone pulumi.StringPtrOutput `pulumi:"timeZone"`
}

// NewScheduleRotation registers a new resource with the given unique name, arguments, and options.
func NewScheduleRotation(ctx *pulumi.Context,
	name string, args *ScheduleRotationArgs, opts ...pulumi.ResourceOption) (*ScheduleRotation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScheduleRotationableAttributes == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleRotationableAttributes'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScheduleRotation
	err := ctx.RegisterResource("rootly:index/scheduleRotation:ScheduleRotation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduleRotation gets an existing ScheduleRotation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduleRotation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleRotationState, opts ...pulumi.ResourceOption) (*ScheduleRotation, error) {
	var resource ScheduleRotation
	err := ctx.ReadResource("rootly:index/scheduleRotation:ScheduleRotation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduleRotation resources.
type scheduleRotationState struct {
	// Schedule rotation active all week?. Value must be one of true or false
	ActiveAllWeek *bool `pulumi:"activeAllWeek"`
	// Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
	ActiveDays []string `pulumi:"activeDays"`
	// Schedule rotation's active times
	ActiveTimeAttributes []ScheduleRotationActiveTimeAttribute `pulumi:"activeTimeAttributes"`
	ActiveTimeType       *string                               `pulumi:"activeTimeType"`
	// The name of the schedule rotation
	Name *string `pulumi:"name"`
	// Position of the schedule rotation
	Position *int `pulumi:"position"`
	// The ID of parent schedule
	ScheduleId                     *string           `pulumi:"scheduleId"`
	ScheduleRotationableAttributes map[string]string `pulumi:"scheduleRotationableAttributes"`
	// Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
	ScheduleRotationableType *string `pulumi:"scheduleRotationableType"`
	// A valid IANA time zone name.
	TimeZone *string `pulumi:"timeZone"`
}

type ScheduleRotationState struct {
	// Schedule rotation active all week?. Value must be one of true or false
	ActiveAllWeek pulumi.BoolPtrInput
	// Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
	ActiveDays pulumi.StringArrayInput
	// Schedule rotation's active times
	ActiveTimeAttributes ScheduleRotationActiveTimeAttributeArrayInput
	ActiveTimeType       pulumi.StringPtrInput
	// The name of the schedule rotation
	Name pulumi.StringPtrInput
	// Position of the schedule rotation
	Position pulumi.IntPtrInput
	// The ID of parent schedule
	ScheduleId                     pulumi.StringPtrInput
	ScheduleRotationableAttributes pulumi.StringMapInput
	// Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
	ScheduleRotationableType pulumi.StringPtrInput
	// A valid IANA time zone name.
	TimeZone pulumi.StringPtrInput
}

func (ScheduleRotationState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleRotationState)(nil)).Elem()
}

type scheduleRotationArgs struct {
	// Schedule rotation active all week?. Value must be one of true or false
	ActiveAllWeek *bool `pulumi:"activeAllWeek"`
	// Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
	ActiveDays []string `pulumi:"activeDays"`
	// Schedule rotation's active times
	ActiveTimeAttributes []ScheduleRotationActiveTimeAttribute `pulumi:"activeTimeAttributes"`
	ActiveTimeType       *string                               `pulumi:"activeTimeType"`
	// The name of the schedule rotation
	Name *string `pulumi:"name"`
	// Position of the schedule rotation
	Position *int `pulumi:"position"`
	// The ID of parent schedule
	ScheduleId                     *string           `pulumi:"scheduleId"`
	ScheduleRotationableAttributes map[string]string `pulumi:"scheduleRotationableAttributes"`
	// Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
	ScheduleRotationableType *string `pulumi:"scheduleRotationableType"`
	// A valid IANA time zone name.
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a ScheduleRotation resource.
type ScheduleRotationArgs struct {
	// Schedule rotation active all week?. Value must be one of true or false
	ActiveAllWeek pulumi.BoolPtrInput
	// Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
	ActiveDays pulumi.StringArrayInput
	// Schedule rotation's active times
	ActiveTimeAttributes ScheduleRotationActiveTimeAttributeArrayInput
	ActiveTimeType       pulumi.StringPtrInput
	// The name of the schedule rotation
	Name pulumi.StringPtrInput
	// Position of the schedule rotation
	Position pulumi.IntPtrInput
	// The ID of parent schedule
	ScheduleId                     pulumi.StringPtrInput
	ScheduleRotationableAttributes pulumi.StringMapInput
	// Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
	ScheduleRotationableType pulumi.StringPtrInput
	// A valid IANA time zone name.
	TimeZone pulumi.StringPtrInput
}

func (ScheduleRotationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleRotationArgs)(nil)).Elem()
}

type ScheduleRotationInput interface {
	pulumi.Input

	ToScheduleRotationOutput() ScheduleRotationOutput
	ToScheduleRotationOutputWithContext(ctx context.Context) ScheduleRotationOutput
}

func (*ScheduleRotation) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleRotation)(nil)).Elem()
}

func (i *ScheduleRotation) ToScheduleRotationOutput() ScheduleRotationOutput {
	return i.ToScheduleRotationOutputWithContext(context.Background())
}

func (i *ScheduleRotation) ToScheduleRotationOutputWithContext(ctx context.Context) ScheduleRotationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleRotationOutput)
}

// ScheduleRotationArrayInput is an input type that accepts ScheduleRotationArray and ScheduleRotationArrayOutput values.
// You can construct a concrete instance of `ScheduleRotationArrayInput` via:
//
//	ScheduleRotationArray{ ScheduleRotationArgs{...} }
type ScheduleRotationArrayInput interface {
	pulumi.Input

	ToScheduleRotationArrayOutput() ScheduleRotationArrayOutput
	ToScheduleRotationArrayOutputWithContext(context.Context) ScheduleRotationArrayOutput
}

type ScheduleRotationArray []ScheduleRotationInput

func (ScheduleRotationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduleRotation)(nil)).Elem()
}

func (i ScheduleRotationArray) ToScheduleRotationArrayOutput() ScheduleRotationArrayOutput {
	return i.ToScheduleRotationArrayOutputWithContext(context.Background())
}

func (i ScheduleRotationArray) ToScheduleRotationArrayOutputWithContext(ctx context.Context) ScheduleRotationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleRotationArrayOutput)
}

// ScheduleRotationMapInput is an input type that accepts ScheduleRotationMap and ScheduleRotationMapOutput values.
// You can construct a concrete instance of `ScheduleRotationMapInput` via:
//
//	ScheduleRotationMap{ "key": ScheduleRotationArgs{...} }
type ScheduleRotationMapInput interface {
	pulumi.Input

	ToScheduleRotationMapOutput() ScheduleRotationMapOutput
	ToScheduleRotationMapOutputWithContext(context.Context) ScheduleRotationMapOutput
}

type ScheduleRotationMap map[string]ScheduleRotationInput

func (ScheduleRotationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduleRotation)(nil)).Elem()
}

func (i ScheduleRotationMap) ToScheduleRotationMapOutput() ScheduleRotationMapOutput {
	return i.ToScheduleRotationMapOutputWithContext(context.Background())
}

func (i ScheduleRotationMap) ToScheduleRotationMapOutputWithContext(ctx context.Context) ScheduleRotationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleRotationMapOutput)
}

type ScheduleRotationOutput struct{ *pulumi.OutputState }

func (ScheduleRotationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleRotation)(nil)).Elem()
}

func (o ScheduleRotationOutput) ToScheduleRotationOutput() ScheduleRotationOutput {
	return o
}

func (o ScheduleRotationOutput) ToScheduleRotationOutputWithContext(ctx context.Context) ScheduleRotationOutput {
	return o
}

// Schedule rotation active all week?. Value must be one of true or false
func (o ScheduleRotationOutput) ActiveAllWeek() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScheduleRotation) pulumi.BoolOutput { return v.ActiveAllWeek }).(pulumi.BoolOutput)
}

// Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
func (o ScheduleRotationOutput) ActiveDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScheduleRotation) pulumi.StringArrayOutput { return v.ActiveDays }).(pulumi.StringArrayOutput)
}

// Schedule rotation's active times
func (o ScheduleRotationOutput) ActiveTimeAttributes() ScheduleRotationActiveTimeAttributeArrayOutput {
	return o.ApplyT(func(v *ScheduleRotation) ScheduleRotationActiveTimeAttributeArrayOutput {
		return v.ActiveTimeAttributes
	}).(ScheduleRotationActiveTimeAttributeArrayOutput)
}

func (o ScheduleRotationOutput) ActiveTimeType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduleRotation) pulumi.StringOutput { return v.ActiveTimeType }).(pulumi.StringOutput)
}

// The name of the schedule rotation
func (o ScheduleRotationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduleRotation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Position of the schedule rotation
func (o ScheduleRotationOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduleRotation) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// The ID of parent schedule
func (o ScheduleRotationOutput) ScheduleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduleRotation) pulumi.StringOutput { return v.ScheduleId }).(pulumi.StringOutput)
}

func (o ScheduleRotationOutput) ScheduleRotationableAttributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ScheduleRotation) pulumi.StringMapOutput { return v.ScheduleRotationableAttributes }).(pulumi.StringMapOutput)
}

// Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
func (o ScheduleRotationOutput) ScheduleRotationableType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleRotation) pulumi.StringPtrOutput { return v.ScheduleRotationableType }).(pulumi.StringPtrOutput)
}

// A valid IANA time zone name.
func (o ScheduleRotationOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleRotation) pulumi.StringPtrOutput { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type ScheduleRotationArrayOutput struct{ *pulumi.OutputState }

func (ScheduleRotationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduleRotation)(nil)).Elem()
}

func (o ScheduleRotationArrayOutput) ToScheduleRotationArrayOutput() ScheduleRotationArrayOutput {
	return o
}

func (o ScheduleRotationArrayOutput) ToScheduleRotationArrayOutputWithContext(ctx context.Context) ScheduleRotationArrayOutput {
	return o
}

func (o ScheduleRotationArrayOutput) Index(i pulumi.IntInput) ScheduleRotationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScheduleRotation {
		return vs[0].([]*ScheduleRotation)[vs[1].(int)]
	}).(ScheduleRotationOutput)
}

type ScheduleRotationMapOutput struct{ *pulumi.OutputState }

func (ScheduleRotationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduleRotation)(nil)).Elem()
}

func (o ScheduleRotationMapOutput) ToScheduleRotationMapOutput() ScheduleRotationMapOutput {
	return o
}

func (o ScheduleRotationMapOutput) ToScheduleRotationMapOutputWithContext(ctx context.Context) ScheduleRotationMapOutput {
	return o
}

func (o ScheduleRotationMapOutput) MapIndex(k pulumi.StringInput) ScheduleRotationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScheduleRotation {
		return vs[0].(map[string]*ScheduleRotation)[vs[1].(string)]
	}).(ScheduleRotationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleRotationInput)(nil)).Elem(), &ScheduleRotation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleRotationArrayInput)(nil)).Elem(), ScheduleRotationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleRotationMapInput)(nil)).Elem(), ScheduleRotationMap{})
	pulumi.RegisterOutputType(ScheduleRotationOutput{})
	pulumi.RegisterOutputType(ScheduleRotationArrayOutput{})
	pulumi.RegisterOutputType(ScheduleRotationMapOutput{})
}
