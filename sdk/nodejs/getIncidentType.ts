// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getIncidentType(args?: GetIncidentTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetIncidentTypeResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getIncidentType:getIncidentType", {
        "color": args.color,
        "createdAt": args.createdAt,
        "name": args.name,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getIncidentType.
 */
export interface GetIncidentTypeArgs {
    color?: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: any};
    name?: string;
    slug?: string;
}

/**
 * A collection of values returned by getIncidentType.
 */
export interface GetIncidentTypeResult {
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
export function getIncidentTypeOutput(args?: GetIncidentTypeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIncidentTypeResult> {
    return pulumi.output(args).apply((a: any) => getIncidentType(a, opts))
}

/**
 * A collection of arguments for invoking getIncidentType.
 */
export interface GetIncidentTypeOutputArgs {
    color?: pulumi.Input<string>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: any}>;
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
}
