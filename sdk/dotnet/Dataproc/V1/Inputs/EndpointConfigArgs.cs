// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// Endpoint config for this cluster
    /// </summary>
    public sealed class EndpointConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. If true, enable http access to specific ports on the cluster from external sources. Defaults to false.
        /// </summary>
        [Input("enableHttpPortAccess")]
        public Input<bool>? EnableHttpPortAccess { get; set; }

        public EndpointConfigArgs()
        {
        }
    }
}
