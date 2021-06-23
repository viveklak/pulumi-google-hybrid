// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1
{
    /// <summary>
    /// Create a CaPool.
    /// </summary>
    [GoogleNativeResourceType("google-native:privateca/v1:CaPool")]
    public partial class CaPool : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. The IssuancePolicy to control how Certificates will be issued from this CaPool.
        /// </summary>
        [Output("issuancePolicy")]
        public Output<Outputs.IssuancePolicyResponse> IssuancePolicy { get; private set; } = null!;

        /// <summary>
        /// Optional. Labels with user-defined metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name for this CaPool in the format `projects/*/locations/*/caPools/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
        /// </summary>
        [Output("publishingOptions")]
        public Output<Outputs.PublishingOptionsResponse> PublishingOptions { get; private set; } = null!;

        /// <summary>
        /// Required. Immutable. The Tier of this CaPool.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;


        /// <summary>
        /// Create a CaPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CaPool(string name, CaPoolArgs args, CustomResourceOptions? options = null)
            : base("google-native:privateca/v1:CaPool", name, args ?? new CaPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CaPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:privateca/v1:CaPool", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CaPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CaPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CaPool(name, id, options);
        }
    }

    public sealed class CaPoolArgs : Pulumi.ResourceArgs
    {
        [Input("caPoolId", required: true)]
        public Input<string> CaPoolId { get; set; } = null!;

        /// <summary>
        /// Optional. The IssuancePolicy to control how Certificates will be issued from this CaPool.
        /// </summary>
        [Input("issuancePolicy")]
        public Input<Inputs.IssuancePolicyArgs>? IssuancePolicy { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Labels with user-defined metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Optional. The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
        /// </summary>
        [Input("publishingOptions")]
        public Input<Inputs.PublishingOptionsArgs>? PublishingOptions { get; set; }

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Required. Immutable. The Tier of this CaPool.
        /// </summary>
        [Input("tier")]
        public Input<Pulumi.GoogleNative.Privateca.V1.CaPoolTier>? Tier { get; set; }

        public CaPoolArgs()
        {
        }
    }
}
