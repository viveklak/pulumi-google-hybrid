// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetImage
    {
        /// <summary>
        /// Returns the specified image. Gets a list of available images by making a list() request.
        /// </summary>
        public static Task<GetImageResult> InvokeAsync(GetImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("google-native:compute/v1:getImage", args ?? new GetImageArgs(), options.WithVersion());
    }


    public sealed class GetImageArgs : Pulumi.InvokeArgs
    {
        [Input("image", required: true)]
        public string Image { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetImageArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetImageResult
    {
        /// <summary>
        /// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
        /// </summary>
        public readonly string ArchiveSizeBytes;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// The deprecation status associated with this image.
        /// </summary>
        public readonly Outputs.DeprecationStatusResponse Deprecated;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Size of the image when restored onto a persistent disk (in GB).
        /// </summary>
        public readonly string DiskSizeGb;
        /// <summary>
        /// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
        /// </summary>
        public readonly string Family;
        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
        /// </summary>
        public readonly ImmutableArray<Outputs.GuestOsFeatureResponse> GuestOsFeatures;
        /// <summary>
        /// Encrypts the image using a customer-supplied encryption key.
        /// 
        /// After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image).
        /// 
        /// Customer-supplied encryption keys do not protect access to metadata of the disk.
        /// 
        /// If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse ImageEncryptionKey;
        /// <summary>
        /// Type of the resource. Always compute#image for images.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve an image.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// Labels to apply to this image. These can be later modified by the setLabels method.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Integer license codes indicating which licenses are attached to this image.
        /// </summary>
        public readonly ImmutableArray<string> LicenseCodes;
        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        public readonly ImmutableArray<string> Licenses;
        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The parameters of the raw disk image.
        /// </summary>
        public readonly Outputs.ImageRawDiskResponse RawDisk;
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Set the secure boot keys of shielded instance.
        /// </summary>
        public readonly Outputs.InitialStateConfigResponse ShieldedInstanceInitialState;
        /// <summary>
        /// URL of the source disk used to create this image. This can be a full or valid partial URL. You must provide either this property or the rawDisk.source property but not both to create an image. For example, the following are valid values:  
        /// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk 
        /// - projects/project/zones/zone/disks/disk 
        /// - zones/zone/disks/disk
        /// </summary>
        public readonly string SourceDisk;
        /// <summary>
        /// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse SourceDiskEncryptionKey;
        /// <summary>
        /// The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
        /// </summary>
        public readonly string SourceDiskId;
        /// <summary>
        /// URL of the source image used to create this image.
        /// 
        /// In order to create an image, you must provide the full or partial URL of one of the following:  
        /// - The selfLink URL  
        /// - This property  
        /// - The rawDisk.source URL  
        /// - The sourceDisk URL
        /// </summary>
        public readonly string SourceImage;
        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse SourceImageEncryptionKey;
        /// <summary>
        /// The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
        /// </summary>
        public readonly string SourceImageId;
        /// <summary>
        /// URL of the source snapshot used to create this image.
        /// 
        /// In order to create an image, you must provide the full or partial URL of one of the following:  
        /// - The selfLink URL  
        /// - This property 
        /// - The sourceImage URL  
        /// - The rawDisk.source URL  
        /// - The sourceDisk URL
        /// </summary>
        public readonly string SourceSnapshot;
        /// <summary>
        /// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse SourceSnapshotEncryptionKey;
        /// <summary>
        /// The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
        /// </summary>
        public readonly string SourceSnapshotId;
        /// <summary>
        /// The type of the image used to create this disk. The default and only value is RAW
        /// </summary>
        public readonly string SourceType;
        /// <summary>
        /// The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Cloud Storage bucket storage location of the image (regional or multi-regional).
        /// </summary>
        public readonly ImmutableArray<string> StorageLocations;

        [OutputConstructor]
        private GetImageResult(
            string archiveSizeBytes,

            string creationTimestamp,

            Outputs.DeprecationStatusResponse deprecated,

            string description,

            string diskSizeGb,

            string family,

            ImmutableArray<Outputs.GuestOsFeatureResponse> guestOsFeatures,

            Outputs.CustomerEncryptionKeyResponse imageEncryptionKey,

            string kind,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<string> licenseCodes,

            ImmutableArray<string> licenses,

            string name,

            Outputs.ImageRawDiskResponse rawDisk,

            bool satisfiesPzs,

            string selfLink,

            Outputs.InitialStateConfigResponse shieldedInstanceInitialState,

            string sourceDisk,

            Outputs.CustomerEncryptionKeyResponse sourceDiskEncryptionKey,

            string sourceDiskId,

            string sourceImage,

            Outputs.CustomerEncryptionKeyResponse sourceImageEncryptionKey,

            string sourceImageId,

            string sourceSnapshot,

            Outputs.CustomerEncryptionKeyResponse sourceSnapshotEncryptionKey,

            string sourceSnapshotId,

            string sourceType,

            string status,

            ImmutableArray<string> storageLocations)
        {
            ArchiveSizeBytes = archiveSizeBytes;
            CreationTimestamp = creationTimestamp;
            Deprecated = deprecated;
            Description = description;
            DiskSizeGb = diskSizeGb;
            Family = family;
            GuestOsFeatures = guestOsFeatures;
            ImageEncryptionKey = imageEncryptionKey;
            Kind = kind;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            LicenseCodes = licenseCodes;
            Licenses = licenses;
            Name = name;
            RawDisk = rawDisk;
            SatisfiesPzs = satisfiesPzs;
            SelfLink = selfLink;
            ShieldedInstanceInitialState = shieldedInstanceInitialState;
            SourceDisk = sourceDisk;
            SourceDiskEncryptionKey = sourceDiskEncryptionKey;
            SourceDiskId = sourceDiskId;
            SourceImage = sourceImage;
            SourceImageEncryptionKey = sourceImageEncryptionKey;
            SourceImageId = sourceImageId;
            SourceSnapshot = sourceSnapshot;
            SourceSnapshotEncryptionKey = sourceSnapshotEncryptionKey;
            SourceSnapshotId = sourceSnapshotId;
            SourceType = sourceType;
            Status = status;
            StorageLocations = storageLocations;
        }
    }
}
