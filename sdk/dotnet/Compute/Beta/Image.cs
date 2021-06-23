// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    /// <summary>
    /// Creates an image in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/beta:Image")]
    public partial class Image : Pulumi.CustomResource
    {
        /// <summary>
        /// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
        /// </summary>
        [Output("archiveSizeBytes")]
        public Output<string> ArchiveSizeBytes { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// The deprecation status associated with this image.
        /// </summary>
        [Output("deprecated")]
        public Output<Outputs.DeprecationStatusResponse> Deprecated { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Size of the image when restored onto a persistent disk (in GB).
        /// </summary>
        [Output("diskSizeGb")]
        public Output<string> DiskSizeGb { get; private set; } = null!;

        /// <summary>
        /// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
        /// </summary>
        [Output("family")]
        public Output<string> Family { get; private set; } = null!;

        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
        /// </summary>
        [Output("guestOsFeatures")]
        public Output<ImmutableArray<Outputs.GuestOsFeatureResponse>> GuestOsFeatures { get; private set; } = null!;

        /// <summary>
        /// Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
        /// </summary>
        [Output("imageEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> ImageEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#image for images.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an image.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this image. These can be later modified by the setLabels method.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Integer license codes indicating which licenses are attached to this image.
        /// </summary>
        [Output("licenseCodes")]
        public Output<ImmutableArray<string>> LicenseCodes { get; private set; } = null!;

        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        [Output("licenses")]
        public Output<ImmutableArray<string>> Licenses { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parameters of the raw disk image.
        /// </summary>
        [Output("rawDisk")]
        public Output<Outputs.ImageRawDiskResponse> RawDisk { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Set the secure boot keys of shielded instance.
        /// </summary>
        [Output("shieldedInstanceInitialState")]
        public Output<Outputs.InitialStateConfigResponse> ShieldedInstanceInitialState { get; private set; } = null!;

        /// <summary>
        /// URL of the source disk used to create this image. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - projects/project/zones/zone/disks/disk - zones/zone/disks/disk In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        /// </summary>
        [Output("sourceDisk")]
        public Output<string> SourceDisk { get; private set; } = null!;

        /// <summary>
        /// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
        /// </summary>
        [Output("sourceDiskEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> SourceDiskEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
        /// </summary>
        [Output("sourceDiskId")]
        public Output<string> SourceDiskId { get; private set; } = null!;

        /// <summary>
        /// URL of the source image used to create this image. In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        /// </summary>
        [Output("sourceImage")]
        public Output<string> SourceImage { get; private set; } = null!;

        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
        /// </summary>
        [Output("sourceImageEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> SourceImageEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
        /// </summary>
        [Output("sourceImageId")]
        public Output<string> SourceImageId { get; private set; } = null!;

        /// <summary>
        /// URL of the source snapshot used to create this image. In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        /// </summary>
        [Output("sourceSnapshot")]
        public Output<string> SourceSnapshot { get; private set; } = null!;

        /// <summary>
        /// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
        /// </summary>
        [Output("sourceSnapshotEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> SourceSnapshotEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
        /// </summary>
        [Output("sourceSnapshotId")]
        public Output<string> SourceSnapshotId { get; private set; } = null!;

        /// <summary>
        /// The type of the image used to create this disk. The default and only value is RAW
        /// </summary>
        [Output("sourceType")]
        public Output<string> SourceType { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Cloud Storage bucket storage location of the image (regional or multi-regional).
        /// </summary>
        [Output("storageLocations")]
        public Output<ImmutableArray<string>> StorageLocations { get; private set; } = null!;


        /// <summary>
        /// Create a Image resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Image(string name, ImageArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:Image", name, args ?? new ImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Image(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:Image", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Image resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Image Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Image(name, id, options);
        }
    }

    public sealed class ImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
        /// </summary>
        [Input("archiveSizeBytes")]
        public Input<string>? ArchiveSizeBytes { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// The deprecation status associated with this image.
        /// </summary>
        [Input("deprecated")]
        public Input<Inputs.DeprecationStatusArgs>? Deprecated { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Size of the image when restored onto a persistent disk (in GB).
        /// </summary>
        [Input("diskSizeGb")]
        public Input<string>? DiskSizeGb { get; set; }

        /// <summary>
        /// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        [Input("forceCreate")]
        public Input<string>? ForceCreate { get; set; }

        [Input("guestOsFeatures")]
        private InputList<Inputs.GuestOsFeatureArgs>? _guestOsFeatures;

        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
        /// </summary>
        public InputList<Inputs.GuestOsFeatureArgs> GuestOsFeatures
        {
            get => _guestOsFeatures ?? (_guestOsFeatures = new InputList<Inputs.GuestOsFeatureArgs>());
            set => _guestOsFeatures = value;
        }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
        /// </summary>
        [Input("imageEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? ImageEncryptionKey { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#image for images.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an image.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this image. These can be later modified by the setLabels method.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("licenseCodes")]
        private InputList<string>? _licenseCodes;

        /// <summary>
        /// Integer license codes indicating which licenses are attached to this image.
        /// </summary>
        public InputList<string> LicenseCodes
        {
            get => _licenseCodes ?? (_licenseCodes = new InputList<string>());
            set => _licenseCodes = value;
        }

        [Input("licenses")]
        private InputList<string>? _licenses;

        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        public InputList<string> Licenses
        {
            get => _licenses ?? (_licenses = new InputList<string>());
            set => _licenses = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The parameters of the raw disk image.
        /// </summary>
        [Input("rawDisk")]
        public Input<Inputs.ImageRawDiskArgs>? RawDisk { get; set; }

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] Reserved for future use.
        /// </summary>
        [Input("satisfiesPzs")]
        public Input<bool>? SatisfiesPzs { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Set the secure boot keys of shielded instance.
        /// </summary>
        [Input("shieldedInstanceInitialState")]
        public Input<Inputs.InitialStateConfigArgs>? ShieldedInstanceInitialState { get; set; }

        /// <summary>
        /// URL of the source disk used to create this image. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - projects/project/zones/zone/disks/disk - zones/zone/disks/disk In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        /// </summary>
        [Input("sourceDisk")]
        public Input<string>? SourceDisk { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
        /// </summary>
        [Input("sourceDiskEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? SourceDiskEncryptionKey { get; set; }

        /// <summary>
        /// [Output Only] The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
        /// </summary>
        [Input("sourceDiskId")]
        public Input<string>? SourceDiskId { get; set; }

        /// <summary>
        /// URL of the source image used to create this image. In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        /// </summary>
        [Input("sourceImage")]
        public Input<string>? SourceImage { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
        /// </summary>
        [Input("sourceImageEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? SourceImageEncryptionKey { get; set; }

        /// <summary>
        /// [Output Only] The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
        /// </summary>
        [Input("sourceImageId")]
        public Input<string>? SourceImageId { get; set; }

        /// <summary>
        /// URL of the source snapshot used to create this image. In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        /// </summary>
        [Input("sourceSnapshot")]
        public Input<string>? SourceSnapshot { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
        /// </summary>
        [Input("sourceSnapshotEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? SourceSnapshotEncryptionKey { get; set; }

        /// <summary>
        /// [Output Only] The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
        /// </summary>
        [Input("sourceSnapshotId")]
        public Input<string>? SourceSnapshotId { get; set; }

        /// <summary>
        /// The type of the image used to create this disk. The default and only value is RAW
        /// </summary>
        [Input("sourceType")]
        public Input<Pulumi.GoogleNative.Compute.Beta.ImageSourceType>? SourceType { get; set; }

        /// <summary>
        /// [Output Only] The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
        /// </summary>
        [Input("status")]
        public Input<Pulumi.GoogleNative.Compute.Beta.ImageStatus>? Status { get; set; }

        [Input("storageLocations")]
        private InputList<string>? _storageLocations;

        /// <summary>
        /// Cloud Storage bucket storage location of the image (regional or multi-regional).
        /// </summary>
        public InputList<string> StorageLocations
        {
            get => _storageLocations ?? (_storageLocations = new InputList<string>());
            set => _storageLocations = value;
        }

        public ImageArgs()
        {
        }
    }
}
