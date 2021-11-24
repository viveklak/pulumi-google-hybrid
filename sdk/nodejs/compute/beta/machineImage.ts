// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a machine image in the specified project using the data that is included in the request. If you are creating a new machine image to update an existing instance, your new machine image should use the same network or, if applicable, the same subnetwork as the original instance.
 */
export class MachineImage extends pulumi.CustomResource {
    /**
     * Get an existing MachineImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MachineImage {
        return new MachineImage(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:MachineImage';

    /**
     * Returns true if the given object is an instance of MachineImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MachineImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MachineImage.__pulumiType;
    }

    /**
     * The creation timestamp for this machine image in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
     */
    public readonly guestFlush!: pulumi.Output<boolean>;
    /**
     * The resource type, which is always compute#machineImage for machine image.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
     */
    public readonly machineImageEncryptionKey!: pulumi.Output<outputs.compute.beta.CustomerEncryptionKeyResponse>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Reserved for future use.
     */
    public /*out*/ readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * The URL for this machine image. The server defines this URL.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
     */
    public readonly sourceDiskEncryptionKeys!: pulumi.Output<outputs.compute.beta.SourceDiskEncryptionKeyResponse[]>;
    /**
     * The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
     */
    public readonly sourceInstance!: pulumi.Output<string>;
    /**
     * Properties of source instance.
     */
    public /*out*/ readonly sourceInstanceProperties!: pulumi.Output<outputs.compute.beta.SourceInstancePropertiesResponse>;
    /**
     * The status of the machine image. One of the following values: INVALID, CREATING, READY, DELETING, and UPLOADING.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
     */
    public readonly storageLocations!: pulumi.Output<string[]>;
    /**
     * Total size of the storage used by the machine image.
     */
    public /*out*/ readonly totalStorageBytes!: pulumi.Output<string>;

    /**
     * Create a MachineImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MachineImageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.sourceInstance === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceInstance'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["guestFlush"] = args ? args.guestFlush : undefined;
            resourceInputs["machineImageEncryptionKey"] = args ? args.machineImageEncryptionKey : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["sourceDiskEncryptionKeys"] = args ? args.sourceDiskEncryptionKeys : undefined;
            resourceInputs["sourceInstance"] = args ? args.sourceInstance : undefined;
            resourceInputs["storageLocations"] = args ? args.storageLocations : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["sourceInstanceProperties"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["totalStorageBytes"] = undefined /*out*/;
        } else {
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["guestFlush"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["machineImageEncryptionKey"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["sourceDiskEncryptionKeys"] = undefined /*out*/;
            resourceInputs["sourceInstance"] = undefined /*out*/;
            resourceInputs["sourceInstanceProperties"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["storageLocations"] = undefined /*out*/;
            resourceInputs["totalStorageBytes"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MachineImage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MachineImage resource.
 */
export interface MachineImageArgs {
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
     */
    guestFlush?: pulumi.Input<boolean>;
    /**
     * Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
     */
    machineImageEncryptionKey?: pulumi.Input<inputs.compute.beta.CustomerEncryptionKeyArgs>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
     */
    sourceDiskEncryptionKeys?: pulumi.Input<pulumi.Input<inputs.compute.beta.SourceDiskEncryptionKeyArgs>[]>;
    /**
     * The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
     */
    sourceInstance: pulumi.Input<string>;
    /**
     * The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
     */
    storageLocations?: pulumi.Input<pulumi.Input<string>[]>;
}
