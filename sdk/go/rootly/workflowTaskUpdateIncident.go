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

// Manages workflow updateIncident task.
type WorkflowTaskUpdateIncident struct {
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
	TaskParams WorkflowTaskUpdateIncidentTaskParamsOutput `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflowTaskUpdateIncident registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTaskUpdateIncident(ctx *pulumi.Context,
	name string, args *WorkflowTaskUpdateIncidentArgs, opts ...pulumi.ResourceOption) (*WorkflowTaskUpdateIncident, error) {
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
	var resource WorkflowTaskUpdateIncident
	err := ctx.RegisterResource("rootly:index/workflowTaskUpdateIncident:WorkflowTaskUpdateIncident", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTaskUpdateIncident gets an existing WorkflowTaskUpdateIncident resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTaskUpdateIncident(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTaskUpdateIncidentState, opts ...pulumi.ResourceOption) (*WorkflowTaskUpdateIncident, error) {
	var resource WorkflowTaskUpdateIncident
	err := ctx.ReadResource("rootly:index/workflowTaskUpdateIncident:WorkflowTaskUpdateIncident", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTaskUpdateIncident resources.
type workflowTaskUpdateIncidentState struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams *WorkflowTaskUpdateIncidentTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId *string `pulumi:"workflowId"`
}

type WorkflowTaskUpdateIncidentState struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskUpdateIncidentTaskParamsPtrInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringPtrInput
}

func (WorkflowTaskUpdateIncidentState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskUpdateIncidentState)(nil)).Elem()
}

type workflowTaskUpdateIncidentArgs struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskUpdateIncidentTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowTaskUpdateIncident resource.
type WorkflowTaskUpdateIncidentArgs struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskUpdateIncidentTaskParamsInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringInput
}

func (WorkflowTaskUpdateIncidentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskUpdateIncidentArgs)(nil)).Elem()
}

type WorkflowTaskUpdateIncidentInput interface {
	pulumi.Input

	ToWorkflowTaskUpdateIncidentOutput() WorkflowTaskUpdateIncidentOutput
	ToWorkflowTaskUpdateIncidentOutputWithContext(ctx context.Context) WorkflowTaskUpdateIncidentOutput
}

func (*WorkflowTaskUpdateIncident) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskUpdateIncident)(nil)).Elem()
}

func (i *WorkflowTaskUpdateIncident) ToWorkflowTaskUpdateIncidentOutput() WorkflowTaskUpdateIncidentOutput {
	return i.ToWorkflowTaskUpdateIncidentOutputWithContext(context.Background())
}

func (i *WorkflowTaskUpdateIncident) ToWorkflowTaskUpdateIncidentOutputWithContext(ctx context.Context) WorkflowTaskUpdateIncidentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskUpdateIncidentOutput)
}

// WorkflowTaskUpdateIncidentArrayInput is an input type that accepts WorkflowTaskUpdateIncidentArray and WorkflowTaskUpdateIncidentArrayOutput values.
// You can construct a concrete instance of `WorkflowTaskUpdateIncidentArrayInput` via:
//
//	WorkflowTaskUpdateIncidentArray{ WorkflowTaskUpdateIncidentArgs{...} }
type WorkflowTaskUpdateIncidentArrayInput interface {
	pulumi.Input

	ToWorkflowTaskUpdateIncidentArrayOutput() WorkflowTaskUpdateIncidentArrayOutput
	ToWorkflowTaskUpdateIncidentArrayOutputWithContext(context.Context) WorkflowTaskUpdateIncidentArrayOutput
}

type WorkflowTaskUpdateIncidentArray []WorkflowTaskUpdateIncidentInput

func (WorkflowTaskUpdateIncidentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskUpdateIncident)(nil)).Elem()
}

func (i WorkflowTaskUpdateIncidentArray) ToWorkflowTaskUpdateIncidentArrayOutput() WorkflowTaskUpdateIncidentArrayOutput {
	return i.ToWorkflowTaskUpdateIncidentArrayOutputWithContext(context.Background())
}

func (i WorkflowTaskUpdateIncidentArray) ToWorkflowTaskUpdateIncidentArrayOutputWithContext(ctx context.Context) WorkflowTaskUpdateIncidentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskUpdateIncidentArrayOutput)
}

// WorkflowTaskUpdateIncidentMapInput is an input type that accepts WorkflowTaskUpdateIncidentMap and WorkflowTaskUpdateIncidentMapOutput values.
// You can construct a concrete instance of `WorkflowTaskUpdateIncidentMapInput` via:
//
//	WorkflowTaskUpdateIncidentMap{ "key": WorkflowTaskUpdateIncidentArgs{...} }
type WorkflowTaskUpdateIncidentMapInput interface {
	pulumi.Input

	ToWorkflowTaskUpdateIncidentMapOutput() WorkflowTaskUpdateIncidentMapOutput
	ToWorkflowTaskUpdateIncidentMapOutputWithContext(context.Context) WorkflowTaskUpdateIncidentMapOutput
}

type WorkflowTaskUpdateIncidentMap map[string]WorkflowTaskUpdateIncidentInput

func (WorkflowTaskUpdateIncidentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskUpdateIncident)(nil)).Elem()
}

func (i WorkflowTaskUpdateIncidentMap) ToWorkflowTaskUpdateIncidentMapOutput() WorkflowTaskUpdateIncidentMapOutput {
	return i.ToWorkflowTaskUpdateIncidentMapOutputWithContext(context.Background())
}

func (i WorkflowTaskUpdateIncidentMap) ToWorkflowTaskUpdateIncidentMapOutputWithContext(ctx context.Context) WorkflowTaskUpdateIncidentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskUpdateIncidentMapOutput)
}

type WorkflowTaskUpdateIncidentOutput struct{ *pulumi.OutputState }

func (WorkflowTaskUpdateIncidentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskUpdateIncident)(nil)).Elem()
}

func (o WorkflowTaskUpdateIncidentOutput) ToWorkflowTaskUpdateIncidentOutput() WorkflowTaskUpdateIncidentOutput {
	return o
}

func (o WorkflowTaskUpdateIncidentOutput) ToWorkflowTaskUpdateIncidentOutputWithContext(ctx context.Context) WorkflowTaskUpdateIncidentOutput {
	return o
}

// Enable/disable this workflow task
func (o WorkflowTaskUpdateIncidentOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateIncident) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Name of the workflow task
func (o WorkflowTaskUpdateIncidentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateIncident) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of the workflow task (1 being top of list)
func (o WorkflowTaskUpdateIncidentOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateIncident) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Skip workflow task if any failures
func (o WorkflowTaskUpdateIncidentOutput) SkipOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateIncident) pulumi.BoolPtrOutput { return v.SkipOnFailure }).(pulumi.BoolPtrOutput)
}

// The parameters for this workflow task.
func (o WorkflowTaskUpdateIncidentOutput) TaskParams() WorkflowTaskUpdateIncidentTaskParamsOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateIncident) WorkflowTaskUpdateIncidentTaskParamsOutput { return v.TaskParams }).(WorkflowTaskUpdateIncidentTaskParamsOutput)
}

// The ID of the parent workflow
func (o WorkflowTaskUpdateIncidentOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateIncident) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

type WorkflowTaskUpdateIncidentArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTaskUpdateIncidentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskUpdateIncident)(nil)).Elem()
}

func (o WorkflowTaskUpdateIncidentArrayOutput) ToWorkflowTaskUpdateIncidentArrayOutput() WorkflowTaskUpdateIncidentArrayOutput {
	return o
}

func (o WorkflowTaskUpdateIncidentArrayOutput) ToWorkflowTaskUpdateIncidentArrayOutputWithContext(ctx context.Context) WorkflowTaskUpdateIncidentArrayOutput {
	return o
}

func (o WorkflowTaskUpdateIncidentArrayOutput) Index(i pulumi.IntInput) WorkflowTaskUpdateIncidentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTaskUpdateIncident {
		return vs[0].([]*WorkflowTaskUpdateIncident)[vs[1].(int)]
	}).(WorkflowTaskUpdateIncidentOutput)
}

type WorkflowTaskUpdateIncidentMapOutput struct{ *pulumi.OutputState }

func (WorkflowTaskUpdateIncidentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskUpdateIncident)(nil)).Elem()
}

func (o WorkflowTaskUpdateIncidentMapOutput) ToWorkflowTaskUpdateIncidentMapOutput() WorkflowTaskUpdateIncidentMapOutput {
	return o
}

func (o WorkflowTaskUpdateIncidentMapOutput) ToWorkflowTaskUpdateIncidentMapOutputWithContext(ctx context.Context) WorkflowTaskUpdateIncidentMapOutput {
	return o
}

func (o WorkflowTaskUpdateIncidentMapOutput) MapIndex(k pulumi.StringInput) WorkflowTaskUpdateIncidentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTaskUpdateIncident {
		return vs[0].(map[string]*WorkflowTaskUpdateIncident)[vs[1].(string)]
	}).(WorkflowTaskUpdateIncidentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskUpdateIncidentInput)(nil)).Elem(), &WorkflowTaskUpdateIncident{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskUpdateIncidentArrayInput)(nil)).Elem(), WorkflowTaskUpdateIncidentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskUpdateIncidentMapInput)(nil)).Elem(), WorkflowTaskUpdateIncidentMap{})
	pulumi.RegisterOutputType(WorkflowTaskUpdateIncidentOutput{})
	pulumi.RegisterOutputType(WorkflowTaskUpdateIncidentArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTaskUpdateIncidentMapOutput{})
}
