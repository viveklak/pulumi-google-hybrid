// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of an accelator for a CDF instance.
type AcceleratorAcceleratorType string

const (
	// Default value, if unspecified.
	AcceleratorAcceleratorTypeAcceleratorTypeUnspecified = AcceleratorAcceleratorType("ACCELERATOR_TYPE_UNSPECIFIED")
	// Change Data Capture accelerator for CDF.
	AcceleratorAcceleratorTypeCdc = AcceleratorAcceleratorType("CDC")
	// Cloud Healthcare accelerator for CDF. This accelerator is to enable Cloud Healthcare specific CDF plugins developed by Healthcare team.
	AcceleratorAcceleratorTypeHealthcare = AcceleratorAcceleratorType("HEALTHCARE")
	// Contact Center AI Insights This accelerator is used to enable import and export pipelines custom built to streamline CCAI Insights processing.
	AcceleratorAcceleratorTypeCcaiInsights = AcceleratorAcceleratorType("CCAI_INSIGHTS")
)

func (AcceleratorAcceleratorType) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceleratorAcceleratorType)(nil)).Elem()
}

func (e AcceleratorAcceleratorType) ToAcceleratorAcceleratorTypeOutput() AcceleratorAcceleratorTypeOutput {
	return pulumi.ToOutput(e).(AcceleratorAcceleratorTypeOutput)
}

func (e AcceleratorAcceleratorType) ToAcceleratorAcceleratorTypeOutputWithContext(ctx context.Context) AcceleratorAcceleratorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AcceleratorAcceleratorTypeOutput)
}

func (e AcceleratorAcceleratorType) ToAcceleratorAcceleratorTypePtrOutput() AcceleratorAcceleratorTypePtrOutput {
	return e.ToAcceleratorAcceleratorTypePtrOutputWithContext(context.Background())
}

func (e AcceleratorAcceleratorType) ToAcceleratorAcceleratorTypePtrOutputWithContext(ctx context.Context) AcceleratorAcceleratorTypePtrOutput {
	return AcceleratorAcceleratorType(e).ToAcceleratorAcceleratorTypeOutputWithContext(ctx).ToAcceleratorAcceleratorTypePtrOutputWithContext(ctx)
}

func (e AcceleratorAcceleratorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AcceleratorAcceleratorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AcceleratorAcceleratorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AcceleratorAcceleratorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AcceleratorAcceleratorTypeOutput struct{ *pulumi.OutputState }

func (AcceleratorAcceleratorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceleratorAcceleratorType)(nil)).Elem()
}

func (o AcceleratorAcceleratorTypeOutput) ToAcceleratorAcceleratorTypeOutput() AcceleratorAcceleratorTypeOutput {
	return o
}

func (o AcceleratorAcceleratorTypeOutput) ToAcceleratorAcceleratorTypeOutputWithContext(ctx context.Context) AcceleratorAcceleratorTypeOutput {
	return o
}

func (o AcceleratorAcceleratorTypeOutput) ToAcceleratorAcceleratorTypePtrOutput() AcceleratorAcceleratorTypePtrOutput {
	return o.ToAcceleratorAcceleratorTypePtrOutputWithContext(context.Background())
}

func (o AcceleratorAcceleratorTypeOutput) ToAcceleratorAcceleratorTypePtrOutputWithContext(ctx context.Context) AcceleratorAcceleratorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AcceleratorAcceleratorType) *AcceleratorAcceleratorType {
		return &v
	}).(AcceleratorAcceleratorTypePtrOutput)
}

func (o AcceleratorAcceleratorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AcceleratorAcceleratorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AcceleratorAcceleratorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AcceleratorAcceleratorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AcceleratorAcceleratorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AcceleratorAcceleratorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AcceleratorAcceleratorTypePtrOutput struct{ *pulumi.OutputState }

func (AcceleratorAcceleratorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AcceleratorAcceleratorType)(nil)).Elem()
}

func (o AcceleratorAcceleratorTypePtrOutput) ToAcceleratorAcceleratorTypePtrOutput() AcceleratorAcceleratorTypePtrOutput {
	return o
}

func (o AcceleratorAcceleratorTypePtrOutput) ToAcceleratorAcceleratorTypePtrOutputWithContext(ctx context.Context) AcceleratorAcceleratorTypePtrOutput {
	return o
}

func (o AcceleratorAcceleratorTypePtrOutput) Elem() AcceleratorAcceleratorTypeOutput {
	return o.ApplyT(func(v *AcceleratorAcceleratorType) AcceleratorAcceleratorType {
		if v != nil {
			return *v
		}
		var ret AcceleratorAcceleratorType
		return ret
	}).(AcceleratorAcceleratorTypeOutput)
}

func (o AcceleratorAcceleratorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AcceleratorAcceleratorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AcceleratorAcceleratorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AcceleratorAcceleratorTypeInput is an input type that accepts AcceleratorAcceleratorTypeArgs and AcceleratorAcceleratorTypeOutput values.
// You can construct a concrete instance of `AcceleratorAcceleratorTypeInput` via:
//
//          AcceleratorAcceleratorTypeArgs{...}
type AcceleratorAcceleratorTypeInput interface {
	pulumi.Input

	ToAcceleratorAcceleratorTypeOutput() AcceleratorAcceleratorTypeOutput
	ToAcceleratorAcceleratorTypeOutputWithContext(context.Context) AcceleratorAcceleratorTypeOutput
}

var acceleratorAcceleratorTypePtrType = reflect.TypeOf((**AcceleratorAcceleratorType)(nil)).Elem()

type AcceleratorAcceleratorTypePtrInput interface {
	pulumi.Input

	ToAcceleratorAcceleratorTypePtrOutput() AcceleratorAcceleratorTypePtrOutput
	ToAcceleratorAcceleratorTypePtrOutputWithContext(context.Context) AcceleratorAcceleratorTypePtrOutput
}

type acceleratorAcceleratorTypePtr string

func AcceleratorAcceleratorTypePtr(v string) AcceleratorAcceleratorTypePtrInput {
	return (*acceleratorAcceleratorTypePtr)(&v)
}

func (*acceleratorAcceleratorTypePtr) ElementType() reflect.Type {
	return acceleratorAcceleratorTypePtrType
}

func (in *acceleratorAcceleratorTypePtr) ToAcceleratorAcceleratorTypePtrOutput() AcceleratorAcceleratorTypePtrOutput {
	return pulumi.ToOutput(in).(AcceleratorAcceleratorTypePtrOutput)
}

func (in *acceleratorAcceleratorTypePtr) ToAcceleratorAcceleratorTypePtrOutputWithContext(ctx context.Context) AcceleratorAcceleratorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AcceleratorAcceleratorTypePtrOutput)
}

// The state of the accelerator
type AcceleratorState string

const (
	// Default value, do not use
	AcceleratorStateStateUnspecified = AcceleratorState("STATE_UNSPECIFIED")
	// Indicates that the accelerator is enabled and available to use
	AcceleratorStateEnabled = AcceleratorState("ENABLED")
	// Indicates that the accelerator is disabled and not available to use
	AcceleratorStateDisabled = AcceleratorState("DISABLED")
	// Indicates that accelerator state is currently unknown. Requests for enable, disable could be retried while in this state
	AcceleratorStateUnknown = AcceleratorState("UNKNOWN")
)

func (AcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceleratorState)(nil)).Elem()
}

func (e AcceleratorState) ToAcceleratorStateOutput() AcceleratorStateOutput {
	return pulumi.ToOutput(e).(AcceleratorStateOutput)
}

func (e AcceleratorState) ToAcceleratorStateOutputWithContext(ctx context.Context) AcceleratorStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AcceleratorStateOutput)
}

func (e AcceleratorState) ToAcceleratorStatePtrOutput() AcceleratorStatePtrOutput {
	return e.ToAcceleratorStatePtrOutputWithContext(context.Background())
}

func (e AcceleratorState) ToAcceleratorStatePtrOutputWithContext(ctx context.Context) AcceleratorStatePtrOutput {
	return AcceleratorState(e).ToAcceleratorStateOutputWithContext(ctx).ToAcceleratorStatePtrOutputWithContext(ctx)
}

func (e AcceleratorState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AcceleratorState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AcceleratorState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AcceleratorState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AcceleratorStateOutput struct{ *pulumi.OutputState }

func (AcceleratorStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceleratorState)(nil)).Elem()
}

func (o AcceleratorStateOutput) ToAcceleratorStateOutput() AcceleratorStateOutput {
	return o
}

func (o AcceleratorStateOutput) ToAcceleratorStateOutputWithContext(ctx context.Context) AcceleratorStateOutput {
	return o
}

func (o AcceleratorStateOutput) ToAcceleratorStatePtrOutput() AcceleratorStatePtrOutput {
	return o.ToAcceleratorStatePtrOutputWithContext(context.Background())
}

func (o AcceleratorStateOutput) ToAcceleratorStatePtrOutputWithContext(ctx context.Context) AcceleratorStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AcceleratorState) *AcceleratorState {
		return &v
	}).(AcceleratorStatePtrOutput)
}

func (o AcceleratorStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AcceleratorStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AcceleratorState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AcceleratorStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AcceleratorStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AcceleratorState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AcceleratorStatePtrOutput struct{ *pulumi.OutputState }

func (AcceleratorStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AcceleratorState)(nil)).Elem()
}

func (o AcceleratorStatePtrOutput) ToAcceleratorStatePtrOutput() AcceleratorStatePtrOutput {
	return o
}

func (o AcceleratorStatePtrOutput) ToAcceleratorStatePtrOutputWithContext(ctx context.Context) AcceleratorStatePtrOutput {
	return o
}

func (o AcceleratorStatePtrOutput) Elem() AcceleratorStateOutput {
	return o.ApplyT(func(v *AcceleratorState) AcceleratorState {
		if v != nil {
			return *v
		}
		var ret AcceleratorState
		return ret
	}).(AcceleratorStateOutput)
}

func (o AcceleratorStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AcceleratorStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AcceleratorState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AcceleratorStateInput is an input type that accepts AcceleratorStateArgs and AcceleratorStateOutput values.
// You can construct a concrete instance of `AcceleratorStateInput` via:
//
//          AcceleratorStateArgs{...}
type AcceleratorStateInput interface {
	pulumi.Input

	ToAcceleratorStateOutput() AcceleratorStateOutput
	ToAcceleratorStateOutputWithContext(context.Context) AcceleratorStateOutput
}

var acceleratorStatePtrType = reflect.TypeOf((**AcceleratorState)(nil)).Elem()

type AcceleratorStatePtrInput interface {
	pulumi.Input

	ToAcceleratorStatePtrOutput() AcceleratorStatePtrOutput
	ToAcceleratorStatePtrOutputWithContext(context.Context) AcceleratorStatePtrOutput
}

type acceleratorStatePtr string

func AcceleratorStatePtr(v string) AcceleratorStatePtrInput {
	return (*acceleratorStatePtr)(&v)
}

func (*acceleratorStatePtr) ElementType() reflect.Type {
	return acceleratorStatePtrType
}

func (in *acceleratorStatePtr) ToAcceleratorStatePtrOutput() AcceleratorStatePtrOutput {
	return pulumi.ToOutput(in).(AcceleratorStatePtrOutput)
}

func (in *acceleratorStatePtr) ToAcceleratorStatePtrOutputWithContext(ctx context.Context) AcceleratorStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AcceleratorStatePtrOutput)
}

// The log type that this config enables.
type AuditLogConfigLogType string

const (
	// Default case. Should never be this.
	AuditLogConfigLogTypeLogTypeUnspecified = AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED")
	// Admin reads. Example: CloudIAM getIamPolicy
	AuditLogConfigLogTypeAdminRead = AuditLogConfigLogType("ADMIN_READ")
	// Data writes. Example: CloudSQL Users create
	AuditLogConfigLogTypeDataWrite = AuditLogConfigLogType("DATA_WRITE")
	// Data reads. Example: CloudSQL Users list
	AuditLogConfigLogTypeDataRead = AuditLogConfigLogType("DATA_READ")
)

func (AuditLogConfigLogType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogConfigLogType)(nil)).Elem()
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput {
	return pulumi.ToOutput(e).(AuditLogConfigLogTypeOutput)
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypeOutputWithContext(ctx context.Context) AuditLogConfigLogTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuditLogConfigLogTypeOutput)
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return e.ToAuditLogConfigLogTypePtrOutputWithContext(context.Background())
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return AuditLogConfigLogType(e).ToAuditLogConfigLogTypeOutputWithContext(ctx).ToAuditLogConfigLogTypePtrOutputWithContext(ctx)
}

func (e AuditLogConfigLogType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogConfigLogType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogConfigLogType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuditLogConfigLogType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuditLogConfigLogTypeOutput struct{ *pulumi.OutputState }

func (AuditLogConfigLogTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogConfigLogType)(nil)).Elem()
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput {
	return o
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypeOutputWithContext(ctx context.Context) AuditLogConfigLogTypeOutput {
	return o
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return o.ToAuditLogConfigLogTypePtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuditLogConfigLogType) *AuditLogConfigLogType {
		return &v
	}).(AuditLogConfigLogTypePtrOutput)
}

func (o AuditLogConfigLogTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogConfigLogType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuditLogConfigLogTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogConfigLogType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuditLogConfigLogTypePtrOutput struct{ *pulumi.OutputState }

func (AuditLogConfigLogTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditLogConfigLogType)(nil)).Elem()
}

func (o AuditLogConfigLogTypePtrOutput) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return o
}

func (o AuditLogConfigLogTypePtrOutput) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return o
}

func (o AuditLogConfigLogTypePtrOutput) Elem() AuditLogConfigLogTypeOutput {
	return o.ApplyT(func(v *AuditLogConfigLogType) AuditLogConfigLogType {
		if v != nil {
			return *v
		}
		var ret AuditLogConfigLogType
		return ret
	}).(AuditLogConfigLogTypeOutput)
}

func (o AuditLogConfigLogTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuditLogConfigLogType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AuditLogConfigLogTypeInput is an input type that accepts AuditLogConfigLogTypeArgs and AuditLogConfigLogTypeOutput values.
// You can construct a concrete instance of `AuditLogConfigLogTypeInput` via:
//
//          AuditLogConfigLogTypeArgs{...}
type AuditLogConfigLogTypeInput interface {
	pulumi.Input

	ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput
	ToAuditLogConfigLogTypeOutputWithContext(context.Context) AuditLogConfigLogTypeOutput
}

var auditLogConfigLogTypePtrType = reflect.TypeOf((**AuditLogConfigLogType)(nil)).Elem()

type AuditLogConfigLogTypePtrInput interface {
	pulumi.Input

	ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput
	ToAuditLogConfigLogTypePtrOutputWithContext(context.Context) AuditLogConfigLogTypePtrOutput
}

type auditLogConfigLogTypePtr string

func AuditLogConfigLogTypePtr(v string) AuditLogConfigLogTypePtrInput {
	return (*auditLogConfigLogTypePtr)(&v)
}

func (*auditLogConfigLogTypePtr) ElementType() reflect.Type {
	return auditLogConfigLogTypePtrType
}

func (in *auditLogConfigLogTypePtr) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return pulumi.ToOutput(in).(AuditLogConfigLogTypePtrOutput)
}

func (in *auditLogConfigLogTypePtr) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuditLogConfigLogTypePtrOutput)
}

// Required. Instance type.
type InstanceType string

const (
	// No type specified. The instance creation will fail.
	InstanceTypeTypeUnspecified = InstanceType("TYPE_UNSPECIFIED")
	// Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines using point and click UI. However, there are certain limitations, such as fewer number of concurrent pipelines, no support for streaming pipelines, etc.
	InstanceTypeBasic = InstanceType("BASIC")
	// Enterprise Data Fusion instance. In Enterprise type, the user will have all features available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
	InstanceTypeEnterprise = InstanceType("ENTERPRISE")
	// Developer Data Fusion instance. In Developer type, the user will have all features available but with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration pipelines at low cost.
	InstanceTypeDeveloper = InstanceType("DEVELOPER")
)

func (InstanceType) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceType)(nil)).Elem()
}

func (e InstanceType) ToInstanceTypeOutput() InstanceTypeOutput {
	return pulumi.ToOutput(e).(InstanceTypeOutput)
}

func (e InstanceType) ToInstanceTypeOutputWithContext(ctx context.Context) InstanceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InstanceTypeOutput)
}

func (e InstanceType) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return e.ToInstanceTypePtrOutputWithContext(context.Background())
}

func (e InstanceType) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return InstanceType(e).ToInstanceTypeOutputWithContext(ctx).ToInstanceTypePtrOutputWithContext(ctx)
}

func (e InstanceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InstanceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InstanceTypeOutput struct{ *pulumi.OutputState }

func (InstanceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceType)(nil)).Elem()
}

func (o InstanceTypeOutput) ToInstanceTypeOutput() InstanceTypeOutput {
	return o
}

func (o InstanceTypeOutput) ToInstanceTypeOutputWithContext(ctx context.Context) InstanceTypeOutput {
	return o
}

func (o InstanceTypeOutput) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return o.ToInstanceTypePtrOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceType) *InstanceType {
		return &v
	}).(InstanceTypePtrOutput)
}

func (o InstanceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InstanceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InstanceTypePtrOutput struct{ *pulumi.OutputState }

func (InstanceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceType)(nil)).Elem()
}

func (o InstanceTypePtrOutput) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return o
}

func (o InstanceTypePtrOutput) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return o
}

func (o InstanceTypePtrOutput) Elem() InstanceTypeOutput {
	return o.ApplyT(func(v *InstanceType) InstanceType {
		if v != nil {
			return *v
		}
		var ret InstanceType
		return ret
	}).(InstanceTypeOutput)
}

func (o InstanceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InstanceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InstanceTypeInput is an input type that accepts InstanceTypeArgs and InstanceTypeOutput values.
// You can construct a concrete instance of `InstanceTypeInput` via:
//
//          InstanceTypeArgs{...}
type InstanceTypeInput interface {
	pulumi.Input

	ToInstanceTypeOutput() InstanceTypeOutput
	ToInstanceTypeOutputWithContext(context.Context) InstanceTypeOutput
}

var instanceTypePtrType = reflect.TypeOf((**InstanceType)(nil)).Elem()

type InstanceTypePtrInput interface {
	pulumi.Input

	ToInstanceTypePtrOutput() InstanceTypePtrOutput
	ToInstanceTypePtrOutputWithContext(context.Context) InstanceTypePtrOutput
}

type instanceTypePtr string

func InstanceTypePtr(v string) InstanceTypePtrInput {
	return (*instanceTypePtr)(&v)
}

func (*instanceTypePtr) ElementType() reflect.Type {
	return instanceTypePtrType
}

func (in *instanceTypePtr) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return pulumi.ToOutput(in).(InstanceTypePtrOutput)
}

func (in *instanceTypePtr) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InstanceTypePtrOutput)
}

// Type represents the release availability of the version
type VersionType string

const (
	// Version does not have availability yet
	VersionTypeTypeUnspecified = VersionType("TYPE_UNSPECIFIED")
	// Version is under development and not considered stable
	VersionTypeTypePreview = VersionType("TYPE_PREVIEW")
	// Version is available for public use
	VersionTypeTypeGeneralAvailability = VersionType("TYPE_GENERAL_AVAILABILITY")
)

func (VersionType) ElementType() reflect.Type {
	return reflect.TypeOf((*VersionType)(nil)).Elem()
}

func (e VersionType) ToVersionTypeOutput() VersionTypeOutput {
	return pulumi.ToOutput(e).(VersionTypeOutput)
}

func (e VersionType) ToVersionTypeOutputWithContext(ctx context.Context) VersionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VersionTypeOutput)
}

func (e VersionType) ToVersionTypePtrOutput() VersionTypePtrOutput {
	return e.ToVersionTypePtrOutputWithContext(context.Background())
}

func (e VersionType) ToVersionTypePtrOutputWithContext(ctx context.Context) VersionTypePtrOutput {
	return VersionType(e).ToVersionTypeOutputWithContext(ctx).ToVersionTypePtrOutputWithContext(ctx)
}

func (e VersionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VersionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VersionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VersionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VersionTypeOutput struct{ *pulumi.OutputState }

func (VersionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VersionType)(nil)).Elem()
}

func (o VersionTypeOutput) ToVersionTypeOutput() VersionTypeOutput {
	return o
}

func (o VersionTypeOutput) ToVersionTypeOutputWithContext(ctx context.Context) VersionTypeOutput {
	return o
}

func (o VersionTypeOutput) ToVersionTypePtrOutput() VersionTypePtrOutput {
	return o.ToVersionTypePtrOutputWithContext(context.Background())
}

func (o VersionTypeOutput) ToVersionTypePtrOutputWithContext(ctx context.Context) VersionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VersionType) *VersionType {
		return &v
	}).(VersionTypePtrOutput)
}

func (o VersionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VersionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VersionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VersionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VersionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VersionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VersionTypePtrOutput struct{ *pulumi.OutputState }

func (VersionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VersionType)(nil)).Elem()
}

func (o VersionTypePtrOutput) ToVersionTypePtrOutput() VersionTypePtrOutput {
	return o
}

func (o VersionTypePtrOutput) ToVersionTypePtrOutputWithContext(ctx context.Context) VersionTypePtrOutput {
	return o
}

func (o VersionTypePtrOutput) Elem() VersionTypeOutput {
	return o.ApplyT(func(v *VersionType) VersionType {
		if v != nil {
			return *v
		}
		var ret VersionType
		return ret
	}).(VersionTypeOutput)
}

func (o VersionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VersionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VersionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// VersionTypeInput is an input type that accepts VersionTypeArgs and VersionTypeOutput values.
// You can construct a concrete instance of `VersionTypeInput` via:
//
//          VersionTypeArgs{...}
type VersionTypeInput interface {
	pulumi.Input

	ToVersionTypeOutput() VersionTypeOutput
	ToVersionTypeOutputWithContext(context.Context) VersionTypeOutput
}

var versionTypePtrType = reflect.TypeOf((**VersionType)(nil)).Elem()

type VersionTypePtrInput interface {
	pulumi.Input

	ToVersionTypePtrOutput() VersionTypePtrOutput
	ToVersionTypePtrOutputWithContext(context.Context) VersionTypePtrOutput
}

type versionTypePtr string

func VersionTypePtr(v string) VersionTypePtrInput {
	return (*versionTypePtr)(&v)
}

func (*versionTypePtr) ElementType() reflect.Type {
	return versionTypePtrType
}

func (in *versionTypePtr) ToVersionTypePtrOutput() VersionTypePtrOutput {
	return pulumi.ToOutput(in).(VersionTypePtrOutput)
}

func (in *versionTypePtr) ToVersionTypePtrOutputWithContext(ctx context.Context) VersionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VersionTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AcceleratorAcceleratorTypeInput)(nil)).Elem(), AcceleratorAcceleratorType("ACCELERATOR_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AcceleratorAcceleratorTypePtrInput)(nil)).Elem(), AcceleratorAcceleratorType("ACCELERATOR_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AcceleratorStateInput)(nil)).Elem(), AcceleratorState("STATE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AcceleratorStatePtrInput)(nil)).Elem(), AcceleratorState("STATE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypeInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypePtrInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTypeInput)(nil)).Elem(), InstanceType("TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTypePtrInput)(nil)).Elem(), InstanceType("TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*VersionTypeInput)(nil)).Elem(), VersionType("TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*VersionTypePtrInput)(nil)).Elem(), VersionType("TYPE_UNSPECIFIED"))
	pulumi.RegisterOutputType(AcceleratorAcceleratorTypeOutput{})
	pulumi.RegisterOutputType(AcceleratorAcceleratorTypePtrOutput{})
	pulumi.RegisterOutputType(AcceleratorStateOutput{})
	pulumi.RegisterOutputType(AcceleratorStatePtrOutput{})
	pulumi.RegisterOutputType(AuditLogConfigLogTypeOutput{})
	pulumi.RegisterOutputType(AuditLogConfigLogTypePtrOutput{})
	pulumi.RegisterOutputType(InstanceTypeOutput{})
	pulumi.RegisterOutputType(InstanceTypePtrOutput{})
	pulumi.RegisterOutputType(VersionTypeOutput{})
	pulumi.RegisterOutputType(VersionTypePtrOutput{})
}
