// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Inputs
{

    public sealed class WorkflowTaskAutoAssignRoleRootlyTaskParamsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("escalationPolicyTarget")]
        private InputMap<string>? _escalationPolicyTarget;

        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public InputMap<string> EscalationPolicyTarget
        {
            get => _escalationPolicyTarget ?? (_escalationPolicyTarget = new InputMap<string>());
            set => _escalationPolicyTarget = value;
        }

        [Input("groupTarget")]
        private InputMap<string>? _groupTarget;

        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public InputMap<string> GroupTarget
        {
            get => _groupTarget ?? (_groupTarget = new InputMap<string>());
            set => _groupTarget = value;
        }

        /// <summary>
        /// The role id
        /// </summary>
        [Input("incidentRoleId", required: true)]
        public Input<string> IncidentRoleId { get; set; } = null!;

        [Input("scheduleTarget")]
        private InputMap<string>? _scheduleTarget;

        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public InputMap<string> ScheduleTarget
        {
            get => _scheduleTarget ?? (_scheduleTarget = new InputMap<string>());
            set => _scheduleTarget = value;
        }

        [Input("serviceTarget")]
        private InputMap<string>? _serviceTarget;

        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public InputMap<string> ServiceTarget
        {
            get => _serviceTarget ?? (_serviceTarget = new InputMap<string>());
            set => _serviceTarget = value;
        }

        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        [Input("userTarget")]
        private InputMap<string>? _userTarget;

        /// <summary>
        /// Map must contain two fields, `id` and `name`.
        /// </summary>
        public InputMap<string> UserTarget
        {
            get => _userTarget ?? (_userTarget = new InputMap<string>());
            set => _userTarget = value;
        }

        public WorkflowTaskAutoAssignRoleRootlyTaskParamsGetArgs()
        {
        }
        public static new WorkflowTaskAutoAssignRoleRootlyTaskParamsGetArgs Empty => new WorkflowTaskAutoAssignRoleRootlyTaskParamsGetArgs();
    }
}
