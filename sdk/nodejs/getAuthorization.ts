// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAuthorization(args?: GetAuthorizationArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthorizationResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getAuthorization:getAuthorization", {
        "authorizableId": args.authorizableId,
        "authorizableType": args.authorizableType,
        "createdAt": args.createdAt,
        "granteeId": args.granteeId,
        "granteeType": args.granteeType,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthorization.
 */
export interface GetAuthorizationArgs {
    authorizableId?: string;
    authorizableType?: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: string};
    granteeId?: string;
    granteeType?: string;
}

/**
 * A collection of values returned by getAuthorization.
 */
export interface GetAuthorizationResult {
    readonly authorizableId: string;
    readonly authorizableType: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly createdAt?: {[key: string]: string};
    readonly granteeId: string;
    readonly granteeType: string;
    /**
     * The ID of this resource.
     */
    readonly id: string;
}
export function getAuthorizationOutput(args?: GetAuthorizationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAuthorizationResult> {
    return pulumi.output(args).apply((a: any) => getAuthorization(a, opts))
}

/**
 * A collection of arguments for invoking getAuthorization.
 */
export interface GetAuthorizationOutputArgs {
    authorizableId?: pulumi.Input<string>;
    authorizableType?: pulumi.Input<string>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    granteeId?: pulumi.Input<string>;
    granteeType?: pulumi.Input<string>;
}
