// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    [RootlyResourceType("rootly:index/escalationLevel:EscalationLevel")]
    public partial class EscalationLevel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Delay before notification targets will be alerted.
        /// </summary>
        [Output("delay")]
        public Output<int> Delay { get; private set; } = null!;

        /// <summary>
        /// The ID of the escalation policy
        /// </summary>
        [Output("escalationPolicyId")]
        public Output<string> EscalationPolicyId { get; private set; } = null!;

        /// <summary>
        /// Escalation level's notification targets
        /// </summary>
        [Output("notificationTargetParams")]
        public Output<ImmutableArray<Outputs.EscalationLevelNotificationTargetParam>> NotificationTargetParams { get; private set; } = null!;

        /// <summary>
        /// Position of the escalation policy level
        /// </summary>
        [Output("position")]
        public Output<int> Position { get; private set; } = null!;


        /// <summary>
        /// Create a EscalationLevel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EscalationLevel(string name, EscalationLevelArgs args, CustomResourceOptions? options = null)
            : base("rootly:index/escalationLevel:EscalationLevel", name, args ?? new EscalationLevelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EscalationLevel(string name, Input<string> id, EscalationLevelState? state = null, CustomResourceOptions? options = null)
            : base("rootly:index/escalationLevel:EscalationLevel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EscalationLevel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EscalationLevel Get(string name, Input<string> id, EscalationLevelState? state = null, CustomResourceOptions? options = null)
        {
            return new EscalationLevel(name, id, state, options);
        }
    }

    public sealed class EscalationLevelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Delay before notification targets will be alerted.
        /// </summary>
        [Input("delay")]
        public Input<int>? Delay { get; set; }

        /// <summary>
        /// The ID of the escalation policy
        /// </summary>
        [Input("escalationPolicyId")]
        public Input<string>? EscalationPolicyId { get; set; }

        [Input("notificationTargetParams", required: true)]
        private InputList<Inputs.EscalationLevelNotificationTargetParamArgs>? _notificationTargetParams;

        /// <summary>
        /// Escalation level's notification targets
        /// </summary>
        public InputList<Inputs.EscalationLevelNotificationTargetParamArgs> NotificationTargetParams
        {
            get => _notificationTargetParams ?? (_notificationTargetParams = new InputList<Inputs.EscalationLevelNotificationTargetParamArgs>());
            set => _notificationTargetParams = value;
        }

        /// <summary>
        /// Position of the escalation policy level
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        public EscalationLevelArgs()
        {
        }
        public static new EscalationLevelArgs Empty => new EscalationLevelArgs();
    }

    public sealed class EscalationLevelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Delay before notification targets will be alerted.
        /// </summary>
        [Input("delay")]
        public Input<int>? Delay { get; set; }

        /// <summary>
        /// The ID of the escalation policy
        /// </summary>
        [Input("escalationPolicyId")]
        public Input<string>? EscalationPolicyId { get; set; }

        [Input("notificationTargetParams")]
        private InputList<Inputs.EscalationLevelNotificationTargetParamGetArgs>? _notificationTargetParams;

        /// <summary>
        /// Escalation level's notification targets
        /// </summary>
        public InputList<Inputs.EscalationLevelNotificationTargetParamGetArgs> NotificationTargetParams
        {
            get => _notificationTargetParams ?? (_notificationTargetParams = new InputList<Inputs.EscalationLevelNotificationTargetParamGetArgs>());
            set => _notificationTargetParams = value;
        }

        /// <summary>
        /// Position of the escalation policy level
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        public EscalationLevelState()
        {
        }
        public static new EscalationLevelState Empty => new EscalationLevelState();
    }
}
