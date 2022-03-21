// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Domains.V1Alpha2.Outputs
{

    /// <summary>
    /// Configuration for using the free DNS zone provided by Google Domains as a `Registration`'s `dns_provider`. You cannot configure the DNS zone itself using the API. To configure the DNS zone, go to [Google Domains](https://domains.google/).
    /// </summary>
    [OutputType]
    public sealed class GoogleDomainsDnsResponse
    {
        /// <summary>
        /// The list of DS records published for this domain. The list is automatically populated when `ds_state` is `DS_RECORDS_PUBLISHED`, otherwise it remains empty.
        /// </summary>
        public readonly ImmutableArray<Outputs.DsRecordResponse> DsRecords;
        /// <summary>
        /// The state of DS records for this domain. Used to enable or disable automatic DNSSEC.
        /// </summary>
        public readonly string DsState;
        /// <summary>
        /// A list of name servers that store the DNS zone for this domain. Each name server is a domain name, with Unicode domain names expressed in Punycode format. This field is automatically populated with the name servers assigned to the Google Domains DNS zone.
        /// </summary>
        public readonly ImmutableArray<string> NameServers;

        [OutputConstructor]
        private GoogleDomainsDnsResponse(
            ImmutableArray<Outputs.DsRecordResponse> dsRecords,

            string dsState,

            ImmutableArray<string> nameServers)
        {
            DsRecords = dsRecords;
            DsState = dsState;
            NameServers = nameServers;
        }
    }
}
