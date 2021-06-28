// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetTargetHttpProxy
    {
        /// <summary>
        /// Returns the specified TargetHttpProxy resource. Gets a list of available target HTTP proxies by making a list() request.
        /// </summary>
        public static Task<GetTargetHttpProxyResult> InvokeAsync(GetTargetHttpProxyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetHttpProxyResult>("google-native:compute/v1:getTargetHttpProxy", args ?? new GetTargetHttpProxyArgs(), options.WithVersion());
    }


    public sealed class GetTargetHttpProxyArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("targetHttpProxy", required: true)]
        public string TargetHttpProxy { get; set; } = null!;

        public GetTargetHttpProxyArgs()
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
        /// Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// 
        /// When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them.
        /// 
        /// The default is false.
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
        /// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
        /// </summary>
        public readonly string UrlMap;

        [OutputConstructor]
        private GetTargetHttpProxyResult(
            string creationTimestamp,

            string description,

            string fingerprint,

            string kind,

            string name,

            bool proxyBind,

            string region,

            string selfLink,

            string urlMap)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Fingerprint = fingerprint;
            Kind = kind;
            Name = name;
            ProxyBind = proxyBind;
            Region = region;
            SelfLink = selfLink;
            UrlMap = urlMap;
        }
    }
}
