// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Outputs
{

    [OutputType]
    public sealed class WorkflowTaskUpdatePagerdutyIncidentTaskParams
    {
        /// <summary>
        /// Escalation level of policy attached to incident
        /// </summary>
        public readonly int? EscalationLevel;
        /// <summary>
        /// Pagerduty incident id
        /// </summary>
        public readonly string PagerdutyIncidentId;
        /// <summary>
        /// PagerDuty incident priority, selecting auto will let Rootly auto map our incident severity
        /// </summary>
        public readonly string? Priority;
        /// <summary>
        /// A message outlining the incident's resolution in PagerDuty
        /// </summary>
        public readonly string? Resolution;
        /// <summary>
        /// Value must be one of `resolved`, `acknowledged`, `auto`.
        /// </summary>
        public readonly string? Status;
        public readonly string? TaskType;
        /// <summary>
        /// Title to update to
        /// </summary>
        public readonly string? Title;
        /// <summary>
        /// PagerDuty incident urgency, selecting auto will let Rootly auto map our incident severity. Value must be one of `high`, `low`, `auto`.
        /// </summary>
        public readonly string? Urgency;

        [OutputConstructor]
        private WorkflowTaskUpdatePagerdutyIncidentTaskParams(
            int? escalationLevel,

            string pagerdutyIncidentId,

            string? priority,

            string? resolution,

            string? status,

            string? taskType,

            string? title,

            string? urgency)
        {
            EscalationLevel = escalationLevel;
            PagerdutyIncidentId = pagerdutyIncidentId;
            Priority = priority;
            Resolution = resolution;
            Status = status;
            TaskType = taskType;
            Title = title;
            Urgency = urgency;
        }
    }
}
