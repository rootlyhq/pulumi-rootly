// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Playbook extends pulumi.CustomResource {
    /**
     * Get an existing Playbook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PlaybookState, opts?: pulumi.CustomResourceOptions): Playbook {
        return new Playbook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/playbook:Playbook';

    /**
     * Returns true if the given object is an instance of Playbook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Playbook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Playbook.__pulumiType;
    }

    /**
     * The Environment ID's to attach to the incident
     */
    public readonly environmentIds!: pulumi.Output<string[]>;
    /**
     * The external url of the playbook
     */
    public readonly externalUrl!: pulumi.Output<string>;
    /**
     * The Functionality ID's to attach to the incident
     */
    public readonly functionalityIds!: pulumi.Output<string[]>;
    /**
     * The Team ID's to attach to the incident
     */
    public readonly groupIds!: pulumi.Output<string[]>;
    /**
     * The Incident Type ID's to attach to the incident
     */
    public readonly incidentTypeIds!: pulumi.Output<string[]>;
    /**
     * The Service ID's to attach to the incident
     */
    public readonly serviceIds!: pulumi.Output<string[]>;
    /**
     * The Severity ID's to attach to the incident
     */
    public readonly severityIds!: pulumi.Output<string[]>;
    /**
     * The summary of the playbook
     */
    public readonly summary!: pulumi.Output<string>;
    /**
     * The title of the playbook
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a Playbook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PlaybookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PlaybookArgs | PlaybookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PlaybookState | undefined;
            resourceInputs["environmentIds"] = state ? state.environmentIds : undefined;
            resourceInputs["externalUrl"] = state ? state.externalUrl : undefined;
            resourceInputs["functionalityIds"] = state ? state.functionalityIds : undefined;
            resourceInputs["groupIds"] = state ? state.groupIds : undefined;
            resourceInputs["incidentTypeIds"] = state ? state.incidentTypeIds : undefined;
            resourceInputs["serviceIds"] = state ? state.serviceIds : undefined;
            resourceInputs["severityIds"] = state ? state.severityIds : undefined;
            resourceInputs["summary"] = state ? state.summary : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as PlaybookArgs | undefined;
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["environmentIds"] = args ? args.environmentIds : undefined;
            resourceInputs["externalUrl"] = args ? args.externalUrl : undefined;
            resourceInputs["functionalityIds"] = args ? args.functionalityIds : undefined;
            resourceInputs["groupIds"] = args ? args.groupIds : undefined;
            resourceInputs["incidentTypeIds"] = args ? args.incidentTypeIds : undefined;
            resourceInputs["serviceIds"] = args ? args.serviceIds : undefined;
            resourceInputs["severityIds"] = args ? args.severityIds : undefined;
            resourceInputs["summary"] = args ? args.summary : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Playbook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Playbook resources.
 */
export interface PlaybookState {
    /**
     * The Environment ID's to attach to the incident
     */
    environmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The external url of the playbook
     */
    externalUrl?: pulumi.Input<string>;
    /**
     * The Functionality ID's to attach to the incident
     */
    functionalityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Team ID's to attach to the incident
     */
    groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Incident Type ID's to attach to the incident
     */
    incidentTypeIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Service ID's to attach to the incident
     */
    serviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Severity ID's to attach to the incident
     */
    severityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The summary of the playbook
     */
    summary?: pulumi.Input<string>;
    /**
     * The title of the playbook
     */
    title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Playbook resource.
 */
export interface PlaybookArgs {
    /**
     * The Environment ID's to attach to the incident
     */
    environmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The external url of the playbook
     */
    externalUrl?: pulumi.Input<string>;
    /**
     * The Functionality ID's to attach to the incident
     */
    functionalityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Team ID's to attach to the incident
     */
    groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Incident Type ID's to attach to the incident
     */
    incidentTypeIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Service ID's to attach to the incident
     */
    serviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Severity ID's to attach to the incident
     */
    severityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The summary of the playbook
     */
    summary?: pulumi.Input<string>;
    /**
     * The title of the playbook
     */
    title: pulumi.Input<string>;
}
