// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified HealthCheck resource. Gets a list of available health checks by making a list() request.
func LookupHealthCheck(ctx *pulumi.Context, args *LookupHealthCheckArgs, opts ...pulumi.InvokeOption) (*LookupHealthCheckResult, error) {
	var rv LookupHealthCheckResult
	err := ctx.Invoke("google-native:compute/v1:getHealthCheck", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHealthCheckArgs struct {
	HealthCheck string  `pulumi:"healthCheck"`
	Project     *string `pulumi:"project"`
}

type LookupHealthCheckResult struct {
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec int `pulumi:"checkIntervalSec"`
	// Creation timestamp in 3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description     string                  `pulumi:"description"`
	GrpcHealthCheck GRPCHealthCheckResponse `pulumi:"grpcHealthCheck"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold int                      `pulumi:"healthyThreshold"`
	Http2HealthCheck HTTP2HealthCheckResponse `pulumi:"http2HealthCheck"`
	HttpHealthCheck  HTTPHealthCheckResponse  `pulumi:"httpHealthCheck"`
	HttpsHealthCheck HTTPSHealthCheckResponse `pulumi:"httpsHealthCheck"`
	// Type of the resource.
	Kind string `pulumi:"kind"`
	// Configure logging on this health check.
	LogConfig HealthCheckLogConfigResponse `pulumi:"logConfig"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. For example, a name that is 1-63 characters long, matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`, and otherwise complies with RFC1035. This regular expression describes a name where the first character is a lowercase letter, and all following characters are a dash, lowercase letter, or digit, except the last character, which isn't a dash.
	Name string `pulumi:"name"`
	// Region where the health check resides. Not applicable to global health checks.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink       string                 `pulumi:"selfLink"`
	SslHealthCheck SSLHealthCheckResponse `pulumi:"sslHealthCheck"`
	TcpHealthCheck TCPHealthCheckResponse `pulumi:"tcpHealthCheck"`
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
	TimeoutSec int `pulumi:"timeoutSec"`
	// Specifies the type of the healthCheck, either TCP, SSL, HTTP, HTTPS or HTTP2. Exactly one of the protocol-specific health check field must be specified, which must match type field.
	Type string `pulumi:"type"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold int `pulumi:"unhealthyThreshold"`
}

func LookupHealthCheckOutput(ctx *pulumi.Context, args LookupHealthCheckOutputArgs, opts ...pulumi.InvokeOption) LookupHealthCheckResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHealthCheckResult, error) {
			args := v.(LookupHealthCheckArgs)
			r, err := LookupHealthCheck(ctx, &args, opts...)
			return *r, err
		}).(LookupHealthCheckResultOutput)
}

type LookupHealthCheckOutputArgs struct {
	HealthCheck pulumi.StringInput    `pulumi:"healthCheck"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupHealthCheckOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHealthCheckArgs)(nil)).Elem()
}

type LookupHealthCheckResultOutput struct{ *pulumi.OutputState }

func (LookupHealthCheckResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHealthCheckResult)(nil)).Elem()
}

func (o LookupHealthCheckResultOutput) ToLookupHealthCheckResultOutput() LookupHealthCheckResultOutput {
	return o
}

func (o LookupHealthCheckResultOutput) ToLookupHealthCheckResultOutputWithContext(ctx context.Context) LookupHealthCheckResultOutput {
	return o
}

// How often (in seconds) to send a health check. The default value is 5 seconds.
func (o LookupHealthCheckResultOutput) CheckIntervalSec() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) int { return v.CheckIntervalSec }).(pulumi.IntOutput)
}

// Creation timestamp in 3339 text format.
func (o LookupHealthCheckResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupHealthCheckResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupHealthCheckResultOutput) GrpcHealthCheck() GRPCHealthCheckResponseOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) GRPCHealthCheckResponse { return v.GrpcHealthCheck }).(GRPCHealthCheckResponseOutput)
}

// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
func (o LookupHealthCheckResultOutput) HealthyThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) int { return v.HealthyThreshold }).(pulumi.IntOutput)
}

func (o LookupHealthCheckResultOutput) Http2HealthCheck() HTTP2HealthCheckResponseOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) HTTP2HealthCheckResponse { return v.Http2HealthCheck }).(HTTP2HealthCheckResponseOutput)
}

func (o LookupHealthCheckResultOutput) HttpHealthCheck() HTTPHealthCheckResponseOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) HTTPHealthCheckResponse { return v.HttpHealthCheck }).(HTTPHealthCheckResponseOutput)
}

func (o LookupHealthCheckResultOutput) HttpsHealthCheck() HTTPSHealthCheckResponseOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) HTTPSHealthCheckResponse { return v.HttpsHealthCheck }).(HTTPSHealthCheckResponseOutput)
}

// Type of the resource.
func (o LookupHealthCheckResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Configure logging on this health check.
func (o LookupHealthCheckResultOutput) LogConfig() HealthCheckLogConfigResponseOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) HealthCheckLogConfigResponse { return v.LogConfig }).(HealthCheckLogConfigResponseOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. For example, a name that is 1-63 characters long, matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`, and otherwise complies with RFC1035. This regular expression describes a name where the first character is a lowercase letter, and all following characters are a dash, lowercase letter, or digit, except the last character, which isn't a dash.
func (o LookupHealthCheckResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) string { return v.Name }).(pulumi.StringOutput)
}

// Region where the health check resides. Not applicable to global health checks.
func (o LookupHealthCheckResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupHealthCheckResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupHealthCheckResultOutput) SslHealthCheck() SSLHealthCheckResponseOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) SSLHealthCheckResponse { return v.SslHealthCheck }).(SSLHealthCheckResponseOutput)
}

func (o LookupHealthCheckResultOutput) TcpHealthCheck() TCPHealthCheckResponseOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) TCPHealthCheckResponse { return v.TcpHealthCheck }).(TCPHealthCheckResponseOutput)
}

// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
func (o LookupHealthCheckResultOutput) TimeoutSec() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) int { return v.TimeoutSec }).(pulumi.IntOutput)
}

// Specifies the type of the healthCheck, either TCP, SSL, HTTP, HTTPS or HTTP2. Exactly one of the protocol-specific health check field must be specified, which must match type field.
func (o LookupHealthCheckResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) string { return v.Type }).(pulumi.StringOutput)
}

// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
func (o LookupHealthCheckResultOutput) UnhealthyThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) int { return v.UnhealthyThreshold }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHealthCheckResultOutput{})
}
