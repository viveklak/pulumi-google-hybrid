// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// A set of Shielded Instance options.
    /// </summary>
    [OutputType]
    public sealed class ShieldedInstanceConfigResponse
    {
        /// <summary>
        /// Defines whether the instance has integrity monitoring enabled. Enabled by default.
        /// </summary>
        public readonly bool EnableIntegrityMonitoring;
        /// <summary>
        /// Defines whether the instance has Secure Boot enabled. Disabled by default.
        /// </summary>
        public readonly bool EnableSecureBoot;
        /// <summary>
        /// Defines whether the instance has the vTPM enabled. Enabled by default.
        /// </summary>
        public readonly bool EnableVtpm;

        [OutputConstructor]
        private ShieldedInstanceConfigResponse(
            bool enableIntegrityMonitoring,

            bool enableSecureBoot,

            bool enableVtpm)
        {
            EnableIntegrityMonitoring = enableIntegrityMonitoring;
            EnableSecureBoot = enableSecureBoot;
            EnableVtpm = enableVtpm;
        }
    }
}
