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
    /// A chart that displays data on a 2D (X and Y axes) plane.
    /// </summary>
    public sealed class XyChartArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Display options for the chart.
        /// </summary>
        [Input("chartOptions")]
        public Input<Inputs.ChartOptionsArgs>? ChartOptions { get; set; }

        [Input("dataSets", required: true)]
        private InputList<Inputs.DataSetArgs>? _dataSets;

        /// <summary>
        /// The data displayed in this chart.
        /// </summary>
        public InputList<Inputs.DataSetArgs> DataSets
        {
            get => _dataSets ?? (_dataSets = new InputList<Inputs.DataSetArgs>());
            set => _dataSets = value;
        }

        [Input("thresholds")]
        private InputList<Inputs.ThresholdArgs>? _thresholds;

        /// <summary>
        /// Threshold lines drawn horizontally across the chart.
        /// </summary>
        public InputList<Inputs.ThresholdArgs> Thresholds
        {
            get => _thresholds ?? (_thresholds = new InputList<Inputs.ThresholdArgs>());
            set => _thresholds = value;
        }

        /// <summary>
        /// The duration used to display a comparison chart. A comparison chart simultaneously shows values from two similar-length time periods (e.g., week-over-week metrics). The duration must be positive, and it can only be applied to charts with data sets of LINE plot type.
        /// </summary>
        [Input("timeshiftDuration")]
        public Input<string>? TimeshiftDuration { get; set; }

        /// <summary>
        /// The properties applied to the X axis.
        /// </summary>
        [Input("xAxis")]
        public Input<Inputs.AxisArgs>? XAxis { get; set; }

        /// <summary>
        /// The properties applied to the Y axis.
        /// </summary>
        [Input("yAxis")]
        public Input<Inputs.AxisArgs>? YAxis { get; set; }

        public XyChartArgs()
        {
        }
    }
}
