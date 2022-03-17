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
    /// For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview). On failover or failback, this field indicates whether connection draining will be honored. Google Cloud has a fixed connection draining timeout of 10 minutes. A setting of true terminates existing TCP connections to the active pool during failover and failback, immediately draining traffic. A setting of false allows existing TCP connections to persist, even on VMs no longer in the active pool, for up to the duration of the connection draining timeout (10 minutes).
    /// </summary>
    [OutputType]
    public sealed class BackendServiceFailoverPolicyResponse
    {
        /// <summary>
        /// This can be set to true only if the protocol is TCP. The default is false.
        /// </summary>
        public readonly bool DisableConnectionDrainOnFailover;
        /// <summary>
        /// If set to true, connections to the load balancer are dropped when all primary and all backup backend VMs are unhealthy.If set to false, connections are distributed among all primary VMs when all primary and all backup backend VMs are unhealthy. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview). The default is false.
        /// </summary>
        public readonly bool DropTrafficIfUnhealthy;
        /// <summary>
        /// The value of the field must be in the range [0, 1]. If the value is 0, the load balancer performs a failover when the number of healthy primary VMs equals zero. For all other values, the load balancer performs a failover when the total number of healthy primary VMs is less than this ratio. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview).
        /// </summary>
        public readonly double FailoverRatio;

        [OutputConstructor]
        private BackendServiceFailoverPolicyResponse(
            bool disableConnectionDrainOnFailover,

            bool dropTrafficIfUnhealthy,

            double failoverRatio)
        {
            DisableConnectionDrainOnFailover = disableConnectionDrainOnFailover;
            DropTrafficIfUnhealthy = dropTrafficIfUnhealthy;
            FailoverRatio = failoverRatio;
        }
    }
}
