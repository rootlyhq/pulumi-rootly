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
    public sealed class WorkflowTaskInviteToSlackChannelVictorOpsTaskParams
    {
        public readonly ImmutableArray<Outputs.WorkflowTaskInviteToSlackChannelVictorOpsTaskParamsChannel> Channels;
        public readonly string? TaskType;
        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Team;

        [OutputConstructor]
        private WorkflowTaskInviteToSlackChannelVictorOpsTaskParams(
            ImmutableArray<Outputs.WorkflowTaskInviteToSlackChannelVictorOpsTaskParamsChannel> channels,

            string? taskType,

            ImmutableDictionary<string, string> team)
        {
            Channels = channels;
            TaskType = taskType;
            Team = team;
        }
    }
}
