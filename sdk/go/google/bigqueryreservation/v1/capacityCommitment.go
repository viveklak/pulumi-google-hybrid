// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new capacity commitment resource.
type CapacityCommitment struct {
	pulumi.CustomResourceState

	// The end of the current commitment period. It is applicable only for ACTIVE capacity commitments.
	CommitmentEndTime pulumi.StringOutput `pulumi:"commitmentEndTime"`
	// The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.
	CommitmentStartTime pulumi.StringOutput `pulumi:"commitmentStartTime"`
	// For FAILED commitment plan, provides the reason of failure.
	FailureStatus StatusResponseOutput `pulumi:"failureStatus"`
	// The resource name of the capacity commitment, e.g., `projects/myproject/locations/US/capacityCommitments/123`
	Name pulumi.StringOutput `pulumi:"name"`
	// Capacity commitment commitment plan.
	Plan pulumi.StringOutput `pulumi:"plan"`
	// The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
	RenewalPlan pulumi.StringOutput `pulumi:"renewalPlan"`
	// Number of slots in this commitment.
	SlotCount pulumi.StringOutput `pulumi:"slotCount"`
	// State of the commitment.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewCapacityCommitment registers a new resource with the given unique name, arguments, and options.
func NewCapacityCommitment(ctx *pulumi.Context,
	name string, args *CapacityCommitmentArgs, opts ...pulumi.ResourceOption) (*CapacityCommitment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource CapacityCommitment
	err := ctx.RegisterResource("google-native:bigqueryreservation/v1:CapacityCommitment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityCommitment gets an existing CapacityCommitment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityCommitment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityCommitmentState, opts ...pulumi.ResourceOption) (*CapacityCommitment, error) {
	var resource CapacityCommitment
	err := ctx.ReadResource("google-native:bigqueryreservation/v1:CapacityCommitment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityCommitment resources.
type capacityCommitmentState struct {
}

type CapacityCommitmentState struct {
}

func (CapacityCommitmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityCommitmentState)(nil)).Elem()
}

type capacityCommitmentArgs struct {
	CapacityCommitmentId            *string `pulumi:"capacityCommitmentId"`
	EnforceSingleAdminProjectPerOrg *string `pulumi:"enforceSingleAdminProjectPerOrg"`
	Location                        string  `pulumi:"location"`
	// Capacity commitment commitment plan.
	Plan    *string `pulumi:"plan"`
	Project string  `pulumi:"project"`
	// The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
	RenewalPlan *string `pulumi:"renewalPlan"`
	// Number of slots in this commitment.
	SlotCount *string `pulumi:"slotCount"`
}

// The set of arguments for constructing a CapacityCommitment resource.
type CapacityCommitmentArgs struct {
	CapacityCommitmentId            pulumi.StringPtrInput
	EnforceSingleAdminProjectPerOrg pulumi.StringPtrInput
	Location                        pulumi.StringInput
	// Capacity commitment commitment plan.
	Plan    *CapacityCommitmentPlan
	Project pulumi.StringInput
	// The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
	RenewalPlan *CapacityCommitmentRenewalPlan
	// Number of slots in this commitment.
	SlotCount pulumi.StringPtrInput
}

func (CapacityCommitmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityCommitmentArgs)(nil)).Elem()
}

type CapacityCommitmentInput interface {
	pulumi.Input

	ToCapacityCommitmentOutput() CapacityCommitmentOutput
	ToCapacityCommitmentOutputWithContext(ctx context.Context) CapacityCommitmentOutput
}

func (*CapacityCommitment) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityCommitment)(nil))
}

func (i *CapacityCommitment) ToCapacityCommitmentOutput() CapacityCommitmentOutput {
	return i.ToCapacityCommitmentOutputWithContext(context.Background())
}

func (i *CapacityCommitment) ToCapacityCommitmentOutputWithContext(ctx context.Context) CapacityCommitmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityCommitmentOutput)
}

type CapacityCommitmentOutput struct {
	*pulumi.OutputState
}

func (CapacityCommitmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityCommitment)(nil))
}

func (o CapacityCommitmentOutput) ToCapacityCommitmentOutput() CapacityCommitmentOutput {
	return o
}

func (o CapacityCommitmentOutput) ToCapacityCommitmentOutputWithContext(ctx context.Context) CapacityCommitmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CapacityCommitmentOutput{})
}
