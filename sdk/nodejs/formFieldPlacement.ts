// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FormFieldPlacement extends pulumi.CustomResource {
    /**
     * Get an existing FormFieldPlacement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FormFieldPlacementState, opts?: pulumi.CustomResourceOptions): FormFieldPlacement {
        return new FormFieldPlacement(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/formFieldPlacement:FormFieldPlacement';

    /**
     * Returns true if the given object is an instance of FormFieldPlacement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FormFieldPlacement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FormFieldPlacement.__pulumiType;
    }

    /**
     * The form this field is placed on.
     */
    public readonly form!: pulumi.Output<string>;
    /**
     * The form field that is placed.
     */
    public readonly formFieldId!: pulumi.Output<string>;
    /**
     * The form set this field is placed in.
     */
    public readonly formSetId!: pulumi.Output<string>;
    /**
     * The position of the field placement.
     */
    public readonly position!: pulumi.Output<number>;
    /**
     * Whether the field is required on this form.. Value must be one of true or false
     */
    public readonly required!: pulumi.Output<boolean>;

    /**
     * Create a FormFieldPlacement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FormFieldPlacementArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FormFieldPlacementArgs | FormFieldPlacementState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FormFieldPlacementState | undefined;
            resourceInputs["form"] = state ? state.form : undefined;
            resourceInputs["formFieldId"] = state ? state.formFieldId : undefined;
            resourceInputs["formSetId"] = state ? state.formSetId : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["required"] = state ? state.required : undefined;
        } else {
            const args = argsOrState as FormFieldPlacementArgs | undefined;
            if ((!args || args.form === undefined) && !opts.urn) {
                throw new Error("Missing required property 'form'");
            }
            if ((!args || args.formSetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'formSetId'");
            }
            resourceInputs["form"] = args ? args.form : undefined;
            resourceInputs["formFieldId"] = args ? args.formFieldId : undefined;
            resourceInputs["formSetId"] = args ? args.formSetId : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["required"] = args ? args.required : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FormFieldPlacement.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FormFieldPlacement resources.
 */
export interface FormFieldPlacementState {
    /**
     * The form this field is placed on.
     */
    form?: pulumi.Input<string>;
    /**
     * The form field that is placed.
     */
    formFieldId?: pulumi.Input<string>;
    /**
     * The form set this field is placed in.
     */
    formSetId?: pulumi.Input<string>;
    /**
     * The position of the field placement.
     */
    position?: pulumi.Input<number>;
    /**
     * Whether the field is required on this form.. Value must be one of true or false
     */
    required?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a FormFieldPlacement resource.
 */
export interface FormFieldPlacementArgs {
    /**
     * The form this field is placed on.
     */
    form: pulumi.Input<string>;
    /**
     * The form field that is placed.
     */
    formFieldId?: pulumi.Input<string>;
    /**
     * The form set this field is placed in.
     */
    formSetId: pulumi.Input<string>;
    /**
     * The position of the field placement.
     */
    position?: pulumi.Input<number>;
    /**
     * Whether the field is required on this form.. Value must be one of true or false
     */
    required?: pulumi.Input<boolean>;
}