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
    /// Specifies how to route matched traffic.
    /// </summary>
    [OutputType]
    public sealed class GrpcRouteRouteActionResponse
    {
        /// <summary>
        /// Optional. The destination services to which traffic should be forwarded. If multiple destinations are specified, traffic will be split between Backend Service(s) according to the weight field of these destinations.
        /// </summary>
        public readonly ImmutableArray<Outputs.GrpcRouteDestinationResponse> Destinations;
        /// <summary>
        /// Optional. The specification for fault injection introduced into traffic to test the resiliency of clients to destination service failure. As part of fault injection, when clients send requests to a destination, delays can be introduced on a percentage of requests before sending those requests to the destination service. Similarly requests from clients can be aborted by for a percentage of requests. timeout and retry_policy will be ignored by clients that are configured with a fault_injection_policy
        /// </summary>
        public readonly Outputs.GrpcRouteFaultInjectionPolicyResponse FaultInjectionPolicy;
        /// <summary>
        /// Optional. Specifies the retry policy associated with this route.
        /// </summary>
        public readonly Outputs.GrpcRouteRetryPolicyResponse RetryPolicy;
        /// <summary>
        /// Optional. Specifies the timeout for selected route. Timeout is computed from the time the request has been fully processed (i.e. end of stream) up until the response has been completely processed. Timeout includes all retries.
        /// </summary>
        public readonly string Timeout;

        [OutputConstructor]
        private GrpcRouteRouteActionResponse(
            ImmutableArray<Outputs.GrpcRouteDestinationResponse> destinations,

            Outputs.GrpcRouteFaultInjectionPolicyResponse faultInjectionPolicy,

            Outputs.GrpcRouteRetryPolicyResponse retryPolicy,

            string timeout)
        {
            Destinations = destinations;
            FaultInjectionPolicy = faultInjectionPolicy;
            RetryPolicy = retryPolicy;
            Timeout = timeout;
        }
    }
}
