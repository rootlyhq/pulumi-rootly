// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getFormFieldPlacementCondition(args?: GetFormFieldPlacementConditionArgs, opts?: pulumi.InvokeOptions): Promise<GetFormFieldPlacementConditionResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getFormFieldPlacementCondition:getFormFieldPlacementCondition", {
        "formFieldId": args.formFieldId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFormFieldPlacementCondition.
 */
export interface GetFormFieldPlacementConditionArgs {
    formFieldId?: string;
}

/**
 * A collection of values returned by getFormFieldPlacementCondition.
 */
export interface GetFormFieldPlacementConditionResult {
    readonly formFieldId: string;
    /**
     * The ID of this resource.
     */
    readonly id: string;
}
export function getFormFieldPlacementConditionOutput(args?: GetFormFieldPlacementConditionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFormFieldPlacementConditionResult> {
    return pulumi.output(args).apply((a: any) => getFormFieldPlacementCondition(a, opts))
}

/**
 * A collection of arguments for invoking getFormFieldPlacementCondition.
 */
export interface GetFormFieldPlacementConditionOutputArgs {
    formFieldId?: pulumi.Input<string>;
}
