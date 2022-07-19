// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages workflow updateServiceNowIncident task.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rootly from "@pulumi/rootly";
 *
 * const foo = new rootly.WorkflowTaskUpdateServiceNowIncident("foo", {
 *     name: "bar",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import rootly:index/workflowTaskUpdateServiceNowIncident:WorkflowTaskUpdateServiceNowIncident foo 11111111-2222-3333-4444-555555555555
 * ```
 */
export class WorkflowTaskUpdateServiceNowIncident extends pulumi.CustomResource {
    /**
     * Get an existing WorkflowTaskUpdateServiceNowIncident resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkflowTaskUpdateServiceNowIncidentState, opts?: pulumi.CustomResourceOptions): WorkflowTaskUpdateServiceNowIncident {
        return new WorkflowTaskUpdateServiceNowIncident(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/workflowTaskUpdateServiceNowIncident:WorkflowTaskUpdateServiceNowIncident';

    /**
     * Returns true if the given object is an instance of WorkflowTaskUpdateServiceNowIncident.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkflowTaskUpdateServiceNowIncident {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkflowTaskUpdateServiceNowIncident.__pulumiType;
    }

    /**
     * The parameters for this workflow task.
     */
    public readonly taskParams!: pulumi.Output<outputs.WorkflowTaskUpdateServiceNowIncidentTaskParams>;
    /**
     * The ID of the parent workflow
     */
    public readonly workflowId!: pulumi.Output<string>;

    /**
     * Create a WorkflowTaskUpdateServiceNowIncident resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkflowTaskUpdateServiceNowIncidentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkflowTaskUpdateServiceNowIncidentArgs | WorkflowTaskUpdateServiceNowIncidentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkflowTaskUpdateServiceNowIncidentState | undefined;
            resourceInputs["taskParams"] = state ? state.taskParams : undefined;
            resourceInputs["workflowId"] = state ? state.workflowId : undefined;
        } else {
            const args = argsOrState as WorkflowTaskUpdateServiceNowIncidentArgs | undefined;
            if ((!args || args.taskParams === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskParams'");
            }
            if ((!args || args.workflowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workflowId'");
            }
            resourceInputs["taskParams"] = args ? args.taskParams : undefined;
            resourceInputs["workflowId"] = args ? args.workflowId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WorkflowTaskUpdateServiceNowIncident.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkflowTaskUpdateServiceNowIncident resources.
 */
export interface WorkflowTaskUpdateServiceNowIncidentState {
    /**
     * The parameters for this workflow task.
     */
    taskParams?: pulumi.Input<inputs.WorkflowTaskUpdateServiceNowIncidentTaskParams>;
    /**
     * The ID of the parent workflow
     */
    workflowId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkflowTaskUpdateServiceNowIncident resource.
 */
export interface WorkflowTaskUpdateServiceNowIncidentArgs {
    /**
     * The parameters for this workflow task.
     */
    taskParams: pulumi.Input<inputs.WorkflowTaskUpdateServiceNowIncidentTaskParams>;
    /**
     * The ID of the parent workflow
     */
    workflowId: pulumi.Input<string>;
}