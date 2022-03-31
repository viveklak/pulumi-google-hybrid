// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a persistent regional disk in the specified project using the data included in the request.
 */
export class RegionDisk extends pulumi.CustomResource {
    /**
     * Get an existing RegionDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionDisk {
        return new RegionDisk(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:RegionDisk';

    /**
     * Returns true if the given object is an instance of RegionDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionDisk.__pulumiType;
    }

    /**
     * The architecture of the disk. Valid values are ARM64 or X86_64.
     */
    public readonly architecture!: pulumi.Output<string>;
    /**
     * Disk asynchronously replicated into this disk.
     */
    public readonly asyncPrimaryDisk!: pulumi.Output<outputs.compute.alpha.DiskAsyncReplicationResponse>;
    /**
     * A list of disks this disk is asynchronously replicated to.
     */
    public /*out*/ readonly asyncSecondaryDisks!: pulumi.Output<{[key: string]: string}>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Encrypts the disk using a customer-supplied encryption key or a customer-managed encryption key. Encryption keys do not protect access to metadata of the disk. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later. For example, to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine. After you encrypt a disk with a customer-managed key, the diskEncryptionKey.kmsKeyName is set to a key *version* name once the disk is created. The disk is encrypted with this version of the key. In the response, diskEncryptionKey.kmsKeyName appears in the following format: "diskEncryptionKey.kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeysVersions/version If you do not provide an encryption key when creating the disk, then the disk is encrypted using an automatically generated key and you don't need to provide a key to use the disk later.
     */
    public readonly diskEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * Specifies whether the disk restored from a source snapshot should erase Windows specific VSS signature.
     */
    public readonly eraseWindowsVssSignature!: pulumi.Output<boolean>;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
     */
    public readonly guestOsFeatures!: pulumi.Output<outputs.compute.alpha.GuestOsFeatureResponse[]>;
    /**
     * [Deprecated] Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
     *
     * @deprecated [Deprecated] Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#disk for disks.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a disk.
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
     * The field indicates if the disk is created from a locked source image. Attachment of a disk created from a locked source image will cause the following operations to become irreversibly prohibited: - R/W or R/O disk attachment to any other instance - Disk detachment. And the disk can only be deleted when the instance is deleted - Creation of images or snapshots - Disk cloning Furthermore, the instance with at least one disk with locked flag set to true will be prohibited from performing the operations below: - Further attachment of secondary disks. - Detachment of any disks - Create machine images - Create instance template - Delete the instance with --keep-disk parameter set to true for locked disks - Attach a locked disk with --auto-delete parameter set to false 
     */
    public /*out*/ readonly locked!: pulumi.Output<boolean>;
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
     * Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
     */
    public readonly provisionedIops!: pulumi.Output<string>;
    /**
     * URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
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
     * Status information for the disk resource.
     */
    public /*out*/ readonly resourceStatus!: pulumi.Output<outputs.compute.alpha.DiskResourceStatusResponse>;
    /**
     * Reserved for future use.
     */
    public /*out*/ readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource's resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
     */
    public readonly sizeGb!: pulumi.Output<string>;
    /**
     * URL of the DiskConsistencyGroupPolicy for a secondary disk that was created using a consistency group.
     */
    public /*out*/ readonly sourceConsistencyGroupPolicy!: pulumi.Output<string>;
    /**
     * ID of the DiskConsistencyGroupPolicy for a secondary disk that was created using a consistency group.
     */
    public /*out*/ readonly sourceConsistencyGroupPolicyId!: pulumi.Output<string>;
    /**
     * The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk 
     */
    public readonly sourceDisk!: pulumi.Output<string>;
    /**
     * The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
     */
    public /*out*/ readonly sourceDiskId!: pulumi.Output<string>;
    /**
     * The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family 
     */
    public readonly sourceImage!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    public readonly sourceImageEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
     */
    public /*out*/ readonly sourceImageId!: pulumi.Output<string>;
    /**
     * The source instant snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot 
     */
    public readonly sourceInstantSnapshot!: pulumi.Output<string>;
    /**
     * The unique ID of the instant snapshot used to create this disk. This value identifies the exact instant snapshot that was used to create this persistent disk. For example, if you created the persistent disk from an instant snapshot that was later deleted and recreated under the same name, the source instant snapshot ID would identify the exact version of the instant snapshot that was used.
     */
    public /*out*/ readonly sourceInstantSnapshotId!: pulumi.Output<string>;
    /**
     * The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot 
     */
    public readonly sourceSnapshot!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    public readonly sourceSnapshotEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
     */
    public /*out*/ readonly sourceSnapshotId!: pulumi.Output<string>;
    /**
     * The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
     */
    public readonly sourceStorageObject!: pulumi.Output<string>;
    /**
     * The status of disk creation. - CREATING: Disk is provisioning. - RESTORING: Source data is being copied into the disk. - FAILED: Disk creation failed. - READY: Disk is ready for use. - DELETING: Disk is deleting. 
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * [Deprecated] Storage type of the persistent disk.
     *
     * @deprecated [Deprecated] Storage type of the persistent disk.
     */
    public readonly storageType!: pulumi.Output<string>;
    /**
     * URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch 
     */
    public readonly userLicenses!: pulumi.Output<string[]>;
    /**
     * Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
     */
    public /*out*/ readonly users!: pulumi.Output<string[]>;
    /**
     * URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public /*out*/ readonly zone!: pulumi.Output<string>;

    /**
     * Create a RegionDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionDiskArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["architecture"] = args ? args.architecture : undefined;
            resourceInputs["asyncPrimaryDisk"] = args ? args.asyncPrimaryDisk : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["diskEncryptionKey"] = args ? args.diskEncryptionKey : undefined;
            resourceInputs["eraseWindowsVssSignature"] = args ? args.eraseWindowsVssSignature : undefined;
            resourceInputs["guestOsFeatures"] = args ? args.guestOsFeatures : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["licenseCodes"] = args ? args.licenseCodes : undefined;
            resourceInputs["licenses"] = args ? args.licenses : undefined;
            resourceInputs["locationHint"] = args ? args.locationHint : undefined;
            resourceInputs["multiWriter"] = args ? args.multiWriter : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["physicalBlockSizeBytes"] = args ? args.physicalBlockSizeBytes : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["provisionedIops"] = args ? args.provisionedIops : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["replicaZones"] = args ? args.replicaZones : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["resourcePolicies"] = args ? args.resourcePolicies : undefined;
            resourceInputs["sizeGb"] = args ? args.sizeGb : undefined;
            resourceInputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            resourceInputs["sourceImage"] = args ? args.sourceImage : undefined;
            resourceInputs["sourceImageEncryptionKey"] = args ? args.sourceImageEncryptionKey : undefined;
            resourceInputs["sourceInstantSnapshot"] = args ? args.sourceInstantSnapshot : undefined;
            resourceInputs["sourceSnapshot"] = args ? args.sourceSnapshot : undefined;
            resourceInputs["sourceSnapshotEncryptionKey"] = args ? args.sourceSnapshotEncryptionKey : undefined;
            resourceInputs["sourceStorageObject"] = args ? args.sourceStorageObject : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userLicenses"] = args ? args.userLicenses : undefined;
            resourceInputs["asyncSecondaryDisks"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["lastAttachTimestamp"] = undefined /*out*/;
            resourceInputs["lastDetachTimestamp"] = undefined /*out*/;
            resourceInputs["locked"] = undefined /*out*/;
            resourceInputs["resourceStatus"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["sourceConsistencyGroupPolicy"] = undefined /*out*/;
            resourceInputs["sourceConsistencyGroupPolicyId"] = undefined /*out*/;
            resourceInputs["sourceDiskId"] = undefined /*out*/;
            resourceInputs["sourceImageId"] = undefined /*out*/;
            resourceInputs["sourceInstantSnapshotId"] = undefined /*out*/;
            resourceInputs["sourceSnapshotId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["users"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        } else {
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["asyncPrimaryDisk"] = undefined /*out*/;
            resourceInputs["asyncSecondaryDisks"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["diskEncryptionKey"] = undefined /*out*/;
            resourceInputs["eraseWindowsVssSignature"] = undefined /*out*/;
            resourceInputs["guestOsFeatures"] = undefined /*out*/;
            resourceInputs["interface"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["lastAttachTimestamp"] = undefined /*out*/;
            resourceInputs["lastDetachTimestamp"] = undefined /*out*/;
            resourceInputs["licenseCodes"] = undefined /*out*/;
            resourceInputs["licenses"] = undefined /*out*/;
            resourceInputs["locationHint"] = undefined /*out*/;
            resourceInputs["locked"] = undefined /*out*/;
            resourceInputs["multiWriter"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["options"] = undefined /*out*/;
            resourceInputs["physicalBlockSizeBytes"] = undefined /*out*/;
            resourceInputs["provisionedIops"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["replicaZones"] = undefined /*out*/;
            resourceInputs["resourcePolicies"] = undefined /*out*/;
            resourceInputs["resourceStatus"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["sizeGb"] = undefined /*out*/;
            resourceInputs["sourceConsistencyGroupPolicy"] = undefined /*out*/;
            resourceInputs["sourceConsistencyGroupPolicyId"] = undefined /*out*/;
            resourceInputs["sourceDisk"] = undefined /*out*/;
            resourceInputs["sourceDiskId"] = undefined /*out*/;
            resourceInputs["sourceImage"] = undefined /*out*/;
            resourceInputs["sourceImageEncryptionKey"] = undefined /*out*/;
            resourceInputs["sourceImageId"] = undefined /*out*/;
            resourceInputs["sourceInstantSnapshot"] = undefined /*out*/;
            resourceInputs["sourceInstantSnapshotId"] = undefined /*out*/;
            resourceInputs["sourceSnapshot"] = undefined /*out*/;
            resourceInputs["sourceSnapshotEncryptionKey"] = undefined /*out*/;
            resourceInputs["sourceSnapshotId"] = undefined /*out*/;
            resourceInputs["sourceStorageObject"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["storageType"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["userLicenses"] = undefined /*out*/;
            resourceInputs["users"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegionDisk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionDisk resource.
 */
export interface RegionDiskArgs {
    /**
     * The architecture of the disk. Valid values are ARM64 or X86_64.
     */
    architecture?: pulumi.Input<enums.compute.alpha.RegionDiskArchitecture>;
    /**
     * Disk asynchronously replicated into this disk.
     */
    asyncPrimaryDisk?: pulumi.Input<inputs.compute.alpha.DiskAsyncReplicationArgs>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Encrypts the disk using a customer-supplied encryption key or a customer-managed encryption key. Encryption keys do not protect access to metadata of the disk. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later. For example, to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine. After you encrypt a disk with a customer-managed key, the diskEncryptionKey.kmsKeyName is set to a key *version* name once the disk is created. The disk is encrypted with this version of the key. In the response, diskEncryptionKey.kmsKeyName appears in the following format: "diskEncryptionKey.kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeysVersions/version If you do not provide an encryption key when creating the disk, then the disk is encrypted using an automatically generated key and you don't need to provide a key to use the disk later.
     */
    diskEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * Specifies whether the disk restored from a source snapshot should erase Windows specific VSS signature.
     */
    eraseWindowsVssSignature?: pulumi.Input<boolean>;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
     */
    guestOsFeatures?: pulumi.Input<pulumi.Input<inputs.compute.alpha.GuestOsFeatureArgs>[]>;
    /**
     * [Deprecated] Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
     *
     * @deprecated [Deprecated] Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
     */
    interface?: pulumi.Input<enums.compute.alpha.RegionDiskInterface>;
    /**
     * Labels to apply to this disk. These can be later modified by the setLabels method.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
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
     * Indicates whether or not the disk can be read/write attached to more than one instance.
     */
    multiWriter?: pulumi.Input<boolean>;
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
    project?: pulumi.Input<string>;
    /**
     * Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
     */
    provisionedIops?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    /**
     * URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
     */
    replicaZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * Resource policies applied to this disk for automatic snapshot creations.
     */
    resourcePolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
     */
    sizeGb?: pulumi.Input<string>;
    /**
     * The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk 
     */
    sourceDisk?: pulumi.Input<string>;
    /**
     * The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family 
     */
    sourceImage?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    sourceImageEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * The source instant snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot 
     */
    sourceInstantSnapshot?: pulumi.Input<string>;
    /**
     * The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot 
     */
    sourceSnapshot?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    sourceSnapshotEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
     */
    sourceStorageObject?: pulumi.Input<string>;
    /**
     * [Deprecated] Storage type of the persistent disk.
     *
     * @deprecated [Deprecated] Storage type of the persistent disk.
     */
    storageType?: pulumi.Input<enums.compute.alpha.RegionDiskStorageType>;
    /**
     * URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
     */
    type?: pulumi.Input<string>;
    /**
     * A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch 
     */
    userLicenses?: pulumi.Input<pulumi.Input<string>[]>;
}
