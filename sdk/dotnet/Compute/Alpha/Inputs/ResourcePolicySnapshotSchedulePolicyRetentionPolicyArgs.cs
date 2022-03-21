// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Policy for retention of scheduled snapshots.
    /// </summary>
    public sealed class ResourcePolicySnapshotSchedulePolicyRetentionPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum age of the snapshot that is allowed to be kept.
        /// </summary>
        [Input("maxRetentionDays")]
        public Input<int>? MaxRetentionDays { get; set; }

        [Input("onPolicySwitch")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.ResourcePolicySnapshotSchedulePolicyRetentionPolicyOnPolicySwitch>? OnPolicySwitch { get; set; }

        /// <summary>
        /// Specifies the behavior to apply to scheduled snapshots when the source disk is deleted.
        /// </summary>
        [Input("onSourceDiskDelete")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.ResourcePolicySnapshotSchedulePolicyRetentionPolicyOnSourceDiskDelete>? OnSourceDiskDelete { get; set; }

        public ResourcePolicySnapshotSchedulePolicyRetentionPolicyArgs()
        {
        }
    }
}
