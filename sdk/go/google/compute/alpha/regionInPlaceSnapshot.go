// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an in-place snapshot in the specified region.
type RegionInPlaceSnapshot struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Size of the source disk, specified in GB.
	DiskSizeGb pulumi.StringOutput `pulumi:"diskSizeGb"`
	// Specifies to create an application consistent in-place snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolOutput `pulumi:"guestFlush"`
	// Type of the resource. Always compute#inPlaceSnapshot for InPlaceSnapshot resources.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this InPlaceSnapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a InPlaceSnapshot.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this InPlaceSnapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the region where the in-place snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource's resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// URL of the source disk used to create this in-place snapshot. Note that the source disk must be in the same zone/region as the in-place snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
	// - projects/project/zones/zone/disks/disk
	// - zones/zone/disks/disk
	SourceDisk pulumi.StringOutput `pulumi:"sourceDisk"`
	// The ID value of the disk used to create this InPlaceSnapshot. This value may be used to determine whether the InPlaceSnapshot was taken from the current or a previous instance of a given disk name.
	SourceDiskId pulumi.StringOutput `pulumi:"sourceDiskId"`
	// The status of the inPlaceSnapshot. This can be CREATING, DELETING, FAILED, or READY.
	Status pulumi.StringOutput `pulumi:"status"`
	// URL of the zone where the in-place snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewRegionInPlaceSnapshot registers a new resource with the given unique name, arguments, and options.
func NewRegionInPlaceSnapshot(ctx *pulumi.Context,
	name string, args *RegionInPlaceSnapshotArgs, opts ...pulumi.ResourceOption) (*RegionInPlaceSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionInPlaceSnapshot
	err := ctx.RegisterResource("google-native:compute/alpha:RegionInPlaceSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionInPlaceSnapshot gets an existing RegionInPlaceSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionInPlaceSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionInPlaceSnapshotState, opts ...pulumi.ResourceOption) (*RegionInPlaceSnapshot, error) {
	var resource RegionInPlaceSnapshot
	err := ctx.ReadResource("google-native:compute/alpha:RegionInPlaceSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionInPlaceSnapshot resources.
type regionInPlaceSnapshotState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Size of the source disk, specified in GB.
	DiskSizeGb *string `pulumi:"diskSizeGb"`
	// Specifies to create an application consistent in-place snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush *bool `pulumi:"guestFlush"`
	// Type of the resource. Always compute#inPlaceSnapshot for InPlaceSnapshot resources.
	Kind *string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this InPlaceSnapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a InPlaceSnapshot.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this InPlaceSnapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// URL of the region where the in-place snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region *string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// Server-defined URL for this resource's resource id.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// URL of the source disk used to create this in-place snapshot. Note that the source disk must be in the same zone/region as the in-place snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
	// - projects/project/zones/zone/disks/disk
	// - zones/zone/disks/disk
	SourceDisk *string `pulumi:"sourceDisk"`
	// The ID value of the disk used to create this InPlaceSnapshot. This value may be used to determine whether the InPlaceSnapshot was taken from the current or a previous instance of a given disk name.
	SourceDiskId *string `pulumi:"sourceDiskId"`
	// The status of the inPlaceSnapshot. This can be CREATING, DELETING, FAILED, or READY.
	Status *string `pulumi:"status"`
	// URL of the zone where the in-place snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Zone *string `pulumi:"zone"`
}

type RegionInPlaceSnapshotState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Size of the source disk, specified in GB.
	DiskSizeGb pulumi.StringPtrInput
	// Specifies to create an application consistent in-place snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolPtrInput
	// Type of the resource. Always compute#inPlaceSnapshot for InPlaceSnapshot resources.
	Kind pulumi.StringPtrInput
	// A fingerprint for the labels being applied to this InPlaceSnapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a InPlaceSnapshot.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this InPlaceSnapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// URL of the region where the in-place snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringPtrInput
	// Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// Server-defined URL for this resource's resource id.
	SelfLinkWithId pulumi.StringPtrInput
	// URL of the source disk used to create this in-place snapshot. Note that the source disk must be in the same zone/region as the in-place snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
	// - projects/project/zones/zone/disks/disk
	// - zones/zone/disks/disk
	SourceDisk pulumi.StringPtrInput
	// The ID value of the disk used to create this InPlaceSnapshot. This value may be used to determine whether the InPlaceSnapshot was taken from the current or a previous instance of a given disk name.
	SourceDiskId pulumi.StringPtrInput
	// The status of the inPlaceSnapshot. This can be CREATING, DELETING, FAILED, or READY.
	Status pulumi.StringPtrInput
	// URL of the zone where the in-place snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Zone pulumi.StringPtrInput
}

func (RegionInPlaceSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionInPlaceSnapshotState)(nil)).Elem()
}

type regionInPlaceSnapshotArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Specifies to create an application consistent in-place snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush *bool `pulumi:"guestFlush"`
	// Labels to apply to this InPlaceSnapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      *string `pulumi:"name"`
	Project   string  `pulumi:"project"`
	Region    string  `pulumi:"region"`
	RequestId *string `pulumi:"requestId"`
	// URL of the source disk used to create this in-place snapshot. Note that the source disk must be in the same zone/region as the in-place snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
	// - projects/project/zones/zone/disks/disk
	// - zones/zone/disks/disk
	SourceDisk *string `pulumi:"sourceDisk"`
}

// The set of arguments for constructing a RegionInPlaceSnapshot resource.
type RegionInPlaceSnapshotArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Specifies to create an application consistent in-place snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolPtrInput
	// Labels to apply to this InPlaceSnapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringInput
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// URL of the source disk used to create this in-place snapshot. Note that the source disk must be in the same zone/region as the in-place snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk
	// - projects/project/zones/zone/disks/disk
	// - zones/zone/disks/disk
	SourceDisk pulumi.StringPtrInput
}

func (RegionInPlaceSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionInPlaceSnapshotArgs)(nil)).Elem()
}

type RegionInPlaceSnapshotInput interface {
	pulumi.Input

	ToRegionInPlaceSnapshotOutput() RegionInPlaceSnapshotOutput
	ToRegionInPlaceSnapshotOutputWithContext(ctx context.Context) RegionInPlaceSnapshotOutput
}

func (*RegionInPlaceSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionInPlaceSnapshot)(nil))
}

func (i *RegionInPlaceSnapshot) ToRegionInPlaceSnapshotOutput() RegionInPlaceSnapshotOutput {
	return i.ToRegionInPlaceSnapshotOutputWithContext(context.Background())
}

func (i *RegionInPlaceSnapshot) ToRegionInPlaceSnapshotOutputWithContext(ctx context.Context) RegionInPlaceSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionInPlaceSnapshotOutput)
}

type RegionInPlaceSnapshotOutput struct {
	*pulumi.OutputState
}

func (RegionInPlaceSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionInPlaceSnapshot)(nil))
}

func (o RegionInPlaceSnapshotOutput) ToRegionInPlaceSnapshotOutput() RegionInPlaceSnapshotOutput {
	return o
}

func (o RegionInPlaceSnapshotOutput) ToRegionInPlaceSnapshotOutputWithContext(ctx context.Context) RegionInPlaceSnapshotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionInPlaceSnapshotOutput{})
}
