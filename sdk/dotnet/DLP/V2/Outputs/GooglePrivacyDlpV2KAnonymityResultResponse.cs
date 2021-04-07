// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.DLP.V2.Outputs
{

    [OutputType]
    public sealed class GooglePrivacyDlpV2KAnonymityResultResponse
    {
        /// <summary>
        /// Histogram of k-anonymity equivalence classes.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2KAnonymityHistogramBucketResponse> EquivalenceClassHistogramBuckets;

        [OutputConstructor]
        private GooglePrivacyDlpV2KAnonymityResultResponse(ImmutableArray<Outputs.GooglePrivacyDlpV2KAnonymityHistogramBucketResponse> equivalenceClassHistogramBuckets)
        {
            EquivalenceClassHistogramBuckets = equivalenceClassHistogramBuckets;
        }
    }
}