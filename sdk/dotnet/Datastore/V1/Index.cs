// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastore.V1
{
    /// <summary>
    /// Creates the specified index. A newly created index's initial state is `CREATING`. On completion of the returned google.longrunning.Operation, the state will be `READY`. If the index already exists, the call will return an `ALREADY_EXISTS` status. During index creation, the process could result in an error, in which case the index will move to the `ERROR` state. The process can be recovered by fixing the data that caused the error, removing the index with delete, then re-creating the index with create. Indexes with a single property cannot be created.
    /// </summary>
    [GoogleNativeResourceType("google-native:datastore/v1:Index")]
    public partial class Index : Pulumi.CustomResource
    {
        /// <summary>
        /// Required. The index's ancestor mode. Must not be ANCESTOR_MODE_UNSPECIFIED.
        /// </summary>
        [Output("ancestor")]
        public Output<string> Ancestor { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the index.
        /// </summary>
        [Output("indexId")]
        public Output<string> IndexId { get; private set; } = null!;

        /// <summary>
        /// Required. The entity kind to which this index applies.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Required. An ordered sequence of property names and their index attributes.
        /// </summary>
        [Output("properties")]
        public Output<ImmutableArray<Outputs.GoogleDatastoreAdminV1IndexedPropertyResponse>> Properties { get; private set; } = null!;

        /// <summary>
        /// The state of the index.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Index resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Index(string name, IndexArgs args, CustomResourceOptions? options = null)
            : base("google-native:datastore/v1:Index", name, args ?? new IndexArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Index(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:datastore/v1:Index", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Index resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Index Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Index(name, id, options);
        }
    }

    public sealed class IndexArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The index's ancestor mode. Must not be ANCESTOR_MODE_UNSPECIFIED.
        /// </summary>
        [Input("ancestor")]
        public Input<string>? Ancestor { get; set; }

        /// <summary>
        /// Required. The entity kind to which this index applies.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("properties")]
        private InputList<Inputs.GoogleDatastoreAdminV1IndexedPropertyArgs>? _properties;

        /// <summary>
        /// Required. An ordered sequence of property names and their index attributes.
        /// </summary>
        public InputList<Inputs.GoogleDatastoreAdminV1IndexedPropertyArgs> Properties
        {
            get => _properties ?? (_properties = new InputList<Inputs.GoogleDatastoreAdminV1IndexedPropertyArgs>());
            set => _properties = value;
        }

        public IndexArgs()
        {
        }
    }
}
