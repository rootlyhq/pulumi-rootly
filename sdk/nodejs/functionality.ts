// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages functionalities (e.g Logging In, Search, Adds items to Cart).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rootly from "@pulumi/rootly";
 *
 * const foo = new rootly.Functionality("foo", {});
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import rootly:index/functionality:Functionality foo 11111111-2222-3333-4444-555555555555
 * ```
 */
export class Functionality extends pulumi.CustomResource {
    /**
     * Get an existing Functionality resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionalityState, opts?: pulumi.CustomResourceOptions): Functionality {
        return new Functionality(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/functionality:Functionality';

    /**
     * Returns true if the given object is an instance of Functionality.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Functionality {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Functionality.__pulumiType;
    }

    /**
     * The color of the severity
     */
    public readonly color!: pulumi.Output<string | undefined>;
    /**
     * The description of the functionality
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the functionality
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The slug of the severity
     */
    public readonly slug!: pulumi.Output<string>;

    /**
     * Create a Functionality resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FunctionalityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionalityArgs | FunctionalityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionalityState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
        } else {
            const args = argsOrState as FunctionalityArgs | undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Functionality.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Functionality resources.
 */
export interface FunctionalityState {
    /**
     * The color of the severity
     */
    color?: pulumi.Input<string>;
    /**
     * The description of the functionality
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the functionality
     */
    name?: pulumi.Input<string>;
    /**
     * The slug of the severity
     */
    slug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Functionality resource.
 */
export interface FunctionalityArgs {
    /**
     * The color of the severity
     */
    color?: pulumi.Input<string>;
    /**
     * The description of the functionality
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the functionality
     */
    name?: pulumi.Input<string>;
    /**
     * The slug of the severity
     */
    slug?: pulumi.Input<string>;
}
