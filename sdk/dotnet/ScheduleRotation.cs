// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    [RootlyResourceType("rootly:index/scheduleRotation:ScheduleRotation")]
    public partial class ScheduleRotation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Schedule rotation active all week?. Value must be one of true or false
        /// </summary>
        [Output("activeAllWeek")]
        public Output<bool> ActiveAllWeek { get; private set; } = null!;

        /// <summary>
        /// Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        /// </summary>
        [Output("activeDays")]
        public Output<ImmutableArray<string>> ActiveDays { get; private set; } = null!;

        /// <summary>
        /// Schedule rotation's active times
        /// </summary>
        [Output("activeTimeAttributes")]
        public Output<ImmutableArray<Outputs.ScheduleRotationActiveTimeAttribute>> ActiveTimeAttributes { get; private set; } = null!;

        [Output("activeTimeType")]
        public Output<string> ActiveTimeType { get; private set; } = null!;

        /// <summary>
        /// The name of the schedule rotation
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Position of the schedule rotation
        /// </summary>
        [Output("position")]
        public Output<int> Position { get; private set; } = null!;

        /// <summary>
        /// The ID of parent schedule
        /// </summary>
        [Output("scheduleId")]
        public Output<string> ScheduleId { get; private set; } = null!;

        [Output("scheduleRotationableAttributes")]
        public Output<ImmutableDictionary<string, string>> ScheduleRotationableAttributes { get; private set; } = null!;

        /// <summary>
        /// Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
        /// </summary>
        [Output("scheduleRotationableType")]
        public Output<string?> ScheduleRotationableType { get; private set; } = null!;

        /// <summary>
        /// A valid IANA time zone name.
        /// </summary>
        [Output("timeZone")]
        public Output<string?> TimeZone { get; private set; } = null!;


        /// <summary>
        /// Create a ScheduleRotation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScheduleRotation(string name, ScheduleRotationArgs args, CustomResourceOptions? options = null)
            : base("rootly:index/scheduleRotation:ScheduleRotation", name, args ?? new ScheduleRotationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScheduleRotation(string name, Input<string> id, ScheduleRotationState? state = null, CustomResourceOptions? options = null)
            : base("rootly:index/scheduleRotation:ScheduleRotation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScheduleRotation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScheduleRotation Get(string name, Input<string> id, ScheduleRotationState? state = null, CustomResourceOptions? options = null)
        {
            return new ScheduleRotation(name, id, state, options);
        }
    }

    public sealed class ScheduleRotationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Schedule rotation active all week?. Value must be one of true or false
        /// </summary>
        [Input("activeAllWeek")]
        public Input<bool>? ActiveAllWeek { get; set; }

        [Input("activeDays")]
        private InputList<string>? _activeDays;

        /// <summary>
        /// Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        /// </summary>
        public InputList<string> ActiveDays
        {
            get => _activeDays ?? (_activeDays = new InputList<string>());
            set => _activeDays = value;
        }

        [Input("activeTimeAttributes")]
        private InputList<Inputs.ScheduleRotationActiveTimeAttributeArgs>? _activeTimeAttributes;

        /// <summary>
        /// Schedule rotation's active times
        /// </summary>
        public InputList<Inputs.ScheduleRotationActiveTimeAttributeArgs> ActiveTimeAttributes
        {
            get => _activeTimeAttributes ?? (_activeTimeAttributes = new InputList<Inputs.ScheduleRotationActiveTimeAttributeArgs>());
            set => _activeTimeAttributes = value;
        }

        [Input("activeTimeType")]
        public Input<string>? ActiveTimeType { get; set; }

        /// <summary>
        /// The name of the schedule rotation
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Position of the schedule rotation
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// The ID of parent schedule
        /// </summary>
        [Input("scheduleId")]
        public Input<string>? ScheduleId { get; set; }

        [Input("scheduleRotationableAttributes", required: true)]
        private InputMap<string>? _scheduleRotationableAttributes;
        public InputMap<string> ScheduleRotationableAttributes
        {
            get => _scheduleRotationableAttributes ?? (_scheduleRotationableAttributes = new InputMap<string>());
            set => _scheduleRotationableAttributes = value;
        }

        /// <summary>
        /// Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
        /// </summary>
        [Input("scheduleRotationableType")]
        public Input<string>? ScheduleRotationableType { get; set; }

        /// <summary>
        /// A valid IANA time zone name.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public ScheduleRotationArgs()
        {
        }
        public static new ScheduleRotationArgs Empty => new ScheduleRotationArgs();
    }

    public sealed class ScheduleRotationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Schedule rotation active all week?. Value must be one of true or false
        /// </summary>
        [Input("activeAllWeek")]
        public Input<bool>? ActiveAllWeek { get; set; }

        [Input("activeDays")]
        private InputList<string>? _activeDays;

        /// <summary>
        /// Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        /// </summary>
        public InputList<string> ActiveDays
        {
            get => _activeDays ?? (_activeDays = new InputList<string>());
            set => _activeDays = value;
        }

        [Input("activeTimeAttributes")]
        private InputList<Inputs.ScheduleRotationActiveTimeAttributeGetArgs>? _activeTimeAttributes;

        /// <summary>
        /// Schedule rotation's active times
        /// </summary>
        public InputList<Inputs.ScheduleRotationActiveTimeAttributeGetArgs> ActiveTimeAttributes
        {
            get => _activeTimeAttributes ?? (_activeTimeAttributes = new InputList<Inputs.ScheduleRotationActiveTimeAttributeGetArgs>());
            set => _activeTimeAttributes = value;
        }

        [Input("activeTimeType")]
        public Input<string>? ActiveTimeType { get; set; }

        /// <summary>
        /// The name of the schedule rotation
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Position of the schedule rotation
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// The ID of parent schedule
        /// </summary>
        [Input("scheduleId")]
        public Input<string>? ScheduleId { get; set; }

        [Input("scheduleRotationableAttributes")]
        private InputMap<string>? _scheduleRotationableAttributes;
        public InputMap<string> ScheduleRotationableAttributes
        {
            get => _scheduleRotationableAttributes ?? (_scheduleRotationableAttributes = new InputMap<string>());
            set => _scheduleRotationableAttributes = value;
        }

        /// <summary>
        /// Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
        /// </summary>
        [Input("scheduleRotationableType")]
        public Input<string>? ScheduleRotationableType { get; set; }

        /// <summary>
        /// A valid IANA time zone name.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public ScheduleRotationState()
        {
        }
        public static new ScheduleRotationState Empty => new ScheduleRotationState();
    }
}
