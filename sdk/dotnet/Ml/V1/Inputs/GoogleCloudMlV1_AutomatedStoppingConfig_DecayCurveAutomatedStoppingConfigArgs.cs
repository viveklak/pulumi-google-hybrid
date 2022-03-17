// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Inputs
{

    public sealed class GoogleCloudMlV1_AutomatedStoppingConfig_DecayCurveAutomatedStoppingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, measurement.elapsed_time is used as the x-axis of each Trials Decay Curve. Otherwise, Measurement.steps will be used as the x-axis.
        /// </summary>
        [Input("useElapsedTime")]
        public Input<bool>? UseElapsedTime { get; set; }

        public GoogleCloudMlV1_AutomatedStoppingConfig_DecayCurveAutomatedStoppingConfigArgs()
        {
        }
    }
}
