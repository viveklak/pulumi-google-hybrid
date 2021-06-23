// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class PathRuleResponse
    {
        /// <summary>
        /// The list of path patterns to match. Each must start with / and the only place a * is allowed is at the end following a /. The string fed to the path matcher does not include any text after the first ? or #, and those chars are not allowed here.
        /// </summary>
        public readonly ImmutableArray<string> Paths;
        /// <summary>
        /// In response to a matching path, the load balancer performs advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If routeAction specifies any weightedBackendServices, service must not be set. Conversely if service is set, routeAction cannot contain any weightedBackendServices. Only one of routeAction or urlRedirect must be set. UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within a pathRule's routeAction.
        /// </summary>
        public readonly Outputs.HttpRouteActionResponse RouteAction;
        /// <summary>
        /// The full or partial URL of the backend service resource to which traffic is directed if this rule is matched. If routeAction is additionally specified, advanced routing actions like URL Rewrites, etc. take effect prior to sending the request to the backend. However, if service is specified, routeAction cannot contain any weightedBackendService s. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of urlRedirect, service or routeAction.weightedBackendService must be set.
        /// </summary>
        public readonly string Service;
        /// <summary>
        /// When a path pattern is matched, the request is redirected to a URL specified by urlRedirect. If urlRedirect is specified, service or routeAction must not be set. Not supported when the URL map is bound to target gRPC proxy.
        /// </summary>
        public readonly Outputs.HttpRedirectActionResponse UrlRedirect;

        [OutputConstructor]
        private PathRuleResponse(
            ImmutableArray<string> paths,

            Outputs.HttpRouteActionResponse routeAction,

            string service,

            Outputs.HttpRedirectActionResponse urlRedirect)
        {
            Paths = paths;
            RouteAction = routeAction;
            Service = service;
            UrlRedirect = urlRedirect;
        }
    }
}
