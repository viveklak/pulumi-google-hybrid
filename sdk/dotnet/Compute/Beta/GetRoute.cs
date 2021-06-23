// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetRoute
    {
        /// <summary>
        /// Returns the specified Route resource. Gets a list of available routes by making a list() request.
        /// </summary>
        public static Task<GetRouteResult> InvokeAsync(GetRouteArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteResult>("google-native:compute/beta:getRoute", args ?? new GetRouteArgs(), options.WithVersion());
    }


    public sealed class GetRouteArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("route", required: true)]
        public string Route { get; set; } = null!;

        public GetRouteArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouteResult
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
        /// </summary>
        public readonly string DestRange;
        /// <summary>
        /// [Output Only] Type of this resource. Always compute#routes for Route resources.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Fully-qualified URL of the network that this route applies to.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL: projects/ project/global/gateways/default-internet-gateway
        /// </summary>
        public readonly string NextHopGateway;
        /// <summary>
        /// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs: - 10.128.0.56 - https://www.googleapis.com/compute/v1/projects/project/regions/region /forwardingRules/forwardingRule - regions/region/forwardingRules/forwardingRule 
        /// </summary>
        public readonly string NextHopIlb;
        /// <summary>
        /// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
        /// </summary>
        public readonly string NextHopInstance;
        /// <summary>
        /// [Output Only] The URL to an InterconnectAttachment which is the next hop for the route. This field will only be populated for the dynamic routes generated by Cloud Router with a linked interconnectAttachment.
        /// </summary>
        public readonly string NextHopInterconnectAttachment;
        /// <summary>
        /// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
        /// </summary>
        public readonly string NextHopIp;
        /// <summary>
        /// The URL of the local network if it should handle matching packets.
        /// </summary>
        public readonly string NextHopNetwork;
        /// <summary>
        /// [Output Only] The network peering name that should handle matching packets, which should conform to RFC1035.
        /// </summary>
        public readonly string NextHopPeering;
        /// <summary>
        /// The URL to a VpnTunnel that should handle matching packets.
        /// </summary>
        public readonly string NextHopVpnTunnel;
        /// <summary>
        /// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// [Output Only] Server-defined fully-qualified URL for this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// A list of instance tags to which this route applies.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// [Output Only] If potential misconfigurations are detected for this route, this field will be populated with warning messages.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouteWarningsItemResponse> Warnings;

        [OutputConstructor]
        private GetRouteResult(
            string creationTimestamp,

            string description,

            string destRange,

            string kind,

            string name,

            string network,

            string nextHopGateway,

            string nextHopIlb,

            string nextHopInstance,

            string nextHopInterconnectAttachment,

            string nextHopIp,

            string nextHopNetwork,

            string nextHopPeering,

            string nextHopVpnTunnel,

            int priority,

            string selfLink,

            ImmutableArray<string> tags,

            ImmutableArray<Outputs.RouteWarningsItemResponse> warnings)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            DestRange = destRange;
            Kind = kind;
            Name = name;
            Network = network;
            NextHopGateway = nextHopGateway;
            NextHopIlb = nextHopIlb;
            NextHopInstance = nextHopInstance;
            NextHopInterconnectAttachment = nextHopInterconnectAttachment;
            NextHopIp = nextHopIp;
            NextHopNetwork = nextHopNetwork;
            NextHopPeering = nextHopPeering;
            NextHopVpnTunnel = nextHopVpnTunnel;
            Priority = priority;
            SelfLink = selfLink;
            Tags = tags;
            Warnings = warnings;
        }
    }
}
