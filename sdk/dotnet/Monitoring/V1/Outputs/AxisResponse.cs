// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Outputs
{

    /// <summary>
    /// A chart axis.
    /// </summary>
    [OutputType]
    public sealed class AxisResponse
    {
        /// <summary>
        /// The label of the axis.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// The axis scale. By default, a linear scale is used.
        /// </summary>
        public readonly string Scale;

        [OutputConstructor]
        private AxisResponse(
            string label,

            string scale)
        {
            Label = label;
            Scale = scale;
        }
    }
}
