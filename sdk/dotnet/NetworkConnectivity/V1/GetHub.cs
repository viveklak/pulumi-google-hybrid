// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.NetworkConnectivity.V1
{
    public static class GetHub
    {
        /// <summary>
        /// Gets details about the specified hub.
        /// </summary>
        public static Task<GetHubResult> InvokeAsync(GetHubArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHubResult>("google-native:networkconnectivity/v1:getHub", args ?? new GetHubArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details about the specified hub.
        /// </summary>
        public static Output<GetHubResult> Invoke(GetHubInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetHubResult>("google-native:networkconnectivity/v1:getHub", args ?? new GetHubInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHubArgs : Pulumi.InvokeArgs
    {
        [Input("hubId", required: true)]
        public string HubId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetHubArgs()
        {
        }
    }

    public sealed class GetHubInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("hubId", required: true)]
        public Input<string> HubId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetHubInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHubResult
    {
        /// <summary>
        /// The time the hub was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// An optional description of the hub.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Immutable. The name of the hub. Hub names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The VPC network associated with this hub's spokes. All of the VPN tunnels, VLAN attachments, and router appliance instances referenced by this hub's spokes must belong to this VPC network. This field is read-only. Network Connectivity Center automatically populates it based on the set of spokes attached to the hub.
        /// </summary>
        public readonly ImmutableArray<Outputs.RoutingVPCResponse> RoutingVpcs;
        /// <summary>
        /// The current lifecycle state of this hub.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The Google-generated UUID for the hub. This value is unique across all hub resources. If a hub is deleted and another with the same name is created, the new hub is assigned a different unique_id.
        /// </summary>
        public readonly string UniqueId;
        /// <summary>
        /// The time the hub was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetHubResult(
            string createTime,

            string description,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<Outputs.RoutingVPCResponse> routingVpcs,

            string state,

            string uniqueId,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Labels = labels;
            Name = name;
            RoutingVpcs = routingVpcs;
            State = state;
            UniqueId = uniqueId;
            UpdateTime = updateTime;
        }
    }
}
