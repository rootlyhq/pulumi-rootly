// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class PlaybookTask extends pulumi.CustomResource {
    /**
     * Get an existing PlaybookTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PlaybookTaskState, opts?: pulumi.CustomResourceOptions): PlaybookTask {
        return new PlaybookTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/playbookTask:PlaybookTask';

    /**
     * Returns true if the given object is an instance of PlaybookTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PlaybookTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PlaybookTask.__pulumiType;
    }

    /**
     * The description of incident task
     */
    public readonly description!: pulumi.Output<string>;
    public readonly playbookId!: pulumi.Output<string>;
    /**
     * The task of the incident task
     */
    public readonly task!: pulumi.Output<string>;

    /**
     * Create a PlaybookTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PlaybookTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PlaybookTaskArgs | PlaybookTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PlaybookTaskState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["playbookId"] = state ? state.playbookId : undefined;
            resourceInputs["task"] = state ? state.task : undefined;
        } else {
            const args = argsOrState as PlaybookTaskArgs | undefined;
            if ((!args || args.task === undefined) && !opts.urn) {
                throw new Error("Missing required property 'task'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["playbookId"] = args ? args.playbookId : undefined;
            resourceInputs["task"] = args ? args.task : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PlaybookTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PlaybookTask resources.
 */
export interface PlaybookTaskState {
    /**
     * The description of incident task
     */
    description?: pulumi.Input<string>;
    playbookId?: pulumi.Input<string>;
    /**
     * The task of the incident task
     */
    task?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PlaybookTask resource.
 */
export interface PlaybookTaskArgs {
    /**
     * The description of incident task
     */
    description?: pulumi.Input<string>;
    playbookId?: pulumi.Input<string>;
    /**
     * The task of the incident task
     */
    task: pulumi.Input<string>;
}