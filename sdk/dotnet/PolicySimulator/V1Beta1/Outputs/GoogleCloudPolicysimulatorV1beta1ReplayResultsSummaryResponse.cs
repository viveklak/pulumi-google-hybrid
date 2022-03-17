// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.PolicySimulator.V1Beta1.Outputs
{

    /// <summary>
    /// Summary statistics about the replayed log entries.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse
    {
        /// <summary>
        /// The number of replayed log entries with a difference between baseline and simulated policies.
        /// </summary>
        public readonly int DifferenceCount;
        /// <summary>
        /// The number of log entries that could not be replayed.
        /// </summary>
        public readonly int ErrorCount;
        /// <summary>
        /// The total number of log entries replayed.
        /// </summary>
        public readonly int LogCount;
        /// <summary>
        /// The date of the newest log entry replayed.
        /// </summary>
        public readonly Outputs.GoogleTypeDateResponse NewestDate;
        /// <summary>
        /// The date of the oldest log entry replayed.
        /// </summary>
        public readonly Outputs.GoogleTypeDateResponse OldestDate;
        /// <summary>
        /// The number of replayed log entries with no difference between baseline and simulated policies.
        /// </summary>
        public readonly int UnchangedCount;

        [OutputConstructor]
        private GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse(
            int differenceCount,

            int errorCount,

            int logCount,

            Outputs.GoogleTypeDateResponse newestDate,

            Outputs.GoogleTypeDateResponse oldestDate,

            int unchangedCount)
        {
            DifferenceCount = differenceCount;
            ErrorCount = errorCount;
            LogCount = logCount;
            NewestDate = newestDate;
            OldestDate = oldestDate;
            UnchangedCount = unchangedCount;
        }
    }
}
