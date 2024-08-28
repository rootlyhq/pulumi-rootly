// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getIncidentPermissionSetBoolean(args?: GetIncidentPermissionSetBooleanArgs, opts?: pulumi.InvokeOptions): Promise<GetIncidentPermissionSetBooleanResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getIncidentPermissionSetBoolean:getIncidentPermissionSetBoolean", {
        "createdAt": args.createdAt,
        "kind": args.kind,
    }, opts);
}

/**
 * A collection of arguments for invoking getIncidentPermissionSetBoolean.
 */
export interface GetIncidentPermissionSetBooleanArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: string};
    kind?: string;
}

/**
 * A collection of values returned by getIncidentPermissionSetBoolean.
 */
export interface GetIncidentPermissionSetBooleanResult {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly createdAt?: {[key: string]: string};
    /**
     * The ID of this resource.
     */
    readonly id: string;
    readonly kind: string;
}
export function getIncidentPermissionSetBooleanOutput(args?: GetIncidentPermissionSetBooleanOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIncidentPermissionSetBooleanResult> {
    return pulumi.output(args).apply((a: any) => getIncidentPermissionSetBoolean(a, opts))
}

/**
 * A collection of arguments for invoking getIncidentPermissionSetBoolean.
 */
export interface GetIncidentPermissionSetBooleanOutputArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    kind?: pulumi.Input<string>;
}
