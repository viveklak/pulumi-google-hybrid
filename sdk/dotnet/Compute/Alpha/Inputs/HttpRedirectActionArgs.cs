// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Specifies settings for an HTTP redirect.
    /// </summary>
    public sealed class HttpRedirectActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The host that is used in the redirect response instead of the one that was supplied in the request. The value must be from 1 to 255 characters.
        /// </summary>
        [Input("hostRedirect")]
        public Input<string>? HostRedirect { get; set; }

        /// <summary>
        /// If set to true, the URL scheme in the redirected request is set to HTTPS. If set to false, the URL scheme of the redirected request remains the same as that of the request. This must only be set for URL maps used in TargetHttpProxys. Setting this true for TargetHttpsProxy is not permitted. The default is set to false.
        /// </summary>
        [Input("httpsRedirect")]
        public Input<bool>? HttpsRedirect { get; set; }

        /// <summary>
        /// The path that is used in the redirect response instead of the one that was supplied in the request. pathRedirect cannot be supplied together with prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the original request is used for the redirect. The value must be from 1 to 1024 characters.
        /// </summary>
        [Input("pathRedirect")]
        public Input<string>? PathRedirect { get; set; }

        /// <summary>
        /// The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch, retaining the remaining portion of the URL before redirecting the request. prefixRedirect cannot be supplied together with pathRedirect. Supply one alone or neither. If neither is supplied, the path of the original request is used for the redirect. The value must be from 1 to 1024 characters.
        /// </summary>
        [Input("prefixRedirect")]
        public Input<string>? PrefixRedirect { get; set; }

        /// <summary>
        /// The HTTP Status code to use for this RedirectAction. Supported values are: - MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301. - FOUND, which corresponds to 302. - SEE_OTHER which corresponds to 303. - TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method is retained. - PERMANENT_REDIRECT, which corresponds to 308. In this case, the request method is retained. 
        /// </summary>
        [Input("redirectResponseCode")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.HttpRedirectActionRedirectResponseCode>? RedirectResponseCode { get; set; }

        /// <summary>
        /// If set to true, any accompanying query portion of the original URL is removed before redirecting the request. If set to false, the query portion of the original URL is retained. The default is set to false. 
        /// </summary>
        [Input("stripQuery")]
        public Input<bool>? StripQuery { get; set; }

        public HttpRedirectActionArgs()
        {
        }
    }
}
