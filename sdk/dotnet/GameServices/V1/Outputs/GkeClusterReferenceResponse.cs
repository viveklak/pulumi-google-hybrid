// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1.Outputs
{

    /// <summary>
    /// A reference to a GKE cluster.
    /// </summary>
    [OutputType]
    public sealed class GkeClusterReferenceResponse
    {
        /// <summary>
        /// The full or partial name of a GKE cluster, using one of the following forms: * `projects/{project}/locations/{locationId}/clusters/{cluster}` * `locations/{locationId}/clusters/{cluster}` * `{cluster}` If project and location are not specified, the project and location of the GameServerCluster resource are used to generate the full name of the GKE cluster.
        /// </summary>
        public readonly string Cluster;

        [OutputConstructor]
        private GkeClusterReferenceResponse(string cluster)
        {
            Cluster = cluster;
        }
    }
}
