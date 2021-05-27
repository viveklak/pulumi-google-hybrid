// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Iap.V1
{
    /// <summary>
    /// Creates an Identity Aware Proxy (IAP) OAuth client. The client is owned by IAP. Requires that the brand for the project exists and that it is set for internal-only use.
    /// </summary>
    [GoogleNativeResourceType("google-native:iap/v1:BrandIdentityAwareProxyClient")]
    public partial class BrandIdentityAwareProxyClient : Pulumi.CustomResource
    {
        /// <summary>
        /// Human-friendly name given to the OAuth client.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Unique identifier of the OAuth client.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Client secret of the OAuth client.
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;


        /// <summary>
        /// Create a BrandIdentityAwareProxyClient resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BrandIdentityAwareProxyClient(string name, BrandIdentityAwareProxyClientArgs args, CustomResourceOptions? options = null)
            : base("google-native:iap/v1:BrandIdentityAwareProxyClient", name, args ?? new BrandIdentityAwareProxyClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BrandIdentityAwareProxyClient(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:iap/v1:BrandIdentityAwareProxyClient", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing BrandIdentityAwareProxyClient resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BrandIdentityAwareProxyClient Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BrandIdentityAwareProxyClient(name, id, options);
        }
    }

    public sealed class BrandIdentityAwareProxyClientArgs : Pulumi.ResourceArgs
    {
        [Input("brandId", required: true)]
        public Input<string> BrandId { get; set; } = null!;

        /// <summary>
        /// Human-friendly name given to the OAuth client.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public BrandIdentityAwareProxyClientArgs()
        {
        }
    }
}
