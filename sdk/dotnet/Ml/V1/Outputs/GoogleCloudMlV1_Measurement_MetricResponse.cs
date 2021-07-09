// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudMlV1_Measurement_MetricResponse
    {
        /// <summary>
        /// Metric name.
        /// </summary>
        public readonly string Metric;
        /// <summary>
        /// The value for this metric.
        /// </summary>
        public readonly double Value;

        [OutputConstructor]
        private GoogleCloudMlV1_Measurement_MetricResponse(
            string metric,

            double value)
        {
            Metric = metric;
            Value = value;
        }
    }
}
