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
    public sealed class WorkflowTaskUpdateActionItemTaskParams
    {
        /// <summary>
        /// Map must contain two fields, `id` and `name`.  The user this action item is assigned to
        /// </summary>
        public readonly ImmutableDictionary<string, string>? AssignedToUser;
        /// <summary>
        /// [DEPRECATED] Use assigned*to*user attribute instead. The user id this action item is assigned to
        /// </summary>
        public readonly string? AssignedToUserId;
        /// <summary>
        /// Attribute of the action item to match against. Value must be one of `id`, `jira_issue_id`, `asana_task_id`, `shortcut_task_id`, `linear_issue_id`, `zendesk_ticket_id`, `trello_card_id`, `airtable_record_id`, `shortcut_story_id`, `github_issue_id`, `gitlab_issue_id`, `freshservice_ticket_id`, `freshservice_task_id`, `clickup_task_id`.
        /// </summary>
        public readonly string AttributeToQueryBy;
        /// <summary>
        /// Custom field mappings. Can contain liquid markup and need to be valid JSON
        /// </summary>
        public readonly string? CustomFieldsMapping;
        /// <summary>
        /// The action item description
        /// </summary>
        public readonly string? Description;
        public readonly ImmutableArray<string> GroupIds;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? PostToIncidentTimeline;
        /// <summary>
        /// The action item priority. Value must be one of `high`, `medium`, `low`.
        /// </summary>
        public readonly string? Priority;
        /// <summary>
        /// Value that attribute*to*query_by to uses to match against
        /// </summary>
        public readonly string QueryValue;
        /// <summary>
        /// The action item status. Value must be one of `open`, `in_progress`, `cancelled`, `done`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Brief description of the action item
        /// </summary>
        public readonly string? Summary;
        public readonly string? TaskType;

        [OutputConstructor]
        private WorkflowTaskUpdateActionItemTaskParams(
            ImmutableDictionary<string, string>? assignedToUser,

            string? assignedToUserId,

            string attributeToQueryBy,

            string? customFieldsMapping,

            string? description,

            ImmutableArray<string> groupIds,

            bool? postToIncidentTimeline,

            string? priority,

            string queryValue,

            string? status,

            string? summary,

            string? taskType)
        {
            AssignedToUser = assignedToUser;
            AssignedToUserId = assignedToUserId;
            AttributeToQueryBy = attributeToQueryBy;
            CustomFieldsMapping = customFieldsMapping;
            Description = description;
            GroupIds = groupIds;
            PostToIncidentTimeline = postToIncidentTimeline;
            Priority = priority;
            QueryValue = queryValue;
            Status = status;
            Summary = summary;
            TaskType = taskType;
        }
    }
}
