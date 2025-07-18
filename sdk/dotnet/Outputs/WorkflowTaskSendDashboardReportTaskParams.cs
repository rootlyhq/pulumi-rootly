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
    public sealed class WorkflowTaskSendDashboardReportTaskParams
    {
        /// <summary>
        /// The email body
        /// </summary>
        public readonly string Body;
        public readonly ImmutableArray<string> DashboardIds;
        /// <summary>
        /// The from email address. Need to use SMTP integration if different than rootly.com
        /// </summary>
        public readonly string? From;
        /// <summary>
        /// The preheader
        /// </summary>
        public readonly string? Preheader;
        /// <summary>
        /// The subject
        /// </summary>
        public readonly string Subject;
        public readonly string? TaskType;
        public readonly ImmutableArray<string> Tos;

        [OutputConstructor]
        private WorkflowTaskSendDashboardReportTaskParams(
            string body,

            ImmutableArray<string> dashboardIds,

            string? from,

            string? preheader,

            string subject,

            string? taskType,

            ImmutableArray<string> tos)
        {
            Body = body;
            DashboardIds = dashboardIds;
            From = from;
            Preheader = preheader;
            Subject = subject;
            TaskType = taskType;
            Tos = tos;
        }
    }
}
