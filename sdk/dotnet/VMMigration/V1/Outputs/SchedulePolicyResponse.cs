// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1.Outputs
{

    /// <summary>
    /// A policy for scheduling replications.
    /// </summary>
    [OutputType]
    public sealed class SchedulePolicyResponse
    {
        /// <summary>
        /// The idle duration between replication stages.
        /// </summary>
        public readonly string IdleDuration;
        /// <summary>
        /// A flag to indicate whether to skip OS adaptation during the replication sync. OS adaptation is a process where the VM's operating system undergoes changes and adaptations to fully function on Compute Engine.
        /// </summary>
        public readonly bool SkipOsAdaptation;

        [OutputConstructor]
        private SchedulePolicyResponse(
            string idleDuration,

            bool skipOsAdaptation)
        {
            IdleDuration = idleDuration;
            SkipOsAdaptation = skipOsAdaptation;
        }
    }
}
