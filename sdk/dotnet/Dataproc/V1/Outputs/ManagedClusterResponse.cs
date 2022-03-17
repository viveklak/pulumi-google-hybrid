// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// Cluster that is managed by the workflow.
    /// </summary>
    [OutputType]
    public sealed class ManagedClusterResponse
    {
        /// <summary>
        /// The cluster name prefix. A unique cluster name will be formed by appending a random suffix.The name must contain only lower-case letters (a-z), numbers (0-9), and hyphens (-). Must begin with a letter. Cannot begin or end with hyphen. Must consist of between 2 and 35 characters.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// The cluster configuration.
        /// </summary>
        public readonly Outputs.ClusterConfigResponse Config;
        /// <summary>
        /// Optional. The labels to associate with this cluster.Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62}Label values must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}\p{N}_-{0,63}No more than 32 labels can be associated with a given cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;

        [OutputConstructor]
        private ManagedClusterResponse(
            string clusterName,

            Outputs.ClusterConfigResponse config,

            ImmutableDictionary<string, string> labels)
        {
            ClusterName = clusterName;
            Config = config;
            Labels = labels;
        }
    }
}
