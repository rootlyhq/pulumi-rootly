// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    /// <summary>
    /// ## Example Usage
    /// </summary>
    [RootlyResourceType("rootly:index/workflowPulse:WorkflowPulse")]
    public partial class WorkflowPulse : global::Pulumi.CustomResource
    {
        [Output("causeIds")]
        public Output<ImmutableArray<string>> CauseIds { get; private set; } = null!;

        /// <summary>
        /// Workflow command
        /// </summary>
        [Output("command")]
        public Output<string> Command { get; private set; } = null!;

        /// <summary>
        /// This will notify you back when the workflow is starting. Value must be one of true or false
        /// </summary>
        [Output("commandFeedbackEnabled")]
        public Output<bool> CommandFeedbackEnabled { get; private set; } = null!;

        /// <summary>
        /// The description of the workflow
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("environmentIds")]
        public Output<ImmutableArray<string>> EnvironmentIds { get; private set; } = null!;

        [Output("functionalityIds")]
        public Output<ImmutableArray<string>> FunctionalityIds { get; private set; } = null!;

        [Output("groupIds")]
        public Output<ImmutableArray<string>> GroupIds { get; private set; } = null!;

        [Output("incidentRoleIds")]
        public Output<ImmutableArray<string>> IncidentRoleIds { get; private set; } = null!;

        [Output("incidentTypeIds")]
        public Output<ImmutableArray<string>> IncidentTypeIds { get; private set; } = null!;

        /// <summary>
        /// Restricts workflow edits to admins when turned on. Only admins can set this field.. Value must be one of true or false
        /// </summary>
        [Output("locked")]
        public Output<bool> Locked { get; private set; } = null!;

        /// <summary>
        /// The title of the workflow
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The order which the workflow should run with other workflows.
        /// </summary>
        [Output("position")]
        public Output<int> Position { get; private set; } = null!;

        /// <summary>
        /// Repeat workflow every duration
        /// </summary>
        [Output("repeatEveryDuration")]
        public Output<string> RepeatEveryDuration { get; private set; } = null!;

        /// <summary>
        /// Repeat on weekdays. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        /// </summary>
        [Output("repeatOns")]
        public Output<ImmutableArray<string>> RepeatOns { get; private set; } = null!;

        [Output("serviceIds")]
        public Output<ImmutableArray<string>> ServiceIds { get; private set; } = null!;

        [Output("severityIds")]
        public Output<ImmutableArray<string>> SeverityIds { get; private set; } = null!;

        /// <summary>
        /// The slug of the workflow
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        [Output("triggerParams")]
        public Output<Outputs.WorkflowPulseTriggerParams> TriggerParams { get; private set; } = null!;

        /// <summary>
        /// Wait this duration before executing
        /// </summary>
        [Output("wait")]
        public Output<string> Wait { get; private set; } = null!;

        /// <summary>
        /// The group this workflow belongs to.
        /// </summary>
        [Output("workflowGroupId")]
        public Output<string> WorkflowGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a WorkflowPulse resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkflowPulse(string name, WorkflowPulseArgs? args = null, CustomResourceOptions? options = null)
            : base("rootly:index/workflowPulse:WorkflowPulse", name, args ?? new WorkflowPulseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkflowPulse(string name, Input<string> id, WorkflowPulseState? state = null, CustomResourceOptions? options = null)
            : base("rootly:index/workflowPulse:WorkflowPulse", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkflowPulse resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkflowPulse Get(string name, Input<string> id, WorkflowPulseState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkflowPulse(name, id, state, options);
        }
    }

    public sealed class WorkflowPulseArgs : global::Pulumi.ResourceArgs
    {
        [Input("causeIds")]
        private InputList<string>? _causeIds;
        public InputList<string> CauseIds
        {
            get => _causeIds ?? (_causeIds = new InputList<string>());
            set => _causeIds = value;
        }

        /// <summary>
        /// Workflow command
        /// </summary>
        [Input("command")]
        public Input<string>? Command { get; set; }

        /// <summary>
        /// This will notify you back when the workflow is starting. Value must be one of true or false
        /// </summary>
        [Input("commandFeedbackEnabled")]
        public Input<bool>? CommandFeedbackEnabled { get; set; }

        /// <summary>
        /// The description of the workflow
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("environmentIds")]
        private InputList<string>? _environmentIds;
        public InputList<string> EnvironmentIds
        {
            get => _environmentIds ?? (_environmentIds = new InputList<string>());
            set => _environmentIds = value;
        }

        [Input("functionalityIds")]
        private InputList<string>? _functionalityIds;
        public InputList<string> FunctionalityIds
        {
            get => _functionalityIds ?? (_functionalityIds = new InputList<string>());
            set => _functionalityIds = value;
        }

        [Input("groupIds")]
        private InputList<string>? _groupIds;
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        [Input("incidentRoleIds")]
        private InputList<string>? _incidentRoleIds;
        public InputList<string> IncidentRoleIds
        {
            get => _incidentRoleIds ?? (_incidentRoleIds = new InputList<string>());
            set => _incidentRoleIds = value;
        }

        [Input("incidentTypeIds")]
        private InputList<string>? _incidentTypeIds;
        public InputList<string> IncidentTypeIds
        {
            get => _incidentTypeIds ?? (_incidentTypeIds = new InputList<string>());
            set => _incidentTypeIds = value;
        }

        /// <summary>
        /// Restricts workflow edits to admins when turned on. Only admins can set this field.. Value must be one of true or false
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        /// <summary>
        /// The title of the workflow
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The order which the workflow should run with other workflows.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// Repeat workflow every duration
        /// </summary>
        [Input("repeatEveryDuration")]
        public Input<string>? RepeatEveryDuration { get; set; }

        [Input("repeatOns")]
        private InputList<string>? _repeatOns;

        /// <summary>
        /// Repeat on weekdays. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        /// </summary>
        public InputList<string> RepeatOns
        {
            get => _repeatOns ?? (_repeatOns = new InputList<string>());
            set => _repeatOns = value;
        }

        [Input("serviceIds")]
        private InputList<string>? _serviceIds;
        public InputList<string> ServiceIds
        {
            get => _serviceIds ?? (_serviceIds = new InputList<string>());
            set => _serviceIds = value;
        }

        [Input("severityIds")]
        private InputList<string>? _severityIds;
        public InputList<string> SeverityIds
        {
            get => _severityIds ?? (_severityIds = new InputList<string>());
            set => _severityIds = value;
        }

        /// <summary>
        /// The slug of the workflow
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("triggerParams")]
        public Input<Inputs.WorkflowPulseTriggerParamsArgs>? TriggerParams { get; set; }

        /// <summary>
        /// Wait this duration before executing
        /// </summary>
        [Input("wait")]
        public Input<string>? Wait { get; set; }

        /// <summary>
        /// The group this workflow belongs to.
        /// </summary>
        [Input("workflowGroupId")]
        public Input<string>? WorkflowGroupId { get; set; }

        public WorkflowPulseArgs()
        {
        }
        public static new WorkflowPulseArgs Empty => new WorkflowPulseArgs();
    }

    public sealed class WorkflowPulseState : global::Pulumi.ResourceArgs
    {
        [Input("causeIds")]
        private InputList<string>? _causeIds;
        public InputList<string> CauseIds
        {
            get => _causeIds ?? (_causeIds = new InputList<string>());
            set => _causeIds = value;
        }

        /// <summary>
        /// Workflow command
        /// </summary>
        [Input("command")]
        public Input<string>? Command { get; set; }

        /// <summary>
        /// This will notify you back when the workflow is starting. Value must be one of true or false
        /// </summary>
        [Input("commandFeedbackEnabled")]
        public Input<bool>? CommandFeedbackEnabled { get; set; }

        /// <summary>
        /// The description of the workflow
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("environmentIds")]
        private InputList<string>? _environmentIds;
        public InputList<string> EnvironmentIds
        {
            get => _environmentIds ?? (_environmentIds = new InputList<string>());
            set => _environmentIds = value;
        }

        [Input("functionalityIds")]
        private InputList<string>? _functionalityIds;
        public InputList<string> FunctionalityIds
        {
            get => _functionalityIds ?? (_functionalityIds = new InputList<string>());
            set => _functionalityIds = value;
        }

        [Input("groupIds")]
        private InputList<string>? _groupIds;
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        [Input("incidentRoleIds")]
        private InputList<string>? _incidentRoleIds;
        public InputList<string> IncidentRoleIds
        {
            get => _incidentRoleIds ?? (_incidentRoleIds = new InputList<string>());
            set => _incidentRoleIds = value;
        }

        [Input("incidentTypeIds")]
        private InputList<string>? _incidentTypeIds;
        public InputList<string> IncidentTypeIds
        {
            get => _incidentTypeIds ?? (_incidentTypeIds = new InputList<string>());
            set => _incidentTypeIds = value;
        }

        /// <summary>
        /// Restricts workflow edits to admins when turned on. Only admins can set this field.. Value must be one of true or false
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        /// <summary>
        /// The title of the workflow
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The order which the workflow should run with other workflows.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// Repeat workflow every duration
        /// </summary>
        [Input("repeatEveryDuration")]
        public Input<string>? RepeatEveryDuration { get; set; }

        [Input("repeatOns")]
        private InputList<string>? _repeatOns;

        /// <summary>
        /// Repeat on weekdays. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        /// </summary>
        public InputList<string> RepeatOns
        {
            get => _repeatOns ?? (_repeatOns = new InputList<string>());
            set => _repeatOns = value;
        }

        [Input("serviceIds")]
        private InputList<string>? _serviceIds;
        public InputList<string> ServiceIds
        {
            get => _serviceIds ?? (_serviceIds = new InputList<string>());
            set => _serviceIds = value;
        }

        [Input("severityIds")]
        private InputList<string>? _severityIds;
        public InputList<string> SeverityIds
        {
            get => _severityIds ?? (_severityIds = new InputList<string>());
            set => _severityIds = value;
        }

        /// <summary>
        /// The slug of the workflow
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("triggerParams")]
        public Input<Inputs.WorkflowPulseTriggerParamsGetArgs>? TriggerParams { get; set; }

        /// <summary>
        /// Wait this duration before executing
        /// </summary>
        [Input("wait")]
        public Input<string>? Wait { get; set; }

        /// <summary>
        /// The group this workflow belongs to.
        /// </summary>
        [Input("workflowGroupId")]
        public Input<string>? WorkflowGroupId { get; set; }

        public WorkflowPulseState()
        {
        }
        public static new WorkflowPulseState Empty => new WorkflowPulseState();
    }
}
