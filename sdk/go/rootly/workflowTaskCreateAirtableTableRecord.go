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

// Manages workflow createAirtableTableRecord task.
type WorkflowTaskCreateAirtableTableRecord struct {
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
	TaskParams WorkflowTaskCreateAirtableTableRecordTaskParamsOutput `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflowTaskCreateAirtableTableRecord registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTaskCreateAirtableTableRecord(ctx *pulumi.Context,
	name string, args *WorkflowTaskCreateAirtableTableRecordArgs, opts ...pulumi.ResourceOption) (*WorkflowTaskCreateAirtableTableRecord, error) {
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
	var resource WorkflowTaskCreateAirtableTableRecord
	err := ctx.RegisterResource("rootly:index/workflowTaskCreateAirtableTableRecord:WorkflowTaskCreateAirtableTableRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTaskCreateAirtableTableRecord gets an existing WorkflowTaskCreateAirtableTableRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTaskCreateAirtableTableRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTaskCreateAirtableTableRecordState, opts ...pulumi.ResourceOption) (*WorkflowTaskCreateAirtableTableRecord, error) {
	var resource WorkflowTaskCreateAirtableTableRecord
	err := ctx.ReadResource("rootly:index/workflowTaskCreateAirtableTableRecord:WorkflowTaskCreateAirtableTableRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTaskCreateAirtableTableRecord resources.
type workflowTaskCreateAirtableTableRecordState struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams *WorkflowTaskCreateAirtableTableRecordTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId *string `pulumi:"workflowId"`
}

type WorkflowTaskCreateAirtableTableRecordState struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateAirtableTableRecordTaskParamsPtrInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringPtrInput
}

func (WorkflowTaskCreateAirtableTableRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskCreateAirtableTableRecordState)(nil)).Elem()
}

type workflowTaskCreateAirtableTableRecordArgs struct {
	// Enable/disable this workflow task
	Enabled *bool `pulumi:"enabled"`
	// Name of the workflow task
	Name *string `pulumi:"name"`
	// The position of the workflow task (1 being top of list)
	Position *int `pulumi:"position"`
	// Skip workflow task if any failures
	SkipOnFailure *bool `pulumi:"skipOnFailure"`
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateAirtableTableRecordTaskParams `pulumi:"taskParams"`
	// The ID of the parent workflow
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowTaskCreateAirtableTableRecord resource.
type WorkflowTaskCreateAirtableTableRecordArgs struct {
	// Enable/disable this workflow task
	Enabled pulumi.BoolPtrInput
	// Name of the workflow task
	Name pulumi.StringPtrInput
	// The position of the workflow task (1 being top of list)
	Position pulumi.IntPtrInput
	// Skip workflow task if any failures
	SkipOnFailure pulumi.BoolPtrInput
	// The parameters for this workflow task.
	TaskParams WorkflowTaskCreateAirtableTableRecordTaskParamsInput
	// The ID of the parent workflow
	WorkflowId pulumi.StringInput
}

func (WorkflowTaskCreateAirtableTableRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTaskCreateAirtableTableRecordArgs)(nil)).Elem()
}

type WorkflowTaskCreateAirtableTableRecordInput interface {
	pulumi.Input

	ToWorkflowTaskCreateAirtableTableRecordOutput() WorkflowTaskCreateAirtableTableRecordOutput
	ToWorkflowTaskCreateAirtableTableRecordOutputWithContext(ctx context.Context) WorkflowTaskCreateAirtableTableRecordOutput
}

func (*WorkflowTaskCreateAirtableTableRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskCreateAirtableTableRecord)(nil)).Elem()
}

func (i *WorkflowTaskCreateAirtableTableRecord) ToWorkflowTaskCreateAirtableTableRecordOutput() WorkflowTaskCreateAirtableTableRecordOutput {
	return i.ToWorkflowTaskCreateAirtableTableRecordOutputWithContext(context.Background())
}

func (i *WorkflowTaskCreateAirtableTableRecord) ToWorkflowTaskCreateAirtableTableRecordOutputWithContext(ctx context.Context) WorkflowTaskCreateAirtableTableRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateAirtableTableRecordOutput)
}

// WorkflowTaskCreateAirtableTableRecordArrayInput is an input type that accepts WorkflowTaskCreateAirtableTableRecordArray and WorkflowTaskCreateAirtableTableRecordArrayOutput values.
// You can construct a concrete instance of `WorkflowTaskCreateAirtableTableRecordArrayInput` via:
//
//	WorkflowTaskCreateAirtableTableRecordArray{ WorkflowTaskCreateAirtableTableRecordArgs{...} }
type WorkflowTaskCreateAirtableTableRecordArrayInput interface {
	pulumi.Input

	ToWorkflowTaskCreateAirtableTableRecordArrayOutput() WorkflowTaskCreateAirtableTableRecordArrayOutput
	ToWorkflowTaskCreateAirtableTableRecordArrayOutputWithContext(context.Context) WorkflowTaskCreateAirtableTableRecordArrayOutput
}

type WorkflowTaskCreateAirtableTableRecordArray []WorkflowTaskCreateAirtableTableRecordInput

func (WorkflowTaskCreateAirtableTableRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskCreateAirtableTableRecord)(nil)).Elem()
}

func (i WorkflowTaskCreateAirtableTableRecordArray) ToWorkflowTaskCreateAirtableTableRecordArrayOutput() WorkflowTaskCreateAirtableTableRecordArrayOutput {
	return i.ToWorkflowTaskCreateAirtableTableRecordArrayOutputWithContext(context.Background())
}

func (i WorkflowTaskCreateAirtableTableRecordArray) ToWorkflowTaskCreateAirtableTableRecordArrayOutputWithContext(ctx context.Context) WorkflowTaskCreateAirtableTableRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateAirtableTableRecordArrayOutput)
}

// WorkflowTaskCreateAirtableTableRecordMapInput is an input type that accepts WorkflowTaskCreateAirtableTableRecordMap and WorkflowTaskCreateAirtableTableRecordMapOutput values.
// You can construct a concrete instance of `WorkflowTaskCreateAirtableTableRecordMapInput` via:
//
//	WorkflowTaskCreateAirtableTableRecordMap{ "key": WorkflowTaskCreateAirtableTableRecordArgs{...} }
type WorkflowTaskCreateAirtableTableRecordMapInput interface {
	pulumi.Input

	ToWorkflowTaskCreateAirtableTableRecordMapOutput() WorkflowTaskCreateAirtableTableRecordMapOutput
	ToWorkflowTaskCreateAirtableTableRecordMapOutputWithContext(context.Context) WorkflowTaskCreateAirtableTableRecordMapOutput
}

type WorkflowTaskCreateAirtableTableRecordMap map[string]WorkflowTaskCreateAirtableTableRecordInput

func (WorkflowTaskCreateAirtableTableRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskCreateAirtableTableRecord)(nil)).Elem()
}

func (i WorkflowTaskCreateAirtableTableRecordMap) ToWorkflowTaskCreateAirtableTableRecordMapOutput() WorkflowTaskCreateAirtableTableRecordMapOutput {
	return i.ToWorkflowTaskCreateAirtableTableRecordMapOutputWithContext(context.Background())
}

func (i WorkflowTaskCreateAirtableTableRecordMap) ToWorkflowTaskCreateAirtableTableRecordMapOutputWithContext(ctx context.Context) WorkflowTaskCreateAirtableTableRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTaskCreateAirtableTableRecordMapOutput)
}

type WorkflowTaskCreateAirtableTableRecordOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateAirtableTableRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTaskCreateAirtableTableRecord)(nil)).Elem()
}

func (o WorkflowTaskCreateAirtableTableRecordOutput) ToWorkflowTaskCreateAirtableTableRecordOutput() WorkflowTaskCreateAirtableTableRecordOutput {
	return o
}

func (o WorkflowTaskCreateAirtableTableRecordOutput) ToWorkflowTaskCreateAirtableTableRecordOutputWithContext(ctx context.Context) WorkflowTaskCreateAirtableTableRecordOutput {
	return o
}

// Enable/disable this workflow task
func (o WorkflowTaskCreateAirtableTableRecordOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateAirtableTableRecord) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Name of the workflow task
func (o WorkflowTaskCreateAirtableTableRecordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateAirtableTableRecord) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of the workflow task (1 being top of list)
func (o WorkflowTaskCreateAirtableTableRecordOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateAirtableTableRecord) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Skip workflow task if any failures
func (o WorkflowTaskCreateAirtableTableRecordOutput) SkipOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateAirtableTableRecord) pulumi.BoolPtrOutput { return v.SkipOnFailure }).(pulumi.BoolPtrOutput)
}

// The parameters for this workflow task.
func (o WorkflowTaskCreateAirtableTableRecordOutput) TaskParams() WorkflowTaskCreateAirtableTableRecordTaskParamsOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateAirtableTableRecord) WorkflowTaskCreateAirtableTableRecordTaskParamsOutput {
		return v.TaskParams
	}).(WorkflowTaskCreateAirtableTableRecordTaskParamsOutput)
}

// The ID of the parent workflow
func (o WorkflowTaskCreateAirtableTableRecordOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTaskCreateAirtableTableRecord) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

type WorkflowTaskCreateAirtableTableRecordArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateAirtableTableRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTaskCreateAirtableTableRecord)(nil)).Elem()
}

func (o WorkflowTaskCreateAirtableTableRecordArrayOutput) ToWorkflowTaskCreateAirtableTableRecordArrayOutput() WorkflowTaskCreateAirtableTableRecordArrayOutput {
	return o
}

func (o WorkflowTaskCreateAirtableTableRecordArrayOutput) ToWorkflowTaskCreateAirtableTableRecordArrayOutputWithContext(ctx context.Context) WorkflowTaskCreateAirtableTableRecordArrayOutput {
	return o
}

func (o WorkflowTaskCreateAirtableTableRecordArrayOutput) Index(i pulumi.IntInput) WorkflowTaskCreateAirtableTableRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTaskCreateAirtableTableRecord {
		return vs[0].([]*WorkflowTaskCreateAirtableTableRecord)[vs[1].(int)]
	}).(WorkflowTaskCreateAirtableTableRecordOutput)
}

type WorkflowTaskCreateAirtableTableRecordMapOutput struct{ *pulumi.OutputState }

func (WorkflowTaskCreateAirtableTableRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTaskCreateAirtableTableRecord)(nil)).Elem()
}

func (o WorkflowTaskCreateAirtableTableRecordMapOutput) ToWorkflowTaskCreateAirtableTableRecordMapOutput() WorkflowTaskCreateAirtableTableRecordMapOutput {
	return o
}

func (o WorkflowTaskCreateAirtableTableRecordMapOutput) ToWorkflowTaskCreateAirtableTableRecordMapOutputWithContext(ctx context.Context) WorkflowTaskCreateAirtableTableRecordMapOutput {
	return o
}

func (o WorkflowTaskCreateAirtableTableRecordMapOutput) MapIndex(k pulumi.StringInput) WorkflowTaskCreateAirtableTableRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTaskCreateAirtableTableRecord {
		return vs[0].(map[string]*WorkflowTaskCreateAirtableTableRecord)[vs[1].(string)]
	}).(WorkflowTaskCreateAirtableTableRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateAirtableTableRecordInput)(nil)).Elem(), &WorkflowTaskCreateAirtableTableRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateAirtableTableRecordArrayInput)(nil)).Elem(), WorkflowTaskCreateAirtableTableRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTaskCreateAirtableTableRecordMapInput)(nil)).Elem(), WorkflowTaskCreateAirtableTableRecordMap{})
	pulumi.RegisterOutputType(WorkflowTaskCreateAirtableTableRecordOutput{})
	pulumi.RegisterOutputType(WorkflowTaskCreateAirtableTableRecordArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTaskCreateAirtableTableRecordMapOutput{})
}
