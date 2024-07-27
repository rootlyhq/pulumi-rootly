// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getFormFieldPosition(args?: GetFormFieldPositionArgs, opts?: pulumi.InvokeOptions): Promise<GetFormFieldPositionResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getFormFieldPosition:getFormFieldPosition", {
        "form": args.form,
    }, opts);
}

/**
 * A collection of arguments for invoking getFormFieldPosition.
 */
export interface GetFormFieldPositionArgs {
    form?: string;
}

/**
 * A collection of values returned by getFormFieldPosition.
 */
export interface GetFormFieldPositionResult {
    readonly form: string;
    /**
     * The ID of this resource.
     */
    readonly id: string;
}
export function getFormFieldPositionOutput(args?: GetFormFieldPositionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFormFieldPositionResult> {
    return pulumi.output(args).apply((a: any) => getFormFieldPosition(a, opts))
}

/**
 * A collection of arguments for invoking getFormFieldPosition.
 */
export interface GetFormFieldPositionOutputArgs {
    form?: pulumi.Input<string>;
}
