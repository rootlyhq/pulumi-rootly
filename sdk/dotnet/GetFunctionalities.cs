// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    public static class GetFunctionalities
    {
        public static Task<GetFunctionalitiesResult> InvokeAsync(GetFunctionalitiesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFunctionalitiesResult>("rootly:index/getFunctionalities:getFunctionalities", args ?? new GetFunctionalitiesArgs(), options.WithDefaults());

        public static Output<GetFunctionalitiesResult> Invoke(GetFunctionalitiesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFunctionalitiesResult>("rootly:index/getFunctionalities:getFunctionalities", args ?? new GetFunctionalitiesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionalitiesArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("opsgenieId")]
        public string? OpsgenieId { get; set; }

        [Input("pagerdutyId")]
        public string? PagerdutyId { get; set; }

        [Input("slug")]
        public string? Slug { get; set; }

        public GetFunctionalitiesArgs()
        {
        }
        public static new GetFunctionalitiesArgs Empty => new GetFunctionalitiesArgs();
    }

    public sealed class GetFunctionalitiesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("opsgenieId")]
        public Input<string>? OpsgenieId { get; set; }

        [Input("pagerdutyId")]
        public Input<string>? PagerdutyId { get; set; }

        [Input("slug")]
        public Input<string>? Slug { get; set; }

        public GetFunctionalitiesInvokeArgs()
        {
        }
        public static new GetFunctionalitiesInvokeArgs Empty => new GetFunctionalitiesInvokeArgs();
    }


    [OutputType]
    public sealed class GetFunctionalitiesResult
    {
        public readonly ImmutableArray<Outputs.GetFunctionalitiesFunctionalityResult> Functionalities;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? OpsgenieId;
        public readonly string? PagerdutyId;
        public readonly string? Slug;

        [OutputConstructor]
        private GetFunctionalitiesResult(
            ImmutableArray<Outputs.GetFunctionalitiesFunctionalityResult> functionalities,

            string id,

            string? name,

            string? opsgenieId,

            string? pagerdutyId,

            string? slug)
        {
            Functionalities = functionalities;
            Id = id;
            Name = name;
            OpsgenieId = opsgenieId;
            PagerdutyId = pagerdutyId;
            Slug = slug;
        }
    }
}
