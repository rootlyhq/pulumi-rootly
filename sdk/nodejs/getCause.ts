// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getCause(args?: GetCauseArgs, opts?: pulumi.InvokeOptions): Promise<GetCauseResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getCause:getCause", {
        "createdAt": args.createdAt,
        "name": args.name,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getCause.
 */
export interface GetCauseArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: string};
    name?: string;
    slug?: string;
}

/**
 * A collection of values returned by getCause.
 */
export interface GetCauseResult {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly createdAt?: {[key: string]: string};
    /**
     * The ID of this resource.
     */
    readonly id: string;
    readonly name: string;
    readonly slug: string;
}
/**
 * ## Example Usage
 */
export function getCauseOutput(args?: GetCauseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCauseResult> {
    return pulumi.output(args).apply((a: any) => getCause(a, opts))
}

/**
 * A collection of arguments for invoking getCause.
 */
export interface GetCauseOutputArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
}
