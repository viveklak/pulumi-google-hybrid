// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Instance in a given project and location.
 * Auto-naming is currently not supported for this resource.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:notebooks/v1:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The hardware accelerator used on this instance. If you use accelerators, make sure that your configuration has [enough vCPUs and memory to support the `machine_type` you have selected](/compute/docs/gpus/#gpus-list).
     */
    public readonly acceleratorConfig!: pulumi.Output<outputs.notebooks.v1.AcceleratorConfigResponse>;
    /**
     * Input only. The size of the boot disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB. If not specified, this defaults to 100.
     */
    public readonly bootDiskSizeGb!: pulumi.Output<string>;
    /**
     * Input only. The type of the boot disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
     */
    public readonly bootDiskType!: pulumi.Output<string>;
    /**
     * Use a container image to start the notebook instance.
     */
    public readonly containerImage!: pulumi.Output<outputs.notebooks.v1.ContainerImageResponse>;
    /**
     * Instance creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Specify a custom Cloud Storage path where the GPU driver is stored. If not specified, we'll automatically choose from official GPU drivers.
     */
    public readonly customGpuDriverPath!: pulumi.Output<string>;
    /**
     * Input only. The size of the data disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). You can choose the size of the data disk based on how big your notebooks and data are. If not specified, this defaults to 100.
     */
    public readonly dataDiskSizeGb!: pulumi.Output<string>;
    /**
     * Input only. The type of the data disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
     */
    public readonly dataDiskType!: pulumi.Output<string>;
    /**
     * Input only. Disk encryption method used on the boot and data disks, defaults to GMEK.
     */
    public readonly diskEncryption!: pulumi.Output<string>;
    /**
     * Attached disks to notebook instance.
     */
    public /*out*/ readonly disks!: pulumi.Output<outputs.notebooks.v1.DiskResponse[]>;
    /**
     * Whether the end user authorizes Google Cloud to install GPU driver on this instance. If this field is empty or set to false, the GPU driver won't be installed. Only applicable to instances with GPUs.
     */
    public readonly installGpuDriver!: pulumi.Output<boolean>;
    /**
     * Input only. The owner of this instance after creation. Format: `alias@example.com` Currently supports one owner only. If not specified, all of the service account users of your VM instance's service account can use the instance.
     */
    public readonly instanceOwners!: pulumi.Output<string[]>;
    /**
     * Input only. The KMS key used to encrypt the disks, only applicable if disk_encryption is CMEK. Format: `projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}` Learn more about [using your own encryption keys](/kms/docs/quickstart).
     */
    public readonly kmsKey!: pulumi.Output<string>;
    /**
     * Labels to apply to this instance. These can be later modified by the setLabels method.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The [Compute Engine machine type](/compute/docs/machine-types) of this instance.
     */
    public readonly machineType!: pulumi.Output<string>;
    /**
     * Custom metadata to apply to this instance.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * The name of this notebook instance. Format: `projects/{project_id}/locations/{location}/instances/{instance_id}`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The name of the VPC that this instance is in. Format: `projects/{project_id}/global/networks/{network_id}`
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Optional. The type of vNIC to be used on this interface. This may be gVNIC or VirtioNet.
     */
    public readonly nicType!: pulumi.Output<string>;
    /**
     * If true, the notebook instance will not register with the proxy.
     */
    public readonly noProxyAccess!: pulumi.Output<boolean>;
    /**
     * If true, no public IP will be assigned to this instance.
     */
    public readonly noPublicIp!: pulumi.Output<boolean>;
    /**
     * Input only. If true, the data disk will not be auto deleted when deleting the instance.
     */
    public readonly noRemoveDataDisk!: pulumi.Output<boolean>;
    /**
     * Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path (`gs://path-to-file/file-name`).
     */
    public readonly postStartupScript!: pulumi.Output<string>;
    /**
     * The proxy endpoint that is used to access the Jupyter notebook.
     */
    public /*out*/ readonly proxyUri!: pulumi.Output<string>;
    /**
     * Optional. The optional reservation affinity. Setting this field will apply the specified [Zonal Compute Reservation](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources) to this notebook instance.
     */
    public readonly reservationAffinity!: pulumi.Output<outputs.notebooks.v1.ReservationAffinityResponse>;
    /**
     * The service account on this instance, giving access to other Google Cloud services. You can use any service account within the same project, but you must have the service account user permission to use the instance. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * Optional. The URIs of service account scopes to be included in Compute Engine instances. If not specified, the following [scopes](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam) are defined: - https://www.googleapis.com/auth/cloud-platform - https://www.googleapis.com/auth/userinfo.email If not using default scopes, you need at least: https://www.googleapis.com/auth/compute
     */
    public readonly serviceAccountScopes!: pulumi.Output<string[]>;
    /**
     * Optional. Shielded VM configuration. [Images using supported Shielded VM features](https://cloud.google.com/compute/docs/instances/modifying-shielded-vm).
     */
    public readonly shieldedInstanceConfig!: pulumi.Output<outputs.notebooks.v1.ShieldedInstanceConfigResponse>;
    /**
     * The state of this instance.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The name of the subnet that this instance is in. Format: `projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}`
     */
    public readonly subnet!: pulumi.Output<string>;
    /**
     * Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * Instance update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The upgrade history of this instance.
     */
    public readonly upgradeHistory!: pulumi.Output<outputs.notebooks.v1.UpgradeHistoryEntryResponse[]>;
    /**
     * Use a Compute Engine VM image to start the notebook instance.
     */
    public readonly vmImage!: pulumi.Output<outputs.notebooks.v1.VmImageResponse>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.machineType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'machineType'");
            }
            resourceInputs["acceleratorConfig"] = args ? args.acceleratorConfig : undefined;
            resourceInputs["bootDiskSizeGb"] = args ? args.bootDiskSizeGb : undefined;
            resourceInputs["bootDiskType"] = args ? args.bootDiskType : undefined;
            resourceInputs["containerImage"] = args ? args.containerImage : undefined;
            resourceInputs["customGpuDriverPath"] = args ? args.customGpuDriverPath : undefined;
            resourceInputs["dataDiskSizeGb"] = args ? args.dataDiskSizeGb : undefined;
            resourceInputs["dataDiskType"] = args ? args.dataDiskType : undefined;
            resourceInputs["diskEncryption"] = args ? args.diskEncryption : undefined;
            resourceInputs["installGpuDriver"] = args ? args.installGpuDriver : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["instanceOwners"] = args ? args.instanceOwners : undefined;
            resourceInputs["kmsKey"] = args ? args.kmsKey : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["machineType"] = args ? args.machineType : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["nicType"] = args ? args.nicType : undefined;
            resourceInputs["noProxyAccess"] = args ? args.noProxyAccess : undefined;
            resourceInputs["noPublicIp"] = args ? args.noPublicIp : undefined;
            resourceInputs["noRemoveDataDisk"] = args ? args.noRemoveDataDisk : undefined;
            resourceInputs["postStartupScript"] = args ? args.postStartupScript : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["reservationAffinity"] = args ? args.reservationAffinity : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["serviceAccountScopes"] = args ? args.serviceAccountScopes : undefined;
            resourceInputs["shieldedInstanceConfig"] = args ? args.shieldedInstanceConfig : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["upgradeHistory"] = args ? args.upgradeHistory : undefined;
            resourceInputs["vmImage"] = args ? args.vmImage : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["disks"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["proxyUri"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["acceleratorConfig"] = undefined /*out*/;
            resourceInputs["bootDiskSizeGb"] = undefined /*out*/;
            resourceInputs["bootDiskType"] = undefined /*out*/;
            resourceInputs["containerImage"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["customGpuDriverPath"] = undefined /*out*/;
            resourceInputs["dataDiskSizeGb"] = undefined /*out*/;
            resourceInputs["dataDiskType"] = undefined /*out*/;
            resourceInputs["diskEncryption"] = undefined /*out*/;
            resourceInputs["disks"] = undefined /*out*/;
            resourceInputs["installGpuDriver"] = undefined /*out*/;
            resourceInputs["instanceOwners"] = undefined /*out*/;
            resourceInputs["kmsKey"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["machineType"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["nicType"] = undefined /*out*/;
            resourceInputs["noProxyAccess"] = undefined /*out*/;
            resourceInputs["noPublicIp"] = undefined /*out*/;
            resourceInputs["noRemoveDataDisk"] = undefined /*out*/;
            resourceInputs["postStartupScript"] = undefined /*out*/;
            resourceInputs["proxyUri"] = undefined /*out*/;
            resourceInputs["reservationAffinity"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["serviceAccountScopes"] = undefined /*out*/;
            resourceInputs["shieldedInstanceConfig"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["subnet"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["upgradeHistory"] = undefined /*out*/;
            resourceInputs["vmImage"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The hardware accelerator used on this instance. If you use accelerators, make sure that your configuration has [enough vCPUs and memory to support the `machine_type` you have selected](/compute/docs/gpus/#gpus-list).
     */
    acceleratorConfig?: pulumi.Input<inputs.notebooks.v1.AcceleratorConfigArgs>;
    /**
     * Input only. The size of the boot disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB. If not specified, this defaults to 100.
     */
    bootDiskSizeGb?: pulumi.Input<string>;
    /**
     * Input only. The type of the boot disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
     */
    bootDiskType?: pulumi.Input<enums.notebooks.v1.InstanceBootDiskType>;
    /**
     * Use a container image to start the notebook instance.
     */
    containerImage?: pulumi.Input<inputs.notebooks.v1.ContainerImageArgs>;
    /**
     * Specify a custom Cloud Storage path where the GPU driver is stored. If not specified, we'll automatically choose from official GPU drivers.
     */
    customGpuDriverPath?: pulumi.Input<string>;
    /**
     * Input only. The size of the data disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). You can choose the size of the data disk based on how big your notebooks and data are. If not specified, this defaults to 100.
     */
    dataDiskSizeGb?: pulumi.Input<string>;
    /**
     * Input only. The type of the data disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
     */
    dataDiskType?: pulumi.Input<enums.notebooks.v1.InstanceDataDiskType>;
    /**
     * Input only. Disk encryption method used on the boot and data disks, defaults to GMEK.
     */
    diskEncryption?: pulumi.Input<enums.notebooks.v1.InstanceDiskEncryption>;
    /**
     * Whether the end user authorizes Google Cloud to install GPU driver on this instance. If this field is empty or set to false, the GPU driver won't be installed. Only applicable to instances with GPUs.
     */
    installGpuDriver?: pulumi.Input<boolean>;
    instanceId: pulumi.Input<string>;
    /**
     * Input only. The owner of this instance after creation. Format: `alias@example.com` Currently supports one owner only. If not specified, all of the service account users of your VM instance's service account can use the instance.
     */
    instanceOwners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Input only. The KMS key used to encrypt the disks, only applicable if disk_encryption is CMEK. Format: `projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}` Learn more about [using your own encryption keys](/kms/docs/quickstart).
     */
    kmsKey?: pulumi.Input<string>;
    /**
     * Labels to apply to this instance. These can be later modified by the setLabels method.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The [Compute Engine machine type](/compute/docs/machine-types) of this instance.
     */
    machineType: pulumi.Input<string>;
    /**
     * Custom metadata to apply to this instance.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the VPC that this instance is in. Format: `projects/{project_id}/global/networks/{network_id}`
     */
    network?: pulumi.Input<string>;
    /**
     * Optional. The type of vNIC to be used on this interface. This may be gVNIC or VirtioNet.
     */
    nicType?: pulumi.Input<enums.notebooks.v1.InstanceNicType>;
    /**
     * If true, the notebook instance will not register with the proxy.
     */
    noProxyAccess?: pulumi.Input<boolean>;
    /**
     * If true, no public IP will be assigned to this instance.
     */
    noPublicIp?: pulumi.Input<boolean>;
    /**
     * Input only. If true, the data disk will not be auto deleted when deleting the instance.
     */
    noRemoveDataDisk?: pulumi.Input<boolean>;
    /**
     * Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path (`gs://path-to-file/file-name`).
     */
    postStartupScript?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. The optional reservation affinity. Setting this field will apply the specified [Zonal Compute Reservation](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources) to this notebook instance.
     */
    reservationAffinity?: pulumi.Input<inputs.notebooks.v1.ReservationAffinityArgs>;
    /**
     * The service account on this instance, giving access to other Google Cloud services. You can use any service account within the same project, but you must have the service account user permission to use the instance. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Optional. The URIs of service account scopes to be included in Compute Engine instances. If not specified, the following [scopes](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam) are defined: - https://www.googleapis.com/auth/cloud-platform - https://www.googleapis.com/auth/userinfo.email If not using default scopes, you need at least: https://www.googleapis.com/auth/compute
     */
    serviceAccountScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. Shielded VM configuration. [Images using supported Shielded VM features](https://cloud.google.com/compute/docs/instances/modifying-shielded-vm).
     */
    shieldedInstanceConfig?: pulumi.Input<inputs.notebooks.v1.ShieldedInstanceConfigArgs>;
    /**
     * The name of the subnet that this instance is in. Format: `projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}`
     */
    subnet?: pulumi.Input<string>;
    /**
     * Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The upgrade history of this instance.
     */
    upgradeHistory?: pulumi.Input<pulumi.Input<inputs.notebooks.v1.UpgradeHistoryEntryArgs>[]>;
    /**
     * Use a Compute Engine VM image to start the notebook instance.
     */
    vmImage?: pulumi.Input<inputs.notebooks.v1.VmImageArgs>;
}
