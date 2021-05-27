// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2
{
    /// <summary>
    /// Creates a cluster in a project. The returned Operation.metadata will be ClusterOperationMetadata (https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1beta2#clusteroperationmetadata).
    /// </summary>
    [GoogleNativeResourceType("google-native:dataproc/v1beta2:RegionCluster")]
    public partial class RegionCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Required. The cluster name. Cluster names within a project must be unique. Names of deleted clusters can be reused.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster.
        /// </summary>
        [Output("clusterUuid")]
        public Output<string> ClusterUuid { get; private set; } = null!;

        /// <summary>
        /// Required. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
        /// </summary>
        [Output("config")]
        public Output<Outputs.ClusterConfigResponse> Config { get; private set; } = null!;

        /// <summary>
        /// Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Contains cluster daemon metrics such as HDFS and YARN stats.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
        /// </summary>
        [Output("metrics")]
        public Output<Outputs.ClusterMetricsResponse> Metrics { get; private set; } = null!;

        /// <summary>
        /// Required. The Google Cloud Platform project ID that the cluster belongs to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Cluster status.
        /// </summary>
        [Output("status")]
        public Output<Outputs.ClusterStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// The previous cluster status.
        /// </summary>
        [Output("statusHistory")]
        public Output<ImmutableArray<Outputs.ClusterStatusResponse>> StatusHistory { get; private set; } = null!;


        /// <summary>
        /// Create a RegionCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionCluster(string name, RegionClusterArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataproc/v1beta2:RegionCluster", name, args ?? new RegionClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataproc/v1beta2:RegionCluster", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RegionCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionCluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegionCluster(name, id, options);
        }
    }

    public sealed class RegionClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The cluster name. Cluster names within a project must be unique. Names of deleted clusters can be reused.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Required. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
        /// </summary>
        [Input("config")]
        public Input<Inputs.ClusterConfigArgs>? Config { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Required. The Google Cloud Platform project ID that the cluster belongs to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        public RegionClusterArgs()
        {
        }
    }
}
