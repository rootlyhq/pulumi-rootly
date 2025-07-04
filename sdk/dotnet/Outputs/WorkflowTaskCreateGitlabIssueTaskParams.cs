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
    public sealed class WorkflowTaskCreateGitlabIssueTaskParams
    {
        /// <summary>
        /// The issue description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The due date
        /// </summary>
        public readonly string? DueDate;
        /// <summary>
        /// The issue type. Value must be one of `issue`, `incident`, `test_case`, `task`.
        /// </summary>
        public readonly string? IssueType;
        /// <summary>
        /// The issue labels
        /// </summary>
        public readonly string? Labels;
        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Repository;
        public readonly string? TaskType;
        /// <summary>
        /// The issue title
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private WorkflowTaskCreateGitlabIssueTaskParams(
            string? description,

            string? dueDate,

            string? issueType,

            string? labels,

            ImmutableDictionary<string, string> repository,

            string? taskType,

            string title)
        {
            Description = description;
            DueDate = dueDate;
            IssueType = issueType;
            Labels = labels;
            Repository = repository;
            TaskType = taskType;
            Title = title;
        }
    }
}
