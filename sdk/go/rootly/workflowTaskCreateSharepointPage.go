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

// Manages workflow createSharepointPage task.
type WorkflowTaskCreateSharepointPage struct {
	pulumi.CustomResourceState

	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Name of the workflow task
	Name pulumi.StringOutput `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntOutput `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrOutput `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateSharepointPageTaskParamsOutput `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflowTaskCreateSharepointPage registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTaskCreateSharepointPage(ctx *pulumi.Context,
	name string, args *WorkflowTaskCreateSharepointPageArgs, opts ...pulumi.ResourceOption) (*WorkflowTaskCreateSharepointPage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TaskParams == nil {
		return nil, errors.New("invalid value for required argument 'TaskParams'")
	}
	if args.WorkflowId == nil {
		return nil, errors.New("invalid value for required argument 'WorkflowId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WorkflowTaskCreateSharepointPage
	err := ctx.RegisterResource("rootly:index/workflowTaskCreateSharepointPage:WorkflowTaskCreateSharepointPage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTaskCreateSharepointPage gets an existing WorkflowTaskCreateSharepointPage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTaskCreateSharepointPage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTaskCreateSharepointPageState, opts ...pulumi.ResourceOption) (*WorkflowTaskCreateSharepointPage, error) {
	var resource WorkflowTaskCreateSharepointPage
	err := ctx.ReadResource("rootly:index/workflowTaskCreateSharepointPage:WorkflowTaskCreateSharepointPage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTaskCreateSharepointPage resources.
type workflowTaskCreateSharepointPageState struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams *WorkflowTaskCreateSharepointPageTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId *string `pulumi:"workflowId"`
}

type WorkflowTaskCreateSharepointPageState struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateSharepointPageTaskParamsPtrInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringPtrInput
}

func (WorkflowTaskCreateSharepointPageState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskCreateSharepointPageState)(nil)).Elem()
}

type workflowTaskCreateSharepointPageArgs struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateSharepointPageTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowTaskCreateSharepointPage resource.
type WorkflowTaskCreateSharepointPageArgs struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateSharepointPageTaskParamsInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringInput
}

func (WorkflowTaskCreateSharepointPageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskCreateSharepointPageArgs)(nil)).Elem()
}

type WorkflowTaskCreateSharepointPageInput interface {
	pulumi.Input

	ToWorkflowTaskCreateSharepointPageOutput() WorkflowTaskCreateSharepointPageOutput
	ToWorkflowTaskCreateSharepointPageOutputWithContext(ctx context.Context) WorkflowTaskCreateSharepointPageOutput
}

func (*WorkflowTaskCreateSharepointPage) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskCreateSharepointPage)(nil)).Elem()
}

func (i *WorkflowTaskCreateSharepointPage) ToWorkflowTaskCreateSharepointPageOutput() WorkflowTaskCreateSharepointPageOutput {
	return i.ToWorkflowTaskCreateSharepointPageOutputWithContext(context.Background())
}

func (i *WorkflowTaskCreateSharepointPage) ToWorkflowTaskCreateSharepointPageOutputWithContext(ctx context.Context) WorkflowTaskCreateSharepointPageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateSharepointPageOutput)
}

// WorkflowTaskCreateSharepointPageArrayInput is an input type that accepts WorkflowTaskCreateSharepointPageArray and WorkflowTaskCreateSharepointPageArrayOutput values.
// You can construct a concrete instance of `WorkflowTaskCreateSharepointPageArrayInput` via:
//
//	WorkflowTaskCreateSharepointPageArray{ WorkflowTaskCreateSharepointPageArgs{...} }
type WorkflowTaskCreateSharepointPageArrayInput interface {
	pulumi.Input

	ToWorkflowTaskCreateSharepointPageArrayOutput() WorkflowTaskCreateSharepointPageArrayOutput
	ToWorkflowTaskCreateSharepointPageArrayOutputWithContext(context.Context) WorkflowTaskCreateSharepointPageArrayOutput
}

type WorkflowTaskCreateSharepointPageArray []WorkflowTaskCreateSharepointPageInput

func (WorkflowTaskCreateSharepointPageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskCreateSharepointPage)(nil)).Elem()
}

func (i WorkflowTaskCreateSharepointPageArray) ToWorkflowTaskCreateSharepointPageArrayOutput() WorkflowTaskCreateSharepointPageArrayOutput {
	return i.ToWorkflowTaskCreateSharepointPageArrayOutputWithContext(context.Background())
}

func (i WorkflowTaskCreateSharepointPageArray) ToWorkflowTaskCreateSharepointPageArrayOutputWithContext(ctx context.Context) WorkflowTaskCreateSharepointPageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateSharepointPageArrayOutput)
}

// WorkflowTaskCreateSharepointPageMapInput is an input type that accepts WorkflowTaskCreateSharepointPageMap and WorkflowTaskCreateSharepointPageMapOutput values.
// You can construct a concrete instance of `WorkflowTaskCreateSharepointPageMapInput` via:
//
//	WorkflowTaskCreateSharepointPageMap{ "key": WorkflowTaskCreateSharepointPageArgs{...} }
type WorkflowTaskCreateSharepointPageMapInput interface {
	pulumi.Input

	ToWorkflowTaskCreateSharepointPageMapOutput() WorkflowTaskCreateSharepointPageMapOutput
	ToWorkflowTaskCreateSharepointPageMapOutputWithContext(context.Context) WorkflowTaskCreateSharepointPageMapOutput
}

type WorkflowTaskCreateSharepointPageMap map[string]WorkflowTaskCreateSharepointPageInput

func (WorkflowTaskCreateSharepointPageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskCreateSharepointPage)(nil)).Elem()
}

func (i WorkflowTaskCreateSharepointPageMap) ToWorkflowTaskCreateSharepointPageMapOutput() WorkflowTaskCreateSharepointPageMapOutput {
	return i.ToWorkflowTaskCreateSharepointPageMapOutputWithContext(context.Background())
}

func (i WorkflowTaskCreateSharepointPageMap) ToWorkflowTaskCreateSharepointPageMapOutputWithContext(ctx context.Context) WorkflowTaskCreateSharepointPageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateSharepointPageMapOutput)
}

type WorkflowTaskCreateSharepointPageOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateSharepointPageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskCreateSharepointPage)(nil)).Elem()
}

func (o WorkflowTaskCreateSharepointPageOutput) ToWorkflowTaskCreateSharepointPageOutput() WorkflowTaskCreateSharepointPageOutput {
	return o
}

func (o WorkflowTaskCreateSharepointPageOutput) ToWorkflowTaskCreateSharepointPageOutputWithContext(ctx context.Context) WorkflowTaskCreateSharepointPageOutput {
	return o
}

// Enable/disable this workflow task
func (o WorkflowTaskCreateSharepointPageOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateSharepointPage) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Name of the workflow task
func (o WorkflowTaskCreateSharepointPageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateSharepointPage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of the workflow task (1 being top of list)
func (o WorkflowTaskCreateSharepointPageOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateSharepointPage) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Skip workflow task if any failures
func (o WorkflowTaskCreateSharepointPageOutput) SkipOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateSharepointPage) pulumi.BoolPtrOutput { return v.SkipOnFailure }).(pulumi.BoolPtrOutput)
}

// The parameters for this workflow task.
func (o WorkflowTaskCreateSharepointPageOutput) TaskParams() WorkflowTaskCreateSharepointPageTaskParamsOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateSharepointPage) WorkflowTaskCreateSharepointPageTaskParamsOutput {
		return v.TaskParams
	}).(WorkflowTaskCreateSharepointPageTaskParamsOutput)
}

// The ID of the parent workflow
func (o WorkflowTaskCreateSharepointPageOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateSharepointPage) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

type WorkflowTaskCreateSharepointPageArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateSharepointPageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskCreateSharepointPage)(nil)).Elem()
}

func (o WorkflowTaskCreateSharepointPageArrayOutput) ToWorkflowTaskCreateSharepointPageArrayOutput() WorkflowTaskCreateSharepointPageArrayOutput {
	return o
}

func (o WorkflowTaskCreateSharepointPageArrayOutput) ToWorkflowTaskCreateSharepointPageArrayOutputWithContext(ctx context.Context) WorkflowTaskCreateSharepointPageArrayOutput {
	return o
}

func (o WorkflowTaskCreateSharepointPageArrayOutput) Index(i pulumi.IntInput) WorkflowTaskCreateSharepointPageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTaskCreateSharepointPage {
		return vs[0].([]*WorkflowTaskCreateSharepointPage)[vs[1].(int)]
	}).(WorkflowTaskCreateSharepointPageOutput)
}

type WorkflowTaskCreateSharepointPageMapOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateSharepointPageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskCreateSharepointPage)(nil)).Elem()
}

func (o WorkflowTaskCreateSharepointPageMapOutput) ToWorkflowTaskCreateSharepointPageMapOutput() WorkflowTaskCreateSharepointPageMapOutput {
	return o
}

func (o WorkflowTaskCreateSharepointPageMapOutput) ToWorkflowTaskCreateSharepointPageMapOutputWithContext(ctx context.Context) WorkflowTaskCreateSharepointPageMapOutput {
	return o
}

func (o WorkflowTaskCreateSharepointPageMapOutput) MapIndex(k pulumi.StringInput) WorkflowTaskCreateSharepointPageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTaskCreateSharepointPage {
		return vs[0].(map[string]*WorkflowTaskCreateSharepointPage)[vs[1].(string)]
	}).(WorkflowTaskCreateSharepointPageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateSharepointPageInput)(nil)).Elem(), &WorkflowTaskCreateSharepointPage{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateSharepointPageArrayInput)(nil)).Elem(), WorkflowTaskCreateSharepointPageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateSharepointPageMapInput)(nil)).Elem(), WorkflowTaskCreateSharepointPageMap{})
	pulumi.RegisterOutputType(WorkflowTaskCreateSharepointPageOutput{})
	pulumi.RegisterOutputType(WorkflowTaskCreateSharepointPageArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTaskCreateSharepointPageMapOutput{})
}
