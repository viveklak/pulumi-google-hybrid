// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The throughput capacity configuration for each partition.
type Capacity struct {
	// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	PublishMibPerSec *int `pulumi:"publishMibPerSec"`
	// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 32.
	SubscribeMibPerSec *int `pulumi:"subscribeMibPerSec"`
}

// CapacityInput is an input type that accepts CapacityArgs and CapacityOutput values.
// You can construct a concrete instance of `CapacityInput` via:
//
//          CapacityArgs{...}
type CapacityInput interface {
	pulumi.Input

	ToCapacityOutput() CapacityOutput
	ToCapacityOutputWithContext(context.Context) CapacityOutput
}

// The throughput capacity configuration for each partition.
type CapacityArgs struct {
	// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	PublishMibPerSec pulumi.IntPtrInput `pulumi:"publishMibPerSec"`
	// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 32.
	SubscribeMibPerSec pulumi.IntPtrInput `pulumi:"subscribeMibPerSec"`
}

func (CapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Capacity)(nil)).Elem()
}

func (i CapacityArgs) ToCapacityOutput() CapacityOutput {
	return i.ToCapacityOutputWithContext(context.Background())
}

func (i CapacityArgs) ToCapacityOutputWithContext(ctx context.Context) CapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityOutput)
}

func (i CapacityArgs) ToCapacityPtrOutput() CapacityPtrOutput {
	return i.ToCapacityPtrOutputWithContext(context.Background())
}

func (i CapacityArgs) ToCapacityPtrOutputWithContext(ctx context.Context) CapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityOutput).ToCapacityPtrOutputWithContext(ctx)
}

// CapacityPtrInput is an input type that accepts CapacityArgs, CapacityPtr and CapacityPtrOutput values.
// You can construct a concrete instance of `CapacityPtrInput` via:
//
//          CapacityArgs{...}
//
//  or:
//
//          nil
type CapacityPtrInput interface {
	pulumi.Input

	ToCapacityPtrOutput() CapacityPtrOutput
	ToCapacityPtrOutputWithContext(context.Context) CapacityPtrOutput
}

type capacityPtrType CapacityArgs

func CapacityPtr(v *CapacityArgs) CapacityPtrInput {
	return (*capacityPtrType)(v)
}

func (*capacityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Capacity)(nil)).Elem()
}

func (i *capacityPtrType) ToCapacityPtrOutput() CapacityPtrOutput {
	return i.ToCapacityPtrOutputWithContext(context.Background())
}

func (i *capacityPtrType) ToCapacityPtrOutputWithContext(ctx context.Context) CapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityPtrOutput)
}

// The throughput capacity configuration for each partition.
type CapacityOutput struct{ *pulumi.OutputState }

func (CapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Capacity)(nil)).Elem()
}

func (o CapacityOutput) ToCapacityOutput() CapacityOutput {
	return o
}

func (o CapacityOutput) ToCapacityOutputWithContext(ctx context.Context) CapacityOutput {
	return o
}

func (o CapacityOutput) ToCapacityPtrOutput() CapacityPtrOutput {
	return o.ToCapacityPtrOutputWithContext(context.Background())
}

func (o CapacityOutput) ToCapacityPtrOutputWithContext(ctx context.Context) CapacityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Capacity) *Capacity {
		return &v
	}).(CapacityPtrOutput)
}

// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
func (o CapacityOutput) PublishMibPerSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Capacity) *int { return v.PublishMibPerSec }).(pulumi.IntPtrOutput)
}

// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 32.
func (o CapacityOutput) SubscribeMibPerSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Capacity) *int { return v.SubscribeMibPerSec }).(pulumi.IntPtrOutput)
}

type CapacityPtrOutput struct{ *pulumi.OutputState }

func (CapacityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Capacity)(nil)).Elem()
}

func (o CapacityPtrOutput) ToCapacityPtrOutput() CapacityPtrOutput {
	return o
}

func (o CapacityPtrOutput) ToCapacityPtrOutputWithContext(ctx context.Context) CapacityPtrOutput {
	return o
}

func (o CapacityPtrOutput) Elem() CapacityOutput {
	return o.ApplyT(func(v *Capacity) Capacity {
		if v != nil {
			return *v
		}
		var ret Capacity
		return ret
	}).(CapacityOutput)
}

// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
func (o CapacityPtrOutput) PublishMibPerSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Capacity) *int {
		if v == nil {
			return nil
		}
		return v.PublishMibPerSec
	}).(pulumi.IntPtrOutput)
}

// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 32.
func (o CapacityPtrOutput) SubscribeMibPerSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Capacity) *int {
		if v == nil {
			return nil
		}
		return v.SubscribeMibPerSec
	}).(pulumi.IntPtrOutput)
}

// The throughput capacity configuration for each partition.
type CapacityResponse struct {
	// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	PublishMibPerSec int `pulumi:"publishMibPerSec"`
	// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 32.
	SubscribeMibPerSec int `pulumi:"subscribeMibPerSec"`
}

// The throughput capacity configuration for each partition.
type CapacityResponseOutput struct{ *pulumi.OutputState }

func (CapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityResponse)(nil)).Elem()
}

func (o CapacityResponseOutput) ToCapacityResponseOutput() CapacityResponseOutput {
	return o
}

func (o CapacityResponseOutput) ToCapacityResponseOutputWithContext(ctx context.Context) CapacityResponseOutput {
	return o
}

// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
func (o CapacityResponseOutput) PublishMibPerSec() pulumi.IntOutput {
	return o.ApplyT(func(v CapacityResponse) int { return v.PublishMibPerSec }).(pulumi.IntOutput)
}

// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 32.
func (o CapacityResponseOutput) SubscribeMibPerSec() pulumi.IntOutput {
	return o.ApplyT(func(v CapacityResponse) int { return v.SubscribeMibPerSec }).(pulumi.IntOutput)
}

// The settings for a subscription's message delivery.
type DeliveryConfig struct {
	// The DeliveryRequirement for this subscription.
	DeliveryRequirement *DeliveryConfigDeliveryRequirement `pulumi:"deliveryRequirement"`
}

// DeliveryConfigInput is an input type that accepts DeliveryConfigArgs and DeliveryConfigOutput values.
// You can construct a concrete instance of `DeliveryConfigInput` via:
//
//          DeliveryConfigArgs{...}
type DeliveryConfigInput interface {
	pulumi.Input

	ToDeliveryConfigOutput() DeliveryConfigOutput
	ToDeliveryConfigOutputWithContext(context.Context) DeliveryConfigOutput
}

// The settings for a subscription's message delivery.
type DeliveryConfigArgs struct {
	// The DeliveryRequirement for this subscription.
	DeliveryRequirement DeliveryConfigDeliveryRequirementPtrInput `pulumi:"deliveryRequirement"`
}

func (DeliveryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryConfig)(nil)).Elem()
}

func (i DeliveryConfigArgs) ToDeliveryConfigOutput() DeliveryConfigOutput {
	return i.ToDeliveryConfigOutputWithContext(context.Background())
}

func (i DeliveryConfigArgs) ToDeliveryConfigOutputWithContext(ctx context.Context) DeliveryConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryConfigOutput)
}

func (i DeliveryConfigArgs) ToDeliveryConfigPtrOutput() DeliveryConfigPtrOutput {
	return i.ToDeliveryConfigPtrOutputWithContext(context.Background())
}

func (i DeliveryConfigArgs) ToDeliveryConfigPtrOutputWithContext(ctx context.Context) DeliveryConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryConfigOutput).ToDeliveryConfigPtrOutputWithContext(ctx)
}

// DeliveryConfigPtrInput is an input type that accepts DeliveryConfigArgs, DeliveryConfigPtr and DeliveryConfigPtrOutput values.
// You can construct a concrete instance of `DeliveryConfigPtrInput` via:
//
//          DeliveryConfigArgs{...}
//
//  or:
//
//          nil
type DeliveryConfigPtrInput interface {
	pulumi.Input

	ToDeliveryConfigPtrOutput() DeliveryConfigPtrOutput
	ToDeliveryConfigPtrOutputWithContext(context.Context) DeliveryConfigPtrOutput
}

type deliveryConfigPtrType DeliveryConfigArgs

func DeliveryConfigPtr(v *DeliveryConfigArgs) DeliveryConfigPtrInput {
	return (*deliveryConfigPtrType)(v)
}

func (*deliveryConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryConfig)(nil)).Elem()
}

func (i *deliveryConfigPtrType) ToDeliveryConfigPtrOutput() DeliveryConfigPtrOutput {
	return i.ToDeliveryConfigPtrOutputWithContext(context.Background())
}

func (i *deliveryConfigPtrType) ToDeliveryConfigPtrOutputWithContext(ctx context.Context) DeliveryConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryConfigPtrOutput)
}

// The settings for a subscription's message delivery.
type DeliveryConfigOutput struct{ *pulumi.OutputState }

func (DeliveryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryConfig)(nil)).Elem()
}

func (o DeliveryConfigOutput) ToDeliveryConfigOutput() DeliveryConfigOutput {
	return o
}

func (o DeliveryConfigOutput) ToDeliveryConfigOutputWithContext(ctx context.Context) DeliveryConfigOutput {
	return o
}

func (o DeliveryConfigOutput) ToDeliveryConfigPtrOutput() DeliveryConfigPtrOutput {
	return o.ToDeliveryConfigPtrOutputWithContext(context.Background())
}

func (o DeliveryConfigOutput) ToDeliveryConfigPtrOutputWithContext(ctx context.Context) DeliveryConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeliveryConfig) *DeliveryConfig {
		return &v
	}).(DeliveryConfigPtrOutput)
}

// The DeliveryRequirement for this subscription.
func (o DeliveryConfigOutput) DeliveryRequirement() DeliveryConfigDeliveryRequirementPtrOutput {
	return o.ApplyT(func(v DeliveryConfig) *DeliveryConfigDeliveryRequirement { return v.DeliveryRequirement }).(DeliveryConfigDeliveryRequirementPtrOutput)
}

type DeliveryConfigPtrOutput struct{ *pulumi.OutputState }

func (DeliveryConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryConfig)(nil)).Elem()
}

func (o DeliveryConfigPtrOutput) ToDeliveryConfigPtrOutput() DeliveryConfigPtrOutput {
	return o
}

func (o DeliveryConfigPtrOutput) ToDeliveryConfigPtrOutputWithContext(ctx context.Context) DeliveryConfigPtrOutput {
	return o
}

func (o DeliveryConfigPtrOutput) Elem() DeliveryConfigOutput {
	return o.ApplyT(func(v *DeliveryConfig) DeliveryConfig {
		if v != nil {
			return *v
		}
		var ret DeliveryConfig
		return ret
	}).(DeliveryConfigOutput)
}

// The DeliveryRequirement for this subscription.
func (o DeliveryConfigPtrOutput) DeliveryRequirement() DeliveryConfigDeliveryRequirementPtrOutput {
	return o.ApplyT(func(v *DeliveryConfig) *DeliveryConfigDeliveryRequirement {
		if v == nil {
			return nil
		}
		return v.DeliveryRequirement
	}).(DeliveryConfigDeliveryRequirementPtrOutput)
}

// The settings for a subscription's message delivery.
type DeliveryConfigResponse struct {
	// The DeliveryRequirement for this subscription.
	DeliveryRequirement string `pulumi:"deliveryRequirement"`
}

// The settings for a subscription's message delivery.
type DeliveryConfigResponseOutput struct{ *pulumi.OutputState }

func (DeliveryConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryConfigResponse)(nil)).Elem()
}

func (o DeliveryConfigResponseOutput) ToDeliveryConfigResponseOutput() DeliveryConfigResponseOutput {
	return o
}

func (o DeliveryConfigResponseOutput) ToDeliveryConfigResponseOutputWithContext(ctx context.Context) DeliveryConfigResponseOutput {
	return o
}

// The DeliveryRequirement for this subscription.
func (o DeliveryConfigResponseOutput) DeliveryRequirement() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryConfigResponse) string { return v.DeliveryRequirement }).(pulumi.StringOutput)
}

// The settings for a topic's partitions.
type PartitionConfig struct {
	// The capacity configuration.
	Capacity *Capacity `pulumi:"capacity"`
	// The number of partitions in the topic. Must be at least 1. Once a topic has been created the number of partitions can be increased but not decreased. Message ordering is not guaranteed across a topic resize. For more information see https://cloud.google.com/pubsub/lite/docs/topics#scaling_capacity
	Count *string `pulumi:"count"`
}

// PartitionConfigInput is an input type that accepts PartitionConfigArgs and PartitionConfigOutput values.
// You can construct a concrete instance of `PartitionConfigInput` via:
//
//          PartitionConfigArgs{...}
type PartitionConfigInput interface {
	pulumi.Input

	ToPartitionConfigOutput() PartitionConfigOutput
	ToPartitionConfigOutputWithContext(context.Context) PartitionConfigOutput
}

// The settings for a topic's partitions.
type PartitionConfigArgs struct {
	// The capacity configuration.
	Capacity CapacityPtrInput `pulumi:"capacity"`
	// The number of partitions in the topic. Must be at least 1. Once a topic has been created the number of partitions can be increased but not decreased. Message ordering is not guaranteed across a topic resize. For more information see https://cloud.google.com/pubsub/lite/docs/topics#scaling_capacity
	Count pulumi.StringPtrInput `pulumi:"count"`
}

func (PartitionConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionConfig)(nil)).Elem()
}

func (i PartitionConfigArgs) ToPartitionConfigOutput() PartitionConfigOutput {
	return i.ToPartitionConfigOutputWithContext(context.Background())
}

func (i PartitionConfigArgs) ToPartitionConfigOutputWithContext(ctx context.Context) PartitionConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionConfigOutput)
}

func (i PartitionConfigArgs) ToPartitionConfigPtrOutput() PartitionConfigPtrOutput {
	return i.ToPartitionConfigPtrOutputWithContext(context.Background())
}

func (i PartitionConfigArgs) ToPartitionConfigPtrOutputWithContext(ctx context.Context) PartitionConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionConfigOutput).ToPartitionConfigPtrOutputWithContext(ctx)
}

// PartitionConfigPtrInput is an input type that accepts PartitionConfigArgs, PartitionConfigPtr and PartitionConfigPtrOutput values.
// You can construct a concrete instance of `PartitionConfigPtrInput` via:
//
//          PartitionConfigArgs{...}
//
//  or:
//
//          nil
type PartitionConfigPtrInput interface {
	pulumi.Input

	ToPartitionConfigPtrOutput() PartitionConfigPtrOutput
	ToPartitionConfigPtrOutputWithContext(context.Context) PartitionConfigPtrOutput
}

type partitionConfigPtrType PartitionConfigArgs

func PartitionConfigPtr(v *PartitionConfigArgs) PartitionConfigPtrInput {
	return (*partitionConfigPtrType)(v)
}

func (*partitionConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PartitionConfig)(nil)).Elem()
}

func (i *partitionConfigPtrType) ToPartitionConfigPtrOutput() PartitionConfigPtrOutput {
	return i.ToPartitionConfigPtrOutputWithContext(context.Background())
}

func (i *partitionConfigPtrType) ToPartitionConfigPtrOutputWithContext(ctx context.Context) PartitionConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionConfigPtrOutput)
}

// The settings for a topic's partitions.
type PartitionConfigOutput struct{ *pulumi.OutputState }

func (PartitionConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionConfig)(nil)).Elem()
}

func (o PartitionConfigOutput) ToPartitionConfigOutput() PartitionConfigOutput {
	return o
}

func (o PartitionConfigOutput) ToPartitionConfigOutputWithContext(ctx context.Context) PartitionConfigOutput {
	return o
}

func (o PartitionConfigOutput) ToPartitionConfigPtrOutput() PartitionConfigPtrOutput {
	return o.ToPartitionConfigPtrOutputWithContext(context.Background())
}

func (o PartitionConfigOutput) ToPartitionConfigPtrOutputWithContext(ctx context.Context) PartitionConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartitionConfig) *PartitionConfig {
		return &v
	}).(PartitionConfigPtrOutput)
}

// The capacity configuration.
func (o PartitionConfigOutput) Capacity() CapacityPtrOutput {
	return o.ApplyT(func(v PartitionConfig) *Capacity { return v.Capacity }).(CapacityPtrOutput)
}

// The number of partitions in the topic. Must be at least 1. Once a topic has been created the number of partitions can be increased but not decreased. Message ordering is not guaranteed across a topic resize. For more information see https://cloud.google.com/pubsub/lite/docs/topics#scaling_capacity
func (o PartitionConfigOutput) Count() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartitionConfig) *string { return v.Count }).(pulumi.StringPtrOutput)
}

type PartitionConfigPtrOutput struct{ *pulumi.OutputState }

func (PartitionConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartitionConfig)(nil)).Elem()
}

func (o PartitionConfigPtrOutput) ToPartitionConfigPtrOutput() PartitionConfigPtrOutput {
	return o
}

func (o PartitionConfigPtrOutput) ToPartitionConfigPtrOutputWithContext(ctx context.Context) PartitionConfigPtrOutput {
	return o
}

func (o PartitionConfigPtrOutput) Elem() PartitionConfigOutput {
	return o.ApplyT(func(v *PartitionConfig) PartitionConfig {
		if v != nil {
			return *v
		}
		var ret PartitionConfig
		return ret
	}).(PartitionConfigOutput)
}

// The capacity configuration.
func (o PartitionConfigPtrOutput) Capacity() CapacityPtrOutput {
	return o.ApplyT(func(v *PartitionConfig) *Capacity {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(CapacityPtrOutput)
}

// The number of partitions in the topic. Must be at least 1. Once a topic has been created the number of partitions can be increased but not decreased. Message ordering is not guaranteed across a topic resize. For more information see https://cloud.google.com/pubsub/lite/docs/topics#scaling_capacity
func (o PartitionConfigPtrOutput) Count() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartitionConfig) *string {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.StringPtrOutput)
}

// The settings for a topic's partitions.
type PartitionConfigResponse struct {
	// The capacity configuration.
	Capacity CapacityResponse `pulumi:"capacity"`
	// The number of partitions in the topic. Must be at least 1. Once a topic has been created the number of partitions can be increased but not decreased. Message ordering is not guaranteed across a topic resize. For more information see https://cloud.google.com/pubsub/lite/docs/topics#scaling_capacity
	Count string `pulumi:"count"`
}

// The settings for a topic's partitions.
type PartitionConfigResponseOutput struct{ *pulumi.OutputState }

func (PartitionConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionConfigResponse)(nil)).Elem()
}

func (o PartitionConfigResponseOutput) ToPartitionConfigResponseOutput() PartitionConfigResponseOutput {
	return o
}

func (o PartitionConfigResponseOutput) ToPartitionConfigResponseOutputWithContext(ctx context.Context) PartitionConfigResponseOutput {
	return o
}

// The capacity configuration.
func (o PartitionConfigResponseOutput) Capacity() CapacityResponseOutput {
	return o.ApplyT(func(v PartitionConfigResponse) CapacityResponse { return v.Capacity }).(CapacityResponseOutput)
}

// The number of partitions in the topic. Must be at least 1. Once a topic has been created the number of partitions can be increased but not decreased. Message ordering is not guaranteed across a topic resize. For more information see https://cloud.google.com/pubsub/lite/docs/topics#scaling_capacity
func (o PartitionConfigResponseOutput) Count() pulumi.StringOutput {
	return o.ApplyT(func(v PartitionConfigResponse) string { return v.Count }).(pulumi.StringOutput)
}

// The settings for this topic's Reservation usage.
type ReservationConfig struct {
	// The Reservation to use for this topic's throughput capacity. Structured like: projects/{project_number}/locations/{location}/reservations/{reservation_id}
	ThroughputReservation *string `pulumi:"throughputReservation"`
}

// ReservationConfigInput is an input type that accepts ReservationConfigArgs and ReservationConfigOutput values.
// You can construct a concrete instance of `ReservationConfigInput` via:
//
//          ReservationConfigArgs{...}
type ReservationConfigInput interface {
	pulumi.Input

	ToReservationConfigOutput() ReservationConfigOutput
	ToReservationConfigOutputWithContext(context.Context) ReservationConfigOutput
}

// The settings for this topic's Reservation usage.
type ReservationConfigArgs struct {
	// The Reservation to use for this topic's throughput capacity. Structured like: projects/{project_number}/locations/{location}/reservations/{reservation_id}
	ThroughputReservation pulumi.StringPtrInput `pulumi:"throughputReservation"`
}

func (ReservationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReservationConfig)(nil)).Elem()
}

func (i ReservationConfigArgs) ToReservationConfigOutput() ReservationConfigOutput {
	return i.ToReservationConfigOutputWithContext(context.Background())
}

func (i ReservationConfigArgs) ToReservationConfigOutputWithContext(ctx context.Context) ReservationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationConfigOutput)
}

func (i ReservationConfigArgs) ToReservationConfigPtrOutput() ReservationConfigPtrOutput {
	return i.ToReservationConfigPtrOutputWithContext(context.Background())
}

func (i ReservationConfigArgs) ToReservationConfigPtrOutputWithContext(ctx context.Context) ReservationConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationConfigOutput).ToReservationConfigPtrOutputWithContext(ctx)
}

// ReservationConfigPtrInput is an input type that accepts ReservationConfigArgs, ReservationConfigPtr and ReservationConfigPtrOutput values.
// You can construct a concrete instance of `ReservationConfigPtrInput` via:
//
//          ReservationConfigArgs{...}
//
//  or:
//
//          nil
type ReservationConfigPtrInput interface {
	pulumi.Input

	ToReservationConfigPtrOutput() ReservationConfigPtrOutput
	ToReservationConfigPtrOutputWithContext(context.Context) ReservationConfigPtrOutput
}

type reservationConfigPtrType ReservationConfigArgs

func ReservationConfigPtr(v *ReservationConfigArgs) ReservationConfigPtrInput {
	return (*reservationConfigPtrType)(v)
}

func (*reservationConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReservationConfig)(nil)).Elem()
}

func (i *reservationConfigPtrType) ToReservationConfigPtrOutput() ReservationConfigPtrOutput {
	return i.ToReservationConfigPtrOutputWithContext(context.Background())
}

func (i *reservationConfigPtrType) ToReservationConfigPtrOutputWithContext(ctx context.Context) ReservationConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationConfigPtrOutput)
}

// The settings for this topic's Reservation usage.
type ReservationConfigOutput struct{ *pulumi.OutputState }

func (ReservationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReservationConfig)(nil)).Elem()
}

func (o ReservationConfigOutput) ToReservationConfigOutput() ReservationConfigOutput {
	return o
}

func (o ReservationConfigOutput) ToReservationConfigOutputWithContext(ctx context.Context) ReservationConfigOutput {
	return o
}

func (o ReservationConfigOutput) ToReservationConfigPtrOutput() ReservationConfigPtrOutput {
	return o.ToReservationConfigPtrOutputWithContext(context.Background())
}

func (o ReservationConfigOutput) ToReservationConfigPtrOutputWithContext(ctx context.Context) ReservationConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReservationConfig) *ReservationConfig {
		return &v
	}).(ReservationConfigPtrOutput)
}

// The Reservation to use for this topic's throughput capacity. Structured like: projects/{project_number}/locations/{location}/reservations/{reservation_id}
func (o ReservationConfigOutput) ThroughputReservation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReservationConfig) *string { return v.ThroughputReservation }).(pulumi.StringPtrOutput)
}

type ReservationConfigPtrOutput struct{ *pulumi.OutputState }

func (ReservationConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReservationConfig)(nil)).Elem()
}

func (o ReservationConfigPtrOutput) ToReservationConfigPtrOutput() ReservationConfigPtrOutput {
	return o
}

func (o ReservationConfigPtrOutput) ToReservationConfigPtrOutputWithContext(ctx context.Context) ReservationConfigPtrOutput {
	return o
}

func (o ReservationConfigPtrOutput) Elem() ReservationConfigOutput {
	return o.ApplyT(func(v *ReservationConfig) ReservationConfig {
		if v != nil {
			return *v
		}
		var ret ReservationConfig
		return ret
	}).(ReservationConfigOutput)
}

// The Reservation to use for this topic's throughput capacity. Structured like: projects/{project_number}/locations/{location}/reservations/{reservation_id}
func (o ReservationConfigPtrOutput) ThroughputReservation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReservationConfig) *string {
		if v == nil {
			return nil
		}
		return v.ThroughputReservation
	}).(pulumi.StringPtrOutput)
}

// The settings for this topic's Reservation usage.
type ReservationConfigResponse struct {
	// The Reservation to use for this topic's throughput capacity. Structured like: projects/{project_number}/locations/{location}/reservations/{reservation_id}
	ThroughputReservation string `pulumi:"throughputReservation"`
}

// The settings for this topic's Reservation usage.
type ReservationConfigResponseOutput struct{ *pulumi.OutputState }

func (ReservationConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReservationConfigResponse)(nil)).Elem()
}

func (o ReservationConfigResponseOutput) ToReservationConfigResponseOutput() ReservationConfigResponseOutput {
	return o
}

func (o ReservationConfigResponseOutput) ToReservationConfigResponseOutputWithContext(ctx context.Context) ReservationConfigResponseOutput {
	return o
}

// The Reservation to use for this topic's throughput capacity. Structured like: projects/{project_number}/locations/{location}/reservations/{reservation_id}
func (o ReservationConfigResponseOutput) ThroughputReservation() pulumi.StringOutput {
	return o.ApplyT(func(v ReservationConfigResponse) string { return v.ThroughputReservation }).(pulumi.StringOutput)
}

// The settings for a topic's message retention.
type RetentionConfig struct {
	// The provisioned storage, in bytes, per partition. If the number of bytes stored in any of the topic's partitions grows beyond this value, older messages will be dropped to make room for newer ones, regardless of the value of `period`.
	PerPartitionBytes *string `pulumi:"perPartitionBytes"`
	// How long a published message is retained. If unset, messages will be retained as long as the bytes retained for each partition is below `per_partition_bytes`.
	Period *string `pulumi:"period"`
}

// RetentionConfigInput is an input type that accepts RetentionConfigArgs and RetentionConfigOutput values.
// You can construct a concrete instance of `RetentionConfigInput` via:
//
//          RetentionConfigArgs{...}
type RetentionConfigInput interface {
	pulumi.Input

	ToRetentionConfigOutput() RetentionConfigOutput
	ToRetentionConfigOutputWithContext(context.Context) RetentionConfigOutput
}

// The settings for a topic's message retention.
type RetentionConfigArgs struct {
	// The provisioned storage, in bytes, per partition. If the number of bytes stored in any of the topic's partitions grows beyond this value, older messages will be dropped to make room for newer ones, regardless of the value of `period`.
	PerPartitionBytes pulumi.StringPtrInput `pulumi:"perPartitionBytes"`
	// How long a published message is retained. If unset, messages will be retained as long as the bytes retained for each partition is below `per_partition_bytes`.
	Period pulumi.StringPtrInput `pulumi:"period"`
}

func (RetentionConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionConfig)(nil)).Elem()
}

func (i RetentionConfigArgs) ToRetentionConfigOutput() RetentionConfigOutput {
	return i.ToRetentionConfigOutputWithContext(context.Background())
}

func (i RetentionConfigArgs) ToRetentionConfigOutputWithContext(ctx context.Context) RetentionConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionConfigOutput)
}

func (i RetentionConfigArgs) ToRetentionConfigPtrOutput() RetentionConfigPtrOutput {
	return i.ToRetentionConfigPtrOutputWithContext(context.Background())
}

func (i RetentionConfigArgs) ToRetentionConfigPtrOutputWithContext(ctx context.Context) RetentionConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionConfigOutput).ToRetentionConfigPtrOutputWithContext(ctx)
}

// RetentionConfigPtrInput is an input type that accepts RetentionConfigArgs, RetentionConfigPtr and RetentionConfigPtrOutput values.
// You can construct a concrete instance of `RetentionConfigPtrInput` via:
//
//          RetentionConfigArgs{...}
//
//  or:
//
//          nil
type RetentionConfigPtrInput interface {
	pulumi.Input

	ToRetentionConfigPtrOutput() RetentionConfigPtrOutput
	ToRetentionConfigPtrOutputWithContext(context.Context) RetentionConfigPtrOutput
}

type retentionConfigPtrType RetentionConfigArgs

func RetentionConfigPtr(v *RetentionConfigArgs) RetentionConfigPtrInput {
	return (*retentionConfigPtrType)(v)
}

func (*retentionConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionConfig)(nil)).Elem()
}

func (i *retentionConfigPtrType) ToRetentionConfigPtrOutput() RetentionConfigPtrOutput {
	return i.ToRetentionConfigPtrOutputWithContext(context.Background())
}

func (i *retentionConfigPtrType) ToRetentionConfigPtrOutputWithContext(ctx context.Context) RetentionConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionConfigPtrOutput)
}

// The settings for a topic's message retention.
type RetentionConfigOutput struct{ *pulumi.OutputState }

func (RetentionConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionConfig)(nil)).Elem()
}

func (o RetentionConfigOutput) ToRetentionConfigOutput() RetentionConfigOutput {
	return o
}

func (o RetentionConfigOutput) ToRetentionConfigOutputWithContext(ctx context.Context) RetentionConfigOutput {
	return o
}

func (o RetentionConfigOutput) ToRetentionConfigPtrOutput() RetentionConfigPtrOutput {
	return o.ToRetentionConfigPtrOutputWithContext(context.Background())
}

func (o RetentionConfigOutput) ToRetentionConfigPtrOutputWithContext(ctx context.Context) RetentionConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionConfig) *RetentionConfig {
		return &v
	}).(RetentionConfigPtrOutput)
}

// The provisioned storage, in bytes, per partition. If the number of bytes stored in any of the topic's partitions grows beyond this value, older messages will be dropped to make room for newer ones, regardless of the value of `period`.
func (o RetentionConfigOutput) PerPartitionBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionConfig) *string { return v.PerPartitionBytes }).(pulumi.StringPtrOutput)
}

// How long a published message is retained. If unset, messages will be retained as long as the bytes retained for each partition is below `per_partition_bytes`.
func (o RetentionConfigOutput) Period() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionConfig) *string { return v.Period }).(pulumi.StringPtrOutput)
}

type RetentionConfigPtrOutput struct{ *pulumi.OutputState }

func (RetentionConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionConfig)(nil)).Elem()
}

func (o RetentionConfigPtrOutput) ToRetentionConfigPtrOutput() RetentionConfigPtrOutput {
	return o
}

func (o RetentionConfigPtrOutput) ToRetentionConfigPtrOutputWithContext(ctx context.Context) RetentionConfigPtrOutput {
	return o
}

func (o RetentionConfigPtrOutput) Elem() RetentionConfigOutput {
	return o.ApplyT(func(v *RetentionConfig) RetentionConfig {
		if v != nil {
			return *v
		}
		var ret RetentionConfig
		return ret
	}).(RetentionConfigOutput)
}

// The provisioned storage, in bytes, per partition. If the number of bytes stored in any of the topic's partitions grows beyond this value, older messages will be dropped to make room for newer ones, regardless of the value of `period`.
func (o RetentionConfigPtrOutput) PerPartitionBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionConfig) *string {
		if v == nil {
			return nil
		}
		return v.PerPartitionBytes
	}).(pulumi.StringPtrOutput)
}

// How long a published message is retained. If unset, messages will be retained as long as the bytes retained for each partition is below `per_partition_bytes`.
func (o RetentionConfigPtrOutput) Period() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionConfig) *string {
		if v == nil {
			return nil
		}
		return v.Period
	}).(pulumi.StringPtrOutput)
}

// The settings for a topic's message retention.
type RetentionConfigResponse struct {
	// The provisioned storage, in bytes, per partition. If the number of bytes stored in any of the topic's partitions grows beyond this value, older messages will be dropped to make room for newer ones, regardless of the value of `period`.
	PerPartitionBytes string `pulumi:"perPartitionBytes"`
	// How long a published message is retained. If unset, messages will be retained as long as the bytes retained for each partition is below `per_partition_bytes`.
	Period string `pulumi:"period"`
}

// The settings for a topic's message retention.
type RetentionConfigResponseOutput struct{ *pulumi.OutputState }

func (RetentionConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionConfigResponse)(nil)).Elem()
}

func (o RetentionConfigResponseOutput) ToRetentionConfigResponseOutput() RetentionConfigResponseOutput {
	return o
}

func (o RetentionConfigResponseOutput) ToRetentionConfigResponseOutputWithContext(ctx context.Context) RetentionConfigResponseOutput {
	return o
}

// The provisioned storage, in bytes, per partition. If the number of bytes stored in any of the topic's partitions grows beyond this value, older messages will be dropped to make room for newer ones, regardless of the value of `period`.
func (o RetentionConfigResponseOutput) PerPartitionBytes() pulumi.StringOutput {
	return o.ApplyT(func(v RetentionConfigResponse) string { return v.PerPartitionBytes }).(pulumi.StringOutput)
}

// How long a published message is retained. If unset, messages will be retained as long as the bytes retained for each partition is below `per_partition_bytes`.
func (o RetentionConfigResponseOutput) Period() pulumi.StringOutput {
	return o.ApplyT(func(v RetentionConfigResponse) string { return v.Period }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CapacityInput)(nil)).Elem(), CapacityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CapacityPtrInput)(nil)).Elem(), CapacityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeliveryConfigInput)(nil)).Elem(), DeliveryConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeliveryConfigPtrInput)(nil)).Elem(), DeliveryConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionConfigInput)(nil)).Elem(), PartitionConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionConfigPtrInput)(nil)).Elem(), PartitionConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReservationConfigInput)(nil)).Elem(), ReservationConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReservationConfigPtrInput)(nil)).Elem(), ReservationConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RetentionConfigInput)(nil)).Elem(), RetentionConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RetentionConfigPtrInput)(nil)).Elem(), RetentionConfigArgs{})
	pulumi.RegisterOutputType(CapacityOutput{})
	pulumi.RegisterOutputType(CapacityPtrOutput{})
	pulumi.RegisterOutputType(CapacityResponseOutput{})
	pulumi.RegisterOutputType(DeliveryConfigOutput{})
	pulumi.RegisterOutputType(DeliveryConfigPtrOutput{})
	pulumi.RegisterOutputType(DeliveryConfigResponseOutput{})
	pulumi.RegisterOutputType(PartitionConfigOutput{})
	pulumi.RegisterOutputType(PartitionConfigPtrOutput{})
	pulumi.RegisterOutputType(PartitionConfigResponseOutput{})
	pulumi.RegisterOutputType(ReservationConfigOutput{})
	pulumi.RegisterOutputType(ReservationConfigPtrOutput{})
	pulumi.RegisterOutputType(ReservationConfigResponseOutput{})
	pulumi.RegisterOutputType(RetentionConfigOutput{})
	pulumi.RegisterOutputType(RetentionConfigPtrOutput{})
	pulumi.RegisterOutputType(RetentionConfigResponseOutput{})
}
