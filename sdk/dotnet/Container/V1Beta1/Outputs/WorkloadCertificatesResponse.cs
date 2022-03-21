// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// Configuration for issuance of mTLS keys and certificates to Kubernetes pods.
    /// </summary>
    [OutputType]
    public sealed class WorkloadCertificatesResponse
    {
        /// <summary>
        /// enable_certificates controls issuance of workload mTLS certificates. If set, the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster, which can then be configured by creating a WorkloadCertificateConfig Custom Resource. Requires Workload Identity (workload_pool must be non-empty).
        /// </summary>
        public readonly bool EnableCertificates;

        [OutputConstructor]
        private WorkloadCertificatesResponse(bool enableCertificates)
        {
            EnableCertificates = enableCertificates;
        }
    }
}
