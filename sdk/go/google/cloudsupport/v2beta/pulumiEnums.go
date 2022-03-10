// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The priority of this case. If this is set, do not set severity.
type CasePriority string

const (
	// Severity is undefined or has not been set yet.
	CasePriorityPriorityUnspecified = CasePriority("PRIORITY_UNSPECIFIED")
	// Extreme impact on a production service. Service is hard down.
	CasePriorityP0 = CasePriority("P0")
	// Critical impact on a production service. Service is currently unusable.
	CasePriorityP1 = CasePriority("P1")
	// Severe impact on a production service. Service is usable but greatly impaired.
	CasePriorityP2 = CasePriority("P2")
	// Medium impact on a production service. Service is available, but moderately impaired.
	CasePriorityP3 = CasePriority("P3")
	// General questions or minor issues. Production service is fully available.
	CasePriorityP4 = CasePriority("P4")
)

func (CasePriority) ElementType() reflect.Type {
	return reflect.TypeOf((*CasePriority)(nil)).Elem()
}

func (e CasePriority) ToCasePriorityOutput() CasePriorityOutput {
	return pulumi.ToOutput(e).(CasePriorityOutput)
}

func (e CasePriority) ToCasePriorityOutputWithContext(ctx context.Context) CasePriorityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CasePriorityOutput)
}

func (e CasePriority) ToCasePriorityPtrOutput() CasePriorityPtrOutput {
	return e.ToCasePriorityPtrOutputWithContext(context.Background())
}

func (e CasePriority) ToCasePriorityPtrOutputWithContext(ctx context.Context) CasePriorityPtrOutput {
	return CasePriority(e).ToCasePriorityOutputWithContext(ctx).ToCasePriorityPtrOutputWithContext(ctx)
}

func (e CasePriority) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CasePriority) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CasePriority) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CasePriority) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CasePriorityOutput struct{ *pulumi.OutputState }

func (CasePriorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CasePriority)(nil)).Elem()
}

func (o CasePriorityOutput) ToCasePriorityOutput() CasePriorityOutput {
	return o
}

func (o CasePriorityOutput) ToCasePriorityOutputWithContext(ctx context.Context) CasePriorityOutput {
	return o
}

func (o CasePriorityOutput) ToCasePriorityPtrOutput() CasePriorityPtrOutput {
	return o.ToCasePriorityPtrOutputWithContext(context.Background())
}

func (o CasePriorityOutput) ToCasePriorityPtrOutputWithContext(ctx context.Context) CasePriorityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CasePriority) *CasePriority {
		return &v
	}).(CasePriorityPtrOutput)
}

func (o CasePriorityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CasePriorityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CasePriority) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CasePriorityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CasePriorityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CasePriority) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CasePriorityPtrOutput struct{ *pulumi.OutputState }

func (CasePriorityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CasePriority)(nil)).Elem()
}

func (o CasePriorityPtrOutput) ToCasePriorityPtrOutput() CasePriorityPtrOutput {
	return o
}

func (o CasePriorityPtrOutput) ToCasePriorityPtrOutputWithContext(ctx context.Context) CasePriorityPtrOutput {
	return o
}

func (o CasePriorityPtrOutput) Elem() CasePriorityOutput {
	return o.ApplyT(func(v *CasePriority) CasePriority {
		if v != nil {
			return *v
		}
		var ret CasePriority
		return ret
	}).(CasePriorityOutput)
}

func (o CasePriorityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CasePriorityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CasePriority) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CasePriorityInput is an input type that accepts CasePriorityArgs and CasePriorityOutput values.
// You can construct a concrete instance of `CasePriorityInput` via:
//
//          CasePriorityArgs{...}
type CasePriorityInput interface {
	pulumi.Input

	ToCasePriorityOutput() CasePriorityOutput
	ToCasePriorityOutputWithContext(context.Context) CasePriorityOutput
}

var casePriorityPtrType = reflect.TypeOf((**CasePriority)(nil)).Elem()

type CasePriorityPtrInput interface {
	pulumi.Input

	ToCasePriorityPtrOutput() CasePriorityPtrOutput
	ToCasePriorityPtrOutputWithContext(context.Context) CasePriorityPtrOutput
}

type casePriorityPtr string

func CasePriorityPtr(v string) CasePriorityPtrInput {
	return (*casePriorityPtr)(&v)
}

func (*casePriorityPtr) ElementType() reflect.Type {
	return casePriorityPtrType
}

func (in *casePriorityPtr) ToCasePriorityPtrOutput() CasePriorityPtrOutput {
	return pulumi.ToOutput(in).(CasePriorityPtrOutput)
}

func (in *casePriorityPtr) ToCasePriorityPtrOutputWithContext(ctx context.Context) CasePriorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CasePriorityPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CasePriorityInput)(nil)).Elem(), CasePriority("PRIORITY_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*CasePriorityPtrInput)(nil)).Elem(), CasePriority("PRIORITY_UNSPECIFIED"))
	pulumi.RegisterOutputType(CasePriorityOutput{})
	pulumi.RegisterOutputType(CasePriorityPtrOutput{})
}