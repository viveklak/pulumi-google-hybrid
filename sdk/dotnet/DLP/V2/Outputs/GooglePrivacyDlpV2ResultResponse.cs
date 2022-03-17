// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// All result fields mentioned below are updated while the job is processing.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2ResultResponse
    {
        /// <summary>
        /// Statistics related to the processing of hybrid inspect.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2HybridInspectStatisticsResponse HybridStats;
        /// <summary>
        /// Statistics of how many instances of each info type were found during inspect job.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2InfoTypeStatsResponse> InfoTypeStats;
        /// <summary>
        /// Total size in bytes that were processed.
        /// </summary>
        public readonly string ProcessedBytes;
        /// <summary>
        /// Estimate of the number of bytes to process.
        /// </summary>
        public readonly string TotalEstimatedBytes;

        [OutputConstructor]
        private GooglePrivacyDlpV2ResultResponse(
            Outputs.GooglePrivacyDlpV2HybridInspectStatisticsResponse hybridStats,

            ImmutableArray<Outputs.GooglePrivacyDlpV2InfoTypeStatsResponse> infoTypeStats,

            string processedBytes,

            string totalEstimatedBytes)
        {
            HybridStats = hybridStats;
            InfoTypeStats = infoTypeStats;
            ProcessedBytes = processedBytes;
            TotalEstimatedBytes = totalEstimatedBytes;
        }
    }
}
