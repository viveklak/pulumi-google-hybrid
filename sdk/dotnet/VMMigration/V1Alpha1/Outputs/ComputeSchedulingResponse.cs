// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1.Outputs
{

    /// <summary>
    /// Scheduling information for VM on maintenance/restart behaviour and node allocation in sole tenant nodes.
    /// </summary>
    [OutputType]
    public sealed class ComputeSchedulingResponse
    {
        public readonly bool AutomaticRestart;
        /// <summary>
        /// The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node. Ignored if no node_affinites are configured.
        /// </summary>
        public readonly int MinNodeCpus;
        /// <summary>
        /// A set of node affinity and anti-affinity configurations for sole tenant nodes.
        /// </summary>
        public readonly ImmutableArray<Outputs.SchedulingNodeAffinityResponse> NodeAffinities;
        /// <summary>
        /// How the instance should behave when the host machine undergoes maintenance that may temporarily impact instance performance.
        /// </summary>
        public readonly string OnHostMaintenance;
        /// <summary>
        /// Whether the Instance should be automatically restarted whenever it is terminated by Compute Engine (not terminated by user). This configuration is identical to `automaticRestart` field in Compute Engine create instance under scheduling. It was changed to an enum (instead of a boolean) to match the default value in Compute Engine which is automatic restart.
        /// </summary>
        public readonly string RestartType;

        [OutputConstructor]
        private ComputeSchedulingResponse(
            bool automaticRestart,

            int minNodeCpus,

            ImmutableArray<Outputs.SchedulingNodeAffinityResponse> nodeAffinities,

            string onHostMaintenance,

            string restartType)
        {
            AutomaticRestart = automaticRestart;
            MinNodeCpus = minNodeCpus;
            NodeAffinities = nodeAffinities;
            OnHostMaintenance = onHostMaintenance;
            RestartType = restartType;
        }
    }
}
