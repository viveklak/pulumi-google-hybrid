// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetTargetHttpProxy
    {
        /// <summary>
        /// Returns the specified TargetHttpProxy resource. Gets a list of available target HTTP proxies by making a list() request.
        /// </summary>
        public static Task<GetTargetHttpProxyResult> InvokeAsync(GetTargetHttpProxyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetHttpProxyResult>("google-native:compute/alpha:getTargetHttpProxy", args ?? new GetTargetHttpProxyArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified TargetHttpProxy resource. Gets a list of available target HTTP proxies by making a list() request.
        /// </summary>
        public static Output<GetTargetHttpProxyResult> Invoke(GetTargetHttpProxyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTargetHttpProxyResult>("google-native:compute/alpha:getTargetHttpProxy", args ?? new GetTargetHttpProxyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetHttpProxyArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("targetHttpProxy", required: true)]
        public string TargetHttpProxy { get; set; } = null!;

        public GetTargetHttpProxyArgs()
        {
        }
    }

    public sealed class GetTargetHttpProxyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("targetHttpProxy", required: true)]
        public Input<string> TargetHttpProxy { get; set; } = null!;

        public GetTargetHttpProxyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTargetHttpProxyResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// URLs to networkservices.HttpFilter resources enabled for xDS clients using this configuration. For example, https://networkservices.googleapis.com/v1alpha1/projects/project/locations/ locationhttpFilters/httpFilter Only filters that handle outbound connection and stream events may be specified. These filters work in conjunction with a default set of HTTP filters that may already be configured by Traffic Director. Traffic Director will determine the final location of these filters within xDS configuration based on the name of the HTTP filter. If Traffic Director positions multiple filters at the same location, those filters will be in the same order as specified in this list. httpFilters only applies for loadbalancers with loadBalancingScheme set to INTERNAL_SELF_MANAGED. See ForwardingRule for more details.
        /// </summary>
        public readonly ImmutableArray<string> HttpFilters;
        /// <summary>
        /// Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
        /// </summary>
        public readonly bool ProxyBind;
        /// <summary>
        /// URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
        /// </summary>
        public readonly string UrlMap;

        [OutputConstructor]
        private GetTargetHttpProxyResult(
            string creationTimestamp,

            string description,

            string fingerprint,

            ImmutableArray<string> httpFilters,

            string kind,

            string name,

            bool proxyBind,

            string region,

            string selfLink,

            string selfLinkWithId,

            string urlMap)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Fingerprint = fingerprint;
            HttpFilters = httpFilters;
            Kind = kind;
            Name = name;
            ProxyBind = proxyBind;
            Region = region;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            UrlMap = urlMap;
        }
    }
}
