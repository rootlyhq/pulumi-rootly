// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getIncidentPermissionSet(args?: GetIncidentPermissionSetArgs, opts?: pulumi.InvokeOptions): Promise<GetIncidentPermissionSetResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getIncidentPermissionSet:getIncidentPermissionSet", {
        "createdAt": args.createdAt,
        "name": args.name,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getIncidentPermissionSet.
 */
export interface GetIncidentPermissionSetArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: string};
    name?: string;
    slug?: string;
}

/**
 * A collection of values returned by getIncidentPermissionSet.
 */
export interface GetIncidentPermissionSetResult {
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
export function getIncidentPermissionSetOutput(args?: GetIncidentPermissionSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIncidentPermissionSetResult> {
    return pulumi.output(args).apply((a: any) => getIncidentPermissionSet(a, opts))
}

/**
 * A collection of arguments for invoking getIncidentPermissionSet.
 */
export interface GetIncidentPermissionSetOutputArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
}
