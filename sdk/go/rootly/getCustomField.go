// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rootly

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rootlyhq/pulumi-rootly/sdk/go/rootly/internal"
)

// DEPRECATED: Please use `formField` data source instead.
//
// ## Example Usage
func LookupCustomField(ctx *pulumi.Context, args *LookupCustomFieldArgs, opts ...pulumi.InvokeOption) (*LookupCustomFieldResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomFieldResult
	err := ctx.Invoke("rootly:index/getCustomField:getCustomField", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomField.
type LookupCustomFieldArgs struct {
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt map[string]string `pulumi:"createdAt"`
	Enabled   *bool             `pulumi:"enabled"`
	Kind      *string           `pulumi:"kind"`
	Label     *string           `pulumi:"label"`
	Slug      *string           `pulumi:"slug"`
}

// A collection of values returned by getCustomField.
type LookupCustomFieldResult struct {
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt map[string]string `pulumi:"createdAt"`
	Enabled   *bool             `pulumi:"enabled"`
	// The ID of this resource.
	Id    string `pulumi:"id"`
	Kind  string `pulumi:"kind"`
	Label string `pulumi:"label"`
	Slug  string `pulumi:"slug"`
}

func LookupCustomFieldOutput(ctx *pulumi.Context, args LookupCustomFieldOutputArgs, opts ...pulumi.InvokeOption) LookupCustomFieldResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomFieldResult, error) {
			args := v.(LookupCustomFieldArgs)
			r, err := LookupCustomField(ctx, &args, opts...)
			var s LookupCustomFieldResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomFieldResultOutput)
}

// A collection of arguments for invoking getCustomField.
type LookupCustomFieldOutputArgs struct {
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt pulumi.StringMapInput `pulumi:"createdAt"`
	Enabled   pulumi.BoolPtrInput   `pulumi:"enabled"`
	Kind      pulumi.StringPtrInput `pulumi:"kind"`
	Label     pulumi.StringPtrInput `pulumi:"label"`
	Slug      pulumi.StringPtrInput `pulumi:"slug"`
}

func (LookupCustomFieldOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomFieldArgs)(nil)).Elem()
}

// A collection of values returned by getCustomField.
type LookupCustomFieldResultOutput struct{ *pulumi.OutputState }

func (LookupCustomFieldResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomFieldResult)(nil)).Elem()
}

func (o LookupCustomFieldResultOutput) ToLookupCustomFieldResultOutput() LookupCustomFieldResultOutput {
	return o
}

func (o LookupCustomFieldResultOutput) ToLookupCustomFieldResultOutputWithContext(ctx context.Context) LookupCustomFieldResultOutput {
	return o
}

// Filter by date range using 'lt' and 'gt'.
func (o LookupCustomFieldResultOutput) CreatedAt() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomFieldResult) map[string]string { return v.CreatedAt }).(pulumi.StringMapOutput)
}

func (o LookupCustomFieldResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCustomFieldResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID of this resource.
func (o LookupCustomFieldResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomFieldResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomFieldResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomFieldResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupCustomFieldResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomFieldResult) string { return v.Label }).(pulumi.StringOutput)
}

func (o LookupCustomFieldResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomFieldResult) string { return v.Slug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomFieldResultOutput{})
}
