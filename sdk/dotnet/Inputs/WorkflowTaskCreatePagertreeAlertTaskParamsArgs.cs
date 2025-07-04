// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Inputs
{

    public sealed class WorkflowTaskCreatePagertreeAlertTaskParamsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of alert as text
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Setting to true makes an alert a Pagertree incident. Value must be one of true or false
        /// </summary>
        [Input("incident")]
        public Input<bool>? Incident { get; set; }

        /// <summary>
        /// Value must be one of `auto`, `SEV-1`, `SEV-2`, `SEV-3`, `SEV-4`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        [Input("teams")]
        private InputList<Inputs.WorkflowTaskCreatePagertreeAlertTaskParamsTeamArgs>? _teams;
        public InputList<Inputs.WorkflowTaskCreatePagertreeAlertTaskParamsTeamArgs> Teams
        {
            get => _teams ?? (_teams = new InputList<Inputs.WorkflowTaskCreatePagertreeAlertTaskParamsTeamArgs>());
            set => _teams = value;
        }

        /// <summary>
        /// Title of alert as text
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Value must be one of `auto`, `critical`, `high`, `medium`, `low`.
        /// </summary>
        [Input("urgency")]
        public Input<string>? Urgency { get; set; }

        [Input("users")]
        private InputList<Inputs.WorkflowTaskCreatePagertreeAlertTaskParamsUserArgs>? _users;
        public InputList<Inputs.WorkflowTaskCreatePagertreeAlertTaskParamsUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.WorkflowTaskCreatePagertreeAlertTaskParamsUserArgs>());
            set => _users = value;
        }

        public WorkflowTaskCreatePagertreeAlertTaskParamsArgs()
        {
        }
        public static new WorkflowTaskCreatePagertreeAlertTaskParamsArgs Empty => new WorkflowTaskCreatePagertreeAlertTaskParamsArgs();
    }
}
