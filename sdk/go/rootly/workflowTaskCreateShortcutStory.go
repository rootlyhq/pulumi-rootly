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

// Manages workflow createShortcutStory task.
type WorkflowTaskCreateShortcutStory struct {
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
	TaskParams WorkflowTaskCreateShortcutStoryTaskParamsOutput `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflowTaskCreateShortcutStory registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTaskCreateShortcutStory(ctx *pulumi.Context,
	name string, args *WorkflowTaskCreateShortcutStoryArgs, opts ...pulumi.ResourceOption) (*WorkflowTaskCreateShortcutStory, error) {
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
	var resource WorkflowTaskCreateShortcutStory
	err := ctx.RegisterResource("rootly:index/workflowTaskCreateShortcutStory:WorkflowTaskCreateShortcutStory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTaskCreateShortcutStory gets an existing WorkflowTaskCreateShortcutStory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTaskCreateShortcutStory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTaskCreateShortcutStoryState, opts ...pulumi.ResourceOption) (*WorkflowTaskCreateShortcutStory, error) {
	var resource WorkflowTaskCreateShortcutStory
	err := ctx.ReadResource("rootly:index/workflowTaskCreateShortcutStory:WorkflowTaskCreateShortcutStory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTaskCreateShortcutStory resources.
type workflowTaskCreateShortcutStoryState struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams *WorkflowTaskCreateShortcutStoryTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId *string `pulumi:"workflowId"`
}

type WorkflowTaskCreateShortcutStoryState struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateShortcutStoryTaskParamsPtrInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringPtrInput
}

func (WorkflowTaskCreateShortcutStoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskCreateShortcutStoryState)(nil)).Elem()
}

type workflowTaskCreateShortcutStoryArgs struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateShortcutStoryTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowTaskCreateShortcutStory resource.
type WorkflowTaskCreateShortcutStoryArgs struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateShortcutStoryTaskParamsInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringInput
}

func (WorkflowTaskCreateShortcutStoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskCreateShortcutStoryArgs)(nil)).Elem()
}

type WorkflowTaskCreateShortcutStoryInput interface {
	pulumi.Input

	ToWorkflowTaskCreateShortcutStoryOutput() WorkflowTaskCreateShortcutStoryOutput
	ToWorkflowTaskCreateShortcutStoryOutputWithContext(ctx context.Context) WorkflowTaskCreateShortcutStoryOutput
}

func (*WorkflowTaskCreateShortcutStory) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskCreateShortcutStory)(nil)).Elem()
}

func (i *WorkflowTaskCreateShortcutStory) ToWorkflowTaskCreateShortcutStoryOutput() WorkflowTaskCreateShortcutStoryOutput {
	return i.ToWorkflowTaskCreateShortcutStoryOutputWithContext(context.Background())
}

func (i *WorkflowTaskCreateShortcutStory) ToWorkflowTaskCreateShortcutStoryOutputWithContext(ctx context.Context) WorkflowTaskCreateShortcutStoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateShortcutStoryOutput)
}

// WorkflowTaskCreateShortcutStoryArrayInput is an input type that accepts WorkflowTaskCreateShortcutStoryArray and WorkflowTaskCreateShortcutStoryArrayOutput values.
// You can construct a concrete instance of `WorkflowTaskCreateShortcutStoryArrayInput` via:
//
//	WorkflowTaskCreateShortcutStoryArray{ WorkflowTaskCreateShortcutStoryArgs{...} }
type WorkflowTaskCreateShortcutStoryArrayInput interface {
	pulumi.Input

	ToWorkflowTaskCreateShortcutStoryArrayOutput() WorkflowTaskCreateShortcutStoryArrayOutput
	ToWorkflowTaskCreateShortcutStoryArrayOutputWithContext(context.Context) WorkflowTaskCreateShortcutStoryArrayOutput
}

type WorkflowTaskCreateShortcutStoryArray []WorkflowTaskCreateShortcutStoryInput

func (WorkflowTaskCreateShortcutStoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskCreateShortcutStory)(nil)).Elem()
}

func (i WorkflowTaskCreateShortcutStoryArray) ToWorkflowTaskCreateShortcutStoryArrayOutput() WorkflowTaskCreateShortcutStoryArrayOutput {
	return i.ToWorkflowTaskCreateShortcutStoryArrayOutputWithContext(context.Background())
}

func (i WorkflowTaskCreateShortcutStoryArray) ToWorkflowTaskCreateShortcutStoryArrayOutputWithContext(ctx context.Context) WorkflowTaskCreateShortcutStoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateShortcutStoryArrayOutput)
}

// WorkflowTaskCreateShortcutStoryMapInput is an input type that accepts WorkflowTaskCreateShortcutStoryMap and WorkflowTaskCreateShortcutStoryMapOutput values.
// You can construct a concrete instance of `WorkflowTaskCreateShortcutStoryMapInput` via:
//
//	WorkflowTaskCreateShortcutStoryMap{ "key": WorkflowTaskCreateShortcutStoryArgs{...} }
type WorkflowTaskCreateShortcutStoryMapInput interface {
	pulumi.Input

	ToWorkflowTaskCreateShortcutStoryMapOutput() WorkflowTaskCreateShortcutStoryMapOutput
	ToWorkflowTaskCreateShortcutStoryMapOutputWithContext(context.Context) WorkflowTaskCreateShortcutStoryMapOutput
}

type WorkflowTaskCreateShortcutStoryMap map[string]WorkflowTaskCreateShortcutStoryInput

func (WorkflowTaskCreateShortcutStoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskCreateShortcutStory)(nil)).Elem()
}

func (i WorkflowTaskCreateShortcutStoryMap) ToWorkflowTaskCreateShortcutStoryMapOutput() WorkflowTaskCreateShortcutStoryMapOutput {
	return i.ToWorkflowTaskCreateShortcutStoryMapOutputWithContext(context.Background())
}

func (i WorkflowTaskCreateShortcutStoryMap) ToWorkflowTaskCreateShortcutStoryMapOutputWithContext(ctx context.Context) WorkflowTaskCreateShortcutStoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateShortcutStoryMapOutput)
}

type WorkflowTaskCreateShortcutStoryOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateShortcutStoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskCreateShortcutStory)(nil)).Elem()
}

func (o WorkflowTaskCreateShortcutStoryOutput) ToWorkflowTaskCreateShortcutStoryOutput() WorkflowTaskCreateShortcutStoryOutput {
	return o
}

func (o WorkflowTaskCreateShortcutStoryOutput) ToWorkflowTaskCreateShortcutStoryOutputWithContext(ctx context.Context) WorkflowTaskCreateShortcutStoryOutput {
	return o
}

// Enable/disable this workflow task
func (o WorkflowTaskCreateShortcutStoryOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateShortcutStory) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Name of the workflow task
func (o WorkflowTaskCreateShortcutStoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateShortcutStory) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of the workflow task (1 being top of list)
func (o WorkflowTaskCreateShortcutStoryOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateShortcutStory) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Skip workflow task if any failures
func (o WorkflowTaskCreateShortcutStoryOutput) SkipOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateShortcutStory) pulumi.BoolPtrOutput { return v.SkipOnFailure }).(pulumi.BoolPtrOutput)
}

// The parameters for this workflow task.
func (o WorkflowTaskCreateShortcutStoryOutput) TaskParams() WorkflowTaskCreateShortcutStoryTaskParamsOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateShortcutStory) WorkflowTaskCreateShortcutStoryTaskParamsOutput {
		return v.TaskParams
	}).(WorkflowTaskCreateShortcutStoryTaskParamsOutput)
}

// The ID of the parent workflow
func (o WorkflowTaskCreateShortcutStoryOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateShortcutStory) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

type WorkflowTaskCreateShortcutStoryArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateShortcutStoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskCreateShortcutStory)(nil)).Elem()
}

func (o WorkflowTaskCreateShortcutStoryArrayOutput) ToWorkflowTaskCreateShortcutStoryArrayOutput() WorkflowTaskCreateShortcutStoryArrayOutput {
	return o
}

func (o WorkflowTaskCreateShortcutStoryArrayOutput) ToWorkflowTaskCreateShortcutStoryArrayOutputWithContext(ctx context.Context) WorkflowTaskCreateShortcutStoryArrayOutput {
	return o
}

func (o WorkflowTaskCreateShortcutStoryArrayOutput) Index(i pulumi.IntInput) WorkflowTaskCreateShortcutStoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTaskCreateShortcutStory {
		return vs[0].([]*WorkflowTaskCreateShortcutStory)[vs[1].(int)]
	}).(WorkflowTaskCreateShortcutStoryOutput)
}

type WorkflowTaskCreateShortcutStoryMapOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateShortcutStoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskCreateShortcutStory)(nil)).Elem()
}

func (o WorkflowTaskCreateShortcutStoryMapOutput) ToWorkflowTaskCreateShortcutStoryMapOutput() WorkflowTaskCreateShortcutStoryMapOutput {
	return o
}

func (o WorkflowTaskCreateShortcutStoryMapOutput) ToWorkflowTaskCreateShortcutStoryMapOutputWithContext(ctx context.Context) WorkflowTaskCreateShortcutStoryMapOutput {
	return o
}

func (o WorkflowTaskCreateShortcutStoryMapOutput) MapIndex(k pulumi.StringInput) WorkflowTaskCreateShortcutStoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTaskCreateShortcutStory {
		return vs[0].(map[string]*WorkflowTaskCreateShortcutStory)[vs[1].(string)]
	}).(WorkflowTaskCreateShortcutStoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateShortcutStoryInput)(nil)).Elem(), &WorkflowTaskCreateShortcutStory{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateShortcutStoryArrayInput)(nil)).Elem(), WorkflowTaskCreateShortcutStoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateShortcutStoryMapInput)(nil)).Elem(), WorkflowTaskCreateShortcutStoryMap{})
	pulumi.RegisterOutputType(WorkflowTaskCreateShortcutStoryOutput{})
	pulumi.RegisterOutputType(WorkflowTaskCreateShortcutStoryArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTaskCreateShortcutStoryMapOutput{})
}
