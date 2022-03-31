// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new reservation resource.
type Reservation struct {
	pulumi.CustomResourceState

	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
	Concurrency pulumi.StringOutput `pulumi:"concurrency"`
	// Creation time of the reservation.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
	IgnoreIdleSlots pulumi.BoolOutput `pulumi:"ignoreIdleSlots"`
	// Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
	MultiRegionAuxiliary pulumi.BoolOutput `pulumi:"multiRegionAuxiliary"`
	// The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. For the reservation id, it must only contain lower case alphanumeric characters or dashes.It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If the new reservation's slot capacity exceed the project's slot capacity or if total slot capacity of the new reservation and its siblings exceeds the project's slot capacity, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
	SlotCapacity pulumi.StringOutput `pulumi:"slotCapacity"`
	// Last update time of the reservation.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewReservation registers a new resource with the given unique name, arguments, and options.
func NewReservation(ctx *pulumi.Context,
	name string, args *ReservationArgs, opts ...pulumi.ResourceOption) (*Reservation, error) {
	if args == nil {
		args = &ReservationArgs{}
	}

	var resource Reservation
	err := ctx.RegisterResource("google-native:bigqueryreservation/v1:Reservation", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:bigqueryreservation/v1:Reservation", name, id, state, &resource, opts...)
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
	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
	Concurrency *string `pulumi:"concurrency"`
	// If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
	IgnoreIdleSlots *bool   `pulumi:"ignoreIdleSlots"`
	Location        *string `pulumi:"location"`
	// Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
	MultiRegionAuxiliary *bool `pulumi:"multiRegionAuxiliary"`
	// The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. For the reservation id, it must only contain lower case alphanumeric characters or dashes.It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The reservation ID. It must only contain lower case alphanumeric characters or dashes.It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
	ReservationId *string `pulumi:"reservationId"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If the new reservation's slot capacity exceed the project's slot capacity or if total slot capacity of the new reservation and its siblings exceeds the project's slot capacity, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
	SlotCapacity *string `pulumi:"slotCapacity"`
}

// The set of arguments for constructing a Reservation resource.
type ReservationArgs struct {
	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
	Concurrency pulumi.StringPtrInput
	// If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
	IgnoreIdleSlots pulumi.BoolPtrInput
	Location        pulumi.StringPtrInput
	// Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
	MultiRegionAuxiliary pulumi.BoolPtrInput
	// The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. For the reservation id, it must only contain lower case alphanumeric characters or dashes.It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The reservation ID. It must only contain lower case alphanumeric characters or dashes.It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
	ReservationId pulumi.StringPtrInput
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If the new reservation's slot capacity exceed the project's slot capacity or if total slot capacity of the new reservation and its siblings exceeds the project's slot capacity, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
	SlotCapacity pulumi.StringPtrInput
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
	return reflect.TypeOf((**Reservation)(nil)).Elem()
}

func (i *Reservation) ToReservationOutput() ReservationOutput {
	return i.ToReservationOutputWithContext(context.Background())
}

func (i *Reservation) ToReservationOutputWithContext(ctx context.Context) ReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationOutput)
}

type ReservationOutput struct{ *pulumi.OutputState }

func (ReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Reservation)(nil)).Elem()
}

func (o ReservationOutput) ToReservationOutput() ReservationOutput {
	return o
}

func (o ReservationOutput) ToReservationOutputWithContext(ctx context.Context) ReservationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReservationInput)(nil)).Elem(), &Reservation{})
	pulumi.RegisterOutputType(ReservationOutput{})
}
