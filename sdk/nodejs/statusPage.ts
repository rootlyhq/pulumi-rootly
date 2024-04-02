// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class StatusPage extends pulumi.CustomResource {
    /**
     * Get an existing StatusPage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StatusPageState, opts?: pulumi.CustomResourceOptions): StatusPage {
        return new StatusPage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rootly:index/statusPage:StatusPage';

    /**
     * Returns true if the given object is an instance of StatusPage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StatusPage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StatusPage.__pulumiType;
    }

    /**
     * Allow search engines to include your public status page in search results. Value must be one of true or false
     */
    public readonly allowSearchEngineIndex!: pulumi.Output<boolean>;
    /**
     * Enable authentication. Value must be one of true or false
     */
    public readonly authenticationEnabled!: pulumi.Output<boolean>;
    /**
     * Authentication password
     */
    public readonly authenticationPassword!: pulumi.Output<string>;
    /**
     * The description of the status page
     */
    public readonly description!: pulumi.Output<string>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Message showing when at least one component is not operational
     */
    public readonly failureMessage!: pulumi.Output<string>;
    /**
     * The color of the footer. Eg. "#1F2F41"
     */
    public readonly footerColor!: pulumi.Output<string>;
    /**
     * Functionalities attached to the status page
     */
    public readonly functionalityIds!: pulumi.Output<string[]>;
    /**
     * Google Analytics tracking ID
     */
    public readonly gaTrackingId!: pulumi.Output<string>;
    /**
     * The color of the header. Eg. "#0061F2"
     */
    public readonly headerColor!: pulumi.Output<string>;
    /**
     * Make the status page accessible to the public. Value must be one of true or false
     */
    public readonly public!: pulumi.Output<boolean>;
    /**
     * The public description of the status page
     */
    public readonly publicDescription!: pulumi.Output<string>;
    /**
     * The public title of the status page
     */
    public readonly publicTitle!: pulumi.Output<string>;
    /**
     * Services attached to the status page
     */
    public readonly serviceIds!: pulumi.Output<string[]>;
    /**
     * Show uptime. Value must be one of true or false
     */
    public readonly showUptime!: pulumi.Output<boolean>;
    /**
     * Show uptime over x days. Value must be one of `30`, `60`, `90`, `180`, `360`.
     */
    public readonly showUptimeLastDays!: pulumi.Output<number>;
    /**
     * Message showing when all components are operational
     */
    public readonly successMessage!: pulumi.Output<string>;
    /**
     * Status Page Timezone. Value must be one of `International Date Line West`, `American Samoa`, `Midway Island`, `Hawaii`, `Alaska`, `Pacific Time (US & Canada)`, `Tijuana`, `Arizona`, `Mazatlan`, `Mountain Time (US & Canada)`, `Central America`, `Central Time (US & Canada)`, `Chihuahua`, `Guadalajara`, `Mexico City`, `Monterrey`, `Saskatchewan`, `Bogota`, `Eastern Time (US & Canada)`, `Indiana (East)`, `Lima`, `Quito`, `Atlantic Time (Canada)`, `Caracas`, `Georgetown`, `La Paz`, `Puerto Rico`, `Santiago`, `Newfoundland`, `Brasilia`, `Buenos Aires`, `Montevideo`, `Greenland`, `Mid-Atlantic`, `Azores`, `Cape Verde Is.`, `Edinburgh`, `Lisbon`, `London`, `Monrovia`, `UTC`, `Amsterdam`, `Belgrade`, `Berlin`, `Bern`, `Bratislava`, `Brussels`, `Budapest`, `Casablanca`, `Copenhagen`, `Dublin`, `Ljubljana`, `Madrid`, `Paris`, `Prague`, `Rome`, `Sarajevo`, `Skopje`, `Stockholm`, `Vienna`, `Warsaw`, `West Central Africa`, `Zagreb`, `Zurich`, `Athens`, `Bucharest`, `Cairo`, `Harare`, `Helsinki`, `Jerusalem`, `Kaliningrad`, `Kyiv`, `Pretoria`, `Riga`, `Sofia`, `Tallinn`, `Vilnius`, `Baghdad`, `Istanbul`, `Kuwait`, `Minsk`, `Moscow`, `Nairobi`, `Riyadh`, `St. Petersburg`, `Volgograd`, `Tehran`, `Abu Dhabi`, `Baku`, `Muscat`, `Samara`, `Tbilisi`, `Yerevan`, `Kabul`, `Almaty`, `Ekaterinburg`, `Islamabad`, `Karachi`, `Tashkent`, `Chennai`, `Kolkata`, `Mumbai`, `New Delhi`, `Sri Jayawardenepura`, `Kathmandu`, `Astana`, `Dhaka`, `Urumqi`, `Rangoon`, `Bangkok`, `Hanoi`, `Jakarta`, `Krasnoyarsk`, `Novosibirsk`, `Beijing`, `Chongqing`, `Hong Kong`, `Irkutsk`, `Kuala Lumpur`, `Perth`, `Singapore`, `Taipei`, `Ulaanbaatar`, `Osaka`, `Sapporo`, `Seoul`, `Tokyo`, `Yakutsk`, `Adelaide`, `Darwin`, `Brisbane`, `Canberra`, `Guam`, `Hobart`, `Melbourne`, `Port Moresby`, `Sydney`, `Vladivostok`, `Magadan`, `New Caledonia`, `Solomon Is.`, `Srednekolymsk`, `Auckland`, `Fiji`, `Kamchatka`, `Marshall Is.`, `Wellington`, `Chatham Is.`, `Nuku'alofa`, `Samoa`, `Tokelau Is.`.
     */
    public readonly timeZone!: pulumi.Output<string | undefined>;
    /**
     * The title of the status page
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Website Privacy URL
     */
    public readonly websitePrivacyUrl!: pulumi.Output<string>;
    /**
     * Website Support URL
     */
    public readonly websiteSupportUrl!: pulumi.Output<string>;
    /**
     * Website URL
     */
    public readonly websiteUrl!: pulumi.Output<string>;

    /**
     * Create a StatusPage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StatusPageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StatusPageArgs | StatusPageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StatusPageState | undefined;
            resourceInputs["allowSearchEngineIndex"] = state ? state.allowSearchEngineIndex : undefined;
            resourceInputs["authenticationEnabled"] = state ? state.authenticationEnabled : undefined;
            resourceInputs["authenticationPassword"] = state ? state.authenticationPassword : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["failureMessage"] = state ? state.failureMessage : undefined;
            resourceInputs["footerColor"] = state ? state.footerColor : undefined;
            resourceInputs["functionalityIds"] = state ? state.functionalityIds : undefined;
            resourceInputs["gaTrackingId"] = state ? state.gaTrackingId : undefined;
            resourceInputs["headerColor"] = state ? state.headerColor : undefined;
            resourceInputs["public"] = state ? state.public : undefined;
            resourceInputs["publicDescription"] = state ? state.publicDescription : undefined;
            resourceInputs["publicTitle"] = state ? state.publicTitle : undefined;
            resourceInputs["serviceIds"] = state ? state.serviceIds : undefined;
            resourceInputs["showUptime"] = state ? state.showUptime : undefined;
            resourceInputs["showUptimeLastDays"] = state ? state.showUptimeLastDays : undefined;
            resourceInputs["successMessage"] = state ? state.successMessage : undefined;
            resourceInputs["timeZone"] = state ? state.timeZone : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["websitePrivacyUrl"] = state ? state.websitePrivacyUrl : undefined;
            resourceInputs["websiteSupportUrl"] = state ? state.websiteSupportUrl : undefined;
            resourceInputs["websiteUrl"] = state ? state.websiteUrl : undefined;
        } else {
            const args = argsOrState as StatusPageArgs | undefined;
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["allowSearchEngineIndex"] = args ? args.allowSearchEngineIndex : undefined;
            resourceInputs["authenticationEnabled"] = args ? args.authenticationEnabled : undefined;
            resourceInputs["authenticationPassword"] = args ? args.authenticationPassword : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["failureMessage"] = args ? args.failureMessage : undefined;
            resourceInputs["footerColor"] = args ? args.footerColor : undefined;
            resourceInputs["functionalityIds"] = args ? args.functionalityIds : undefined;
            resourceInputs["gaTrackingId"] = args ? args.gaTrackingId : undefined;
            resourceInputs["headerColor"] = args ? args.headerColor : undefined;
            resourceInputs["public"] = args ? args.public : undefined;
            resourceInputs["publicDescription"] = args ? args.publicDescription : undefined;
            resourceInputs["publicTitle"] = args ? args.publicTitle : undefined;
            resourceInputs["serviceIds"] = args ? args.serviceIds : undefined;
            resourceInputs["showUptime"] = args ? args.showUptime : undefined;
            resourceInputs["showUptimeLastDays"] = args ? args.showUptimeLastDays : undefined;
            resourceInputs["successMessage"] = args ? args.successMessage : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["websitePrivacyUrl"] = args ? args.websitePrivacyUrl : undefined;
            resourceInputs["websiteSupportUrl"] = args ? args.websiteSupportUrl : undefined;
            resourceInputs["websiteUrl"] = args ? args.websiteUrl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StatusPage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StatusPage resources.
 */
export interface StatusPageState {
    /**
     * Allow search engines to include your public status page in search results. Value must be one of true or false
     */
    allowSearchEngineIndex?: pulumi.Input<boolean>;
    /**
     * Enable authentication. Value must be one of true or false
     */
    authenticationEnabled?: pulumi.Input<boolean>;
    /**
     * Authentication password
     */
    authenticationPassword?: pulumi.Input<string>;
    /**
     * The description of the status page
     */
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    /**
     * Message showing when at least one component is not operational
     */
    failureMessage?: pulumi.Input<string>;
    /**
     * The color of the footer. Eg. "#1F2F41"
     */
    footerColor?: pulumi.Input<string>;
    /**
     * Functionalities attached to the status page
     */
    functionalityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Google Analytics tracking ID
     */
    gaTrackingId?: pulumi.Input<string>;
    /**
     * The color of the header. Eg. "#0061F2"
     */
    headerColor?: pulumi.Input<string>;
    /**
     * Make the status page accessible to the public. Value must be one of true or false
     */
    public?: pulumi.Input<boolean>;
    /**
     * The public description of the status page
     */
    publicDescription?: pulumi.Input<string>;
    /**
     * The public title of the status page
     */
    publicTitle?: pulumi.Input<string>;
    /**
     * Services attached to the status page
     */
    serviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Show uptime. Value must be one of true or false
     */
    showUptime?: pulumi.Input<boolean>;
    /**
     * Show uptime over x days. Value must be one of `30`, `60`, `90`, `180`, `360`.
     */
    showUptimeLastDays?: pulumi.Input<number>;
    /**
     * Message showing when all components are operational
     */
    successMessage?: pulumi.Input<string>;
    /**
     * Status Page Timezone. Value must be one of `International Date Line West`, `American Samoa`, `Midway Island`, `Hawaii`, `Alaska`, `Pacific Time (US & Canada)`, `Tijuana`, `Arizona`, `Mazatlan`, `Mountain Time (US & Canada)`, `Central America`, `Central Time (US & Canada)`, `Chihuahua`, `Guadalajara`, `Mexico City`, `Monterrey`, `Saskatchewan`, `Bogota`, `Eastern Time (US & Canada)`, `Indiana (East)`, `Lima`, `Quito`, `Atlantic Time (Canada)`, `Caracas`, `Georgetown`, `La Paz`, `Puerto Rico`, `Santiago`, `Newfoundland`, `Brasilia`, `Buenos Aires`, `Montevideo`, `Greenland`, `Mid-Atlantic`, `Azores`, `Cape Verde Is.`, `Edinburgh`, `Lisbon`, `London`, `Monrovia`, `UTC`, `Amsterdam`, `Belgrade`, `Berlin`, `Bern`, `Bratislava`, `Brussels`, `Budapest`, `Casablanca`, `Copenhagen`, `Dublin`, `Ljubljana`, `Madrid`, `Paris`, `Prague`, `Rome`, `Sarajevo`, `Skopje`, `Stockholm`, `Vienna`, `Warsaw`, `West Central Africa`, `Zagreb`, `Zurich`, `Athens`, `Bucharest`, `Cairo`, `Harare`, `Helsinki`, `Jerusalem`, `Kaliningrad`, `Kyiv`, `Pretoria`, `Riga`, `Sofia`, `Tallinn`, `Vilnius`, `Baghdad`, `Istanbul`, `Kuwait`, `Minsk`, `Moscow`, `Nairobi`, `Riyadh`, `St. Petersburg`, `Volgograd`, `Tehran`, `Abu Dhabi`, `Baku`, `Muscat`, `Samara`, `Tbilisi`, `Yerevan`, `Kabul`, `Almaty`, `Ekaterinburg`, `Islamabad`, `Karachi`, `Tashkent`, `Chennai`, `Kolkata`, `Mumbai`, `New Delhi`, `Sri Jayawardenepura`, `Kathmandu`, `Astana`, `Dhaka`, `Urumqi`, `Rangoon`, `Bangkok`, `Hanoi`, `Jakarta`, `Krasnoyarsk`, `Novosibirsk`, `Beijing`, `Chongqing`, `Hong Kong`, `Irkutsk`, `Kuala Lumpur`, `Perth`, `Singapore`, `Taipei`, `Ulaanbaatar`, `Osaka`, `Sapporo`, `Seoul`, `Tokyo`, `Yakutsk`, `Adelaide`, `Darwin`, `Brisbane`, `Canberra`, `Guam`, `Hobart`, `Melbourne`, `Port Moresby`, `Sydney`, `Vladivostok`, `Magadan`, `New Caledonia`, `Solomon Is.`, `Srednekolymsk`, `Auckland`, `Fiji`, `Kamchatka`, `Marshall Is.`, `Wellington`, `Chatham Is.`, `Nuku'alofa`, `Samoa`, `Tokelau Is.`.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * The title of the status page
     */
    title?: pulumi.Input<string>;
    /**
     * Website Privacy URL
     */
    websitePrivacyUrl?: pulumi.Input<string>;
    /**
     * Website Support URL
     */
    websiteSupportUrl?: pulumi.Input<string>;
    /**
     * Website URL
     */
    websiteUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StatusPage resource.
 */
export interface StatusPageArgs {
    /**
     * Allow search engines to include your public status page in search results. Value must be one of true or false
     */
    allowSearchEngineIndex?: pulumi.Input<boolean>;
    /**
     * Enable authentication. Value must be one of true or false
     */
    authenticationEnabled?: pulumi.Input<boolean>;
    /**
     * Authentication password
     */
    authenticationPassword?: pulumi.Input<string>;
    /**
     * The description of the status page
     */
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    /**
     * Message showing when at least one component is not operational
     */
    failureMessage?: pulumi.Input<string>;
    /**
     * The color of the footer. Eg. "#1F2F41"
     */
    footerColor?: pulumi.Input<string>;
    /**
     * Functionalities attached to the status page
     */
    functionalityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Google Analytics tracking ID
     */
    gaTrackingId?: pulumi.Input<string>;
    /**
     * The color of the header. Eg. "#0061F2"
     */
    headerColor?: pulumi.Input<string>;
    /**
     * Make the status page accessible to the public. Value must be one of true or false
     */
    public?: pulumi.Input<boolean>;
    /**
     * The public description of the status page
     */
    publicDescription?: pulumi.Input<string>;
    /**
     * The public title of the status page
     */
    publicTitle?: pulumi.Input<string>;
    /**
     * Services attached to the status page
     */
    serviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Show uptime. Value must be one of true or false
     */
    showUptime?: pulumi.Input<boolean>;
    /**
     * Show uptime over x days. Value must be one of `30`, `60`, `90`, `180`, `360`.
     */
    showUptimeLastDays?: pulumi.Input<number>;
    /**
     * Message showing when all components are operational
     */
    successMessage?: pulumi.Input<string>;
    /**
     * Status Page Timezone. Value must be one of `International Date Line West`, `American Samoa`, `Midway Island`, `Hawaii`, `Alaska`, `Pacific Time (US & Canada)`, `Tijuana`, `Arizona`, `Mazatlan`, `Mountain Time (US & Canada)`, `Central America`, `Central Time (US & Canada)`, `Chihuahua`, `Guadalajara`, `Mexico City`, `Monterrey`, `Saskatchewan`, `Bogota`, `Eastern Time (US & Canada)`, `Indiana (East)`, `Lima`, `Quito`, `Atlantic Time (Canada)`, `Caracas`, `Georgetown`, `La Paz`, `Puerto Rico`, `Santiago`, `Newfoundland`, `Brasilia`, `Buenos Aires`, `Montevideo`, `Greenland`, `Mid-Atlantic`, `Azores`, `Cape Verde Is.`, `Edinburgh`, `Lisbon`, `London`, `Monrovia`, `UTC`, `Amsterdam`, `Belgrade`, `Berlin`, `Bern`, `Bratislava`, `Brussels`, `Budapest`, `Casablanca`, `Copenhagen`, `Dublin`, `Ljubljana`, `Madrid`, `Paris`, `Prague`, `Rome`, `Sarajevo`, `Skopje`, `Stockholm`, `Vienna`, `Warsaw`, `West Central Africa`, `Zagreb`, `Zurich`, `Athens`, `Bucharest`, `Cairo`, `Harare`, `Helsinki`, `Jerusalem`, `Kaliningrad`, `Kyiv`, `Pretoria`, `Riga`, `Sofia`, `Tallinn`, `Vilnius`, `Baghdad`, `Istanbul`, `Kuwait`, `Minsk`, `Moscow`, `Nairobi`, `Riyadh`, `St. Petersburg`, `Volgograd`, `Tehran`, `Abu Dhabi`, `Baku`, `Muscat`, `Samara`, `Tbilisi`, `Yerevan`, `Kabul`, `Almaty`, `Ekaterinburg`, `Islamabad`, `Karachi`, `Tashkent`, `Chennai`, `Kolkata`, `Mumbai`, `New Delhi`, `Sri Jayawardenepura`, `Kathmandu`, `Astana`, `Dhaka`, `Urumqi`, `Rangoon`, `Bangkok`, `Hanoi`, `Jakarta`, `Krasnoyarsk`, `Novosibirsk`, `Beijing`, `Chongqing`, `Hong Kong`, `Irkutsk`, `Kuala Lumpur`, `Perth`, `Singapore`, `Taipei`, `Ulaanbaatar`, `Osaka`, `Sapporo`, `Seoul`, `Tokyo`, `Yakutsk`, `Adelaide`, `Darwin`, `Brisbane`, `Canberra`, `Guam`, `Hobart`, `Melbourne`, `Port Moresby`, `Sydney`, `Vladivostok`, `Magadan`, `New Caledonia`, `Solomon Is.`, `Srednekolymsk`, `Auckland`, `Fiji`, `Kamchatka`, `Marshall Is.`, `Wellington`, `Chatham Is.`, `Nuku'alofa`, `Samoa`, `Tokelau Is.`.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * The title of the status page
     */
    title: pulumi.Input<string>;
    /**
     * Website Privacy URL
     */
    websitePrivacyUrl?: pulumi.Input<string>;
    /**
     * Website Support URL
     */
    websiteSupportUrl?: pulumi.Input<string>;
    /**
     * Website URL
     */
    websiteUrl?: pulumi.Input<string>;
}
