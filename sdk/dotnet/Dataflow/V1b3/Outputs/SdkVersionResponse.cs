// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Outputs
{

    /// <summary>
    /// The version of the SDK used to run the job.
    /// </summary>
    [OutputType]
    public sealed class SdkVersionResponse
    {
        /// <summary>
        /// The support status for this SDK version.
        /// </summary>
        public readonly string SdkSupportStatus;
        /// <summary>
        /// The version of the SDK used to run the job.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// A readable string describing the version of the SDK.
        /// </summary>
        public readonly string VersionDisplayName;

        [OutputConstructor]
        private SdkVersionResponse(
            string sdkSupportStatus,

            string version,

            string versionDisplayName)
        {
            SdkSupportStatus = sdkSupportStatus;
            Version = version;
            VersionDisplayName = versionDisplayName;
        }
    }
}
