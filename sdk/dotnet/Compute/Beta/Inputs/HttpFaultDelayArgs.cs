// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Specifies the delay introduced by the load balancer before forwarding the request to the backend service as part of fault injection.
    /// </summary>
    public sealed class HttpFaultDelayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the value of the fixed delay interval.
        /// </summary>
        [Input("fixedDelay")]
        public Input<Inputs.DurationArgs>? FixedDelay { get; set; }

        /// <summary>
        /// The percentage of traffic for connections, operations, or requests for which a delay is introduced as part of fault injection. The value must be from 0.0 to 100.0 inclusive.
        /// </summary>
        [Input("percentage")]
        public Input<double>? Percentage { get; set; }

        public HttpFaultDelayArgs()
        {
        }
    }
}
