// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Error describes why the execution was abnormally terminated.
type ErrorResponse struct {
	// Human-readable stack trace string.
	Context string `pulumi:"context"`
	// Error message and data returned represented as a JSON string.
	Payload string `pulumi:"payload"`
	// Stack trace with detailed information of where error was generated.
	StackTrace StackTraceResponse `pulumi:"stackTrace"`
}

// Error describes why the execution was abnormally terminated.
type ErrorResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponse)(nil)).Elem()
}

func (o ErrorResponseOutput) ToErrorResponseOutput() ErrorResponseOutput {
	return o
}

func (o ErrorResponseOutput) ToErrorResponseOutputWithContext(ctx context.Context) ErrorResponseOutput {
	return o
}

// Human-readable stack trace string.
func (o ErrorResponseOutput) Context() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponse) string { return v.Context }).(pulumi.StringOutput)
}

// Error message and data returned represented as a JSON string.
func (o ErrorResponseOutput) Payload() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponse) string { return v.Payload }).(pulumi.StringOutput)
}

// Stack trace with detailed information of where error was generated.
func (o ErrorResponseOutput) StackTrace() StackTraceResponseOutput {
	return o.ApplyT(func(v ErrorResponse) StackTraceResponse { return v.StackTrace }).(StackTraceResponseOutput)
}

// Position contains source position information about the stack trace element such as line number, column number and length of the code block in bytes.
type PositionResponse struct {
	// The source code column position (of the line) the current instruction was generated from.
	Column string `pulumi:"column"`
	// The number of bytes of source code making up this stack trace element.
	Length string `pulumi:"length"`
	// The source code line number the current instruction was generated from.
	Line string `pulumi:"line"`
}

// Position contains source position information about the stack trace element such as line number, column number and length of the code block in bytes.
type PositionResponseOutput struct{ *pulumi.OutputState }

func (PositionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PositionResponse)(nil)).Elem()
}

func (o PositionResponseOutput) ToPositionResponseOutput() PositionResponseOutput {
	return o
}

func (o PositionResponseOutput) ToPositionResponseOutputWithContext(ctx context.Context) PositionResponseOutput {
	return o
}

// The source code column position (of the line) the current instruction was generated from.
func (o PositionResponseOutput) Column() pulumi.StringOutput {
	return o.ApplyT(func(v PositionResponse) string { return v.Column }).(pulumi.StringOutput)
}

// The number of bytes of source code making up this stack trace element.
func (o PositionResponseOutput) Length() pulumi.StringOutput {
	return o.ApplyT(func(v PositionResponse) string { return v.Length }).(pulumi.StringOutput)
}

// The source code line number the current instruction was generated from.
func (o PositionResponseOutput) Line() pulumi.StringOutput {
	return o.ApplyT(func(v PositionResponse) string { return v.Line }).(pulumi.StringOutput)
}

// A single stack element (frame) where an error occurred.
type StackTraceElementResponse struct {
	// The source position information of the stack trace element.
	Position PositionResponse `pulumi:"position"`
	// The routine where the error occurred.
	Routine string `pulumi:"routine"`
	// The step the error occurred at.
	Step string `pulumi:"step"`
}

// A single stack element (frame) where an error occurred.
type StackTraceElementResponseOutput struct{ *pulumi.OutputState }

func (StackTraceElementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackTraceElementResponse)(nil)).Elem()
}

func (o StackTraceElementResponseOutput) ToStackTraceElementResponseOutput() StackTraceElementResponseOutput {
	return o
}

func (o StackTraceElementResponseOutput) ToStackTraceElementResponseOutputWithContext(ctx context.Context) StackTraceElementResponseOutput {
	return o
}

// The source position information of the stack trace element.
func (o StackTraceElementResponseOutput) Position() PositionResponseOutput {
	return o.ApplyT(func(v StackTraceElementResponse) PositionResponse { return v.Position }).(PositionResponseOutput)
}

// The routine where the error occurred.
func (o StackTraceElementResponseOutput) Routine() pulumi.StringOutput {
	return o.ApplyT(func(v StackTraceElementResponse) string { return v.Routine }).(pulumi.StringOutput)
}

// The step the error occurred at.
func (o StackTraceElementResponseOutput) Step() pulumi.StringOutput {
	return o.ApplyT(func(v StackTraceElementResponse) string { return v.Step }).(pulumi.StringOutput)
}

type StackTraceElementResponseArrayOutput struct{ *pulumi.OutputState }

func (StackTraceElementResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StackTraceElementResponse)(nil)).Elem()
}

func (o StackTraceElementResponseArrayOutput) ToStackTraceElementResponseArrayOutput() StackTraceElementResponseArrayOutput {
	return o
}

func (o StackTraceElementResponseArrayOutput) ToStackTraceElementResponseArrayOutputWithContext(ctx context.Context) StackTraceElementResponseArrayOutput {
	return o
}

func (o StackTraceElementResponseArrayOutput) Index(i pulumi.IntInput) StackTraceElementResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StackTraceElementResponse {
		return vs[0].([]StackTraceElementResponse)[vs[1].(int)]
	}).(StackTraceElementResponseOutput)
}

// A collection of stack elements (frames) where an error occurred.
type StackTraceResponse struct {
	// An array of stack elements.
	Elements []StackTraceElementResponse `pulumi:"elements"`
}

// A collection of stack elements (frames) where an error occurred.
type StackTraceResponseOutput struct{ *pulumi.OutputState }

func (StackTraceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackTraceResponse)(nil)).Elem()
}

func (o StackTraceResponseOutput) ToStackTraceResponseOutput() StackTraceResponseOutput {
	return o
}

func (o StackTraceResponseOutput) ToStackTraceResponseOutputWithContext(ctx context.Context) StackTraceResponseOutput {
	return o
}

// An array of stack elements.
func (o StackTraceResponseOutput) Elements() StackTraceElementResponseArrayOutput {
	return o.ApplyT(func(v StackTraceResponse) []StackTraceElementResponse { return v.Elements }).(StackTraceElementResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ErrorResponseOutput{})
	pulumi.RegisterOutputType(PositionResponseOutput{})
	pulumi.RegisterOutputType(StackTraceElementResponseOutput{})
	pulumi.RegisterOutputType(StackTraceElementResponseArrayOutput{})
	pulumi.RegisterOutputType(StackTraceResponseOutput{})
}
