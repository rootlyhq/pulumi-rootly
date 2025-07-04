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
    public sealed class WorkflowTaskCreateDropboxPaperPageTaskParams
    {
        /// <summary>
        /// The page content
        /// </summary>
        public readonly string? Content;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? MarkPostMortemAsPublished;
        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Namespace;
        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ParentFolder;
        /// <summary>
        /// Retrospective template to use when creating page task, if desired
        /// </summary>
        public readonly string? PostMortemTemplateId;
        public readonly string? TaskType;
        /// <summary>
        /// The page task title
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private WorkflowTaskCreateDropboxPaperPageTaskParams(
            string? content,

            bool? markPostMortemAsPublished,

            ImmutableDictionary<string, string>? @namespace,

            ImmutableDictionary<string, string>? parentFolder,

            string? postMortemTemplateId,

            string? taskType,

            string title)
        {
            Content = content;
            MarkPostMortemAsPublished = markPostMortemAsPublished;
            Namespace = @namespace;
            ParentFolder = parentFolder;
            PostMortemTemplateId = postMortemTemplateId;
            TaskType = taskType;
            Title = title;
        }
    }
}
