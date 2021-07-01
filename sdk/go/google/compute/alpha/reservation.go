// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new reservation. For more information, read Reserving zonal resources.
type Reservation struct {
	pulumi.CustomResourceState

	// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
	Commitment pulumi.StringOutput `pulumi:"commitment"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of the resource. Always compute#reservations for reservations.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// Share-settings for shared-reservation
	ShareSettings AllocationShareSettingsResponseOutput `pulumi:"shareSettings"`
	// Reservation for instances with specific machine shapes.
	SpecificReservation AllocationSpecificSKUReservationResponseOutput `pulumi:"specificReservation"`
	// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
	SpecificReservationRequired pulumi.BoolOutput `pulumi:"specificReservationRequired"`
	// The status of the reservation.
	Status pulumi.StringOutput `pulumi:"status"`
	// Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewReservation registers a new resource with the given unique name, arguments, and options.
func NewReservation(ctx *pulumi.Context,
	name string, args *ReservationArgs, opts ...pulumi.ResourceOption) (*Reservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource Reservation
	err := ctx.RegisterResource("google-native:compute/alpha:Reservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReservation gets an existing Reservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReservationState, opts ...pulumi.ResourceOption) (*Reservation, error) {
	var resource Reservation
	err := ctx.ReadResource("google-native:compute/alpha:Reservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Reservation resources.
type reservationState struct {
}

type ReservationState struct {
}

func (ReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationState)(nil)).Elem()
}

type reservationArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      *string `pulumi:"name"`
	Project   string  `pulumi:"project"`
	RequestId *string `pulumi:"requestId"`
	// Share-settings for shared-reservation
	ShareSettings *AllocationShareSettings `pulumi:"shareSettings"`
	// Reservation for instances with specific machine shapes.
	SpecificReservation *AllocationSpecificSKUReservation `pulumi:"specificReservation"`
	// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
	SpecificReservationRequired *bool `pulumi:"specificReservationRequired"`
	// Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Reservation resource.
type ReservationArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// Share-settings for shared-reservation
	ShareSettings AllocationShareSettingsPtrInput
	// Reservation for instances with specific machine shapes.
	SpecificReservation AllocationSpecificSKUReservationPtrInput
	// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
	SpecificReservationRequired pulumi.BoolPtrInput
	// Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
	Zone pulumi.StringInput
}

func (ReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationArgs)(nil)).Elem()
}

type ReservationInput interface {
	pulumi.Input

	ToReservationOutput() ReservationOutput
	ToReservationOutputWithContext(ctx context.Context) ReservationOutput
}

func (*Reservation) ElementType() reflect.Type {
	return reflect.TypeOf((*Reservation)(nil))
}

func (i *Reservation) ToReservationOutput() ReservationOutput {
	return i.ToReservationOutputWithContext(context.Background())
}

func (i *Reservation) ToReservationOutputWithContext(ctx context.Context) ReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationOutput)
}

type ReservationOutput struct {
	*pulumi.OutputState
}

func (ReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Reservation)(nil))
}

func (o ReservationOutput) ToReservationOutput() ReservationOutput {
	return o
}

func (o ReservationOutput) ToReservationOutputWithContext(ctx context.Context) ReservationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReservationOutput{})
}
