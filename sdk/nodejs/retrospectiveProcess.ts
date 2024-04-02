// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class RetrospectiveProcess extends pulumi.CustomResource {
    /**
     * Get an existing RetrospectiveProcess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RetrospectiveProcessState, opts?: pulumi.CustomResourceOptions): RetrospectiveProcess {
        return new RetrospectiveProcess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/retrospectiveProcess:RetrospectiveProcess';

    /**
     * Returns true if the given object is an instance of RetrospectiveProcess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RetrospectiveProcess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RetrospectiveProcess.__pulumiType;
    }

    /**
     * Retrospective process ID from which retrospective steps have to be copied. To use starter template for retrospective steps provide value: 'starter_template'
     */
    public readonly copyFrom!: pulumi.Output<string | undefined>;
    /**
     * The description of the retrospective process
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Is the retrospective process default?. Value must be one of true or false
     */
    public readonly isDefault!: pulumi.Output<boolean>;
    /**
     * The name of the retrospective process
     */
    public readonly name!: pulumi.Output<string>;
    public readonly retrospectiveProcessMatchingCriteria!: pulumi.Output<outputs.RetrospectiveProcessRetrospectiveProcessMatchingCriteria>;

    /**
     * Create a RetrospectiveProcess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RetrospectiveProcessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RetrospectiveProcessArgs | RetrospectiveProcessState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RetrospectiveProcessState | undefined;
            resourceInputs["copyFrom"] = state ? state.copyFrom : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["retrospectiveProcessMatchingCriteria"] = state ? state.retrospectiveProcessMatchingCriteria : undefined;
        } else {
            const args = argsOrState as RetrospectiveProcessArgs | undefined;
            if ((!args || args.retrospectiveProcessMatchingCriteria === undefined) && !opts.urn) {
                throw new Error("Missing required property 'retrospectiveProcessMatchingCriteria'");
            }
            resourceInputs["copyFrom"] = args ? args.copyFrom : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isDefault"] = args ? args.isDefault : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["retrospectiveProcessMatchingCriteria"] = args ? args.retrospectiveProcessMatchingCriteria : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RetrospectiveProcess.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RetrospectiveProcess resources.
 */
export interface RetrospectiveProcessState {
    /**
     * Retrospective process ID from which retrospective steps have to be copied. To use starter template for retrospective steps provide value: 'starter_template'
     */
    copyFrom?: pulumi.Input<string>;
    /**
     * The description of the retrospective process
     */
    description?: pulumi.Input<string>;
    /**
     * Is the retrospective process default?. Value must be one of true or false
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * The name of the retrospective process
     */
    name?: pulumi.Input<string>;
    retrospectiveProcessMatchingCriteria?: pulumi.Input<inputs.RetrospectiveProcessRetrospectiveProcessMatchingCriteria>;
}

/**
 * The set of arguments for constructing a RetrospectiveProcess resource.
 */
export interface RetrospectiveProcessArgs {
    /**
     * Retrospective process ID from which retrospective steps have to be copied. To use starter template for retrospective steps provide value: 'starter_template'
     */
    copyFrom?: pulumi.Input<string>;
    /**
     * The description of the retrospective process
     */
    description?: pulumi.Input<string>;
    /**
     * Is the retrospective process default?. Value must be one of true or false
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * The name of the retrospective process
     */
    name?: pulumi.Input<string>;
    retrospectiveProcessMatchingCriteria: pulumi.Input<inputs.RetrospectiveProcessRetrospectiveProcessMatchingCriteria>;
}
