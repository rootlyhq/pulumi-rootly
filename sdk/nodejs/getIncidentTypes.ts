// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getIncidentTypes(args?: GetIncidentTypesArgs, opts?: pulumi.InvokeOptions): Promise<GetIncidentTypesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("rootly:index/getIncidentTypes:getIncidentTypes", {
        "name": args.name,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getIncidentTypes.
 */
export interface GetIncidentTypesArgs {
    name?: string;
    slug?: string;
}

/**
 * A collection of values returned by getIncidentTypes.
 */
export interface GetIncidentTypesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly incidentTypes: outputs.GetIncidentTypesIncidentType[];
    readonly name?: string;
    readonly slug?: string;
}

export function getIncidentTypesOutput(args?: GetIncidentTypesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIncidentTypesResult> {
    return pulumi.output(args).apply(a => getIncidentTypes(a, opts))
}

/**
 * A collection of arguments for invoking getIncidentTypes.
 */
export interface GetIncidentTypesOutputArgs {
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
}