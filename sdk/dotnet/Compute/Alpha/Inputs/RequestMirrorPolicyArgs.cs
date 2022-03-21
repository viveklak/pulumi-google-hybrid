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
    /// A policy that specifies how requests intended for the route's backends are shadowed to a separate mirrored backend service. The load balancer doesn't wait for responses from the shadow service. Before sending traffic to the shadow service, the host or authority header is suffixed with -shadow.
    /// </summary>
    public sealed class RequestMirrorPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full or partial URL to the BackendService resource being mirrored to.
        /// </summary>
        [Input("backendService")]
        public Input<string>? BackendService { get; set; }

        public RequestMirrorPolicyArgs()
        {
        }
    }
}
