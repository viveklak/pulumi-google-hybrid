// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// Identity-Aware Proxy
    /// </summary>
    [OutputType]
    public sealed class BackendServiceIAPResponse
    {
        /// <summary>
        /// Whether the serving infrastructure will authenticate and authorize all incoming requests. If true, the oauth2ClientId and oauth2ClientSecret fields must be non-empty.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// OAuth2 client ID to use for the authentication flow.
        /// </summary>
        public readonly string Oauth2ClientId;
        /// <summary>
        /// OAuth2 client secret to use for the authentication flow. For security reasons, this value cannot be retrieved via the API. Instead, the SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field. @InputOnly
        /// </summary>
        public readonly string Oauth2ClientSecret;
        /// <summary>
        /// SHA256 hash value for the field oauth2_client_secret above.
        /// </summary>
        public readonly string Oauth2ClientSecretSha256;

        [OutputConstructor]
        private BackendServiceIAPResponse(
            bool enabled,

            string oauth2ClientId,

            string oauth2ClientSecret,

            string oauth2ClientSecretSha256)
        {
            Enabled = enabled;
            Oauth2ClientId = oauth2ClientId;
            Oauth2ClientSecret = oauth2ClientSecret;
            Oauth2ClientSecretSha256 = oauth2ClientSecretSha256;
        }
    }
}
