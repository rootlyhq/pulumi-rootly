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
func LookupSeverity(ctx *pulumi.Context, args *LookupSeverityArgs, opts ...pulumi.InvokeOption) (*LookupSeverityResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSeverityResult
	err := ctx.Invoke("rootly:index/getSeverity:getSeverity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSeverity.
type LookupSeverityArgs struct {
	Color *string `pulumi:"color"`
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt map[string]string `pulumi:"createdAt"`
	Name      *string           `pulumi:"name"`
	Severity  *string           `pulumi:"severity"`
	Slug      *string           `pulumi:"slug"`
}

// A collection of values returned by getSeverity.
type LookupSeverityResult struct {
	Color string `pulumi:"color"`
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt map[string]string `pulumi:"createdAt"`
	// The ID of this resource.
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	Severity string `pulumi:"severity"`
	Slug     string `pulumi:"slug"`
}

func LookupSeverityOutput(ctx *pulumi.Context, args LookupSeverityOutputArgs, opts ...pulumi.InvokeOption) LookupSeverityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSeverityResult, error) {
			args := v.(LookupSeverityArgs)
			r, err := LookupSeverity(ctx, &args, opts...)
			var s LookupSeverityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSeverityResultOutput)
}

// A collection of arguments for invoking getSeverity.
type LookupSeverityOutputArgs struct {
	Color pulumi.StringPtrInput `pulumi:"color"`
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt pulumi.StringMapInput `pulumi:"createdAt"`
	Name      pulumi.StringPtrInput `pulumi:"name"`
	Severity  pulumi.StringPtrInput `pulumi:"severity"`
	Slug      pulumi.StringPtrInput `pulumi:"slug"`
}

func (LookupSeverityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSeverityArgs)(nil)).Elem()
}

// A collection of values returned by getSeverity.
type LookupSeverityResultOutput struct{ *pulumi.OutputState }

func (LookupSeverityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSeverityResult)(nil)).Elem()
}

func (o LookupSeverityResultOutput) ToLookupSeverityResultOutput() LookupSeverityResultOutput {
	return o
}

func (o LookupSeverityResultOutput) ToLookupSeverityResultOutputWithContext(ctx context.Context) LookupSeverityResultOutput {
	return o
}

func (o LookupSeverityResultOutput) Color() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSeverityResult) string { return v.Color }).(pulumi.StringOutput)
}

// Filter by date range using 'lt' and 'gt'.
func (o LookupSeverityResultOutput) CreatedAt() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSeverityResult) map[string]string { return v.CreatedAt }).(pulumi.StringMapOutput)
}

// The ID of this resource.
func (o LookupSeverityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSeverityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSeverityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSeverityResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSeverityResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSeverityResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupSeverityResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSeverityResult) string { return v.Slug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSeverityResultOutput{})
}
