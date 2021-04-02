// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Containeranalysis.V1beta1.Outputs
{

    [OutputType]
    public sealed class DiscoveredResponse
    {
        /// <summary>
        /// The status of discovery for the resource.
        /// </summary>
        public readonly string AnalysisStatus;
        /// <summary>
        /// When an error is encountered this will contain a LocalizedMessage under details to show to the user. The LocalizedMessage is output only and populated by the API.
        /// </summary>
        public readonly Outputs.StatusResponse AnalysisStatusError;
        /// <summary>
        /// Whether the resource is continuously analyzed.
        /// </summary>
        public readonly string ContinuousAnalysis;
        /// <summary>
        /// The last time continuous analysis was done for this resource. Deprecated, do not use.
        /// </summary>
        public readonly string LastAnalysisTime;

        [OutputConstructor]
        private DiscoveredResponse(
            string analysisStatus,

            Outputs.StatusResponse analysisStatusError,

            string continuousAnalysis,

            string lastAnalysisTime)
        {
            AnalysisStatus = analysisStatus;
            AnalysisStatusError = analysisStatusError;
            ContinuousAnalysis = continuousAnalysis;
            LastAnalysisTime = lastAnalysisTime;
        }
    }
}