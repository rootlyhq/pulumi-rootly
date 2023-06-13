// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rootly from "@pulumi/rootly";
 *
 * const my_status_page = pulumi.output(rootly.getStatusPage({
 *     slug: "my-status-page",
 * }));
 * ```
 */
export function getStatusPage(args?: GetStatusPageArgs, opts?: pulumi.InvokeOptions): Promise<GetStatusPageResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("rootly:index/getStatusPage:getStatusPage", {
        "createdAt": args.createdAt,
    }, opts);
}

/**
 * A collection of arguments for invoking getStatusPage.
 */
export interface GetStatusPageArgs {
    createdAt?: {[key: string]: any};
}

/**
 * A collection of values returned by getStatusPage.
 */
export interface GetStatusPageResult {
    readonly createdAt?: {[key: string]: any};
    readonly id: string;
}

export function getStatusPageOutput(args?: GetStatusPageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStatusPageResult> {
    return pulumi.output(args).apply(a => getStatusPage(a, opts))
}

/**
 * A collection of arguments for invoking getStatusPage.
 */
export interface GetStatusPageOutputArgs {
    createdAt?: pulumi.Input<{[key: string]: any}>;
}