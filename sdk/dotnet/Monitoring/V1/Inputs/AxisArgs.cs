// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Monitoring.V1.Inputs
{

    /// <summary>
    /// A chart axis.
    /// </summary>
    public sealed class AxisArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The label of the axis.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The axis scale. By default, a linear scale is used.
        /// </summary>
        [Input("scale")]
        public Input<string>? Scale { get; set; }

        public AxisArgs()
        {
        }
    }
}