// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class EscalationLevel extends pulumi.CustomResource {
    /**
     * Get an existing EscalationLevel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EscalationLevelState, opts?: pulumi.CustomResourceOptions): EscalationLevel {
        return new EscalationLevel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/escalationLevel:EscalationLevel';

    /**
     * Returns true if the given object is an instance of EscalationLevel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EscalationLevel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EscalationLevel.__pulumiType;
    }

    /**
     * Delay before notification targets will be alerted.
     */
    public readonly delay!: pulumi.Output<number>;
    /**
     * The ID of the escalation policy
     */
    public readonly escalationPolicyId!: pulumi.Output<string>;
    /**
     * Escalation level's notification targets
     */
    public readonly notificationTargetParams!: pulumi.Output<outputs.EscalationLevelNotificationTargetParam[]>;
    /**
     * Position of the escalation policy level
     */
    public readonly position!: pulumi.Output<number>;

    /**
     * Create a EscalationLevel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EscalationLevelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EscalationLevelArgs | EscalationLevelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EscalationLevelState | undefined;
            resourceInputs["delay"] = state ? state.delay : undefined;
            resourceInputs["escalationPolicyId"] = state ? state.escalationPolicyId : undefined;
            resourceInputs["notificationTargetParams"] = state ? state.notificationTargetParams : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
        } else {
            const args = argsOrState as EscalationLevelArgs | undefined;
            if ((!args || args.notificationTargetParams === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notificationTargetParams'");
            }
            if ((!args || args.position === undefined) && !opts.urn) {
                throw new Error("Missing required property 'position'");
            }
            resourceInputs["delay"] = args ? args.delay : undefined;
            resourceInputs["escalationPolicyId"] = args ? args.escalationPolicyId : undefined;
            resourceInputs["notificationTargetParams"] = args ? args.notificationTargetParams : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EscalationLevel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EscalationLevel resources.
 */
export interface EscalationLevelState {
    /**
     * Delay before notification targets will be alerted.
     */
    delay?: pulumi.Input<number>;
    /**
     * The ID of the escalation policy
     */
    escalationPolicyId?: pulumi.Input<string>;
    /**
     * Escalation level's notification targets
     */
    notificationTargetParams?: pulumi.Input<pulumi.Input<inputs.EscalationLevelNotificationTargetParam>[]>;
    /**
     * Position of the escalation policy level
     */
    position?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a EscalationLevel resource.
 */
export interface EscalationLevelArgs {
    /**
     * Delay before notification targets will be alerted.
     */
    delay?: pulumi.Input<number>;
    /**
     * The ID of the escalation policy
     */
    escalationPolicyId?: pulumi.Input<string>;
    /**
     * Escalation level's notification targets
     */
    notificationTargetParams: pulumi.Input<pulumi.Input<inputs.EscalationLevelNotificationTargetParam>[]>;
    /**
     * Position of the escalation policy level
     */
    position: pulumi.Input<number>;
}
