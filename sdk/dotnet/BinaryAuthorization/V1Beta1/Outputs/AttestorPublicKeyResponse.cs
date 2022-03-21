// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BinaryAuthorization.V1Beta1.Outputs
{

    /// <summary>
    /// An attestor public key that will be used to verify attestations signed by this attestor.
    /// </summary>
    [OutputType]
    public sealed class AttestorPublicKeyResponse
    {
        /// <summary>
        /// ASCII-armored representation of a PGP public key, as the entire output by the command `gpg --export --armor foo@example.com` (either LF or CRLF line endings). When using this field, `id` should be left blank. The BinAuthz API handlers will calculate the ID and fill it in automatically. BinAuthz computes this ID as the OpenPGP RFC4880 V4 fingerprint, represented as upper-case hex. If `id` is provided by the caller, it will be overwritten by the API-calculated ID.
        /// </summary>
        public readonly string AsciiArmoredPgpPublicKey;
        /// <summary>
        /// Optional. A descriptive comment. This field may be updated.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// A raw PKIX SubjectPublicKeyInfo format public key. NOTE: `id` may be explicitly provided by the caller when using this type of public key, but it MUST be a valid RFC3986 URI. If `id` is left blank, a default one will be computed based on the digest of the DER encoding of the public key.
        /// </summary>
        public readonly Outputs.PkixPublicKeyResponse PkixPublicKey;

        [OutputConstructor]
        private AttestorPublicKeyResponse(
            string asciiArmoredPgpPublicKey,

            string comment,

            Outputs.PkixPublicKeyResponse pkixPublicKey)
        {
            AsciiArmoredPgpPublicKey = asciiArmoredPgpPublicKey;
            Comment = comment;
            PkixPublicKey = pkixPublicKey;
        }
    }
}
