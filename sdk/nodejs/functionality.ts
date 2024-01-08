// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rootly from "@pulumi/rootly";
 *
 * const addItemsToCard = new rootly.Functionality("addItemsToCard", {
 *     color: "#800080",
 *     notifyEmails: [
 *         "foo@acme.com",
 *         "bar@acme.com",
 *     ],
 *     slackAliases: [{
 *         id: "S0614TZR7",
 *         name: "Alias 1",
 *     }],
 *     slackChannels: [
 *         {
 *             id: "C06A4RZR9",
 *             name: "Channel 1",
 *         },
 *         {
 *             id: "C02T4RYR2",
 *             name: "Channel 2",
 *         },
 *     ],
 * });
 * const loggingIn = new rootly.Functionality("loggingIn", {
 *     color: "#800080",
 *     notifyEmails: [
 *         "foo@acme.com",
 *         "bar@acme.com",
 *     ],
 *     slackAliases: [{
 *         id: "S0614TZR7",
 *         name: "Alias 1",
 *     }],
 *     slackChannels: [
 *         {
 *             id: "C06A4RZR9",
 *             name: "Channel 1",
 *         },
 *         {
 *             id: "C02T4RYR2",
 *             name: "Channel 2",
 *         },
 *     ],
 * });
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
     * The hex color of the functionality
     */
    public readonly color!: pulumi.Output<string>;
    /**
     * The description of the functionality
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Environments associated with this functionality
     */
    public readonly environmentIds!: pulumi.Output<string[]>;
    /**
     * The name of the functionality
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Emails attached to the functionality
     */
    public readonly notifyEmails!: pulumi.Output<string[]>;
    /**
     * Owner Teams associated with this functionality
     */
    public readonly ownersGroupIds!: pulumi.Output<string[]>;
    /**
     * Owner Users associated with this service
     */
    public readonly ownersUserIds!: pulumi.Output<number[]>;
    /**
     * Position of the functionality
     */
    public readonly position!: pulumi.Output<number>;
    /**
     * The public description of the functionality
     */
    public readonly publicDescription!: pulumi.Output<string>;
    /**
     * Services associated with this functionality
     */
    public readonly serviceIds!: pulumi.Output<string[]>;
    /**
     * Slack Aliases associated with this service
     */
    public readonly slackAliases!: pulumi.Output<outputs.FunctionalitySlackAlias[]>;
    /**
     * Slack Channels associated with this service
     */
    public readonly slackChannels!: pulumi.Output<outputs.FunctionalitySlackChannel[]>;
    /**
     * The slug of the functionality
     */
    public readonly slug!: pulumi.Output<string>;
    /**
     * The status of the functionality. Value must be one of `operational`, `impacted`, `outage`, `partialOutage`, `majorOutage`.
     */
    public readonly status!: pulumi.Output<string | undefined>;

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
            resourceInputs["environmentIds"] = state ? state.environmentIds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notifyEmails"] = state ? state.notifyEmails : undefined;
            resourceInputs["ownersGroupIds"] = state ? state.ownersGroupIds : undefined;
            resourceInputs["ownersUserIds"] = state ? state.ownersUserIds : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["publicDescription"] = state ? state.publicDescription : undefined;
            resourceInputs["serviceIds"] = state ? state.serviceIds : undefined;
            resourceInputs["slackAliases"] = state ? state.slackAliases : undefined;
            resourceInputs["slackChannels"] = state ? state.slackChannels : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as FunctionalityArgs | undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environmentIds"] = args ? args.environmentIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notifyEmails"] = args ? args.notifyEmails : undefined;
            resourceInputs["ownersGroupIds"] = args ? args.ownersGroupIds : undefined;
            resourceInputs["ownersUserIds"] = args ? args.ownersUserIds : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["publicDescription"] = args ? args.publicDescription : undefined;
            resourceInputs["serviceIds"] = args ? args.serviceIds : undefined;
            resourceInputs["slackAliases"] = args ? args.slackAliases : undefined;
            resourceInputs["slackChannels"] = args ? args.slackChannels : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
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
     * The hex color of the functionality
     */
    color?: pulumi.Input<string>;
    /**
     * The description of the functionality
     */
    description?: pulumi.Input<string>;
    /**
     * Environments associated with this functionality
     */
    environmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the functionality
     */
    name?: pulumi.Input<string>;
    /**
     * Emails attached to the functionality
     */
    notifyEmails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Owner Teams associated with this functionality
     */
    ownersGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Owner Users associated with this service
     */
    ownersUserIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Position of the functionality
     */
    position?: pulumi.Input<number>;
    /**
     * The public description of the functionality
     */
    publicDescription?: pulumi.Input<string>;
    /**
     * Services associated with this functionality
     */
    serviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Slack Aliases associated with this service
     */
    slackAliases?: pulumi.Input<pulumi.Input<inputs.FunctionalitySlackAlias>[]>;
    /**
     * Slack Channels associated with this service
     */
    slackChannels?: pulumi.Input<pulumi.Input<inputs.FunctionalitySlackChannel>[]>;
    /**
     * The slug of the functionality
     */
    slug?: pulumi.Input<string>;
    /**
     * The status of the functionality. Value must be one of `operational`, `impacted`, `outage`, `partialOutage`, `majorOutage`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Functionality resource.
 */
export interface FunctionalityArgs {
    /**
     * The hex color of the functionality
     */
    color?: pulumi.Input<string>;
    /**
     * The description of the functionality
     */
    description?: pulumi.Input<string>;
    /**
     * Environments associated with this functionality
     */
    environmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the functionality
     */
    name?: pulumi.Input<string>;
    /**
     * Emails attached to the functionality
     */
    notifyEmails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Owner Teams associated with this functionality
     */
    ownersGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Owner Users associated with this service
     */
    ownersUserIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Position of the functionality
     */
    position?: pulumi.Input<number>;
    /**
     * The public description of the functionality
     */
    publicDescription?: pulumi.Input<string>;
    /**
     * Services associated with this functionality
     */
    serviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Slack Aliases associated with this service
     */
    slackAliases?: pulumi.Input<pulumi.Input<inputs.FunctionalitySlackAlias>[]>;
    /**
     * Slack Channels associated with this service
     */
    slackChannels?: pulumi.Input<pulumi.Input<inputs.FunctionalitySlackChannel>[]>;
    /**
     * The slug of the functionality
     */
    slug?: pulumi.Input<string>;
    /**
     * The status of the functionality. Value must be one of `operational`, `impacted`, `outage`, `partialOutage`, `majorOutage`.
     */
    status?: pulumi.Input<string>;
}
