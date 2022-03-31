// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The DeliveryRequirement for this subscription.
type DeliveryConfigDeliveryRequirement string

const (
	// Default value. This value is unused.
	DeliveryConfigDeliveryRequirementDeliveryRequirementUnspecified = DeliveryConfigDeliveryRequirement("DELIVERY_REQUIREMENT_UNSPECIFIED")
	// The server does not wait for a published message to be successfully written to storage before delivering it to subscribers.
	DeliveryConfigDeliveryRequirementDeliverImmediately = DeliveryConfigDeliveryRequirement("DELIVER_IMMEDIATELY")
	// The server will not deliver a published message to subscribers until the message has been successfully written to storage. This will result in higher end-to-end latency, but consistent delivery.
	DeliveryConfigDeliveryRequirementDeliverAfterStored = DeliveryConfigDeliveryRequirement("DELIVER_AFTER_STORED")
)

func (DeliveryConfigDeliveryRequirement) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryConfigDeliveryRequirement)(nil)).Elem()
}

func (e DeliveryConfigDeliveryRequirement) ToDeliveryConfigDeliveryRequirementOutput() DeliveryConfigDeliveryRequirementOutput {
	return pulumi.ToOutput(e).(DeliveryConfigDeliveryRequirementOutput)
}

func (e DeliveryConfigDeliveryRequirement) ToDeliveryConfigDeliveryRequirementOutputWithContext(ctx context.Context) DeliveryConfigDeliveryRequirementOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeliveryConfigDeliveryRequirementOutput)
}

func (e DeliveryConfigDeliveryRequirement) ToDeliveryConfigDeliveryRequirementPtrOutput() DeliveryConfigDeliveryRequirementPtrOutput {
	return e.ToDeliveryConfigDeliveryRequirementPtrOutputWithContext(context.Background())
}

func (e DeliveryConfigDeliveryRequirement) ToDeliveryConfigDeliveryRequirementPtrOutputWithContext(ctx context.Context) DeliveryConfigDeliveryRequirementPtrOutput {
	return DeliveryConfigDeliveryRequirement(e).ToDeliveryConfigDeliveryRequirementOutputWithContext(ctx).ToDeliveryConfigDeliveryRequirementPtrOutputWithContext(ctx)
}

func (e DeliveryConfigDeliveryRequirement) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeliveryConfigDeliveryRequirement) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeliveryConfigDeliveryRequirement) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeliveryConfigDeliveryRequirement) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeliveryConfigDeliveryRequirementOutput struct{ *pulumi.OutputState }

func (DeliveryConfigDeliveryRequirementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryConfigDeliveryRequirement)(nil)).Elem()
}

func (o DeliveryConfigDeliveryRequirementOutput) ToDeliveryConfigDeliveryRequirementOutput() DeliveryConfigDeliveryRequirementOutput {
	return o
}

func (o DeliveryConfigDeliveryRequirementOutput) ToDeliveryConfigDeliveryRequirementOutputWithContext(ctx context.Context) DeliveryConfigDeliveryRequirementOutput {
	return o
}

func (o DeliveryConfigDeliveryRequirementOutput) ToDeliveryConfigDeliveryRequirementPtrOutput() DeliveryConfigDeliveryRequirementPtrOutput {
	return o.ToDeliveryConfigDeliveryRequirementPtrOutputWithContext(context.Background())
}

func (o DeliveryConfigDeliveryRequirementOutput) ToDeliveryConfigDeliveryRequirementPtrOutputWithContext(ctx context.Context) DeliveryConfigDeliveryRequirementPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeliveryConfigDeliveryRequirement) *DeliveryConfigDeliveryRequirement {
		return &v
	}).(DeliveryConfigDeliveryRequirementPtrOutput)
}

func (o DeliveryConfigDeliveryRequirementOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeliveryConfigDeliveryRequirementOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeliveryConfigDeliveryRequirement) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeliveryConfigDeliveryRequirementOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeliveryConfigDeliveryRequirementOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeliveryConfigDeliveryRequirement) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeliveryConfigDeliveryRequirementPtrOutput struct{ *pulumi.OutputState }

func (DeliveryConfigDeliveryRequirementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryConfigDeliveryRequirement)(nil)).Elem()
}

func (o DeliveryConfigDeliveryRequirementPtrOutput) ToDeliveryConfigDeliveryRequirementPtrOutput() DeliveryConfigDeliveryRequirementPtrOutput {
	return o
}

func (o DeliveryConfigDeliveryRequirementPtrOutput) ToDeliveryConfigDeliveryRequirementPtrOutputWithContext(ctx context.Context) DeliveryConfigDeliveryRequirementPtrOutput {
	return o
}

func (o DeliveryConfigDeliveryRequirementPtrOutput) Elem() DeliveryConfigDeliveryRequirementOutput {
	return o.ApplyT(func(v *DeliveryConfigDeliveryRequirement) DeliveryConfigDeliveryRequirement {
		if v != nil {
			return *v
		}
		var ret DeliveryConfigDeliveryRequirement
		return ret
	}).(DeliveryConfigDeliveryRequirementOutput)
}

func (o DeliveryConfigDeliveryRequirementPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeliveryConfigDeliveryRequirementPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeliveryConfigDeliveryRequirement) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DeliveryConfigDeliveryRequirementInput is an input type that accepts DeliveryConfigDeliveryRequirementArgs and DeliveryConfigDeliveryRequirementOutput values.
// You can construct a concrete instance of `DeliveryConfigDeliveryRequirementInput` via:
//
//          DeliveryConfigDeliveryRequirementArgs{...}
type DeliveryConfigDeliveryRequirementInput interface {
	pulumi.Input

	ToDeliveryConfigDeliveryRequirementOutput() DeliveryConfigDeliveryRequirementOutput
	ToDeliveryConfigDeliveryRequirementOutputWithContext(context.Context) DeliveryConfigDeliveryRequirementOutput
}

var deliveryConfigDeliveryRequirementPtrType = reflect.TypeOf((**DeliveryConfigDeliveryRequirement)(nil)).Elem()

type DeliveryConfigDeliveryRequirementPtrInput interface {
	pulumi.Input

	ToDeliveryConfigDeliveryRequirementPtrOutput() DeliveryConfigDeliveryRequirementPtrOutput
	ToDeliveryConfigDeliveryRequirementPtrOutputWithContext(context.Context) DeliveryConfigDeliveryRequirementPtrOutput
}

type deliveryConfigDeliveryRequirementPtr string

func DeliveryConfigDeliveryRequirementPtr(v string) DeliveryConfigDeliveryRequirementPtrInput {
	return (*deliveryConfigDeliveryRequirementPtr)(&v)
}

func (*deliveryConfigDeliveryRequirementPtr) ElementType() reflect.Type {
	return deliveryConfigDeliveryRequirementPtrType
}

func (in *deliveryConfigDeliveryRequirementPtr) ToDeliveryConfigDeliveryRequirementPtrOutput() DeliveryConfigDeliveryRequirementPtrOutput {
	return pulumi.ToOutput(in).(DeliveryConfigDeliveryRequirementPtrOutput)
}

func (in *deliveryConfigDeliveryRequirementPtr) ToDeliveryConfigDeliveryRequirementPtrOutputWithContext(ctx context.Context) DeliveryConfigDeliveryRequirementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeliveryConfigDeliveryRequirementPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeliveryConfigDeliveryRequirementInput)(nil)).Elem(), DeliveryConfigDeliveryRequirement("DELIVERY_REQUIREMENT_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeliveryConfigDeliveryRequirementPtrInput)(nil)).Elem(), DeliveryConfigDeliveryRequirement("DELIVERY_REQUIREMENT_UNSPECIFIED"))
	pulumi.RegisterOutputType(DeliveryConfigDeliveryRequirementOutput{})
	pulumi.RegisterOutputType(DeliveryConfigDeliveryRequirementPtrOutput{})
}
