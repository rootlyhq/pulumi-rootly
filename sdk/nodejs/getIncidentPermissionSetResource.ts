// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getIncidentPermissionSetResource(args?: GetIncidentPermissionSetResourceArgs, opts?: pulumi.InvokeOptions): Promise<GetIncidentPermissionSetResourceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getIncidentPermissionSetResource:getIncidentPermissionSetResource", {
        "createdAt": args.createdAt,
        "kind": args.kind,
    }, opts);
}

/**
 * A collection of arguments for invoking getIncidentPermissionSetResource.
 */
export interface GetIncidentPermissionSetResourceArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: any};
    kind?: string;
}

/**
 * A collection of values returned by getIncidentPermissionSetResource.
 */
export interface GetIncidentPermissionSetResourceResult {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly createdAt?: {[key: string]: any};
    /**
     * The ID of this resource.
     */
    readonly id: string;
    readonly kind: string;
}
export function getIncidentPermissionSetResourceOutput(args?: GetIncidentPermissionSetResourceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIncidentPermissionSetResourceResult> {
    return pulumi.output(args).apply((a: any) => getIncidentPermissionSetResource(a, opts))
}

/**
 * A collection of arguments for invoking getIncidentPermissionSetResource.
 */
export interface GetIncidentPermissionSetResourceOutputArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: any}>;
    kind?: pulumi.Input<string>;
}
