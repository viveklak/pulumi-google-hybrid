// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Uploads a ZIP-formatted shared flow configuration bundle to an organization. If the shared flow already exists, this creates a new revision of it. If the shared flow does not exist, this creates it. Once imported, the shared flow revision must be deployed before it can be accessed at runtime. The size limit of a shared flow bundle is 15 MB.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:OrganizationSharedflow")]
    public partial class OrganizationSharedflow : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the most recently created revision for this shared flow.
        /// </summary>
        [Output("latestRevisionId")]
        public Output<string> LatestRevisionId { get; private set; } = null!;

        /// <summary>
        /// Metadata describing the shared flow.
        /// </summary>
        [Output("metaData")]
        public Output<Outputs.GoogleCloudApigeeV1EntityMetadataResponse> MetaData { get; private set; } = null!;

        /// <summary>
        /// The ID of the shared flow.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of revisions of this shared flow.
        /// </summary>
        [Output("revision")]
        public Output<ImmutableArray<string>> Revision { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationSharedflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationSharedflow(string name, OrganizationSharedflowArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationSharedflow", name, args ?? new OrganizationSharedflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationSharedflow(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationSharedflow", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationSharedflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationSharedflow Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationSharedflow(name, id, options);
        }
    }

    public sealed class OrganizationSharedflowArgs : Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The HTTP Content-Type header value specifying the content type of the body.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// The HTTP request/response body as raw binary.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        [Input("extensions")]
        private InputList<ImmutableDictionary<string, string>>? _extensions;

        /// <summary>
        /// Application specific response metadata. Must be set in the first response for streaming APIs.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Extensions
        {
            get => _extensions ?? (_extensions = new InputList<ImmutableDictionary<string, string>>());
            set => _extensions = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public OrganizationSharedflowArgs()
        {
        }
    }
}
