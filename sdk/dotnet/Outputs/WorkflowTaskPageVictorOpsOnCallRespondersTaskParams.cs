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
    public sealed class WorkflowTaskPageVictorOpsOnCallRespondersTaskParams
    {
        public readonly ImmutableArray<Outputs.WorkflowTaskPageVictorOpsOnCallRespondersTaskParamsEscalationPolicy> EscalationPolicies;
        public readonly string? TaskType;
        /// <summary>
        /// Alert title.
        /// </summary>
        public readonly string? Title;
        public readonly ImmutableArray<Outputs.WorkflowTaskPageVictorOpsOnCallRespondersTaskParamsUser> Users;

        [OutputConstructor]
        private WorkflowTaskPageVictorOpsOnCallRespondersTaskParams(
            ImmutableArray<Outputs.WorkflowTaskPageVictorOpsOnCallRespondersTaskParamsEscalationPolicy> escalationPolicies,

            string? taskType,

            string? title,

            ImmutableArray<Outputs.WorkflowTaskPageVictorOpsOnCallRespondersTaskParamsUser> users)
        {
            EscalationPolicies = escalationPolicies;
            TaskType = taskType;
            Title = title;
            Users = users;
        }
    }
}
