// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// A matcher for the path portion of the URL. The BackendService from the longest-matched rule will serve the URL. If no rule was matched, the default service is used.
    /// </summary>
    [OutputType]
    public sealed class PathMatcherResponse
    {
        /// <summary>
        /// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within a path matcher's defaultRouteAction.
        /// </summary>
        public readonly Outputs.HttpRouteActionResponse DefaultRouteAction;
        /// <summary>
        /// The full or partial URL to the BackendService resource. This URL is used if none of the pathRules or routeRules defined by this PathMatcher are matched. For example, the following are all valid URLs to a BackendService resource: - https://www.googleapis.com/compute/v1/projects/project /global/backendServices/backendService - compute/v1/projects/project/global/backendServices/backendService - global/backendServices/backendService If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if defaultRouteAction specifies any weightedBackendServices, defaultService must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. Authorization requires one or more of the following Google IAM permissions on the specified resource default_service: - compute.backendBuckets.use - compute.backendServices.use 
        /// </summary>
        public readonly string DefaultService;
        /// <summary>
        /// When none of the specified pathRules or routeRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
        /// </summary>
        public readonly Outputs.HttpRedirectActionResponse DefaultUrlRedirect;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Specifies changes to request and response headers that need to take effect for the selected backend service. HeaderAction specified here are applied after the matching HttpRouteRule HeaderAction and before the HeaderAction in the UrlMap HeaderAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        public readonly Outputs.HttpHeaderActionResponse HeaderAction;
        /// <summary>
        /// The name to which this PathMatcher is referred by the HostRule.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of path rules. Use this list instead of routeRules when routing based on simple path matching is all that's required. The order by which path rules are specified does not matter. Matches are always done on the longest-path-first basis. For example: a pathRule with a path /a/b/c/* will match before /a/b/* irrespective of the order in which those paths appear in this list. Within a given pathMatcher, only one of pathRules or routeRules must be set.
        /// </summary>
        public readonly ImmutableArray<Outputs.PathRuleResponse> PathRules;
        /// <summary>
        /// The list of HTTP route rules. Use this list instead of pathRules when advanced route matching and routing actions are desired. routeRules are evaluated in order of priority, from the lowest to highest number. Within a given pathMatcher, you can set only one of pathRules or routeRules.
        /// </summary>
        public readonly ImmutableArray<Outputs.HttpRouteRuleResponse> RouteRules;

        [OutputConstructor]
        private PathMatcherResponse(
            Outputs.HttpRouteActionResponse defaultRouteAction,

            string defaultService,

            Outputs.HttpRedirectActionResponse defaultUrlRedirect,

            string description,

            Outputs.HttpHeaderActionResponse headerAction,

            string name,

            ImmutableArray<Outputs.PathRuleResponse> pathRules,

            ImmutableArray<Outputs.HttpRouteRuleResponse> routeRules)
        {
            DefaultRouteAction = defaultRouteAction;
            DefaultService = defaultService;
            DefaultUrlRedirect = defaultUrlRedirect;
            Description = description;
            HeaderAction = headerAction;
            Name = name;
            PathRules = pathRules;
            RouteRules = routeRules;
        }
    }
}
