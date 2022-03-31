// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2.Outputs
{

    /// <summary>
    /// Contains cluster daemon metrics, such as HDFS and YARN stats.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
    /// </summary>
    [OutputType]
    public sealed class ClusterMetricsResponse
    {
        /// <summary>
        /// The HDFS metrics.
        /// </summary>
        public readonly ImmutableDictionary<string, string> HdfsMetrics;
        /// <summary>
        /// The YARN metrics.
        /// </summary>
        public readonly ImmutableDictionary<string, string> YarnMetrics;

        [OutputConstructor]
        private ClusterMetricsResponse(
            ImmutableDictionary<string, string> hdfsMetrics,

            ImmutableDictionary<string, string> yarnMetrics)
        {
            HdfsMetrics = hdfsMetrics;
            YarnMetrics = yarnMetrics;
        }
    }
}
