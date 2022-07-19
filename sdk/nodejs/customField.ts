// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages custom fields.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rootly from "@pulumi/rootly";
 *
 * const foo = new rootly.CustomField("foo", {
 *     name: "bar",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import rootly:index/customField:CustomField foo 11111111-2222-3333-4444-555555555555
 * ```
 */
export class CustomField extends pulumi.CustomResource {
    /**
     * Get an existing CustomField resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomFieldState, opts?: pulumi.CustomResourceOptions): CustomField {
        return new CustomField(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/customField:CustomField';

    /**
     * Returns true if the given object is an instance of CustomField.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomField {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomField.__pulumiType;
    }

    /**
     * The description of the custom field
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the custom field is enabled or not
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The kind of the custom field
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The name of the custom field
     */
    public readonly label!: pulumi.Output<string>;
    /**
     * Where the custom field is required.
     */
    public readonly requireds!: pulumi.Output<string[] | undefined>;
    /**
     * Where the custom field is shown.
     */
    public readonly showns!: pulumi.Output<string[] | undefined>;

    /**
     * Create a CustomField resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomFieldArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomFieldArgs | CustomFieldState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomFieldState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["label"] = state ? state.label : undefined;
            resourceInputs["requireds"] = state ? state.requireds : undefined;
            resourceInputs["showns"] = state ? state.showns : undefined;
        } else {
            const args = argsOrState as CustomFieldArgs | undefined;
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.label === undefined) && !opts.urn) {
                throw new Error("Missing required property 'label'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["label"] = args ? args.label : undefined;
            resourceInputs["requireds"] = args ? args.requireds : undefined;
            resourceInputs["showns"] = args ? args.showns : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomField.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomField resources.
 */
export interface CustomFieldState {
    /**
     * The description of the custom field
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the custom field is enabled or not
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The kind of the custom field
     */
    kind?: pulumi.Input<string>;
    /**
     * The name of the custom field
     */
    label?: pulumi.Input<string>;
    /**
     * Where the custom field is required.
     */
    requireds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where the custom field is shown.
     */
    showns?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a CustomField resource.
 */
export interface CustomFieldArgs {
    /**
     * The description of the custom field
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the custom field is enabled or not
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The kind of the custom field
     */
    kind: pulumi.Input<string>;
    /**
     * The name of the custom field
     */
    label: pulumi.Input<string>;
    /**
     * Where the custom field is required.
     */
    requireds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where the custom field is shown.
     */
    showns?: pulumi.Input<pulumi.Input<string>[]>;
}