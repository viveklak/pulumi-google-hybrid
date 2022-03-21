// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Outputs
{

    /// <summary>
    /// User signature.
    /// </summary>
    [OutputType]
    public sealed class SignatureResponse
    {
        /// <summary>
        /// Optional. An image of the user's signature.
        /// </summary>
        public readonly Outputs.ImageResponse Image;
        /// <summary>
        /// Optional. Metadata associated with the user's signature. For example, the user's name or the user's title.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// Optional. Timestamp of the signature.
        /// </summary>
        public readonly string SignatureTime;
        /// <summary>
        /// User's UUID provided by the client.
        /// </summary>
        public readonly string UserId;

        [OutputConstructor]
        private SignatureResponse(
            Outputs.ImageResponse image,

            ImmutableDictionary<string, string> metadata,

            string signatureTime,

            string userId)
        {
            Image = image;
            Metadata = metadata;
            SignatureTime = signatureTime;
            UserId = userId;
        }
    }
}
