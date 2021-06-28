// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an image in the specified project using the data included in the request.
type Image struct {
	pulumi.CustomResourceState

	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes pulumi.StringOutput `pulumi:"archiveSizeBytes"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// The deprecation status associated with this image.
	Deprecated DeprecationStatusResponseOutput `pulumi:"deprecated"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb pulumi.StringOutput `pulumi:"diskSizeGb"`
	// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
	Family pulumi.StringOutput `pulumi:"family"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
	GuestOsFeatures GuestOsFeatureResponseArrayOutput `pulumi:"guestOsFeatures"`
	// Encrypts the image using a customer-supplied encryption key.
	//
	// After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image).
	//
	// Customer-supplied encryption keys do not protect access to metadata of the disk.
	//
	// If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
	ImageEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"imageEncryptionKey"`
	// Type of the resource. Always compute#image for images.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve an image.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this image. These can be later modified by the setLabels method.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Integer license codes indicating which licenses are attached to this image.
	LicenseCodes pulumi.StringArrayOutput `pulumi:"licenseCodes"`
	// Any applicable license URI.
	Licenses pulumi.StringArrayOutput `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters of the raw disk image.
	RawDisk ImageRawDiskResponseOutput `pulumi:"rawDisk"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Set the secure boot keys of shielded instance.
	ShieldedInstanceInitialState InitialStateConfigResponseOutput `pulumi:"shieldedInstanceInitialState"`
	// URL of the source disk used to create this image. This can be a full or valid partial URL. You must provide either this property or the rawDisk.source property but not both to create an image. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
	// - projects/project/zones/zone/disks/disk
	// - zones/zone/disks/disk
	SourceDisk pulumi.StringOutput `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"sourceDiskEncryptionKey"`
	// The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
	SourceDiskId pulumi.StringOutput `pulumi:"sourceDiskId"`
	// URL of the source image used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// - The selfLink URL
	// - This property
	// - The rawDisk.source URL
	// - The sourceDisk URL
	SourceImage pulumi.StringOutput `pulumi:"sourceImage"`
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"sourceImageEncryptionKey"`
	// The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
	SourceImageId pulumi.StringOutput `pulumi:"sourceImageId"`
	// URL of the source snapshot used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// - The selfLink URL
	// - This property
	// - The sourceImage URL
	// - The rawDisk.source URL
	// - The sourceDisk URL
	SourceSnapshot pulumi.StringOutput `pulumi:"sourceSnapshot"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
	SourceSnapshotEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"sourceSnapshotEncryptionKey"`
	// The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
	SourceSnapshotId pulumi.StringOutput `pulumi:"sourceSnapshotId"`
	// The type of the image used to create this disk. The default and only value is RAW
	SourceType pulumi.StringOutput `pulumi:"sourceType"`
	// The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
	Status pulumi.StringOutput `pulumi:"status"`
	// Cloud Storage bucket storage location of the image (regional or multi-regional).
	StorageLocations pulumi.StringArrayOutput `pulumi:"storageLocations"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Image
	err := ctx.RegisterResource("google-native:compute/beta:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("google-native:compute/beta:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes *string `pulumi:"archiveSizeBytes"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// The deprecation status associated with this image.
	Deprecated *DeprecationStatusResponse `pulumi:"deprecated"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb *string `pulumi:"diskSizeGb"`
	// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
	Family *string `pulumi:"family"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
	GuestOsFeatures []GuestOsFeatureResponse `pulumi:"guestOsFeatures"`
	// Encrypts the image using a customer-supplied encryption key.
	//
	// After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image).
	//
	// Customer-supplied encryption keys do not protect access to metadata of the disk.
	//
	// If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
	ImageEncryptionKey *CustomerEncryptionKeyResponse `pulumi:"imageEncryptionKey"`
	// Type of the resource. Always compute#image for images.
	Kind *string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve an image.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this image. These can be later modified by the setLabels method.
	Labels map[string]string `pulumi:"labels"`
	// Integer license codes indicating which licenses are attached to this image.
	LicenseCodes []string `pulumi:"licenseCodes"`
	// Any applicable license URI.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The parameters of the raw disk image.
	RawDisk *ImageRawDiskResponse `pulumi:"rawDisk"`
	// Reserved for future use.
	SatisfiesPzs *bool `pulumi:"satisfiesPzs"`
	// Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// Set the secure boot keys of shielded instance.
	ShieldedInstanceInitialState *InitialStateConfigResponse `pulumi:"shieldedInstanceInitialState"`
	// URL of the source disk used to create this image. This can be a full or valid partial URL. You must provide either this property or the rawDisk.source property but not both to create an image. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
	// - projects/project/zones/zone/disks/disk
	// - zones/zone/disks/disk
	SourceDisk *string `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKey *CustomerEncryptionKeyResponse `pulumi:"sourceDiskEncryptionKey"`
	// The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
	SourceDiskId *string `pulumi:"sourceDiskId"`
	// URL of the source image used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// - The selfLink URL
	// - This property
	// - The rawDisk.source URL
	// - The sourceDisk URL
	SourceImage *string `pulumi:"sourceImage"`
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey *CustomerEncryptionKeyResponse `pulumi:"sourceImageEncryptionKey"`
	// The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
	SourceImageId *string `pulumi:"sourceImageId"`
	// URL of the source snapshot used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// - The selfLink URL
	// - This property
	// - The sourceImage URL
	// - The rawDisk.source URL
	// - The sourceDisk URL
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
	SourceSnapshotEncryptionKey *CustomerEncryptionKeyResponse `pulumi:"sourceSnapshotEncryptionKey"`
	// The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
	SourceSnapshotId *string `pulumi:"sourceSnapshotId"`
	// The type of the image used to create this disk. The default and only value is RAW
	SourceType *string `pulumi:"sourceType"`
	// The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
	Status *string `pulumi:"status"`
	// Cloud Storage bucket storage location of the image (regional or multi-regional).
	StorageLocations []string `pulumi:"storageLocations"`
}

type ImageState struct {
	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// The deprecation status associated with this image.
	Deprecated DeprecationStatusResponsePtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb pulumi.StringPtrInput
	// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
	Family pulumi.StringPtrInput
	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
	GuestOsFeatures GuestOsFeatureResponseArrayInput
	// Encrypts the image using a customer-supplied encryption key.
	//
	// After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image).
	//
	// Customer-supplied encryption keys do not protect access to metadata of the disk.
	//
	// If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
	ImageEncryptionKey CustomerEncryptionKeyResponsePtrInput
	// Type of the resource. Always compute#image for images.
	Kind pulumi.StringPtrInput
	// A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve an image.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this image. These can be later modified by the setLabels method.
	Labels pulumi.StringMapInput
	// Integer license codes indicating which licenses are attached to this image.
	LicenseCodes pulumi.StringArrayInput
	// Any applicable license URI.
	Licenses pulumi.StringArrayInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The parameters of the raw disk image.
	RawDisk ImageRawDiskResponsePtrInput
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolPtrInput
	// Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// Set the secure boot keys of shielded instance.
	ShieldedInstanceInitialState InitialStateConfigResponsePtrInput
	// URL of the source disk used to create this image. This can be a full or valid partial URL. You must provide either this property or the rawDisk.source property but not both to create an image. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
	// - projects/project/zones/zone/disks/disk
	// - zones/zone/disks/disk
	SourceDisk pulumi.StringPtrInput
	// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKey CustomerEncryptionKeyResponsePtrInput
	// The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
	SourceDiskId pulumi.StringPtrInput
	// URL of the source image used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// - The selfLink URL
	// - This property
	// - The rawDisk.source URL
	// - The sourceDisk URL
	SourceImage pulumi.StringPtrInput
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey CustomerEncryptionKeyResponsePtrInput
	// The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
	SourceImageId pulumi.StringPtrInput
	// URL of the source snapshot used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// - The selfLink URL
	// - This property
	// - The sourceImage URL
	// - The rawDisk.source URL
	// - The sourceDisk URL
	SourceSnapshot pulumi.StringPtrInput
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
	SourceSnapshotEncryptionKey CustomerEncryptionKeyResponsePtrInput
	// The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
	SourceSnapshotId pulumi.StringPtrInput
	// The type of the image used to create this disk. The default and only value is RAW
	SourceType pulumi.StringPtrInput
	// The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
	Status pulumi.StringPtrInput
	// Cloud Storage bucket storage location of the image (regional or multi-regional).
	StorageLocations pulumi.StringArrayInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes *string `pulumi:"archiveSizeBytes"`
	// The deprecation status associated with this image.
	Deprecated *DeprecationStatus `pulumi:"deprecated"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb *string `pulumi:"diskSizeGb"`
	// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
	Family      *string `pulumi:"family"`
	ForceCreate *string `pulumi:"forceCreate"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
	GuestOsFeatures []GuestOsFeature `pulumi:"guestOsFeatures"`
	// Encrypts the image using a customer-supplied encryption key.
	//
	// After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image).
	//
	// Customer-supplied encryption keys do not protect access to metadata of the disk.
	//
	// If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
	ImageEncryptionKey *CustomerEncryptionKey `pulumi:"imageEncryptionKey"`
	// Labels to apply to this image. These can be later modified by the setLabels method.
	Labels map[string]string `pulumi:"labels"`
	// Integer license codes indicating which licenses are attached to this image.
	LicenseCodes []string `pulumi:"licenseCodes"`
	// Any applicable license URI.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// The parameters of the raw disk image.
	RawDisk   *ImageRawDisk `pulumi:"rawDisk"`
	RequestId *string       `pulumi:"requestId"`
	// Set the secure boot keys of shielded instance.
	ShieldedInstanceInitialState *InitialStateConfig `pulumi:"shieldedInstanceInitialState"`
	// URL of the source disk used to create this image. This can be a full or valid partial URL. You must provide either this property or the rawDisk.source property but not both to create an image. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
	// - projects/project/zones/zone/disks/disk
	// - zones/zone/disks/disk
	SourceDisk *string `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKey *CustomerEncryptionKey `pulumi:"sourceDiskEncryptionKey"`
	// URL of the source image used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// - The selfLink URL
	// - This property
	// - The rawDisk.source URL
	// - The sourceDisk URL
	SourceImage *string `pulumi:"sourceImage"`
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey *CustomerEncryptionKey `pulumi:"sourceImageEncryptionKey"`
	// URL of the source snapshot used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// - The selfLink URL
	// - This property
	// - The sourceImage URL
	// - The rawDisk.source URL
	// - The sourceDisk URL
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
	SourceSnapshotEncryptionKey *CustomerEncryptionKey `pulumi:"sourceSnapshotEncryptionKey"`
	// The type of the image used to create this disk. The default and only value is RAW
	SourceType *string `pulumi:"sourceType"`
	// Cloud Storage bucket storage location of the image (regional or multi-regional).
	StorageLocations []string `pulumi:"storageLocations"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes pulumi.StringPtrInput
	// The deprecation status associated with this image.
	Deprecated DeprecationStatusPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb pulumi.StringPtrInput
	// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
	Family      pulumi.StringPtrInput
	ForceCreate pulumi.StringPtrInput
	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
	GuestOsFeatures GuestOsFeatureArrayInput
	// Encrypts the image using a customer-supplied encryption key.
	//
	// After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image).
	//
	// Customer-supplied encryption keys do not protect access to metadata of the disk.
	//
	// If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
	ImageEncryptionKey CustomerEncryptionKeyPtrInput
	// Labels to apply to this image. These can be later modified by the setLabels method.
	Labels pulumi.StringMapInput
	// Integer license codes indicating which licenses are attached to this image.
	LicenseCodes pulumi.StringArrayInput
	// Any applicable license URI.
	Licenses pulumi.StringArrayInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// The parameters of the raw disk image.
	RawDisk   ImageRawDiskPtrInput
	RequestId pulumi.StringPtrInput
	// Set the secure boot keys of shielded instance.
	ShieldedInstanceInitialState InitialStateConfigPtrInput
	// URL of the source disk used to create this image. This can be a full or valid partial URL. You must provide either this property or the rawDisk.source property but not both to create an image. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
	// - projects/project/zones/zone/disks/disk
	// - zones/zone/disks/disk
	SourceDisk pulumi.StringPtrInput
	// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKey CustomerEncryptionKeyPtrInput
	// URL of the source image used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// - The selfLink URL
	// - This property
	// - The rawDisk.source URL
	// - The sourceDisk URL
	SourceImage pulumi.StringPtrInput
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey CustomerEncryptionKeyPtrInput
	// URL of the source snapshot used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// - The selfLink URL
	// - This property
	// - The sourceImage URL
	// - The rawDisk.source URL
	// - The sourceDisk URL
	SourceSnapshot pulumi.StringPtrInput
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
	SourceSnapshotEncryptionKey CustomerEncryptionKeyPtrInput
	// The type of the image used to create this disk. The default and only value is RAW
	SourceType *ImageSourceType
	// Cloud Storage bucket storage location of the image (regional or multi-regional).
	StorageLocations pulumi.StringArrayInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil))
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct {
	*pulumi.OutputState
}

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil))
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ImageOutput{})
}
