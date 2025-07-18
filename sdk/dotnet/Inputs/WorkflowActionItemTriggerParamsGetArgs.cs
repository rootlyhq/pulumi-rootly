// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Inputs
{

    public sealed class WorkflowActionItemTriggerParamsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value must be one off `ALL`, `ANY`, `NONE`.
        /// </summary>
        [Input("incidentActionItemCondition")]
        public Input<string>? IncidentActionItemCondition { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentActionItemConditionGroup")]
        public Input<string>? IncidentActionItemConditionGroup { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentActionItemConditionKind")]
        public Input<string>? IncidentActionItemConditionKind { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentActionItemConditionPriority")]
        public Input<string>? IncidentActionItemConditionPriority { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentActionItemConditionStatus")]
        public Input<string>? IncidentActionItemConditionStatus { get; set; }

        [Input("incidentActionItemGroupIds")]
        private InputList<string>? _incidentActionItemGroupIds;
        public InputList<string> IncidentActionItemGroupIds
        {
            get => _incidentActionItemGroupIds ?? (_incidentActionItemGroupIds = new InputList<string>());
            set => _incidentActionItemGroupIds = value;
        }

        [Input("incidentActionItemKinds")]
        private InputList<string>? _incidentActionItemKinds;

        /// <summary>
        /// Value must be one of `task`, `follow_up`.
        /// </summary>
        public InputList<string> IncidentActionItemKinds
        {
            get => _incidentActionItemKinds ?? (_incidentActionItemKinds = new InputList<string>());
            set => _incidentActionItemKinds = value;
        }

        [Input("incidentActionItemPriorities")]
        private InputList<string>? _incidentActionItemPriorities;

        /// <summary>
        /// Value must be one of `high`, `medium`, `low`.
        /// </summary>
        public InputList<string> IncidentActionItemPriorities
        {
            get => _incidentActionItemPriorities ?? (_incidentActionItemPriorities = new InputList<string>());
            set => _incidentActionItemPriorities = value;
        }

        [Input("incidentActionItemStatuses")]
        private InputList<string>? _incidentActionItemStatuses;

        /// <summary>
        /// Value must be one of `open`, `in_progress`, `cancelled`, `done`.
        /// </summary>
        public InputList<string> IncidentActionItemStatuses
        {
            get => _incidentActionItemStatuses ?? (_incidentActionItemStatuses = new InputList<string>());
            set => _incidentActionItemStatuses = value;
        }

        /// <summary>
        /// Value must be one off `ALL`, `ANY`, `NONE`.
        /// </summary>
        [Input("incidentCondition")]
        public Input<string>? IncidentCondition { get; set; }

        [Input("incidentConditionAcknowledgedAt")]
        public Input<string>? IncidentConditionAcknowledgedAt { get; set; }

        [Input("incidentConditionDetectedAt")]
        public Input<string>? IncidentConditionDetectedAt { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentConditionEnvironment")]
        public Input<string>? IncidentConditionEnvironment { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentConditionFunctionality")]
        public Input<string>? IncidentConditionFunctionality { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentConditionGroup")]
        public Input<string>? IncidentConditionGroup { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentConditionIncidentRoles")]
        public Input<string>? IncidentConditionIncidentRoles { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentConditionIncidentType")]
        public Input<string>? IncidentConditionIncidentType { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentConditionKind")]
        public Input<string>? IncidentConditionKind { get; set; }

        [Input("incidentConditionMitigatedAt")]
        public Input<string>? IncidentConditionMitigatedAt { get; set; }

        [Input("incidentConditionResolvedAt")]
        public Input<string>? IncidentConditionResolvedAt { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentConditionService")]
        public Input<string>? IncidentConditionService { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentConditionSeverity")]
        public Input<string>? IncidentConditionSeverity { get; set; }

        [Input("incidentConditionStartedAt")]
        public Input<string>? IncidentConditionStartedAt { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentConditionStatus")]
        public Input<string>? IncidentConditionStatus { get; set; }

        [Input("incidentConditionSummary")]
        public Input<string>? IncidentConditionSummary { get; set; }

        /// <summary>
        /// Value must be one off `IS`, `ANY`, `CONTAINS`, `CONTAINS_ALL`, `CONTAINS_NONE`, `NONE`, `SET`, `UNSET`.
        /// </summary>
        [Input("incidentConditionVisibility")]
        public Input<string>? IncidentConditionVisibility { get; set; }

        [Input("incidentConditionalInactivity")]
        public Input<string>? IncidentConditionalInactivity { get; set; }

        /// <summary>
        /// ex. 10 min, 1h, 3 days, 2 weeks
        /// </summary>
        [Input("incidentInactivityDuration")]
        public Input<string>? IncidentInactivityDuration { get; set; }

        [Input("incidentKinds")]
        private InputList<string>? _incidentKinds;

        /// <summary>
        /// Value must be one of `test`, `test_sub`, `example`, `example_sub`, `normal`, `normal_sub`, `backfilled`, `scheduled`.
        /// </summary>
        public InputList<string> IncidentKinds
        {
            get => _incidentKinds ?? (_incidentKinds = new InputList<string>());
            set => _incidentKinds = value;
        }

        [Input("incidentStatuses")]
        private InputList<string>? _incidentStatuses;

        /// <summary>
        /// Value must be one of `in_triage`, `started`, `detected`, `acknowledged`, `mitigated`, `resolved`, `cancelled`, `scheduled`, `in_progress`, `completed`.
        /// </summary>
        public InputList<string> IncidentStatuses
        {
            get => _incidentStatuses ?? (_incidentStatuses = new InputList<string>());
            set => _incidentStatuses = value;
        }

        [Input("incidentVisibilities")]
        private InputList<string>? _incidentVisibilities;
        public InputList<string> IncidentVisibilities
        {
            get => _incidentVisibilities ?? (_incidentVisibilities = new InputList<string>());
            set => _incidentVisibilities = value;
        }

        /// <summary>
        /// Value must be one off `action_item`.
        /// </summary>
        [Input("triggerType")]
        public Input<string>? TriggerType { get; set; }

        [Input("triggers")]
        private InputList<string>? _triggers;

        /// <summary>
        /// Actions that trigger the workflow. One of custom*fields.\n\n.updated, incident*updated, action*item*created, action*item*updated, assigned*user*updated, summary*updated, description*updated, status*updated, priority*updated, due*date*updated, teams*updated, slack*command
        /// </summary>
        public InputList<string> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<string>());
            set => _triggers = value;
        }

        public WorkflowActionItemTriggerParamsGetArgs()
        {
        }
        public static new WorkflowActionItemTriggerParamsGetArgs Empty => new WorkflowActionItemTriggerParamsGetArgs();
    }
}
