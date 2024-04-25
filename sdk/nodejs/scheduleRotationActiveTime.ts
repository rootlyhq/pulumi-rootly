// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ScheduleRotationActiveTime extends pulumi.CustomResource {
    /**
     * Get an existing ScheduleRotationActiveTime resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScheduleRotationActiveTimeState, opts?: pulumi.CustomResourceOptions): ScheduleRotationActiveTime {
        return new ScheduleRotationActiveTime(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/scheduleRotationActiveTime:ScheduleRotationActiveTime';

    /**
     * Returns true if the given object is an instance of ScheduleRotationActiveTime.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScheduleRotationActiveTime {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScheduleRotationActiveTime.__pulumiType;
    }

    /**
     * End time for schedule rotation active time
     */
    public readonly endTime!: pulumi.Output<string>;
    public readonly scheduleRotationId!: pulumi.Output<string>;
    /**
     * Start time for schedule rotation active time
     */
    public readonly startTime!: pulumi.Output<string>;

    /**
     * Create a ScheduleRotationActiveTime resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduleRotationActiveTimeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduleRotationActiveTimeArgs | ScheduleRotationActiveTimeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScheduleRotationActiveTimeState | undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["scheduleRotationId"] = state ? state.scheduleRotationId : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
        } else {
            const args = argsOrState as ScheduleRotationActiveTimeArgs | undefined;
            if ((!args || args.endTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endTime'");
            }
            if ((!args || args.startTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startTime'");
            }
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["scheduleRotationId"] = args ? args.scheduleRotationId : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScheduleRotationActiveTime.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScheduleRotationActiveTime resources.
 */
export interface ScheduleRotationActiveTimeState {
    /**
     * End time for schedule rotation active time
     */
    endTime?: pulumi.Input<string>;
    scheduleRotationId?: pulumi.Input<string>;
    /**
     * Start time for schedule rotation active time
     */
    startTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScheduleRotationActiveTime resource.
 */
export interface ScheduleRotationActiveTimeArgs {
    /**
     * End time for schedule rotation active time
     */
    endTime: pulumi.Input<string>;
    scheduleRotationId?: pulumi.Input<string>;
    /**
     * Start time for schedule rotation active time
     */
    startTime: pulumi.Input<string>;
}
