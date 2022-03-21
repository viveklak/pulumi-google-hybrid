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
    /// The state of the Kubernetes cluster.
    /// </summary>
    [OutputType]
    public sealed class KubernetesClusterStateResponse
    {
        /// <summary>
        /// The version of Agones currently installed in the registered Kubernetes cluster.
        /// </summary>
        public readonly string AgonesVersionInstalled;
        /// <summary>
        /// The version of Agones that is targeted to be installed in the cluster.
        /// </summary>
        public readonly string AgonesVersionTargeted;
        /// <summary>
        /// The state for the installed versions of Agones/Kubernetes.
        /// </summary>
        public readonly string InstallationState;
        /// <summary>
        /// The version of Kubernetes that is currently used in the registered Kubernetes cluster (as detected by the Cloud Game Servers service).
        /// </summary>
        public readonly string KubernetesVersionInstalled;
        /// <summary>
        /// The cloud provider type reported by the first node's `providerID` in the list of nodes on the Kubernetes endpoint. On Kubernetes platforms that support zero-node clusters (like GKE-on-GCP), the provider type will be empty.
        /// </summary>
        public readonly string Provider;
        /// <summary>
        /// The detailed error message for the installed versions of Agones/Kubernetes.
        /// </summary>
        public readonly string VersionInstalledErrorMessage;

        [OutputConstructor]
        private KubernetesClusterStateResponse(
            string agonesVersionInstalled,

            string agonesVersionTargeted,

            string installationState,

            string kubernetesVersionInstalled,

            string provider,

            string versionInstalledErrorMessage)
        {
            AgonesVersionInstalled = agonesVersionInstalled;
            AgonesVersionTargeted = agonesVersionTargeted;
            InstallationState = installationState;
            KubernetesVersionInstalled = kubernetesVersionInstalled;
            Provider = provider;
            VersionInstalledErrorMessage = versionInstalledErrorMessage;
        }
    }
}
