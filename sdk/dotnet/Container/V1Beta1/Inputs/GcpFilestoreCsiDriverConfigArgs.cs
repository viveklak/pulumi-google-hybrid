// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// Configuration for the GCP Filestore CSI driver.
    /// </summary>
    public sealed class GcpFilestoreCsiDriverConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the GCP Filestore CSI driver is enabled for this cluster.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public GcpFilestoreCsiDriverConfigArgs()
        {
        }
    }
}
