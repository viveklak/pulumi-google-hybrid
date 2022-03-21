// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The call logging level associated to this execution.
type ExecutionCallLogLevel string

const (
	// No call logging specified.
	ExecutionCallLogLevelCallLogLevelUnspecified = ExecutionCallLogLevel("CALL_LOG_LEVEL_UNSPECIFIED")
	// Log all call steps within workflows, all call returns, and all exceptions raised.
	ExecutionCallLogLevelLogAllCalls = ExecutionCallLogLevel("LOG_ALL_CALLS")
	// Log only exceptions that are raised from call steps within workflows.
	ExecutionCallLogLevelLogErrorsOnly = ExecutionCallLogLevel("LOG_ERRORS_ONLY")
)

func (ExecutionCallLogLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*ExecutionCallLogLevel)(nil)).Elem()
}

func (e ExecutionCallLogLevel) ToExecutionCallLogLevelOutput() ExecutionCallLogLevelOutput {
	return pulumi.ToOutput(e).(ExecutionCallLogLevelOutput)
}

func (e ExecutionCallLogLevel) ToExecutionCallLogLevelOutputWithContext(ctx context.Context) ExecutionCallLogLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExecutionCallLogLevelOutput)
}

func (e ExecutionCallLogLevel) ToExecutionCallLogLevelPtrOutput() ExecutionCallLogLevelPtrOutput {
	return e.ToExecutionCallLogLevelPtrOutputWithContext(context.Background())
}

func (e ExecutionCallLogLevel) ToExecutionCallLogLevelPtrOutputWithContext(ctx context.Context) ExecutionCallLogLevelPtrOutput {
	return ExecutionCallLogLevel(e).ToExecutionCallLogLevelOutputWithContext(ctx).ToExecutionCallLogLevelPtrOutputWithContext(ctx)
}

func (e ExecutionCallLogLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExecutionCallLogLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExecutionCallLogLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExecutionCallLogLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExecutionCallLogLevelOutput struct{ *pulumi.OutputState }

func (ExecutionCallLogLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExecutionCallLogLevel)(nil)).Elem()
}

func (o ExecutionCallLogLevelOutput) ToExecutionCallLogLevelOutput() ExecutionCallLogLevelOutput {
	return o
}

func (o ExecutionCallLogLevelOutput) ToExecutionCallLogLevelOutputWithContext(ctx context.Context) ExecutionCallLogLevelOutput {
	return o
}

func (o ExecutionCallLogLevelOutput) ToExecutionCallLogLevelPtrOutput() ExecutionCallLogLevelPtrOutput {
	return o.ToExecutionCallLogLevelPtrOutputWithContext(context.Background())
}

func (o ExecutionCallLogLevelOutput) ToExecutionCallLogLevelPtrOutputWithContext(ctx context.Context) ExecutionCallLogLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExecutionCallLogLevel) *ExecutionCallLogLevel {
		return &v
	}).(ExecutionCallLogLevelPtrOutput)
}

func (o ExecutionCallLogLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExecutionCallLogLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExecutionCallLogLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExecutionCallLogLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExecutionCallLogLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExecutionCallLogLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExecutionCallLogLevelPtrOutput struct{ *pulumi.OutputState }

func (ExecutionCallLogLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExecutionCallLogLevel)(nil)).Elem()
}

func (o ExecutionCallLogLevelPtrOutput) ToExecutionCallLogLevelPtrOutput() ExecutionCallLogLevelPtrOutput {
	return o
}

func (o ExecutionCallLogLevelPtrOutput) ToExecutionCallLogLevelPtrOutputWithContext(ctx context.Context) ExecutionCallLogLevelPtrOutput {
	return o
}

func (o ExecutionCallLogLevelPtrOutput) Elem() ExecutionCallLogLevelOutput {
	return o.ApplyT(func(v *ExecutionCallLogLevel) ExecutionCallLogLevel {
		if v != nil {
			return *v
		}
		var ret ExecutionCallLogLevel
		return ret
	}).(ExecutionCallLogLevelOutput)
}

func (o ExecutionCallLogLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExecutionCallLogLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExecutionCallLogLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ExecutionCallLogLevelInput is an input type that accepts ExecutionCallLogLevelArgs and ExecutionCallLogLevelOutput values.
// You can construct a concrete instance of `ExecutionCallLogLevelInput` via:
//
//          ExecutionCallLogLevelArgs{...}
type ExecutionCallLogLevelInput interface {
	pulumi.Input

	ToExecutionCallLogLevelOutput() ExecutionCallLogLevelOutput
	ToExecutionCallLogLevelOutputWithContext(context.Context) ExecutionCallLogLevelOutput
}

var executionCallLogLevelPtrType = reflect.TypeOf((**ExecutionCallLogLevel)(nil)).Elem()

type ExecutionCallLogLevelPtrInput interface {
	pulumi.Input

	ToExecutionCallLogLevelPtrOutput() ExecutionCallLogLevelPtrOutput
	ToExecutionCallLogLevelPtrOutputWithContext(context.Context) ExecutionCallLogLevelPtrOutput
}

type executionCallLogLevelPtr string

func ExecutionCallLogLevelPtr(v string) ExecutionCallLogLevelPtrInput {
	return (*executionCallLogLevelPtr)(&v)
}

func (*executionCallLogLevelPtr) ElementType() reflect.Type {
	return executionCallLogLevelPtrType
}

func (in *executionCallLogLevelPtr) ToExecutionCallLogLevelPtrOutput() ExecutionCallLogLevelPtrOutput {
	return pulumi.ToOutput(in).(ExecutionCallLogLevelPtrOutput)
}

func (in *executionCallLogLevelPtr) ToExecutionCallLogLevelPtrOutputWithContext(ctx context.Context) ExecutionCallLogLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExecutionCallLogLevelPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionCallLogLevelInput)(nil)).Elem(), ExecutionCallLogLevel("CALL_LOG_LEVEL_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionCallLogLevelPtrInput)(nil)).Elem(), ExecutionCallLogLevel("CALL_LOG_LEVEL_UNSPECIFIED"))
	pulumi.RegisterOutputType(ExecutionCallLogLevelOutput{})
	pulumi.RegisterOutputType(ExecutionCallLogLevelPtrOutput{})
}
