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
 * const my-form-field = rootly.getFormField({
 *     slug: "my-form-field",
 * });
 * ```
 */
export function getFormField(args?: GetFormFieldArgs, opts?: pulumi.InvokeOptions): Promise<GetFormFieldResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getFormField:getFormField", {
        "createdAt": args.createdAt,
        "enabled": args.enabled,
        "kind": args.kind,
        "name": args.name,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getFormField.
 */
export interface GetFormFieldArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: any};
    enabled?: boolean;
    kind?: string;
    name?: string;
    slug?: string;
}

/**
 * A collection of values returned by getFormField.
 */
export interface GetFormFieldResult {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly createdAt?: {[key: string]: any};
    readonly enabled?: boolean;
    /**
     * The ID of this resource.
     */
    readonly id: string;
    readonly kind: string;
    readonly name: string;
    readonly slug: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rootly from "@pulumi/rootly";
 *
 * const my-form-field = rootly.getFormField({
 *     slug: "my-form-field",
 * });
 * ```
 */
export function getFormFieldOutput(args?: GetFormFieldOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFormFieldResult> {
    return pulumi.output(args).apply((a: any) => getFormField(a, opts))
}

/**
 * A collection of arguments for invoking getFormField.
 */
export interface GetFormFieldOutputArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: any}>;
    enabled?: pulumi.Input<boolean>;
    kind?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
}
