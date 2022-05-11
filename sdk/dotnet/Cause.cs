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
    /// Manages incident causes (e.g Bug, Load, Human Error, 3rd party Outage, Configuration Change).
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
    ///         var foo = new Rootly.Cause("foo", new Rootly.CauseArgs
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
    ///  $ pulumi import rootly:index/cause:Cause foo 11111111-2222-3333-4444-555555555555
    /// ```
    /// </summary>
    [RootlyResourceType("rootly:index/cause:Cause")]
    public partial class Cause : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the cause
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the cause
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Cause resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cause(string name, CauseArgs? args = null, CustomResourceOptions? options = null)
            : base("rootly:index/cause:Cause", name, args ?? new CauseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cause(string name, Input<string> id, CauseState? state = null, CustomResourceOptions? options = null)
            : base("rootly:index/cause:Cause", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cause resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cause Get(string name, Input<string> id, CauseState? state = null, CustomResourceOptions? options = null)
        {
            return new Cause(name, id, state, options);
        }
    }

    public sealed class CauseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the cause
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the cause
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CauseArgs()
        {
        }
    }

    public sealed class CauseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the cause
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the cause
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CauseState()
        {
        }
    }
}
