// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Inputs
{

    /// <summary>
    /// A message representing a measurement.
    /// </summary>
    public sealed class GoogleCloudMlV1__MeasurementArgs : Pulumi.ResourceArgs
    {
        [Input("metrics")]
        private InputList<Inputs.GoogleCloudMlV1_Measurement_MetricArgs>? _metrics;

        /// <summary>
        /// Provides a list of metrics that act as inputs into the objective function.
        /// </summary>
        public InputList<Inputs.GoogleCloudMlV1_Measurement_MetricArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.GoogleCloudMlV1_Measurement_MetricArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// The number of steps a machine learning model has been trained for. Must be non-negative.
        /// </summary>
        [Input("stepCount")]
        public Input<string>? StepCount { get; set; }

        public GoogleCloudMlV1__MeasurementArgs()
        {
        }
    }
}
