// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Inputs
{

    public sealed class WorkflowTaskAttachDatadogDashboardsTaskParamsArgs : global::Pulumi.ResourceArgs
    {
        [Input("dashboards", required: true)]
        private InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsDashboardArgs>? _dashboards;
        public InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsDashboardArgs> Dashboards
        {
            get => _dashboards ?? (_dashboards = new InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsDashboardArgs>());
            set => _dashboards = value;
        }

        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        [Input("postToIncidentTimeline")]
        public Input<bool>? PostToIncidentTimeline { get; set; }

        [Input("postToSlackChannels")]
        private InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsPostToSlackChannelArgs>? _postToSlackChannels;
        public InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsPostToSlackChannelArgs> PostToSlackChannels
        {
            get => _postToSlackChannels ?? (_postToSlackChannels = new InputList<Inputs.WorkflowTaskAttachDatadogDashboardsTaskParamsPostToSlackChannelArgs>());
            set => _postToSlackChannels = value;
        }

        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        public WorkflowTaskAttachDatadogDashboardsTaskParamsArgs()
        {
        }
        public static new WorkflowTaskAttachDatadogDashboardsTaskParamsArgs Empty => new WorkflowTaskAttachDatadogDashboardsTaskParamsArgs();
    }
}
