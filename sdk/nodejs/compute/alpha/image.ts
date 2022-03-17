// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an image in the specified project using the data included in the request.
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    /**
     * The architecture of the image. Valid values are ARM64 or X86_64.
     */
    public readonly architecture!: pulumi.Output<string>;
    /**
     * Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
     */
    public readonly archiveSizeBytes!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * The deprecation status associated with this image.
     */
    public readonly deprecated!: pulumi.Output<outputs.compute.alpha.DeprecationStatusResponse>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Size of the image when restored onto a persistent disk (in GB).
     */
    public readonly diskSizeGb!: pulumi.Output<string>;
    /**
     * The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
     */
    public readonly family!: pulumi.Output<string>;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. To see a list of available options, see the guestOSfeatures[].type parameter.
     */
    public readonly guestOsFeatures!: pulumi.Output<outputs.compute.alpha.GuestOsFeatureResponse[]>;
    /**
     * Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
     */
    public readonly imageEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * Type of the resource. Always compute#image for images.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an image.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this image. These can be later modified by the setLabels method.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Integer license codes indicating which licenses are attached to this image.
     */
    public readonly licenseCodes!: pulumi.Output<string[]>;
    /**
     * Any applicable license URI.
     */
    public readonly licenses!: pulumi.Output<string[]>;
    /**
     * A flag for marketplace VM disk created from the image, which is designed for marketplace VM disk to prevent the proprietary data on the disk from being accessed unwantedly. The flag will be inherited by the disk created from the image. The disk with locked flag set to true will be prohibited from performing the operations below: - R/W or R/O disk attach - Disk detach, if disk is created via create-on-create - Create images - Create snapshots - Create disk clone (create disk from the current disk) The image with the locked field set to true will be prohibited from performing the operations below: - Create images from the current image - Update the locked field for the current image The instance with at least one disk with locked flag set to true will be prohibited from performing the operations below: - Secondary disk attach - Create instant snapshot - Create machine images - Create instance template - Delete the instance with --keep-disk parameter set to true 
     */
    public readonly locked!: pulumi.Output<boolean>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The parameters of the raw disk image.
     */
    public readonly rawDisk!: pulumi.Output<outputs.compute.alpha.ImageRawDiskResponse>;
    /**
     * A rollout policy to apply to this image. When specified, the rollout policy overrides per-zone references to the image via the associated image family. The rollout policy restricts the zones where this image is accessible when using a zonal image family reference. When the rollout policy does not include the user specified zone, or if the zone is rolled out, this image is accessible. The rollout policy for this image is read-only, except for allowlisted users. This field might not be configured. To view the latest non-deprecated image in a specific zone, use the imageFamilyViews.get method.
     */
    public readonly rolloutOverride!: pulumi.Output<outputs.compute.alpha.RolloutPolicyResponse>;
    /**
     * Reserved for future use.
     */
    public /*out*/ readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource's resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * Set the secure boot keys of shielded instance.
     */
    public readonly shieldedInstanceInitialState!: pulumi.Output<outputs.compute.alpha.InitialStateConfigResponse>;
    /**
     * URL of the source disk used to create this image. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - projects/project/zones/zone/disks/disk - zones/zone/disks/disk In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
     */
    public readonly sourceDisk!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
     */
    public readonly sourceDiskEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
     */
    public /*out*/ readonly sourceDiskId!: pulumi.Output<string>;
    /**
     * URL of the source image used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ images/image_name - projects/project_id/global/images/image_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
     */
    public readonly sourceImage!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    public readonly sourceImageEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
     */
    public /*out*/ readonly sourceImageId!: pulumi.Output<string>;
    /**
     * URL of the source snapshot used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ snapshots/snapshot_name - projects/project_id/global/snapshots/snapshot_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
     */
    public readonly sourceSnapshot!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    public readonly sourceSnapshotEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
     */
    public /*out*/ readonly sourceSnapshotId!: pulumi.Output<string>;
    /**
     * The type of the image used to create this disk. The default and only value is RAW
     */
    public readonly sourceType!: pulumi.Output<string>;
    /**
     * The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Cloud Storage bucket storage location of the image (regional or multi-regional).
     */
    public readonly storageLocations!: pulumi.Output<string[]>;
    /**
     * A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch 
     */
    public readonly userLicenses!: pulumi.Output<string[]>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ImageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["architecture"] = args ? args.architecture : undefined;
            resourceInputs["archiveSizeBytes"] = args ? args.archiveSizeBytes : undefined;
            resourceInputs["deprecated"] = args ? args.deprecated : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["diskSizeGb"] = args ? args.diskSizeGb : undefined;
            resourceInputs["family"] = args ? args.family : undefined;
            resourceInputs["forceCreate"] = args ? args.forceCreate : undefined;
            resourceInputs["guestOsFeatures"] = args ? args.guestOsFeatures : undefined;
            resourceInputs["imageEncryptionKey"] = args ? args.imageEncryptionKey : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["licenseCodes"] = args ? args.licenseCodes : undefined;
            resourceInputs["licenses"] = args ? args.licenses : undefined;
            resourceInputs["locked"] = args ? args.locked : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["rawDisk"] = args ? args.rawDisk : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["rolloutOverride"] = args ? args.rolloutOverride : undefined;
            resourceInputs["shieldedInstanceInitialState"] = args ? args.shieldedInstanceInitialState : undefined;
            resourceInputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            resourceInputs["sourceDiskEncryptionKey"] = args ? args.sourceDiskEncryptionKey : undefined;
            resourceInputs["sourceImage"] = args ? args.sourceImage : undefined;
            resourceInputs["sourceImageEncryptionKey"] = args ? args.sourceImageEncryptionKey : undefined;
            resourceInputs["sourceSnapshot"] = args ? args.sourceSnapshot : undefined;
            resourceInputs["sourceSnapshotEncryptionKey"] = args ? args.sourceSnapshotEncryptionKey : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["storageLocations"] = args ? args.storageLocations : undefined;
            resourceInputs["userLicenses"] = args ? args.userLicenses : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["sourceDiskId"] = undefined /*out*/;
            resourceInputs["sourceImageId"] = undefined /*out*/;
            resourceInputs["sourceSnapshotId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["archiveSizeBytes"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["deprecated"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["diskSizeGb"] = undefined /*out*/;
            resourceInputs["family"] = undefined /*out*/;
            resourceInputs["guestOsFeatures"] = undefined /*out*/;
            resourceInputs["imageEncryptionKey"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["licenseCodes"] = undefined /*out*/;
            resourceInputs["licenses"] = undefined /*out*/;
            resourceInputs["locked"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["rawDisk"] = undefined /*out*/;
            resourceInputs["rolloutOverride"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["shieldedInstanceInitialState"] = undefined /*out*/;
            resourceInputs["sourceDisk"] = undefined /*out*/;
            resourceInputs["sourceDiskEncryptionKey"] = undefined /*out*/;
            resourceInputs["sourceDiskId"] = undefined /*out*/;
            resourceInputs["sourceImage"] = undefined /*out*/;
            resourceInputs["sourceImageEncryptionKey"] = undefined /*out*/;
            resourceInputs["sourceImageId"] = undefined /*out*/;
            resourceInputs["sourceSnapshot"] = undefined /*out*/;
            resourceInputs["sourceSnapshotEncryptionKey"] = undefined /*out*/;
            resourceInputs["sourceSnapshotId"] = undefined /*out*/;
            resourceInputs["sourceType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["storageLocations"] = undefined /*out*/;
            resourceInputs["userLicenses"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * The architecture of the image. Valid values are ARM64 or X86_64.
     */
    architecture?: pulumi.Input<enums.compute.alpha.ImageArchitecture>;
    /**
     * Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
     */
    archiveSizeBytes?: pulumi.Input<string>;
    /**
     * The deprecation status associated with this image.
     */
    deprecated?: pulumi.Input<inputs.compute.alpha.DeprecationStatusArgs>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Size of the image when restored onto a persistent disk (in GB).
     */
    diskSizeGb?: pulumi.Input<string>;
    /**
     * The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
     */
    family?: pulumi.Input<string>;
    /**
     * Force image creation if true.
     */
    forceCreate?: pulumi.Input<string>;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. To see a list of available options, see the guestOSfeatures[].type parameter.
     */
    guestOsFeatures?: pulumi.Input<pulumi.Input<inputs.compute.alpha.GuestOsFeatureArgs>[]>;
    /**
     * Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
     */
    imageEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * Labels to apply to this image. These can be later modified by the setLabels method.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Integer license codes indicating which licenses are attached to this image.
     */
    licenseCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Any applicable license URI.
     */
    licenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A flag for marketplace VM disk created from the image, which is designed for marketplace VM disk to prevent the proprietary data on the disk from being accessed unwantedly. The flag will be inherited by the disk created from the image. The disk with locked flag set to true will be prohibited from performing the operations below: - R/W or R/O disk attach - Disk detach, if disk is created via create-on-create - Create images - Create snapshots - Create disk clone (create disk from the current disk) The image with the locked field set to true will be prohibited from performing the operations below: - Create images from the current image - Update the locked field for the current image The instance with at least one disk with locked flag set to true will be prohibited from performing the operations below: - Secondary disk attach - Create instant snapshot - Create machine images - Create instance template - Delete the instance with --keep-disk parameter set to true 
     */
    locked?: pulumi.Input<boolean>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The parameters of the raw disk image.
     */
    rawDisk?: pulumi.Input<inputs.compute.alpha.ImageRawDiskArgs>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * A rollout policy to apply to this image. When specified, the rollout policy overrides per-zone references to the image via the associated image family. The rollout policy restricts the zones where this image is accessible when using a zonal image family reference. When the rollout policy does not include the user specified zone, or if the zone is rolled out, this image is accessible. The rollout policy for this image is read-only, except for allowlisted users. This field might not be configured. To view the latest non-deprecated image in a specific zone, use the imageFamilyViews.get method.
     */
    rolloutOverride?: pulumi.Input<inputs.compute.alpha.RolloutPolicyArgs>;
    /**
     * Set the secure boot keys of shielded instance.
     */
    shieldedInstanceInitialState?: pulumi.Input<inputs.compute.alpha.InitialStateConfigArgs>;
    /**
     * URL of the source disk used to create this image. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - projects/project/zones/zone/disks/disk - zones/zone/disks/disk In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
     */
    sourceDisk?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
     */
    sourceDiskEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * URL of the source image used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ images/image_name - projects/project_id/global/images/image_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
     */
    sourceImage?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    sourceImageEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * URL of the source snapshot used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ snapshots/snapshot_name - projects/project_id/global/snapshots/snapshot_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
     */
    sourceSnapshot?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    sourceSnapshotEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * The type of the image used to create this disk. The default and only value is RAW
     */
    sourceType?: pulumi.Input<enums.compute.alpha.ImageSourceType>;
    /**
     * Cloud Storage bucket storage location of the image (regional or multi-regional).
     */
    storageLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch 
     */
    userLicenses?: pulumi.Input<pulumi.Input<string>[]>;
}
