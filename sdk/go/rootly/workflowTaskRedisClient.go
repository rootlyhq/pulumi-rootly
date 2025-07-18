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

// Manages workflow redisClient task.
type WorkflowTaskRedisClient struct {
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
	TaskParams WorkflowTaskRedisClientTaskParamsOutput `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflowTaskRedisClient registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTaskRedisClient(ctx *pulumi.Context,
	name string, args *WorkflowTaskRedisClientArgs, opts ...pulumi.ResourceOption) (*WorkflowTaskRedisClient, error) {
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
	var resource WorkflowTaskRedisClient
	err := ctx.RegisterResource("rootly:index/workflowTaskRedisClient:WorkflowTaskRedisClient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTaskRedisClient gets an existing WorkflowTaskRedisClient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTaskRedisClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTaskRedisClientState, opts ...pulumi.ResourceOption) (*WorkflowTaskRedisClient, error) {
	var resource WorkflowTaskRedisClient
	err := ctx.ReadResource("rootly:index/workflowTaskRedisClient:WorkflowTaskRedisClient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTaskRedisClient resources.
type workflowTaskRedisClientState struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams *WorkflowTaskRedisClientTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId *string `pulumi:"workflowId"`
}

type WorkflowTaskRedisClientState struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskRedisClientTaskParamsPtrInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringPtrInput
}

func (WorkflowTaskRedisClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskRedisClientState)(nil)).Elem()
}

type workflowTaskRedisClientArgs struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskRedisClientTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowTaskRedisClient resource.
type WorkflowTaskRedisClientArgs struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskRedisClientTaskParamsInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringInput
}

func (WorkflowTaskRedisClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskRedisClientArgs)(nil)).Elem()
}

type WorkflowTaskRedisClientInput interface {
	pulumi.Input

	ToWorkflowTaskRedisClientOutput() WorkflowTaskRedisClientOutput
	ToWorkflowTaskRedisClientOutputWithContext(ctx context.Context) WorkflowTaskRedisClientOutput
}

func (*WorkflowTaskRedisClient) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskRedisClient)(nil)).Elem()
}

func (i *WorkflowTaskRedisClient) ToWorkflowTaskRedisClientOutput() WorkflowTaskRedisClientOutput {
	return i.ToWorkflowTaskRedisClientOutputWithContext(context.Background())
}

func (i *WorkflowTaskRedisClient) ToWorkflowTaskRedisClientOutputWithContext(ctx context.Context) WorkflowTaskRedisClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskRedisClientOutput)
}

// WorkflowTaskRedisClientArrayInput is an input type that accepts WorkflowTaskRedisClientArray and WorkflowTaskRedisClientArrayOutput values.
// You can construct a concrete instance of `WorkflowTaskRedisClientArrayInput` via:
//
//	WorkflowTaskRedisClientArray{ WorkflowTaskRedisClientArgs{...} }
type WorkflowTaskRedisClientArrayInput interface {
	pulumi.Input

	ToWorkflowTaskRedisClientArrayOutput() WorkflowTaskRedisClientArrayOutput
	ToWorkflowTaskRedisClientArrayOutputWithContext(context.Context) WorkflowTaskRedisClientArrayOutput
}

type WorkflowTaskRedisClientArray []WorkflowTaskRedisClientInput

func (WorkflowTaskRedisClientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskRedisClient)(nil)).Elem()
}

func (i WorkflowTaskRedisClientArray) ToWorkflowTaskRedisClientArrayOutput() WorkflowTaskRedisClientArrayOutput {
	return i.ToWorkflowTaskRedisClientArrayOutputWithContext(context.Background())
}

func (i WorkflowTaskRedisClientArray) ToWorkflowTaskRedisClientArrayOutputWithContext(ctx context.Context) WorkflowTaskRedisClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskRedisClientArrayOutput)
}

// WorkflowTaskRedisClientMapInput is an input type that accepts WorkflowTaskRedisClientMap and WorkflowTaskRedisClientMapOutput values.
// You can construct a concrete instance of `WorkflowTaskRedisClientMapInput` via:
//
//	WorkflowTaskRedisClientMap{ "key": WorkflowTaskRedisClientArgs{...} }
type WorkflowTaskRedisClientMapInput interface {
	pulumi.Input

	ToWorkflowTaskRedisClientMapOutput() WorkflowTaskRedisClientMapOutput
	ToWorkflowTaskRedisClientMapOutputWithContext(context.Context) WorkflowTaskRedisClientMapOutput
}

type WorkflowTaskRedisClientMap map[string]WorkflowTaskRedisClientInput

func (WorkflowTaskRedisClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskRedisClient)(nil)).Elem()
}

func (i WorkflowTaskRedisClientMap) ToWorkflowTaskRedisClientMapOutput() WorkflowTaskRedisClientMapOutput {
	return i.ToWorkflowTaskRedisClientMapOutputWithContext(context.Background())
}

func (i WorkflowTaskRedisClientMap) ToWorkflowTaskRedisClientMapOutputWithContext(ctx context.Context) WorkflowTaskRedisClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskRedisClientMapOutput)
}

type WorkflowTaskRedisClientOutput struct{ *pulumi.OutputState }

func (WorkflowTaskRedisClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskRedisClient)(nil)).Elem()
}

func (o WorkflowTaskRedisClientOutput) ToWorkflowTaskRedisClientOutput() WorkflowTaskRedisClientOutput {
	return o
}

func (o WorkflowTaskRedisClientOutput) ToWorkflowTaskRedisClientOutputWithContext(ctx context.Context) WorkflowTaskRedisClientOutput {
	return o
}

// Enable/disable this workflow task
func (o WorkflowTaskRedisClientOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskRedisClient) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Name of the workflow task
func (o WorkflowTaskRedisClientOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskRedisClient) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of the workflow task (1 being top of list)
func (o WorkflowTaskRedisClientOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTaskRedisClient) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Skip workflow task if any failures
func (o WorkflowTaskRedisClientOutput) SkipOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskRedisClient) pulumi.BoolPtrOutput { return v.SkipOnFailure }).(pulumi.BoolPtrOutput)
}

// The parameters for this workflow task.
func (o WorkflowTaskRedisClientOutput) TaskParams() WorkflowTaskRedisClientTaskParamsOutput {
	return o.ApplyT(func(v *WorkflowTaskRedisClient) WorkflowTaskRedisClientTaskParamsOutput { return v.TaskParams }).(WorkflowTaskRedisClientTaskParamsOutput)
}

// The ID of the parent workflow
func (o WorkflowTaskRedisClientOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskRedisClient) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

type WorkflowTaskRedisClientArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTaskRedisClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskRedisClient)(nil)).Elem()
}

func (o WorkflowTaskRedisClientArrayOutput) ToWorkflowTaskRedisClientArrayOutput() WorkflowTaskRedisClientArrayOutput {
	return o
}

func (o WorkflowTaskRedisClientArrayOutput) ToWorkflowTaskRedisClientArrayOutputWithContext(ctx context.Context) WorkflowTaskRedisClientArrayOutput {
	return o
}

func (o WorkflowTaskRedisClientArrayOutput) Index(i pulumi.IntInput) WorkflowTaskRedisClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTaskRedisClient {
		return vs[0].([]*WorkflowTaskRedisClient)[vs[1].(int)]
	}).(WorkflowTaskRedisClientOutput)
}

type WorkflowTaskRedisClientMapOutput struct{ *pulumi.OutputState }

func (WorkflowTaskRedisClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskRedisClient)(nil)).Elem()
}

func (o WorkflowTaskRedisClientMapOutput) ToWorkflowTaskRedisClientMapOutput() WorkflowTaskRedisClientMapOutput {
	return o
}

func (o WorkflowTaskRedisClientMapOutput) ToWorkflowTaskRedisClientMapOutputWithContext(ctx context.Context) WorkflowTaskRedisClientMapOutput {
	return o
}

func (o WorkflowTaskRedisClientMapOutput) MapIndex(k pulumi.StringInput) WorkflowTaskRedisClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTaskRedisClient {
		return vs[0].(map[string]*WorkflowTaskRedisClient)[vs[1].(string)]
	}).(WorkflowTaskRedisClientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskRedisClientInput)(nil)).Elem(), &WorkflowTaskRedisClient{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskRedisClientArrayInput)(nil)).Elem(), WorkflowTaskRedisClientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskRedisClientMapInput)(nil)).Elem(), WorkflowTaskRedisClientMap{})
	pulumi.RegisterOutputType(WorkflowTaskRedisClientOutput{})
	pulumi.RegisterOutputType(WorkflowTaskRedisClientArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTaskRedisClientMapOutput{})
}
