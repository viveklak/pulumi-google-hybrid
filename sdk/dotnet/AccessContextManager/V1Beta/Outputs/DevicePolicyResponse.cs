// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1Beta.Outputs
{

    /// <summary>
    /// `DevicePolicy` specifies device specific restrictions necessary to acquire a given access level. A `DevicePolicy` specifies requirements for requests from devices to be granted access levels, it does not do any enforcement on the device. `DevicePolicy` acts as an AND over all specified fields, and each repeated field is an OR over its elements. Any unset fields are ignored. For example, if the proto is { os_type : DESKTOP_WINDOWS, os_type : DESKTOP_LINUX, encryption_status: ENCRYPTED}, then the DevicePolicy will be true for requests originating from encrypted Linux desktops and encrypted Windows desktops.
    /// </summary>
    [OutputType]
    public sealed class DevicePolicyResponse
    {
        /// <summary>
        /// Allowed device management levels, an empty list allows all management levels.
        /// </summary>
        public readonly ImmutableArray<string> AllowedDeviceManagementLevels;
        /// <summary>
        /// Allowed encryptions statuses, an empty list allows all statuses.
        /// </summary>
        public readonly ImmutableArray<string> AllowedEncryptionStatuses;
        /// <summary>
        /// Allowed OS versions, an empty list allows all types and all versions.
        /// </summary>
        public readonly ImmutableArray<Outputs.OsConstraintResponse> OsConstraints;
        /// <summary>
        /// Whether the device needs to be approved by the customer admin.
        /// </summary>
        public readonly bool RequireAdminApproval;
        /// <summary>
        /// Whether the device needs to be corp owned.
        /// </summary>
        public readonly bool RequireCorpOwned;
        /// <summary>
        /// Whether or not screenlock is required for the DevicePolicy to be true. Defaults to `false`.
        /// </summary>
        public readonly bool RequireScreenlock;

        [OutputConstructor]
        private DevicePolicyResponse(
            ImmutableArray<string> allowedDeviceManagementLevels,

            ImmutableArray<string> allowedEncryptionStatuses,

            ImmutableArray<Outputs.OsConstraintResponse> osConstraints,

            bool requireAdminApproval,

            bool requireCorpOwned,

            bool requireScreenlock)
        {
            AllowedDeviceManagementLevels = allowedDeviceManagementLevels;
            AllowedEncryptionStatuses = allowedEncryptionStatuses;
            OsConstraints = osConstraints;
            RequireAdminApproval = requireAdminApproval;
            RequireCorpOwned = requireCorpOwned;
            RequireScreenlock = requireScreenlock;
        }
    }
}
