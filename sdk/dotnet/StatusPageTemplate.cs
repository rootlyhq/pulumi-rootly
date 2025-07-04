// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    [RootlyResourceType("rootly:index/statusPageTemplate:StatusPageTemplate")]
    public partial class StatusPageTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the event the template will populate
        /// </summary>
        [Output("body")]
        public Output<string> Body { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The kind of the status page template. Value must be one of `normal`, `scheduled`.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Position of the workflow task
        /// </summary>
        [Output("position")]
        public Output<int> Position { get; private set; } = null!;

        /// <summary>
        /// Controls if incident subscribers should be notified. Value must be one of true or false
        /// </summary>
        [Output("shouldNotifySubscribers")]
        public Output<bool> ShouldNotifySubscribers { get; private set; } = null!;

        [Output("statusPageId")]
        public Output<string> StatusPageId { get; private set; } = null!;

        /// <summary>
        /// Title of the template
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// Status of the event the template will populate
        /// </summary>
        [Output("updateStatus")]
        public Output<string> UpdateStatus { get; private set; } = null!;


        /// <summary>
        /// Create a StatusPageTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StatusPageTemplate(string name, StatusPageTemplateArgs args, CustomResourceOptions? options = null)
            : base("rootly:index/statusPageTemplate:StatusPageTemplate", name, args ?? new StatusPageTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StatusPageTemplate(string name, Input<string> id, StatusPageTemplateState? state = null, CustomResourceOptions? options = null)
            : base("rootly:index/statusPageTemplate:StatusPageTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StatusPageTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StatusPageTemplate Get(string name, Input<string> id, StatusPageTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new StatusPageTemplate(name, id, state, options);
        }
    }

    public sealed class StatusPageTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the event the template will populate
        /// </summary>
        [Input("body", required: true)]
        public Input<string> Body { get; set; } = null!;

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The kind of the status page template. Value must be one of `normal`, `scheduled`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Position of the workflow task
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// Controls if incident subscribers should be notified. Value must be one of true or false
        /// </summary>
        [Input("shouldNotifySubscribers")]
        public Input<bool>? ShouldNotifySubscribers { get; set; }

        [Input("statusPageId")]
        public Input<string>? StatusPageId { get; set; }

        /// <summary>
        /// Title of the template
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// Status of the event the template will populate
        /// </summary>
        [Input("updateStatus")]
        public Input<string>? UpdateStatus { get; set; }

        public StatusPageTemplateArgs()
        {
        }
        public static new StatusPageTemplateArgs Empty => new StatusPageTemplateArgs();
    }

    public sealed class StatusPageTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the event the template will populate
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The kind of the status page template. Value must be one of `normal`, `scheduled`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Position of the workflow task
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// Controls if incident subscribers should be notified. Value must be one of true or false
        /// </summary>
        [Input("shouldNotifySubscribers")]
        public Input<bool>? ShouldNotifySubscribers { get; set; }

        [Input("statusPageId")]
        public Input<string>? StatusPageId { get; set; }

        /// <summary>
        /// Title of the template
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Status of the event the template will populate
        /// </summary>
        [Input("updateStatus")]
        public Input<string>? UpdateStatus { get; set; }

        public StatusPageTemplateState()
        {
        }
        public static new StatusPageTemplateState Empty => new StatusPageTemplateState();
    }
}
