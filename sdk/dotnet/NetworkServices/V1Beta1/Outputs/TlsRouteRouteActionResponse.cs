// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1.Outputs
{

    /// <summary>
    /// The specifications for routing traffic and applying associated policies.
    /// </summary>
    [OutputType]
    public sealed class TlsRouteRouteActionResponse
    {
        /// <summary>
        /// The destination services to which traffic should be forwarded. At least one destination service is required.
        /// </summary>
        public readonly ImmutableArray<Outputs.TlsRouteRouteDestinationResponse> Destinations;

        [OutputConstructor]
        private TlsRouteRouteActionResponse(ImmutableArray<Outputs.TlsRouteRouteDestinationResponse> destinations)
        {
            Destinations = destinations;
        }
    }
}
