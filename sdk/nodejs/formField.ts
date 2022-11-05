// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FormField extends pulumi.CustomResource {
    /**
     * Get an existing FormField resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FormFieldState, opts?: pulumi.CustomResourceOptions): FormField {
        return new FormField(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/formField:FormField';

    /**
     * Returns true if the given object is an instance of FormField.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FormField {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FormField.__pulumiType;
    }

    public readonly defaultValues!: pulumi.Output<string[]>;
    /**
     * The description of the form field
     */
    public readonly description!: pulumi.Output<string>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The input kind of the form field. Value must be one of `text`, `textarea`, `select`, `multi_select`, `date`, `datetime`,
     * `users`, `number`, `checkbox`, `tags`.
     */
    public readonly inputKind!: pulumi.Output<string | undefined>;
    /**
     * The kind of the form field. Value must be one of `custom`, `title`, `summary`, `severity`, `environments`, `types`,
     * `services`, `functionalities`, `teams`, `visibility`, `mark_as_test`, `mark_as_backfilled`, `labels`, `notify_emails`,
     * `trigger_manual_workflows`, `show_ongoing_incidents`, `attach_alerts`, `manual_starting_datetime_field`.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The name of the form field
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * . Value must be one of `web_new_incident_form`, `web_update_incident_form`, `web_incident_post_mortem_form`,
     * `web_incident_mitigation_form`, `web_incident_resolution_form`, `slack_new_incident_form`, `slack_update_incident_form`,
     * `slack_update_incident_status_form`, `slack_incident_mitigation_form`, `slack_incident_resolution_form`.
     */
    public readonly requireds!: pulumi.Output<string[]>;
    /**
     * . Value must be one of `web_new_incident_form`, `web_update_incident_form`, `web_incident_post_mortem_form`,
     * `web_incident_mitigation_form`, `web_incident_resolution_form`, `incident_post_mortem`, `slack_new_incident_form`,
     * `slack_update_incident_form`, `slack_update_incident_status_form`, `slack_incident_mitigation_form`,
     * `slack_incident_resolution_form`.
     */
    public readonly showns!: pulumi.Output<string[]>;
    /**
     * The slug of the form field
     */
    public readonly slug!: pulumi.Output<string>;

    /**
     * Create a FormField resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FormFieldArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FormFieldArgs | FormFieldState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FormFieldState | undefined;
            resourceInputs["defaultValues"] = state ? state.defaultValues : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["inputKind"] = state ? state.inputKind : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["requireds"] = state ? state.requireds : undefined;
            resourceInputs["showns"] = state ? state.showns : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
        } else {
            const args = argsOrState as FormFieldArgs | undefined;
            resourceInputs["defaultValues"] = args ? args.defaultValues : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["inputKind"] = args ? args.inputKind : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["requireds"] = args ? args.requireds : undefined;
            resourceInputs["showns"] = args ? args.showns : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FormField.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FormField resources.
 */
export interface FormFieldState {
    defaultValues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the form field
     */
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    /**
     * The input kind of the form field. Value must be one of `text`, `textarea`, `select`, `multi_select`, `date`, `datetime`,
     * `users`, `number`, `checkbox`, `tags`.
     */
    inputKind?: pulumi.Input<string>;
    /**
     * The kind of the form field. Value must be one of `custom`, `title`, `summary`, `severity`, `environments`, `types`,
     * `services`, `functionalities`, `teams`, `visibility`, `mark_as_test`, `mark_as_backfilled`, `labels`, `notify_emails`,
     * `trigger_manual_workflows`, `show_ongoing_incidents`, `attach_alerts`, `manual_starting_datetime_field`.
     */
    kind?: pulumi.Input<string>;
    /**
     * The name of the form field
     */
    name?: pulumi.Input<string>;
    /**
     * . Value must be one of `web_new_incident_form`, `web_update_incident_form`, `web_incident_post_mortem_form`,
     * `web_incident_mitigation_form`, `web_incident_resolution_form`, `slack_new_incident_form`, `slack_update_incident_form`,
     * `slack_update_incident_status_form`, `slack_incident_mitigation_form`, `slack_incident_resolution_form`.
     */
    requireds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * . Value must be one of `web_new_incident_form`, `web_update_incident_form`, `web_incident_post_mortem_form`,
     * `web_incident_mitigation_form`, `web_incident_resolution_form`, `incident_post_mortem`, `slack_new_incident_form`,
     * `slack_update_incident_form`, `slack_update_incident_status_form`, `slack_incident_mitigation_form`,
     * `slack_incident_resolution_form`.
     */
    showns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The slug of the form field
     */
    slug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FormField resource.
 */
export interface FormFieldArgs {
    defaultValues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the form field
     */
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    /**
     * The input kind of the form field. Value must be one of `text`, `textarea`, `select`, `multi_select`, `date`, `datetime`,
     * `users`, `number`, `checkbox`, `tags`.
     */
    inputKind?: pulumi.Input<string>;
    /**
     * The kind of the form field. Value must be one of `custom`, `title`, `summary`, `severity`, `environments`, `types`,
     * `services`, `functionalities`, `teams`, `visibility`, `mark_as_test`, `mark_as_backfilled`, `labels`, `notify_emails`,
     * `trigger_manual_workflows`, `show_ongoing_incidents`, `attach_alerts`, `manual_starting_datetime_field`.
     */
    kind?: pulumi.Input<string>;
    /**
     * The name of the form field
     */
    name?: pulumi.Input<string>;
    /**
     * . Value must be one of `web_new_incident_form`, `web_update_incident_form`, `web_incident_post_mortem_form`,
     * `web_incident_mitigation_form`, `web_incident_resolution_form`, `slack_new_incident_form`, `slack_update_incident_form`,
     * `slack_update_incident_status_form`, `slack_incident_mitigation_form`, `slack_incident_resolution_form`.
     */
    requireds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * . Value must be one of `web_new_incident_form`, `web_update_incident_form`, `web_incident_post_mortem_form`,
     * `web_incident_mitigation_form`, `web_incident_resolution_form`, `incident_post_mortem`, `slack_new_incident_form`,
     * `slack_update_incident_form`, `slack_update_incident_status_form`, `slack_incident_mitigation_form`,
     * `slack_incident_resolution_form`.
     */
    showns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The slug of the form field
     */
    slug?: pulumi.Input<string>;
}