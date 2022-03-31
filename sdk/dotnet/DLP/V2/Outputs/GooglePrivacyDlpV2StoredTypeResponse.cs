// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// A reference to a StoredInfoType to use with scanning.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2StoredTypeResponse
    {
        /// <summary>
        /// Timestamp indicating when the version of the `StoredInfoType` used for inspection was created. Output-only field, populated by the system.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Resource name of the requested `StoredInfoType`, for example `organizations/433245324/storedInfoTypes/432452342` or `projects/project-id/storedInfoTypes/432452342`.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GooglePrivacyDlpV2StoredTypeResponse(
            string createTime,

            string name)
        {
            CreateTime = createTime;
            Name = name;
        }
    }
}
