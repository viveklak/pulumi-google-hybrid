// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// DEPRECATED: Please use compute#instanceProperties instead. New properties will not be added to this field.
    /// </summary>
    [OutputType]
    public sealed class SourceInstancePropertiesResponse
    {
        /// <summary>
        /// Enables instances created based on this machine image to send packets with source IP addresses other than their own and receive packets with destination IP addresses other than their own. If these instances will be used as an IP gateway or it will be set as the next-hop in a Route resource, specify true. If unsure, leave this set to false. See the Enable IP forwarding documentation for more information.
        /// </summary>
        public readonly bool CanIpForward;
        /// <summary>
        /// Whether the instance created from this machine image should be protected against deletion.
        /// </summary>
        public readonly bool DeletionProtection;
        /// <summary>
        /// An optional text description for the instances that are created from this machine image.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// An array of disks that are associated with the instances that are created from this machine image.
        /// </summary>
        public readonly ImmutableArray<Outputs.SavedAttachedDiskResponse> Disks;
        /// <summary>
        /// A list of guest accelerator cards' type and count to use for instances created from this machine image.
        /// </summary>
        public readonly ImmutableArray<Outputs.AcceleratorConfigResponse> GuestAccelerators;
        /// <summary>
        /// Labels to apply to instances that are created from this machine image.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The machine type to use for instances that are created from this machine image.
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// The metadata key/value pairs to assign to instances that are created from this machine image. These pairs can consist of custom metadata or predefined keys. See Project and instance metadata for more information.
        /// </summary>
        public readonly Outputs.MetadataResponse Metadata;
        /// <summary>
        /// Minimum cpu/platform to be used by instances created from this machine image. The instance may be scheduled on the specified or newer cpu/platform. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge". For more information, read Specifying a Minimum CPU Platform.
        /// </summary>
        public readonly string MinCpuPlatform;
        /// <summary>
        /// An array of network access configurations for this interface.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceResponse> NetworkInterfaces;
        /// <summary>
        /// PostKeyRevocationActionType of the instance.
        /// </summary>
        public readonly string PostKeyRevocationActionType;
        /// <summary>
        /// Specifies the scheduling options for the instances that are created from this machine image.
        /// </summary>
        public readonly Outputs.SchedulingResponse Scheduling;
        /// <summary>
        /// A list of service accounts with specified scopes. Access tokens for these service accounts are available to the instances that are created from this machine image. Use metadata queries to obtain the access tokens for these instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceAccountResponse> ServiceAccounts;
        /// <summary>
        /// A list of tags to apply to the instances that are created from this machine image. The tags identify valid sources or targets for network firewalls. The setTags method can modify this list of tags. Each tag within the list must comply with RFC1035.
        /// </summary>
        public readonly Outputs.TagsResponse Tags;

        [OutputConstructor]
        private SourceInstancePropertiesResponse(
            bool canIpForward,

            bool deletionProtection,

            string description,

            ImmutableArray<Outputs.SavedAttachedDiskResponse> disks,

            ImmutableArray<Outputs.AcceleratorConfigResponse> guestAccelerators,

            ImmutableDictionary<string, string> labels,

            string machineType,

            Outputs.MetadataResponse metadata,

            string minCpuPlatform,

            ImmutableArray<Outputs.NetworkInterfaceResponse> networkInterfaces,

            string postKeyRevocationActionType,

            Outputs.SchedulingResponse scheduling,

            ImmutableArray<Outputs.ServiceAccountResponse> serviceAccounts,

            Outputs.TagsResponse tags)
        {
            CanIpForward = canIpForward;
            DeletionProtection = deletionProtection;
            Description = description;
            Disks = disks;
            GuestAccelerators = guestAccelerators;
            Labels = labels;
            MachineType = machineType;
            Metadata = metadata;
            MinCpuPlatform = minCpuPlatform;
            NetworkInterfaces = networkInterfaces;
            PostKeyRevocationActionType = postKeyRevocationActionType;
            Scheduling = scheduling;
            ServiceAccounts = serviceAccounts;
            Tags = tags;
        }
    }
}
