// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rootly

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rootlyhq/pulumi-rootly/sdk/go/rootly/internal"
)

func GetCauses(ctx *pulumi.Context, args *GetCausesArgs, opts ...pulumi.InvokeOption) (*GetCausesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCausesResult
	err := ctx.Invoke("rootly:index/getCauses:getCauses", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCauses.
type GetCausesArgs struct {
	Name *string `pulumi:"name"`
	Slug *string `pulumi:"slug"`
}

// A collection of values returned by getCauses.
type GetCausesResult struct {
	Causes []GetCausesCause `pulumi:"causes"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	Slug *string `pulumi:"slug"`
}

func GetCausesOutput(ctx *pulumi.Context, args GetCausesOutputArgs, opts ...pulumi.InvokeOption) GetCausesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCausesResult, error) {
			args := v.(GetCausesArgs)
			r, err := GetCauses(ctx, &args, opts...)
			var s GetCausesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCausesResultOutput)
}

// A collection of arguments for invoking getCauses.
type GetCausesOutputArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Slug pulumi.StringPtrInput `pulumi:"slug"`
}

func (GetCausesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCausesArgs)(nil)).Elem()
}

// A collection of values returned by getCauses.
type GetCausesResultOutput struct{ *pulumi.OutputState }

func (GetCausesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCausesResult)(nil)).Elem()
}

func (o GetCausesResultOutput) ToGetCausesResultOutput() GetCausesResultOutput {
	return o
}

func (o GetCausesResultOutput) ToGetCausesResultOutputWithContext(ctx context.Context) GetCausesResultOutput {
	return o
}

func (o GetCausesResultOutput) Causes() GetCausesCauseArrayOutput {
	return o.ApplyT(func(v GetCausesResult) []GetCausesCause { return v.Causes }).(GetCausesCauseArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCausesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCausesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCausesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCausesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetCausesResultOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCausesResult) *string { return v.Slug }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCausesResultOutput{})
}
