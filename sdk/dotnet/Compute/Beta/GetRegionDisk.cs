// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetRegionDisk
    {
        /// <summary>
        /// Returns a specified regional persistent disk.
        /// </summary>
        public static Task<GetRegionDiskResult> InvokeAsync(GetRegionDiskArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionDiskResult>("google-native:compute/beta:getRegionDisk", args ?? new GetRegionDiskArgs(), options.WithVersion());
    }


    public sealed class GetRegionDiskArgs : Pulumi.InvokeArgs
    {
        [Input("disk", required: true)]
        public string Disk { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetRegionDiskArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionDiskResult
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later (e.g. to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse DiskEncryptionKey;
        /// <summary>
        /// Specifies whether the disk restored from a source snapshot should erase Windows specific VSS signature.
        /// </summary>
        public readonly bool EraseWindowsVssSignature;
        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
        /// </summary>
        public readonly ImmutableArray<Outputs.GuestOsFeatureResponse> GuestOsFeatures;
        /// <summary>
        /// Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// [Output Only] Type of the resource. Always compute#disk for disks.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a disk.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// Labels to apply to this disk. These can be later modified by the setLabels method.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// [Output Only] Last attach timestamp in RFC3339 text format.
        /// </summary>
        public readonly string LastAttachTimestamp;
        /// <summary>
        /// [Output Only] Last detach timestamp in RFC3339 text format.
        /// </summary>
        public readonly string LastDetachTimestamp;
        /// <summary>
        /// Integer license codes indicating which licenses are attached to this disk.
        /// </summary>
        public readonly ImmutableArray<string> LicenseCodes;
        /// <summary>
        /// A list of publicly visible licenses. Reserved for Google's use.
        /// </summary>
        public readonly ImmutableArray<string> Licenses;
        /// <summary>
        /// An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
        /// </summary>
        public readonly string LocationHint;
        /// <summary>
        /// Indicates whether or not the disk can be read/write attached to more than one instance.
        /// </summary>
        public readonly bool MultiWriter;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Internal use only.
        /// </summary>
        public readonly string Options;
        /// <summary>
        /// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
        /// </summary>
        public readonly string PhysicalBlockSizeBytes;
        /// <summary>
        /// Indicates how many IOPS must be provisioned for the disk.
        /// </summary>
        public readonly string ProvisionedIops;
        /// <summary>
        /// [Output Only] URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
        /// </summary>
        public readonly ImmutableArray<string> ReplicaZones;
        /// <summary>
        /// Resource policies applied to this disk for automatic snapshot creations.
        /// </summary>
        public readonly ImmutableArray<string> ResourcePolicies;
        /// <summary>
        /// [Output Only] Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// [Output Only] Server-defined fully-qualified URL for this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
        /// </summary>
        public readonly string SizeGb;
        /// <summary>
        /// The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk 
        /// </summary>
        public readonly string SourceDisk;
        /// <summary>
        /// [Output Only] The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
        /// </summary>
        public readonly string SourceDiskId;
        /// <summary>
        /// The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family 
        /// </summary>
        public readonly string SourceImage;
        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse SourceImageEncryptionKey;
        /// <summary>
        /// [Output Only] The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
        /// </summary>
        public readonly string SourceImageId;
        /// <summary>
        /// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot 
        /// </summary>
        public readonly string SourceSnapshot;
        /// <summary>
        /// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse SourceSnapshotEncryptionKey;
        /// <summary>
        /// [Output Only] The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
        /// </summary>
        public readonly string SourceSnapshotId;
        /// <summary>
        /// The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
        /// </summary>
        public readonly string SourceStorageObject;
        /// <summary>
        /// [Output Only] The status of disk creation. - CREATING: Disk is provisioning. - RESTORING: Source data is being copied into the disk. - FAILED: Disk creation failed. - READY: Disk is ready for use. - DELETING: Disk is deleting. 
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-standard or pd-ssd
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// [Output Only] Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
        /// </summary>
        public readonly ImmutableArray<string> Users;
        /// <summary>
        /// [Output Only] URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetRegionDiskResult(
            string creationTimestamp,

            string description,

            Outputs.CustomerEncryptionKeyResponse diskEncryptionKey,

            bool eraseWindowsVssSignature,

            ImmutableArray<Outputs.GuestOsFeatureResponse> guestOsFeatures,

            string @interface,

            string kind,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            string lastAttachTimestamp,

            string lastDetachTimestamp,

            ImmutableArray<string> licenseCodes,

            ImmutableArray<string> licenses,

            string locationHint,

            bool multiWriter,

            string name,

            string options,

            string physicalBlockSizeBytes,

            string provisionedIops,

            string region,

            ImmutableArray<string> replicaZones,

            ImmutableArray<string> resourcePolicies,

            bool satisfiesPzs,

            string selfLink,

            string sizeGb,

            string sourceDisk,

            string sourceDiskId,

            string sourceImage,

            Outputs.CustomerEncryptionKeyResponse sourceImageEncryptionKey,

            string sourceImageId,

            string sourceSnapshot,

            Outputs.CustomerEncryptionKeyResponse sourceSnapshotEncryptionKey,

            string sourceSnapshotId,

            string sourceStorageObject,

            string status,

            string type,

            ImmutableArray<string> users,

            string zone)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            DiskEncryptionKey = diskEncryptionKey;
            EraseWindowsVssSignature = eraseWindowsVssSignature;
            GuestOsFeatures = guestOsFeatures;
            Interface = @interface;
            Kind = kind;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            LastAttachTimestamp = lastAttachTimestamp;
            LastDetachTimestamp = lastDetachTimestamp;
            LicenseCodes = licenseCodes;
            Licenses = licenses;
            LocationHint = locationHint;
            MultiWriter = multiWriter;
            Name = name;
            Options = options;
            PhysicalBlockSizeBytes = physicalBlockSizeBytes;
            ProvisionedIops = provisionedIops;
            Region = region;
            ReplicaZones = replicaZones;
            ResourcePolicies = resourcePolicies;
            SatisfiesPzs = satisfiesPzs;
            SelfLink = selfLink;
            SizeGb = sizeGb;
            SourceDisk = sourceDisk;
            SourceDiskId = sourceDiskId;
            SourceImage = sourceImage;
            SourceImageEncryptionKey = sourceImageEncryptionKey;
            SourceImageId = sourceImageId;
            SourceSnapshot = sourceSnapshot;
            SourceSnapshotEncryptionKey = sourceSnapshotEncryptionKey;
            SourceSnapshotId = sourceSnapshotId;
            SourceStorageObject = sourceStorageObject;
            Status = status;
            Type = type;
            Users = users;
            Zone = zone;
        }
    }
}
