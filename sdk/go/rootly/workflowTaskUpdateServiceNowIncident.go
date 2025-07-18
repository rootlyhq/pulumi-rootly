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

// Manages workflow updateServiceNowIncident task.
type WorkflowTaskUpdateServiceNowIncident struct {
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
	TaskParams WorkflowTaskUpdateServiceNowIncidentTaskParamsOutput `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflowTaskUpdateServiceNowIncident registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTaskUpdateServiceNowIncident(ctx *pulumi.Context,
	name string, args *WorkflowTaskUpdateServiceNowIncidentArgs, opts ...pulumi.ResourceOption) (*WorkflowTaskUpdateServiceNowIncident, error) {
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
	var resource WorkflowTaskUpdateServiceNowIncident
	err := ctx.RegisterResource("rootly:index/workflowTaskUpdateServiceNowIncident:WorkflowTaskUpdateServiceNowIncident", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTaskUpdateServiceNowIncident gets an existing WorkflowTaskUpdateServiceNowIncident resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTaskUpdateServiceNowIncident(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTaskUpdateServiceNowIncidentState, opts ...pulumi.ResourceOption) (*WorkflowTaskUpdateServiceNowIncident, error) {
	var resource WorkflowTaskUpdateServiceNowIncident
	err := ctx.ReadResource("rootly:index/workflowTaskUpdateServiceNowIncident:WorkflowTaskUpdateServiceNowIncident", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTaskUpdateServiceNowIncident resources.
type workflowTaskUpdateServiceNowIncidentState struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams *WorkflowTaskUpdateServiceNowIncidentTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId *string `pulumi:"workflowId"`
}

type WorkflowTaskUpdateServiceNowIncidentState struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskUpdateServiceNowIncidentTaskParamsPtrInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringPtrInput
}

func (WorkflowTaskUpdateServiceNowIncidentState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskUpdateServiceNowIncidentState)(nil)).Elem()
}

type workflowTaskUpdateServiceNowIncidentArgs struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskUpdateServiceNowIncidentTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowTaskUpdateServiceNowIncident resource.
type WorkflowTaskUpdateServiceNowIncidentArgs struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskUpdateServiceNowIncidentTaskParamsInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringInput
}

func (WorkflowTaskUpdateServiceNowIncidentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskUpdateServiceNowIncidentArgs)(nil)).Elem()
}

type WorkflowTaskUpdateServiceNowIncidentInput interface {
	pulumi.Input

	ToWorkflowTaskUpdateServiceNowIncidentOutput() WorkflowTaskUpdateServiceNowIncidentOutput
	ToWorkflowTaskUpdateServiceNowIncidentOutputWithContext(ctx context.Context) WorkflowTaskUpdateServiceNowIncidentOutput
}

func (*WorkflowTaskUpdateServiceNowIncident) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskUpdateServiceNowIncident)(nil)).Elem()
}

func (i *WorkflowTaskUpdateServiceNowIncident) ToWorkflowTaskUpdateServiceNowIncidentOutput() WorkflowTaskUpdateServiceNowIncidentOutput {
	return i.ToWorkflowTaskUpdateServiceNowIncidentOutputWithContext(context.Background())
}

func (i *WorkflowTaskUpdateServiceNowIncident) ToWorkflowTaskUpdateServiceNowIncidentOutputWithContext(ctx context.Context) WorkflowTaskUpdateServiceNowIncidentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskUpdateServiceNowIncidentOutput)
}

// WorkflowTaskUpdateServiceNowIncidentArrayInput is an input type that accepts WorkflowTaskUpdateServiceNowIncidentArray and WorkflowTaskUpdateServiceNowIncidentArrayOutput values.
// You can construct a concrete instance of `WorkflowTaskUpdateServiceNowIncidentArrayInput` via:
//
//	WorkflowTaskUpdateServiceNowIncidentArray{ WorkflowTaskUpdateServiceNowIncidentArgs{...} }
type WorkflowTaskUpdateServiceNowIncidentArrayInput interface {
	pulumi.Input

	ToWorkflowTaskUpdateServiceNowIncidentArrayOutput() WorkflowTaskUpdateServiceNowIncidentArrayOutput
	ToWorkflowTaskUpdateServiceNowIncidentArrayOutputWithContext(context.Context) WorkflowTaskUpdateServiceNowIncidentArrayOutput
}

type WorkflowTaskUpdateServiceNowIncidentArray []WorkflowTaskUpdateServiceNowIncidentInput

func (WorkflowTaskUpdateServiceNowIncidentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskUpdateServiceNowIncident)(nil)).Elem()
}

func (i WorkflowTaskUpdateServiceNowIncidentArray) ToWorkflowTaskUpdateServiceNowIncidentArrayOutput() WorkflowTaskUpdateServiceNowIncidentArrayOutput {
	return i.ToWorkflowTaskUpdateServiceNowIncidentArrayOutputWithContext(context.Background())
}

func (i WorkflowTaskUpdateServiceNowIncidentArray) ToWorkflowTaskUpdateServiceNowIncidentArrayOutputWithContext(ctx context.Context) WorkflowTaskUpdateServiceNowIncidentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskUpdateServiceNowIncidentArrayOutput)
}

// WorkflowTaskUpdateServiceNowIncidentMapInput is an input type that accepts WorkflowTaskUpdateServiceNowIncidentMap and WorkflowTaskUpdateServiceNowIncidentMapOutput values.
// You can construct a concrete instance of `WorkflowTaskUpdateServiceNowIncidentMapInput` via:
//
//	WorkflowTaskUpdateServiceNowIncidentMap{ "key": WorkflowTaskUpdateServiceNowIncidentArgs{...} }
type WorkflowTaskUpdateServiceNowIncidentMapInput interface {
	pulumi.Input

	ToWorkflowTaskUpdateServiceNowIncidentMapOutput() WorkflowTaskUpdateServiceNowIncidentMapOutput
	ToWorkflowTaskUpdateServiceNowIncidentMapOutputWithContext(context.Context) WorkflowTaskUpdateServiceNowIncidentMapOutput
}

type WorkflowTaskUpdateServiceNowIncidentMap map[string]WorkflowTaskUpdateServiceNowIncidentInput

func (WorkflowTaskUpdateServiceNowIncidentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskUpdateServiceNowIncident)(nil)).Elem()
}

func (i WorkflowTaskUpdateServiceNowIncidentMap) ToWorkflowTaskUpdateServiceNowIncidentMapOutput() WorkflowTaskUpdateServiceNowIncidentMapOutput {
	return i.ToWorkflowTaskUpdateServiceNowIncidentMapOutputWithContext(context.Background())
}

func (i WorkflowTaskUpdateServiceNowIncidentMap) ToWorkflowTaskUpdateServiceNowIncidentMapOutputWithContext(ctx context.Context) WorkflowTaskUpdateServiceNowIncidentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskUpdateServiceNowIncidentMapOutput)
}

type WorkflowTaskUpdateServiceNowIncidentOutput struct{ *pulumi.OutputState }

func (WorkflowTaskUpdateServiceNowIncidentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskUpdateServiceNowIncident)(nil)).Elem()
}

func (o WorkflowTaskUpdateServiceNowIncidentOutput) ToWorkflowTaskUpdateServiceNowIncidentOutput() WorkflowTaskUpdateServiceNowIncidentOutput {
	return o
}

func (o WorkflowTaskUpdateServiceNowIncidentOutput) ToWorkflowTaskUpdateServiceNowIncidentOutputWithContext(ctx context.Context) WorkflowTaskUpdateServiceNowIncidentOutput {
	return o
}

// Enable/disable this workflow task
func (o WorkflowTaskUpdateServiceNowIncidentOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateServiceNowIncident) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Name of the workflow task
func (o WorkflowTaskUpdateServiceNowIncidentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateServiceNowIncident) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of the workflow task (1 being top of list)
func (o WorkflowTaskUpdateServiceNowIncidentOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateServiceNowIncident) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Skip workflow task if any failures
func (o WorkflowTaskUpdateServiceNowIncidentOutput) SkipOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateServiceNowIncident) pulumi.BoolPtrOutput { return v.SkipOnFailure }).(pulumi.BoolPtrOutput)
}

// The parameters for this workflow task.
func (o WorkflowTaskUpdateServiceNowIncidentOutput) TaskParams() WorkflowTaskUpdateServiceNowIncidentTaskParamsOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateServiceNowIncident) WorkflowTaskUpdateServiceNowIncidentTaskParamsOutput {
		return v.TaskParams
	}).(WorkflowTaskUpdateServiceNowIncidentTaskParamsOutput)
}

// The ID of the parent workflow
func (o WorkflowTaskUpdateServiceNowIncidentOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskUpdateServiceNowIncident) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

type WorkflowTaskUpdateServiceNowIncidentArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTaskUpdateServiceNowIncidentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskUpdateServiceNowIncident)(nil)).Elem()
}

func (o WorkflowTaskUpdateServiceNowIncidentArrayOutput) ToWorkflowTaskUpdateServiceNowIncidentArrayOutput() WorkflowTaskUpdateServiceNowIncidentArrayOutput {
	return o
}

func (o WorkflowTaskUpdateServiceNowIncidentArrayOutput) ToWorkflowTaskUpdateServiceNowIncidentArrayOutputWithContext(ctx context.Context) WorkflowTaskUpdateServiceNowIncidentArrayOutput {
	return o
}

func (o WorkflowTaskUpdateServiceNowIncidentArrayOutput) Index(i pulumi.IntInput) WorkflowTaskUpdateServiceNowIncidentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTaskUpdateServiceNowIncident {
		return vs[0].([]*WorkflowTaskUpdateServiceNowIncident)[vs[1].(int)]
	}).(WorkflowTaskUpdateServiceNowIncidentOutput)
}

type WorkflowTaskUpdateServiceNowIncidentMapOutput struct{ *pulumi.OutputState }

func (WorkflowTaskUpdateServiceNowIncidentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskUpdateServiceNowIncident)(nil)).Elem()
}

func (o WorkflowTaskUpdateServiceNowIncidentMapOutput) ToWorkflowTaskUpdateServiceNowIncidentMapOutput() WorkflowTaskUpdateServiceNowIncidentMapOutput {
	return o
}

func (o WorkflowTaskUpdateServiceNowIncidentMapOutput) ToWorkflowTaskUpdateServiceNowIncidentMapOutputWithContext(ctx context.Context) WorkflowTaskUpdateServiceNowIncidentMapOutput {
	return o
}

func (o WorkflowTaskUpdateServiceNowIncidentMapOutput) MapIndex(k pulumi.StringInput) WorkflowTaskUpdateServiceNowIncidentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTaskUpdateServiceNowIncident {
		return vs[0].(map[string]*WorkflowTaskUpdateServiceNowIncident)[vs[1].(string)]
	}).(WorkflowTaskUpdateServiceNowIncidentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskUpdateServiceNowIncidentInput)(nil)).Elem(), &WorkflowTaskUpdateServiceNowIncident{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskUpdateServiceNowIncidentArrayInput)(nil)).Elem(), WorkflowTaskUpdateServiceNowIncidentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskUpdateServiceNowIncidentMapInput)(nil)).Elem(), WorkflowTaskUpdateServiceNowIncidentMap{})
	pulumi.RegisterOutputType(WorkflowTaskUpdateServiceNowIncidentOutput{})
	pulumi.RegisterOutputType(WorkflowTaskUpdateServiceNowIncidentArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTaskUpdateServiceNowIncidentMapOutput{})
}
