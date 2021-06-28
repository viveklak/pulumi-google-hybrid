// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    /// <summary>
    /// Creates an instance resource in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/beta:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// Controls for advanced machine-related behavior features.
        /// </summary>
        [Output("advancedMachineFeatures")]
        public Output<Outputs.AdvancedMachineFeaturesResponse> AdvancedMachineFeatures { get; private set; } = null!;

        /// <summary>
        /// Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding.
        /// </summary>
        [Output("canIpForward")]
        public Output<bool> CanIpForward { get; private set; } = null!;

        [Output("confidentialInstanceConfig")]
        public Output<Outputs.ConfidentialInstanceConfigResponse> ConfidentialInstanceConfig { get; private set; } = null!;

        /// <summary>
        /// The CPU platform used by this instance.
        /// </summary>
        [Output("cpuPlatform")]
        public Output<string> CpuPlatform { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// Whether the resource should be protected against deletion.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Array of disks associated with this instance. Persistent disks must be created before you can assign them.
        /// </summary>
        [Output("disks")]
        public Output<ImmutableArray<Outputs.AttachedDiskResponse>> Disks { get; private set; } = null!;

        /// <summary>
        /// Enables display device for the instance.
        /// </summary>
        [Output("displayDevice")]
        public Output<Outputs.DisplayDeviceResponse> DisplayDevice { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
        /// </summary>
        [Output("eraseWindowsVssSignature")]
        public Output<bool> EraseWindowsVssSignature { get; private set; } = null!;

        /// <summary>
        /// Specifies a fingerprint for this resource, which is essentially a hash of the instance's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update the instance. You must always provide an up-to-date fingerprint hash in order to update the instance.
        /// 
        /// To see the latest fingerprint, make get() request to the instance.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// A list of the type and count of accelerator cards attached to the instance.
        /// </summary>
        [Output("guestAccelerators")]
        public Output<ImmutableArray<Outputs.AcceleratorConfigResponse>> GuestAccelerators { get; private set; } = null!;

        /// <summary>
        /// Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// Type of the resource. Always compute#instance for instances.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A fingerprint for this request, which is essentially a hash of the label's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels.
        /// 
        /// To see the latest fingerprint, make get() request to the instance.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this instance. These can be later modified by the setLabels method.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Last start timestamp in RFC3339 text format.
        /// </summary>
        [Output("lastStartTimestamp")]
        public Output<string> LastStartTimestamp { get; private set; } = null!;

        /// <summary>
        /// Last stop timestamp in RFC3339 text format.
        /// </summary>
        [Output("lastStopTimestamp")]
        public Output<string> LastStopTimestamp { get; private set; } = null!;

        /// <summary>
        /// Last suspended timestamp in RFC3339 text format.
        /// </summary>
        [Output("lastSuspendedTimestamp")]
        public Output<string> LastSuspendedTimestamp { get; private set; } = null!;

        /// <summary>
        /// Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type:
        /// zones/us-central1-f/machineTypes/n1-standard-1
        /// 
        /// 
        /// To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB):
        /// zones/zone/machineTypes/custom-CPUS-MEMORY
        /// 
        /// 
        /// For example: zones/us-central1-f/machineTypes/custom-4-5120 
        /// 
        /// For a full list of restrictions, read the Specifications for custom machine types.
        /// </summary>
        [Output("machineType")]
        public Output<string> MachineType { get; private set; } = null!;

        /// <summary>
        /// The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
        /// </summary>
        [Output("metadata")]
        public Output<Outputs.MetadataResponse> Metadata { get; private set; } = null!;

        /// <summary>
        /// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
        /// </summary>
        [Output("minCpuPlatform")]
        public Output<string> MinCpuPlatform { get; private set; } = null!;

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
        /// </summary>
        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.NetworkInterfaceResponse>> NetworkInterfaces { get; private set; } = null!;

        [Output("networkPerformanceConfig")]
        public Output<Outputs.NetworkPerformanceConfigResponse> NetworkPerformanceConfig { get; private set; } = null!;

        /// <summary>
        /// PostKeyRevocationActionType of the instance.
        /// </summary>
        [Output("postKeyRevocationActionType")]
        public Output<string> PostKeyRevocationActionType { get; private set; } = null!;

        /// <summary>
        /// The private IPv6 google access type for the VM. If not specified, use  INHERIT_FROM_SUBNETWORK as default.
        /// </summary>
        [Output("privateIpv6GoogleAccess")]
        public Output<string> PrivateIpv6GoogleAccess { get; private set; } = null!;

        /// <summary>
        /// Specifies the reservations that this instance can consume from.
        /// </summary>
        [Output("reservationAffinity")]
        public Output<Outputs.ReservationAffinityResponse> ReservationAffinity { get; private set; } = null!;

        /// <summary>
        /// Resource policies applied to this instance.
        /// </summary>
        [Output("resourcePolicies")]
        public Output<ImmutableArray<string>> ResourcePolicies { get; private set; } = null!;

        /// <summary>
        /// Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// Sets the scheduling options for this instance.
        /// </summary>
        [Output("scheduling")]
        public Output<Outputs.SchedulingResponse> Scheduling { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for this resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported.
        /// 
        /// Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
        /// </summary>
        [Output("serviceAccounts")]
        public Output<ImmutableArray<Outputs.ServiceAccountResponse>> ServiceAccounts { get; private set; } = null!;

        [Output("shieldedInstanceConfig")]
        public Output<Outputs.ShieldedInstanceConfigResponse> ShieldedInstanceConfig { get; private set; } = null!;

        [Output("shieldedInstanceIntegrityPolicy")]
        public Output<Outputs.ShieldedInstanceIntegrityPolicyResponse> ShieldedInstanceIntegrityPolicy { get; private set; } = null!;

        /// <summary>
        /// Deprecating, please use shielded_instance_config.
        /// </summary>
        [Output("shieldedVmConfig")]
        public Output<Outputs.ShieldedVmConfigResponse> ShieldedVmConfig { get; private set; } = null!;

        /// <summary>
        /// Deprecating, please use shielded_instance_integrity_policy.
        /// </summary>
        [Output("shieldedVmIntegrityPolicy")]
        public Output<Outputs.ShieldedVmIntegrityPolicyResponse> ShieldedVmIntegrityPolicy { get; private set; } = null!;

        /// <summary>
        /// Source machine image
        /// </summary>
        [Output("sourceMachineImage")]
        public Output<string> SourceMachineImage { get; private set; } = null!;

        /// <summary>
        /// Source machine image encryption key when creating an instance from a machine image.
        /// </summary>
        [Output("sourceMachineImageEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> SourceMachineImageEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// Whether a VM has been restricted for start because Compute Engine has detected suspicious activity.
        /// </summary>
        [Output("startRestricted")]
        public Output<bool> StartRestricted { get; private set; } = null!;

        /// <summary>
        /// The status of the instance. One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see  Instance life cycle.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// An optional, human-readable explanation of the status.
        /// </summary>
        [Output("statusMessage")]
        public Output<string> StatusMessage { get; private set; } = null!;

        /// <summary>
        /// Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
        /// </summary>
        [Output("tags")]
        public Output<Outputs.TagsResponse> Tags { get; private set; } = null!;

        /// <summary>
        /// URL of the zone where the instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:Instance", name, null, MakeResourceOptions(options, id))
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
        /// Controls for advanced machine-related behavior features.
        /// </summary>
        [Input("advancedMachineFeatures")]
        public Input<Inputs.AdvancedMachineFeaturesArgs>? AdvancedMachineFeatures { get; set; }

        /// <summary>
        /// Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding.
        /// </summary>
        [Input("canIpForward")]
        public Input<bool>? CanIpForward { get; set; }

        [Input("confidentialInstanceConfig")]
        public Input<Inputs.ConfidentialInstanceConfigArgs>? ConfidentialInstanceConfig { get; set; }

        /// <summary>
        /// Whether the resource should be protected against deletion.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disks")]
        private InputList<Inputs.AttachedDiskArgs>? _disks;

        /// <summary>
        /// Array of disks associated with this instance. Persistent disks must be created before you can assign them.
        /// </summary>
        public InputList<Inputs.AttachedDiskArgs> Disks
        {
            get => _disks ?? (_disks = new InputList<Inputs.AttachedDiskArgs>());
            set => _disks = value;
        }

        /// <summary>
        /// Enables display device for the instance.
        /// </summary>
        [Input("displayDevice")]
        public Input<Inputs.DisplayDeviceArgs>? DisplayDevice { get; set; }

        /// <summary>
        /// Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
        /// </summary>
        [Input("eraseWindowsVssSignature")]
        public Input<bool>? EraseWindowsVssSignature { get; set; }

        [Input("guestAccelerators")]
        private InputList<Inputs.AcceleratorConfigArgs>? _guestAccelerators;

        /// <summary>
        /// A list of the type and count of accelerator cards attached to the instance.
        /// </summary>
        public InputList<Inputs.AcceleratorConfigArgs> GuestAccelerators
        {
            get => _guestAccelerators ?? (_guestAccelerators = new InputList<Inputs.AcceleratorConfigArgs>());
            set => _guestAccelerators = value;
        }

        /// <summary>
        /// Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

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

        /// <summary>
        /// Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type:
        /// zones/us-central1-f/machineTypes/n1-standard-1
        /// 
        /// 
        /// To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB):
        /// zones/zone/machineTypes/custom-CPUS-MEMORY
        /// 
        /// 
        /// For example: zones/us-central1-f/machineTypes/custom-4-5120 
        /// 
        /// For a full list of restrictions, read the Specifications for custom machine types.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        /// <summary>
        /// The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
        /// </summary>
        [Input("metadata")]
        public Input<Inputs.MetadataArgs>? Metadata { get; set; }

        /// <summary>
        /// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
        /// </summary>
        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.NetworkInterfaceArgs>? _networkInterfaces;

        /// <summary>
        /// An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.NetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        [Input("networkPerformanceConfig")]
        public Input<Inputs.NetworkPerformanceConfigArgs>? NetworkPerformanceConfig { get; set; }

        /// <summary>
        /// PostKeyRevocationActionType of the instance.
        /// </summary>
        [Input("postKeyRevocationActionType")]
        public Input<Pulumi.GoogleNative.Compute.Beta.InstancePostKeyRevocationActionType>? PostKeyRevocationActionType { get; set; }

        /// <summary>
        /// The private IPv6 google access type for the VM. If not specified, use  INHERIT_FROM_SUBNETWORK as default.
        /// </summary>
        [Input("privateIpv6GoogleAccess")]
        public Input<Pulumi.GoogleNative.Compute.Beta.InstancePrivateIpv6GoogleAccess>? PrivateIpv6GoogleAccess { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Specifies the reservations that this instance can consume from.
        /// </summary>
        [Input("reservationAffinity")]
        public Input<Inputs.ReservationAffinityArgs>? ReservationAffinity { get; set; }

        [Input("resourcePolicies")]
        private InputList<string>? _resourcePolicies;

        /// <summary>
        /// Resource policies applied to this instance.
        /// </summary>
        public InputList<string> ResourcePolicies
        {
            get => _resourcePolicies ?? (_resourcePolicies = new InputList<string>());
            set => _resourcePolicies = value;
        }

        /// <summary>
        /// Sets the scheduling options for this instance.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.SchedulingArgs>? Scheduling { get; set; }

        [Input("serviceAccounts")]
        private InputList<Inputs.ServiceAccountArgs>? _serviceAccounts;

        /// <summary>
        /// A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported.
        /// 
        /// Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
        /// </summary>
        public InputList<Inputs.ServiceAccountArgs> ServiceAccounts
        {
            get => _serviceAccounts ?? (_serviceAccounts = new InputList<Inputs.ServiceAccountArgs>());
            set => _serviceAccounts = value;
        }

        [Input("shieldedInstanceConfig")]
        public Input<Inputs.ShieldedInstanceConfigArgs>? ShieldedInstanceConfig { get; set; }

        [Input("shieldedInstanceIntegrityPolicy")]
        public Input<Inputs.ShieldedInstanceIntegrityPolicyArgs>? ShieldedInstanceIntegrityPolicy { get; set; }

        /// <summary>
        /// Deprecating, please use shielded_instance_config.
        /// </summary>
        [Input("shieldedVmConfig")]
        public Input<Inputs.ShieldedVmConfigArgs>? ShieldedVmConfig { get; set; }

        /// <summary>
        /// Deprecating, please use shielded_instance_integrity_policy.
        /// </summary>
        [Input("shieldedVmIntegrityPolicy")]
        public Input<Inputs.ShieldedVmIntegrityPolicyArgs>? ShieldedVmIntegrityPolicy { get; set; }

        [Input("sourceInstanceTemplate")]
        public Input<string>? SourceInstanceTemplate { get; set; }

        /// <summary>
        /// Source machine image
        /// </summary>
        [Input("sourceMachineImage")]
        public Input<string>? SourceMachineImage { get; set; }

        /// <summary>
        /// Source machine image encryption key when creating an instance from a machine image.
        /// </summary>
        [Input("sourceMachineImageEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? SourceMachineImageEncryptionKey { get; set; }

        /// <summary>
        /// Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
        /// </summary>
        [Input("tags")]
        public Input<Inputs.TagsArgs>? Tags { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public InstanceArgs()
        {
        }
    }
}
