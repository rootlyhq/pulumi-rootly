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
    public sealed class WorkflowTaskCreateTrelloCardTaskParams
    {
        /// <summary>
        /// Map must contain two fields, `id` and `name`. The archivation id and display name
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Archivation;
        /// <summary>
        /// Map must contain two fields, `id` and `name`. The board id and display name
        /// </summary>
        public readonly ImmutableDictionary<string, string> Board;
        /// <summary>
        /// The card description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The due date
        /// </summary>
        public readonly string? DueDate;
        public readonly ImmutableArray<Outputs.WorkflowTaskCreateTrelloCardTaskParamsLabel> Labels;
        /// <summary>
        /// Map must contain two fields, `id` and `name`. The list id and display name
        /// </summary>
        public readonly ImmutableDictionary<string, string> List;
        public readonly string? TaskType;
        /// <summary>
        /// The card title
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private WorkflowTaskCreateTrelloCardTaskParams(
            ImmutableDictionary<string, string>? archivation,

            ImmutableDictionary<string, string> board,

            string? description,

            string? dueDate,

            ImmutableArray<Outputs.WorkflowTaskCreateTrelloCardTaskParamsLabel> labels,

            ImmutableDictionary<string, string> list,

            string? taskType,

            string title)
        {
            Archivation = archivation;
            Board = board;
            Description = description;
            DueDate = dueDate;
            Labels = labels;
            List = list;
            TaskType = taskType;
            Title = title;
        }
    }
}
