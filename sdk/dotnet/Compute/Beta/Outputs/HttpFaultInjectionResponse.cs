// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// The specification for fault injection introduced into traffic to test the resiliency of clients to backend service failure. As part of fault injection, when clients send requests to a backend service, delays can be introduced by the load balancer on a percentage of requests before sending those request to the backend service. Similarly requests from clients can be aborted by the load balancer for a percentage of requests.
    /// </summary>
    [OutputType]
    public sealed class HttpFaultInjectionResponse
    {
        /// <summary>
        /// The specification for how client requests are aborted as part of fault injection.
        /// </summary>
        public readonly Outputs.HttpFaultAbortResponse Abort;
        /// <summary>
        /// The specification for how client requests are delayed as part of fault injection, before being sent to a backend service.
        /// </summary>
        public readonly Outputs.HttpFaultDelayResponse Delay;

        [OutputConstructor]
        private HttpFaultInjectionResponse(
            Outputs.HttpFaultAbortResponse abort,

            Outputs.HttpFaultDelayResponse delay)
        {
            Abort = abort;
            Delay = delay;
        }
    }
}
