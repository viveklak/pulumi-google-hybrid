// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Outputs
{

    [OutputType]
    public sealed class BiEngineStatisticsResponse
    {
        /// <summary>
        /// Specifies which mode of BI Engine acceleration was performed (if any).
        /// </summary>
        public readonly string BiEngineMode;
        /// <summary>
        /// In case of DISABLED or PARTIAL bi_engine_mode, these contain the explanatory reasons as to why BI Engine could not accelerate. In case the full query was accelerated, this field is not populated.
        /// </summary>
        public readonly ImmutableArray<Outputs.BiEngineReasonResponse> BiEngineReasons;

        [OutputConstructor]
        private BiEngineStatisticsResponse(
            string biEngineMode,

            ImmutableArray<Outputs.BiEngineReasonResponse> biEngineReasons)
        {
            BiEngineMode = biEngineMode;
            BiEngineReasons = biEngineReasons;
        }
    }
}
