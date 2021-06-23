// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1Beta
{
    public static class GetGameServerCluster
    {
        /// <summary>
        /// Gets details of a single game server cluster.
        /// </summary>
        public static Task<GetGameServerClusterResult> InvokeAsync(GetGameServerClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGameServerClusterResult>("google-native:gameservices/v1beta:getGameServerCluster", args ?? new GetGameServerClusterArgs(), options.WithVersion());
    }


    public sealed class GetGameServerClusterArgs : Pulumi.InvokeArgs
    {
        [Input("gameServerClusterId", required: true)]
        public string GameServerClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        [Input("view")]
        public string? View { get; set; }

        public GetGameServerClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGameServerClusterResult
    {
        /// <summary>
        /// Optional. The allocation priority assigned to the game server cluster. Game server clusters receive new game server allocations based on the relative allocation priorites set for each cluster, if the realm is configured for multicluster allocation.
        /// </summary>
        public readonly string AllocationPriority;
        /// <summary>
        /// The state of the Kubernetes cluster, this will be available if 'view' is set to `FULL` in the relevant List/Get/Preview request.
        /// </summary>
        public readonly Outputs.KubernetesClusterStateResponse ClusterState;
        /// <summary>
        /// The game server cluster connection information. This information is used to manage game server clusters.
        /// </summary>
        public readonly Outputs.GameServerClusterConnectionInfoResponse ConnectionInfo;
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Human readable description of the cluster.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ETag of the resource.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The labels associated with this game server cluster. Each label is a key-value pair.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Required. The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The last-modified time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetGameServerClusterResult(
            string allocationPriority,

            Outputs.KubernetesClusterStateResponse clusterState,

            Outputs.GameServerClusterConnectionInfoResponse connectionInfo,

            string createTime,

            string description,

            string etag,

            ImmutableDictionary<string, string> labels,

            string name,

            string updateTime)
        {
            AllocationPriority = allocationPriority;
            ClusterState = clusterState;
            ConnectionInfo = connectionInfo;
            CreateTime = createTime;
            Description = description;
            Etag = etag;
            Labels = labels;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}
