// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Inputs
{

    /// <summary>
    /// CommonFeatureSpec contains Hub-wide configuration information
    /// </summary>
    public sealed class CommonFeatureSpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Anthos Observability spec
        /// </summary>
        [Input("anthosobservability")]
        public Input<Inputs.AnthosObservabilityFeatureSpecArgs>? Anthosobservability { get; set; }

        /// <summary>
        /// Appdevexperience specific spec.
        /// </summary>
        [Input("appdevexperience")]
        public Input<Inputs.AppDevExperienceFeatureSpecArgs>? Appdevexperience { get; set; }

        /// <summary>
        /// Cloud Audit Logging-specific spec.
        /// </summary>
        [Input("cloudauditlogging")]
        public Input<Inputs.CloudAuditLoggingFeatureSpecArgs>? Cloudauditlogging { get; set; }

        /// <summary>
        /// Multicluster Ingress-specific spec.
        /// </summary>
        [Input("multiclusteringress")]
        public Input<Inputs.MultiClusterIngressFeatureSpecArgs>? Multiclusteringress { get; set; }

        /// <summary>
        /// Workload Certificate spec.
        /// </summary>
        [Input("workloadcertificate")]
        public Input<Inputs.FeatureSpecArgs>? Workloadcertificate { get; set; }

        public CommonFeatureSpecArgs()
        {
        }
    }
}
