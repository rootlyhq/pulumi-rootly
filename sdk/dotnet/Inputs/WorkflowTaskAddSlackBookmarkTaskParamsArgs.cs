// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Inputs
{

    public sealed class WorkflowTaskAddSlackBookmarkTaskParamsArgs : global::Pulumi.ResourceArgs
    {
        [Input("channel", required: true)]
        private InputMap<string>? _channel;

        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public InputMap<string> Channel
        {
            get => _channel ?? (_channel = new InputMap<string>());
            set => _channel = value;
        }

        /// <summary>
        /// The bookmark emoji
        /// </summary>
        [Input("emoji")]
        public Input<string>? Emoji { get; set; }

        /// <summary>
        /// The bookmark link. Required if not a playbook bookmark
        /// </summary>
        [Input("link")]
        public Input<string>? Link { get; set; }

        /// <summary>
        /// The playbook id if bookmark is of an incident playbook
        /// </summary>
        [Input("playbookId")]
        public Input<string>? PlaybookId { get; set; }

        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        /// <summary>
        /// The bookmark title. Required if not a playbook bookmark
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public WorkflowTaskAddSlackBookmarkTaskParamsArgs()
        {
        }
        public static new WorkflowTaskAddSlackBookmarkTaskParamsArgs Empty => new WorkflowTaskAddSlackBookmarkTaskParamsArgs();
    }
}
