// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Alpha
{
    /// <summary>
    /// Creates an instance resource in the specified project using the data included in the request.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:compute/alpha:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:compute/alpha:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:compute/alpha:Instance", name, null, MakeResourceOptions(options, id))
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
        /// [Output Only] The CPU platform used by this instance.
        /// </summary>
        [Input("cpuPlatform")]
        public Input<string>? CpuPlatform { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

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

        /// <summary>
        /// Specifies a fingerprint for this resource, which is essentially a hash of the instance's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update the instance. You must always provide an up-to-date fingerprint hash in order to update the instance.
        /// 
        /// To see the latest fingerprint, make get() request to the instance.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

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

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// Encrypts or decrypts data for an instance with a customer-supplied encryption key.
        /// 
        /// If you are creating a new instance, this field encrypts the local SSD and in-memory contents of the instance using a key that you provide.
        /// 
        /// If you are restarting an instance protected with a customer-supplied encryption key, you must provide the correct key in order to successfully restart the instance.
        /// 
        /// If you do not provide an encryption key when creating the instance, then the local SSD and in-memory contents will be encrypted using an automatically generated key and you do not need to provide a key to start the instance later.
        /// 
        /// Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt local SSDs and in-memory content in a managed instance group.
        /// </summary>
        [Input("instanceEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? InstanceEncryptionKey { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#instance for instances.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// A fingerprint for this request, which is essentially a hash of the label's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels.
        /// 
        /// To see the latest fingerprint, make get() request to the instance.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

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
        /// [Output Only] Last start timestamp in RFC3339 text format.
        /// </summary>
        [Input("lastStartTimestamp")]
        public Input<string>? LastStartTimestamp { get; set; }

        /// <summary>
        /// [Output Only] Last stop timestamp in RFC3339 text format.
        /// </summary>
        [Input("lastStopTimestamp")]
        public Input<string>? LastStopTimestamp { get; set; }

        /// <summary>
        /// [Output Only] Last suspended timestamp in RFC3339 text format.
        /// </summary>
        [Input("lastSuspendedTimestamp")]
        public Input<string>? LastSuspendedTimestamp { get; set; }

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
        public Input<string>? PostKeyRevocationActionType { get; set; }

        /// <summary>
        /// Total amount of preserved state for SUSPENDED instances. Read-only in the api.
        /// </summary>
        [Input("preservedStateSizeGb")]
        public Input<string>? PreservedStateSizeGb { get; set; }

        /// <summary>
        /// The private IPv6 google access type for the VM. If not specified, use  INHERIT_FROM_SUBNETWORK as default.
        /// </summary>
        [Input("privateIpv6GoogleAccess")]
        public Input<string>? PrivateIpv6GoogleAccess { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

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
        /// [Output Only] Specifies values set for instance attributes as compared to the values requested by user in the corresponding input only field.
        /// </summary>
        [Input("resourceStatus")]
        public Input<Inputs.ResourceStatusArgs>? ResourceStatus { get; set; }

        /// <summary>
        /// [Output Only] Reserved for future use.
        /// </summary>
        [Input("satisfiesPzs")]
        public Input<bool>? SatisfiesPzs { get; set; }

        /// <summary>
        /// Sets the scheduling options for this instance.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.SchedulingArgs>? Scheduling { get; set; }

        [Input("secureLabels")]
        private InputList<string>? _secureLabels;

        /// <summary>
        /// Secure labels to apply to this instance. These can be later modified by the update method. Maximum number of secure labels allowed is 300.
        /// </summary>
        public InputList<string> SecureLabels
        {
            get => _secureLabels ?? (_secureLabels = new InputList<string>());
            set => _secureLabels = value;
        }

        [Input("secureTags")]
        private InputList<string>? _secureTags;

        /// <summary>
        /// Secure tags to apply to this instance. These can be later modified by the update method. Maximum number of secure tags allowed is 300.
        /// </summary>
        public InputList<string> SecureTags
        {
            get => _secureTags ?? (_secureTags = new InputList<string>());
            set => _secureTags = value;
        }

        /// <summary>
        /// [Output Only] Server-defined URL for this resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Input("selfLinkWithId")]
        public Input<string>? SelfLinkWithId { get; set; }

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
        /// [Output Only] Whether a VM has been restricted for start because Compute Engine has detected suspicious activity.
        /// </summary>
        [Input("startRestricted")]
        public Input<bool>? StartRestricted { get; set; }

        /// <summary>
        /// [Output Only] The status of the instance. One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see  Instance life cycle.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// [Output Only] An optional, human-readable explanation of the status.
        /// </summary>
        [Input("statusMessage")]
        public Input<string>? StatusMessage { get; set; }

        /// <summary>
        /// Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
        /// </summary>
        [Input("tags")]
        public Input<Inputs.TagsArgs>? Tags { get; set; }

        /// <summary>
        /// [Output Only] Specifies upcoming maintenance for the instance.
        /// </summary>
        [Input("upcomingMaintenance")]
        public Input<Inputs.UpcomingMaintenanceArgs>? UpcomingMaintenance { get; set; }

        /// <summary>
        /// [Output Only] URL of the zone where the instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public InstanceArgs()
        {
        }
    }
}
