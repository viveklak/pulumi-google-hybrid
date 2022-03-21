// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2.Outputs
{

    [OutputType]
    public sealed class ManagedZoneForwardingConfigResponse
    {
        public readonly string Kind;
        /// <summary>
        /// List of target name servers to forward to. Cloud DNS selects the best available name server if more than one target is given.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedZoneForwardingConfigNameServerTargetResponse> TargetNameServers;

        [OutputConstructor]
        private ManagedZoneForwardingConfigResponse(
            string kind,

            ImmutableArray<Outputs.ManagedZoneForwardingConfigNameServerTargetResponse> targetNameServers)
        {
            Kind = kind;
            TargetNameServers = targetNameServers;
        }
    }
}
