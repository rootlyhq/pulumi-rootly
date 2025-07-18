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

// Manages workflow createGithubIssue task.
type WorkflowTaskCreateGithubIssue struct {
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
	TaskParams WorkflowTaskCreateGithubIssueTaskParamsOutput `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflowTaskCreateGithubIssue registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTaskCreateGithubIssue(ctx *pulumi.Context,
	name string, args *WorkflowTaskCreateGithubIssueArgs, opts ...pulumi.ResourceOption) (*WorkflowTaskCreateGithubIssue, error) {
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
	var resource WorkflowTaskCreateGithubIssue
	err := ctx.RegisterResource("rootly:index/workflowTaskCreateGithubIssue:WorkflowTaskCreateGithubIssue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTaskCreateGithubIssue gets an existing WorkflowTaskCreateGithubIssue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTaskCreateGithubIssue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTaskCreateGithubIssueState, opts ...pulumi.ResourceOption) (*WorkflowTaskCreateGithubIssue, error) {
	var resource WorkflowTaskCreateGithubIssue
	err := ctx.ReadResource("rootly:index/workflowTaskCreateGithubIssue:WorkflowTaskCreateGithubIssue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTaskCreateGithubIssue resources.
type workflowTaskCreateGithubIssueState struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams *WorkflowTaskCreateGithubIssueTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId *string `pulumi:"workflowId"`
}

type WorkflowTaskCreateGithubIssueState struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateGithubIssueTaskParamsPtrInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringPtrInput
}

func (WorkflowTaskCreateGithubIssueState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskCreateGithubIssueState)(nil)).Elem()
}

type workflowTaskCreateGithubIssueArgs struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateGithubIssueTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowTaskCreateGithubIssue resource.
type WorkflowTaskCreateGithubIssueArgs struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateGithubIssueTaskParamsInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringInput
}

func (WorkflowTaskCreateGithubIssueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskCreateGithubIssueArgs)(nil)).Elem()
}

type WorkflowTaskCreateGithubIssueInput interface {
	pulumi.Input

	ToWorkflowTaskCreateGithubIssueOutput() WorkflowTaskCreateGithubIssueOutput
	ToWorkflowTaskCreateGithubIssueOutputWithContext(ctx context.Context) WorkflowTaskCreateGithubIssueOutput
}

func (*WorkflowTaskCreateGithubIssue) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskCreateGithubIssue)(nil)).Elem()
}

func (i *WorkflowTaskCreateGithubIssue) ToWorkflowTaskCreateGithubIssueOutput() WorkflowTaskCreateGithubIssueOutput {
	return i.ToWorkflowTaskCreateGithubIssueOutputWithContext(context.Background())
}

func (i *WorkflowTaskCreateGithubIssue) ToWorkflowTaskCreateGithubIssueOutputWithContext(ctx context.Context) WorkflowTaskCreateGithubIssueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateGithubIssueOutput)
}

// WorkflowTaskCreateGithubIssueArrayInput is an input type that accepts WorkflowTaskCreateGithubIssueArray and WorkflowTaskCreateGithubIssueArrayOutput values.
// You can construct a concrete instance of `WorkflowTaskCreateGithubIssueArrayInput` via:
//
//	WorkflowTaskCreateGithubIssueArray{ WorkflowTaskCreateGithubIssueArgs{...} }
type WorkflowTaskCreateGithubIssueArrayInput interface {
	pulumi.Input

	ToWorkflowTaskCreateGithubIssueArrayOutput() WorkflowTaskCreateGithubIssueArrayOutput
	ToWorkflowTaskCreateGithubIssueArrayOutputWithContext(context.Context) WorkflowTaskCreateGithubIssueArrayOutput
}

type WorkflowTaskCreateGithubIssueArray []WorkflowTaskCreateGithubIssueInput

func (WorkflowTaskCreateGithubIssueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskCreateGithubIssue)(nil)).Elem()
}

func (i WorkflowTaskCreateGithubIssueArray) ToWorkflowTaskCreateGithubIssueArrayOutput() WorkflowTaskCreateGithubIssueArrayOutput {
	return i.ToWorkflowTaskCreateGithubIssueArrayOutputWithContext(context.Background())
}

func (i WorkflowTaskCreateGithubIssueArray) ToWorkflowTaskCreateGithubIssueArrayOutputWithContext(ctx context.Context) WorkflowTaskCreateGithubIssueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateGithubIssueArrayOutput)
}

// WorkflowTaskCreateGithubIssueMapInput is an input type that accepts WorkflowTaskCreateGithubIssueMap and WorkflowTaskCreateGithubIssueMapOutput values.
// You can construct a concrete instance of `WorkflowTaskCreateGithubIssueMapInput` via:
//
//	WorkflowTaskCreateGithubIssueMap{ "key": WorkflowTaskCreateGithubIssueArgs{...} }
type WorkflowTaskCreateGithubIssueMapInput interface {
	pulumi.Input

	ToWorkflowTaskCreateGithubIssueMapOutput() WorkflowTaskCreateGithubIssueMapOutput
	ToWorkflowTaskCreateGithubIssueMapOutputWithContext(context.Context) WorkflowTaskCreateGithubIssueMapOutput
}

type WorkflowTaskCreateGithubIssueMap map[string]WorkflowTaskCreateGithubIssueInput

func (WorkflowTaskCreateGithubIssueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskCreateGithubIssue)(nil)).Elem()
}

func (i WorkflowTaskCreateGithubIssueMap) ToWorkflowTaskCreateGithubIssueMapOutput() WorkflowTaskCreateGithubIssueMapOutput {
	return i.ToWorkflowTaskCreateGithubIssueMapOutputWithContext(context.Background())
}

func (i WorkflowTaskCreateGithubIssueMap) ToWorkflowTaskCreateGithubIssueMapOutputWithContext(ctx context.Context) WorkflowTaskCreateGithubIssueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateGithubIssueMapOutput)
}

type WorkflowTaskCreateGithubIssueOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateGithubIssueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskCreateGithubIssue)(nil)).Elem()
}

func (o WorkflowTaskCreateGithubIssueOutput) ToWorkflowTaskCreateGithubIssueOutput() WorkflowTaskCreateGithubIssueOutput {
	return o
}

func (o WorkflowTaskCreateGithubIssueOutput) ToWorkflowTaskCreateGithubIssueOutputWithContext(ctx context.Context) WorkflowTaskCreateGithubIssueOutput {
	return o
}

// Enable/disable this workflow task
func (o WorkflowTaskCreateGithubIssueOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateGithubIssue) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Name of the workflow task
func (o WorkflowTaskCreateGithubIssueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateGithubIssue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of the workflow task (1 being top of list)
func (o WorkflowTaskCreateGithubIssueOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateGithubIssue) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Skip workflow task if any failures
func (o WorkflowTaskCreateGithubIssueOutput) SkipOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateGithubIssue) pulumi.BoolPtrOutput { return v.SkipOnFailure }).(pulumi.BoolPtrOutput)
}

// The parameters for this workflow task.
func (o WorkflowTaskCreateGithubIssueOutput) TaskParams() WorkflowTaskCreateGithubIssueTaskParamsOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateGithubIssue) WorkflowTaskCreateGithubIssueTaskParamsOutput {
		return v.TaskParams
	}).(WorkflowTaskCreateGithubIssueTaskParamsOutput)
}

// The ID of the parent workflow
func (o WorkflowTaskCreateGithubIssueOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateGithubIssue) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

type WorkflowTaskCreateGithubIssueArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateGithubIssueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskCreateGithubIssue)(nil)).Elem()
}

func (o WorkflowTaskCreateGithubIssueArrayOutput) ToWorkflowTaskCreateGithubIssueArrayOutput() WorkflowTaskCreateGithubIssueArrayOutput {
	return o
}

func (o WorkflowTaskCreateGithubIssueArrayOutput) ToWorkflowTaskCreateGithubIssueArrayOutputWithContext(ctx context.Context) WorkflowTaskCreateGithubIssueArrayOutput {
	return o
}

func (o WorkflowTaskCreateGithubIssueArrayOutput) Index(i pulumi.IntInput) WorkflowTaskCreateGithubIssueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTaskCreateGithubIssue {
		return vs[0].([]*WorkflowTaskCreateGithubIssue)[vs[1].(int)]
	}).(WorkflowTaskCreateGithubIssueOutput)
}

type WorkflowTaskCreateGithubIssueMapOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateGithubIssueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskCreateGithubIssue)(nil)).Elem()
}

func (o WorkflowTaskCreateGithubIssueMapOutput) ToWorkflowTaskCreateGithubIssueMapOutput() WorkflowTaskCreateGithubIssueMapOutput {
	return o
}

func (o WorkflowTaskCreateGithubIssueMapOutput) ToWorkflowTaskCreateGithubIssueMapOutputWithContext(ctx context.Context) WorkflowTaskCreateGithubIssueMapOutput {
	return o
}

func (o WorkflowTaskCreateGithubIssueMapOutput) MapIndex(k pulumi.StringInput) WorkflowTaskCreateGithubIssueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTaskCreateGithubIssue {
		return vs[0].(map[string]*WorkflowTaskCreateGithubIssue)[vs[1].(string)]
	}).(WorkflowTaskCreateGithubIssueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateGithubIssueInput)(nil)).Elem(), &WorkflowTaskCreateGithubIssue{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateGithubIssueArrayInput)(nil)).Elem(), WorkflowTaskCreateGithubIssueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateGithubIssueMapInput)(nil)).Elem(), WorkflowTaskCreateGithubIssueMap{})
	pulumi.RegisterOutputType(WorkflowTaskCreateGithubIssueOutput{})
	pulumi.RegisterOutputType(WorkflowTaskCreateGithubIssueArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTaskCreateGithubIssueMapOutput{})
}
