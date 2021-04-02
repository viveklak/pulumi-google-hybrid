// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dlp.V2.Outputs
{

    [OutputType]
    public sealed class GooglePrivacyDlpV2CategoricalStatsHistogramBucketResponse
    {
        /// <summary>
        /// Total number of values in this bucket.
        /// </summary>
        public readonly string BucketSize;
        /// <summary>
        /// Total number of distinct values in this bucket.
        /// </summary>
        public readonly string BucketValueCount;
        /// <summary>
        /// Sample of value frequencies in this bucket. The total number of values returned per bucket is capped at 20.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2ValueFrequencyResponse> BucketValues;
        /// <summary>
        /// Lower bound on the value frequency of the values in this bucket.
        /// </summary>
        public readonly string ValueFrequencyLowerBound;
        /// <summary>
        /// Upper bound on the value frequency of the values in this bucket.
        /// </summary>
        public readonly string ValueFrequencyUpperBound;

        [OutputConstructor]
        private GooglePrivacyDlpV2CategoricalStatsHistogramBucketResponse(
            string bucketSize,

            string bucketValueCount,

            ImmutableArray<Outputs.GooglePrivacyDlpV2ValueFrequencyResponse> bucketValues,

            string valueFrequencyLowerBound,

            string valueFrequencyUpperBound)
        {
            BucketSize = bucketSize;
            BucketValueCount = bucketValueCount;
            BucketValues = bucketValues;
            ValueFrequencyLowerBound = valueFrequencyLowerBound;
            ValueFrequencyUpperBound = valueFrequencyUpperBound;
        }
    }
}