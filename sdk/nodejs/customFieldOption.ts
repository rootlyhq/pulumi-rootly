// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages custom field options.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rootly from "@pulumi/rootly";
 *
 * const foo = new rootly.CustomFieldOption("foo", {
 *     name: "bar",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import rootly:index/customFieldOption:CustomFieldOption foo 11111111-2222-3333-4444-555555555555
 * ```
 */
export class CustomFieldOption extends pulumi.CustomResource {
    /**
     * Get an existing CustomFieldOption resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomFieldOptionState, opts?: pulumi.CustomResourceOptions): CustomFieldOption {
        return new CustomFieldOption(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/customFieldOption:CustomFieldOption';

    /**
     * Returns true if the given object is an instance of CustomFieldOption.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomFieldOption {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomFieldOption.__pulumiType;
    }

    /**
     * The color of the custom field option
     */
    public readonly color!: pulumi.Output<string | undefined>;
    /**
     * The ID of the parent custom field
     */
    public readonly customFieldId!: pulumi.Output<number>;
    /**
     * The value of the custom field option
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a CustomFieldOption resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomFieldOptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomFieldOptionArgs | CustomFieldOptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomFieldOptionState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["customFieldId"] = state ? state.customFieldId : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as CustomFieldOptionArgs | undefined;
            if ((!args || args.customFieldId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customFieldId'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["customFieldId"] = args ? args.customFieldId : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomFieldOption.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomFieldOption resources.
 */
export interface CustomFieldOptionState {
    /**
     * The color of the custom field option
     */
    color?: pulumi.Input<string>;
    /**
     * The ID of the parent custom field
     */
    customFieldId?: pulumi.Input<number>;
    /**
     * The value of the custom field option
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomFieldOption resource.
 */
export interface CustomFieldOptionArgs {
    /**
     * The color of the custom field option
     */
    color?: pulumi.Input<string>;
    /**
     * The ID of the parent custom field
     */
    customFieldId: pulumi.Input<number>;
    /**
     * The value of the custom field option
     */
    value: pulumi.Input<string>;
}
