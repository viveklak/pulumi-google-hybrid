// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a BackendBucket resource in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:BackendBucket")]
    public partial class BackendBucket : Pulumi.CustomResource
    {
        /// <summary>
        /// Cloud Storage bucket name.
        /// </summary>
        [Output("bucketName")]
        public Output<string> BucketName { get; private set; } = null!;

        /// <summary>
        /// Cloud CDN configuration for this BackendBucket.
        /// </summary>
        [Output("cdnPolicy")]
        public Output<Outputs.BackendBucketCdnPolicyResponse> CdnPolicy { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied responses.
        /// </summary>
        [Output("customResponseHeaders")]
        public Output<ImmutableArray<string>> CustomResponseHeaders { get; private set; } = null!;

        /// <summary>
        /// An optional textual description of the resource; provided by the client when the resource is created.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// If true, enable Cloud CDN for this BackendBucket.
        /// </summary>
        [Output("enableCdn")]
        public Output<bool> EnableCdn { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a BackendBucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackendBucket(string name, BackendBucketArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:BackendBucket", name, args ?? new BackendBucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackendBucket(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:BackendBucket", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing BackendBucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackendBucket Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BackendBucket(name, id, options);
        }
    }

    public sealed class BackendBucketArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud Storage bucket name.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// Cloud CDN configuration for this BackendBucket.
        /// </summary>
        [Input("cdnPolicy")]
        public Input<Inputs.BackendBucketCdnPolicyArgs>? CdnPolicy { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        [Input("customResponseHeaders")]
        private InputList<string>? _customResponseHeaders;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied responses.
        /// </summary>
        public InputList<string> CustomResponseHeaders
        {
            get => _customResponseHeaders ?? (_customResponseHeaders = new InputList<string>());
            set => _customResponseHeaders = value;
        }

        /// <summary>
        /// An optional textual description of the resource; provided by the client when the resource is created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If true, enable Cloud CDN for this BackendBucket.
        /// </summary>
        [Input("enableCdn")]
        public Input<bool>? EnableCdn { get; set; }

        /// <summary>
        /// [Output Only] Unique identifier for the resource; defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public BackendBucketArgs()
        {
        }
    }
}
