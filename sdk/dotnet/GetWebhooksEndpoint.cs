// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    public static class GetWebhooksEndpoint
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```shell
        /// data "rootly_webhooks_endpoint" "my-webhook-endpoint" {
        ///   slug = "my-webhookd-endpoint"
        /// }
        /// ```
        /// </summary>
        public static Task<GetWebhooksEndpointResult> InvokeAsync(GetWebhooksEndpointArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebhooksEndpointResult>("rootly:index/getWebhooksEndpoint:getWebhooksEndpoint", args ?? new GetWebhooksEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```shell
        /// data "rootly_webhooks_endpoint" "my-webhook-endpoint" {
        ///   slug = "my-webhookd-endpoint"
        /// }
        /// ```
        /// </summary>
        public static Output<GetWebhooksEndpointResult> Invoke(GetWebhooksEndpointInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebhooksEndpointResult>("rootly:index/getWebhooksEndpoint:getWebhooksEndpoint", args ?? new GetWebhooksEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebhooksEndpointArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("slug")]
        public string? Slug { get; set; }

        public GetWebhooksEndpointArgs()
        {
        }
        public static new GetWebhooksEndpointArgs Empty => new GetWebhooksEndpointArgs();
    }

    public sealed class GetWebhooksEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("slug")]
        public Input<string>? Slug { get; set; }

        public GetWebhooksEndpointInvokeArgs()
        {
        }
        public static new GetWebhooksEndpointInvokeArgs Empty => new GetWebhooksEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebhooksEndpointResult
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Slug;

        [OutputConstructor]
        private GetWebhooksEndpointResult(
            string id,

            string name,

            string slug)
        {
            Id = id;
            Name = name;
            Slug = slug;
        }
    }
}
