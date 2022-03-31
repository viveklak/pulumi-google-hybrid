// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkConnectivity.V1
{
    /// <summary>
    /// Creates a spoke in the specified project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:networkconnectivity/v1:Spoke")]
    public partial class Spoke : Pulumi.CustomResource
    {
        /// <summary>
        /// The time the spoke was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// An optional description of the spoke.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Immutable. The name of the hub that this spoke is attached to.
        /// </summary>
        [Output("hub")]
        public Output<string> Hub { get; private set; } = null!;

        /// <summary>
        /// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// VLAN attachments that are associated with the spoke.
        /// </summary>
        [Output("linkedInterconnectAttachments")]
        public Output<Outputs.LinkedInterconnectAttachmentsResponse> LinkedInterconnectAttachments { get; private set; } = null!;

        /// <summary>
        /// Router appliance instances that are associated with the spoke.
        /// </summary>
        [Output("linkedRouterApplianceInstances")]
        public Output<Outputs.LinkedRouterApplianceInstancesResponse> LinkedRouterApplianceInstances { get; private set; } = null!;

        /// <summary>
        /// VPN tunnels that are associated with the spoke.
        /// </summary>
        [Output("linkedVpnTunnels")]
        public Output<Outputs.LinkedVpnTunnelsResponse> LinkedVpnTunnels { get; private set; } = null!;

        /// <summary>
        /// Immutable. The name of the spoke. Spoke names must be unique. They use the following form: `projects/{project_number}/locations/{region}/spokes/{spoke_id}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current lifecycle state of this spoke.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.
        /// </summary>
        [Output("uniqueId")]
        public Output<string> UniqueId { get; private set; } = null!;

        /// <summary>
        /// The time the spoke was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Spoke resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Spoke(string name, SpokeArgs args, CustomResourceOptions? options = null)
            : base("google-native:networkconnectivity/v1:Spoke", name, args ?? new SpokeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Spoke(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:networkconnectivity/v1:Spoke", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Spoke resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Spoke Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Spoke(name, id, options);
        }
    }

    public sealed class SpokeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of the spoke.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Immutable. The name of the hub that this spoke is attached to.
        /// </summary>
        [Input("hub")]
        public Input<string>? Hub { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// VLAN attachments that are associated with the spoke.
        /// </summary>
        [Input("linkedInterconnectAttachments")]
        public Input<Inputs.LinkedInterconnectAttachmentsArgs>? LinkedInterconnectAttachments { get; set; }

        /// <summary>
        /// Router appliance instances that are associated with the spoke.
        /// </summary>
        [Input("linkedRouterApplianceInstances")]
        public Input<Inputs.LinkedRouterApplianceInstancesArgs>? LinkedRouterApplianceInstances { get; set; }

        /// <summary>
        /// VPN tunnels that are associated with the spoke.
        /// </summary>
        [Input("linkedVpnTunnels")]
        public Input<Inputs.LinkedVpnTunnelsArgs>? LinkedVpnTunnels { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. The name of the spoke. Spoke names must be unique. They use the following form: `projects/{project_number}/locations/{region}/spokes/{spoke_id}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. A unique request ID (optional). If you specify this ID, you can use it in cases when you need to retry your request. When you need to retry, this ID lets the server know that it can ignore the request if it has already been completed. The server guarantees that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check to see whether the original operation was received. If it was, the server ignores the second request. This behavior prevents clients from mistakenly creating duplicate commitments. The request ID must be a valid UUID, with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Required. Unique id for the spoke to create.
        /// </summary>
        [Input("spokeId", required: true)]
        public Input<string> SpokeId { get; set; } = null!;

        public SpokeArgs()
        {
        }
    }
}
