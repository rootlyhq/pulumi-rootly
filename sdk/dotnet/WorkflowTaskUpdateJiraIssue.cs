// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    /// <summary>
    /// Manages workflow update_jira_issue task.
    /// </summary>
    [RootlyResourceType("rootly:index/workflowTaskUpdateJiraIssue:WorkflowTaskUpdateJiraIssue")]
    public partial class WorkflowTaskUpdateJiraIssue : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable this workflow task
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Name of the workflow task
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The position of the workflow task (1 being top of list)
        /// </summary>
        [Output("position")]
        public Output<int> Position { get; private set; } = null!;

        /// <summary>
        /// Skip workflow task if any failures
        /// </summary>
        [Output("skipOnFailure")]
        public Output<bool?> SkipOnFailure { get; private set; } = null!;

        /// <summary>
        /// The parameters for this workflow task.
        /// </summary>
        [Output("taskParams")]
        public Output<Outputs.WorkflowTaskUpdateJiraIssueTaskParams> TaskParams { get; private set; } = null!;

        /// <summary>
        /// The ID of the parent workflow
        /// </summary>
        [Output("workflowId")]
        public Output<string> WorkflowId { get; private set; } = null!;


        /// <summary>
        /// Create a WorkflowTaskUpdateJiraIssue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkflowTaskUpdateJiraIssue(string name, WorkflowTaskUpdateJiraIssueArgs args, CustomResourceOptions? options = null)
            : base("rootly:index/workflowTaskUpdateJiraIssue:WorkflowTaskUpdateJiraIssue", name, args ?? new WorkflowTaskUpdateJiraIssueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkflowTaskUpdateJiraIssue(string name, Input<string> id, WorkflowTaskUpdateJiraIssueState? state = null, CustomResourceOptions? options = null)
            : base("rootly:index/workflowTaskUpdateJiraIssue:WorkflowTaskUpdateJiraIssue", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/rootlyhq/pulumi-rootly/releases/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkflowTaskUpdateJiraIssue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkflowTaskUpdateJiraIssue Get(string name, Input<string> id, WorkflowTaskUpdateJiraIssueState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkflowTaskUpdateJiraIssue(name, id, state, options);
        }
    }

    public sealed class WorkflowTaskUpdateJiraIssueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable this workflow task
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Name of the workflow task
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The position of the workflow task (1 being top of list)
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// Skip workflow task if any failures
        /// </summary>
        [Input("skipOnFailure")]
        public Input<bool>? SkipOnFailure { get; set; }

        /// <summary>
        /// The parameters for this workflow task.
        /// </summary>
        [Input("taskParams", required: true)]
        public Input<Inputs.WorkflowTaskUpdateJiraIssueTaskParamsArgs> TaskParams { get; set; } = null!;

        /// <summary>
        /// The ID of the parent workflow
        /// </summary>
        [Input("workflowId", required: true)]
        public Input<string> WorkflowId { get; set; } = null!;

        public WorkflowTaskUpdateJiraIssueArgs()
        {
        }
        public static new WorkflowTaskUpdateJiraIssueArgs Empty => new WorkflowTaskUpdateJiraIssueArgs();
    }

    public sealed class WorkflowTaskUpdateJiraIssueState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable this workflow task
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Name of the workflow task
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The position of the workflow task (1 being top of list)
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// Skip workflow task if any failures
        /// </summary>
        [Input("skipOnFailure")]
        public Input<bool>? SkipOnFailure { get; set; }

        /// <summary>
        /// The parameters for this workflow task.
        /// </summary>
        [Input("taskParams")]
        public Input<Inputs.WorkflowTaskUpdateJiraIssueTaskParamsGetArgs>? TaskParams { get; set; }

        /// <summary>
        /// The ID of the parent workflow
        /// </summary>
        [Input("workflowId")]
        public Input<string>? WorkflowId { get; set; }

        public WorkflowTaskUpdateJiraIssueState()
        {
        }
        public static new WorkflowTaskUpdateJiraIssueState Empty => new WorkflowTaskUpdateJiraIssueState();
    }
}
