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
    public sealed class WorkflowTaskSendSlackMessageTaskParams
    {
        /// <summary>
        /// Value must be one of `update_summary`, `update_status`, `archive_channel`, `manage_incident_roles`, `update_incident`, `all_commands`, `leave_feedback`, `manage_form_fields`, `manage_action_items`, `view_tasks`, `add_pagerduty_responders`, `add_opsgenie_responders`, `add_victor_ops_responders`, `snooze_reminder`, `pause_reminder`, `restart_reminder`, `update_status_page`, `cancel_incident`.
        /// </summary>
        public readonly ImmutableArray<string> Actionables;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? BroadcastThreadReplyToChannel;
        public readonly ImmutableArray<Outputs.WorkflowTaskSendSlackMessageTaskParamsChannel> Channels;
        /// <summary>
        /// A hex color
        /// </summary>
        public readonly string? Color;
        /// <summary>
        /// Map must contain two fields, `id` and `name`. A hash where [id] is the task id of the parent task that sent a message, and [name] is the name of the parent task
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ParentMessageThreadTask;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? PinToChannel;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? SendAsEphemeral;
        /// <summary>
        /// When set to true, if the parent for this threaded message cannot be found the message will be skipped.. Value must be one of true or false
        /// </summary>
        public readonly bool? SendOnlyAsThreadedMessage;
        public readonly ImmutableArray<Outputs.WorkflowTaskSendSlackMessageTaskParamsSlackUserGroup> SlackUserGroups;
        public readonly ImmutableArray<Outputs.WorkflowTaskSendSlackMessageTaskParamsSlackUser> SlackUsers;
        public readonly string? TaskType;
        /// <summary>
        /// The message text
        /// </summary>
        public readonly string Text;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? UpdateParentMessage;

        [OutputConstructor]
        private WorkflowTaskSendSlackMessageTaskParams(
            ImmutableArray<string> actionables,

            bool? broadcastThreadReplyToChannel,

            ImmutableArray<Outputs.WorkflowTaskSendSlackMessageTaskParamsChannel> channels,

            string? color,

            ImmutableDictionary<string, string>? parentMessageThreadTask,

            bool? pinToChannel,

            bool? sendAsEphemeral,

            bool? sendOnlyAsThreadedMessage,

            ImmutableArray<Outputs.WorkflowTaskSendSlackMessageTaskParamsSlackUserGroup> slackUserGroups,

            ImmutableArray<Outputs.WorkflowTaskSendSlackMessageTaskParamsSlackUser> slackUsers,

            string? taskType,

            string text,

            bool? updateParentMessage)
        {
            Actionables = actionables;
            BroadcastThreadReplyToChannel = broadcastThreadReplyToChannel;
            Channels = channels;
            Color = color;
            ParentMessageThreadTask = parentMessageThreadTask;
            PinToChannel = pinToChannel;
            SendAsEphemeral = sendAsEphemeral;
            SendOnlyAsThreadedMessage = sendOnlyAsThreadedMessage;
            SlackUserGroups = slackUserGroups;
            SlackUsers = slackUsers;
            TaskType = taskType;
            Text = text;
            UpdateParentMessage = updateParentMessage;
        }
    }
}
