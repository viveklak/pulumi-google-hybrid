// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1Beta1.Inputs
{

    /// <summary>
    /// Specification of traffic destination attributes.
    /// </summary>
    public sealed class DestinationArgs : Pulumi.ResourceArgs
    {
        [Input("hosts", required: true)]
        private InputList<string>? _hosts;

        /// <summary>
        /// List of host names to match. Matched against the ":authority" header in http requests. At least one host should match. Each host can be an exact match, or a prefix match (example "mydomain.*") or a suffix match (example // *.myorg.com") or a presence(any) match "*".
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// Optional. Match against key:value pair in http header. Provides a flexible match based on HTTP headers, for potentially advanced use cases. At least one header should match. Avoid using header matches to make authorization decisions unless there is a strong guarantee that requests arrive through a trusted client or proxy.
        /// </summary>
        [Input("httpHeaderMatch")]
        public Input<Inputs.HttpHeaderMatchArgs>? HttpHeaderMatch { get; set; }

        [Input("methods")]
        private InputList<string>? _methods;

        /// <summary>
        /// Optional. A list of HTTP methods to match. At least one method should match. Should not be set for gRPC services.
        /// </summary>
        public InputList<string> Methods
        {
            get => _methods ?? (_methods = new InputList<string>());
            set => _methods = value;
        }

        [Input("ports", required: true)]
        private InputList<int>? _ports;

        /// <summary>
        /// List of destination ports to match. At least one port should match.
        /// </summary>
        public InputList<int> Ports
        {
            get => _ports ?? (_ports = new InputList<int>());
            set => _ports = value;
        }

        public DestinationArgs()
        {
        }
    }
}
