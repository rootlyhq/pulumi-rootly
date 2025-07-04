// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Inputs
{

    public sealed class WorkflowTaskAttachDatadogDashboardsTaskParamsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dashboards", required: true)]
        private InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsDashboardGetArgs>? _dashboards;
        public InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsDashboardGetArgs> Dashboards
        {
            get => _dashboards ?? (_dashboards = new InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsDashboardGetArgs>());
            set => _dashboards = value;
        }

        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        [Input("postToIncidentTimeline")]
        public Input<bool>? PostToIncidentTimeline { get; set; }

        [Input("postToSlackChannels")]
        private InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsPostToSlackChannelGetArgs>? _postToSlackChannels;
        public InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsPostToSlackChannelGetArgs> PostToSlackChannels
        {
            get => _postToSlackChannels ?? (_postToSlackChannels = new InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsPostToSlackChannelGetArgs>());
            set => _postToSlackChannels = value;
        }

        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        public WorkflowTaskAttachDatadogDashboardsTaskParamsGetArgs()
        {
        }
        public static new WorkflowTaskAttachDatadogDashboardsTaskParamsGetArgs Empty => new WorkflowTaskAttachDatadogDashboardsTaskParamsGetArgs();
    }
}
