// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// The spec for modifying the path before sending the request to the matched backend service.
    /// </summary>
    public sealed class UrlRewriteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Before forwarding the request to the selected service, the request's host header is replaced with contents of hostRewrite. The value must be from 1 to 255 characters.
        /// </summary>
        [Input("hostRewrite")]
        public Input<string>? HostRewrite { get; set; }

        /// <summary>
        /// Before forwarding the request to the selected backend service, the matching portion of the request's path is replaced by pathPrefixRewrite. The value must be from 1 to 1024 characters.
        /// </summary>
        [Input("pathPrefixRewrite")]
        public Input<string>? PathPrefixRewrite { get; set; }

        public UrlRewriteArgs()
        {
        }
    }
}
