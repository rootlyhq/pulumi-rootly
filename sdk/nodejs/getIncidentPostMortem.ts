// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getIncidentPostMortem(args?: GetIncidentPostMortemArgs, opts?: pulumi.InvokeOptions): Promise<GetIncidentPostMortemResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getIncidentPostMortem:getIncidentPostMortem", {
        "createdAt": args.createdAt,
        "mitigatedAt": args.mitigatedAt,
        "resolvedAt": args.resolvedAt,
        "startedAt": args.startedAt,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getIncidentPostMortem.
 */
export interface GetIncidentPostMortemArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    mitigatedAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    resolvedAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    startedAt?: {[key: string]: string};
    status?: string;
}

/**
 * A collection of values returned by getIncidentPostMortem.
 */
export interface GetIncidentPostMortemResult {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly createdAt?: {[key: string]: string};
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly mitigatedAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly resolvedAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly startedAt?: {[key: string]: string};
    readonly status: string;
}
/**
 * ## Example Usage
 */
export function getIncidentPostMortemOutput(args?: GetIncidentPostMortemOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIncidentPostMortemResult> {
    return pulumi.output(args).apply((a: any) => getIncidentPostMortem(a, opts))
}

/**
 * A collection of arguments for invoking getIncidentPostMortem.
 */
export interface GetIncidentPostMortemOutputArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    mitigatedAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    resolvedAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    startedAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    status?: pulumi.Input<string>;
}
