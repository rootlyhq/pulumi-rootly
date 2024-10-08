// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getStatusPage(args?: GetStatusPageArgs, opts?: pulumi.InvokeOptions): Promise<GetStatusPageResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getStatusPage:getStatusPage", {
        "createdAt": args.createdAt,
    }, opts);
}

/**
 * A collection of arguments for invoking getStatusPage.
 */
export interface GetStatusPageArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: string};
}

/**
 * A collection of values returned by getStatusPage.
 */
export interface GetStatusPageResult {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly createdAt?: {[key: string]: string};
    /**
     * The ID of this resource.
     */
    readonly id: string;
}
/**
 * ## Example Usage
 */
export function getStatusPageOutput(args?: GetStatusPageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStatusPageResult> {
    return pulumi.output(args).apply((a: any) => getStatusPage(a, opts))
}

/**
 * A collection of arguments for invoking getStatusPage.
 */
export interface GetStatusPageOutputArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
