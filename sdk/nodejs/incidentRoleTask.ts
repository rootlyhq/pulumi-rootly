// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IncidentRoleTask extends pulumi.CustomResource {
    /**
     * Get an existing IncidentRoleTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IncidentRoleTaskState, opts?: pulumi.CustomResourceOptions): IncidentRoleTask {
        return new IncidentRoleTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/incidentRoleTask:IncidentRoleTask';

    /**
     * Returns true if the given object is an instance of IncidentRoleTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IncidentRoleTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IncidentRoleTask.__pulumiType;
    }

    /**
     * The description of incident task
     */
    public readonly description!: pulumi.Output<string>;
    public readonly incidentRoleId!: pulumi.Output<string>;
    /**
     * The priority of the incident task. Value must be one of `high`, `medium`, `low`.
     */
    public readonly priority!: pulumi.Output<string | undefined>;
    /**
     * The task of the incident task
     */
    public readonly task!: pulumi.Output<string>;

    /**
     * Create a IncidentRoleTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IncidentRoleTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IncidentRoleTaskArgs | IncidentRoleTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IncidentRoleTaskState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["incidentRoleId"] = state ? state.incidentRoleId : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["task"] = state ? state.task : undefined;
        } else {
            const args = argsOrState as IncidentRoleTaskArgs | undefined;
            if ((!args || args.task === undefined) && !opts.urn) {
                throw new Error("Missing required property 'task'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["incidentRoleId"] = args ? args.incidentRoleId : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["task"] = args ? args.task : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IncidentRoleTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IncidentRoleTask resources.
 */
export interface IncidentRoleTaskState {
    /**
     * The description of incident task
     */
    description?: pulumi.Input<string>;
    incidentRoleId?: pulumi.Input<string>;
    /**
     * The priority of the incident task. Value must be one of `high`, `medium`, `low`.
     */
    priority?: pulumi.Input<string>;
    /**
     * The task of the incident task
     */
    task?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IncidentRoleTask resource.
 */
export interface IncidentRoleTaskArgs {
    /**
     * The description of incident task
     */
    description?: pulumi.Input<string>;
    incidentRoleId?: pulumi.Input<string>;
    /**
     * The priority of the incident task. Value must be one of `high`, `medium`, `low`.
     */
    priority?: pulumi.Input<string>;
    /**
     * The task of the incident task
     */
    task: pulumi.Input<string>;
}