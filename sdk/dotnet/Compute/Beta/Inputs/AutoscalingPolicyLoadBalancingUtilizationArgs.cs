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
    /// Configuration parameters of autoscaling based on load balancing.
    /// </summary>
    public sealed class AutoscalingPolicyLoadBalancingUtilizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fraction of backend capacity utilization (set in HTTP(S) load balancing configuration) that the autoscaler maintains. Must be a positive float value. If not defined, the default is 0.8.
        /// </summary>
        [Input("utilizationTarget")]
        public Input<double>? UtilizationTarget { get; set; }

        public AutoscalingPolicyLoadBalancingUtilizationArgs()
        {
        }
    }
}
