// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class OverrideShift extends pulumi.CustomResource {
    /**
     * Get an existing OverrideShift resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OverrideShiftState, opts?: pulumi.CustomResourceOptions): OverrideShift {
        return new OverrideShift(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/overrideShift:OverrideShift';

    /**
     * Returns true if the given object is an instance of OverrideShift.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OverrideShift {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OverrideShift.__pulumiType;
    }

    /**
     * End datetime of shift
     */
    public readonly endsAt!: pulumi.Output<string>;
    /**
     * Denotes shift is an override shift. Value must be one of true or false
     */
    public readonly isOverride!: pulumi.Output<boolean>;
    /**
     * ID of rotation
     */
    public readonly rotationId!: pulumi.Output<string>;
    /**
     * ID of schedule
     */
    public readonly scheduleId!: pulumi.Output<string>;
    /**
     * Override metadata
     */
    public readonly shiftOverride!: pulumi.Output<{[key: string]: string}>;
    /**
     * Start datetime of shift
     */
    public readonly startsAt!: pulumi.Output<string>;
    /**
     * User metadata
     */
    public readonly user!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a OverrideShift resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OverrideShiftArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OverrideShiftArgs | OverrideShiftState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OverrideShiftState | undefined;
            resourceInputs["endsAt"] = state ? state.endsAt : undefined;
            resourceInputs["isOverride"] = state ? state.isOverride : undefined;
            resourceInputs["rotationId"] = state ? state.rotationId : undefined;
            resourceInputs["scheduleId"] = state ? state.scheduleId : undefined;
            resourceInputs["shiftOverride"] = state ? state.shiftOverride : undefined;
            resourceInputs["startsAt"] = state ? state.startsAt : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as OverrideShiftArgs | undefined;
            if ((!args || args.endsAt === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endsAt'");
            }
            if ((!args || args.startsAt === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startsAt'");
            }
            resourceInputs["endsAt"] = args ? args.endsAt : undefined;
            resourceInputs["isOverride"] = args ? args.isOverride : undefined;
            resourceInputs["rotationId"] = args ? args.rotationId : undefined;
            resourceInputs["scheduleId"] = args ? args.scheduleId : undefined;
            resourceInputs["shiftOverride"] = args ? args.shiftOverride : undefined;
            resourceInputs["startsAt"] = args ? args.startsAt : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OverrideShift.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OverrideShift resources.
 */
export interface OverrideShiftState {
    /**
     * End datetime of shift
     */
    endsAt?: pulumi.Input<string>;
    /**
     * Denotes shift is an override shift. Value must be one of true or false
     */
    isOverride?: pulumi.Input<boolean>;
    /**
     * ID of rotation
     */
    rotationId?: pulumi.Input<string>;
    /**
     * ID of schedule
     */
    scheduleId?: pulumi.Input<string>;
    /**
     * Override metadata
     */
    shiftOverride?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Start datetime of shift
     */
    startsAt?: pulumi.Input<string>;
    /**
     * User metadata
     */
    user?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a OverrideShift resource.
 */
export interface OverrideShiftArgs {
    /**
     * End datetime of shift
     */
    endsAt: pulumi.Input<string>;
    /**
     * Denotes shift is an override shift. Value must be one of true or false
     */
    isOverride?: pulumi.Input<boolean>;
    /**
     * ID of rotation
     */
    rotationId?: pulumi.Input<string>;
    /**
     * ID of schedule
     */
    scheduleId?: pulumi.Input<string>;
    /**
     * Override metadata
     */
    shiftOverride?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Start datetime of shift
     */
    startsAt: pulumi.Input<string>;
    /**
     * User metadata
     */
    user?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
