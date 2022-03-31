// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Inputs
{

    /// <summary>
    /// Information required for a TCP Uptime check request.
    /// </summary>
    public sealed class TcpCheckArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The TCP port on the server against which to run the check. Will be combined with host (specified within the monitored_resource) to construct the full URL. Required.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public TcpCheckArgs()
        {
        }
    }
}
