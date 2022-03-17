// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Iap.V1
{
    public static class GetIdentityAwareProxyClient
    {
        /// <summary>
        /// Retrieves an Identity Aware Proxy (IAP) OAuth client. Requires that the client is owned by IAP.
        /// </summary>
        public static Task<GetIdentityAwareProxyClientResult> InvokeAsync(GetIdentityAwareProxyClientArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIdentityAwareProxyClientResult>("google-native:iap/v1:getIdentityAwareProxyClient", args ?? new GetIdentityAwareProxyClientArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves an Identity Aware Proxy (IAP) OAuth client. Requires that the client is owned by IAP.
        /// </summary>
        public static Output<GetIdentityAwareProxyClientResult> Invoke(GetIdentityAwareProxyClientInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIdentityAwareProxyClientResult>("google-native:iap/v1:getIdentityAwareProxyClient", args ?? new GetIdentityAwareProxyClientInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdentityAwareProxyClientArgs : Pulumi.InvokeArgs
    {
        [Input("brandId", required: true)]
        public string BrandId { get; set; } = null!;

        [Input("identityAwareProxyClientId", required: true)]
        public string IdentityAwareProxyClientId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetIdentityAwareProxyClientArgs()
        {
        }
    }

    public sealed class GetIdentityAwareProxyClientInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("brandId", required: true)]
        public Input<string> BrandId { get; set; } = null!;

        [Input("identityAwareProxyClientId", required: true)]
        public Input<string> IdentityAwareProxyClientId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetIdentityAwareProxyClientInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIdentityAwareProxyClientResult
    {
        /// <summary>
        /// Human-friendly name given to the OAuth client.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Unique identifier of the OAuth client.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Client secret of the OAuth client.
        /// </summary>
        public readonly string Secret;

        [OutputConstructor]
        private GetIdentityAwareProxyClientResult(
            string displayName,

            string name,

            string secret)
        {
            DisplayName = displayName;
            Name = name;
            Secret = secret;
        }
    }
}
