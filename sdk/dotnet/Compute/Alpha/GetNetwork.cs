// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetNetwork
    {
        /// <summary>
        /// Returns the specified network. Gets a list of available networks by making a list() request.
        /// </summary>
        public static Task<GetNetworkResult> InvokeAsync(GetNetworkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkResult>("google-native:compute/alpha:getNetwork", args ?? new GetNetworkArgs(), options.WithVersion());
    }


    public sealed class GetNetworkArgs : Pulumi.InvokeArgs
    {
        [Input("network", required: true)]
        public string Network { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetNetworkArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkResult
    {
        /// <summary>
        /// Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
        /// </summary>
        public readonly bool AutoCreateSubnetworks;
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// [Output Only] URL of the firewall policy the network is associated with.
        /// </summary>
        public readonly string FirewallPolicy;
        /// <summary>
        /// [Output Only] The gateway address for default routing out of the network, selected by GCP.
        /// </summary>
        public readonly string GatewayIPv4;
        /// <summary>
        /// [Output Only] Type of the resource. Always compute#network for networks.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes.
        /// </summary>
        public readonly int Mtu;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// [Output Only] A list of network peerings for the resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkPeeringResponse> Peerings;
        /// <summary>
        /// The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
        /// </summary>
        public readonly Outputs.NetworkRoutingConfigResponse RoutingConfig;
        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// [Output Only] Server-defined fully-qualified URLs for all subnetworks in this VPC network.
        /// </summary>
        public readonly ImmutableArray<string> Subnetworks;

        [OutputConstructor]
        private GetNetworkResult(
            bool autoCreateSubnetworks,

            string creationTimestamp,

            string description,

            string firewallPolicy,

            string gatewayIPv4,

            string kind,

            int mtu,

            string name,

            ImmutableArray<Outputs.NetworkPeeringResponse> peerings,

            Outputs.NetworkRoutingConfigResponse routingConfig,

            string selfLink,

            string selfLinkWithId,

            ImmutableArray<string> subnetworks)
        {
            AutoCreateSubnetworks = autoCreateSubnetworks;
            CreationTimestamp = creationTimestamp;
            Description = description;
            FirewallPolicy = firewallPolicy;
            GatewayIPv4 = gatewayIPv4;
            Kind = kind;
            Mtu = mtu;
            Name = name;
            Peerings = peerings;
            RoutingConfig = routingConfig;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            Subnetworks = subnetworks;
        }
    }
}
