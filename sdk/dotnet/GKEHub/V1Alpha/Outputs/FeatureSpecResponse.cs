// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Outputs
{

    /// <summary>
    /// **Workload Certificate**: The Hub-wide input for the WorkloadCertificate feature.
    /// </summary>
    [OutputType]
    public sealed class FeatureSpecResponse
    {
        /// <summary>
        /// Specifies default membership spec. Users can override the default in the member_configs for each member.
        /// </summary>
        public readonly Outputs.MembershipSpecResponse DefaultConfig;
        /// <summary>
        /// Immutable. Specifies CA configuration.
        /// </summary>
        public readonly string ProvisionGoogleCa;

        [OutputConstructor]
        private FeatureSpecResponse(
            Outputs.MembershipSpecResponse defaultConfig,

            string provisionGoogleCa)
        {
            DefaultConfig = defaultConfig;
            ProvisionGoogleCa = provisionGoogleCa;
        }
    }
}
