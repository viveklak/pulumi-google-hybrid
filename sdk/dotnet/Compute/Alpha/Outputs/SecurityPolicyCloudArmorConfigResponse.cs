// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// Configuration options for Cloud Armor.
    /// </summary>
    [OutputType]
    public sealed class SecurityPolicyCloudArmorConfigResponse
    {
        /// <summary>
        /// If set to true, enables Cloud Armor Machine Learning.
        /// </summary>
        public readonly bool EnableMl;

        [OutputConstructor]
        private SecurityPolicyCloudArmorConfigResponse(bool enableMl)
        {
            EnableMl = enableMl;
        }
    }
}
