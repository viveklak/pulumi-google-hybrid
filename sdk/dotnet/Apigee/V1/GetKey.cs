// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetKey
    {
        /// <summary>
        /// Gets details for a consumer key for a developer app, including the key and secret value, associated API products, and other information.
        /// </summary>
        public static Task<GetKeyResult> InvokeAsync(GetKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("google-native:apigee/v1:getKey", args ?? new GetKeyArgs(), options.WithVersion());
    }


    public sealed class GetKeyArgs : Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        [Input("developerId", required: true)]
        public string DeveloperId { get; set; } = null!;

        [Input("keyId", required: true)]
        public string KeyId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetKeyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeyResult
    {
        /// <summary>
        /// List of API products for which the credential can be used. **Note**: Do not specify the list of API products when creating a consumer key and secret for a developer app. Instead, use the UpdateDeveloperAppKey API to make the association after the consumer key and secret are created.
        /// </summary>
        public readonly ImmutableArray<object> ApiProducts;
        /// <summary>
        /// List of attributes associated with the credential.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudApigeeV1AttributeResponse> Attributes;
        /// <summary>
        /// Consumer key.
        /// </summary>
        public readonly string ConsumerKey;
        /// <summary>
        /// Secret key.
        /// </summary>
        public readonly string ConsumerSecret;
        /// <summary>
        /// Time the developer app expires in milliseconds since epoch.
        /// </summary>
        public readonly string ExpiresAt;
        /// <summary>
        /// Input only. Expiration time, in seconds, for the consumer key. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
        /// </summary>
        public readonly string ExpiresInSeconds;
        /// <summary>
        /// Time the developer app was created in milliseconds since epoch.
        /// </summary>
        public readonly string IssuedAt;
        /// <summary>
        /// Scopes to apply to the app. The specified scope names must already be defined for the API product that you associate with the app.
        /// </summary>
        public readonly ImmutableArray<string> Scopes;
        /// <summary>
        /// Status of the credential. Valid values include `approved` or `revoked`.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetKeyResult(
            ImmutableArray<object> apiProducts,

            ImmutableArray<Outputs.GoogleCloudApigeeV1AttributeResponse> attributes,

            string consumerKey,

            string consumerSecret,

            string expiresAt,

            string expiresInSeconds,

            string issuedAt,

            ImmutableArray<string> scopes,

            string status)
        {
            ApiProducts = apiProducts;
            Attributes = attributes;
            ConsumerKey = consumerKey;
            ConsumerSecret = consumerSecret;
            ExpiresAt = expiresAt;
            ExpiresInSeconds = expiresInSeconds;
            IssuedAt = issuedAt;
            Scopes = scopes;
            Status = status;
        }
    }
}
