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
 * const elasticsearchProd = new rootly.Service("elasticsearchProd", {
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
 * const customerPostgresqlProd = new rootly.Service("customerPostgresqlProd", {
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
     * The Backstage entity id associated to this service. eg: :namespace/:kind/:entity_name
     */
    public readonly backstageId!: pulumi.Output<string>;
    /**
     * The hex color of the service
     */
    public readonly color!: pulumi.Output<string>;
    /**
     * The description of the service
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Environments associated with this service
     */
    public readonly environmentIds!: pulumi.Output<string[]>;
    /**
     * The GitHub repository branch associated to this service. eg: main
     */
    public readonly githubRepositoryBranch!: pulumi.Output<string>;
    /**
     * The GitHub repository name associated to this service. eg: rootlyhq/my-service
     */
    public readonly githubRepositoryName!: pulumi.Output<string>;
    /**
     * The Gitlab repository branch associated to this service. eg: main
     */
    public readonly gitlabRepositoryBranch!: pulumi.Output<string>;
    /**
     * The Gitlab repository name associated to this service. eg: rootlyhq/my-service
     */
    public readonly gitlabRepositoryName!: pulumi.Output<string>;
    /**
     * The name of the service
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Emails attached to the service
     */
    public readonly notifyEmails!: pulumi.Output<string[]>;
    /**
     * The Opsgenie service id associated to this service
     */
    public readonly opsgenieId!: pulumi.Output<string>;
    /**
     * Owner Teams associated with this service
     */
    public readonly ownersGroupIds!: pulumi.Output<string[]>;
    /**
     * Owner Users associated with this service
     */
    public readonly ownersUserIds!: pulumi.Output<number[]>;
    /**
     * The PagerDuty service id associated to this service
     */
    public readonly pagerdutyId!: pulumi.Output<string>;
    /**
     * Position of the service
     */
    public readonly position!: pulumi.Output<number>;
    /**
     * The public description of the service
     */
    public readonly publicDescription!: pulumi.Output<string>;
    /**
     * Services dependent on this service
     */
    public readonly serviceIds!: pulumi.Output<string[]>;
    /**
     * Slack Aliases associated with this service
     */
    public readonly slackAliases!: pulumi.Output<outputs.ServiceSlackAlias[]>;
    /**
     * Slack Channels associated with this service
     */
    public readonly slackChannels!: pulumi.Output<outputs.ServiceSlackChannel[]>;
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
            resourceInputs["backstageId"] = state ? state.backstageId : undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environmentIds"] = state ? state.environmentIds : undefined;
            resourceInputs["githubRepositoryBranch"] = state ? state.githubRepositoryBranch : undefined;
            resourceInputs["githubRepositoryName"] = state ? state.githubRepositoryName : undefined;
            resourceInputs["gitlabRepositoryBranch"] = state ? state.gitlabRepositoryBranch : undefined;
            resourceInputs["gitlabRepositoryName"] = state ? state.gitlabRepositoryName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notifyEmails"] = state ? state.notifyEmails : undefined;
            resourceInputs["opsgenieId"] = state ? state.opsgenieId : undefined;
            resourceInputs["ownersGroupIds"] = state ? state.ownersGroupIds : undefined;
            resourceInputs["ownersUserIds"] = state ? state.ownersUserIds : undefined;
            resourceInputs["pagerdutyId"] = state ? state.pagerdutyId : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["publicDescription"] = state ? state.publicDescription : undefined;
            resourceInputs["serviceIds"] = state ? state.serviceIds : undefined;
            resourceInputs["slackAliases"] = state ? state.slackAliases : undefined;
            resourceInputs["slackChannels"] = state ? state.slackChannels : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            resourceInputs["backstageId"] = args ? args.backstageId : undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environmentIds"] = args ? args.environmentIds : undefined;
            resourceInputs["githubRepositoryBranch"] = args ? args.githubRepositoryBranch : undefined;
            resourceInputs["githubRepositoryName"] = args ? args.githubRepositoryName : undefined;
            resourceInputs["gitlabRepositoryBranch"] = args ? args.gitlabRepositoryBranch : undefined;
            resourceInputs["gitlabRepositoryName"] = args ? args.gitlabRepositoryName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notifyEmails"] = args ? args.notifyEmails : undefined;
            resourceInputs["opsgenieId"] = args ? args.opsgenieId : undefined;
            resourceInputs["ownersGroupIds"] = args ? args.ownersGroupIds : undefined;
            resourceInputs["ownersUserIds"] = args ? args.ownersUserIds : undefined;
            resourceInputs["pagerdutyId"] = args ? args.pagerdutyId : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["publicDescription"] = args ? args.publicDescription : undefined;
            resourceInputs["serviceIds"] = args ? args.serviceIds : undefined;
            resourceInputs["slackAliases"] = args ? args.slackAliases : undefined;
            resourceInputs["slackChannels"] = args ? args.slackChannels : undefined;
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
     * The Backstage entity id associated to this service. eg: :namespace/:kind/:entity_name
     */
    backstageId?: pulumi.Input<string>;
    /**
     * The hex color of the service
     */
    color?: pulumi.Input<string>;
    /**
     * The description of the service
     */
    description?: pulumi.Input<string>;
    /**
     * Environments associated with this service
     */
    environmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The GitHub repository branch associated to this service. eg: main
     */
    githubRepositoryBranch?: pulumi.Input<string>;
    /**
     * The GitHub repository name associated to this service. eg: rootlyhq/my-service
     */
    githubRepositoryName?: pulumi.Input<string>;
    /**
     * The Gitlab repository branch associated to this service. eg: main
     */
    gitlabRepositoryBranch?: pulumi.Input<string>;
    /**
     * The Gitlab repository name associated to this service. eg: rootlyhq/my-service
     */
    gitlabRepositoryName?: pulumi.Input<string>;
    /**
     * The name of the service
     */
    name?: pulumi.Input<string>;
    /**
     * Emails attached to the service
     */
    notifyEmails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Opsgenie service id associated to this service
     */
    opsgenieId?: pulumi.Input<string>;
    /**
     * Owner Teams associated with this service
     */
    ownersGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Owner Users associated with this service
     */
    ownersUserIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The PagerDuty service id associated to this service
     */
    pagerdutyId?: pulumi.Input<string>;
    /**
     * Position of the service
     */
    position?: pulumi.Input<number>;
    /**
     * The public description of the service
     */
    publicDescription?: pulumi.Input<string>;
    /**
     * Services dependent on this service
     */
    serviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Slack Aliases associated with this service
     */
    slackAliases?: pulumi.Input<pulumi.Input<inputs.ServiceSlackAlias>[]>;
    /**
     * Slack Channels associated with this service
     */
    slackChannels?: pulumi.Input<pulumi.Input<inputs.ServiceSlackChannel>[]>;
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
     * The Backstage entity id associated to this service. eg: :namespace/:kind/:entity_name
     */
    backstageId?: pulumi.Input<string>;
    /**
     * The hex color of the service
     */
    color?: pulumi.Input<string>;
    /**
     * The description of the service
     */
    description?: pulumi.Input<string>;
    /**
     * Environments associated with this service
     */
    environmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The GitHub repository branch associated to this service. eg: main
     */
    githubRepositoryBranch?: pulumi.Input<string>;
    /**
     * The GitHub repository name associated to this service. eg: rootlyhq/my-service
     */
    githubRepositoryName?: pulumi.Input<string>;
    /**
     * The Gitlab repository branch associated to this service. eg: main
     */
    gitlabRepositoryBranch?: pulumi.Input<string>;
    /**
     * The Gitlab repository name associated to this service. eg: rootlyhq/my-service
     */
    gitlabRepositoryName?: pulumi.Input<string>;
    /**
     * The name of the service
     */
    name?: pulumi.Input<string>;
    /**
     * Emails attached to the service
     */
    notifyEmails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Opsgenie service id associated to this service
     */
    opsgenieId?: pulumi.Input<string>;
    /**
     * Owner Teams associated with this service
     */
    ownersGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Owner Users associated with this service
     */
    ownersUserIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The PagerDuty service id associated to this service
     */
    pagerdutyId?: pulumi.Input<string>;
    /**
     * Position of the service
     */
    position?: pulumi.Input<number>;
    /**
     * The public description of the service
     */
    publicDescription?: pulumi.Input<string>;
    /**
     * Services dependent on this service
     */
    serviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Slack Aliases associated with this service
     */
    slackAliases?: pulumi.Input<pulumi.Input<inputs.ServiceSlackAlias>[]>;
    /**
     * Slack Channels associated with this service
     */
    slackChannels?: pulumi.Input<pulumi.Input<inputs.ServiceSlackChannel>[]>;
    /**
     * The slug of the service
     */
    slug?: pulumi.Input<string>;
}
