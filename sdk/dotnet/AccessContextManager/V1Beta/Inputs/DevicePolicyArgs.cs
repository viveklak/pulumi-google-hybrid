// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1Beta.Inputs
{

    /// <summary>
    /// `DevicePolicy` specifies device specific restrictions necessary to acquire a given access level. A `DevicePolicy` specifies requirements for requests from devices to be granted access levels, it does not do any enforcement on the device. `DevicePolicy` acts as an AND over all specified fields, and each repeated field is an OR over its elements. Any unset fields are ignored. For example, if the proto is { os_type : DESKTOP_WINDOWS, os_type : DESKTOP_LINUX, encryption_status: ENCRYPTED}, then the DevicePolicy will be true for requests originating from encrypted Linux desktops and encrypted Windows desktops.
    /// </summary>
    public sealed class DevicePolicyArgs : Pulumi.ResourceArgs
    {
        [Input("allowedDeviceManagementLevels")]
        private InputList<Pulumi.GoogleNative.AccessContextManager.V1Beta.DevicePolicyAllowedDeviceManagementLevelsItem>? _allowedDeviceManagementLevels;

        /// <summary>
        /// Allowed device management levels, an empty list allows all management levels.
        /// </summary>
        public InputList<Pulumi.GoogleNative.AccessContextManager.V1Beta.DevicePolicyAllowedDeviceManagementLevelsItem> AllowedDeviceManagementLevels
        {
            get => _allowedDeviceManagementLevels ?? (_allowedDeviceManagementLevels = new InputList<Pulumi.GoogleNative.AccessContextManager.V1Beta.DevicePolicyAllowedDeviceManagementLevelsItem>());
            set => _allowedDeviceManagementLevels = value;
        }

        [Input("allowedEncryptionStatuses")]
        private InputList<Pulumi.GoogleNative.AccessContextManager.V1Beta.DevicePolicyAllowedEncryptionStatusesItem>? _allowedEncryptionStatuses;

        /// <summary>
        /// Allowed encryptions statuses, an empty list allows all statuses.
        /// </summary>
        public InputList<Pulumi.GoogleNative.AccessContextManager.V1Beta.DevicePolicyAllowedEncryptionStatusesItem> AllowedEncryptionStatuses
        {
            get => _allowedEncryptionStatuses ?? (_allowedEncryptionStatuses = new InputList<Pulumi.GoogleNative.AccessContextManager.V1Beta.DevicePolicyAllowedEncryptionStatusesItem>());
            set => _allowedEncryptionStatuses = value;
        }

        [Input("osConstraints")]
        private InputList<Inputs.OsConstraintArgs>? _osConstraints;

        /// <summary>
        /// Allowed OS versions, an empty list allows all types and all versions.
        /// </summary>
        public InputList<Inputs.OsConstraintArgs> OsConstraints
        {
            get => _osConstraints ?? (_osConstraints = new InputList<Inputs.OsConstraintArgs>());
            set => _osConstraints = value;
        }

        /// <summary>
        /// Whether the device needs to be approved by the customer admin.
        /// </summary>
        [Input("requireAdminApproval")]
        public Input<bool>? RequireAdminApproval { get; set; }

        /// <summary>
        /// Whether the device needs to be corp owned.
        /// </summary>
        [Input("requireCorpOwned")]
        public Input<bool>? RequireCorpOwned { get; set; }

        /// <summary>
        /// Whether or not screenlock is required for the DevicePolicy to be true. Defaults to `false`.
        /// </summary>
        [Input("requireScreenlock")]
        public Input<bool>? RequireScreenlock { get; set; }

        public DevicePolicyArgs()
        {
        }
    }
}
