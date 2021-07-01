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
    public static readonly __pulumiType = 'google-native:compute/beta:Image';

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
    public readonly deprecated!: pulumi.Output<outputs.compute.beta.DeprecationStatusResponse>;
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
     * A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
     */
    public readonly guestOsFeatures!: pulumi.Output<outputs.compute.beta.GuestOsFeatureResponse[]>;
    /**
     * Encrypts the image using a customer-supplied encryption key.
     *
     * After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image).
     *
     * Customer-supplied encryption keys do not protect access to metadata of the disk.
     *
     * If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
     */
    public readonly imageEncryptionKey!: pulumi.Output<outputs.compute.beta.CustomerEncryptionKeyResponse>;
    /**
     * Type of the resource. Always compute#image for images.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve an image.
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
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The parameters of the raw disk image.
     */
    public readonly rawDisk!: pulumi.Output<outputs.compute.beta.ImageRawDiskResponse>;
    /**
     * Reserved for future use.
     */
    public /*out*/ readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Set the secure boot keys of shielded instance.
     */
    public readonly shieldedInstanceInitialState!: pulumi.Output<outputs.compute.beta.InitialStateConfigResponse>;
    /**
     * URL of the source disk used to create this image. This can be a full or valid partial URL. You must provide either this property or the rawDisk.source property but not both to create an image. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk 
     * - projects/project/zones/zone/disks/disk 
     * - zones/zone/disks/disk
     */
    public readonly sourceDisk!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
     */
    public readonly sourceDiskEncryptionKey!: pulumi.Output<outputs.compute.beta.CustomerEncryptionKeyResponse>;
    /**
     * The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
     */
    public /*out*/ readonly sourceDiskId!: pulumi.Output<string>;
    /**
     * URL of the source image used to create this image.
     *
     * In order to create an image, you must provide the full or partial URL of one of the following:  
     * - The selfLink URL  
     * - This property  
     * - The rawDisk.source URL  
     * - The sourceDisk URL
     */
    public readonly sourceImage!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    public readonly sourceImageEncryptionKey!: pulumi.Output<outputs.compute.beta.CustomerEncryptionKeyResponse>;
    /**
     * The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
     */
    public /*out*/ readonly sourceImageId!: pulumi.Output<string>;
    /**
     * URL of the source snapshot used to create this image.
     *
     * In order to create an image, you must provide the full or partial URL of one of the following:  
     * - The selfLink URL  
     * - This property 
     * - The sourceImage URL  
     * - The rawDisk.source URL  
     * - The sourceDisk URL
     */
    public readonly sourceSnapshot!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    public readonly sourceSnapshotEncryptionKey!: pulumi.Output<outputs.compute.beta.CustomerEncryptionKeyResponse>;
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
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["archiveSizeBytes"] = args ? args.archiveSizeBytes : undefined;
            inputs["deprecated"] = args ? args.deprecated : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["diskSizeGb"] = args ? args.diskSizeGb : undefined;
            inputs["family"] = args ? args.family : undefined;
            inputs["forceCreate"] = args ? args.forceCreate : undefined;
            inputs["guestOsFeatures"] = args ? args.guestOsFeatures : undefined;
            inputs["imageEncryptionKey"] = args ? args.imageEncryptionKey : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["licenseCodes"] = args ? args.licenseCodes : undefined;
            inputs["licenses"] = args ? args.licenses : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["rawDisk"] = args ? args.rawDisk : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["shieldedInstanceInitialState"] = args ? args.shieldedInstanceInitialState : undefined;
            inputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            inputs["sourceDiskEncryptionKey"] = args ? args.sourceDiskEncryptionKey : undefined;
            inputs["sourceImage"] = args ? args.sourceImage : undefined;
            inputs["sourceImageEncryptionKey"] = args ? args.sourceImageEncryptionKey : undefined;
            inputs["sourceSnapshot"] = args ? args.sourceSnapshot : undefined;
            inputs["sourceSnapshotEncryptionKey"] = args ? args.sourceSnapshotEncryptionKey : undefined;
            inputs["sourceType"] = args ? args.sourceType : undefined;
            inputs["storageLocations"] = args ? args.storageLocations : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["satisfiesPzs"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["sourceDiskId"] = undefined /*out*/;
            inputs["sourceImageId"] = undefined /*out*/;
            inputs["sourceSnapshotId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        } else {
            inputs["archiveSizeBytes"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["deprecated"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["diskSizeGb"] = undefined /*out*/;
            inputs["family"] = undefined /*out*/;
            inputs["guestOsFeatures"] = undefined /*out*/;
            inputs["imageEncryptionKey"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["licenseCodes"] = undefined /*out*/;
            inputs["licenses"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["rawDisk"] = undefined /*out*/;
            inputs["satisfiesPzs"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["shieldedInstanceInitialState"] = undefined /*out*/;
            inputs["sourceDisk"] = undefined /*out*/;
            inputs["sourceDiskEncryptionKey"] = undefined /*out*/;
            inputs["sourceDiskId"] = undefined /*out*/;
            inputs["sourceImage"] = undefined /*out*/;
            inputs["sourceImageEncryptionKey"] = undefined /*out*/;
            inputs["sourceImageId"] = undefined /*out*/;
            inputs["sourceSnapshot"] = undefined /*out*/;
            inputs["sourceSnapshotEncryptionKey"] = undefined /*out*/;
            inputs["sourceSnapshotId"] = undefined /*out*/;
            inputs["sourceType"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["storageLocations"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Image.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
     */
    archiveSizeBytes?: pulumi.Input<string>;
    /**
     * The deprecation status associated with this image.
     */
    deprecated?: pulumi.Input<inputs.compute.beta.DeprecationStatusArgs>;
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
    forceCreate?: pulumi.Input<string>;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
     */
    guestOsFeatures?: pulumi.Input<pulumi.Input<inputs.compute.beta.GuestOsFeatureArgs>[]>;
    /**
     * Encrypts the image using a customer-supplied encryption key.
     *
     * After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image).
     *
     * Customer-supplied encryption keys do not protect access to metadata of the disk.
     *
     * If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
     */
    imageEncryptionKey?: pulumi.Input<inputs.compute.beta.CustomerEncryptionKeyArgs>;
    /**
     * Labels to apply to this image. These can be later modified by the setLabels method.
     */
    labels?: pulumi.Input<{[key: string]: string}>;
    /**
     * Integer license codes indicating which licenses are attached to this image.
     */
    licenseCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Any applicable license URI.
     */
    licenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    /**
     * The parameters of the raw disk image.
     */
    rawDisk?: pulumi.Input<inputs.compute.beta.ImageRawDiskArgs>;
    requestId?: pulumi.Input<string>;
    /**
     * Set the secure boot keys of shielded instance.
     */
    shieldedInstanceInitialState?: pulumi.Input<inputs.compute.beta.InitialStateConfigArgs>;
    /**
     * URL of the source disk used to create this image. This can be a full or valid partial URL. You must provide either this property or the rawDisk.source property but not both to create an image. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk 
     * - projects/project/zones/zone/disks/disk 
     * - zones/zone/disks/disk
     */
    sourceDisk?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
     */
    sourceDiskEncryptionKey?: pulumi.Input<inputs.compute.beta.CustomerEncryptionKeyArgs>;
    /**
     * URL of the source image used to create this image.
     *
     * In order to create an image, you must provide the full or partial URL of one of the following:  
     * - The selfLink URL  
     * - This property  
     * - The rawDisk.source URL  
     * - The sourceDisk URL
     */
    sourceImage?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    sourceImageEncryptionKey?: pulumi.Input<inputs.compute.beta.CustomerEncryptionKeyArgs>;
    /**
     * URL of the source snapshot used to create this image.
     *
     * In order to create an image, you must provide the full or partial URL of one of the following:  
     * - The selfLink URL  
     * - This property 
     * - The sourceImage URL  
     * - The rawDisk.source URL  
     * - The sourceDisk URL
     */
    sourceSnapshot?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    sourceSnapshotEncryptionKey?: pulumi.Input<inputs.compute.beta.CustomerEncryptionKeyArgs>;
    /**
     * The type of the image used to create this disk. The default and only value is RAW
     */
    sourceType?: pulumi.Input<enums.compute.beta.ImageSourceType>;
    /**
     * Cloud Storage bucket storage location of the image (regional or multi-regional).
     */
    storageLocations?: pulumi.Input<pulumi.Input<string>[]>;
}
