// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an instant snapshot in the specified zone.
type InstantSnapshot struct {
	pulumi.CustomResourceState

	// The architecture of the instant snapshot. Valid values are ARM64 or X86_64.
	Architecture pulumi.StringOutput `pulumi:"architecture"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Size of the source disk, specified in GB.
	DiskSizeGb pulumi.StringOutput `pulumi:"diskSizeGb"`
	// Whether to attempt an application consistent instant snapshot by informing the OS to prepare for the snapshot process.
	GuestFlush pulumi.BoolOutput `pulumi:"guestFlush"`
	// Type of the resource. Always compute#instantSnapshot for InstantSnapshot resources.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this InstantSnapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a InstantSnapshot.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this InstantSnapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the region where the instant snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// Status information for the instant snapshot resource.
	ResourceStatus InstantSnapshotResourceStatusResponseOutput `pulumi:"resourceStatus"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource's resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// URL of the source disk used to create this instant snapshot. Note that the source disk must be in the same zone/region as the instant snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk
	SourceDisk pulumi.StringOutput `pulumi:"sourceDisk"`
	// The ID value of the disk used to create this InstantSnapshot. This value may be used to determine whether the InstantSnapshot was taken from the current or a previous instance of a given disk name.
	SourceDiskId pulumi.StringOutput `pulumi:"sourceDiskId"`
	// The status of the instantSnapshot. This can be CREATING, DELETING, FAILED, or READY.
	Status pulumi.StringOutput `pulumi:"status"`
	// URL of the zone where the instant snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstantSnapshot registers a new resource with the given unique name, arguments, and options.
func NewInstantSnapshot(ctx *pulumi.Context,
	name string, args *InstantSnapshotArgs, opts ...pulumi.ResourceOption) (*InstantSnapshot, error) {
	if args == nil {
		args = &InstantSnapshotArgs{}
	}

	var resource InstantSnapshot
	err := ctx.RegisterResource("google-native:compute/alpha:InstantSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstantSnapshot gets an existing InstantSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstantSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstantSnapshotState, opts ...pulumi.ResourceOption) (*InstantSnapshot, error) {
	var resource InstantSnapshot
	err := ctx.ReadResource("google-native:compute/alpha:InstantSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstantSnapshot resources.
type instantSnapshotState struct {
}

type InstantSnapshotState struct {
}

func (InstantSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*instantSnapshotState)(nil)).Elem()
}

type instantSnapshotArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Whether to attempt an application consistent instant snapshot by informing the OS to prepare for the snapshot process.
	GuestFlush *bool `pulumi:"guestFlush"`
	// Labels to apply to this InstantSnapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// URL of the source disk used to create this instant snapshot. Note that the source disk must be in the same zone/region as the instant snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk
	SourceDisk *string `pulumi:"sourceDisk"`
	Zone       *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstantSnapshot resource.
type InstantSnapshotArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Whether to attempt an application consistent instant snapshot by informing the OS to prepare for the snapshot process.
	GuestFlush pulumi.BoolPtrInput
	// Labels to apply to this InstantSnapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// URL of the source disk used to create this instant snapshot. Note that the source disk must be in the same zone/region as the instant snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk
	SourceDisk pulumi.StringPtrInput
	Zone       pulumi.StringPtrInput
}

func (InstantSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instantSnapshotArgs)(nil)).Elem()
}

type InstantSnapshotInput interface {
	pulumi.Input

	ToInstantSnapshotOutput() InstantSnapshotOutput
	ToInstantSnapshotOutputWithContext(ctx context.Context) InstantSnapshotOutput
}

func (*InstantSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**InstantSnapshot)(nil)).Elem()
}

func (i *InstantSnapshot) ToInstantSnapshotOutput() InstantSnapshotOutput {
	return i.ToInstantSnapshotOutputWithContext(context.Background())
}

func (i *InstantSnapshot) ToInstantSnapshotOutputWithContext(ctx context.Context) InstantSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstantSnapshotOutput)
}

type InstantSnapshotOutput struct{ *pulumi.OutputState }

func (InstantSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstantSnapshot)(nil)).Elem()
}

func (o InstantSnapshotOutput) ToInstantSnapshotOutput() InstantSnapshotOutput {
	return o
}

func (o InstantSnapshotOutput) ToInstantSnapshotOutputWithContext(ctx context.Context) InstantSnapshotOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstantSnapshotInput)(nil)).Elem(), &InstantSnapshot{})
	pulumi.RegisterOutputType(InstantSnapshotOutput{})
}
