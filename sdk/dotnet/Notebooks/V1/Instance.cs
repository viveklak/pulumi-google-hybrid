// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V1
{
    /// <summary>
    /// Creates a new Instance in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:notebooks/v1:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// The hardware accelerator used on this instance. If you use accelerators, make sure that your configuration has [enough vCPUs and memory to support the `machine_type` you have selected](/compute/docs/gpus/#gpus-list).
        /// </summary>
        [Output("acceleratorConfig")]
        public Output<Outputs.AcceleratorConfigResponse> AcceleratorConfig { get; private set; } = null!;

        /// <summary>
        /// Input only. The size of the boot disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB. If not specified, this defaults to 100.
        /// </summary>
        [Output("bootDiskSizeGb")]
        public Output<string> BootDiskSizeGb { get; private set; } = null!;

        /// <summary>
        /// Input only. The type of the boot disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
        /// </summary>
        [Output("bootDiskType")]
        public Output<string> BootDiskType { get; private set; } = null!;

        /// <summary>
        /// Use a container image to start the notebook instance.
        /// </summary>
        [Output("containerImage")]
        public Output<Outputs.ContainerImageResponse> ContainerImage { get; private set; } = null!;

        /// <summary>
        /// Instance creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Specify a custom Cloud Storage path where the GPU driver is stored. If not specified, we'll automatically choose from official GPU drivers.
        /// </summary>
        [Output("customGpuDriverPath")]
        public Output<string> CustomGpuDriverPath { get; private set; } = null!;

        /// <summary>
        /// Input only. The size of the data disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). You can choose the size of the data disk based on how big your notebooks and data are. If not specified, this defaults to 100.
        /// </summary>
        [Output("dataDiskSizeGb")]
        public Output<string> DataDiskSizeGb { get; private set; } = null!;

        /// <summary>
        /// Input only. The type of the data disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
        /// </summary>
        [Output("dataDiskType")]
        public Output<string> DataDiskType { get; private set; } = null!;

        /// <summary>
        /// Input only. Disk encryption method used on the boot and data disks, defaults to GMEK.
        /// </summary>
        [Output("diskEncryption")]
        public Output<string> DiskEncryption { get; private set; } = null!;

        /// <summary>
        /// Attached disks to notebook instance.
        /// </summary>
        [Output("disks")]
        public Output<ImmutableArray<Outputs.DiskResponse>> Disks { get; private set; } = null!;

        /// <summary>
        /// Whether the end user authorizes Google Cloud to install GPU driver on this instance. If this field is empty or set to false, the GPU driver won't be installed. Only applicable to instances with GPUs.
        /// </summary>
        [Output("installGpuDriver")]
        public Output<bool> InstallGpuDriver { get; private set; } = null!;

        /// <summary>
        /// Input only. The owner of this instance after creation. Format: `alias@example.com` Currently supports one owner only. If not specified, all of the service account users of your VM instance's service account can use the instance.
        /// </summary>
        [Output("instanceOwners")]
        public Output<ImmutableArray<string>> InstanceOwners { get; private set; } = null!;

        /// <summary>
        /// Input only. The KMS key used to encrypt the disks, only applicable if disk_encryption is CMEK. Format: `projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}` Learn more about [using your own encryption keys](/kms/docs/quickstart).
        /// </summary>
        [Output("kmsKey")]
        public Output<string> KmsKey { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this instance. These can be later modified by the setLabels method.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The [Compute Engine machine type](/compute/docs/machine-types) of this instance.
        /// </summary>
        [Output("machineType")]
        public Output<string> MachineType { get; private set; } = null!;

        /// <summary>
        /// Custom metadata to apply to this instance.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// The name of this notebook instance. Format: `projects/{project_id}/locations/{location}/instances/{instance_id}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the VPC that this instance is in. Format: `projects/{project_id}/global/networks/{network_id}`
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Optional. The type of vNIC to be used on this interface. This may be gVNIC or VirtioNet.
        /// </summary>
        [Output("nicType")]
        public Output<string> NicType { get; private set; } = null!;

        /// <summary>
        /// If true, the notebook instance will not register with the proxy.
        /// </summary>
        [Output("noProxyAccess")]
        public Output<bool> NoProxyAccess { get; private set; } = null!;

        /// <summary>
        /// If true, no public IP will be assigned to this instance.
        /// </summary>
        [Output("noPublicIp")]
        public Output<bool> NoPublicIp { get; private set; } = null!;

        /// <summary>
        /// Input only. If true, the data disk will not be auto deleted when deleting the instance.
        /// </summary>
        [Output("noRemoveDataDisk")]
        public Output<bool> NoRemoveDataDisk { get; private set; } = null!;

        /// <summary>
        /// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path (gs://path-to-file/file-name).
        /// </summary>
        [Output("postStartupScript")]
        public Output<string> PostStartupScript { get; private set; } = null!;

        /// <summary>
        /// The proxy endpoint that is used to access the Jupyter notebook.
        /// </summary>
        [Output("proxyUri")]
        public Output<string> ProxyUri { get; private set; } = null!;

        /// <summary>
        /// The service account on this instance, giving access to other Google Cloud services. You can use any service account within the same project, but you must have the service account user permission to use the instance. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Optional. The URIs of service account scopes to be included in Compute Engine instances. If not specified, the following [scopes](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam) are defined: - https://www.googleapis.com/auth/cloud-platform - https://www.googleapis.com/auth/userinfo.email If not using default scopes, you need at least: https://www.googleapis.com/auth/compute
        /// </summary>
        [Output("serviceAccountScopes")]
        public Output<ImmutableArray<string>> ServiceAccountScopes { get; private set; } = null!;

        /// <summary>
        /// Optional. Shielded VM configuration. [Images using supported Shielded VM features] (https://cloud.google.com/compute/docs/instances/modifying-shielded-vm).
        /// </summary>
        [Output("shieldedInstanceConfig")]
        public Output<Outputs.ShieldedInstanceConfigResponse> ShieldedInstanceConfig { get; private set; } = null!;

        /// <summary>
        /// The state of this instance.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The name of the subnet that this instance is in. Format: `projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}`
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Instance update time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The upgrade history of this instance.
        /// </summary>
        [Output("upgradeHistory")]
        public Output<ImmutableArray<Outputs.UpgradeHistoryEntryResponse>> UpgradeHistory { get; private set; } = null!;

        /// <summary>
        /// Use a Compute Engine VM image to start the notebook instance.
        /// </summary>
        [Output("vmImage")]
        public Output<Outputs.VmImageResponse> VmImage { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("google-native:notebooks/v1:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:notebooks/v1:Instance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hardware accelerator used on this instance. If you use accelerators, make sure that your configuration has [enough vCPUs and memory to support the `machine_type` you have selected](/compute/docs/gpus/#gpus-list).
        /// </summary>
        [Input("acceleratorConfig")]
        public Input<Inputs.AcceleratorConfigArgs>? AcceleratorConfig { get; set; }

        /// <summary>
        /// Input only. The size of the boot disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB. If not specified, this defaults to 100.
        /// </summary>
        [Input("bootDiskSizeGb")]
        public Input<string>? BootDiskSizeGb { get; set; }

        /// <summary>
        /// Input only. The type of the boot disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
        /// </summary>
        [Input("bootDiskType")]
        public Input<Pulumi.GoogleNative.Notebooks.V1.InstanceBootDiskType>? BootDiskType { get; set; }

        /// <summary>
        /// Use a container image to start the notebook instance.
        /// </summary>
        [Input("containerImage")]
        public Input<Inputs.ContainerImageArgs>? ContainerImage { get; set; }

        /// <summary>
        /// Specify a custom Cloud Storage path where the GPU driver is stored. If not specified, we'll automatically choose from official GPU drivers.
        /// </summary>
        [Input("customGpuDriverPath")]
        public Input<string>? CustomGpuDriverPath { get; set; }

        /// <summary>
        /// Input only. The size of the data disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB). You can choose the size of the data disk based on how big your notebooks and data are. If not specified, this defaults to 100.
        /// </summary>
        [Input("dataDiskSizeGb")]
        public Input<string>? DataDiskSizeGb { get; set; }

        /// <summary>
        /// Input only. The type of the data disk attached to this instance, defaults to standard persistent disk (`PD_STANDARD`).
        /// </summary>
        [Input("dataDiskType")]
        public Input<Pulumi.GoogleNative.Notebooks.V1.InstanceDataDiskType>? DataDiskType { get; set; }

        /// <summary>
        /// Input only. Disk encryption method used on the boot and data disks, defaults to GMEK.
        /// </summary>
        [Input("diskEncryption")]
        public Input<Pulumi.GoogleNative.Notebooks.V1.InstanceDiskEncryption>? DiskEncryption { get; set; }

        /// <summary>
        /// Whether the end user authorizes Google Cloud to install GPU driver on this instance. If this field is empty or set to false, the GPU driver won't be installed. Only applicable to instances with GPUs.
        /// </summary>
        [Input("installGpuDriver")]
        public Input<bool>? InstallGpuDriver { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("instanceOwners")]
        private InputList<string>? _instanceOwners;

        /// <summary>
        /// Input only. The owner of this instance after creation. Format: `alias@example.com` Currently supports one owner only. If not specified, all of the service account users of your VM instance's service account can use the instance.
        /// </summary>
        public InputList<string> InstanceOwners
        {
            get => _instanceOwners ?? (_instanceOwners = new InputList<string>());
            set => _instanceOwners = value;
        }

        /// <summary>
        /// Input only. The KMS key used to encrypt the disks, only applicable if disk_encryption is CMEK. Format: `projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}` Learn more about [using your own encryption keys](/kms/docs/quickstart).
        /// </summary>
        [Input("kmsKey")]
        public Input<string>? KmsKey { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this instance. These can be later modified by the setLabels method.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The [Compute Engine machine type](/compute/docs/machine-types) of this instance.
        /// </summary>
        [Input("machineType", required: true)]
        public Input<string> MachineType { get; set; } = null!;

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Custom metadata to apply to this instance.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the VPC that this instance is in. Format: `projects/{project_id}/global/networks/{network_id}`
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Optional. The type of vNIC to be used on this interface. This may be gVNIC or VirtioNet.
        /// </summary>
        [Input("nicType")]
        public Input<Pulumi.GoogleNative.Notebooks.V1.InstanceNicType>? NicType { get; set; }

        /// <summary>
        /// If true, the notebook instance will not register with the proxy.
        /// </summary>
        [Input("noProxyAccess")]
        public Input<bool>? NoProxyAccess { get; set; }

        /// <summary>
        /// If true, no public IP will be assigned to this instance.
        /// </summary>
        [Input("noPublicIp")]
        public Input<bool>? NoPublicIp { get; set; }

        /// <summary>
        /// Input only. If true, the data disk will not be auto deleted when deleting the instance.
        /// </summary>
        [Input("noRemoveDataDisk")]
        public Input<bool>? NoRemoveDataDisk { get; set; }

        /// <summary>
        /// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path (gs://path-to-file/file-name).
        /// </summary>
        [Input("postStartupScript")]
        public Input<string>? PostStartupScript { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The service account on this instance, giving access to other Google Cloud services. You can use any service account within the same project, but you must have the service account user permission to use the instance. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("serviceAccountScopes")]
        private InputList<string>? _serviceAccountScopes;

        /// <summary>
        /// Optional. The URIs of service account scopes to be included in Compute Engine instances. If not specified, the following [scopes](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam) are defined: - https://www.googleapis.com/auth/cloud-platform - https://www.googleapis.com/auth/userinfo.email If not using default scopes, you need at least: https://www.googleapis.com/auth/compute
        /// </summary>
        public InputList<string> ServiceAccountScopes
        {
            get => _serviceAccountScopes ?? (_serviceAccountScopes = new InputList<string>());
            set => _serviceAccountScopes = value;
        }

        /// <summary>
        /// Optional. Shielded VM configuration. [Images using supported Shielded VM features] (https://cloud.google.com/compute/docs/instances/modifying-shielded-vm).
        /// </summary>
        [Input("shieldedInstanceConfig")]
        public Input<Inputs.ShieldedInstanceConfigArgs>? ShieldedInstanceConfig { get; set; }

        /// <summary>
        /// The name of the subnet that this instance is in. Format: `projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}`
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("upgradeHistory")]
        private InputList<Inputs.UpgradeHistoryEntryArgs>? _upgradeHistory;

        /// <summary>
        /// The upgrade history of this instance.
        /// </summary>
        public InputList<Inputs.UpgradeHistoryEntryArgs> UpgradeHistory
        {
            get => _upgradeHistory ?? (_upgradeHistory = new InputList<Inputs.UpgradeHistoryEntryArgs>());
            set => _upgradeHistory = value;
        }

        /// <summary>
        /// Use a Compute Engine VM image to start the notebook instance.
        /// </summary>
        [Input("vmImage")]
        public Input<Inputs.VmImageArgs>? VmImage { get; set; }

        public InstanceArgs()
        {
        }
    }
}
