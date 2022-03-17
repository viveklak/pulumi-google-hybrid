// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha2.Outputs
{

    /// <summary>
    /// EdgeCluster contains information specific to Google Edge Clusters.
    /// </summary>
    [OutputType]
    public sealed class EdgeClusterResponse
    {
        /// <summary>
        /// Immutable. Self-link of the GCP resource for the Edge Cluster. For example: //edgecontainer.googleapis.com/projects/my-project/locations/us-west1-a/clusters/my-cluster
        /// </summary>
        public readonly string ResourceLink;

        [OutputConstructor]
        private EdgeClusterResponse(string resourceLink)
        {
            ResourceLink = resourceLink;
        }
    }
}
