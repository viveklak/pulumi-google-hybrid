// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetTargetVpnGateway
    {
        /// <summary>
        /// Returns the specified target VPN gateway. Gets a list of available target VPN gateways by making a list() request.
        /// </summary>
        public static Task<GetTargetVpnGatewayResult> InvokeAsync(GetTargetVpnGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetVpnGatewayResult>("google-native:compute/v1:getTargetVpnGateway", args ?? new GetTargetVpnGatewayArgs(), options.WithVersion());
    }


    public sealed class GetTargetVpnGatewayArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("targetVpnGateway", required: true)]
        public string TargetVpnGateway { get; set; } = null!;

        public GetTargetVpnGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTargetVpnGatewayResult
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
        /// A list of URLs to the ForwardingRule resources. ForwardingRules are created using compute.forwardingRules.insert and associated with a VPN gateway.
        /// </summary>
        public readonly ImmutableArray<string> ForwardingRules;
        /// <summary>
        /// Type of resource. Always compute#targetVpnGateway for target VPN gateways.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// URL of the region where the target VPN gateway resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The status of the VPN gateway, which can be one of the following: CREATING, READY, FAILED, or DELETING.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A list of URLs to VpnTunnel resources. VpnTunnels are created using the compute.vpntunnels.insert method and associated with a VPN gateway.
        /// </summary>
        public readonly ImmutableArray<string> Tunnels;

        [OutputConstructor]
        private GetTargetVpnGatewayResult(
            string creationTimestamp,

            string description,

            ImmutableArray<string> forwardingRules,

            string kind,

            string name,

            string network,

            string region,

            string selfLink,

            string status,

            ImmutableArray<string> tunnels)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            ForwardingRules = forwardingRules;
            Kind = kind;
            Name = name;
            Network = network;
            Region = region;
            SelfLink = selfLink;
            Status = status;
            Tunnels = tunnels;
        }
    }
}
