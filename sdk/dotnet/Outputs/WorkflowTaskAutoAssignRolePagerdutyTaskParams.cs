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
    public sealed class WorkflowTaskAutoAssignRolePagerdutyTaskParams
    {
        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? EscalationPolicy;
        /// <summary>
        /// The role id
        /// </summary>
        public readonly string IncidentRoleId;
        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Schedule;
        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Service;
        public readonly string? TaskType;

        [OutputConstructor]
        private WorkflowTaskAutoAssignRolePagerdutyTaskParams(
            ImmutableDictionary<string, string>? escalationPolicy,

            string incidentRoleId,

            ImmutableDictionary<string, string>? schedule,

            ImmutableDictionary<string, string>? service,

            string? taskType)
        {
            EscalationPolicy = escalationPolicy;
            IncidentRoleId = incidentRoleId;
            Schedule = schedule;
            Service = service;
            TaskType = taskType;
        }
    }
}
