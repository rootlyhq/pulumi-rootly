// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getTeam(args?: GetTeamArgs, opts?: pulumi.InvokeOptions): Promise<GetTeamResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getTeam:getTeam", {
        "color": args.color,
        "createdAt": args.createdAt,
        "name": args.name,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamArgs {
    color?: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: any};
    name?: string;
    slug?: string;
}

/**
 * A collection of values returned by getTeam.
 */
export interface GetTeamResult {
    readonly color: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly createdAt?: {[key: string]: any};
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
export function getTeamOutput(args?: GetTeamOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTeamResult> {
    return pulumi.output(args).apply((a: any) => getTeam(a, opts))
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamOutputArgs {
    color?: pulumi.Input<string>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: any}>;
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
}
