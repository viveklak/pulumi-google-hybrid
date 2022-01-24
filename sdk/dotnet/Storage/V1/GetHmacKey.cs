// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Storage.V1
{
    public static class GetHmacKey
    {
        /// <summary>
        /// Retrieves an HMAC key's metadata
        /// </summary>
        public static Task<GetHmacKeyResult> InvokeAsync(GetHmacKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHmacKeyResult>("google-native:storage/v1:getHmacKey", args ?? new GetHmacKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves an HMAC key's metadata
        /// </summary>
        public static Output<GetHmacKeyResult> Invoke(GetHmacKeyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetHmacKeyResult>("google-native:storage/v1:getHmacKey", args ?? new GetHmacKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHmacKeyArgs : Pulumi.InvokeArgs
    {
        [Input("accessId", required: true)]
        public string AccessId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("userProject")]
        public string? UserProject { get; set; }

        public GetHmacKeyArgs()
        {
        }
    }

    public sealed class GetHmacKeyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("accessId", required: true)]
        public Input<string> AccessId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("userProject")]
        public Input<string>? UserProject { get; set; }

        public GetHmacKeyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHmacKeyResult
    {
        /// <summary>
        /// The ID of the HMAC Key.
        /// </summary>
        public readonly string AccessId;
        /// <summary>
        /// HTTP 1.1 Entity tag for the HMAC key.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The kind of item this is. For HMAC Key metadata, this is always storage#hmacKeyMetadata.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Project ID owning the service account to which the key authenticates.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The link to this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The email address of the key's associated service account.
        /// </summary>
        public readonly string ServiceAccountEmail;
        /// <summary>
        /// The state of the key. Can be one of ACTIVE, INACTIVE, or DELETED.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The creation time of the HMAC key in RFC 3339 format.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The last modification time of the HMAC key metadata in RFC 3339 format.
        /// </summary>
        public readonly string Updated;

        [OutputConstructor]
        private GetHmacKeyResult(
            string accessId,

            string etag,

            string kind,

            string project,

            string selfLink,

            string serviceAccountEmail,

            string state,

            string timeCreated,

            string updated)
        {
            AccessId = accessId;
            Etag = etag;
            Kind = kind;
            Project = project;
            SelfLink = selfLink;
            ServiceAccountEmail = serviceAccountEmail;
            State = state;
            TimeCreated = timeCreated;
            Updated = updated;
        }
    }
}
