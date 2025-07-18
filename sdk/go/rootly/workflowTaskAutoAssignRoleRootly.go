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

// Manages workflow autoAssignRoleRootly task.
type WorkflowTaskAutoAssignRoleRootly struct {
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
	TaskParams WorkflowTaskAutoAssignRoleRootlyTaskParamsOutput `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflowTaskAutoAssignRoleRootly registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTaskAutoAssignRoleRootly(ctx *pulumi.Context,
	name string, args *WorkflowTaskAutoAssignRoleRootlyArgs, opts ...pulumi.ResourceOption) (*WorkflowTaskAutoAssignRoleRootly, error) {
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
	var resource WorkflowTaskAutoAssignRoleRootly
	err := ctx.RegisterResource("rootly:index/workflowTaskAutoAssignRoleRootly:WorkflowTaskAutoAssignRoleRootly", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTaskAutoAssignRoleRootly gets an existing WorkflowTaskAutoAssignRoleRootly resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTaskAutoAssignRoleRootly(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTaskAutoAssignRoleRootlyState, opts ...pulumi.ResourceOption) (*WorkflowTaskAutoAssignRoleRootly, error) {
	var resource WorkflowTaskAutoAssignRoleRootly
	err := ctx.ReadResource("rootly:index/workflowTaskAutoAssignRoleRootly:WorkflowTaskAutoAssignRoleRootly", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTaskAutoAssignRoleRootly resources.
type workflowTaskAutoAssignRoleRootlyState struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams *WorkflowTaskAutoAssignRoleRootlyTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId *string `pulumi:"workflowId"`
}

type WorkflowTaskAutoAssignRoleRootlyState struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskAutoAssignRoleRootlyTaskParamsPtrInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringPtrInput
}

func (WorkflowTaskAutoAssignRoleRootlyState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskAutoAssignRoleRootlyState)(nil)).Elem()
}

type workflowTaskAutoAssignRoleRootlyArgs struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskAutoAssignRoleRootlyTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowTaskAutoAssignRoleRootly resource.
type WorkflowTaskAutoAssignRoleRootlyArgs struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskAutoAssignRoleRootlyTaskParamsInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringInput
}

func (WorkflowTaskAutoAssignRoleRootlyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskAutoAssignRoleRootlyArgs)(nil)).Elem()
}

type WorkflowTaskAutoAssignRoleRootlyInput interface {
	pulumi.Input

	ToWorkflowTaskAutoAssignRoleRootlyOutput() WorkflowTaskAutoAssignRoleRootlyOutput
	ToWorkflowTaskAutoAssignRoleRootlyOutputWithContext(ctx context.Context) WorkflowTaskAutoAssignRoleRootlyOutput
}

func (*WorkflowTaskAutoAssignRoleRootly) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskAutoAssignRoleRootly)(nil)).Elem()
}

func (i *WorkflowTaskAutoAssignRoleRootly) ToWorkflowTaskAutoAssignRoleRootlyOutput() WorkflowTaskAutoAssignRoleRootlyOutput {
	return i.ToWorkflowTaskAutoAssignRoleRootlyOutputWithContext(context.Background())
}

func (i *WorkflowTaskAutoAssignRoleRootly) ToWorkflowTaskAutoAssignRoleRootlyOutputWithContext(ctx context.Context) WorkflowTaskAutoAssignRoleRootlyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskAutoAssignRoleRootlyOutput)
}

// WorkflowTaskAutoAssignRoleRootlyArrayInput is an input type that accepts WorkflowTaskAutoAssignRoleRootlyArray and WorkflowTaskAutoAssignRoleRootlyArrayOutput values.
// You can construct a concrete instance of `WorkflowTaskAutoAssignRoleRootlyArrayInput` via:
//
//	WorkflowTaskAutoAssignRoleRootlyArray{ WorkflowTaskAutoAssignRoleRootlyArgs{...} }
type WorkflowTaskAutoAssignRoleRootlyArrayInput interface {
	pulumi.Input

	ToWorkflowTaskAutoAssignRoleRootlyArrayOutput() WorkflowTaskAutoAssignRoleRootlyArrayOutput
	ToWorkflowTaskAutoAssignRoleRootlyArrayOutputWithContext(context.Context) WorkflowTaskAutoAssignRoleRootlyArrayOutput
}

type WorkflowTaskAutoAssignRoleRootlyArray []WorkflowTaskAutoAssignRoleRootlyInput

func (WorkflowTaskAutoAssignRoleRootlyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskAutoAssignRoleRootly)(nil)).Elem()
}

func (i WorkflowTaskAutoAssignRoleRootlyArray) ToWorkflowTaskAutoAssignRoleRootlyArrayOutput() WorkflowTaskAutoAssignRoleRootlyArrayOutput {
	return i.ToWorkflowTaskAutoAssignRoleRootlyArrayOutputWithContext(context.Background())
}

func (i WorkflowTaskAutoAssignRoleRootlyArray) ToWorkflowTaskAutoAssignRoleRootlyArrayOutputWithContext(ctx context.Context) WorkflowTaskAutoAssignRoleRootlyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskAutoAssignRoleRootlyArrayOutput)
}

// WorkflowTaskAutoAssignRoleRootlyMapInput is an input type that accepts WorkflowTaskAutoAssignRoleRootlyMap and WorkflowTaskAutoAssignRoleRootlyMapOutput values.
// You can construct a concrete instance of `WorkflowTaskAutoAssignRoleRootlyMapInput` via:
//
//	WorkflowTaskAutoAssignRoleRootlyMap{ "key": WorkflowTaskAutoAssignRoleRootlyArgs{...} }
type WorkflowTaskAutoAssignRoleRootlyMapInput interface {
	pulumi.Input

	ToWorkflowTaskAutoAssignRoleRootlyMapOutput() WorkflowTaskAutoAssignRoleRootlyMapOutput
	ToWorkflowTaskAutoAssignRoleRootlyMapOutputWithContext(context.Context) WorkflowTaskAutoAssignRoleRootlyMapOutput
}

type WorkflowTaskAutoAssignRoleRootlyMap map[string]WorkflowTaskAutoAssignRoleRootlyInput

func (WorkflowTaskAutoAssignRoleRootlyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskAutoAssignRoleRootly)(nil)).Elem()
}

func (i WorkflowTaskAutoAssignRoleRootlyMap) ToWorkflowTaskAutoAssignRoleRootlyMapOutput() WorkflowTaskAutoAssignRoleRootlyMapOutput {
	return i.ToWorkflowTaskAutoAssignRoleRootlyMapOutputWithContext(context.Background())
}

func (i WorkflowTaskAutoAssignRoleRootlyMap) ToWorkflowTaskAutoAssignRoleRootlyMapOutputWithContext(ctx context.Context) WorkflowTaskAutoAssignRoleRootlyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskAutoAssignRoleRootlyMapOutput)
}

type WorkflowTaskAutoAssignRoleRootlyOutput struct{ *pulumi.OutputState }

func (WorkflowTaskAutoAssignRoleRootlyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskAutoAssignRoleRootly)(nil)).Elem()
}

func (o WorkflowTaskAutoAssignRoleRootlyOutput) ToWorkflowTaskAutoAssignRoleRootlyOutput() WorkflowTaskAutoAssignRoleRootlyOutput {
	return o
}

func (o WorkflowTaskAutoAssignRoleRootlyOutput) ToWorkflowTaskAutoAssignRoleRootlyOutputWithContext(ctx context.Context) WorkflowTaskAutoAssignRoleRootlyOutput {
	return o
}

// Enable/disable this workflow task
func (o WorkflowTaskAutoAssignRoleRootlyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskAutoAssignRoleRootly) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Name of the workflow task
func (o WorkflowTaskAutoAssignRoleRootlyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskAutoAssignRoleRootly) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of the workflow task (1 being top of list)
func (o WorkflowTaskAutoAssignRoleRootlyOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTaskAutoAssignRoleRootly) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Skip workflow task if any failures
func (o WorkflowTaskAutoAssignRoleRootlyOutput) SkipOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskAutoAssignRoleRootly) pulumi.BoolPtrOutput { return v.SkipOnFailure }).(pulumi.BoolPtrOutput)
}

// The parameters for this workflow task.
func (o WorkflowTaskAutoAssignRoleRootlyOutput) TaskParams() WorkflowTaskAutoAssignRoleRootlyTaskParamsOutput {
	return o.ApplyT(func(v *WorkflowTaskAutoAssignRoleRootly) WorkflowTaskAutoAssignRoleRootlyTaskParamsOutput {
		return v.TaskParams
	}).(WorkflowTaskAutoAssignRoleRootlyTaskParamsOutput)
}

// The ID of the parent workflow
func (o WorkflowTaskAutoAssignRoleRootlyOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskAutoAssignRoleRootly) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

type WorkflowTaskAutoAssignRoleRootlyArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTaskAutoAssignRoleRootlyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskAutoAssignRoleRootly)(nil)).Elem()
}

func (o WorkflowTaskAutoAssignRoleRootlyArrayOutput) ToWorkflowTaskAutoAssignRoleRootlyArrayOutput() WorkflowTaskAutoAssignRoleRootlyArrayOutput {
	return o
}

func (o WorkflowTaskAutoAssignRoleRootlyArrayOutput) ToWorkflowTaskAutoAssignRoleRootlyArrayOutputWithContext(ctx context.Context) WorkflowTaskAutoAssignRoleRootlyArrayOutput {
	return o
}

func (o WorkflowTaskAutoAssignRoleRootlyArrayOutput) Index(i pulumi.IntInput) WorkflowTaskAutoAssignRoleRootlyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTaskAutoAssignRoleRootly {
		return vs[0].([]*WorkflowTaskAutoAssignRoleRootly)[vs[1].(int)]
	}).(WorkflowTaskAutoAssignRoleRootlyOutput)
}

type WorkflowTaskAutoAssignRoleRootlyMapOutput struct{ *pulumi.OutputState }

func (WorkflowTaskAutoAssignRoleRootlyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskAutoAssignRoleRootly)(nil)).Elem()
}

func (o WorkflowTaskAutoAssignRoleRootlyMapOutput) ToWorkflowTaskAutoAssignRoleRootlyMapOutput() WorkflowTaskAutoAssignRoleRootlyMapOutput {
	return o
}

func (o WorkflowTaskAutoAssignRoleRootlyMapOutput) ToWorkflowTaskAutoAssignRoleRootlyMapOutputWithContext(ctx context.Context) WorkflowTaskAutoAssignRoleRootlyMapOutput {
	return o
}

func (o WorkflowTaskAutoAssignRoleRootlyMapOutput) MapIndex(k pulumi.StringInput) WorkflowTaskAutoAssignRoleRootlyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTaskAutoAssignRoleRootly {
		return vs[0].(map[string]*WorkflowTaskAutoAssignRoleRootly)[vs[1].(string)]
	}).(WorkflowTaskAutoAssignRoleRootlyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskAutoAssignRoleRootlyInput)(nil)).Elem(), &WorkflowTaskAutoAssignRoleRootly{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskAutoAssignRoleRootlyArrayInput)(nil)).Elem(), WorkflowTaskAutoAssignRoleRootlyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskAutoAssignRoleRootlyMapInput)(nil)).Elem(), WorkflowTaskAutoAssignRoleRootlyMap{})
	pulumi.RegisterOutputType(WorkflowTaskAutoAssignRoleRootlyOutput{})
	pulumi.RegisterOutputType(WorkflowTaskAutoAssignRoleRootlyArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTaskAutoAssignRoleRootlyMapOutput{})
}
