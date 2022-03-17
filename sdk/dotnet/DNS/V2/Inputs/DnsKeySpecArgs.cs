// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V2.Inputs
{

    /// <summary>
    /// Parameters for DnsKey key generation. Used for generating initial keys for a new ManagedZone and as default when adding a new DnsKey.
    /// </summary>
    public sealed class DnsKeySpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String mnemonic specifying the DNSSEC algorithm of this key.
        /// </summary>
        [Input("algorithm")]
        public Input<Pulumi.GoogleNative.DNS.V2.DnsKeySpecAlgorithm>? Algorithm { get; set; }

        /// <summary>
        /// Length of the keys in bits.
        /// </summary>
        [Input("keyLength")]
        public Input<int>? KeyLength { get; set; }

        /// <summary>
        /// Specifies whether this is a key signing key (KSK) or a zone signing key (ZSK). Key signing keys have the Secure Entry Point flag set and, when active, are only used to sign resource record sets of type DNSKEY. Zone signing keys do not have the Secure Entry Point flag set and are used to sign all other types of resource record sets.
        /// </summary>
        [Input("keyType")]
        public Input<Pulumi.GoogleNative.DNS.V2.DnsKeySpecKeyType>? KeyType { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        public DnsKeySpecArgs()
        {
        }
    }
}
