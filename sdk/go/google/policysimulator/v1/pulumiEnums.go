// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The logs to use as input for the Replay.
type GoogleCloudPolicysimulatorV1ReplayConfigLogSource string

const (
	// An unspecified log source. If the log source is unspecified, the Replay defaults to using `RECENT_ACCESSES`.
	GoogleCloudPolicysimulatorV1ReplayConfigLogSourceLogSourceUnspecified = GoogleCloudPolicysimulatorV1ReplayConfigLogSource("LOG_SOURCE_UNSPECIFIED")
	// All access logs from the last 90 days. These logs may not include logs from the most recent 7 days.
	GoogleCloudPolicysimulatorV1ReplayConfigLogSourceRecentAccesses = GoogleCloudPolicysimulatorV1ReplayConfigLogSource("RECENT_ACCESSES")
)

func (GoogleCloudPolicysimulatorV1ReplayConfigLogSource) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudPolicysimulatorV1ReplayConfigLogSource)(nil)).Elem()
}

func (e GoogleCloudPolicysimulatorV1ReplayConfigLogSource) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput() GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput {
	return pulumi.ToOutput(e).(GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput)
}

func (e GoogleCloudPolicysimulatorV1ReplayConfigLogSource) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutputWithContext(ctx context.Context) GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput)
}

func (e GoogleCloudPolicysimulatorV1ReplayConfigLogSource) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput() GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput {
	return e.ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutputWithContext(context.Background())
}

func (e GoogleCloudPolicysimulatorV1ReplayConfigLogSource) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutputWithContext(ctx context.Context) GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput {
	return GoogleCloudPolicysimulatorV1ReplayConfigLogSource(e).ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutputWithContext(ctx).ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutputWithContext(ctx)
}

func (e GoogleCloudPolicysimulatorV1ReplayConfigLogSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudPolicysimulatorV1ReplayConfigLogSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudPolicysimulatorV1ReplayConfigLogSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudPolicysimulatorV1ReplayConfigLogSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput struct{ *pulumi.OutputState }

func (GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudPolicysimulatorV1ReplayConfigLogSource)(nil)).Elem()
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput() GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput {
	return o
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutputWithContext(ctx context.Context) GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput {
	return o
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput() GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput {
	return o.ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutputWithContext(context.Background())
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutputWithContext(ctx context.Context) GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GoogleCloudPolicysimulatorV1ReplayConfigLogSource) *GoogleCloudPolicysimulatorV1ReplayConfigLogSource {
		return &v
	}).(GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput)
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleCloudPolicysimulatorV1ReplayConfigLogSource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleCloudPolicysimulatorV1ReplayConfigLogSource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput struct{ *pulumi.OutputState }

func (GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleCloudPolicysimulatorV1ReplayConfigLogSource)(nil)).Elem()
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput() GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput {
	return o
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutputWithContext(ctx context.Context) GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput {
	return o
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput) Elem() GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput {
	return o.ApplyT(func(v *GoogleCloudPolicysimulatorV1ReplayConfigLogSource) GoogleCloudPolicysimulatorV1ReplayConfigLogSource {
		if v != nil {
			return *v
		}
		var ret GoogleCloudPolicysimulatorV1ReplayConfigLogSource
		return ret
	}).(GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput)
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GoogleCloudPolicysimulatorV1ReplayConfigLogSource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// GoogleCloudPolicysimulatorV1ReplayConfigLogSourceInput is an input type that accepts GoogleCloudPolicysimulatorV1ReplayConfigLogSourceArgs and GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput values.
// You can construct a concrete instance of `GoogleCloudPolicysimulatorV1ReplayConfigLogSourceInput` via:
//
//          GoogleCloudPolicysimulatorV1ReplayConfigLogSourceArgs{...}
type GoogleCloudPolicysimulatorV1ReplayConfigLogSourceInput interface {
	pulumi.Input

	ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput() GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput
	ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutputWithContext(context.Context) GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput
}

var googleCloudPolicysimulatorV1ReplayConfigLogSourcePtrType = reflect.TypeOf((**GoogleCloudPolicysimulatorV1ReplayConfigLogSource)(nil)).Elem()

type GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrInput interface {
	pulumi.Input

	ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput() GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput
	ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutputWithContext(context.Context) GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput
}

type googleCloudPolicysimulatorV1ReplayConfigLogSourcePtr string

func GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtr(v string) GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrInput {
	return (*googleCloudPolicysimulatorV1ReplayConfigLogSourcePtr)(&v)
}

func (*googleCloudPolicysimulatorV1ReplayConfigLogSourcePtr) ElementType() reflect.Type {
	return googleCloudPolicysimulatorV1ReplayConfigLogSourcePtrType
}

func (in *googleCloudPolicysimulatorV1ReplayConfigLogSourcePtr) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput() GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput {
	return pulumi.ToOutput(in).(GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput)
}

func (in *googleCloudPolicysimulatorV1ReplayConfigLogSourcePtr) ToGoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutputWithContext(ctx context.Context) GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleCloudPolicysimulatorV1ReplayConfigLogSourceInput)(nil)).Elem(), GoogleCloudPolicysimulatorV1ReplayConfigLogSource("LOG_SOURCE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrInput)(nil)).Elem(), GoogleCloudPolicysimulatorV1ReplayConfigLogSource("LOG_SOURCE_UNSPECIFIED"))
	pulumi.RegisterOutputType(GoogleCloudPolicysimulatorV1ReplayConfigLogSourceOutput{})
	pulumi.RegisterOutputType(GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput{})
}
