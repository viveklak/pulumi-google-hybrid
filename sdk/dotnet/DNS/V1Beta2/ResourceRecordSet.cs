// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2
{
    /// <summary>
    /// Creates a new ResourceRecordSet.
    /// </summary>
    [GoogleNativeResourceType("google-native:dns/v1beta2:ResourceRecordSet")]
    public partial class ResourceRecordSet : Pulumi.CustomResource
    {
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// For example, www.example.com.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configures dynamic query responses based on geo location of querying user or a weighted round robin based routing policy. A ResourceRecordSet should only have either rrdata (static) or routing_policy (dynamic). An error is returned otherwise.
        /// </summary>
        [Output("routingPolicy")]
        public Output<Outputs.RRSetRoutingPolicyResponse> RoutingPolicy { get; private set; } = null!;

        /// <summary>
        /// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        /// </summary>
        [Output("rrdatas")]
        public Output<ImmutableArray<string>> Rrdatas { get; private set; } = null!;

        /// <summary>
        /// As defined in RFC 4034 (section 3.2).
        /// </summary>
        [Output("signatureRrdatas")]
        public Output<ImmutableArray<string>> SignatureRrdatas { get; private set; } = null!;

        /// <summary>
        /// Number of seconds that this ResourceRecordSet can be cached by resolvers.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// The identifier of a supported record type. See the list of Supported DNS record types.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceRecordSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceRecordSet(string name, ResourceRecordSetArgs args, CustomResourceOptions? options = null)
            : base("google-native:dns/v1beta2:ResourceRecordSet", name, args ?? new ResourceRecordSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceRecordSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dns/v1beta2:ResourceRecordSet", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceRecordSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceRecordSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourceRecordSet(name, id, options);
        }
    }

    public sealed class ResourceRecordSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
        /// </summary>
        [Input("clientOperationId")]
        public Input<string>? ClientOperationId { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("managedZone", required: true)]
        public Input<string> ManagedZone { get; set; } = null!;

        /// <summary>
        /// For example, www.example.com.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Configures dynamic query responses based on geo location of querying user or a weighted round robin based routing policy. A ResourceRecordSet should only have either rrdata (static) or routing_policy (dynamic). An error is returned otherwise.
        /// </summary>
        [Input("routingPolicy")]
        public Input<Inputs.RRSetRoutingPolicyArgs>? RoutingPolicy { get; set; }

        [Input("rrdatas")]
        private InputList<string>? _rrdatas;

        /// <summary>
        /// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        /// </summary>
        public InputList<string> Rrdatas
        {
            get => _rrdatas ?? (_rrdatas = new InputList<string>());
            set => _rrdatas = value;
        }

        [Input("signatureRrdatas")]
        private InputList<string>? _signatureRrdatas;

        /// <summary>
        /// As defined in RFC 4034 (section 3.2).
        /// </summary>
        public InputList<string> SignatureRrdatas
        {
            get => _signatureRrdatas ?? (_signatureRrdatas = new InputList<string>());
            set => _signatureRrdatas = value;
        }

        /// <summary>
        /// Number of seconds that this ResourceRecordSet can be cached by resolvers.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The identifier of a supported record type. See the list of Supported DNS record types.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ResourceRecordSetArgs()
        {
        }
    }
}
