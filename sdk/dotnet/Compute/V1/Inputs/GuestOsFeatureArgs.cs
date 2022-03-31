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
    /// Guest OS features.
    /// </summary>
    public sealed class GuestOsFeatureArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of a supported feature. To add multiple values, use commas to separate values. Set to one or more of the following values: - VIRTIO_SCSI_MULTIQUEUE - WINDOWS - MULTI_IP_SUBNET - UEFI_COMPATIBLE - GVNIC - SEV_CAPABLE - SUSPEND_RESUME_COMPATIBLE - SEV_SNP_CAPABLE For more information, see Enabling guest operating system features.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Compute.V1.GuestOsFeatureType>? Type { get; set; }

        public GuestOsFeatureArgs()
        {
        }
    }
}
