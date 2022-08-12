// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages workflows.
 */
export class WorkflowPulse extends pulumi.CustomResource {
    /**
     * Get an existing WorkflowPulse resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkflowPulseState, opts?: pulumi.CustomResourceOptions): WorkflowPulse {
        return new WorkflowPulse(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/workflowPulse:WorkflowPulse';

    /**
     * Returns true if the given object is an instance of WorkflowPulse.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkflowPulse {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkflowPulse.__pulumiType;
    }

    /**
     * The workflow command.
     */
    public readonly command!: pulumi.Output<string>;
    /**
     * The description of the workflow
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the workflow is enabled or not
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Environment IDs required to trigger workflow.
     */
    public readonly environmentIds!: pulumi.Output<string[] | undefined>;
    /**
     * Group IDs required to trigger workflow.
     */
    public readonly groupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Incident type IDs required to trigger workflow.
     */
    public readonly incidentTypeIds!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the workflow
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The position of the workflow (1 being top of list)
     */
    public readonly position!: pulumi.Output<number>;
    /**
     * Repeat workflow every duration.
     */
    public readonly repeatEveryDuration!: pulumi.Output<string | undefined>;
    /**
     * Repeat workflow on days.
     */
    public readonly repeatOns!: pulumi.Output<string[]>;
    /**
     * Service IDs required to trigger workflow.
     */
    public readonly serviceIds!: pulumi.Output<string[] | undefined>;
    /**
     * Severity IDs required to trigger workflow.
     */
    public readonly severityIds!: pulumi.Output<string[] | undefined>;
    /**
     * The conditions for triggering this workflow.
     */
    public readonly triggerParams!: pulumi.Output<outputs.WorkflowPulseTriggerParams>;
    /**
     * Wait before running workflow.
     */
    public readonly wait!: pulumi.Output<string | undefined>;
    /**
     * The workflow group this workflow belongs to.
     */
    public readonly workflowGroupId!: pulumi.Output<string>;

    /**
     * Create a WorkflowPulse resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WorkflowPulseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkflowPulseArgs | WorkflowPulseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkflowPulseState | undefined;
            resourceInputs["command"] = state ? state.command : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["environmentIds"] = state ? state.environmentIds : undefined;
            resourceInputs["groupIds"] = state ? state.groupIds : undefined;
            resourceInputs["incidentTypeIds"] = state ? state.incidentTypeIds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["repeatEveryDuration"] = state ? state.repeatEveryDuration : undefined;
            resourceInputs["repeatOns"] = state ? state.repeatOns : undefined;
            resourceInputs["serviceIds"] = state ? state.serviceIds : undefined;
            resourceInputs["severityIds"] = state ? state.severityIds : undefined;
            resourceInputs["triggerParams"] = state ? state.triggerParams : undefined;
            resourceInputs["wait"] = state ? state.wait : undefined;
            resourceInputs["workflowGroupId"] = state ? state.workflowGroupId : undefined;
        } else {
            const args = argsOrState as WorkflowPulseArgs | undefined;
            resourceInputs["command"] = args ? args.command : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["environmentIds"] = args ? args.environmentIds : undefined;
            resourceInputs["groupIds"] = args ? args.groupIds : undefined;
            resourceInputs["incidentTypeIds"] = args ? args.incidentTypeIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["repeatEveryDuration"] = args ? args.repeatEveryDuration : undefined;
            resourceInputs["repeatOns"] = args ? args.repeatOns : undefined;
            resourceInputs["serviceIds"] = args ? args.serviceIds : undefined;
            resourceInputs["severityIds"] = args ? args.severityIds : undefined;
            resourceInputs["triggerParams"] = args ? args.triggerParams : undefined;
            resourceInputs["wait"] = args ? args.wait : undefined;
            resourceInputs["workflowGroupId"] = args ? args.workflowGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WorkflowPulse.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkflowPulse resources.
 */
export interface WorkflowPulseState {
    /**
     * The workflow command.
     */
    command?: pulumi.Input<string>;
    /**
     * The description of the workflow
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the workflow is enabled or not
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Environment IDs required to trigger workflow.
     */
    environmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Group IDs required to trigger workflow.
     */
    groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Incident type IDs required to trigger workflow.
     */
    incidentTypeIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the workflow
     */
    name?: pulumi.Input<string>;
    /**
     * The position of the workflow (1 being top of list)
     */
    position?: pulumi.Input<number>;
    /**
     * Repeat workflow every duration.
     */
    repeatEveryDuration?: pulumi.Input<string>;
    /**
     * Repeat workflow on days.
     */
    repeatOns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Service IDs required to trigger workflow.
     */
    serviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Severity IDs required to trigger workflow.
     */
    severityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The conditions for triggering this workflow.
     */
    triggerParams?: pulumi.Input<inputs.WorkflowPulseTriggerParams>;
    /**
     * Wait before running workflow.
     */
    wait?: pulumi.Input<string>;
    /**
     * The workflow group this workflow belongs to.
     */
    workflowGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkflowPulse resource.
 */
export interface WorkflowPulseArgs {
    /**
     * The workflow command.
     */
    command?: pulumi.Input<string>;
    /**
     * The description of the workflow
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the workflow is enabled or not
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Environment IDs required to trigger workflow.
     */
    environmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Group IDs required to trigger workflow.
     */
    groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Incident type IDs required to trigger workflow.
     */
    incidentTypeIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the workflow
     */
    name?: pulumi.Input<string>;
    /**
     * The position of the workflow (1 being top of list)
     */
    position?: pulumi.Input<number>;
    /**
     * Repeat workflow every duration.
     */
    repeatEveryDuration?: pulumi.Input<string>;
    /**
     * Repeat workflow on days.
     */
    repeatOns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Service IDs required to trigger workflow.
     */
    serviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Severity IDs required to trigger workflow.
     */
    severityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The conditions for triggering this workflow.
     */
    triggerParams?: pulumi.Input<inputs.WorkflowPulseTriggerParams>;
    /**
     * Wait before running workflow.
     */
    wait?: pulumi.Input<string>;
    /**
     * The workflow group this workflow belongs to.
     */
    workflowGroupId?: pulumi.Input<string>;
}
