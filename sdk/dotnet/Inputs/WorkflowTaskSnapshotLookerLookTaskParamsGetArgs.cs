// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Inputs
{

    public sealed class WorkflowTaskSnapshotLookerLookTaskParamsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dashboards", required: true)]
        private InputList<Inputs.WorkflowTaskSnapshotLookerLookTaskParamsDashboardGetArgs>? _dashboards;
        public InputList<Inputs.WorkflowTaskSnapshotLookerLookTaskParamsDashboardGetArgs> Dashboards
        {
            get => _dashboards ?? (_dashboards = new InputList<Inputs.WorkflowTaskSnapshotLookerLookTaskParamsDashboardGetArgs>());
            set => _dashboards = value;
        }

        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        [Input("postToIncidentTimeline")]
        public Input<bool>? PostToIncidentTimeline { get; set; }

        [Input("postToSlackChannels")]
        private InputList<Inputs.WorkflowTaskSnapshotLookerLookTaskParamsPostToSlackChannelGetArgs>? _postToSlackChannels;
        public InputList<Inputs.WorkflowTaskSnapshotLookerLookTaskParamsPostToSlackChannelGetArgs> PostToSlackChannels
        {
            get => _postToSlackChannels ?? (_postToSlackChannels = new InputList<Inputs.WorkflowTaskSnapshotLookerLookTaskParamsPostToSlackChannelGetArgs>());
            set => _postToSlackChannels = value;
        }

        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        public WorkflowTaskSnapshotLookerLookTaskParamsGetArgs()
        {
        }
        public static new WorkflowTaskSnapshotLookerLookTaskParamsGetArgs Empty => new WorkflowTaskSnapshotLookerLookTaskParamsGetArgs();
    }
}
