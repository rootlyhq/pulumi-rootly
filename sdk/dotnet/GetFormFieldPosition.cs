// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    public static class GetFormFieldPosition
    {
        public static Task<GetFormFieldPositionResult> InvokeAsync(GetFormFieldPositionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFormFieldPositionResult>("rootly:index/getFormFieldPosition:getFormFieldPosition", args ?? new GetFormFieldPositionArgs(), options.WithDefaults());

        public static Output<GetFormFieldPositionResult> Invoke(GetFormFieldPositionInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFormFieldPositionResult>("rootly:index/getFormFieldPosition:getFormFieldPosition", args ?? new GetFormFieldPositionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFormFieldPositionArgs : global::Pulumi.InvokeArgs
    {
        [Input("form")]
        public string? Form { get; set; }

        public GetFormFieldPositionArgs()
        {
        }
        public static new GetFormFieldPositionArgs Empty => new GetFormFieldPositionArgs();
    }

    public sealed class GetFormFieldPositionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("form")]
        public Input<string>? Form { get; set; }

        public GetFormFieldPositionInvokeArgs()
        {
        }
        public static new GetFormFieldPositionInvokeArgs Empty => new GetFormFieldPositionInvokeArgs();
    }


    [OutputType]
    public sealed class GetFormFieldPositionResult
    {
        public readonly string Form;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetFormFieldPositionResult(
            string form,

            string id)
        {
            Form = form;
            Id = id;
        }
    }
}
