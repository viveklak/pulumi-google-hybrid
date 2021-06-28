// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetInstance
    {
        /// <summary>
        /// Returns the specified Instance resource. Gets a list of available instances by making a list() request.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("google-native:compute/alpha:getInstance", args ?? new GetInstanceArgs(), options.WithVersion());
    }


    public sealed class GetInstanceArgs : Pulumi.InvokeArgs
    {
        [Input("instance", required: true)]
        public string Instance { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetInstanceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// Controls for advanced machine-related behavior features.
        /// </summary>
        public readonly Outputs.AdvancedMachineFeaturesResponse AdvancedMachineFeatures;
        /// <summary>
        /// Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding.
        /// </summary>
        public readonly bool CanIpForward;
        public readonly Outputs.ConfidentialInstanceConfigResponse ConfidentialInstanceConfig;
        /// <summary>
        /// The CPU platform used by this instance.
        /// </summary>
        public readonly string CpuPlatform;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// Whether the resource should be protected against deletion.
        /// </summary>
        public readonly bool DeletionProtection;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Array of disks associated with this instance. Persistent disks must be created before you can assign them.
        /// </summary>
        public readonly ImmutableArray<Outputs.AttachedDiskResponse> Disks;
        /// <summary>
        /// Enables display device for the instance.
        /// </summary>
        public readonly Outputs.DisplayDeviceResponse DisplayDevice;
        /// <summary>
        /// Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
        /// </summary>
        public readonly bool EraseWindowsVssSignature;
        /// <summary>
        /// Specifies a fingerprint for this resource, which is essentially a hash of the instance's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update the instance. You must always provide an up-to-date fingerprint hash in order to update the instance.
        /// 
        /// To see the latest fingerprint, make get() request to the instance.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// A list of the type and count of accelerator cards attached to the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.AcceleratorConfigResponse> GuestAccelerators;
        /// <summary>
        /// Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
        /// </summary>
        public readonly string Hostname;
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
        public readonly Outputs.CustomerEncryptionKeyResponse InstanceEncryptionKey;
        /// <summary>
        /// Type of the resource. Always compute#instance for instances.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A fingerprint for this request, which is essentially a hash of the label's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels.
        /// 
        /// To see the latest fingerprint, make get() request to the instance.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// Labels to apply to this instance. These can be later modified by the setLabels method.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Last start timestamp in RFC3339 text format.
        /// </summary>
        public readonly string LastStartTimestamp;
        /// <summary>
        /// Last stop timestamp in RFC3339 text format.
        /// </summary>
        public readonly string LastStopTimestamp;
        /// <summary>
        /// Last suspended timestamp in RFC3339 text format.
        /// </summary>
        public readonly string LastSuspendedTimestamp;
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
        public readonly string MachineType;
        /// <summary>
        /// The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
        /// </summary>
        public readonly Outputs.MetadataResponse Metadata;
        /// <summary>
        /// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
        /// </summary>
        public readonly string MinCpuPlatform;
        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceResponse> NetworkInterfaces;
        public readonly Outputs.NetworkPerformanceConfigResponse NetworkPerformanceConfig;
        /// <summary>
        /// PostKeyRevocationActionType of the instance.
        /// </summary>
        public readonly string PostKeyRevocationActionType;
        /// <summary>
        /// Total amount of preserved state for SUSPENDED instances. Read-only in the api.
        /// </summary>
        public readonly string PreservedStateSizeGb;
        /// <summary>
        /// The private IPv6 google access type for the VM. If not specified, use  INHERIT_FROM_SUBNETWORK as default.
        /// </summary>
        public readonly string PrivateIpv6GoogleAccess;
        /// <summary>
        /// Specifies the reservations that this instance can consume from.
        /// </summary>
        public readonly Outputs.ReservationAffinityResponse ReservationAffinity;
        /// <summary>
        /// Resource policies applied to this instance.
        /// </summary>
        public readonly ImmutableArray<string> ResourcePolicies;
        /// <summary>
        /// Specifies values set for instance attributes as compared to the values requested by user in the corresponding input only field.
        /// </summary>
        public readonly Outputs.ResourceStatusResponse ResourceStatus;
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// Sets the scheduling options for this instance.
        /// </summary>
        public readonly Outputs.SchedulingResponse Scheduling;
        /// <summary>
        /// Secure labels to apply to this instance. These can be later modified by the update method. Maximum number of secure labels allowed is 300.
        /// </summary>
        public readonly ImmutableArray<string> SecureLabels;
        /// <summary>
        /// Secure tags to apply to this instance. These can be later modified by the update method. Maximum number of secure tags allowed is 300.
        /// </summary>
        public readonly ImmutableArray<string> SecureTags;
        /// <summary>
        /// Server-defined URL for this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported.
        /// 
        /// Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceAccountResponse> ServiceAccounts;
        public readonly Outputs.ShieldedInstanceConfigResponse ShieldedInstanceConfig;
        public readonly Outputs.ShieldedInstanceIntegrityPolicyResponse ShieldedInstanceIntegrityPolicy;
        /// <summary>
        /// Deprecating, please use shielded_instance_config.
        /// </summary>
        public readonly Outputs.ShieldedVmConfigResponse ShieldedVmConfig;
        /// <summary>
        /// Deprecating, please use shielded_instance_integrity_policy.
        /// </summary>
        public readonly Outputs.ShieldedVmIntegrityPolicyResponse ShieldedVmIntegrityPolicy;
        /// <summary>
        /// Source machine image
        /// </summary>
        public readonly string SourceMachineImage;
        /// <summary>
        /// Source machine image encryption key when creating an instance from a machine image.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse SourceMachineImageEncryptionKey;
        /// <summary>
        /// Whether a VM has been restricted for start because Compute Engine has detected suspicious activity.
        /// </summary>
        public readonly bool StartRestricted;
        /// <summary>
        /// The status of the instance. One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see  Instance life cycle.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// An optional, human-readable explanation of the status.
        /// </summary>
        public readonly string StatusMessage;
        /// <summary>
        /// Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
        /// </summary>
        public readonly Outputs.TagsResponse Tags;
        /// <summary>
        /// Specifies upcoming maintenance for the instance.
        /// </summary>
        public readonly Outputs.UpcomingMaintenanceResponse UpcomingMaintenance;
        /// <summary>
        /// URL of the zone where the instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetInstanceResult(
            Outputs.AdvancedMachineFeaturesResponse advancedMachineFeatures,

            bool canIpForward,

            Outputs.ConfidentialInstanceConfigResponse confidentialInstanceConfig,

            string cpuPlatform,

            string creationTimestamp,

            bool deletionProtection,

            string description,

            ImmutableArray<Outputs.AttachedDiskResponse> disks,

            Outputs.DisplayDeviceResponse displayDevice,

            bool eraseWindowsVssSignature,

            string fingerprint,

            ImmutableArray<Outputs.AcceleratorConfigResponse> guestAccelerators,

            string hostname,

            Outputs.CustomerEncryptionKeyResponse instanceEncryptionKey,

            string kind,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            string lastStartTimestamp,

            string lastStopTimestamp,

            string lastSuspendedTimestamp,

            string machineType,

            Outputs.MetadataResponse metadata,

            string minCpuPlatform,

            string name,

            ImmutableArray<Outputs.NetworkInterfaceResponse> networkInterfaces,

            Outputs.NetworkPerformanceConfigResponse networkPerformanceConfig,

            string postKeyRevocationActionType,

            string preservedStateSizeGb,

            string privateIpv6GoogleAccess,

            Outputs.ReservationAffinityResponse reservationAffinity,

            ImmutableArray<string> resourcePolicies,

            Outputs.ResourceStatusResponse resourceStatus,

            bool satisfiesPzs,

            Outputs.SchedulingResponse scheduling,

            ImmutableArray<string> secureLabels,

            ImmutableArray<string> secureTags,

            string selfLink,

            string selfLinkWithId,

            ImmutableArray<Outputs.ServiceAccountResponse> serviceAccounts,

            Outputs.ShieldedInstanceConfigResponse shieldedInstanceConfig,

            Outputs.ShieldedInstanceIntegrityPolicyResponse shieldedInstanceIntegrityPolicy,

            Outputs.ShieldedVmConfigResponse shieldedVmConfig,

            Outputs.ShieldedVmIntegrityPolicyResponse shieldedVmIntegrityPolicy,

            string sourceMachineImage,

            Outputs.CustomerEncryptionKeyResponse sourceMachineImageEncryptionKey,

            bool startRestricted,

            string status,

            string statusMessage,

            Outputs.TagsResponse tags,

            Outputs.UpcomingMaintenanceResponse upcomingMaintenance,

            string zone)
        {
            AdvancedMachineFeatures = advancedMachineFeatures;
            CanIpForward = canIpForward;
            ConfidentialInstanceConfig = confidentialInstanceConfig;
            CpuPlatform = cpuPlatform;
            CreationTimestamp = creationTimestamp;
            DeletionProtection = deletionProtection;
            Description = description;
            Disks = disks;
            DisplayDevice = displayDevice;
            EraseWindowsVssSignature = eraseWindowsVssSignature;
            Fingerprint = fingerprint;
            GuestAccelerators = guestAccelerators;
            Hostname = hostname;
            InstanceEncryptionKey = instanceEncryptionKey;
            Kind = kind;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            LastStartTimestamp = lastStartTimestamp;
            LastStopTimestamp = lastStopTimestamp;
            LastSuspendedTimestamp = lastSuspendedTimestamp;
            MachineType = machineType;
            Metadata = metadata;
            MinCpuPlatform = minCpuPlatform;
            Name = name;
            NetworkInterfaces = networkInterfaces;
            NetworkPerformanceConfig = networkPerformanceConfig;
            PostKeyRevocationActionType = postKeyRevocationActionType;
            PreservedStateSizeGb = preservedStateSizeGb;
            PrivateIpv6GoogleAccess = privateIpv6GoogleAccess;
            ReservationAffinity = reservationAffinity;
            ResourcePolicies = resourcePolicies;
            ResourceStatus = resourceStatus;
            SatisfiesPzs = satisfiesPzs;
            Scheduling = scheduling;
            SecureLabels = secureLabels;
            SecureTags = secureTags;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            ServiceAccounts = serviceAccounts;
            ShieldedInstanceConfig = shieldedInstanceConfig;
            ShieldedInstanceIntegrityPolicy = shieldedInstanceIntegrityPolicy;
            ShieldedVmConfig = shieldedVmConfig;
            ShieldedVmIntegrityPolicy = shieldedVmIntegrityPolicy;
            SourceMachineImage = sourceMachineImage;
            SourceMachineImageEncryptionKey = sourceMachineImageEncryptionKey;
            StartRestricted = startRestricted;
            Status = status;
            StatusMessage = statusMessage;
            Tags = tags;
            UpcomingMaintenance = upcomingMaintenance;
            Zone = zone;
        }
    }
}
