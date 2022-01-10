// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a persistent regional disk in the specified project using the data included in the request.
type RegionDisk struct {
	pulumi.CustomResourceState

	// The architecture of the disk. Valid values are ARM64 or X86_64.
	Architecture pulumi.StringOutput `pulumi:"architecture"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Encrypts the disk using a customer-supplied encryption key or a customer-managed encryption key. Encryption keys do not protect access to metadata of the disk. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later. For example, to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine. After you encrypt a disk with a customer-managed key, the diskEncryptionKey.kmsKeyName is set to a key *version* name once the disk is created. The disk is encrypted with this version of the key. In the response, diskEncryptionKey.kmsKeyName appears in the following format: "diskEncryptionKey.kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeysVersions/version If you do not provide an encryption key when creating the disk, then the disk is encrypted using an automatically generated key and you don't need to provide a key to use the disk later.
	DiskEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"diskEncryptionKey"`
	// Specifies whether the disk restored from a source snapshot should erase Windows specific VSS signature.
	EraseWindowsVssSignature pulumi.BoolOutput `pulumi:"eraseWindowsVssSignature"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
	GuestOsFeatures GuestOsFeatureResponseArrayOutput `pulumi:"guestOsFeatures"`
	// Type of the resource. Always compute#disk for disks.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a disk.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this disk. These can be later modified by the setLabels method.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp pulumi.StringOutput `pulumi:"lastAttachTimestamp"`
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp pulumi.StringOutput `pulumi:"lastDetachTimestamp"`
	// Integer license codes indicating which licenses are attached to this disk.
	LicenseCodes pulumi.StringArrayOutput `pulumi:"licenseCodes"`
	// A list of publicly visible licenses. Reserved for Google's use.
	Licenses pulumi.StringArrayOutput `pulumi:"licenses"`
	// An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
	LocationHint pulumi.StringOutput `pulumi:"locationHint"`
	// The field indicates if the disk is created from a locked source image. Attachment of a disk created from a locked source image will cause the following operations to become irreversibly prohibited: - R/W or R/O disk attachment to any other instance - Disk detachment. And the disk can only be deleted when the instance is deleted - Creation of images or snapshots - Disk cloning Furthermore, the instance with at least one disk with locked flag set to true will be prohibited from performing the operations below: - Further attachment of secondary disks. - Detachment of any disks - Create machine images - Create instance template - Delete the instance with --keep-disk parameter set to true for locked disks - Attach a locked disk with --auto-delete parameter set to false
	Locked pulumi.BoolOutput `pulumi:"locked"`
	// Indicates whether or not the disk can be read/write attached to more than one instance.
	MultiWriter pulumi.BoolOutput `pulumi:"multiWriter"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Internal use only.
	Options pulumi.StringOutput `pulumi:"options"`
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.StringOutput `pulumi:"physicalBlockSizeBytes"`
	// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
	ProvisionedIops pulumi.StringOutput `pulumi:"provisionedIops"`
	// URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
	ReplicaZones pulumi.StringArrayOutput `pulumi:"replicaZones"`
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies pulumi.StringArrayOutput `pulumi:"resourcePolicies"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource's resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
	SizeGb pulumi.StringOutput `pulumi:"sizeGb"`
	// The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk
	SourceDisk pulumi.StringOutput `pulumi:"sourceDisk"`
	// The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
	SourceDiskId pulumi.StringOutput `pulumi:"sourceDiskId"`
	// The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family
	SourceImage pulumi.StringOutput `pulumi:"sourceImage"`
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"sourceImageEncryptionKey"`
	// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
	SourceImageId pulumi.StringOutput `pulumi:"sourceImageId"`
	// The source instant snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot
	SourceInstantSnapshot pulumi.StringOutput `pulumi:"sourceInstantSnapshot"`
	// The unique ID of the instant snapshot used to create this disk. This value identifies the exact instant snapshot that was used to create this persistent disk. For example, if you created the persistent disk from an instant snapshot that was later deleted and recreated under the same name, the source instant snapshot ID would identify the exact version of the instant snapshot that was used.
	SourceInstantSnapshotId pulumi.StringOutput `pulumi:"sourceInstantSnapshotId"`
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot
	SourceSnapshot pulumi.StringOutput `pulumi:"sourceSnapshot"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
	SourceSnapshotEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"sourceSnapshotEncryptionKey"`
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId pulumi.StringOutput `pulumi:"sourceSnapshotId"`
	// The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
	SourceStorageObject pulumi.StringOutput `pulumi:"sourceStorageObject"`
	// The status of disk creation. - CREATING: Disk is provisioning. - RESTORING: Source data is being copied into the disk. - FAILED: Disk creation failed. - READY: Disk is ready for use. - DELETING: Disk is deleting.
	Status pulumi.StringOutput `pulumi:"status"`
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch
	UserLicenses pulumi.StringArrayOutput `pulumi:"userLicenses"`
	// Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
	Users pulumi.StringArrayOutput `pulumi:"users"`
	// URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewRegionDisk registers a new resource with the given unique name, arguments, and options.
func NewRegionDisk(ctx *pulumi.Context,
	name string, args *RegionDiskArgs, opts ...pulumi.ResourceOption) (*RegionDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionDisk
	err := ctx.RegisterResource("google-native:compute/alpha:RegionDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionDisk gets an existing RegionDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionDiskState, opts ...pulumi.ResourceOption) (*RegionDisk, error) {
	var resource RegionDisk
	err := ctx.ReadResource("google-native:compute/alpha:RegionDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionDisk resources.
type regionDiskState struct {
}

type RegionDiskState struct {
}

func (RegionDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskState)(nil)).Elem()
}

type regionDiskArgs struct {
	// The architecture of the disk. Valid values are ARM64 or X86_64.
	Architecture *RegionDiskArchitecture `pulumi:"architecture"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Encrypts the disk using a customer-supplied encryption key or a customer-managed encryption key. Encryption keys do not protect access to metadata of the disk. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later. For example, to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine. After you encrypt a disk with a customer-managed key, the diskEncryptionKey.kmsKeyName is set to a key *version* name once the disk is created. The disk is encrypted with this version of the key. In the response, diskEncryptionKey.kmsKeyName appears in the following format: "diskEncryptionKey.kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeysVersions/version If you do not provide an encryption key when creating the disk, then the disk is encrypted using an automatically generated key and you don't need to provide a key to use the disk later.
	DiskEncryptionKey *CustomerEncryptionKey `pulumi:"diskEncryptionKey"`
	// Specifies whether the disk restored from a source snapshot should erase Windows specific VSS signature.
	EraseWindowsVssSignature *bool `pulumi:"eraseWindowsVssSignature"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
	GuestOsFeatures []GuestOsFeature `pulumi:"guestOsFeatures"`
	// Labels to apply to this disk. These can be later modified by the setLabels method.
	Labels map[string]string `pulumi:"labels"`
	// Integer license codes indicating which licenses are attached to this disk.
	LicenseCodes []string `pulumi:"licenseCodes"`
	// A list of publicly visible licenses. Reserved for Google's use.
	Licenses []string `pulumi:"licenses"`
	// An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
	LocationHint *string `pulumi:"locationHint"`
	// Indicates whether or not the disk can be read/write attached to more than one instance.
	MultiWriter *bool `pulumi:"multiWriter"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Internal use only.
	Options *string `pulumi:"options"`
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes *string `pulumi:"physicalBlockSizeBytes"`
	Project                *string `pulumi:"project"`
	// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
	ProvisionedIops *string `pulumi:"provisionedIops"`
	Region          string  `pulumi:"region"`
	// URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
	ReplicaZones []string `pulumi:"replicaZones"`
	RequestId    *string  `pulumi:"requestId"`
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies []string `pulumi:"resourcePolicies"`
	// Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
	SizeGb *string `pulumi:"sizeGb"`
	// The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk
	SourceDisk *string `pulumi:"sourceDisk"`
	// The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family
	SourceImage *string `pulumi:"sourceImage"`
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey *CustomerEncryptionKey `pulumi:"sourceImageEncryptionKey"`
	// The source instant snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot
	SourceInstantSnapshot *string `pulumi:"sourceInstantSnapshot"`
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
	SourceSnapshotEncryptionKey *CustomerEncryptionKey `pulumi:"sourceSnapshotEncryptionKey"`
	// The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
	SourceStorageObject *string `pulumi:"sourceStorageObject"`
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
	Type *string `pulumi:"type"`
	// A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch
	UserLicenses []string `pulumi:"userLicenses"`
}

// The set of arguments for constructing a RegionDisk resource.
type RegionDiskArgs struct {
	// The architecture of the disk. Valid values are ARM64 or X86_64.
	Architecture RegionDiskArchitecturePtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Encrypts the disk using a customer-supplied encryption key or a customer-managed encryption key. Encryption keys do not protect access to metadata of the disk. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later. For example, to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine. After you encrypt a disk with a customer-managed key, the diskEncryptionKey.kmsKeyName is set to a key *version* name once the disk is created. The disk is encrypted with this version of the key. In the response, diskEncryptionKey.kmsKeyName appears in the following format: "diskEncryptionKey.kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeysVersions/version If you do not provide an encryption key when creating the disk, then the disk is encrypted using an automatically generated key and you don't need to provide a key to use the disk later.
	DiskEncryptionKey CustomerEncryptionKeyPtrInput
	// Specifies whether the disk restored from a source snapshot should erase Windows specific VSS signature.
	EraseWindowsVssSignature pulumi.BoolPtrInput
	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
	GuestOsFeatures GuestOsFeatureArrayInput
	// Labels to apply to this disk. These can be later modified by the setLabels method.
	Labels pulumi.StringMapInput
	// Integer license codes indicating which licenses are attached to this disk.
	LicenseCodes pulumi.StringArrayInput
	// A list of publicly visible licenses. Reserved for Google's use.
	Licenses pulumi.StringArrayInput
	// An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
	LocationHint pulumi.StringPtrInput
	// Indicates whether or not the disk can be read/write attached to more than one instance.
	MultiWriter pulumi.BoolPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Internal use only.
	Options pulumi.StringPtrInput
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.StringPtrInput
	Project                pulumi.StringPtrInput
	// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
	ProvisionedIops pulumi.StringPtrInput
	Region          pulumi.StringInput
	// URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
	ReplicaZones pulumi.StringArrayInput
	RequestId    pulumi.StringPtrInput
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies pulumi.StringArrayInput
	// Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
	SizeGb pulumi.StringPtrInput
	// The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk
	SourceDisk pulumi.StringPtrInput
	// The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family
	SourceImage pulumi.StringPtrInput
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey CustomerEncryptionKeyPtrInput
	// The source instant snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot
	SourceInstantSnapshot pulumi.StringPtrInput
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot
	SourceSnapshot pulumi.StringPtrInput
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
	SourceSnapshotEncryptionKey CustomerEncryptionKeyPtrInput
	// The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
	SourceStorageObject pulumi.StringPtrInput
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
	Type pulumi.StringPtrInput
	// A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch
	UserLicenses pulumi.StringArrayInput
}

func (RegionDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskArgs)(nil)).Elem()
}

type RegionDiskInput interface {
	pulumi.Input

	ToRegionDiskOutput() RegionDiskOutput
	ToRegionDiskOutputWithContext(ctx context.Context) RegionDiskOutput
}

func (*RegionDisk) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionDisk)(nil)).Elem()
}

func (i *RegionDisk) ToRegionDiskOutput() RegionDiskOutput {
	return i.ToRegionDiskOutputWithContext(context.Background())
}

func (i *RegionDisk) ToRegionDiskOutputWithContext(ctx context.Context) RegionDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskOutput)
}

type RegionDiskOutput struct{ *pulumi.OutputState }

func (RegionDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionDisk)(nil)).Elem()
}

func (o RegionDiskOutput) ToRegionDiskOutput() RegionDiskOutput {
	return o
}

func (o RegionDiskOutput) ToRegionDiskOutputWithContext(ctx context.Context) RegionDiskOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionDiskInput)(nil)).Elem(), &RegionDisk{})
	pulumi.RegisterOutputType(RegionDiskOutput{})
}
