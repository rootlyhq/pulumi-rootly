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
    public sealed class WorkflowTaskCreateGithubIssueTaskParams
    {
        /// <summary>
        /// The issue body
        /// </summary>
        public readonly string? Body;
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
        private WorkflowTaskCreateGithubIssueTaskParams(
            string? body,

            ImmutableDictionary<string, string> repository,

            string? taskType,

            string title)
        {
            Body = body;
            Repository = repository;
            TaskType = taskType;
            Title = title;
        }
    }
}
