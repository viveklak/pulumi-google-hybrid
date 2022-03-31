// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha.Outputs
{

    /// <summary>
    /// Used for time-dependent conditions. Example: Want to have rule applied for week long sale.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2alphaConditionTimeRangeResponse
    {
        /// <summary>
        /// End of time range. Range is inclusive.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Start of time range. Range is inclusive.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private GoogleCloudRetailV2alphaConditionTimeRangeResponse(
            string endTime,

            string startTime)
        {
            EndTime = endTime;
            StartTime = startTime;
        }
    }
}
