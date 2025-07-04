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

// Manages workflow sendDashboardReport task.
type WorkflowTaskSendDashboardReport struct {
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
	TaskParams WorkflowTaskSendDashboardReportTaskParamsOutput `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflowTaskSendDashboardReport registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTaskSendDashboardReport(ctx *pulumi.Context,
	name string, args *WorkflowTaskSendDashboardReportArgs, opts ...pulumi.ResourceOption) (*WorkflowTaskSendDashboardReport, error) {
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
	var resource WorkflowTaskSendDashboardReport
	err := ctx.RegisterResource("rootly:index/workflowTaskSendDashboardReport:WorkflowTaskSendDashboardReport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTaskSendDashboardReport gets an existing WorkflowTaskSendDashboardReport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTaskSendDashboardReport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTaskSendDashboardReportState, opts ...pulumi.ResourceOption) (*WorkflowTaskSendDashboardReport, error) {
	var resource WorkflowTaskSendDashboardReport
	err := ctx.ReadResource("rootly:index/workflowTaskSendDashboardReport:WorkflowTaskSendDashboardReport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTaskSendDashboardReport resources.
type workflowTaskSendDashboardReportState struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams *WorkflowTaskSendDashboardReportTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId *string `pulumi:"workflowId"`
}

type WorkflowTaskSendDashboardReportState struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskSendDashboardReportTaskParamsPtrInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringPtrInput
}

func (WorkflowTaskSendDashboardReportState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskSendDashboardReportState)(nil)).Elem()
}

type workflowTaskSendDashboardReportArgs struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskSendDashboardReportTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowTaskSendDashboardReport resource.
type WorkflowTaskSendDashboardReportArgs struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskSendDashboardReportTaskParamsInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringInput
}

func (WorkflowTaskSendDashboardReportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskSendDashboardReportArgs)(nil)).Elem()
}

type WorkflowTaskSendDashboardReportInput interface {
	pulumi.Input

	ToWorkflowTaskSendDashboardReportOutput() WorkflowTaskSendDashboardReportOutput
	ToWorkflowTaskSendDashboardReportOutputWithContext(ctx context.Context) WorkflowTaskSendDashboardReportOutput
}

func (*WorkflowTaskSendDashboardReport) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskSendDashboardReport)(nil)).Elem()
}

func (i *WorkflowTaskSendDashboardReport) ToWorkflowTaskSendDashboardReportOutput() WorkflowTaskSendDashboardReportOutput {
	return i.ToWorkflowTaskSendDashboardReportOutputWithContext(context.Background())
}

func (i *WorkflowTaskSendDashboardReport) ToWorkflowTaskSendDashboardReportOutputWithContext(ctx context.Context) WorkflowTaskSendDashboardReportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskSendDashboardReportOutput)
}

// WorkflowTaskSendDashboardReportArrayInput is an input type that accepts WorkflowTaskSendDashboardReportArray and WorkflowTaskSendDashboardReportArrayOutput values.
// You can construct a concrete instance of `WorkflowTaskSendDashboardReportArrayInput` via:
//
//	WorkflowTaskSendDashboardReportArray{ WorkflowTaskSendDashboardReportArgs{...} }
type WorkflowTaskSendDashboardReportArrayInput interface {
	pulumi.Input

	ToWorkflowTaskSendDashboardReportArrayOutput() WorkflowTaskSendDashboardReportArrayOutput
	ToWorkflowTaskSendDashboardReportArrayOutputWithContext(context.Context) WorkflowTaskSendDashboardReportArrayOutput
}

type WorkflowTaskSendDashboardReportArray []WorkflowTaskSendDashboardReportInput

func (WorkflowTaskSendDashboardReportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskSendDashboardReport)(nil)).Elem()
}

func (i WorkflowTaskSendDashboardReportArray) ToWorkflowTaskSendDashboardReportArrayOutput() WorkflowTaskSendDashboardReportArrayOutput {
	return i.ToWorkflowTaskSendDashboardReportArrayOutputWithContext(context.Background())
}

func (i WorkflowTaskSendDashboardReportArray) ToWorkflowTaskSendDashboardReportArrayOutputWithContext(ctx context.Context) WorkflowTaskSendDashboardReportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskSendDashboardReportArrayOutput)
}

// WorkflowTaskSendDashboardReportMapInput is an input type that accepts WorkflowTaskSendDashboardReportMap and WorkflowTaskSendDashboardReportMapOutput values.
// You can construct a concrete instance of `WorkflowTaskSendDashboardReportMapInput` via:
//
//	WorkflowTaskSendDashboardReportMap{ "key": WorkflowTaskSendDashboardReportArgs{...} }
type WorkflowTaskSendDashboardReportMapInput interface {
	pulumi.Input

	ToWorkflowTaskSendDashboardReportMapOutput() WorkflowTaskSendDashboardReportMapOutput
	ToWorkflowTaskSendDashboardReportMapOutputWithContext(context.Context) WorkflowTaskSendDashboardReportMapOutput
}

type WorkflowTaskSendDashboardReportMap map[string]WorkflowTaskSendDashboardReportInput

func (WorkflowTaskSendDashboardReportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskSendDashboardReport)(nil)).Elem()
}

func (i WorkflowTaskSendDashboardReportMap) ToWorkflowTaskSendDashboardReportMapOutput() WorkflowTaskSendDashboardReportMapOutput {
	return i.ToWorkflowTaskSendDashboardReportMapOutputWithContext(context.Background())
}

func (i WorkflowTaskSendDashboardReportMap) ToWorkflowTaskSendDashboardReportMapOutputWithContext(ctx context.Context) WorkflowTaskSendDashboardReportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskSendDashboardReportMapOutput)
}

type WorkflowTaskSendDashboardReportOutput struct{ *pulumi.OutputState }

func (WorkflowTaskSendDashboardReportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskSendDashboardReport)(nil)).Elem()
}

func (o WorkflowTaskSendDashboardReportOutput) ToWorkflowTaskSendDashboardReportOutput() WorkflowTaskSendDashboardReportOutput {
	return o
}

func (o WorkflowTaskSendDashboardReportOutput) ToWorkflowTaskSendDashboardReportOutputWithContext(ctx context.Context) WorkflowTaskSendDashboardReportOutput {
	return o
}

// Enable/disable this workflow task
func (o WorkflowTaskSendDashboardReportOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskSendDashboardReport) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Name of the workflow task
func (o WorkflowTaskSendDashboardReportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskSendDashboardReport) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of the workflow task (1 being top of list)
func (o WorkflowTaskSendDashboardReportOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTaskSendDashboardReport) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Skip workflow task if any failures
func (o WorkflowTaskSendDashboardReportOutput) SkipOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskSendDashboardReport) pulumi.BoolPtrOutput { return v.SkipOnFailure }).(pulumi.BoolPtrOutput)
}

// The parameters for this workflow task.
func (o WorkflowTaskSendDashboardReportOutput) TaskParams() WorkflowTaskSendDashboardReportTaskParamsOutput {
	return o.ApplyT(func(v *WorkflowTaskSendDashboardReport) WorkflowTaskSendDashboardReportTaskParamsOutput {
		return v.TaskParams
	}).(WorkflowTaskSendDashboardReportTaskParamsOutput)
}

// The ID of the parent workflow
func (o WorkflowTaskSendDashboardReportOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskSendDashboardReport) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

type WorkflowTaskSendDashboardReportArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTaskSendDashboardReportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskSendDashboardReport)(nil)).Elem()
}

func (o WorkflowTaskSendDashboardReportArrayOutput) ToWorkflowTaskSendDashboardReportArrayOutput() WorkflowTaskSendDashboardReportArrayOutput {
	return o
}

func (o WorkflowTaskSendDashboardReportArrayOutput) ToWorkflowTaskSendDashboardReportArrayOutputWithContext(ctx context.Context) WorkflowTaskSendDashboardReportArrayOutput {
	return o
}

func (o WorkflowTaskSendDashboardReportArrayOutput) Index(i pulumi.IntInput) WorkflowTaskSendDashboardReportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTaskSendDashboardReport {
		return vs[0].([]*WorkflowTaskSendDashboardReport)[vs[1].(int)]
	}).(WorkflowTaskSendDashboardReportOutput)
}

type WorkflowTaskSendDashboardReportMapOutput struct{ *pulumi.OutputState }

func (WorkflowTaskSendDashboardReportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskSendDashboardReport)(nil)).Elem()
}

func (o WorkflowTaskSendDashboardReportMapOutput) ToWorkflowTaskSendDashboardReportMapOutput() WorkflowTaskSendDashboardReportMapOutput {
	return o
}

func (o WorkflowTaskSendDashboardReportMapOutput) ToWorkflowTaskSendDashboardReportMapOutputWithContext(ctx context.Context) WorkflowTaskSendDashboardReportMapOutput {
	return o
}

func (o WorkflowTaskSendDashboardReportMapOutput) MapIndex(k pulumi.StringInput) WorkflowTaskSendDashboardReportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTaskSendDashboardReport {
		return vs[0].(map[string]*WorkflowTaskSendDashboardReport)[vs[1].(string)]
	}).(WorkflowTaskSendDashboardReportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskSendDashboardReportInput)(nil)).Elem(), &WorkflowTaskSendDashboardReport{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskSendDashboardReportArrayInput)(nil)).Elem(), WorkflowTaskSendDashboardReportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskSendDashboardReportMapInput)(nil)).Elem(), WorkflowTaskSendDashboardReportMap{})
	pulumi.RegisterOutputType(WorkflowTaskSendDashboardReportOutput{})
	pulumi.RegisterOutputType(WorkflowTaskSendDashboardReportArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTaskSendDashboardReportMapOutput{})
}
