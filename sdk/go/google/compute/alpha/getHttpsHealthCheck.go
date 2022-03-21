// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified HttpsHealthCheck resource. Gets a list of available HTTPS health checks by making a list() request.
func LookupHttpsHealthCheck(ctx *pulumi.Context, args *LookupHttpsHealthCheckArgs, opts ...pulumi.InvokeOption) (*LookupHttpsHealthCheckResult, error) {
	var rv LookupHttpsHealthCheckResult
	err := ctx.Invoke("google-native:compute/alpha:getHttpsHealthCheck", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHttpsHealthCheckArgs struct {
	HttpsHealthCheck string  `pulumi:"httpsHealthCheck"`
	Project          *string `pulumi:"project"`
}

type LookupHttpsHealthCheckResult struct {
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec int `pulumi:"checkIntervalSec"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold int `pulumi:"healthyThreshold"`
	// The value of the host header in the HTTPS health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
	Host string `pulumi:"host"`
	// Type of the resource.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The TCP port number for the HTTPS health check request. The default value is 443.
	Port int `pulumi:"port"`
	// The request path of the HTTPS health check request. The default value is "/".
	RequestPath string `pulumi:"requestPath"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have a greater value than checkIntervalSec.
	TimeoutSec int `pulumi:"timeoutSec"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold int `pulumi:"unhealthyThreshold"`
}

func LookupHttpsHealthCheckOutput(ctx *pulumi.Context, args LookupHttpsHealthCheckOutputArgs, opts ...pulumi.InvokeOption) LookupHttpsHealthCheckResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHttpsHealthCheckResult, error) {
			args := v.(LookupHttpsHealthCheckArgs)
			r, err := LookupHttpsHealthCheck(ctx, &args, opts...)
			return *r, err
		}).(LookupHttpsHealthCheckResultOutput)
}

type LookupHttpsHealthCheckOutputArgs struct {
	HttpsHealthCheck pulumi.StringInput    `pulumi:"httpsHealthCheck"`
	Project          pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupHttpsHealthCheckOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHttpsHealthCheckArgs)(nil)).Elem()
}

type LookupHttpsHealthCheckResultOutput struct{ *pulumi.OutputState }

func (LookupHttpsHealthCheckResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHttpsHealthCheckResult)(nil)).Elem()
}

func (o LookupHttpsHealthCheckResultOutput) ToLookupHttpsHealthCheckResultOutput() LookupHttpsHealthCheckResultOutput {
	return o
}

func (o LookupHttpsHealthCheckResultOutput) ToLookupHttpsHealthCheckResultOutputWithContext(ctx context.Context) LookupHttpsHealthCheckResultOutput {
	return o
}

// How often (in seconds) to send a health check. The default value is 5 seconds.
func (o LookupHttpsHealthCheckResultOutput) CheckIntervalSec() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) int { return v.CheckIntervalSec }).(pulumi.IntOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupHttpsHealthCheckResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupHttpsHealthCheckResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) string { return v.Description }).(pulumi.StringOutput)
}

// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
func (o LookupHttpsHealthCheckResultOutput) HealthyThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) int { return v.HealthyThreshold }).(pulumi.IntOutput)
}

// The value of the host header in the HTTPS health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
func (o LookupHttpsHealthCheckResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) string { return v.Host }).(pulumi.StringOutput)
}

// Type of the resource.
func (o LookupHttpsHealthCheckResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupHttpsHealthCheckResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) string { return v.Name }).(pulumi.StringOutput)
}

// The TCP port number for the HTTPS health check request. The default value is 443.
func (o LookupHttpsHealthCheckResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) int { return v.Port }).(pulumi.IntOutput)
}

// The request path of the HTTPS health check request. The default value is "/".
func (o LookupHttpsHealthCheckResultOutput) RequestPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) string { return v.RequestPath }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupHttpsHealthCheckResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupHttpsHealthCheckResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have a greater value than checkIntervalSec.
func (o LookupHttpsHealthCheckResultOutput) TimeoutSec() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) int { return v.TimeoutSec }).(pulumi.IntOutput)
}

// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
func (o LookupHttpsHealthCheckResultOutput) UnhealthyThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHttpsHealthCheckResult) int { return v.UnhealthyThreshold }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHttpsHealthCheckResultOutput{})
}
