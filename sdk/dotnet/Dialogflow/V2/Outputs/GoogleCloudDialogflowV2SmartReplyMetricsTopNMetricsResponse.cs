// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Outputs
{

    /// <summary>
    /// Evaluation metrics when retrieving `n` smart replies with the model.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2SmartReplyMetricsTopNMetricsResponse
    {
        /// <summary>
        /// Number of retrieved smart replies. For example, when `n` is 3, this evaluation contains metrics for when Dialogflow retrieves 3 smart replies with the model.
        /// </summary>
        public readonly int N;
        /// <summary>
        /// Defined as `number of queries whose top n smart replies have at least one similar (token match similarity above the defined threshold) reply as the real reply` divided by `number of queries with at least one smart reply`. Value ranges from 0.0 to 1.0 inclusive.
        /// </summary>
        public readonly double Recall;

        [OutputConstructor]
        private GoogleCloudDialogflowV2SmartReplyMetricsTopNMetricsResponse(
            int n,

            double recall)
        {
            N = n;
            Recall = recall;
        }
    }
}
