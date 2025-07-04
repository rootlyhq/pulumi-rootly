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
    public sealed class WorkflowTaskInviteToSlackChannelTaskParams
    {
        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Channel;
        public readonly ImmutableArray<Outputs.WorkflowTaskInviteToSlackChannelTaskParamsSlackUserGroup> SlackUserGroups;
        public readonly ImmutableArray<Outputs.WorkflowTaskInviteToSlackChannelTaskParamsSlackUser> SlackUsers;
        public readonly string? TaskType;

        [OutputConstructor]
        private WorkflowTaskInviteToSlackChannelTaskParams(
            ImmutableDictionary<string, string> channel,

            ImmutableArray<Outputs.WorkflowTaskInviteToSlackChannelTaskParamsSlackUserGroup> slackUserGroups,

            ImmutableArray<Outputs.WorkflowTaskInviteToSlackChannelTaskParamsSlackUser> slackUsers,

            string? taskType)
        {
            Channel = channel;
            SlackUserGroups = slackUserGroups;
            SlackUsers = slackUsers;
            TaskType = taskType;
        }
    }
}
