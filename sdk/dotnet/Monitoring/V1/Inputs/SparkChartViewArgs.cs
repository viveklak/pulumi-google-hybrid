// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Inputs
{

    /// <summary>
    /// A sparkChart is a small chart suitable for inclusion in a table-cell or inline in text. This message contains the configuration for a sparkChart to show up on a Scorecard, showing recent trends of the scorecard's timeseries.
    /// </summary>
    public sealed class SparkChartViewArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The lower bound on data point frequency in the chart implemented by specifying the minimum alignment period to use in a time series query. For example, if the data is published once every 10 minutes it would not make sense to fetch and align data at one minute intervals. This field is optional and exists only as a hint.
        /// </summary>
        [Input("minAlignmentPeriod")]
        public Input<string>? MinAlignmentPeriod { get; set; }

        /// <summary>
        /// The type of sparkchart to show in this chartView.
        /// </summary>
        [Input("sparkChartType", required: true)]
        public Input<Pulumi.GoogleNative.Monitoring.V1.SparkChartViewSparkChartType> SparkChartType { get; set; } = null!;

        public SparkChartViewArgs()
        {
        }
    }
}
