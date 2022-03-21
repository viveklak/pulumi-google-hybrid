// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1.Inputs
{

    /// <summary>
    /// Target scaling by disk usage. Only applicable in the App Engine flexible environment.
    /// </summary>
    public sealed class DiskUtilizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target bytes read per second.
        /// </summary>
        [Input("targetReadBytesPerSecond")]
        public Input<int>? TargetReadBytesPerSecond { get; set; }

        /// <summary>
        /// Target ops read per seconds.
        /// </summary>
        [Input("targetReadOpsPerSecond")]
        public Input<int>? TargetReadOpsPerSecond { get; set; }

        /// <summary>
        /// Target bytes written per second.
        /// </summary>
        [Input("targetWriteBytesPerSecond")]
        public Input<int>? TargetWriteBytesPerSecond { get; set; }

        /// <summary>
        /// Target ops written per second.
        /// </summary>
        [Input("targetWriteOpsPerSecond")]
        public Input<int>? TargetWriteOpsPerSecond { get; set; }

        public DiskUtilizationArgs()
        {
        }
    }
}
