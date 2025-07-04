// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    [RootlyResourceType("rootly:index/incidentType:IncidentType")]
    public partial class IncidentType : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The hex color of the incident type
        /// </summary>
        [Output("color")]
        public Output<string> Color { get; private set; } = null!;

        /// <summary>
        /// The description of the incident type
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the incident type
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Emails to attach to the incident type
        /// </summary>
        [Output("notifyEmails")]
        public Output<ImmutableArray<string>> NotifyEmails { get; private set; } = null!;

        /// <summary>
        /// Position of the incident type
        /// </summary>
        [Output("position")]
        public Output<int> Position { get; private set; } = null!;

        /// <summary>
        /// Slack Aliases associated with this incident type
        /// </summary>
        [Output("slackAliases")]
        public Output<ImmutableArray<Outputs.IncidentTypeSlackAlias>> SlackAliases { get; private set; } = null!;

        /// <summary>
        /// Slack Channels associated with this incident type
        /// </summary>
        [Output("slackChannels")]
        public Output<ImmutableArray<Outputs.IncidentTypeSlackChannel>> SlackChannels { get; private set; } = null!;

        /// <summary>
        /// The slug of the incident type
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;


        /// <summary>
        /// Create a IncidentType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IncidentType(string name, IncidentTypeArgs? args = null, CustomResourceOptions? options = null)
            : base("rootly:index/incidentType:IncidentType", name, args ?? new IncidentTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IncidentType(string name, Input<string> id, IncidentTypeState? state = null, CustomResourceOptions? options = null)
            : base("rootly:index/incidentType:IncidentType", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/rootlyhq/pulumi-rootly/releases/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IncidentType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IncidentType Get(string name, Input<string> id, IncidentTypeState? state = null, CustomResourceOptions? options = null)
        {
            return new IncidentType(name, id, state, options);
        }
    }

    public sealed class IncidentTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hex color of the incident type
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// The description of the incident type
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the incident type
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notifyEmails")]
        private InputList<string>? _notifyEmails;

        /// <summary>
        /// Emails to attach to the incident type
        /// </summary>
        public InputList<string> NotifyEmails
        {
            get => _notifyEmails ?? (_notifyEmails = new InputList<string>());
            set => _notifyEmails = value;
        }

        /// <summary>
        /// Position of the incident type
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        [Input("slackAliases")]
        private InputList<Inputs.IncidentTypeSlackAliasArgs>? _slackAliases;

        /// <summary>
        /// Slack Aliases associated with this incident type
        /// </summary>
        public InputList<Inputs.IncidentTypeSlackAliasArgs> SlackAliases
        {
            get => _slackAliases ?? (_slackAliases = new InputList<Inputs.IncidentTypeSlackAliasArgs>());
            set => _slackAliases = value;
        }

        [Input("slackChannels")]
        private InputList<Inputs.IncidentTypeSlackChannelArgs>? _slackChannels;

        /// <summary>
        /// Slack Channels associated with this incident type
        /// </summary>
        public InputList<Inputs.IncidentTypeSlackChannelArgs> SlackChannels
        {
            get => _slackChannels ?? (_slackChannels = new InputList<Inputs.IncidentTypeSlackChannelArgs>());
            set => _slackChannels = value;
        }

        /// <summary>
        /// The slug of the incident type
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        public IncidentTypeArgs()
        {
        }
        public static new IncidentTypeArgs Empty => new IncidentTypeArgs();
    }

    public sealed class IncidentTypeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hex color of the incident type
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// The description of the incident type
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the incident type
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notifyEmails")]
        private InputList<string>? _notifyEmails;

        /// <summary>
        /// Emails to attach to the incident type
        /// </summary>
        public InputList<string> NotifyEmails
        {
            get => _notifyEmails ?? (_notifyEmails = new InputList<string>());
            set => _notifyEmails = value;
        }

        /// <summary>
        /// Position of the incident type
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        [Input("slackAliases")]
        private InputList<Inputs.IncidentTypeSlackAliasGetArgs>? _slackAliases;

        /// <summary>
        /// Slack Aliases associated with this incident type
        /// </summary>
        public InputList<Inputs.IncidentTypeSlackAliasGetArgs> SlackAliases
        {
            get => _slackAliases ?? (_slackAliases = new InputList<Inputs.IncidentTypeSlackAliasGetArgs>());
            set => _slackAliases = value;
        }

        [Input("slackChannels")]
        private InputList<Inputs.IncidentTypeSlackChannelGetArgs>? _slackChannels;

        /// <summary>
        /// Slack Channels associated with this incident type
        /// </summary>
        public InputList<Inputs.IncidentTypeSlackChannelGetArgs> SlackChannels
        {
            get => _slackChannels ?? (_slackChannels = new InputList<Inputs.IncidentTypeSlackChannelGetArgs>());
            set => _slackChannels = value;
        }

        /// <summary>
        /// The slug of the incident type
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        public IncidentTypeState()
        {
        }
        public static new IncidentTypeState Empty => new IncidentTypeState();
    }
}
