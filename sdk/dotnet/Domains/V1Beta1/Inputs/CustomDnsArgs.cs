// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Domains.V1Beta1.Inputs
{

    /// <summary>
    /// Configuration for an arbitrary DNS provider.
    /// </summary>
    public sealed class CustomDnsArgs : Pulumi.ResourceArgs
    {
        [Input("dsRecords")]
        private InputList<Inputs.DsRecordArgs>? _dsRecords;

        /// <summary>
        /// The list of DS records for this domain, which are used to enable DNSSEC. The domain's DNS provider can provide the values to set here. If this field is empty, DNSSEC is disabled.
        /// </summary>
        public InputList<Inputs.DsRecordArgs> DsRecords
        {
            get => _dsRecords ?? (_dsRecords = new InputList<Inputs.DsRecordArgs>());
            set => _dsRecords = value;
        }

        [Input("nameServers", required: true)]
        private InputList<string>? _nameServers;

        /// <summary>
        /// A list of name servers that store the DNS zone for this domain. Each name server is a domain name, with Unicode domain names expressed in Punycode format.
        /// </summary>
        public InputList<string> NameServers
        {
            get => _nameServers ?? (_nameServers = new InputList<string>());
            set => _nameServers = value;
        }

        public CustomDnsArgs()
        {
        }
    }
}
