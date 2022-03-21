// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2.Outputs
{

    /// <summary>
    /// ResourceRecordSet data for one geo location.
    /// </summary>
    [OutputType]
    public sealed class RRSetRoutingPolicyGeoPolicyGeoPolicyItemResponse
    {
        public readonly string Kind;
        /// <summary>
        /// The geo-location granularity is a GCP region. This location string should correspond to a GCP region. e.g. "us-east1", "southamerica-east1", "asia-east1", etc.
        /// </summary>
        public readonly string Location;
        public readonly ImmutableArray<string> Rrdatas;
        /// <summary>
        /// DNSSEC generated signatures for all the rrdata within this item. Note that if health checked targets are provided for DNSSEC enabled zones, there's a restriction of 1 ip per item. .
        /// </summary>
        public readonly ImmutableArray<string> SignatureRrdatas;

        [OutputConstructor]
        private RRSetRoutingPolicyGeoPolicyGeoPolicyItemResponse(
            string kind,

            string location,

            ImmutableArray<string> rrdatas,

            ImmutableArray<string> signatureRrdatas)
        {
            Kind = kind;
            Location = location;
            Rrdatas = rrdatas;
            SignatureRrdatas = signatureRrdatas;
        }
    }
}
