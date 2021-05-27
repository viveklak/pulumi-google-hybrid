// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a License resource in the specified project.  Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
type License struct {
	pulumi.CustomResourceState

	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringOutput `pulumi:"description"`
	// [Output Only] Type of resource. Always compute#license for licenses.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// [Output Only] The unique code used to attach this license to images, snapshots, and disks.
	LicenseCode pulumi.StringOutput `pulumi:"licenseCode"`
	// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
	Name                 pulumi.StringOutput                       `pulumi:"name"`
	ResourceRequirements LicenseResourceRequirementsResponseOutput `pulumi:"resourceRequirements"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
	Transferable pulumi.BoolOutput `pulumi:"transferable"`
}

// NewLicense registers a new resource with the given unique name, arguments, and options.
func NewLicense(ctx *pulumi.Context,
	name string, args *LicenseArgs, opts ...pulumi.ResourceOption) (*License, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource License
	err := ctx.RegisterResource("google-native:compute/beta:License", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLicense gets an existing License resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLicense(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LicenseState, opts ...pulumi.ResourceOption) (*License, error) {
	var resource License
	err := ctx.ReadResource("google-native:compute/beta:License", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering License resources.
type licenseState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description *string `pulumi:"description"`
	// [Output Only] Type of resource. Always compute#license for licenses.
	Kind *string `pulumi:"kind"`
	// [Output Only] The unique code used to attach this license to images, snapshots, and disks.
	LicenseCode *string `pulumi:"licenseCode"`
	// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
	Name                 *string                              `pulumi:"name"`
	ResourceRequirements *LicenseResourceRequirementsResponse `pulumi:"resourceRequirements"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
	Transferable *bool `pulumi:"transferable"`
}

type LicenseState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringPtrInput
	// [Output Only] Type of resource. Always compute#license for licenses.
	Kind pulumi.StringPtrInput
	// [Output Only] The unique code used to attach this license to images, snapshots, and disks.
	LicenseCode pulumi.StringPtrInput
	// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
	Name                 pulumi.StringPtrInput
	ResourceRequirements LicenseResourceRequirementsResponsePtrInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
	Transferable pulumi.BoolPtrInput
}

func (LicenseState) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseState)(nil)).Elem()
}

type licenseArgs struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description *string `pulumi:"description"`
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id *string `pulumi:"id"`
	// [Output Only] Type of resource. Always compute#license for licenses.
	Kind *string `pulumi:"kind"`
	// [Output Only] The unique code used to attach this license to images, snapshots, and disks.
	LicenseCode *string `pulumi:"licenseCode"`
	// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
	Name                 *string                      `pulumi:"name"`
	Project              string                       `pulumi:"project"`
	RequestId            *string                      `pulumi:"requestId"`
	ResourceRequirements *LicenseResourceRequirements `pulumi:"resourceRequirements"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
	Transferable *bool `pulumi:"transferable"`
}

// The set of arguments for constructing a License resource.
type LicenseArgs struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringPtrInput
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id pulumi.StringPtrInput
	// [Output Only] Type of resource. Always compute#license for licenses.
	Kind pulumi.StringPtrInput
	// [Output Only] The unique code used to attach this license to images, snapshots, and disks.
	LicenseCode pulumi.StringPtrInput
	// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
	Name                 pulumi.StringPtrInput
	Project              pulumi.StringInput
	RequestId            pulumi.StringPtrInput
	ResourceRequirements LicenseResourceRequirementsPtrInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
	Transferable pulumi.BoolPtrInput
}

func (LicenseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseArgs)(nil)).Elem()
}

type LicenseInput interface {
	pulumi.Input

	ToLicenseOutput() LicenseOutput
	ToLicenseOutputWithContext(ctx context.Context) LicenseOutput
}

func (*License) ElementType() reflect.Type {
	return reflect.TypeOf((*License)(nil))
}

func (i *License) ToLicenseOutput() LicenseOutput {
	return i.ToLicenseOutputWithContext(context.Background())
}

func (i *License) ToLicenseOutputWithContext(ctx context.Context) LicenseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseOutput)
}

type LicenseOutput struct {
	*pulumi.OutputState
}

func (LicenseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*License)(nil))
}

func (o LicenseOutput) ToLicenseOutput() LicenseOutput {
	return o
}

func (o LicenseOutput) ToLicenseOutputWithContext(ctx context.Context) LicenseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LicenseOutput{})
}
