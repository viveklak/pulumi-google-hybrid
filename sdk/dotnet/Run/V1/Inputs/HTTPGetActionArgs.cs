// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Inputs
{

    /// <summary>
    /// Not supported by Cloud Run HTTPGetAction describes an action based on HTTP Get requests.
    /// </summary>
    public sealed class HTTPGetActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("httpHeaders")]
        private InputList<Inputs.HTTPHeaderArgs>? _httpHeaders;

        /// <summary>
        /// (Optional) Custom headers to set in the request. HTTP allows repeated headers.
        /// </summary>
        public InputList<Inputs.HTTPHeaderArgs> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputList<Inputs.HTTPHeaderArgs>());
            set => _httpHeaders = value;
        }

        /// <summary>
        /// (Optional) Path to access on the HTTP server.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// (Optional) Scheme to use for connecting to the host. Defaults to HTTP.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        public HTTPGetActionArgs()
        {
        }
    }
}
