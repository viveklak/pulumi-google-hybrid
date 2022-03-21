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
    /// [Input Only] Specifies the parameters for a new disk that will be created alongside the new instance. Use initialization parameters to create boot disks or local SSDs attached to the new instance. This property is mutually exclusive with the source property; you can only define one or the other, but not both.
    /// </summary>
    [OutputType]
    public sealed class AttachedDiskInitializeParamsResponse
    {
        /// <summary>
        /// The architecture of the attached disk. Valid values are arm64 or x86_64.
        /// </summary>
        public readonly string Architecture;
        /// <summary>
        /// An optional description. Provide this property when creating the disk.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Specifies the disk name. If not specified, the default is to use the name of the instance. If a disk with the same name already exists in the given region, the existing disk is attached to the new instance and the new disk is not created.
        /// </summary>
        public readonly string DiskName;
        /// <summary>
        /// Specifies the size of the disk in base-2 GB. The size must be at least 10 GB. If you specify a sourceImage, which is required for boot disks, the default size is the size of the sourceImage. If you do not specify a sourceImage, the default disk size is 500 GB.
        /// </summary>
        public readonly string DiskSizeGb;
        /// <summary>
        /// Specifies the disk type to use to create the instance. If not specified, the default is pd-standard, specified using the full URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone /diskTypes/pd-standard For a full list of acceptable values, see Persistent disk types. If you define this field, you can provide either the full or partial URL. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /diskTypes/diskType - projects/project/zones/zone/diskTypes/diskType - zones/zone/diskTypes/diskType Note that for InstanceTemplate, this is the name of the disk type, not URL.
        /// </summary>
        public readonly string DiskType;
        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options. Guest OS features are applied by merging initializeParams.guestOsFeatures and disks.guestOsFeatures
        /// </summary>
        public readonly ImmutableArray<Outputs.GuestOsFeatureResponse> GuestOsFeatures;
        /// <summary>
        /// [Deprecated] Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// Labels to apply to this disk. These can be later modified by the disks.setLabels method. This field is only applicable for persistent disks.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Integer license codes indicating which licenses are attached to this disk.
        /// </summary>
        public readonly ImmutableArray<string> LicenseCodes;
        /// <summary>
        /// A list of publicly visible licenses. Reserved for Google's use.
        /// </summary>
        public readonly ImmutableArray<string> Licenses;
        /// <summary>
        /// Indicates whether or not the disk can be read/write attached to more than one instance.
        /// </summary>
        public readonly bool MultiWriter;
        /// <summary>
        /// Specifies which action to take on instance update with this disk. Default is to use the existing disk.
        /// </summary>
        public readonly string OnUpdateAction;
        /// <summary>
        /// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
        /// </summary>
        public readonly string ProvisionedIops;
        /// <summary>
        /// URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
        /// </summary>
        public readonly ImmutableArray<string> ReplicaZones;
        /// <summary>
        /// Resource policies applied to this disk for automatic snapshot creations. Specified using the full or partial URL. For instance template, specify only the resource policy name.
        /// </summary>
        public readonly ImmutableArray<string> ResourcePolicies;
        /// <summary>
        /// The source image to create this disk. When creating a new instance, one of initializeParams.sourceImage or initializeParams.sourceSnapshot or disks.source is required except for local SSD. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family If the source image is deleted later, this field will not be set.
        /// </summary>
        public readonly string SourceImage;
        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key. Instance templates do not store customer-supplied encryption keys, so you cannot create disks for instances in a managed instance group if the source images are encrypted with your own keys.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse SourceImageEncryptionKey;
        /// <summary>
        /// The source instant-snapshot to create this disk. When creating a new instance, one of initializeParams.sourceSnapshot or initializeParams.sourceInstantSnapshot initializeParams.sourceImage or disks.source is required except for local SSD. To create a disk with a snapshot that you created, specify the snapshot name in the following format: us-central1-a/instantSnapshots/my-backup If the source instant-snapshot is deleted later, this field will not be set.
        /// </summary>
        public readonly string SourceInstantSnapshot;
        /// <summary>
        /// The source snapshot to create this disk. When creating a new instance, one of initializeParams.sourceSnapshot or initializeParams.sourceImage or disks.source is required except for local SSD. To create a disk with a snapshot that you created, specify the snapshot name in the following format: global/snapshots/my-backup If the source snapshot is deleted later, this field will not be set.
        /// </summary>
        public readonly string SourceSnapshot;
        /// <summary>
        /// The customer-supplied encryption key of the source snapshot.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse SourceSnapshotEncryptionKey;

        [OutputConstructor]
        private AttachedDiskInitializeParamsResponse(
            string architecture,

            string description,

            string diskName,

            string diskSizeGb,

            string diskType,

            ImmutableArray<Outputs.GuestOsFeatureResponse> guestOsFeatures,

            string @interface,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<string> licenseCodes,

            ImmutableArray<string> licenses,

            bool multiWriter,

            string onUpdateAction,

            string provisionedIops,

            ImmutableArray<string> replicaZones,

            ImmutableArray<string> resourcePolicies,

            string sourceImage,

            Outputs.CustomerEncryptionKeyResponse sourceImageEncryptionKey,

            string sourceInstantSnapshot,

            string sourceSnapshot,

            Outputs.CustomerEncryptionKeyResponse sourceSnapshotEncryptionKey)
        {
            Architecture = architecture;
            Description = description;
            DiskName = diskName;
            DiskSizeGb = diskSizeGb;
            DiskType = diskType;
            GuestOsFeatures = guestOsFeatures;
            Interface = @interface;
            Labels = labels;
            LicenseCodes = licenseCodes;
            Licenses = licenses;
            MultiWriter = multiWriter;
            OnUpdateAction = onUpdateAction;
            ProvisionedIops = provisionedIops;
            ReplicaZones = replicaZones;
            ResourcePolicies = resourcePolicies;
            SourceImage = sourceImage;
            SourceImageEncryptionKey = sourceImageEncryptionKey;
            SourceInstantSnapshot = sourceInstantSnapshot;
            SourceSnapshot = sourceSnapshot;
            SourceSnapshotEncryptionKey = sourceSnapshotEncryptionKey;
        }
    }
}
