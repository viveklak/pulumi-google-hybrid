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
    /// **Workload Certificate**: The Hub-wide input for the WorkloadCertificate feature.
    /// </summary>
    public sealed class FeatureSpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies default membership spec. Users can override the default in the member_configs for each member.
        /// </summary>
        [Input("defaultConfig")]
        public Input<Inputs.MembershipSpecArgs>? DefaultConfig { get; set; }

        /// <summary>
        /// Immutable. Specifies CA configuration.
        /// </summary>
        [Input("provisionGoogleCa")]
        public Input<Pulumi.GoogleNative.GKEHub.V1Alpha.FeatureSpecProvisionGoogleCa>? ProvisionGoogleCa { get; set; }

        public FeatureSpecArgs()
        {
        }
    }
}
