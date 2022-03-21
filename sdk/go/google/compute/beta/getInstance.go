// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified Instance resource. Gets a list of available instances by making a list() request.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("google-native:compute/beta:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	Instance string  `pulumi:"instance"`
	Project  *string `pulumi:"project"`
	Zone     string  `pulumi:"zone"`
}

type LookupInstanceResult struct {
	// Controls for advanced machine-related behavior features.
	AdvancedMachineFeatures AdvancedMachineFeaturesResponse `pulumi:"advancedMachineFeatures"`
	// Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding .
	CanIpForward               bool                               `pulumi:"canIpForward"`
	ConfidentialInstanceConfig ConfidentialInstanceConfigResponse `pulumi:"confidentialInstanceConfig"`
	// The CPU platform used by this instance.
	CpuPlatform string `pulumi:"cpuPlatform"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// Whether the resource should be protected against deletion.
	DeletionProtection bool `pulumi:"deletionProtection"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Array of disks associated with this instance. Persistent disks must be created before you can assign them.
	Disks []AttachedDiskResponse `pulumi:"disks"`
	// Enables display device for the instance.
	DisplayDevice DisplayDeviceResponse `pulumi:"displayDevice"`
	// Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
	EraseWindowsVssSignature bool `pulumi:"eraseWindowsVssSignature"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the instance's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update the instance. You must always provide an up-to-date fingerprint hash in order to update the instance. To see the latest fingerprint, make get() request to the instance.
	Fingerprint string `pulumi:"fingerprint"`
	// A list of the type and count of accelerator cards attached to the instance.
	GuestAccelerators []AcceleratorConfigResponse `pulumi:"guestAccelerators"`
	// Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
	Hostname string `pulumi:"hostname"`
	// Type of the resource. Always compute#instance for instances.
	Kind string `pulumi:"kind"`
	// A fingerprint for this request, which is essentially a hash of the label's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the instance.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels to apply to this instance. These can be later modified by the setLabels method.
	Labels map[string]string `pulumi:"labels"`
	// Last start timestamp in RFC3339 text format.
	LastStartTimestamp string `pulumi:"lastStartTimestamp"`
	// Last stop timestamp in RFC3339 text format.
	LastStopTimestamp string `pulumi:"lastStopTimestamp"`
	// Last suspended timestamp in RFC3339 text format.
	LastSuspendedTimestamp string `pulumi:"lastSuspendedTimestamp"`
	// Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type: zones/us-central1-f/machineTypes/n1-standard-1 To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB): zones/zone/machineTypes/custom-CPUS-MEMORY For example: zones/us-central1-f/machineTypes/custom-4-5120 For a full list of restrictions, read the Specifications for custom machine types.
	MachineType string `pulumi:"machineType"`
	// The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
	Metadata MetadataResponse `pulumi:"metadata"`
	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
	MinCpuPlatform string `pulumi:"minCpuPlatform"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
	NetworkInterfaces        []NetworkInterfaceResponse       `pulumi:"networkInterfaces"`
	NetworkPerformanceConfig NetworkPerformanceConfigResponse `pulumi:"networkPerformanceConfig"`
	// Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
	Params InstanceParamsResponse `pulumi:"params"`
	// PostKeyRevocationActionType of the instance.
	PostKeyRevocationActionType string `pulumi:"postKeyRevocationActionType"`
	// The private IPv6 google access type for the VM. If not specified, use INHERIT_FROM_SUBNETWORK as default.
	PrivateIpv6GoogleAccess string `pulumi:"privateIpv6GoogleAccess"`
	// Specifies the reservations that this instance can consume from.
	ReservationAffinity ReservationAffinityResponse `pulumi:"reservationAffinity"`
	// Resource policies applied to this instance.
	ResourcePolicies []string `pulumi:"resourcePolicies"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// Sets the scheduling options for this instance.
	Scheduling SchedulingResponse `pulumi:"scheduling"`
	// Server-defined URL for this resource.
	SelfLink string `pulumi:"selfLink"`
	// A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported. Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
	ServiceAccounts                 []ServiceAccountResponse                `pulumi:"serviceAccounts"`
	ShieldedInstanceConfig          ShieldedInstanceConfigResponse          `pulumi:"shieldedInstanceConfig"`
	ShieldedInstanceIntegrityPolicy ShieldedInstanceIntegrityPolicyResponse `pulumi:"shieldedInstanceIntegrityPolicy"`
	// Deprecating, please use shielded_instance_config.
	ShieldedVmConfig ShieldedVmConfigResponse `pulumi:"shieldedVmConfig"`
	// Deprecating, please use shielded_instance_integrity_policy.
	ShieldedVmIntegrityPolicy ShieldedVmIntegrityPolicyResponse `pulumi:"shieldedVmIntegrityPolicy"`
	// Source machine image
	SourceMachineImage string `pulumi:"sourceMachineImage"`
	// Source machine image encryption key when creating an instance from a machine image.
	SourceMachineImageEncryptionKey CustomerEncryptionKeyResponse `pulumi:"sourceMachineImageEncryptionKey"`
	// Whether a VM has been restricted for start because Compute Engine has detected suspicious activity.
	StartRestricted bool `pulumi:"startRestricted"`
	// The status of the instance. One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see Instance life cycle.
	Status string `pulumi:"status"`
	// An optional, human-readable explanation of the status.
	StatusMessage string `pulumi:"statusMessage"`
	// Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
	Tags TagsResponse `pulumi:"tags"`
	// URL of the zone where the instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Zone string `pulumi:"zone"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	Instance pulumi.StringInput    `pulumi:"instance"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	Zone     pulumi.StringInput    `pulumi:"zone"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

// Controls for advanced machine-related behavior features.
func (o LookupInstanceResultOutput) AdvancedMachineFeatures() AdvancedMachineFeaturesResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) AdvancedMachineFeaturesResponse { return v.AdvancedMachineFeatures }).(AdvancedMachineFeaturesResponseOutput)
}

// Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding .
func (o LookupInstanceResultOutput) CanIpForward() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.CanIpForward }).(pulumi.BoolOutput)
}

func (o LookupInstanceResultOutput) ConfidentialInstanceConfig() ConfidentialInstanceConfigResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) ConfidentialInstanceConfigResponse { return v.ConfidentialInstanceConfig }).(ConfidentialInstanceConfigResponseOutput)
}

// The CPU platform used by this instance.
func (o LookupInstanceResultOutput) CpuPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CpuPlatform }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupInstanceResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// Whether the resource should be protected against deletion.
func (o LookupInstanceResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupInstanceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Description }).(pulumi.StringOutput)
}

// Array of disks associated with this instance. Persistent disks must be created before you can assign them.
func (o LookupInstanceResultOutput) Disks() AttachedDiskResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []AttachedDiskResponse { return v.Disks }).(AttachedDiskResponseArrayOutput)
}

// Enables display device for the instance.
func (o LookupInstanceResultOutput) DisplayDevice() DisplayDeviceResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) DisplayDeviceResponse { return v.DisplayDevice }).(DisplayDeviceResponseOutput)
}

// Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
func (o LookupInstanceResultOutput) EraseWindowsVssSignature() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.EraseWindowsVssSignature }).(pulumi.BoolOutput)
}

// Specifies a fingerprint for this resource, which is essentially a hash of the instance's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update the instance. You must always provide an up-to-date fingerprint hash in order to update the instance. To see the latest fingerprint, make get() request to the instance.
func (o LookupInstanceResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// A list of the type and count of accelerator cards attached to the instance.
func (o LookupInstanceResultOutput) GuestAccelerators() AcceleratorConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []AcceleratorConfigResponse { return v.GuestAccelerators }).(AcceleratorConfigResponseArrayOutput)
}

// Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
func (o LookupInstanceResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#instance for instances.
func (o LookupInstanceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for this request, which is essentially a hash of the label's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the instance.
func (o LookupInstanceResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this instance. These can be later modified by the setLabels method.
func (o LookupInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Last start timestamp in RFC3339 text format.
func (o LookupInstanceResultOutput) LastStartTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.LastStartTimestamp }).(pulumi.StringOutput)
}

// Last stop timestamp in RFC3339 text format.
func (o LookupInstanceResultOutput) LastStopTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.LastStopTimestamp }).(pulumi.StringOutput)
}

// Last suspended timestamp in RFC3339 text format.
func (o LookupInstanceResultOutput) LastSuspendedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.LastSuspendedTimestamp }).(pulumi.StringOutput)
}

// Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type: zones/us-central1-f/machineTypes/n1-standard-1 To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB): zones/zone/machineTypes/custom-CPUS-MEMORY For example: zones/us-central1-f/machineTypes/custom-4-5120 For a full list of restrictions, read the Specifications for custom machine types.
func (o LookupInstanceResultOutput) MachineType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.MachineType }).(pulumi.StringOutput)
}

// The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
func (o LookupInstanceResultOutput) Metadata() MetadataResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) MetadataResponse { return v.Metadata }).(MetadataResponseOutput)
}

// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
func (o LookupInstanceResultOutput) MinCpuPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.MinCpuPlatform }).(pulumi.StringOutput)
}

// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
func (o LookupInstanceResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o LookupInstanceResultOutput) NetworkPerformanceConfig() NetworkPerformanceConfigResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) NetworkPerformanceConfigResponse { return v.NetworkPerformanceConfig }).(NetworkPerformanceConfigResponseOutput)
}

// Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
func (o LookupInstanceResultOutput) Params() InstanceParamsResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) InstanceParamsResponse { return v.Params }).(InstanceParamsResponseOutput)
}

// PostKeyRevocationActionType of the instance.
func (o LookupInstanceResultOutput) PostKeyRevocationActionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.PostKeyRevocationActionType }).(pulumi.StringOutput)
}

// The private IPv6 google access type for the VM. If not specified, use INHERIT_FROM_SUBNETWORK as default.
func (o LookupInstanceResultOutput) PrivateIpv6GoogleAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.PrivateIpv6GoogleAccess }).(pulumi.StringOutput)
}

// Specifies the reservations that this instance can consume from.
func (o LookupInstanceResultOutput) ReservationAffinity() ReservationAffinityResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) ReservationAffinityResponse { return v.ReservationAffinity }).(ReservationAffinityResponseOutput)
}

// Resource policies applied to this instance.
func (o LookupInstanceResultOutput) ResourcePolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.ResourcePolicies }).(pulumi.StringArrayOutput)
}

// Reserved for future use.
func (o LookupInstanceResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// Sets the scheduling options for this instance.
func (o LookupInstanceResultOutput) Scheduling() SchedulingResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) SchedulingResponse { return v.Scheduling }).(SchedulingResponseOutput)
}

// Server-defined URL for this resource.
func (o LookupInstanceResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported. Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
func (o LookupInstanceResultOutput) ServiceAccounts() ServiceAccountResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []ServiceAccountResponse { return v.ServiceAccounts }).(ServiceAccountResponseArrayOutput)
}

func (o LookupInstanceResultOutput) ShieldedInstanceConfig() ShieldedInstanceConfigResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) ShieldedInstanceConfigResponse { return v.ShieldedInstanceConfig }).(ShieldedInstanceConfigResponseOutput)
}

func (o LookupInstanceResultOutput) ShieldedInstanceIntegrityPolicy() ShieldedInstanceIntegrityPolicyResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) ShieldedInstanceIntegrityPolicyResponse {
		return v.ShieldedInstanceIntegrityPolicy
	}).(ShieldedInstanceIntegrityPolicyResponseOutput)
}

// Deprecating, please use shielded_instance_config.
func (o LookupInstanceResultOutput) ShieldedVmConfig() ShieldedVmConfigResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) ShieldedVmConfigResponse { return v.ShieldedVmConfig }).(ShieldedVmConfigResponseOutput)
}

// Deprecating, please use shielded_instance_integrity_policy.
func (o LookupInstanceResultOutput) ShieldedVmIntegrityPolicy() ShieldedVmIntegrityPolicyResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) ShieldedVmIntegrityPolicyResponse { return v.ShieldedVmIntegrityPolicy }).(ShieldedVmIntegrityPolicyResponseOutput)
}

// Source machine image
func (o LookupInstanceResultOutput) SourceMachineImage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.SourceMachineImage }).(pulumi.StringOutput)
}

// Source machine image encryption key when creating an instance from a machine image.
func (o LookupInstanceResultOutput) SourceMachineImageEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) CustomerEncryptionKeyResponse { return v.SourceMachineImageEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// Whether a VM has been restricted for start because Compute Engine has detected suspicious activity.
func (o LookupInstanceResultOutput) StartRestricted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.StartRestricted }).(pulumi.BoolOutput)
}

// The status of the instance. One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see Instance life cycle.
func (o LookupInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

// An optional, human-readable explanation of the status.
func (o LookupInstanceResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

// Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
func (o LookupInstanceResultOutput) Tags() TagsResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) TagsResponse { return v.Tags }).(TagsResponseOutput)
}

// URL of the zone where the instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupInstanceResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
