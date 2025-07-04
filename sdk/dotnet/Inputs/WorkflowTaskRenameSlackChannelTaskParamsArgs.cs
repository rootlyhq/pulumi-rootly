// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Inputs
{

    public sealed class WorkflowTaskRenameSlackChannelTaskParamsArgs : global::Pulumi.ResourceArgs
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

        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public WorkflowTaskRenameSlackChannelTaskParamsArgs()
        {
        }
        public static new WorkflowTaskRenameSlackChannelTaskParamsArgs Empty => new WorkflowTaskRenameSlackChannelTaskParamsArgs();
    }
}
