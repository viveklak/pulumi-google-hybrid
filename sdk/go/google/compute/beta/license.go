// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a License resource in the specified project. *Caution* This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
type License struct {
	pulumi.CustomResourceState

	// Deprecated. This field no longer reflects whether a license charges a usage fee.
	//
	// Deprecated: [Output Only] Deprecated. This field no longer reflects whether a license charges a usage fee.
	ChargesUseFee pulumi.BoolOutput `pulumi:"chargesUseFee"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of resource. Always compute#license for licenses.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The unique code used to attach this license to images, snapshots, and disks.
	LicenseCode pulumi.StringOutput `pulumi:"licenseCode"`
	// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
	Name                 pulumi.StringOutput                       `pulumi:"name"`
	ResourceRequirements LicenseResourceRequirementsResponseOutput `pulumi:"resourceRequirements"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
	Transferable pulumi.BoolOutput `pulumi:"transferable"`
}

// NewLicense registers a new resource with the given unique name, arguments, and options.
func NewLicense(ctx *pulumi.Context,
	name string, args *LicenseArgs, opts ...pulumi.ResourceOption) (*License, error) {
	if args == nil {
		args = &LicenseArgs{}
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
}

type LicenseState struct {
}

func (LicenseState) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseState)(nil)).Elem()
}

type licenseArgs struct {
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description *string `pulumi:"description"`
	// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId            *string                      `pulumi:"requestId"`
	ResourceRequirements *LicenseResourceRequirements `pulumi:"resourceRequirements"`
	// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
	Transferable *bool `pulumi:"transferable"`
}

// The set of arguments for constructing a License resource.
type LicenseArgs struct {
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringPtrInput
	// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId            pulumi.StringPtrInput
	ResourceRequirements LicenseResourceRequirementsPtrInput
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
	return reflect.TypeOf((**License)(nil)).Elem()
}

func (i *License) ToLicenseOutput() LicenseOutput {
	return i.ToLicenseOutputWithContext(context.Background())
}

func (i *License) ToLicenseOutputWithContext(ctx context.Context) LicenseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseOutput)
}

type LicenseOutput struct{ *pulumi.OutputState }

func (LicenseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**License)(nil)).Elem()
}

func (o LicenseOutput) ToLicenseOutput() LicenseOutput {
	return o
}

func (o LicenseOutput) ToLicenseOutputWithContext(ctx context.Context) LicenseOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseInput)(nil)).Elem(), &License{})
	pulumi.RegisterOutputType(LicenseOutput{})
}
