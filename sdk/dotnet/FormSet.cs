// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly
{
    [RootlyResourceType("rootly:index/formSet:FormSet")]
    public partial class FormSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The forms included in the form set. Add custom forms using the custom form's `slug` field. Or choose a built-in form: `web_new_incident_form`, `web_update_incident_form`, `web_incident_post_mortem_form`, `web_incident_mitigation_form`, `web_incident_resolution_form`, `web_incident_cancellation_form`, `web_scheduled_incident_form`, `web_update_scheduled_incident_form`, `slack_new_incident_form`, `slack_update_incident_form`, `slack_update_incident_status_form`, `slack_incident_mitigation_form`, `slack_incident_resolution_form`, `slack_incident_cancellation_form`, `slack_scheduled_incident_form`, `slack_update_scheduled_incident_form`
        /// </summary>
        [Output("forms")]
        public Output<ImmutableArray<string>> Forms { get; private set; } = null!;

        /// <summary>
        /// Whether the form set is default. Value must be one of true or false
        /// </summary>
        [Output("isDefault")]
        public Output<bool> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The name of the form set
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The slug of the form set
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;


        /// <summary>
        /// Create a FormSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FormSet(string name, FormSetArgs args, CustomResourceOptions? options = null)
            : base("rootly:index/formSet:FormSet", name, args ?? new FormSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FormSet(string name, Input<string> id, FormSetState? state = null, CustomResourceOptions? options = null)
            : base("rootly:index/formSet:FormSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FormSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FormSet Get(string name, Input<string> id, FormSetState? state = null, CustomResourceOptions? options = null)
        {
            return new FormSet(name, id, state, options);
        }
    }

    public sealed class FormSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("forms", required: true)]
        private InputList<string>? _forms;

        /// <summary>
        /// The forms included in the form set. Add custom forms using the custom form's `slug` field. Or choose a built-in form: `web_new_incident_form`, `web_update_incident_form`, `web_incident_post_mortem_form`, `web_incident_mitigation_form`, `web_incident_resolution_form`, `web_incident_cancellation_form`, `web_scheduled_incident_form`, `web_update_scheduled_incident_form`, `slack_new_incident_form`, `slack_update_incident_form`, `slack_update_incident_status_form`, `slack_incident_mitigation_form`, `slack_incident_resolution_form`, `slack_incident_cancellation_form`, `slack_scheduled_incident_form`, `slack_update_scheduled_incident_form`
        /// </summary>
        public InputList<string> Forms
        {
            get => _forms ?? (_forms = new InputList<string>());
            set => _forms = value;
        }

        /// <summary>
        /// Whether the form set is default. Value must be one of true or false
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The name of the form set
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The slug of the form set
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        public FormSetArgs()
        {
        }
        public static new FormSetArgs Empty => new FormSetArgs();
    }

    public sealed class FormSetState : global::Pulumi.ResourceArgs
    {
        [Input("forms")]
        private InputList<string>? _forms;

        /// <summary>
        /// The forms included in the form set. Add custom forms using the custom form's `slug` field. Or choose a built-in form: `web_new_incident_form`, `web_update_incident_form`, `web_incident_post_mortem_form`, `web_incident_mitigation_form`, `web_incident_resolution_form`, `web_incident_cancellation_form`, `web_scheduled_incident_form`, `web_update_scheduled_incident_form`, `slack_new_incident_form`, `slack_update_incident_form`, `slack_update_incident_status_form`, `slack_incident_mitigation_form`, `slack_incident_resolution_form`, `slack_incident_cancellation_form`, `slack_scheduled_incident_form`, `slack_update_scheduled_incident_form`
        /// </summary>
        public InputList<string> Forms
        {
            get => _forms ?? (_forms = new InputList<string>());
            set => _forms = value;
        }

        /// <summary>
        /// Whether the form set is default. Value must be one of true or false
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The name of the form set
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The slug of the form set
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        public FormSetState()
        {
        }
        public static new FormSetState Empty => new FormSetState();
    }
}
