// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Outputs
{

    /// <summary>
    /// **Multi-cluster Ingress**: The configuration for the MultiClusterIngress feature.
    /// </summary>
    [OutputType]
    public sealed class MultiClusterIngressFeatureSpecResponse
    {
        /// <summary>
        /// Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: `projects/foo-proj/locations/global/memberships/bar`
        /// </summary>
        public readonly string ConfigMembership;

        [OutputConstructor]
        private MultiClusterIngressFeatureSpecResponse(string configMembership)
        {
            ConfigMembership = configMembership;
        }
    }
}
