// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified machine image. Gets a list of available machine images by making a list() request.
 */
export function getMachineImage(args: GetMachineImageArgs, opts?: pulumi.InvokeOptions): Promise<GetMachineImageResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/alpha:getMachineImage", {
        "machineImage": args.machineImage,
        "project": args.project,
    }, opts);
}

export interface GetMachineImageArgs {
    machineImage: string;
    project?: string;
}

export interface GetMachineImageResult {
    /**
     * The creation timestamp for this machine image in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
     */
    readonly guestFlush: boolean;
    /**
     * Properties of source instance
     */
    readonly instanceProperties: outputs.compute.alpha.InstancePropertiesResponse;
    /**
     * The resource type, which is always compute#machineImage for machine image.
     */
    readonly kind: string;
    /**
     * Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
     */
    readonly machineImageEncryptionKey: outputs.compute.alpha.CustomerEncryptionKeyResponse;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Reserved for future use.
     */
    readonly satisfiesPzs: boolean;
    /**
     * An array of Machine Image specific properties for disks attached to the source instance
     */
    readonly savedDisks: outputs.compute.alpha.SavedDiskResponse[];
    /**
     * The URL for this machine image. The server defines this URL.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
     */
    readonly sourceDiskEncryptionKeys: outputs.compute.alpha.SourceDiskEncryptionKeyResponse[];
    /**
     * The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
     */
    readonly sourceInstance: string;
    /**
     * DEPRECATED: Please use instance_properties instead for source instance related properties. New properties will not be added to this field.
     */
    readonly sourceInstanceProperties: outputs.compute.alpha.SourceInstancePropertiesResponse;
    /**
     * The status of the machine image. One of the following values: INVALID, CREATING, READY, DELETING, and UPLOADING.
     */
    readonly status: string;
    /**
     * The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
     */
    readonly storageLocations: string[];
    /**
     * Total size of the storage used by the machine image.
     */
    readonly totalStorageBytes: string;
}

export function getMachineImageOutput(args: GetMachineImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMachineImageResult> {
    return pulumi.output(args).apply(a => getMachineImage(a, opts))
}

export interface GetMachineImageOutputArgs {
    machineImage: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
