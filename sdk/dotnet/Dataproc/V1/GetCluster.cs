// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1
{
    public static class GetCluster
    {
        /// <summary>
        /// Gets the resource representation for a cluster in a project.
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("google-native:dataproc/v1:getCluster", args ?? new GetClusterArgs(), options.WithVersion());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The cluster name. Cluster names within a project must be unique. Names of deleted clusters can be reused.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster.
        /// </summary>
        public readonly string ClusterUuid;
        /// <summary>
        /// The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
        /// </summary>
        public readonly Outputs.ClusterConfigResponse Config;
        /// <summary>
        /// Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Contains cluster daemon metrics such as HDFS and YARN stats.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
        /// </summary>
        public readonly Outputs.ClusterMetricsResponse Metrics;
        /// <summary>
        /// The Google Cloud Platform project ID that the cluster belongs to.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Cluster status.
        /// </summary>
        public readonly Outputs.ClusterStatusResponse Status;
        /// <summary>
        /// The previous cluster status.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterStatusResponse> StatusHistory;

        [OutputConstructor]
        private GetClusterResult(
            string clusterName,

            string clusterUuid,

            Outputs.ClusterConfigResponse config,

            ImmutableDictionary<string, string> labels,

            Outputs.ClusterMetricsResponse metrics,

            string project,

            Outputs.ClusterStatusResponse status,

            ImmutableArray<Outputs.ClusterStatusResponse> statusHistory)
        {
            ClusterName = clusterName;
            ClusterUuid = clusterUuid;
            Config = config;
            Labels = labels;
            Metrics = metrics;
            Project = project;
            Status = status;
            StatusHistory = statusHistory;
        }
    }
}
