// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Authorization extends pulumi.CustomResource {
    /**
     * Get an existing Authorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthorizationState, opts?: pulumi.CustomResourceOptions): Authorization {
        return new Authorization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/authorization:Authorization';

    /**
     * Returns true if the given object is an instance of Authorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Authorization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Authorization.__pulumiType;
    }

    /**
     * The id of the resource being accessed.
     */
    public readonly authorizableId!: pulumi.Output<string>;
    /**
     * The type of resource being accessed.. Value must be one of `Dashboard`.
     */
    public readonly authorizableType!: pulumi.Output<string | undefined>;
    /**
     * The resource id granted access.
     */
    public readonly granteeId!: pulumi.Output<string>;
    /**
     * The type of resource granted access.. Value must be one of `User`, `Team`.
     */
    public readonly granteeType!: pulumi.Output<string | undefined>;
    /**
     * Value must be one of `read`, `update`, `authorize`, `destroy`.
     */
    public readonly permissions!: pulumi.Output<string[]>;

    /**
     * Create a Authorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthorizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthorizationArgs | AuthorizationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthorizationState | undefined;
            resourceInputs["authorizableId"] = state ? state.authorizableId : undefined;
            resourceInputs["authorizableType"] = state ? state.authorizableType : undefined;
            resourceInputs["granteeId"] = state ? state.granteeId : undefined;
            resourceInputs["granteeType"] = state ? state.granteeType : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
        } else {
            const args = argsOrState as AuthorizationArgs | undefined;
            if ((!args || args.authorizableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizableId'");
            }
            if ((!args || args.granteeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'granteeId'");
            }
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            resourceInputs["authorizableId"] = args ? args.authorizableId : undefined;
            resourceInputs["authorizableType"] = args ? args.authorizableType : undefined;
            resourceInputs["granteeId"] = args ? args.granteeId : undefined;
            resourceInputs["granteeType"] = args ? args.granteeType : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Authorization.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Authorization resources.
 */
export interface AuthorizationState {
    /**
     * The id of the resource being accessed.
     */
    authorizableId?: pulumi.Input<string>;
    /**
     * The type of resource being accessed.. Value must be one of `Dashboard`.
     */
    authorizableType?: pulumi.Input<string>;
    /**
     * The resource id granted access.
     */
    granteeId?: pulumi.Input<string>;
    /**
     * The type of resource granted access.. Value must be one of `User`, `Team`.
     */
    granteeType?: pulumi.Input<string>;
    /**
     * Value must be one of `read`, `update`, `authorize`, `destroy`.
     */
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Authorization resource.
 */
export interface AuthorizationArgs {
    /**
     * The id of the resource being accessed.
     */
    authorizableId: pulumi.Input<string>;
    /**
     * The type of resource being accessed.. Value must be one of `Dashboard`.
     */
    authorizableType?: pulumi.Input<string>;
    /**
     * The resource id granted access.
     */
    granteeId: pulumi.Input<string>;
    /**
     * The type of resource granted access.. Value must be one of `User`, `Team`.
     */
    granteeType?: pulumi.Input<string>;
    /**
     * Value must be one of `read`, `update`, `authorize`, `destroy`.
     */
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}
