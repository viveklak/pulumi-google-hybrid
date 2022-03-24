// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// Metric source to enable along with any optional metrics for this source that override the dataproc defaults
    /// </summary>
    [OutputType]
    public sealed class MetricResponse
    {
        /// <summary>
        /// Optional. Optional Metrics to override the dataproc default metrics configured for the metric source
        /// </summary>
        public readonly ImmutableArray<string> MetricOverrides;
        /// <summary>
        /// MetricSource that should be enabled
        /// </summary>
        public readonly string MetricSource;

        [OutputConstructor]
        private MetricResponse(
            ImmutableArray<string> metricOverrides,

            string metricSource)
        {
            MetricOverrides = metricOverrides;
            MetricSource = metricSource;
        }
    }
}