// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetTargetTcpProxy
    {
        /// <summary>
        /// Returns the specified TargetTcpProxy resource. Gets a list of available target TCP proxies by making a list() request.
        /// </summary>
        public static Task<GetTargetTcpProxyResult> InvokeAsync(GetTargetTcpProxyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetTcpProxyResult>("google-native:compute/beta:getTargetTcpProxy", args ?? new GetTargetTcpProxyArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified TargetTcpProxy resource. Gets a list of available target TCP proxies by making a list() request.
        /// </summary>
        public static Output<GetTargetTcpProxyResult> Invoke(GetTargetTcpProxyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTargetTcpProxyResult>("google-native:compute/beta:getTargetTcpProxy", args ?? new GetTargetTcpProxyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetTcpProxyArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("targetTcpProxy", required: true)]
        public string TargetTcpProxy { get; set; } = null!;

        public GetTargetTcpProxyArgs()
        {
        }
    }

    public sealed class GetTargetTcpProxyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("targetTcpProxy", required: true)]
        public Input<string> TargetTcpProxy { get; set; } = null!;

        public GetTargetTcpProxyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTargetTcpProxyResult
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
        /// Type of the resource. Always compute#targetTcpProxy for target TCP proxies.
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
        /// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
        /// </summary>
        public readonly string ProxyHeader;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// URL to the BackendService resource.
        /// </summary>
        public readonly string Service;

        [OutputConstructor]
        private GetTargetTcpProxyResult(
            string creationTimestamp,

            string description,

            string kind,

            string name,

            bool proxyBind,

            string proxyHeader,

            string selfLink,

            string service)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Kind = kind;
            Name = name;
            ProxyBind = proxyBind;
            ProxyHeader = proxyHeader;
            SelfLink = selfLink;
            Service = service;
        }
    }
}
