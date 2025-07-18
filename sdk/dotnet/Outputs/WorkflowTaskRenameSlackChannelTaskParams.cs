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
    public sealed class WorkflowTaskRenameSlackChannelTaskParams
    {
        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Channel;
        public readonly string? TaskType;
        public readonly string Title;

        [OutputConstructor]
        private WorkflowTaskRenameSlackChannelTaskParams(
            ImmutableDictionary<string, string> channel,

            string? taskType,

            string title)
        {
            Channel = channel;
            TaskType = taskType;
            Title = title;
        }
    }
}
