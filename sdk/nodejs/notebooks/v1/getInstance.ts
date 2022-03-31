// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Instance.
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:notebooks/v1:getInstance", {
        "instanceId": args.instanceId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetInstanceArgs {
    instanceId: string;
    location: string;
    project?: string;
}

export interface GetInstanceResult {
    /**
     * The hardware accelerator used on this instance. If you use accelerators, make sure that your configuration has [enough vCPUs and memory to support the `machine_type` you have selected](/compute/docs/gpus/#gpus-list).
     */
    readonly acceleratorConfig: outputs.notebooks.v1.AcceleratorConfigResponse;
    /**
     * Input only. The size of the boot disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB. If not specified, this defaults to 100.
     */
    readonly bootDiskSizeGb: string;
    /**
     * Input only. The type of the boot disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
     */
    readonly bootDiskType: string;
    /**
     * Optional. Flag to enable ip forwarding or not, default false/off. https://cloud.google.com/vpc/docs/using-routes#canipforward
     */
    readonly canIpForward: boolean;
    /**
     * Use a container image to start the notebook instance.
     */
    readonly containerImage: outputs.notebooks.v1.ContainerImageResponse;
    /**
     * Instance creation time.
     */
    readonly createTime: string;
    /**
     * Email address of entity that sent original CreateInstance request.
     */
    readonly creator: string;
    /**
     * Specify a custom Cloud Storage path where the GPU driver is stored. If not specified, we'll automatically choose from official GPU drivers.
     */
    readonly customGpuDriverPath: string;
    /**
     * Input only. The size of the data disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). You can choose the size of the data disk based on how big your notebooks and data are. If not specified, this defaults to 100.
     */
    readonly dataDiskSizeGb: string;
    /**
     * Input only. The type of the data disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
     */
    readonly dataDiskType: string;
    /**
     * Input only. Disk encryption method used on the boot and data disks, defaults to GMEK.
     */
    readonly diskEncryption: string;
    /**
     * Attached disks to notebook instance.
     */
    readonly disks: outputs.notebooks.v1.DiskResponse[];
    /**
     * Whether the end user authorizes Google Cloud to install GPU driver on this instance. If this field is empty or set to false, the GPU driver won't be installed. Only applicable to instances with GPUs.
     */
    readonly installGpuDriver: boolean;
    /**
     * Input only. The owner of this instance after creation. Format: `alias@example.com` Currently supports one owner only. If not specified, all of the service account users of your VM instance's service account can use the instance.
     */
    readonly instanceOwners: string[];
    /**
     * Input only. The KMS key used to encrypt the disks, only applicable if disk_encryption is CMEK. Format: `projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}` Learn more about [using your own encryption keys](/kms/docs/quickstart).
     */
    readonly kmsKey: string;
    /**
     * Labels to apply to this instance. These can be later modified by the setLabels method.
     */
    readonly labels: {[key: string]: string};
    /**
     * The [Compute Engine machine type](/compute/docs/machine-types) of this instance.
     */
    readonly machineType: string;
    /**
     * Custom metadata to apply to this instance.
     */
    readonly metadata: {[key: string]: string};
    /**
     * The name of this notebook instance. Format: `projects/{project_id}/locations/{location}/instances/{instance_id}`
     */
    readonly name: string;
    /**
     * The name of the VPC that this instance is in. Format: `projects/{project_id}/global/networks/{network_id}`
     */
    readonly network: string;
    /**
     * Optional. The type of vNIC to be used on this interface. This may be gVNIC or VirtioNet.
     */
    readonly nicType: string;
    /**
     * If true, the notebook instance will not register with the proxy.
     */
    readonly noProxyAccess: boolean;
    /**
     * If true, no public IP will be assigned to this instance.
     */
    readonly noPublicIp: boolean;
    /**
     * Input only. If true, the data disk will not be auto deleted when deleting the instance.
     */
    readonly noRemoveDataDisk: boolean;
    /**
     * Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path (`gs://path-to-file/file-name`).
     */
    readonly postStartupScript: string;
    /**
     * The proxy endpoint that is used to access the Jupyter notebook.
     */
    readonly proxyUri: string;
    /**
     * Optional. The optional reservation affinity. Setting this field will apply the specified [Zonal Compute Reservation](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources) to this notebook instance.
     */
    readonly reservationAffinity: outputs.notebooks.v1.ReservationAffinityResponse;
    /**
     * The service account on this instance, giving access to other Google Cloud services. You can use any service account within the same project, but you must have the service account user permission to use the instance. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.
     */
    readonly serviceAccount: string;
    /**
     * Optional. The URIs of service account scopes to be included in Compute Engine instances. If not specified, the following [scopes](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam) are defined: - https://www.googleapis.com/auth/cloud-platform - https://www.googleapis.com/auth/userinfo.email If not using default scopes, you need at least: https://www.googleapis.com/auth/compute
     */
    readonly serviceAccountScopes: string[];
    /**
     * Optional. Shielded VM configuration. [Images using supported Shielded VM features](https://cloud.google.com/compute/docs/instances/modifying-shielded-vm).
     */
    readonly shieldedInstanceConfig: outputs.notebooks.v1.ShieldedInstanceConfigResponse;
    /**
     * The state of this instance.
     */
    readonly state: string;
    /**
     * The name of the subnet that this instance is in. Format: `projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}`
     */
    readonly subnet: string;
    /**
     * Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
     */
    readonly tags: string[];
    /**
     * Instance update time.
     */
    readonly updateTime: string;
    /**
     * The upgrade history of this instance.
     */
    readonly upgradeHistory: outputs.notebooks.v1.UpgradeHistoryEntryResponse[];
    /**
     * Use a Compute Engine VM image to start the notebook instance.
     */
    readonly vmImage: outputs.notebooks.v1.VmImageResponse;
}

export function getInstanceOutput(args: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply(a => getInstance(a, opts))
}

export interface GetInstanceOutputArgs {
    instanceId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
