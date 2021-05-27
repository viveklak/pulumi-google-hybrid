// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
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
    public static readonly __pulumiType = 'google-native:compute/beta:Disk';

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
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    public readonly creationTimestamp!: pulumi.Output<string>;
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
    public readonly diskEncryptionKey!: pulumi.Output<outputs.compute.beta.CustomerEncryptionKeyResponse>;
    /**
     * Specifies whether the disk restored from a source snapshot should erase Windows specific VSS signature.
     */
    public readonly eraseWindowsVssSignature!: pulumi.Output<boolean>;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
     */
    public readonly guestOsFeatures!: pulumi.Output<outputs.compute.beta.GuestOsFeatureResponse[]>;
    /**
     * Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * [Output Only] Type of the resource. Always compute#disk for disks.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a disk.
     */
    public readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this disk. These can be later modified by the setLabels method.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * [Output Only] Last attach timestamp in RFC3339 text format.
     */
    public readonly lastAttachTimestamp!: pulumi.Output<string>;
    /**
     * [Output Only] Last detach timestamp in RFC3339 text format.
     */
    public readonly lastDetachTimestamp!: pulumi.Output<string>;
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
     * Indicates whether or not the disk can be read/write attached to more than one instance.
     */
    public readonly multiWriter!: pulumi.Output<boolean>;
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
     * [Output Only] URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
     */
    public readonly replicaZones!: pulumi.Output<string[]>;
    /**
     * Resource policies applied to this disk for automatic snapshot creations.
     */
    public readonly resourcePolicies!: pulumi.Output<string[]>;
    /**
     * [Output Only] Reserved for future use.
     */
    public readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * [Output Only] Server-defined fully-qualified URL for this resource.
     */
    public readonly selfLink!: pulumi.Output<string>;
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
     * [Output Only] The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
     */
    public readonly sourceDiskId!: pulumi.Output<string>;
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
    public readonly sourceImageEncryptionKey!: pulumi.Output<outputs.compute.beta.CustomerEncryptionKeyResponse>;
    /**
     * [Output Only] The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
     */
    public readonly sourceImageId!: pulumi.Output<string>;
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
    public readonly sourceSnapshotEncryptionKey!: pulumi.Output<outputs.compute.beta.CustomerEncryptionKeyResponse>;
    /**
     * [Output Only] The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
     */
    public readonly sourceSnapshotId!: pulumi.Output<string>;
    /**
     * The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
     */
    public readonly sourceStorageObject!: pulumi.Output<string>;
    /**
     * [Output Only] The status of disk creation.  
     * - CREATING: Disk is provisioning. 
     * - RESTORING: Source data is being copied into the disk. 
     * - FAILED: Disk creation failed. 
     * - READY: Disk is ready for use. 
     * - DELETING: Disk is deleting.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * [Deprecated] Storage type of the persistent disk.
     */
    public readonly storageType!: pulumi.Output<string>;
    /**
     * URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project/zones/zone/diskTypes/pd-standard  or pd-ssd
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * [Output Only] Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
     */
    public readonly users!: pulumi.Output<string[]>;
    /**
     * [Output Only] URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
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
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["diskEncryptionKey"] = args ? args.diskEncryptionKey : undefined;
            inputs["eraseWindowsVssSignature"] = args ? args.eraseWindowsVssSignature : undefined;
            inputs["guestOsFeatures"] = args ? args.guestOsFeatures : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["labelFingerprint"] = args ? args.labelFingerprint : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["lastAttachTimestamp"] = args ? args.lastAttachTimestamp : undefined;
            inputs["lastDetachTimestamp"] = args ? args.lastDetachTimestamp : undefined;
            inputs["licenseCodes"] = args ? args.licenseCodes : undefined;
            inputs["licenses"] = args ? args.licenses : undefined;
            inputs["locationHint"] = args ? args.locationHint : undefined;
            inputs["multiWriter"] = args ? args.multiWriter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["physicalBlockSizeBytes"] = args ? args.physicalBlockSizeBytes : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["provisionedIops"] = args ? args.provisionedIops : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["replicaZones"] = args ? args.replicaZones : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["resourcePolicies"] = args ? args.resourcePolicies : undefined;
            inputs["satisfiesPzs"] = args ? args.satisfiesPzs : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["sizeGb"] = args ? args.sizeGb : undefined;
            inputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            inputs["sourceDiskId"] = args ? args.sourceDiskId : undefined;
            inputs["sourceImage"] = args ? args.sourceImage : undefined;
            inputs["sourceImageEncryptionKey"] = args ? args.sourceImageEncryptionKey : undefined;
            inputs["sourceImageId"] = args ? args.sourceImageId : undefined;
            inputs["sourceSnapshot"] = args ? args.sourceSnapshot : undefined;
            inputs["sourceSnapshotEncryptionKey"] = args ? args.sourceSnapshotEncryptionKey : undefined;
            inputs["sourceSnapshotId"] = args ? args.sourceSnapshotId : undefined;
            inputs["sourceStorageObject"] = args ? args.sourceStorageObject : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["storageType"] = args ? args.storageType : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["users"] = args ? args.users : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        } else {
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["diskEncryptionKey"] = undefined /*out*/;
            inputs["eraseWindowsVssSignature"] = undefined /*out*/;
            inputs["guestOsFeatures"] = undefined /*out*/;
            inputs["interface"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["lastAttachTimestamp"] = undefined /*out*/;
            inputs["lastDetachTimestamp"] = undefined /*out*/;
            inputs["licenseCodes"] = undefined /*out*/;
            inputs["licenses"] = undefined /*out*/;
            inputs["locationHint"] = undefined /*out*/;
            inputs["multiWriter"] = undefined /*out*/;
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
            inputs["storageType"] = undefined /*out*/;
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
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Encrypts the disk using a customer-supplied encryption key.
     *
     * After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later (e.g. to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine).
     *
     * Customer-supplied encryption keys do not protect access to metadata of the disk.
     *
     * If you do not provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later.
     */
    readonly diskEncryptionKey?: pulumi.Input<inputs.compute.beta.CustomerEncryptionKeyArgs>;
    /**
     * Specifies whether the disk restored from a source snapshot should erase Windows specific VSS signature.
     */
    readonly eraseWindowsVssSignature?: pulumi.Input<boolean>;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. Read  Enabling guest operating system features to see a list of available options.
     */
    readonly guestOsFeatures?: pulumi.Input<pulumi.Input<inputs.compute.beta.GuestOsFeatureArgs>[]>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
     */
    readonly interface?: pulumi.Input<string>;
    /**
     * [Output Only] Type of the resource. Always compute#disk for disks.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a disk.
     */
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels to apply to this disk. These can be later modified by the setLabels method.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * [Output Only] Last attach timestamp in RFC3339 text format.
     */
    readonly lastAttachTimestamp?: pulumi.Input<string>;
    /**
     * [Output Only] Last detach timestamp in RFC3339 text format.
     */
    readonly lastDetachTimestamp?: pulumi.Input<string>;
    /**
     * Integer license codes indicating which licenses are attached to this disk.
     */
    readonly licenseCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of publicly visible licenses. Reserved for Google's use.
     */
    readonly licenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
     */
    readonly locationHint?: pulumi.Input<string>;
    /**
     * Indicates whether or not the disk can be read/write attached to more than one instance.
     */
    readonly multiWriter?: pulumi.Input<boolean>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Internal use only.
     */
    readonly options?: pulumi.Input<string>;
    /**
     * Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
     */
    readonly physicalBlockSizeBytes?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * Indicates how many IOPS must be provisioned for the disk.
     */
    readonly provisionedIops?: pulumi.Input<string>;
    /**
     * [Output Only] URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
     */
    readonly replicaZones?: pulumi.Input<pulumi.Input<string>[]>;
    readonly requestId?: pulumi.Input<string>;
    /**
     * Resource policies applied to this disk for automatic snapshot creations.
     */
    readonly resourcePolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [Output Only] Reserved for future use.
     */
    readonly satisfiesPzs?: pulumi.Input<boolean>;
    /**
     * [Output Only] Server-defined fully-qualified URL for this resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk.
     *
     * If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
     */
    readonly sizeGb?: pulumi.Input<string>;
    /**
     * The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk  
     * - https://www.googleapis.com/compute/v1/projects/project/regions/region/disks/disk  
     * - projects/project/zones/zone/disks/disk  
     * - projects/project/regions/region/disks/disk  
     * - zones/zone/disks/disk  
     * - regions/region/disks/disk
     */
    readonly sourceDisk?: pulumi.Input<string>;
    /**
     * [Output Only] The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
     */
    readonly sourceDiskId?: pulumi.Input<string>;
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
    readonly sourceImage?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    readonly sourceImageEncryptionKey?: pulumi.Input<inputs.compute.beta.CustomerEncryptionKeyArgs>;
    /**
     * [Output Only] The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
     */
    readonly sourceImageId?: pulumi.Input<string>;
    /**
     * The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot 
     * - projects/project/global/snapshots/snapshot 
     * - global/snapshots/snapshot
     */
    readonly sourceSnapshot?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    readonly sourceSnapshotEncryptionKey?: pulumi.Input<inputs.compute.beta.CustomerEncryptionKeyArgs>;
    /**
     * [Output Only] The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
     */
    readonly sourceSnapshotId?: pulumi.Input<string>;
    /**
     * The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
     */
    readonly sourceStorageObject?: pulumi.Input<string>;
    /**
     * [Output Only] The status of disk creation.  
     * - CREATING: Disk is provisioning. 
     * - RESTORING: Source data is being copied into the disk. 
     * - FAILED: Disk creation failed. 
     * - READY: Disk is ready for use. 
     * - DELETING: Disk is deleting.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * [Deprecated] Storage type of the persistent disk.
     */
    readonly storageType?: pulumi.Input<string>;
    /**
     * URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project/zones/zone/diskTypes/pd-standard  or pd-ssd
     */
    readonly type?: pulumi.Input<string>;
    /**
     * [Output Only] Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
     */
    readonly users?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [Output Only] URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly zone: pulumi.Input<string>;
}
