// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getIncident(args?: GetIncidentArgs, opts?: pulumi.InvokeOptions): Promise<GetIncidentResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rootly:index/getIncident:getIncident", {
        "acknowledgedAt": args.acknowledgedAt,
        "createdAt": args.createdAt,
        "detectedAt": args.detectedAt,
        "environments": args.environments,
        "functionalities": args.functionalities,
        "inTriageAt": args.inTriageAt,
        "kind": args.kind,
        "labels": args.labels,
        "mitigatedAt": args.mitigatedAt,
        "private": args.private,
        "resolvedAt": args.resolvedAt,
        "services": args.services,
        "severity": args.severity,
        "startedAt": args.startedAt,
        "updatedAt": args.updatedAt,
    }, opts);
}

/**
 * A collection of arguments for invoking getIncident.
 */
export interface GetIncidentArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    acknowledgedAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    detectedAt?: {[key: string]: string};
    environments?: string;
    functionalities?: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    inTriageAt?: {[key: string]: string};
    kind?: string;
    labels?: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    mitigatedAt?: {[key: string]: string};
    private?: boolean;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    resolvedAt?: {[key: string]: string};
    services?: string;
    severity?: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    startedAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    updatedAt?: {[key: string]: string};
}

/**
 * A collection of values returned by getIncident.
 */
export interface GetIncidentResult {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly acknowledgedAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly createdAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly detectedAt?: {[key: string]: string};
    readonly environments: string;
    readonly functionalities: string;
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly inTriageAt?: {[key: string]: string};
    readonly kind: string;
    readonly labels: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly mitigatedAt?: {[key: string]: string};
    readonly private: boolean;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly resolvedAt?: {[key: string]: string};
    readonly services: string;
    readonly severity: string;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly startedAt?: {[key: string]: string};
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    readonly updatedAt?: {[key: string]: string};
}
/**
 * ## Example Usage
 */
export function getIncidentOutput(args?: GetIncidentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIncidentResult> {
    return pulumi.output(args).apply((a: any) => getIncident(a, opts))
}

/**
 * A collection of arguments for invoking getIncident.
 */
export interface GetIncidentOutputArgs {
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    acknowledgedAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    createdAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    detectedAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    environments?: pulumi.Input<string>;
    functionalities?: pulumi.Input<string>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    inTriageAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    kind?: pulumi.Input<string>;
    labels?: pulumi.Input<string>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    mitigatedAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    private?: pulumi.Input<boolean>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    resolvedAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    services?: pulumi.Input<string>;
    severity?: pulumi.Input<string>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    startedAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Filter by date range using 'lt' and 'gt'.
     */
    updatedAt?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
