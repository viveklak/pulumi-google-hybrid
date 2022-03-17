// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Domains.V1Beta1.Outputs
{

    /// <summary>
    /// Defines a Delegation Signer (DS) record, which is needed to enable DNSSEC for a domain. It contains a digest (hash) of a DNSKEY record that must be present in the domain's DNS zone.
    /// </summary>
    [OutputType]
    public sealed class DsRecordResponse
    {
        /// <summary>
        /// The algorithm used to generate the referenced DNSKEY.
        /// </summary>
        public readonly string Algorithm;
        /// <summary>
        /// The digest generated from the referenced DNSKEY.
        /// </summary>
        public readonly string Digest;
        /// <summary>
        /// The hash function used to generate the digest of the referenced DNSKEY.
        /// </summary>
        public readonly string DigestType;
        /// <summary>
        /// The key tag of the record. Must be set in range 0 -- 65535.
        /// </summary>
        public readonly int KeyTag;

        [OutputConstructor]
        private DsRecordResponse(
            string algorithm,

            string digest,

            string digestType,

            int keyTag)
        {
            Algorithm = algorithm;
            Digest = digest;
            DigestType = digestType;
            KeyTag = keyTag;
        }
    }
}
