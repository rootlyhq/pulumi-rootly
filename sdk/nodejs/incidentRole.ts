// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IncidentRole extends pulumi.CustomResource {
    /**
     * Get an existing IncidentRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IncidentRoleState, opts?: pulumi.CustomResourceOptions): IncidentRole {
        return new IncidentRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/incidentRole:IncidentRole';

    /**
     * Returns true if the given object is an instance of IncidentRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IncidentRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IncidentRole.__pulumiType;
    }

    /**
     * Value must be one of true or false
     */
    public readonly allowMultiUserAssignment!: pulumi.Output<boolean>;
    /**
     * The description of the incident role
     */
    public readonly description!: pulumi.Output<string>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the incident role
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Value must be one of true or false
     */
    public readonly optional!: pulumi.Output<boolean>;
    /**
     * Position of the incident role
     */
    public readonly position!: pulumi.Output<number>;
    /**
     * The slug of the incident role
     */
    public readonly slug!: pulumi.Output<string>;
    /**
     * The summary of the incident role
     */
    public readonly summary!: pulumi.Output<string>;

    /**
     * Create a IncidentRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IncidentRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IncidentRoleArgs | IncidentRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IncidentRoleState | undefined;
            resourceInputs["allowMultiUserAssignment"] = state ? state.allowMultiUserAssignment : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["optional"] = state ? state.optional : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["summary"] = state ? state.summary : undefined;
        } else {
            const args = argsOrState as IncidentRoleArgs | undefined;
            resourceInputs["allowMultiUserAssignment"] = args ? args.allowMultiUserAssignment : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["optional"] = args ? args.optional : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["summary"] = args ? args.summary : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IncidentRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IncidentRole resources.
 */
export interface IncidentRoleState {
    /**
     * Value must be one of true or false
     */
    allowMultiUserAssignment?: pulumi.Input<boolean>;
    /**
     * The description of the incident role
     */
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of the incident role
     */
    name?: pulumi.Input<string>;
    /**
     * Value must be one of true or false
     */
    optional?: pulumi.Input<boolean>;
    /**
     * Position of the incident role
     */
    position?: pulumi.Input<number>;
    /**
     * The slug of the incident role
     */
    slug?: pulumi.Input<string>;
    /**
     * The summary of the incident role
     */
    summary?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IncidentRole resource.
 */
export interface IncidentRoleArgs {
    /**
     * Value must be one of true or false
     */
    allowMultiUserAssignment?: pulumi.Input<boolean>;
    /**
     * The description of the incident role
     */
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of the incident role
     */
    name?: pulumi.Input<string>;
    /**
     * Value must be one of true or false
     */
    optional?: pulumi.Input<boolean>;
    /**
     * Position of the incident role
     */
    position?: pulumi.Input<number>;
    /**
     * The slug of the incident role
     */
    slug?: pulumi.Input<string>;
    /**
     * The summary of the incident role
     */
    summary?: pulumi.Input<string>;
}
