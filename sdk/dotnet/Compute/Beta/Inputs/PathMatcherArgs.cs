// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// A matcher for the path portion of the URL. The BackendService from the longest-matched rule will serve the URL. If no rule was matched, the default service will be used.
    /// </summary>
    public sealed class PathMatcherArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within a pathMatcher's defaultRouteAction.
        /// </summary>
        [Input("defaultRouteAction")]
        public Input<Inputs.HttpRouteActionArgs>? DefaultRouteAction { get; set; }

        /// <summary>
        /// The full or partial URL to the BackendService resource. This will be used if none of the pathRules or routeRules defined by this PathMatcher are matched. For example, the following are all valid URLs to a BackendService resource: - https://www.googleapis.com/compute/v1/projects/project /global/backendServices/backendService - compute/v1/projects/project/global/backendServices/backendService - global/backendServices/backendService If defaultRouteAction is additionally specified, advanced routing actions like URL Rewrites, etc. take effect prior to sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if defaultRouteAction specifies any weightedBackendServices, defaultService must not be specified. Only one of defaultService, defaultUrlRedirect or defaultRouteAction.weightedBackendService must be set. Authorization requires one or more of the following Google IAM permissions on the specified resource default_service: - compute.backendBuckets.use - compute.backendServices.use 
        /// </summary>
        [Input("defaultService")]
        public Input<string>? DefaultService { get; set; }

        /// <summary>
        /// When none of the specified pathRules or routeRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to target gRPC proxy.
        /// </summary>
        [Input("defaultUrlRedirect")]
        public Input<Inputs.HttpRedirectActionArgs>? DefaultUrlRedirect { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies changes to request and response headers that need to take effect for the selected backendService. HeaderAction specified here are applied after the matching HttpRouteRule HeaderAction and before the HeaderAction in the UrlMap Note that headerAction is not supported for Loadbalancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("headerAction")]
        public Input<Inputs.HttpHeaderActionArgs>? HeaderAction { get; set; }

        /// <summary>
        /// The name to which this PathMatcher is referred by the HostRule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pathRules")]
        private InputList<Inputs.PathRuleArgs>? _pathRules;

        /// <summary>
        /// The list of path rules. Use this list instead of routeRules when routing based on simple path matching is all that's required. The order by which path rules are specified does not matter. Matches are always done on the longest-path-first basis. For example: a pathRule with a path /a/b/c/* will match before /a/b/* irrespective of the order in which those paths appear in this list. Within a given pathMatcher, only one of pathRules or routeRules must be set.
        /// </summary>
        public InputList<Inputs.PathRuleArgs> PathRules
        {
            get => _pathRules ?? (_pathRules = new InputList<Inputs.PathRuleArgs>());
            set => _pathRules = value;
        }

        [Input("routeRules")]
        private InputList<Inputs.HttpRouteRuleArgs>? _routeRules;

        /// <summary>
        /// The list of HTTP route rules. Use this list instead of pathRules when advanced route matching and routing actions are desired. routeRules are evaluated in order of priority, from the lowest to highest number. Within a given pathMatcher, you can set only one of pathRules or routeRules.
        /// </summary>
        public InputList<Inputs.HttpRouteRuleArgs> RouteRules
        {
            get => _routeRules ?? (_routeRules = new InputList<Inputs.HttpRouteRuleArgs>());
            set => _routeRules = value;
        }

        public PathMatcherArgs()
        {
        }
    }
}
