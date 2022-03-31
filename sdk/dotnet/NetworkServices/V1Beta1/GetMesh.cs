// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1
{
    public static class GetMesh
    {
        /// <summary>
        /// Gets details of a single Mesh.
        /// </summary>
        public static Task<GetMeshResult> InvokeAsync(GetMeshArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMeshResult>("google-native:networkservices/v1beta1:getMesh", args ?? new GetMeshArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Mesh.
        /// </summary>
        public static Output<GetMeshResult> Invoke(GetMeshInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMeshResult>("google-native:networkservices/v1beta1:getMesh", args ?? new GetMeshInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMeshArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("meshId", required: true)]
        public string MeshId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetMeshArgs()
        {
        }
    }

    public sealed class GetMeshInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("meshId", required: true)]
        public Input<string> MeshId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetMeshInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMeshResult
    {
        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to be redirected to this port regardless of its actual ip:port destination. If unset, a port '15001' is used as the interception port. This will is applicable only for sidecar proxy deployments.
        /// </summary>
        public readonly int InterceptionPort;
        /// <summary>
        /// Optional. Set of label tags associated with the Mesh resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Name of the Mesh resource. It matches pattern `projects/*/locations/global/meshes/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Server-defined URL of this resource
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetMeshResult(
            string createTime,

            string description,

            int interceptionPort,

            ImmutableDictionary<string, string> labels,

            string name,

            string selfLink,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            InterceptionPort = interceptionPort;
            Labels = labels;
            Name = name;
            SelfLink = selfLink;
            UpdateTime = updateTime;
        }
    }
}
