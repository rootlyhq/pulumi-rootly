// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages Services (e.g elasticsearch-prod, redis-preprod, customer-postgresql-prod).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rootly from "@pulumi/rootly";
 *
 * const foo = new rootly.Service("foo", {});
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import rootly:index/service:Service foo 11111111-2222-3333-4444-555555555555
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * The color chosen for the service
     */
    public readonly color!: pulumi.Output<string | undefined>;
    /**
     * For internal use only
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the service
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * This will be displayed on your status pages to explain to your customer the use of this service.
     */
    public readonly publicDescription!: pulumi.Output<string | undefined>;
    /**
     * The slug of the service
     */
    public readonly slug!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["publicDescription"] = state ? state.publicDescription : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publicDescription"] = args ? args.publicDescription : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * The color chosen for the service
     */
    color?: pulumi.Input<string>;
    /**
     * For internal use only
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the service
     */
    name?: pulumi.Input<string>;
    /**
     * This will be displayed on your status pages to explain to your customer the use of this service.
     */
    publicDescription?: pulumi.Input<string>;
    /**
     * The slug of the service
     */
    slug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * The color chosen for the service
     */
    color?: pulumi.Input<string>;
    /**
     * For internal use only
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the service
     */
    name?: pulumi.Input<string>;
    /**
     * This will be displayed on your status pages to explain to your customer the use of this service.
     */
    publicDescription?: pulumi.Input<string>;
    /**
     * The slug of the service
     */
    slug?: pulumi.Input<string>;
}
