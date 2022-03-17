// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V2.Outputs
{

    [OutputType]
    public sealed class ManagedZoneDnsSecConfigResponse
    {
        /// <summary>
        /// Specifies parameters for generating initial DnsKeys for this ManagedZone. Can only be changed while the state is OFF.
        /// </summary>
        public readonly ImmutableArray<Outputs.DnsKeySpecResponse> DefaultKeySpecs;
        public readonly string Kind;
        /// <summary>
        /// Specifies the mechanism for authenticated denial-of-existence responses. Can only be changed while the state is OFF.
        /// </summary>
        public readonly string NonExistence;
        /// <summary>
        /// Specifies whether DNSSEC is enabled, and what mode it is in.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private ManagedZoneDnsSecConfigResponse(
            ImmutableArray<Outputs.DnsKeySpecResponse> defaultKeySpecs,

            string kind,

            string nonExistence,

            string state)
        {
            DefaultKeySpecs = defaultKeySpecs;
            Kind = kind;
            NonExistence = nonExistence;
            State = state;
        }
    }
}
