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

// Manages workflow sendEmail task.
//
// ## Example Usage
type WorkflowTaskSendEmail struct {
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
	TaskParams WorkflowTaskSendEmailTaskParamsOutput `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflowTaskSendEmail registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTaskSendEmail(ctx *pulumi.Context,
	name string, args *WorkflowTaskSendEmailArgs, opts ...pulumi.ResourceOption) (*WorkflowTaskSendEmail, error) {
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
	var resource WorkflowTaskSendEmail
	err := ctx.RegisterResource("rootly:index/workflowTaskSendEmail:WorkflowTaskSendEmail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTaskSendEmail gets an existing WorkflowTaskSendEmail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTaskSendEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTaskSendEmailState, opts ...pulumi.ResourceOption) (*WorkflowTaskSendEmail, error) {
	var resource WorkflowTaskSendEmail
	err := ctx.ReadResource("rootly:index/workflowTaskSendEmail:WorkflowTaskSendEmail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTaskSendEmail resources.
type workflowTaskSendEmailState struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams *WorkflowTaskSendEmailTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId *string `pulumi:"workflowId"`
}

type WorkflowTaskSendEmailState struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskSendEmailTaskParamsPtrInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringPtrInput
}

func (WorkflowTaskSendEmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskSendEmailState)(nil)).Elem()
}

type workflowTaskSendEmailArgs struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskSendEmailTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowTaskSendEmail resource.
type WorkflowTaskSendEmailArgs struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskSendEmailTaskParamsInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringInput
}

func (WorkflowTaskSendEmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskSendEmailArgs)(nil)).Elem()
}

type WorkflowTaskSendEmailInput interface {
	pulumi.Input

	ToWorkflowTaskSendEmailOutput() WorkflowTaskSendEmailOutput
	ToWorkflowTaskSendEmailOutputWithContext(ctx context.Context) WorkflowTaskSendEmailOutput
}

func (*WorkflowTaskSendEmail) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskSendEmail)(nil)).Elem()
}

func (i *WorkflowTaskSendEmail) ToWorkflowTaskSendEmailOutput() WorkflowTaskSendEmailOutput {
	return i.ToWorkflowTaskSendEmailOutputWithContext(context.Background())
}

func (i *WorkflowTaskSendEmail) ToWorkflowTaskSendEmailOutputWithContext(ctx context.Context) WorkflowTaskSendEmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskSendEmailOutput)
}

// WorkflowTaskSendEmailArrayInput is an input type that accepts WorkflowTaskSendEmailArray and WorkflowTaskSendEmailArrayOutput values.
// You can construct a concrete instance of `WorkflowTaskSendEmailArrayInput` via:
//
//	WorkflowTaskSendEmailArray{ WorkflowTaskSendEmailArgs{...} }
type WorkflowTaskSendEmailArrayInput interface {
	pulumi.Input

	ToWorkflowTaskSendEmailArrayOutput() WorkflowTaskSendEmailArrayOutput
	ToWorkflowTaskSendEmailArrayOutputWithContext(context.Context) WorkflowTaskSendEmailArrayOutput
}

type WorkflowTaskSendEmailArray []WorkflowTaskSendEmailInput

func (WorkflowTaskSendEmailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskSendEmail)(nil)).Elem()
}

func (i WorkflowTaskSendEmailArray) ToWorkflowTaskSendEmailArrayOutput() WorkflowTaskSendEmailArrayOutput {
	return i.ToWorkflowTaskSendEmailArrayOutputWithContext(context.Background())
}

func (i WorkflowTaskSendEmailArray) ToWorkflowTaskSendEmailArrayOutputWithContext(ctx context.Context) WorkflowTaskSendEmailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskSendEmailArrayOutput)
}

// WorkflowTaskSendEmailMapInput is an input type that accepts WorkflowTaskSendEmailMap and WorkflowTaskSendEmailMapOutput values.
// You can construct a concrete instance of `WorkflowTaskSendEmailMapInput` via:
//
//	WorkflowTaskSendEmailMap{ "key": WorkflowTaskSendEmailArgs{...} }
type WorkflowTaskSendEmailMapInput interface {
	pulumi.Input

	ToWorkflowTaskSendEmailMapOutput() WorkflowTaskSendEmailMapOutput
	ToWorkflowTaskSendEmailMapOutputWithContext(context.Context) WorkflowTaskSendEmailMapOutput
}

type WorkflowTaskSendEmailMap map[string]WorkflowTaskSendEmailInput

func (WorkflowTaskSendEmailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskSendEmail)(nil)).Elem()
}

func (i WorkflowTaskSendEmailMap) ToWorkflowTaskSendEmailMapOutput() WorkflowTaskSendEmailMapOutput {
	return i.ToWorkflowTaskSendEmailMapOutputWithContext(context.Background())
}

func (i WorkflowTaskSendEmailMap) ToWorkflowTaskSendEmailMapOutputWithContext(ctx context.Context) WorkflowTaskSendEmailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskSendEmailMapOutput)
}

type WorkflowTaskSendEmailOutput struct{ *pulumi.OutputState }

func (WorkflowTaskSendEmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskSendEmail)(nil)).Elem()
}

func (o WorkflowTaskSendEmailOutput) ToWorkflowTaskSendEmailOutput() WorkflowTaskSendEmailOutput {
	return o
}

func (o WorkflowTaskSendEmailOutput) ToWorkflowTaskSendEmailOutputWithContext(ctx context.Context) WorkflowTaskSendEmailOutput {
	return o
}

// Enable/disable this workflow task
func (o WorkflowTaskSendEmailOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskSendEmail) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Name of the workflow task
func (o WorkflowTaskSendEmailOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskSendEmail) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of the workflow task (1 being top of list)
func (o WorkflowTaskSendEmailOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTaskSendEmail) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Skip workflow task if any failures
func (o WorkflowTaskSendEmailOutput) SkipOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskSendEmail) pulumi.BoolPtrOutput { return v.SkipOnFailure }).(pulumi.BoolPtrOutput)
}

// The parameters for this workflow task.
func (o WorkflowTaskSendEmailOutput) TaskParams() WorkflowTaskSendEmailTaskParamsOutput {
	return o.ApplyT(func(v *WorkflowTaskSendEmail) WorkflowTaskSendEmailTaskParamsOutput { return v.TaskParams }).(WorkflowTaskSendEmailTaskParamsOutput)
}

// The ID of the parent workflow
func (o WorkflowTaskSendEmailOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskSendEmail) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

type WorkflowTaskSendEmailArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTaskSendEmailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskSendEmail)(nil)).Elem()
}

func (o WorkflowTaskSendEmailArrayOutput) ToWorkflowTaskSendEmailArrayOutput() WorkflowTaskSendEmailArrayOutput {
	return o
}

func (o WorkflowTaskSendEmailArrayOutput) ToWorkflowTaskSendEmailArrayOutputWithContext(ctx context.Context) WorkflowTaskSendEmailArrayOutput {
	return o
}

func (o WorkflowTaskSendEmailArrayOutput) Index(i pulumi.IntInput) WorkflowTaskSendEmailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTaskSendEmail {
		return vs[0].([]*WorkflowTaskSendEmail)[vs[1].(int)]
	}).(WorkflowTaskSendEmailOutput)
}

type WorkflowTaskSendEmailMapOutput struct{ *pulumi.OutputState }

func (WorkflowTaskSendEmailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskSendEmail)(nil)).Elem()
}

func (o WorkflowTaskSendEmailMapOutput) ToWorkflowTaskSendEmailMapOutput() WorkflowTaskSendEmailMapOutput {
	return o
}

func (o WorkflowTaskSendEmailMapOutput) ToWorkflowTaskSendEmailMapOutputWithContext(ctx context.Context) WorkflowTaskSendEmailMapOutput {
	return o
}

func (o WorkflowTaskSendEmailMapOutput) MapIndex(k pulumi.StringInput) WorkflowTaskSendEmailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTaskSendEmail {
		return vs[0].(map[string]*WorkflowTaskSendEmail)[vs[1].(string)]
	}).(WorkflowTaskSendEmailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskSendEmailInput)(nil)).Elem(), &WorkflowTaskSendEmail{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskSendEmailArrayInput)(nil)).Elem(), WorkflowTaskSendEmailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskSendEmailMapInput)(nil)).Elem(), WorkflowTaskSendEmailMap{})
	pulumi.RegisterOutputType(WorkflowTaskSendEmailOutput{})
	pulumi.RegisterOutputType(WorkflowTaskSendEmailArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTaskSendEmailMapOutput{})
}
