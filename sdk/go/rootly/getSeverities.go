// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rootly

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rootlyhq/pulumi-rootly/sdk/go/rootly/internal"
)

func GetSeverities(ctx *pulumi.Context, args *GetSeveritiesArgs, opts ...pulumi.InvokeOption) (*GetSeveritiesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSeveritiesResult
	err := ctx.Invoke("rootly:index/getSeverities:getSeverities", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSeverities.
type GetSeveritiesArgs struct {
	Name *string `pulumi:"name"`
	Slug *string `pulumi:"slug"`
}

// A collection of values returned by getSeverities.
type GetSeveritiesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                  `pulumi:"id"`
	Name       *string                 `pulumi:"name"`
	Severities []GetSeveritiesSeverity `pulumi:"severities"`
	Slug       *string                 `pulumi:"slug"`
}

func GetSeveritiesOutput(ctx *pulumi.Context, args GetSeveritiesOutputArgs, opts ...pulumi.InvokeOption) GetSeveritiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSeveritiesResult, error) {
			args := v.(GetSeveritiesArgs)
			r, err := GetSeverities(ctx, &args, opts...)
			var s GetSeveritiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSeveritiesResultOutput)
}

// A collection of arguments for invoking getSeverities.
type GetSeveritiesOutputArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Slug pulumi.StringPtrInput `pulumi:"slug"`
}

func (GetSeveritiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSeveritiesArgs)(nil)).Elem()
}

// A collection of values returned by getSeverities.
type GetSeveritiesResultOutput struct{ *pulumi.OutputState }

func (GetSeveritiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSeveritiesResult)(nil)).Elem()
}

func (o GetSeveritiesResultOutput) ToGetSeveritiesResultOutput() GetSeveritiesResultOutput {
	return o
}

func (o GetSeveritiesResultOutput) ToGetSeveritiesResultOutputWithContext(ctx context.Context) GetSeveritiesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSeveritiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSeveritiesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSeveritiesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSeveritiesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetSeveritiesResultOutput) Severities() GetSeveritiesSeverityArrayOutput {
	return o.ApplyT(func(v GetSeveritiesResult) []GetSeveritiesSeverity { return v.Severities }).(GetSeveritiesSeverityArrayOutput)
}

func (o GetSeveritiesResultOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSeveritiesResult) *string { return v.Slug }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSeveritiesResultOutput{})
}
