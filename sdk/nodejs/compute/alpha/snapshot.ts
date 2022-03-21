// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a snapshot in the specified project using the data included in the request. For regular snapshot creation, consider using this method instead of disks.createSnapshot, as this method supports more features, such as creating snapshots in a project different from the source disk project.
 */
export class Snapshot extends pulumi.CustomResource {
    /**
     * Get an existing Snapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Snapshot {
        return new Snapshot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:Snapshot';

    /**
     * Returns true if the given object is an instance of Snapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snapshot.__pulumiType;
    }

    /**
     * The architecture of the snapshot. Valid values are ARM64 or X86_64.
     */
    public /*out*/ readonly architecture!: pulumi.Output<string>;
    /**
     * Set to true if snapshots are automatically created by applying resource policy on the target disk.
     */
    public /*out*/ readonly autoCreated!: pulumi.Output<boolean>;
    /**
     * Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
     */
    public readonly chainName!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Size of the source disk, specified in GB.
     */
    public /*out*/ readonly diskSizeGb!: pulumi.Output<string>;
    /**
     * Number of bytes downloaded to restore a snapshot to a disk.
     */
    public /*out*/ readonly downloadBytes!: pulumi.Output<string>;
    /**
     * [Input Only] Whether to attempt an application consistent snapshot by informing the OS to prepare for the snapshot process.
     */
    public readonly guestFlush!: pulumi.Output<boolean>;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
     */
    public /*out*/ readonly guestOsFeatures!: pulumi.Output<outputs.compute.alpha.GuestOsFeatureResponse[]>;
    /**
     * Type of the resource. Always compute#snapshot for Snapshot resources.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for the labels being applied to this snapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a snapshot.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Integer license codes indicating which licenses are attached to this snapshot.
     */
    public /*out*/ readonly licenseCodes!: pulumi.Output<string[]>;
    /**
     * A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses attached (such as a Windows image).
     */
    public /*out*/ readonly licenses!: pulumi.Output<string[]>;
    /**
     * An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
     */
    public readonly locationHint!: pulumi.Output<string>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
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
     * Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
     */
    public readonly snapshotEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * Indicates the type of the snapshot.
     */
    public readonly snapshotType!: pulumi.Output<string>;
    /**
     * The source disk used to create this snapshot.
     */
    public readonly sourceDisk!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
     */
    public readonly sourceDiskEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * The ID value of the disk used to create this snapshot. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given disk name.
     */
    public /*out*/ readonly sourceDiskId!: pulumi.Output<string>;
    /**
     * The source instant snapshot used to create this snapshot. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot 
     */
    public readonly sourceInstantSnapshot!: pulumi.Output<string>;
    /**
     * The unique ID of the instant snapshot used to create this snapshot. This value identifies the exact instant snapshot that was used to create this persistent disk. For example, if you created the persistent disk from an instant snapshot that was later deleted and recreated under the same name, the source instant snapshot ID would identify the exact instant snapshot that was used.
     */
    public /*out*/ readonly sourceInstantSnapshotId!: pulumi.Output<string>;
    /**
     * URL of the resource policy which created this scheduled snapshot.
     */
    public /*out*/ readonly sourceSnapshotSchedulePolicy!: pulumi.Output<string>;
    /**
     * ID of the resource policy which created this scheduled snapshot.
     */
    public /*out*/ readonly sourceSnapshotSchedulePolicyId!: pulumi.Output<string>;
    /**
     * The status of the snapshot. This can be CREATING, DELETING, FAILED, READY, or UPLOADING.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.
     */
    public /*out*/ readonly storageBytes!: pulumi.Output<string>;
    /**
     * An indicator whether storageBytes is in a stable state or it is being adjusted as a result of shared storage reallocation. This status can either be UPDATING, meaning the size of the snapshot is being updated, or UP_TO_DATE, meaning the size of the snapshot is up-to-date.
     */
    public /*out*/ readonly storageBytesStatus!: pulumi.Output<string>;
    /**
     * Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
     */
    public readonly storageLocations!: pulumi.Output<string[]>;
    /**
     * A list of user provided licenses represented by a list of URLs to the license resource.
     */
    public /*out*/ readonly userLicenses!: pulumi.Output<string[]>;

    /**
     * Create a Snapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SnapshotArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["chainName"] = args ? args.chainName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["guestFlush"] = args ? args.guestFlush : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["locationHint"] = args ? args.locationHint : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["snapshotEncryptionKey"] = args ? args.snapshotEncryptionKey : undefined;
            resourceInputs["snapshotType"] = args ? args.snapshotType : undefined;
            resourceInputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            resourceInputs["sourceDiskEncryptionKey"] = args ? args.sourceDiskEncryptionKey : undefined;
            resourceInputs["sourceInstantSnapshot"] = args ? args.sourceInstantSnapshot : undefined;
            resourceInputs["storageLocations"] = args ? args.storageLocations : undefined;
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["autoCreated"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["diskSizeGb"] = undefined /*out*/;
            resourceInputs["downloadBytes"] = undefined /*out*/;
            resourceInputs["guestOsFeatures"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["licenseCodes"] = undefined /*out*/;
            resourceInputs["licenses"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["sourceDiskId"] = undefined /*out*/;
            resourceInputs["sourceInstantSnapshotId"] = undefined /*out*/;
            resourceInputs["sourceSnapshotSchedulePolicy"] = undefined /*out*/;
            resourceInputs["sourceSnapshotSchedulePolicyId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["storageBytes"] = undefined /*out*/;
            resourceInputs["storageBytesStatus"] = undefined /*out*/;
            resourceInputs["userLicenses"] = undefined /*out*/;
        } else {
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["autoCreated"] = undefined /*out*/;
            resourceInputs["chainName"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["diskSizeGb"] = undefined /*out*/;
            resourceInputs["downloadBytes"] = undefined /*out*/;
            resourceInputs["guestFlush"] = undefined /*out*/;
            resourceInputs["guestOsFeatures"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["licenseCodes"] = undefined /*out*/;
            resourceInputs["licenses"] = undefined /*out*/;
            resourceInputs["locationHint"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["snapshotEncryptionKey"] = undefined /*out*/;
            resourceInputs["snapshotType"] = undefined /*out*/;
            resourceInputs["sourceDisk"] = undefined /*out*/;
            resourceInputs["sourceDiskEncryptionKey"] = undefined /*out*/;
            resourceInputs["sourceDiskId"] = undefined /*out*/;
            resourceInputs["sourceInstantSnapshot"] = undefined /*out*/;
            resourceInputs["sourceInstantSnapshotId"] = undefined /*out*/;
            resourceInputs["sourceSnapshotSchedulePolicy"] = undefined /*out*/;
            resourceInputs["sourceSnapshotSchedulePolicyId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["storageBytes"] = undefined /*out*/;
            resourceInputs["storageBytesStatus"] = undefined /*out*/;
            resourceInputs["storageLocations"] = undefined /*out*/;
            resourceInputs["userLicenses"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Snapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Snapshot resource.
 */
export interface SnapshotArgs {
    /**
     * Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
     */
    chainName?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * [Input Only] Whether to attempt an application consistent snapshot by informing the OS to prepare for the snapshot process.
     */
    guestFlush?: pulumi.Input<boolean>;
    /**
     * Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
     */
    locationHint?: pulumi.Input<string>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
     */
    snapshotEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * Indicates the type of the snapshot.
     */
    snapshotType?: pulumi.Input<enums.compute.alpha.SnapshotSnapshotType>;
    /**
     * The source disk used to create this snapshot.
     */
    sourceDisk?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
     */
    sourceDiskEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * The source instant snapshot used to create this snapshot. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot 
     */
    sourceInstantSnapshot?: pulumi.Input<string>;
    /**
     * Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
     */
    storageLocations?: pulumi.Input<pulumi.Input<string>[]>;
}
