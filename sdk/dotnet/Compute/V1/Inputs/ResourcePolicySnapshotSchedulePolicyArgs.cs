// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// A snapshot schedule policy specifies when and how frequently snapshots are to be created for the target disk. Also specifies how many and how long these scheduled snapshots should be retained.
    /// </summary>
    public sealed class ResourcePolicySnapshotSchedulePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Retention policy applied to snapshots created by this resource policy.
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.ResourcePolicySnapshotSchedulePolicyRetentionPolicyArgs>? RetentionPolicy { get; set; }

        /// <summary>
        /// A Vm Maintenance Policy specifies what kind of infrastructure maintenance we are allowed to perform on this VM and when. Schedule that is applied to disks covered by this policy.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.ResourcePolicySnapshotSchedulePolicyScheduleArgs>? Schedule { get; set; }

        /// <summary>
        /// Properties with which snapshots are created such as labels, encryption keys.
        /// </summary>
        [Input("snapshotProperties")]
        public Input<Inputs.ResourcePolicySnapshotSchedulePolicySnapshotPropertiesArgs>? SnapshotProperties { get; set; }

        public ResourcePolicySnapshotSchedulePolicyArgs()
        {
        }
    }
}
