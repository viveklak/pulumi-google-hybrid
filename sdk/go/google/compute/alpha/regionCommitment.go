// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a commitment in the specified project using the data included in the request.
type RegionCommitment struct {
	pulumi.CustomResourceState
}

// NewRegionCommitment registers a new resource with the given unique name, arguments, and options.
func NewRegionCommitment(ctx *pulumi.Context,
	name string, args *RegionCommitmentArgs, opts ...pulumi.ResourceOption) (*RegionCommitment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionCommitment
	err := ctx.RegisterResource("google-cloud:compute/alpha:RegionCommitment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-cloud:compute/alpha:RegionCommitment", name, id, state, &resource, opts...)
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
	// The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
	Category *string `pulumi:"category"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// [Output Only] Commitment end time in RFC3339 text format.
	EndTimestamp *string `pulumi:"endTimestamp"`
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id *string `pulumi:"id"`
	// [Output Only] Type of the resource. Always compute#commitment for commitments.
	Kind *string `pulumi:"kind"`
	// The license specification required as part of a license commitment.
	LicenseResource *LicenseResourceCommitment `pulumi:"licenseResource"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
	Plan    *string `pulumi:"plan"`
	Project string  `pulumi:"project"`
	// [Output Only] URL of the region where this commitment may be used.
	Region string `pulumi:"region"`
	// List of reservations in this commitment.
	Reservations []ReservationType `pulumi:"reservations"`
	// A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
	Resources []ResourceCommitment `pulumi:"resources"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// [Output Only] Commitment start time in RFC3339 text format.
	StartTimestamp *string `pulumi:"startTimestamp"`
	// [Output Only] Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
	Status *string `pulumi:"status"`
	// [Output Only] An optional, human-readable explanation of the status.
	StatusMessage *string `pulumi:"statusMessage"`
	// The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a RegionCommitment resource.
type RegionCommitmentArgs struct {
	// The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
	Category pulumi.StringPtrInput
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// [Output Only] Commitment end time in RFC3339 text format.
	EndTimestamp pulumi.StringPtrInput
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id pulumi.StringPtrInput
	// [Output Only] Type of the resource. Always compute#commitment for commitments.
	Kind pulumi.StringPtrInput
	// The license specification required as part of a license commitment.
	LicenseResource LicenseResourceCommitmentPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
	Plan    pulumi.StringPtrInput
	Project pulumi.StringInput
	// [Output Only] URL of the region where this commitment may be used.
	Region pulumi.StringInput
	// List of reservations in this commitment.
	Reservations ReservationTypeArrayInput
	// A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
	Resources ResourceCommitmentArrayInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	// [Output Only] Commitment start time in RFC3339 text format.
	StartTimestamp pulumi.StringPtrInput
	// [Output Only] Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
	Status pulumi.StringPtrInput
	// [Output Only] An optional, human-readable explanation of the status.
	StatusMessage pulumi.StringPtrInput
	// The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
	Type pulumi.StringPtrInput
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
	return reflect.TypeOf((*RegionCommitment)(nil))
}

func (i *RegionCommitment) ToRegionCommitmentOutput() RegionCommitmentOutput {
	return i.ToRegionCommitmentOutputWithContext(context.Background())
}

func (i *RegionCommitment) ToRegionCommitmentOutputWithContext(ctx context.Context) RegionCommitmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionCommitmentOutput)
}

type RegionCommitmentOutput struct {
	*pulumi.OutputState
}

func (RegionCommitmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionCommitment)(nil))
}

func (o RegionCommitmentOutput) ToRegionCommitmentOutput() RegionCommitmentOutput {
	return o
}

func (o RegionCommitmentOutput) ToRegionCommitmentOutputWithContext(ctx context.Context) RegionCommitmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionCommitmentOutput{})
}
