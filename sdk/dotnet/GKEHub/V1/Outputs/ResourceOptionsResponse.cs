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
    /// ResourceOptions represent options for Kubernetes resource generation.
    /// </summary>
    [OutputType]
    public sealed class ResourceOptionsResponse
    {
        /// <summary>
        /// Optional. The Connect agent version to use for connect_resources. Defaults to the latest GKE Connect version. The version must be a currently supported version, obsolete versions will be rejected.
        /// </summary>
        public readonly string ConnectVersion;
        /// <summary>
        /// Optional. Major version of the Kubernetes cluster. This is only used to determine which version to use for the CustomResourceDefinition resources, `apiextensions/v1beta1` or`apiextensions/v1`.
        /// </summary>
        public readonly string K8sVersion;
        /// <summary>
        /// Optional. Use `apiextensions/v1beta1` instead of `apiextensions/v1` for CustomResourceDefinition resources. This option should be set for clusters with Kubernetes apiserver versions &lt;1.16.
        /// </summary>
        public readonly bool V1beta1Crd;

        [OutputConstructor]
        private ResourceOptionsResponse(
            string connectVersion,

            string k8sVersion,

            bool v1beta1Crd)
        {
            ConnectVersion = connectVersion;
            K8sVersion = k8sVersion;
            V1beta1Crd = v1beta1Crd;
        }
    }
}
