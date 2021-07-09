// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Outputs
{

    [OutputType]
    public sealed class SparkChartViewResponse
    {
        /// <summary>
        /// The lower bound on data point frequency in the chart implemented by specifying the minimum alignment period to use in a time series query. For example, if the data is published once every 10 minutes it would not make sense to fetch and align data at one minute intervals. This field is optional and exists only as a hint.
        /// </summary>
        public readonly string MinAlignmentPeriod;
        /// <summary>
        /// The type of sparkchart to show in this chartView.
        /// </summary>
        public readonly string SparkChartType;

        [OutputConstructor]
        private SparkChartViewResponse(
            string minAlignmentPeriod,

            string sparkChartType)
        {
            MinAlignmentPeriod = minAlignmentPeriod;
            SparkChartType = sparkChartType;
        }
    }
}
