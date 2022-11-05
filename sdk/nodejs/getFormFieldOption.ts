// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getFormFieldOption(args?: GetFormFieldOptionArgs, opts?: pulumi.InvokeOptions): Promise<GetFormFieldOptionResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("rootly:index/getFormFieldOption:getFormFieldOption", {
        "color": args.color,
        "value": args.value,
    }, opts);
}

/**
 * A collection of arguments for invoking getFormFieldOption.
 */
export interface GetFormFieldOptionArgs {
    color?: string;
    value?: string;
}

/**
 * A collection of values returned by getFormFieldOption.
 */
export interface GetFormFieldOptionResult {
    readonly color: string;
    readonly id: string;
    readonly value: string;
}

export function getFormFieldOptionOutput(args?: GetFormFieldOptionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFormFieldOptionResult> {
    return pulumi.output(args).apply(a => getFormFieldOption(a, opts))
}

/**
 * A collection of arguments for invoking getFormFieldOption.
 */
export interface GetFormFieldOptionOutputArgs {
    color?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}