// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetVpnGateway
    {
        /// <summary>
        /// Returns the specified VPN gateway. Gets a list of available VPN gateways by making a list() request.
        /// </summary>
        public static Task<GetVpnGatewayResult> InvokeAsync(GetVpnGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpnGatewayResult>("google-native:compute/alpha:getVpnGateway", args ?? new GetVpnGatewayArgs(), options.WithVersion());
    }


    public sealed class GetVpnGatewayArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("vpnGateway", required: true)]
        public string VpnGateway { get; set; } = null!;

        public GetVpnGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpnGatewayResult
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// [Output Only] Type of resource. Always compute#vpnGateway for VPN gateways.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A fingerprint for the labels being applied to this VpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an VpnGateway.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// [Output Only] URL of the region where the VPN gateway resides.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The list of VPN interfaces associated with this VPN gateway.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpnGatewayVpnGatewayInterfaceResponse> VpnInterfaces;

        [OutputConstructor]
        private GetVpnGatewayResult(
            string creationTimestamp,

            string description,

            string kind,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            string name,

            string network,

            string region,

            string selfLink,

            ImmutableArray<Outputs.VpnGatewayVpnGatewayInterfaceResponse> vpnInterfaces)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Kind = kind;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            Name = name;
            Network = network;
            Region = region;
            SelfLink = selfLink;
            VpnInterfaces = vpnInterfaces;
        }
    }
}
