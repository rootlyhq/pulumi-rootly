// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rootly

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rootlyhq/pulumi-rootly/sdk/go/rootly/internal"
)

func LookupCustomForm(ctx *pulumi.Context, args *LookupCustomFormArgs, opts ...pulumi.InvokeOption) (*LookupCustomFormResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomFormResult
	err := ctx.Invoke("rootly:index/getCustomForm:getCustomForm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomForm.
type LookupCustomFormArgs struct {
	Command *string `pulumi:"command"`
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt map[string]string `pulumi:"createdAt"`
	Name      *string           `pulumi:"name"`
	Slug      *string           `pulumi:"slug"`
}

// A collection of values returned by getCustomForm.
type LookupCustomFormResult struct {
	Command string `pulumi:"command"`
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt map[string]string `pulumi:"createdAt"`
	// The ID of this resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	Slug string `pulumi:"slug"`
}

func LookupCustomFormOutput(ctx *pulumi.Context, args LookupCustomFormOutputArgs, opts ...pulumi.InvokeOption) LookupCustomFormResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomFormResult, error) {
			args := v.(LookupCustomFormArgs)
			r, err := LookupCustomForm(ctx, &args, opts...)
			var s LookupCustomFormResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomFormResultOutput)
}

// A collection of arguments for invoking getCustomForm.
type LookupCustomFormOutputArgs struct {
	Command pulumi.StringPtrInput `pulumi:"command"`
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt pulumi.StringMapInput `pulumi:"createdAt"`
	Name      pulumi.StringPtrInput `pulumi:"name"`
	Slug      pulumi.StringPtrInput `pulumi:"slug"`
}

func (LookupCustomFormOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomFormArgs)(nil)).Elem()
}

// A collection of values returned by getCustomForm.
type LookupCustomFormResultOutput struct{ *pulumi.OutputState }

func (LookupCustomFormResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomFormResult)(nil)).Elem()
}

func (o LookupCustomFormResultOutput) ToLookupCustomFormResultOutput() LookupCustomFormResultOutput {
	return o
}

func (o LookupCustomFormResultOutput) ToLookupCustomFormResultOutputWithContext(ctx context.Context) LookupCustomFormResultOutput {
	return o
}

func (o LookupCustomFormResultOutput) Command() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomFormResult) string { return v.Command }).(pulumi.StringOutput)
}

// Filter by date range using 'lt' and 'gt'.
func (o LookupCustomFormResultOutput) CreatedAt() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomFormResult) map[string]string { return v.CreatedAt }).(pulumi.StringMapOutput)
}

// The ID of this resource.
func (o LookupCustomFormResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomFormResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomFormResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomFormResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomFormResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomFormResult) string { return v.Slug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomFormResultOutput{})
}
