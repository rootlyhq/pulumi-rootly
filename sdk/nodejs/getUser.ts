// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getUser(args?: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getUser:getUser", {
        "createdAt": args.createdAt,
        "email": args.email,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: string};
    email?: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly createdAt?: {[key: string]: string};
    readonly email: string;
    /**
     * The ID of this resource.
     */
    readonly id: string;
}
/**
 * ## Example Usage
 */
export function getUserOutput(args?: GetUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserResult> {
    return pulumi.output(args).apply((a: any) => getUser(a, opts))
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserOutputArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    email?: pulumi.Input<string>;
}
