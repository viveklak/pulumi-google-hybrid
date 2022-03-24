// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1.Outputs
{

    /// <summary>
    /// The specification for modifying the URL of the request, prior to forwarding the request to the destination.
    /// </summary>
    [OutputType]
    public sealed class HttpRouteURLRewriteResponse
    {
        /// <summary>
        /// Prior to forwarding the request to the selected destination, the requests host header is replaced by this value.
        /// </summary>
        public readonly string HostRewrite;
        /// <summary>
        /// Prior to forwarding the request to the selected destination, the matching portion of the requests path is replaced by this value.
        /// </summary>
        public readonly string PathPrefixRewrite;

        [OutputConstructor]
        private HttpRouteURLRewriteResponse(
            string hostRewrite,

            string pathPrefixRewrite)
        {
            HostRewrite = hostRewrite;
            PathPrefixRewrite = pathPrefixRewrite;
        }
    }
}