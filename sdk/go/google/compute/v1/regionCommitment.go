// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a commitment in the specified project using the data included in the request.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type RegionCommitment struct {
	pulumi.CustomResourceState

	// Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
	AutoRenew pulumi.BoolOutput `pulumi:"autoRenew"`
	// The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
	Category pulumi.StringOutput `pulumi:"category"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Commitment end time in RFC3339 text format.
	EndTimestamp pulumi.StringOutput `pulumi:"endTimestamp"`
	// Type of the resource. Always compute#commitment for commitments.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The license specification required as part of a license commitment.
	LicenseResource LicenseResourceCommitmentResponseOutput `pulumi:"licenseResource"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
	Plan pulumi.StringOutput `pulumi:"plan"`
	// URL of the region where this commitment may be used.
	Region pulumi.StringOutput `pulumi:"region"`
	// List of reservations in this commitment.
	Reservations ReservationResponseArrayOutput `pulumi:"reservations"`
	// A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
	Resources ResourceCommitmentResponseArrayOutput `pulumi:"resources"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Commitment start time in RFC3339 text format.
	StartTimestamp pulumi.StringOutput `pulumi:"startTimestamp"`
	// Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
	Status pulumi.StringOutput `pulumi:"status"`
	// An optional, human-readable explanation of the status.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegionCommitment registers a new resource with the given unique name, arguments, and options.
func NewRegionCommitment(ctx *pulumi.Context,
	name string, args *RegionCommitmentArgs, opts ...pulumi.ResourceOption) (*RegionCommitment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionCommitment
	err := ctx.RegisterResource("google-native:compute/v1:RegionCommitment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionCommitment gets an existing RegionCommitment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionCommitment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionCommitmentState, opts ...pulumi.ResourceOption) (*RegionCommitment, error) {
	var resource RegionCommitment
	err := ctx.ReadResource("google-native:compute/v1:RegionCommitment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionCommitment resources.
type regionCommitmentState struct {
}

type RegionCommitmentState struct {
}

func (RegionCommitmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionCommitmentState)(nil)).Elem()
}

type regionCommitmentArgs struct {
	// Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
	AutoRenew *bool `pulumi:"autoRenew"`
	// The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
	Category *RegionCommitmentCategory `pulumi:"category"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The license specification required as part of a license commitment.
	LicenseResource *LicenseResourceCommitment `pulumi:"licenseResource"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
	Plan    *RegionCommitmentPlan `pulumi:"plan"`
	Project *string               `pulumi:"project"`
	Region  string                `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// List of reservations in this commitment.
	Reservations []ReservationType `pulumi:"reservations"`
	// A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
	Resources []ResourceCommitment `pulumi:"resources"`
	// The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
	Type *RegionCommitmentType `pulumi:"type"`
}

// The set of arguments for constructing a RegionCommitment resource.
type RegionCommitmentArgs struct {
	// Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
	AutoRenew pulumi.BoolPtrInput
	// The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
	Category RegionCommitmentCategoryPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The license specification required as part of a license commitment.
	LicenseResource LicenseResourceCommitmentPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
	Plan    RegionCommitmentPlanPtrInput
	Project pulumi.StringPtrInput
	Region  pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// List of reservations in this commitment.
	Reservations ReservationTypeArrayInput
	// A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
	Resources ResourceCommitmentArrayInput
	// The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
	Type RegionCommitmentTypePtrInput
}

func (RegionCommitmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionCommitmentArgs)(nil)).Elem()
}

type RegionCommitmentInput interface {
	pulumi.Input

	ToRegionCommitmentOutput() RegionCommitmentOutput
	ToRegionCommitmentOutputWithContext(ctx context.Context) RegionCommitmentOutput
}

func (*RegionCommitment) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionCommitment)(nil)).Elem()
}

func (i *RegionCommitment) ToRegionCommitmentOutput() RegionCommitmentOutput {
	return i.ToRegionCommitmentOutputWithContext(context.Background())
}

func (i *RegionCommitment) ToRegionCommitmentOutputWithContext(ctx context.Context) RegionCommitmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionCommitmentOutput)
}

type RegionCommitmentOutput struct{ *pulumi.OutputState }

func (RegionCommitmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionCommitment)(nil)).Elem()
}

func (o RegionCommitmentOutput) ToRegionCommitmentOutput() RegionCommitmentOutput {
	return o
}

func (o RegionCommitmentOutput) ToRegionCommitmentOutputWithContext(ctx context.Context) RegionCommitmentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionCommitmentInput)(nil)).Elem(), &RegionCommitment{})
	pulumi.RegisterOutputType(RegionCommitmentOutput{})
}
