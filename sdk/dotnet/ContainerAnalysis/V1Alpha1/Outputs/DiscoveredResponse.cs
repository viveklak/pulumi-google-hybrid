// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// Provides information about the scan status of a discovered resource.
    /// </summary>
    [OutputType]
    public sealed class DiscoveredResponse
    {
        /// <summary>
        /// The status of discovery for the resource.
        /// </summary>
        public readonly string AnalysisStatus;
        /// <summary>
        /// When an error is encountered this will contain a LocalizedMessage under details to show to the user. The LocalizedMessage output only and populated by the API.
        /// </summary>
        public readonly Outputs.StatusResponse AnalysisStatusError;
        /// <summary>
        /// The time occurrences related to this discovery occurrence were archived.
        /// </summary>
        public readonly string ArchiveTime;
        /// <summary>
        /// Whether the resource is continuously analyzed.
        /// </summary>
        public readonly string ContinuousAnalysis;
        /// <summary>
        /// The CPE of the resource being scanned.
        /// </summary>
        public readonly string Cpe;
        /// <summary>
        /// The last time this resource was scanned.
        /// </summary>
        public readonly string LastScanTime;
        /// <summary>
        /// An operation that indicates the status of the current scan. This field is deprecated, do not use.
        /// </summary>
        public readonly Outputs.OperationResponse Operation;

        [OutputConstructor]
        private DiscoveredResponse(
            string analysisStatus,

            Outputs.StatusResponse analysisStatusError,

            string archiveTime,

            string continuousAnalysis,

            string cpe,

            string lastScanTime,

            Outputs.OperationResponse operation)
        {
            AnalysisStatus = analysisStatus;
            AnalysisStatusError = analysisStatusError;
            ArchiveTime = archiveTime;
            ContinuousAnalysis = continuousAnalysis;
            Cpe = cpe;
            LastScanTime = lastScanTime;
            Operation = operation;
        }
    }
}
