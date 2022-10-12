// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getService(args?: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("rootly:index/getService:getService", {
        "backstageId": args.backstageId,
        "createdAt": args.createdAt,
        "name": args.name,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceArgs {
    backstageId?: string;
    createdAt?: {[key: string]: any};
    name?: string;
    slug?: string;
}

/**
 * A collection of values returned by getService.
 */
export interface GetServiceResult {
    readonly backstageId: string;
    readonly createdAt?: {[key: string]: any};
    readonly id: string;
    readonly name: string;
    readonly slug: string;
}

export function getServiceOutput(args?: GetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceResult> {
    return pulumi.output(args).apply(a => getService(a, opts))
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceOutputArgs {
    backstageId?: pulumi.Input<string>;
    createdAt?: pulumi.Input<{[key: string]: any}>;
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
}