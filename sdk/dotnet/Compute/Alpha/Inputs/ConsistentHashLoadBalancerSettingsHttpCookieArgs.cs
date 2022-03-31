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
    /// The information about the HTTP Cookie on which the hash function is based for load balancing policies that use a consistent hash.
    /// </summary>
    public sealed class ConsistentHashLoadBalancerSettingsHttpCookieArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the cookie.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Path to set for the cookie.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Lifetime of the cookie.
        /// </summary>
        [Input("ttl")]
        public Input<Inputs.DurationArgs>? Ttl { get; set; }

        public ConsistentHashLoadBalancerSettingsHttpCookieArgs()
        {
        }
    }
}
