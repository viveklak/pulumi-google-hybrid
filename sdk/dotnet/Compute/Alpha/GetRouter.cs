// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetRouter
    {
        /// <summary>
        /// Returns the specified Router resource. Gets a list of available routers by making a list() request.
        /// </summary>
        public static Task<GetRouterResult> InvokeAsync(GetRouterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouterResult>("google-native:compute/alpha:getRouter", args ?? new GetRouterArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified Router resource. Gets a list of available routers by making a list() request.
        /// </summary>
        public static Output<GetRouterResult> Invoke(GetRouterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRouterResult>("google-native:compute/alpha:getRouter", args ?? new GetRouterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouterArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("router", required: true)]
        public string Router { get; set; } = null!;

        public GetRouterArgs()
        {
        }
    }

    public sealed class GetRouterInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("router", required: true)]
        public Input<string> Router { get; set; } = null!;

        public GetRouterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouterResult
    {
        /// <summary>
        /// BGP information specific to this router.
        /// </summary>
        public readonly Outputs.RouterBgpResponse Bgp;
        /// <summary>
        /// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterBgpPeerResponse> BgpPeers;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments). Not currently available publicly. 
        /// </summary>
        public readonly bool EncryptedInterconnectRouter;
        /// <summary>
        /// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterInterfaceResponse> Interfaces;
        /// <summary>
        /// Type of resource. Always compute#router for routers.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Keys used for MD5 authentication.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterMd5AuthenticationKeyResponse> Md5AuthenticationKeys;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of NAT services created in this router.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterNatResponse> Nats;
        /// <summary>
        /// URI of the network to which this router belongs.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// URI of the region where the router resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
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

        [OutputConstructor]
        private GetRouterResult(
            Outputs.RouterBgpResponse bgp,

            ImmutableArray<Outputs.RouterBgpPeerResponse> bgpPeers,

            string creationTimestamp,

            string description,

            bool encryptedInterconnectRouter,

            ImmutableArray<Outputs.RouterInterfaceResponse> interfaces,

            string kind,

            ImmutableArray<Outputs.RouterMd5AuthenticationKeyResponse> md5AuthenticationKeys,

            string name,

            ImmutableArray<Outputs.RouterNatResponse> nats,

            string network,

            string region,

            string selfLink,

            string selfLinkWithId)
        {
            Bgp = bgp;
            BgpPeers = bgpPeers;
            CreationTimestamp = creationTimestamp;
            Description = description;
            EncryptedInterconnectRouter = encryptedInterconnectRouter;
            Interfaces = interfaces;
            Kind = kind;
            Md5AuthenticationKeys = md5AuthenticationKeys;
            Name = name;
            Nats = nats;
            Network = network;
            Region = region;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
        }
    }
}
