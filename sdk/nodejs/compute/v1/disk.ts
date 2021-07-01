// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a persistent disk in the specified project using the data in the request. You can create a disk from a source (sourceImage, sourceSnapshot, or sourceDisk) or create an empty 500 GB data disk by omitting all properties. You can also create a disk that is larger than the default size by specifying the sizeGb property.
 */
export class Disk extends pulumi.CustomResource {
    /**
     * Get an existing Disk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Disk {
        return new Disk(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/v1:Disk';

    /**
     * Returns true if the given object is an instance of Disk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Disk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Disk.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Encrypts the disk using a customer-supplied encryption key.
     *
     * After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later (e.g. to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine).
     *
     * Customer-supplied encryption keys do not protect access to metadata of the disk.
     *
     * If you do not provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later.
     */
    public readonly diskEncryptionKey!: pulumi.Output<outputs.compute.v1.CustomerEncryptionKeyResponse>;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
     */
    public readonly guestOsFeatures!: pulumi.Output<outputs.compute.v1.GuestOsFeatureResponse[]>;
    /**
     * Type of the resource. Always compute#disk for disks.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a disk.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this disk. These can be later modified by the setLabels method.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Last attach timestamp in RFC3339 text format.
     */
    public /*out*/ readonly lastAttachTimestamp!: pulumi.Output<string>;
    /**
     * Last detach timestamp in RFC3339 text format.
     */
    public /*out*/ readonly lastDetachTimestamp!: pulumi.Output<string>;
    /**
     * Integer license codes indicating which licenses are attached to this disk.
     */
    public readonly licenseCodes!: pulumi.Output<string[]>;
    /**
     * A list of publicly visible licenses. Reserved for Google's use.
     */
    public readonly licenses!: pulumi.Output<string[]>;
    /**
     * An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
     */
    public readonly locationHint!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Internal use only.
     */
    public readonly options!: pulumi.Output<string>;
    /**
     * Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
     */
    public readonly physicalBlockSizeBytes!: pulumi.Output<string>;
    /**
     * Indicates how many IOPS must be provisioned for the disk.
     */
    public readonly provisionedIops!: pulumi.Output<string>;
    /**
     * URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
     */
    public readonly replicaZones!: pulumi.Output<string[]>;
    /**
     * Resource policies applied to this disk for automatic snapshot creations.
     */
    public readonly resourcePolicies!: pulumi.Output<string[]>;
    /**
     * Reserved for future use.
     */
    public /*out*/ readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk.
     *
     * If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
     */
    public readonly sizeGb!: pulumi.Output<string>;
    /**
     * The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk  
     * - https://www.googleapis.com/compute/v1/projects/project/regions/region/disks/disk  
     * - projects/project/zones/zone/disks/disk  
     * - projects/project/regions/region/disks/disk  
     * - zones/zone/disks/disk  
     * - regions/region/disks/disk
     */
    public readonly sourceDisk!: pulumi.Output<string>;
    /**
     * The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
     */
    public /*out*/ readonly sourceDiskId!: pulumi.Output<string>;
    /**
     * The source image used to create this disk. If the source image is deleted, this field will not be set.
     *
     * To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image:
     * projects/debian-cloud/global/images/family/debian-9
     *
     *
     * Alternatively, use a specific version of a public operating system image:
     * projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD
     *
     *
     * To create a disk with a custom image that you created, specify the image name in the following format:
     * global/images/my-custom-image
     *
     *
     * You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name:
     * global/images/family/my-image-family
     */
    public readonly sourceImage!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    public readonly sourceImageEncryptionKey!: pulumi.Output<outputs.compute.v1.CustomerEncryptionKeyResponse>;
    /**
     * The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
     */
    public /*out*/ readonly sourceImageId!: pulumi.Output<string>;
    /**
     * The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot 
     * - projects/project/global/snapshots/snapshot 
     * - global/snapshots/snapshot
     */
    public readonly sourceSnapshot!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    public readonly sourceSnapshotEncryptionKey!: pulumi.Output<outputs.compute.v1.CustomerEncryptionKeyResponse>;
    /**
     * The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
     */
    public /*out*/ readonly sourceSnapshotId!: pulumi.Output<string>;
    /**
     * The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
     */
    public readonly sourceStorageObject!: pulumi.Output<string>;
    /**
     * The status of disk creation.  
     * - CREATING: Disk is provisioning. 
     * - RESTORING: Source data is being copied into the disk. 
     * - FAILED: Disk creation failed. 
     * - READY: Disk is ready for use. 
     * - DELETING: Disk is deleting.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project/zones/zone/diskTypes/pd-standard  or pd-ssd
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
     */
    public /*out*/ readonly users!: pulumi.Output<string[]>;
    /**
     * URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Disk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["diskEncryptionKey"] = args ? args.diskEncryptionKey : undefined;
            inputs["guestOsFeatures"] = args ? args.guestOsFeatures : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["licenseCodes"] = args ? args.licenseCodes : undefined;
            inputs["licenses"] = args ? args.licenses : undefined;
            inputs["locationHint"] = args ? args.locationHint : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["physicalBlockSizeBytes"] = args ? args.physicalBlockSizeBytes : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["provisionedIops"] = args ? args.provisionedIops : undefined;
            inputs["replicaZones"] = args ? args.replicaZones : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["resourcePolicies"] = args ? args.resourcePolicies : undefined;
            inputs["sizeGb"] = args ? args.sizeGb : undefined;
            inputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            inputs["sourceImage"] = args ? args.sourceImage : undefined;
            inputs["sourceImageEncryptionKey"] = args ? args.sourceImageEncryptionKey : undefined;
            inputs["sourceSnapshot"] = args ? args.sourceSnapshot : undefined;
            inputs["sourceSnapshotEncryptionKey"] = args ? args.sourceSnapshotEncryptionKey : undefined;
            inputs["sourceStorageObject"] = args ? args.sourceStorageObject : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["lastAttachTimestamp"] = undefined /*out*/;
            inputs["lastDetachTimestamp"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["satisfiesPzs"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["sourceDiskId"] = undefined /*out*/;
            inputs["sourceImageId"] = undefined /*out*/;
            inputs["sourceSnapshotId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["users"] = undefined /*out*/;
        } else {
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["diskEncryptionKey"] = undefined /*out*/;
            inputs["guestOsFeatures"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["lastAttachTimestamp"] = undefined /*out*/;
            inputs["lastDetachTimestamp"] = undefined /*out*/;
            inputs["licenseCodes"] = undefined /*out*/;
            inputs["licenses"] = undefined /*out*/;
            inputs["locationHint"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["options"] = undefined /*out*/;
            inputs["physicalBlockSizeBytes"] = undefined /*out*/;
            inputs["provisionedIops"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["replicaZones"] = undefined /*out*/;
            inputs["resourcePolicies"] = undefined /*out*/;
            inputs["satisfiesPzs"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["sizeGb"] = undefined /*out*/;
            inputs["sourceDisk"] = undefined /*out*/;
            inputs["sourceDiskId"] = undefined /*out*/;
            inputs["sourceImage"] = undefined /*out*/;
            inputs["sourceImageEncryptionKey"] = undefined /*out*/;
            inputs["sourceImageId"] = undefined /*out*/;
            inputs["sourceSnapshot"] = undefined /*out*/;
            inputs["sourceSnapshotEncryptionKey"] = undefined /*out*/;
            inputs["sourceSnapshotId"] = undefined /*out*/;
            inputs["sourceStorageObject"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["users"] = undefined /*out*/;
            inputs["zone"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Disk.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Disk resource.
 */
export interface DiskArgs {
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Encrypts the disk using a customer-supplied encryption key.
     *
     * After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later (e.g. to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine).
     *
     * Customer-supplied encryption keys do not protect access to metadata of the disk.
     *
     * If you do not provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later.
     */
    diskEncryptionKey?: pulumi.Input<inputs.compute.v1.CustomerEncryptionKeyArgs>;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
     */
    guestOsFeatures?: pulumi.Input<pulumi.Input<inputs.compute.v1.GuestOsFeatureArgs>[]>;
    /**
     * Labels to apply to this disk. These can be later modified by the setLabels method.
     */
    labels?: pulumi.Input<{[key: string]: string}>;
    /**
     * Integer license codes indicating which licenses are attached to this disk.
     */
    licenseCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of publicly visible licenses. Reserved for Google's use.
     */
    licenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
     */
    locationHint?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * Internal use only.
     */
    options?: pulumi.Input<string>;
    /**
     * Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
     */
    physicalBlockSizeBytes?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    /**
     * Indicates how many IOPS must be provisioned for the disk.
     */
    provisionedIops?: pulumi.Input<string>;
    /**
     * URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
     */
    replicaZones?: pulumi.Input<pulumi.Input<string>[]>;
    requestId?: pulumi.Input<string>;
    /**
     * Resource policies applied to this disk for automatic snapshot creations.
     */
    resourcePolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk.
     *
     * If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
     */
    sizeGb?: pulumi.Input<string>;
    /**
     * The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk  
     * - https://www.googleapis.com/compute/v1/projects/project/regions/region/disks/disk  
     * - projects/project/zones/zone/disks/disk  
     * - projects/project/regions/region/disks/disk  
     * - zones/zone/disks/disk  
     * - regions/region/disks/disk
     */
    sourceDisk?: pulumi.Input<string>;
    /**
     * The source image used to create this disk. If the source image is deleted, this field will not be set.
     *
     * To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image:
     * projects/debian-cloud/global/images/family/debian-9
     *
     *
     * Alternatively, use a specific version of a public operating system image:
     * projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD
     *
     *
     * To create a disk with a custom image that you created, specify the image name in the following format:
     * global/images/my-custom-image
     *
     *
     * You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name:
     * global/images/family/my-image-family
     */
    sourceImage?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    sourceImageEncryptionKey?: pulumi.Input<inputs.compute.v1.CustomerEncryptionKeyArgs>;
    /**
     * The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot 
     * - projects/project/global/snapshots/snapshot 
     * - global/snapshots/snapshot
     */
    sourceSnapshot?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    sourceSnapshotEncryptionKey?: pulumi.Input<inputs.compute.v1.CustomerEncryptionKeyArgs>;
    /**
     * The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
     */
    sourceStorageObject?: pulumi.Input<string>;
    /**
     * URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project/zones/zone/diskTypes/pd-standard  or pd-ssd
     */
    type?: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
