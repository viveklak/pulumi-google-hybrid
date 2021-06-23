// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Settings controlling the volume of connections to a backend service.
    /// </summary>
    public sealed class CircuitBreakersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timeout for new network connections to hosts.
        /// </summary>
        [Input("connectTimeout")]
        public Input<Inputs.DurationArgs>? ConnectTimeout { get; set; }

        /// <summary>
        /// Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("maxPendingRequests")]
        public Input<int>? MaxPendingRequests { get; set; }

        /// <summary>
        /// The maximum number of parallel requests that allowed to the backend service. If not specified, there is no limit.
        /// </summary>
        [Input("maxRequests")]
        public Input<int>? MaxRequests { get; set; }

        /// <summary>
        /// Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("maxRequestsPerConnection")]
        public Input<int>? MaxRequestsPerConnection { get; set; }

        /// <summary>
        /// Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        public CircuitBreakersArgs()
        {
        }
    }
}
