// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Inputs
{

    public sealed class WorkflowTaskCreateDatadogNotebookTaskParamsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The notebook content
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The notebook kind. Value must be one of `postmortem`, `runbook`, `investigation`, `documentation`, `report`.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        [Input("markPostMortemAsPublished")]
        public Input<bool>? MarkPostMortemAsPublished { get; set; }

        /// <summary>
        /// Retrospective template to use when creating notebook, if desired
        /// </summary>
        [Input("postMortemTemplateId")]
        public Input<string>? PostMortemTemplateId { get; set; }

        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        [Input("template")]
        private InputMap<string>? _template;

        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public InputMap<string> Template
        {
            get => _template ?? (_template = new InputMap<string>());
            set => _template = value;
        }

        /// <summary>
        /// The notebook title
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public WorkflowTaskCreateDatadogNotebookTaskParamsArgs()
        {
        }
        public static new WorkflowTaskCreateDatadogNotebookTaskParamsArgs Empty => new WorkflowTaskCreateDatadogNotebookTaskParamsArgs();
    }
}
