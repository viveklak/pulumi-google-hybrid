// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2.Outputs
{

    /// <summary>
    /// The Autoscaling targets for a Cluster. These determine the recommended nodes.
    /// </summary>
    [OutputType]
    public sealed class AutoscalingTargetsResponse
    {
        /// <summary>
        /// The cpu utilization that the Autoscaler should be trying to achieve. This number is on a scale from 0 (no utilization) to 100 (total utilization), and is limited between 10 and 80.
        /// </summary>
        public readonly int CpuUtilizationPercent;

        [OutputConstructor]
        private AutoscalingTargetsResponse(int cpuUtilizationPercent)
        {
            CpuUtilizationPercent = cpuUtilizationPercent;
        }
    }
}