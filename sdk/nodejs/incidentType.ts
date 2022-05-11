// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages incident types (e.g Cloud, Customer Facing, Security, Training).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rootly from "@pulumi/rootly";
 *
 * const foo = new rootly.IncidentType("foo", {});
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import rootly:index/incidentType:IncidentType foo 11111111-2222-3333-4444-555555555555
 * ```
 */
export class IncidentType extends pulumi.CustomResource {
    /**
     * Get an existing IncidentType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IncidentTypeState, opts?: pulumi.CustomResourceOptions): IncidentType {
        return new IncidentType(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/incidentType:IncidentType';

    /**
     * Returns true if the given object is an instance of IncidentType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IncidentType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IncidentType.__pulumiType;
    }

    /**
     * The cikir of the incident type
     */
    public readonly color!: pulumi.Output<string | undefined>;
    /**
     * The description of the incident type
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the incident type
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a IncidentType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IncidentTypeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IncidentTypeArgs | IncidentTypeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IncidentTypeState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as IncidentTypeArgs | undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IncidentType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IncidentType resources.
 */
export interface IncidentTypeState {
    /**
     * The cikir of the incident type
     */
    color?: pulumi.Input<string>;
    /**
     * The description of the incident type
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the incident type
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IncidentType resource.
 */
export interface IncidentTypeArgs {
    /**
     * The cikir of the incident type
     */
    color?: pulumi.Input<string>;
    /**
     * The description of the incident type
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the incident type
     */
    name?: pulumi.Input<string>;
}
