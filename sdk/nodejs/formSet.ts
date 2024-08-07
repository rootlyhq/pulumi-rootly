// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FormSet extends pulumi.CustomResource {
    /**
     * Get an existing FormSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FormSetState, opts?: pulumi.CustomResourceOptions): FormSet {
        return new FormSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/formSet:FormSet';

    /**
     * Returns true if the given object is an instance of FormSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FormSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FormSet.__pulumiType;
    }

    /**
     * The forms included in the form set. Add custom forms using the custom form's `slug` field. Or choose a built-in form: `webNewIncidentForm`, `webUpdateIncidentForm`, `webIncidentPostMortemForm`, `webIncidentMitigationForm`, `webIncidentResolutionForm`, `webIncidentCancellationForm`, `webScheduledIncidentForm`, `webUpdateScheduledIncidentForm`, `slackNewIncidentForm`, `slackUpdateIncidentForm`, `slackUpdateIncidentStatusForm`, `slackIncidentMitigationForm`, `slackIncidentResolutionForm`, `slackIncidentCancellationForm`, `slackScheduledIncidentForm`, `slackUpdateScheduledIncidentForm`
     */
    public readonly forms!: pulumi.Output<string[]>;
    /**
     * Whether the form set is default. Value must be one of true or false
     */
    public readonly isDefault!: pulumi.Output<boolean>;
    /**
     * The name of the form set
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The slug of the form set
     */
    public readonly slug!: pulumi.Output<string>;

    /**
     * Create a FormSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FormSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FormSetArgs | FormSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FormSetState | undefined;
            resourceInputs["forms"] = state ? state.forms : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
        } else {
            const args = argsOrState as FormSetArgs | undefined;
            if ((!args || args.forms === undefined) && !opts.urn) {
                throw new Error("Missing required property 'forms'");
            }
            resourceInputs["forms"] = args ? args.forms : undefined;
            resourceInputs["isDefault"] = args ? args.isDefault : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FormSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FormSet resources.
 */
export interface FormSetState {
    /**
     * The forms included in the form set. Add custom forms using the custom form's `slug` field. Or choose a built-in form: `webNewIncidentForm`, `webUpdateIncidentForm`, `webIncidentPostMortemForm`, `webIncidentMitigationForm`, `webIncidentResolutionForm`, `webIncidentCancellationForm`, `webScheduledIncidentForm`, `webUpdateScheduledIncidentForm`, `slackNewIncidentForm`, `slackUpdateIncidentForm`, `slackUpdateIncidentStatusForm`, `slackIncidentMitigationForm`, `slackIncidentResolutionForm`, `slackIncidentCancellationForm`, `slackScheduledIncidentForm`, `slackUpdateScheduledIncidentForm`
     */
    forms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the form set is default. Value must be one of true or false
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * The name of the form set
     */
    name?: pulumi.Input<string>;
    /**
     * The slug of the form set
     */
    slug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FormSet resource.
 */
export interface FormSetArgs {
    /**
     * The forms included in the form set. Add custom forms using the custom form's `slug` field. Or choose a built-in form: `webNewIncidentForm`, `webUpdateIncidentForm`, `webIncidentPostMortemForm`, `webIncidentMitigationForm`, `webIncidentResolutionForm`, `webIncidentCancellationForm`, `webScheduledIncidentForm`, `webUpdateScheduledIncidentForm`, `slackNewIncidentForm`, `slackUpdateIncidentForm`, `slackUpdateIncidentStatusForm`, `slackIncidentMitigationForm`, `slackIncidentResolutionForm`, `slackIncidentCancellationForm`, `slackScheduledIncidentForm`, `slackUpdateScheduledIncidentForm`
     */
    forms: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the form set is default. Value must be one of true or false
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * The name of the form set
     */
    name?: pulumi.Input<string>;
    /**
     * The slug of the form set
     */
    slug?: pulumi.Input<string>;
}
