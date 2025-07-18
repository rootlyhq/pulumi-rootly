// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    [RootlyResourceType("rootly:index/scheduleRotationActiveDay:ScheduleRotationActiveDay")]
    public partial class ScheduleRotationActiveDay : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Schedule rotation active times per day
        /// </summary>
        [Output("activeTimeAttributes")]
        public Output<ImmutableArray<Outputs.ScheduleRotationActiveDayActiveTimeAttribute>> ActiveTimeAttributes { get; private set; } = null!;

        /// <summary>
        /// Schedule rotation day name for which active times to be created. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        /// </summary>
        [Output("dayName")]
        public Output<string?> DayName { get; private set; } = null!;

        [Output("scheduleRotationId")]
        public Output<string> ScheduleRotationId { get; private set; } = null!;


        /// <summary>
        /// Create a ScheduleRotationActiveDay resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScheduleRotationActiveDay(string name, ScheduleRotationActiveDayArgs args, CustomResourceOptions? options = null)
            : base("rootly:index/scheduleRotationActiveDay:ScheduleRotationActiveDay", name, args ?? new ScheduleRotationActiveDayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScheduleRotationActiveDay(string name, Input<string> id, ScheduleRotationActiveDayState? state = null, CustomResourceOptions? options = null)
            : base("rootly:index/scheduleRotationActiveDay:ScheduleRotationActiveDay", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScheduleRotationActiveDay resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScheduleRotationActiveDay Get(string name, Input<string> id, ScheduleRotationActiveDayState? state = null, CustomResourceOptions? options = null)
        {
            return new ScheduleRotationActiveDay(name, id, state, options);
        }
    }

    public sealed class ScheduleRotationActiveDayArgs : global::Pulumi.ResourceArgs
    {
        [Input("activeTimeAttributes", required: true)]
        private InputList<Inputs.ScheduleRotationActiveDayActiveTimeAttributeArgs>? _activeTimeAttributes;

        /// <summary>
        /// Schedule rotation active times per day
        /// </summary>
        public InputList<Inputs.ScheduleRotationActiveDayActiveTimeAttributeArgs> ActiveTimeAttributes
        {
            get => _activeTimeAttributes ?? (_activeTimeAttributes = new InputList<Inputs.ScheduleRotationActiveDayActiveTimeAttributeArgs>());
            set => _activeTimeAttributes = value;
        }

        /// <summary>
        /// Schedule rotation day name for which active times to be created. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        /// </summary>
        [Input("dayName")]
        public Input<string>? DayName { get; set; }

        [Input("scheduleRotationId")]
        public Input<string>? ScheduleRotationId { get; set; }

        public ScheduleRotationActiveDayArgs()
        {
        }
        public static new ScheduleRotationActiveDayArgs Empty => new ScheduleRotationActiveDayArgs();
    }

    public sealed class ScheduleRotationActiveDayState : global::Pulumi.ResourceArgs
    {
        [Input("activeTimeAttributes")]
        private InputList<Inputs.ScheduleRotationActiveDayActiveTimeAttributeGetArgs>? _activeTimeAttributes;

        /// <summary>
        /// Schedule rotation active times per day
        /// </summary>
        public InputList<Inputs.ScheduleRotationActiveDayActiveTimeAttributeGetArgs> ActiveTimeAttributes
        {
            get => _activeTimeAttributes ?? (_activeTimeAttributes = new InputList<Inputs.ScheduleRotationActiveDayActiveTimeAttributeGetArgs>());
            set => _activeTimeAttributes = value;
        }

        /// <summary>
        /// Schedule rotation day name for which active times to be created. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        /// </summary>
        [Input("dayName")]
        public Input<string>? DayName { get; set; }

        [Input("scheduleRotationId")]
        public Input<string>? ScheduleRotationId { get; set; }

        public ScheduleRotationActiveDayState()
        {
        }
        public static new ScheduleRotationActiveDayState Empty => new ScheduleRotationActiveDayState();
    }
}
