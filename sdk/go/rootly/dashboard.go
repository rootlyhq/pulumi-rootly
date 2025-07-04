// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rootly

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rootlyhq/pulumi-rootly/sdk/go/rootly/internal"
)

// Manages dashboards.
//
// ## Example Usage
type Dashboard struct {
	pulumi.CustomResourceState

	// The name of the dashboard
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the dashboard is owned by organization or user.
	Owner pulumi.StringPtrOutput `pulumi:"owner"`
	// Whether the dashboard is public
	Public pulumi.BoolPtrOutput `pulumi:"public"`
	// The slug of the dashboard
	Slug pulumi.StringOutput `pulumi:"slug"`
	// The userId of the dashboard, for dashboards with owner == user
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		args = &DashboardArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Dashboard
	err := ctx.RegisterResource("rootly:index/dashboard:Dashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboard gets an existing Dashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardState, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	var resource Dashboard
	err := ctx.ReadResource("rootly:index/dashboard:Dashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dashboard resources.
type dashboardState struct {
	// The name of the dashboard
	Name *string `pulumi:"name"`
	// Whether the dashboard is owned by organization or user.
	Owner *string `pulumi:"owner"`
	// Whether the dashboard is public
	Public *bool `pulumi:"public"`
	// The slug of the dashboard
	Slug *string `pulumi:"slug"`
	// The userId of the dashboard, for dashboards with owner == user
	UserId *int `pulumi:"userId"`
}

type DashboardState struct {
	// The name of the dashboard
	Name pulumi.StringPtrInput
	// Whether the dashboard is owned by organization or user.
	Owner pulumi.StringPtrInput
	// Whether the dashboard is public
	Public pulumi.BoolPtrInput
	// The slug of the dashboard
	Slug pulumi.StringPtrInput
	// The userId of the dashboard, for dashboards with owner == user
	UserId pulumi.IntPtrInput
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	// The name of the dashboard
	Name *string `pulumi:"name"`
	// Whether the dashboard is owned by organization or user.
	Owner *string `pulumi:"owner"`
	// Whether the dashboard is public
	Public *bool `pulumi:"public"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// The name of the dashboard
	Name pulumi.StringPtrInput
	// Whether the dashboard is owned by organization or user.
	Owner pulumi.StringPtrInput
	// Whether the dashboard is public
	Public pulumi.BoolPtrInput
}

func (DashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardArgs)(nil)).Elem()
}

type DashboardInput interface {
	pulumi.Input

	ToDashboardOutput() DashboardOutput
	ToDashboardOutputWithContext(ctx context.Context) DashboardOutput
}

func (*Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (i *Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

// DashboardArrayInput is an input type that accepts DashboardArray and DashboardArrayOutput values.
// You can construct a concrete instance of `DashboardArrayInput` via:
//
//	DashboardArray{ DashboardArgs{...} }
type DashboardArrayInput interface {
	pulumi.Input

	ToDashboardArrayOutput() DashboardArrayOutput
	ToDashboardArrayOutputWithContext(context.Context) DashboardArrayOutput
}

type DashboardArray []DashboardInput

func (DashboardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dashboard)(nil)).Elem()
}

func (i DashboardArray) ToDashboardArrayOutput() DashboardArrayOutput {
	return i.ToDashboardArrayOutputWithContext(context.Background())
}

func (i DashboardArray) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardArrayOutput)
}

// DashboardMapInput is an input type that accepts DashboardMap and DashboardMapOutput values.
// You can construct a concrete instance of `DashboardMapInput` via:
//
//	DashboardMap{ "key": DashboardArgs{...} }
type DashboardMapInput interface {
	pulumi.Input

	ToDashboardMapOutput() DashboardMapOutput
	ToDashboardMapOutputWithContext(context.Context) DashboardMapOutput
}

type DashboardMap map[string]DashboardInput

func (DashboardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dashboard)(nil)).Elem()
}

func (i DashboardMap) ToDashboardMapOutput() DashboardMapOutput {
	return i.ToDashboardMapOutputWithContext(context.Background())
}

func (i DashboardMap) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardMapOutput)
}

type DashboardOutput struct{ *pulumi.OutputState }

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

// The name of the dashboard
func (o DashboardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether the dashboard is owned by organization or user.
func (o DashboardOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringPtrOutput { return v.Owner }).(pulumi.StringPtrOutput)
}

// Whether the dashboard is public
func (o DashboardOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.BoolPtrOutput { return v.Public }).(pulumi.BoolPtrOutput)
}

// The slug of the dashboard
func (o DashboardOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// The userId of the dashboard, for dashboards with owner == user
func (o DashboardOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type DashboardArrayOutput struct{ *pulumi.OutputState }

func (DashboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dashboard)(nil)).Elem()
}

func (o DashboardArrayOutput) ToDashboardArrayOutput() DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) Index(i pulumi.IntInput) DashboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dashboard {
		return vs[0].([]*Dashboard)[vs[1].(int)]
	}).(DashboardOutput)
}

type DashboardMapOutput struct{ *pulumi.OutputState }

func (DashboardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dashboard)(nil)).Elem()
}

func (o DashboardMapOutput) ToDashboardMapOutput() DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) MapIndex(k pulumi.StringInput) DashboardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dashboard {
		return vs[0].(map[string]*Dashboard)[vs[1].(string)]
	}).(DashboardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardInput)(nil)).Elem(), &Dashboard{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardArrayInput)(nil)).Elem(), DashboardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardMapInput)(nil)).Elem(), DashboardMap{})
	pulumi.RegisterOutputType(DashboardOutput{})
	pulumi.RegisterOutputType(DashboardArrayOutput{})
	pulumi.RegisterOutputType(DashboardMapOutput{})
}
