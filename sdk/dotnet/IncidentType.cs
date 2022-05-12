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
    /// Manages incident types (e.g Cloud, Customer Facing, Security, Training).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Rootly = Pulumi.Rootly;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Rootly.IncidentType("foo", new Rootly.IncidentTypeArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import rootly:index/incidentType:IncidentType foo 11111111-2222-3333-4444-555555555555
    /// ```
    /// </summary>
    [RootlyResourceType("rootly:index/incidentType:IncidentType")]
    public partial class IncidentType : Pulumi.CustomResource
    {
        /// <summary>
        /// The cikir of the incident type
        /// </summary>
        [Output("color")]
        public Output<string?> Color { get; private set; } = null!;

        /// <summary>
        /// The description of the incident type
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the incident type
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a IncidentType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IncidentType(string name, IncidentTypeArgs? args = null, CustomResourceOptions? options = null)
            : base("rootly:index/incidentType:IncidentType", name, args ?? new IncidentTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IncidentType(string name, Input<string> id, IncidentTypeState? state = null, CustomResourceOptions? options = null)
            : base("rootly:index/incidentType:IncidentType", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/rootlyhq/pulumi-rootly/releases/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IncidentType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IncidentType Get(string name, Input<string> id, IncidentTypeState? state = null, CustomResourceOptions? options = null)
        {
            return new IncidentType(name, id, state, options);
        }
    }

    public sealed class IncidentTypeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cikir of the incident type
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// The description of the incident type
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the incident type
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public IncidentTypeArgs()
        {
        }
    }

    public sealed class IncidentTypeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cikir of the incident type
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// The description of the incident type
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the incident type
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public IncidentTypeState()
        {
        }
    }
}