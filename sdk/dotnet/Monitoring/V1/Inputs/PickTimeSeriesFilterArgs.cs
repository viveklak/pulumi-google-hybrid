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
    /// Describes a ranking-based time series filter. Each input time series is ranked with an aligner. The filter will allow up to num_time_series time series to pass through it, selecting them based on the relative ranking.For example, if ranking_method is METHOD_MEAN,direction is BOTTOM, and num_time_series is 3, then the 3 times series with the lowest mean values will pass through the filter.
    /// </summary>
    public sealed class PickTimeSeriesFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to use the ranking to select time series that pass through the filter.
        /// </summary>
        [Input("direction")]
        public Input<Pulumi.GoogleNative.Monitoring.V1.PickTimeSeriesFilterDirection>? Direction { get; set; }

        /// <summary>
        /// How many time series to allow to pass through the filter.
        /// </summary>
        [Input("numTimeSeries")]
        public Input<int>? NumTimeSeries { get; set; }

        /// <summary>
        /// ranking_method is applied to each time series independently to produce the value which will be used to compare the time series to other time series.
        /// </summary>
        [Input("rankingMethod")]
        public Input<Pulumi.GoogleNative.Monitoring.V1.PickTimeSeriesFilterRankingMethod>? RankingMethod { get; set; }

        public PickTimeSeriesFilterArgs()
        {
        }
    }
}
