// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages workflow createGitlabIssue task.
 */
export class WorkflowTaskCreateGitlabIssue extends pulumi.CustomResource {
    /**
     * Get an existing WorkflowTaskCreateGitlabIssue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkflowTaskCreateGitlabIssueState, opts?: pulumi.CustomResourceOptions): WorkflowTaskCreateGitlabIssue {
        return new WorkflowTaskCreateGitlabIssue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/workflowTaskCreateGitlabIssue:WorkflowTaskCreateGitlabIssue';

    /**
     * Returns true if the given object is an instance of WorkflowTaskCreateGitlabIssue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkflowTaskCreateGitlabIssue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkflowTaskCreateGitlabIssue.__pulumiType;
    }

    /**
     * Enable/disable this workflow task
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the workflow task
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The position of the workflow task (1 being top of list)
     */
    public readonly position!: pulumi.Output<number>;
    /**
     * Skip workflow task if any failures
     */
    public readonly skipOnFailure!: pulumi.Output<boolean | undefined>;
    /**
     * The parameters for this workflow task.
     */
    public readonly taskParams!: pulumi.Output<outputs.WorkflowTaskCreateGitlabIssueTaskParams>;
    /**
     * The ID of the parent workflow
     */
    public readonly workflowId!: pulumi.Output<string>;

    /**
     * Create a WorkflowTaskCreateGitlabIssue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkflowTaskCreateGitlabIssueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkflowTaskCreateGitlabIssueArgs | WorkflowTaskCreateGitlabIssueState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkflowTaskCreateGitlabIssueState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["skipOnFailure"] = state ? state.skipOnFailure : undefined;
            resourceInputs["taskParams"] = state ? state.taskParams : undefined;
            resourceInputs["workflowId"] = state ? state.workflowId : undefined;
        } else {
            const args = argsOrState as WorkflowTaskCreateGitlabIssueArgs | undefined;
            if ((!args || args.taskParams === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskParams'");
            }
            if ((!args || args.workflowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workflowId'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["skipOnFailure"] = args ? args.skipOnFailure : undefined;
            resourceInputs["taskParams"] = args ? args.taskParams : undefined;
            resourceInputs["workflowId"] = args ? args.workflowId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WorkflowTaskCreateGitlabIssue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkflowTaskCreateGitlabIssue resources.
 */
export interface WorkflowTaskCreateGitlabIssueState {
    /**
     * Enable/disable this workflow task
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Name of the workflow task
     */
    name?: pulumi.Input<string>;
    /**
     * The position of the workflow task (1 being top of list)
     */
    position?: pulumi.Input<number>;
    /**
     * Skip workflow task if any failures
     */
    skipOnFailure?: pulumi.Input<boolean>;
    /**
     * The parameters for this workflow task.
     */
    taskParams?: pulumi.Input<inputs.WorkflowTaskCreateGitlabIssueTaskParams>;
    /**
     * The ID of the parent workflow
     */
    workflowId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkflowTaskCreateGitlabIssue resource.
 */
export interface WorkflowTaskCreateGitlabIssueArgs {
    /**
     * Enable/disable this workflow task
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Name of the workflow task
     */
    name?: pulumi.Input<string>;
    /**
     * The position of the workflow task (1 being top of list)
     */
    position?: pulumi.Input<number>;
    /**
     * Skip workflow task if any failures
     */
    skipOnFailure?: pulumi.Input<boolean>;
    /**
     * The parameters for this workflow task.
     */
    taskParams: pulumi.Input<inputs.WorkflowTaskCreateGitlabIssueTaskParams>;
    /**
     * The ID of the parent workflow
     */
    workflowId: pulumi.Input<string>;
}
