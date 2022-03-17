// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.V2Beta
{
    /// <summary>
    /// Creates a composite type.
    /// </summary>
    [GoogleNativeResourceType("google-native:deploymentmanager/v2beta:CompositeType")]
    public partial class CompositeType : Pulumi.CustomResource
    {
        /// <summary>
        /// An optional textual description of the resource; provided by the client when the resource is created.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("insertTime")]
        public Output<string> InsertTime { get; private set; } = null!;

        /// <summary>
        /// Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.CompositeTypeLabelEntryResponse>> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Operation that most recently ran, or is currently running, on this composite type.
        /// </summary>
        [Output("operation")]
        public Output<Outputs.OperationResponse> Operation { get; private set; } = null!;

        /// <summary>
        /// Server defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Files for the template type.
        /// </summary>
        [Output("templateContents")]
        public Output<Outputs.TemplateContentsResponse> TemplateContents { get; private set; } = null!;


        /// <summary>
        /// Create a CompositeType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CompositeType(string name, CompositeTypeArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:deploymentmanager/v2beta:CompositeType", name, args ?? new CompositeTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CompositeType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:deploymentmanager/v2beta:CompositeType", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CompositeType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CompositeType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CompositeType(name, id, options);
        }
    }

    public sealed class CompositeTypeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional textual description of the resource; provided by the client when the resource is created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("labels")]
        private InputList<Inputs.CompositeTypeLabelEntryArgs>? _labels;

        /// <summary>
        /// Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        /// </summary>
        public InputList<Inputs.CompositeTypeLabelEntryArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.CompositeTypeLabelEntryArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("status")]
        public Input<Pulumi.GoogleNative.DeploymentManager.V2Beta.CompositeTypeStatus>? Status { get; set; }

        /// <summary>
        /// Files for the template type.
        /// </summary>
        [Input("templateContents")]
        public Input<Inputs.TemplateContentsArgs>? TemplateContents { get; set; }

        public CompositeTypeArgs()
        {
        }
    }
}
