// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rootly

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rootlyhq/pulumi-rootly/sdk/go/rootly/internal"
)

func LookupIncidentPermissionSet(ctx *pulumi.Context, args *LookupIncidentPermissionSetArgs, opts ...pulumi.InvokeOption) (*LookupIncidentPermissionSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIncidentPermissionSetResult
	err := ctx.Invoke("rootly:index/getIncidentPermissionSet:getIncidentPermissionSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIncidentPermissionSet.
type LookupIncidentPermissionSetArgs struct {
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt map[string]string `pulumi:"createdAt"`
	Name      *string           `pulumi:"name"`
	Slug      *string           `pulumi:"slug"`
}

// A collection of values returned by getIncidentPermissionSet.
type LookupIncidentPermissionSetResult struct {
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt map[string]string `pulumi:"createdAt"`
	// The ID of this resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	Slug string `pulumi:"slug"`
}

func LookupIncidentPermissionSetOutput(ctx *pulumi.Context, args LookupIncidentPermissionSetOutputArgs, opts ...pulumi.InvokeOption) LookupIncidentPermissionSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIncidentPermissionSetResult, error) {
			args := v.(LookupIncidentPermissionSetArgs)
			r, err := LookupIncidentPermissionSet(ctx, &args, opts...)
			var s LookupIncidentPermissionSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIncidentPermissionSetResultOutput)
}

// A collection of arguments for invoking getIncidentPermissionSet.
type LookupIncidentPermissionSetOutputArgs struct {
	// Filter by date range using 'lt' and 'gt'.
	CreatedAt pulumi.StringMapInput `pulumi:"createdAt"`
	Name      pulumi.StringPtrInput `pulumi:"name"`
	Slug      pulumi.StringPtrInput `pulumi:"slug"`
}

func (LookupIncidentPermissionSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentPermissionSetArgs)(nil)).Elem()
}

// A collection of values returned by getIncidentPermissionSet.
type LookupIncidentPermissionSetResultOutput struct{ *pulumi.OutputState }

func (LookupIncidentPermissionSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentPermissionSetResult)(nil)).Elem()
}

func (o LookupIncidentPermissionSetResultOutput) ToLookupIncidentPermissionSetResultOutput() LookupIncidentPermissionSetResultOutput {
	return o
}

func (o LookupIncidentPermissionSetResultOutput) ToLookupIncidentPermissionSetResultOutputWithContext(ctx context.Context) LookupIncidentPermissionSetResultOutput {
	return o
}

// Filter by date range using 'lt' and 'gt'.
func (o LookupIncidentPermissionSetResultOutput) CreatedAt() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIncidentPermissionSetResult) map[string]string { return v.CreatedAt }).(pulumi.StringMapOutput)
}

// The ID of this resource.
func (o LookupIncidentPermissionSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentPermissionSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIncidentPermissionSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentPermissionSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIncidentPermissionSetResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentPermissionSetResult) string { return v.Slug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIncidentPermissionSetResultOutput{})
}
