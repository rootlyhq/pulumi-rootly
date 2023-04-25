// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getWebhooksEndpoint(args?: GetWebhooksEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetWebhooksEndpointResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("rootly:index/getWebhooksEndpoint:getWebhooksEndpoint", {
        "name": args.name,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getWebhooksEndpoint.
 */
export interface GetWebhooksEndpointArgs {
    name?: string;
    slug?: string;
}

/**
 * A collection of values returned by getWebhooksEndpoint.
 */
export interface GetWebhooksEndpointResult {
    readonly id: string;
    readonly name: string;
    readonly slug: string;
}

export function getWebhooksEndpointOutput(args?: GetWebhooksEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWebhooksEndpointResult> {
    return pulumi.output(args).apply(a => getWebhooksEndpoint(a, opts))
}

/**
 * A collection of arguments for invoking getWebhooksEndpoint.
 */
export interface GetWebhooksEndpointOutputArgs {
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
}