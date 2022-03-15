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
    /// Creates a persistent regional disk in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/beta:RegionDisk")]
    public partial class RegionDisk : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Encrypts the disk using a customer-supplied encryption key or a customer-managed encryption key. Encryption keys do not protect access to metadata of the disk. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later. For example, to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine. After you encrypt a disk with a customer-managed key, the diskEncryptionKey.kmsKeyName is set to a key *version* name once the disk is created. The disk is encrypted with this version of the key. In the response, diskEncryptionKey.kmsKeyName appears in the following format: "diskEncryptionKey.kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeysVersions/version If you do not provide an encryption key when creating the disk, then the disk is encrypted using an automatically generated key and you don't need to provide a key to use the disk later.
        /// </summary>
        [Output("diskEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> DiskEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the disk restored from a source snapshot should erase Windows specific VSS signature.
        /// </summary>
        [Output("eraseWindowsVssSignature")]
        public Output<bool> EraseWindowsVssSignature { get; private set; } = null!;

        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
        /// </summary>
        [Output("guestOsFeatures")]
        public Output<ImmutableArray<Outputs.GuestOsFeatureResponse>> GuestOsFeatures { get; private set; } = null!;

        /// <summary>
        /// [Deprecated] Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Type of the resource. Always compute#disk for disks.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a disk.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this disk. These can be later modified by the setLabels method.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Last attach timestamp in RFC3339 text format.
        /// </summary>
        [Output("lastAttachTimestamp")]
        public Output<string> LastAttachTimestamp { get; private set; } = null!;

        /// <summary>
        /// Last detach timestamp in RFC3339 text format.
        /// </summary>
        [Output("lastDetachTimestamp")]
        public Output<string> LastDetachTimestamp { get; private set; } = null!;

        /// <summary>
        /// Integer license codes indicating which licenses are attached to this disk.
        /// </summary>
        [Output("licenseCodes")]
        public Output<ImmutableArray<string>> LicenseCodes { get; private set; } = null!;

        /// <summary>
        /// A list of publicly visible licenses. Reserved for Google's use.
        /// </summary>
        [Output("licenses")]
        public Output<ImmutableArray<string>> Licenses { get; private set; } = null!;

        /// <summary>
        /// An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
        /// </summary>
        [Output("locationHint")]
        public Output<string> LocationHint { get; private set; } = null!;

        /// <summary>
        /// The field indicates if the disk is created from a locked source image. Attachment of a disk created from a locked source image will cause the following operations to become irreversibly prohibited: - R/W or R/O disk attachment to any other instance - Disk detachment. And the disk can only be deleted when the instance is deleted - Creation of images or snapshots - Disk cloning Furthermore, the instance with at least one disk with locked flag set to true will be prohibited from performing the operations below: - Further attachment of secondary disks. - Detachment of any disks - Create machine images - Create instance template - Delete the instance with --keep-disk parameter set to true for locked disks - Attach a locked disk with --auto-delete parameter set to false 
        /// </summary>
        [Output("locked")]
        public Output<bool> Locked { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not the disk can be read/write attached to more than one instance.
        /// </summary>
        [Output("multiWriter")]
        public Output<bool> MultiWriter { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Internal use only.
        /// </summary>
        [Output("options")]
        public Output<string> Options { get; private set; } = null!;

        /// <summary>
        /// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
        /// </summary>
        [Output("physicalBlockSizeBytes")]
        public Output<string> PhysicalBlockSizeBytes { get; private set; } = null!;

        /// <summary>
        /// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
        /// </summary>
        [Output("provisionedIops")]
        public Output<string> ProvisionedIops { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
        /// </summary>
        [Output("replicaZones")]
        public Output<ImmutableArray<string>> ReplicaZones { get; private set; } = null!;

        /// <summary>
        /// Resource policies applied to this disk for automatic snapshot creations.
        /// </summary>
        [Output("resourcePolicies")]
        public Output<ImmutableArray<string>> ResourcePolicies { get; private set; } = null!;

        /// <summary>
        /// Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// Server-defined fully-qualified URL for this resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
        /// </summary>
        [Output("sizeGb")]
        public Output<string> SizeGb { get; private set; } = null!;

        /// <summary>
        /// The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk 
        /// </summary>
        [Output("sourceDisk")]
        public Output<string> SourceDisk { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
        /// </summary>
        [Output("sourceDiskId")]
        public Output<string> SourceDiskId { get; private set; } = null!;

        /// <summary>
        /// The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family 
        /// </summary>
        [Output("sourceImage")]
        public Output<string> SourceImage { get; private set; } = null!;

        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
        /// </summary>
        [Output("sourceImageEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> SourceImageEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
        /// </summary>
        [Output("sourceImageId")]
        public Output<string> SourceImageId { get; private set; } = null!;

        /// <summary>
        /// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot 
        /// </summary>
        [Output("sourceSnapshot")]
        public Output<string> SourceSnapshot { get; private set; } = null!;

        /// <summary>
        /// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
        /// </summary>
        [Output("sourceSnapshotEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> SourceSnapshotEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
        /// </summary>
        [Output("sourceSnapshotId")]
        public Output<string> SourceSnapshotId { get; private set; } = null!;

        /// <summary>
        /// The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
        /// </summary>
        [Output("sourceStorageObject")]
        public Output<string> SourceStorageObject { get; private set; } = null!;

        /// <summary>
        /// The status of disk creation. - CREATING: Disk is provisioning. - RESTORING: Source data is being copied into the disk. - FAILED: Disk creation failed. - READY: Disk is ready for use. - DELETING: Disk is deleting. 
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// [Deprecated] Storage type of the persistent disk.
        /// </summary>
        [Output("storageType")]
        public Output<string> StorageType { get; private set; } = null!;

        /// <summary>
        /// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch 
        /// </summary>
        [Output("userLicenses")]
        public Output<ImmutableArray<string>> UserLicenses { get; private set; } = null!;

        /// <summary>
        /// Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;

        /// <summary>
        /// URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a RegionDisk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionDisk(string name, RegionDiskArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:RegionDisk", name, args ?? new RegionDiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionDisk(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:RegionDisk", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RegionDisk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionDisk Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegionDisk(name, id, options);
        }
    }

    public sealed class RegionDiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Encrypts the disk using a customer-supplied encryption key or a customer-managed encryption key. Encryption keys do not protect access to metadata of the disk. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later. For example, to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine. After you encrypt a disk with a customer-managed key, the diskEncryptionKey.kmsKeyName is set to a key *version* name once the disk is created. The disk is encrypted with this version of the key. In the response, diskEncryptionKey.kmsKeyName appears in the following format: "diskEncryptionKey.kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeysVersions/version If you do not provide an encryption key when creating the disk, then the disk is encrypted using an automatically generated key and you don't need to provide a key to use the disk later.
        /// </summary>
        [Input("diskEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? DiskEncryptionKey { get; set; }

        /// <summary>
        /// Specifies whether the disk restored from a source snapshot should erase Windows specific VSS signature.
        /// </summary>
        [Input("eraseWindowsVssSignature")]
        public Input<bool>? EraseWindowsVssSignature { get; set; }

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
        /// [Deprecated] Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
        /// </summary>
        [Input("interface")]
        public Input<Pulumi.GoogleNative.Compute.Beta.RegionDiskInterface>? Interface { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this disk. These can be later modified by the setLabels method.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("licenseCodes")]
        private InputList<string>? _licenseCodes;

        /// <summary>
        /// Integer license codes indicating which licenses are attached to this disk.
        /// </summary>
        public InputList<string> LicenseCodes
        {
            get => _licenseCodes ?? (_licenseCodes = new InputList<string>());
            set => _licenseCodes = value;
        }

        [Input("licenses")]
        private InputList<string>? _licenses;

        /// <summary>
        /// A list of publicly visible licenses. Reserved for Google's use.
        /// </summary>
        public InputList<string> Licenses
        {
            get => _licenses ?? (_licenses = new InputList<string>());
            set => _licenses = value;
        }

        /// <summary>
        /// An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
        /// </summary>
        [Input("locationHint")]
        public Input<string>? LocationHint { get; set; }

        /// <summary>
        /// Indicates whether or not the disk can be read/write attached to more than one instance.
        /// </summary>
        [Input("multiWriter")]
        public Input<bool>? MultiWriter { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Internal use only.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
        /// </summary>
        [Input("physicalBlockSizeBytes")]
        public Input<string>? PhysicalBlockSizeBytes { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
        /// </summary>
        [Input("provisionedIops")]
        public Input<string>? ProvisionedIops { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("replicaZones")]
        private InputList<string>? _replicaZones;

        /// <summary>
        /// URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
        /// </summary>
        public InputList<string> ReplicaZones
        {
            get => _replicaZones ?? (_replicaZones = new InputList<string>());
            set => _replicaZones = value;
        }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("resourcePolicies")]
        private InputList<string>? _resourcePolicies;

        /// <summary>
        /// Resource policies applied to this disk for automatic snapshot creations.
        /// </summary>
        public InputList<string> ResourcePolicies
        {
            get => _resourcePolicies ?? (_resourcePolicies = new InputList<string>());
            set => _resourcePolicies = value;
        }

        /// <summary>
        /// Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
        /// </summary>
        [Input("sizeGb")]
        public Input<string>? SizeGb { get; set; }

        /// <summary>
        /// The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk 
        /// </summary>
        [Input("sourceDisk")]
        public Input<string>? SourceDisk { get; set; }

        /// <summary>
        /// The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family 
        /// </summary>
        [Input("sourceImage")]
        public Input<string>? SourceImage { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
        /// </summary>
        [Input("sourceImageEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? SourceImageEncryptionKey { get; set; }

        /// <summary>
        /// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot 
        /// </summary>
        [Input("sourceSnapshot")]
        public Input<string>? SourceSnapshot { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
        /// </summary>
        [Input("sourceSnapshotEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? SourceSnapshotEncryptionKey { get; set; }

        /// <summary>
        /// The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
        /// </summary>
        [Input("sourceStorageObject")]
        public Input<string>? SourceStorageObject { get; set; }

        /// <summary>
        /// [Deprecated] Storage type of the persistent disk.
        /// </summary>
        [Input("storageType")]
        public Input<Pulumi.GoogleNative.Compute.Beta.RegionDiskStorageType>? StorageType { get; set; }

        /// <summary>
        /// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("userLicenses")]
        private InputList<string>? _userLicenses;

        /// <summary>
        /// A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch 
        /// </summary>
        public InputList<string> UserLicenses
        {
            get => _userLicenses ?? (_userLicenses = new InputList<string>());
            set => _userLicenses = value;
        }

        public RegionDiskArgs()
        {
        }
    }
}
