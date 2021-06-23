// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDatacatalogV1UsageSignalResponse
    {
        /// <summary>
        /// The end timestamp of the duration of usage statistics.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Usage statistics over each of the predefined time ranges. Supported time ranges are `{"24H", "7D", "30D"}`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> UsageWithinTimeRange;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1UsageSignalResponse(
            string updateTime,

            ImmutableDictionary<string, string> usageWithinTimeRange)
        {
            UpdateTime = updateTime;
            UsageWithinTimeRange = usageWithinTimeRange;
        }
    }
}
