// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rootly

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rootlyhq/pulumi-rootly/sdk/go/rootly/internal"
)

// ## Example Usage
type WorkflowPulse struct {
	pulumi.CustomResourceState

	CauseIds pulumi.StringArrayOutput `pulumi:"causeIds"`
	// Workflow command
	Command pulumi.StringOutput `pulumi:"command"`
	// This will notify you back when the workflow is starting. Value must be one of true or false
	CommandFeedbackEnabled pulumi.BoolOutput `pulumi:"commandFeedbackEnabled"`
	// The description of the workflow
	Description      pulumi.StringOutput      `pulumi:"description"`
	Enabled          pulumi.BoolPtrOutput     `pulumi:"enabled"`
	EnvironmentIds   pulumi.StringArrayOutput `pulumi:"environmentIds"`
	FunctionalityIds pulumi.StringArrayOutput `pulumi:"functionalityIds"`
	GroupIds         pulumi.StringArrayOutput `pulumi:"groupIds"`
	IncidentRoleIds  pulumi.StringArrayOutput `pulumi:"incidentRoleIds"`
	IncidentTypeIds  pulumi.StringArrayOutput `pulumi:"incidentTypeIds"`
	// Restricts workflow edits to admins when turned on. Only admins can set this field.. Value must be one of true or false
	Locked pulumi.BoolOutput `pulumi:"locked"`
	// The title of the workflow
	Name pulumi.StringOutput `pulumi:"name"`
	// The order which the workflow should run with other workflows.
	Position pulumi.IntOutput `pulumi:"position"`
	// Repeat workflow every duration
	RepeatEveryDuration pulumi.StringOutput `pulumi:"repeatEveryDuration"`
	// Repeat on weekdays. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
	RepeatOns   pulumi.StringArrayOutput `pulumi:"repeatOns"`
	ServiceIds  pulumi.StringArrayOutput `pulumi:"serviceIds"`
	SeverityIds pulumi.StringArrayOutput `pulumi:"severityIds"`
	// The slug of the workflow
	Slug          pulumi.StringOutput              `pulumi:"slug"`
	TriggerParams WorkflowPulseTriggerParamsOutput `pulumi:"triggerParams"`
	// Wait this duration before executing
	Wait pulumi.StringOutput `pulumi:"wait"`
	// The group this workflow belongs to.
	WorkflowGroupId pulumi.StringOutput `pulumi:"workflowGroupId"`
}

// NewWorkflowPulse registers a new resource with the given unique name, arguments, and options.
func NewWorkflowPulse(ctx *pulumi.Context,
	name string, args *WorkflowPulseArgs, opts ...pulumi.ResourceOption) (*WorkflowPulse, error) {
	if args == nil {
		args = &WorkflowPulseArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WorkflowPulse
	err := ctx.RegisterResource("rootly:index/workflowPulse:WorkflowPulse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowPulse gets an existing WorkflowPulse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowPulse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowPulseState, opts ...pulumi.ResourceOption) (*WorkflowPulse, error) {
	var resource WorkflowPulse
	err := ctx.ReadResource("rootly:index/workflowPulse:WorkflowPulse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowPulse resources.
type workflowPulseState struct {
	CauseIds []string `pulumi:"causeIds"`
	// Workflow command
	Command *string `pulumi:"command"`
	// This will notify you back when the workflow is starting. Value must be one of true or false
	CommandFeedbackEnabled *bool `pulumi:"commandFeedbackEnabled"`
	// The description of the workflow
	Description      *string  `pulumi:"description"`
	Enabled          *bool    `pulumi:"enabled"`
	EnvironmentIds   []string `pulumi:"environmentIds"`
	FunctionalityIds []string `pulumi:"functionalityIds"`
	GroupIds         []string `pulumi:"groupIds"`
	IncidentRoleIds  []string `pulumi:"incidentRoleIds"`
	IncidentTypeIds  []string `pulumi:"incidentTypeIds"`
	// Restricts workflow edits to admins when turned on. Only admins can set this field.. Value must be one of true or false
	Locked *bool `pulumi:"locked"`
	// The title of the workflow
	Name *string `pulumi:"name"`
	// The order which the workflow should run with other workflows.
	Position *int `pulumi:"position"`
	// Repeat workflow every duration
	RepeatEveryDuration *string `pulumi:"repeatEveryDuration"`
	// Repeat on weekdays. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
	RepeatOns   []string `pulumi:"repeatOns"`
	ServiceIds  []string `pulumi:"serviceIds"`
	SeverityIds []string `pulumi:"severityIds"`
	// The slug of the workflow
	Slug          *string                     `pulumi:"slug"`
	TriggerParams *WorkflowPulseTriggerParams `pulumi:"triggerParams"`
	// Wait this duration before executing
	Wait *string `pulumi:"wait"`
	// The group this workflow belongs to.
	WorkflowGroupId *string `pulumi:"workflowGroupId"`
}

type WorkflowPulseState struct {
	CauseIds pulumi.StringArrayInput
	// Workflow command
	Command pulumi.StringPtrInput
	// This will notify you back when the workflow is starting. Value must be one of true or false
	CommandFeedbackEnabled pulumi.BoolPtrInput
	// The description of the workflow
	Description      pulumi.StringPtrInput
	Enabled          pulumi.BoolPtrInput
	EnvironmentIds   pulumi.StringArrayInput
	FunctionalityIds pulumi.StringArrayInput
	GroupIds         pulumi.StringArrayInput
	IncidentRoleIds  pulumi.StringArrayInput
	IncidentTypeIds  pulumi.StringArrayInput
	// Restricts workflow edits to admins when turned on. Only admins can set this field.. Value must be one of true or false
	Locked pulumi.BoolPtrInput
	// The title of the workflow
	Name pulumi.StringPtrInput
	// The order which the workflow should run with other workflows.
	Position pulumi.IntPtrInput
	// Repeat workflow every duration
	RepeatEveryDuration pulumi.StringPtrInput
	// Repeat on weekdays. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
	RepeatOns   pulumi.StringArrayInput
	ServiceIds  pulumi.StringArrayInput
	SeverityIds pulumi.StringArrayInput
	// The slug of the workflow
	Slug          pulumi.StringPtrInput
	TriggerParams WorkflowPulseTriggerParamsPtrInput
	// Wait this duration before executing
	Wait pulumi.StringPtrInput
	// The group this workflow belongs to.
	WorkflowGroupId pulumi.StringPtrInput
}

func (WorkflowPulseState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowPulseState)(nil)).Elem()
}

type workflowPulseArgs struct {
	CauseIds []string `pulumi:"causeIds"`
	// Workflow command
	Command *string `pulumi:"command"`
	// This will notify you back when the workflow is starting. Value must be one of true or false
	CommandFeedbackEnabled *bool `pulumi:"commandFeedbackEnabled"`
	// The description of the workflow
	Description      *string  `pulumi:"description"`
	Enabled          *bool    `pulumi:"enabled"`
	EnvironmentIds   []string `pulumi:"environmentIds"`
	FunctionalityIds []string `pulumi:"functionalityIds"`
	GroupIds         []string `pulumi:"groupIds"`
	IncidentRoleIds  []string `pulumi:"incidentRoleIds"`
	IncidentTypeIds  []string `pulumi:"incidentTypeIds"`
	// Restricts workflow edits to admins when turned on. Only admins can set this field.. Value must be one of true or false
	Locked *bool `pulumi:"locked"`
	// The title of the workflow
	Name *string `pulumi:"name"`
	// The order which the workflow should run with other workflows.
	Position *int `pulumi:"position"`
	// Repeat workflow every duration
	RepeatEveryDuration *string `pulumi:"repeatEveryDuration"`
	// Repeat on weekdays. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
	RepeatOns   []string `pulumi:"repeatOns"`
	ServiceIds  []string `pulumi:"serviceIds"`
	SeverityIds []string `pulumi:"severityIds"`
	// The slug of the workflow
	Slug          *string                     `pulumi:"slug"`
	TriggerParams *WorkflowPulseTriggerParams `pulumi:"triggerParams"`
	// Wait this duration before executing
	Wait *string `pulumi:"wait"`
	// The group this workflow belongs to.
	WorkflowGroupId *string `pulumi:"workflowGroupId"`
}

// The set of arguments for constructing a WorkflowPulse resource.
type WorkflowPulseArgs struct {
	CauseIds pulumi.StringArrayInput
	// Workflow command
	Command pulumi.StringPtrInput
	// This will notify you back when the workflow is starting. Value must be one of true or false
	CommandFeedbackEnabled pulumi.BoolPtrInput
	// The description of the workflow
	Description      pulumi.StringPtrInput
	Enabled          pulumi.BoolPtrInput
	EnvironmentIds   pulumi.StringArrayInput
	FunctionalityIds pulumi.StringArrayInput
	GroupIds         pulumi.StringArrayInput
	IncidentRoleIds  pulumi.StringArrayInput
	IncidentTypeIds  pulumi.StringArrayInput
	// Restricts workflow edits to admins when turned on. Only admins can set this field.. Value must be one of true or false
	Locked pulumi.BoolPtrInput
	// The title of the workflow
	Name pulumi.StringPtrInput
	// The order which the workflow should run with other workflows.
	Position pulumi.IntPtrInput
	// Repeat workflow every duration
	RepeatEveryDuration pulumi.StringPtrInput
	// Repeat on weekdays. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
	RepeatOns   pulumi.StringArrayInput
	ServiceIds  pulumi.StringArrayInput
	SeverityIds pulumi.StringArrayInput
	// The slug of the workflow
	Slug          pulumi.StringPtrInput
	TriggerParams WorkflowPulseTriggerParamsPtrInput
	// Wait this duration before executing
	Wait pulumi.StringPtrInput
	// The group this workflow belongs to.
	WorkflowGroupId pulumi.StringPtrInput
}

func (WorkflowPulseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowPulseArgs)(nil)).Elem()
}

type WorkflowPulseInput interface {
	pulumi.Input

	ToWorkflowPulseOutput() WorkflowPulseOutput
	ToWorkflowPulseOutputWithContext(ctx context.Context) WorkflowPulseOutput
}

func (*WorkflowPulse) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowPulse)(nil)).Elem()
}

func (i *WorkflowPulse) ToWorkflowPulseOutput() WorkflowPulseOutput {
	return i.ToWorkflowPulseOutputWithContext(context.Background())
}

func (i *WorkflowPulse) ToWorkflowPulseOutputWithContext(ctx context.Context) WorkflowPulseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowPulseOutput)
}

// WorkflowPulseArrayInput is an input type that accepts WorkflowPulseArray and WorkflowPulseArrayOutput values.
// You can construct a concrete instance of `WorkflowPulseArrayInput` via:
//
//	WorkflowPulseArray{ WorkflowPulseArgs{...} }
type WorkflowPulseArrayInput interface {
	pulumi.Input

	ToWorkflowPulseArrayOutput() WorkflowPulseArrayOutput
	ToWorkflowPulseArrayOutputWithContext(context.Context) WorkflowPulseArrayOutput
}

type WorkflowPulseArray []WorkflowPulseInput

func (WorkflowPulseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowPulse)(nil)).Elem()
}

func (i WorkflowPulseArray) ToWorkflowPulseArrayOutput() WorkflowPulseArrayOutput {
	return i.ToWorkflowPulseArrayOutputWithContext(context.Background())
}

func (i WorkflowPulseArray) ToWorkflowPulseArrayOutputWithContext(ctx context.Context) WorkflowPulseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowPulseArrayOutput)
}

// WorkflowPulseMapInput is an input type that accepts WorkflowPulseMap and WorkflowPulseMapOutput values.
// You can construct a concrete instance of `WorkflowPulseMapInput` via:
//
//	WorkflowPulseMap{ "key": WorkflowPulseArgs{...} }
type WorkflowPulseMapInput interface {
	pulumi.Input

	ToWorkflowPulseMapOutput() WorkflowPulseMapOutput
	ToWorkflowPulseMapOutputWithContext(context.Context) WorkflowPulseMapOutput
}

type WorkflowPulseMap map[string]WorkflowPulseInput

func (WorkflowPulseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowPulse)(nil)).Elem()
}

func (i WorkflowPulseMap) ToWorkflowPulseMapOutput() WorkflowPulseMapOutput {
	return i.ToWorkflowPulseMapOutputWithContext(context.Background())
}

func (i WorkflowPulseMap) ToWorkflowPulseMapOutputWithContext(ctx context.Context) WorkflowPulseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowPulseMapOutput)
}

type WorkflowPulseOutput struct{ *pulumi.OutputState }

func (WorkflowPulseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowPulse)(nil)).Elem()
}

func (o WorkflowPulseOutput) ToWorkflowPulseOutput() WorkflowPulseOutput {
	return o
}

func (o WorkflowPulseOutput) ToWorkflowPulseOutputWithContext(ctx context.Context) WorkflowPulseOutput {
	return o
}

func (o WorkflowPulseOutput) CauseIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringArrayOutput { return v.CauseIds }).(pulumi.StringArrayOutput)
}

// Workflow command
func (o WorkflowPulseOutput) Command() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringOutput { return v.Command }).(pulumi.StringOutput)
}

// This will notify you back when the workflow is starting. Value must be one of true or false
func (o WorkflowPulseOutput) CommandFeedbackEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.BoolOutput { return v.CommandFeedbackEnabled }).(pulumi.BoolOutput)
}

// The description of the workflow
func (o WorkflowPulseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o WorkflowPulseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o WorkflowPulseOutput) EnvironmentIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringArrayOutput { return v.EnvironmentIds }).(pulumi.StringArrayOutput)
}

func (o WorkflowPulseOutput) FunctionalityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringArrayOutput { return v.FunctionalityIds }).(pulumi.StringArrayOutput)
}

func (o WorkflowPulseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o WorkflowPulseOutput) IncidentRoleIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringArrayOutput { return v.IncidentRoleIds }).(pulumi.StringArrayOutput)
}

func (o WorkflowPulseOutput) IncidentTypeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringArrayOutput { return v.IncidentTypeIds }).(pulumi.StringArrayOutput)
}

// Restricts workflow edits to admins when turned on. Only admins can set this field.. Value must be one of true or false
func (o WorkflowPulseOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.BoolOutput { return v.Locked }).(pulumi.BoolOutput)
}

// The title of the workflow
func (o WorkflowPulseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The order which the workflow should run with other workflows.
func (o WorkflowPulseOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Repeat workflow every duration
func (o WorkflowPulseOutput) RepeatEveryDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringOutput { return v.RepeatEveryDuration }).(pulumi.StringOutput)
}

// Repeat on weekdays. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
func (o WorkflowPulseOutput) RepeatOns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringArrayOutput { return v.RepeatOns }).(pulumi.StringArrayOutput)
}

func (o WorkflowPulseOutput) ServiceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringArrayOutput { return v.ServiceIds }).(pulumi.StringArrayOutput)
}

func (o WorkflowPulseOutput) SeverityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringArrayOutput { return v.SeverityIds }).(pulumi.StringArrayOutput)
}

// The slug of the workflow
func (o WorkflowPulseOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

func (o WorkflowPulseOutput) TriggerParams() WorkflowPulseTriggerParamsOutput {
	return o.ApplyT(func(v *WorkflowPulse) WorkflowPulseTriggerParamsOutput { return v.TriggerParams }).(WorkflowPulseTriggerParamsOutput)
}

// Wait this duration before executing
func (o WorkflowPulseOutput) Wait() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringOutput { return v.Wait }).(pulumi.StringOutput)
}

// The group this workflow belongs to.
func (o WorkflowPulseOutput) WorkflowGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowPulse) pulumi.StringOutput { return v.WorkflowGroupId }).(pulumi.StringOutput)
}

type WorkflowPulseArrayOutput struct{ *pulumi.OutputState }

func (WorkflowPulseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowPulse)(nil)).Elem()
}

func (o WorkflowPulseArrayOutput) ToWorkflowPulseArrayOutput() WorkflowPulseArrayOutput {
	return o
}

func (o WorkflowPulseArrayOutput) ToWorkflowPulseArrayOutputWithContext(ctx context.Context) WorkflowPulseArrayOutput {
	return o
}

func (o WorkflowPulseArrayOutput) Index(i pulumi.IntInput) WorkflowPulseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowPulse {
		return vs[0].([]*WorkflowPulse)[vs[1].(int)]
	}).(WorkflowPulseOutput)
}

type WorkflowPulseMapOutput struct{ *pulumi.OutputState }

func (WorkflowPulseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowPulse)(nil)).Elem()
}

func (o WorkflowPulseMapOutput) ToWorkflowPulseMapOutput() WorkflowPulseMapOutput {
	return o
}

func (o WorkflowPulseMapOutput) ToWorkflowPulseMapOutputWithContext(ctx context.Context) WorkflowPulseMapOutput {
	return o
}

func (o WorkflowPulseMapOutput) MapIndex(k pulumi.StringInput) WorkflowPulseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowPulse {
		return vs[0].(map[string]*WorkflowPulse)[vs[1].(string)]
	}).(WorkflowPulseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowPulseInput)(nil)).Elem(), &WorkflowPulse{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowPulseArrayInput)(nil)).Elem(), WorkflowPulseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowPulseMapInput)(nil)).Elem(), WorkflowPulseMap{})
	pulumi.RegisterOutputType(WorkflowPulseOutput{})
	pulumi.RegisterOutputType(WorkflowPulseArrayOutput{})
	pulumi.RegisterOutputType(WorkflowPulseMapOutput{})
}
