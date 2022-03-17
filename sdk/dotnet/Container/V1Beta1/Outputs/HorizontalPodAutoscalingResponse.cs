// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// Configuration options for the horizontal pod autoscaling feature, which increases or decreases the number of replica pods a replication controller has based on the resource usage of the existing pods.
    /// </summary>
    [OutputType]
    public sealed class HorizontalPodAutoscalingResponse
    {
        /// <summary>
        /// Whether the Horizontal Pod Autoscaling feature is enabled in the cluster. When enabled, it ensures that metrics are collected into Stackdriver Monitoring.
        /// </summary>
        public readonly bool Disabled;

        [OutputConstructor]
        private HorizontalPodAutoscalingResponse(bool disabled)
        {
            Disabled = disabled;
        }
    }
}
