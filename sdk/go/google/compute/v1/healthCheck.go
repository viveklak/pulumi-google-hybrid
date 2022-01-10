// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a HealthCheck resource in the specified project using the data included in the request.
type HealthCheck struct {
	pulumi.CustomResourceState

	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec pulumi.IntOutput `pulumi:"checkIntervalSec"`
	// Creation timestamp in 3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description     pulumi.StringOutput           `pulumi:"description"`
	GrpcHealthCheck GRPCHealthCheckResponseOutput `pulumi:"grpcHealthCheck"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntOutput               `pulumi:"healthyThreshold"`
	Http2HealthCheck HTTP2HealthCheckResponseOutput `pulumi:"http2HealthCheck"`
	HttpHealthCheck  HTTPHealthCheckResponseOutput  `pulumi:"httpHealthCheck"`
	HttpsHealthCheck HTTPSHealthCheckResponseOutput `pulumi:"httpsHealthCheck"`
	// Type of the resource.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Configure logging on this health check.
	LogConfig HealthCheckLogConfigResponseOutput `pulumi:"logConfig"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. For example, a name that is 1-63 characters long, matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`, and otherwise complies with RFC1035. This regular expression describes a name where the first character is a lowercase letter, and all following characters are a dash, lowercase letter, or digit, except the last character, which isn't a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Region where the health check resides. Not applicable to global health checks.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink       pulumi.StringOutput          `pulumi:"selfLink"`
	SslHealthCheck SSLHealthCheckResponseOutput `pulumi:"sslHealthCheck"`
	TcpHealthCheck TCPHealthCheckResponseOutput `pulumi:"tcpHealthCheck"`
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
	TimeoutSec pulumi.IntOutput `pulumi:"timeoutSec"`
	// Specifies the type of the healthCheck, either TCP, SSL, HTTP, HTTPS or HTTP2. Exactly one of the protocol-specific health check field must be specified, which must match type field.
	Type pulumi.StringOutput `pulumi:"type"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntOutput `pulumi:"unhealthyThreshold"`
}

// NewHealthCheck registers a new resource with the given unique name, arguments, and options.
func NewHealthCheck(ctx *pulumi.Context,
	name string, args *HealthCheckArgs, opts ...pulumi.ResourceOption) (*HealthCheck, error) {
	if args == nil {
		args = &HealthCheckArgs{}
	}

	var resource HealthCheck
	err := ctx.RegisterResource("google-native:compute/v1:HealthCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHealthCheck gets an existing HealthCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHealthCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HealthCheckState, opts ...pulumi.ResourceOption) (*HealthCheck, error) {
	var resource HealthCheck
	err := ctx.ReadResource("google-native:compute/v1:HealthCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HealthCheck resources.
type healthCheckState struct {
}

type HealthCheckState struct {
}

func (HealthCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckState)(nil)).Elem()
}

type healthCheckArgs struct {
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec *int `pulumi:"checkIntervalSec"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description     *string          `pulumi:"description"`
	GrpcHealthCheck *GRPCHealthCheck `pulumi:"grpcHealthCheck"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold *int                  `pulumi:"healthyThreshold"`
	Http2HealthCheck *HTTP2HealthCheck     `pulumi:"http2HealthCheck"`
	HttpHealthCheck  *HTTPHealthCheckType  `pulumi:"httpHealthCheck"`
	HttpsHealthCheck *HTTPSHealthCheckType `pulumi:"httpsHealthCheck"`
	// Type of the resource.
	Kind *string `pulumi:"kind"`
	// Configure logging on this health check.
	LogConfig *HealthCheckLogConfig `pulumi:"logConfig"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. For example, a name that is 1-63 characters long, matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`, and otherwise complies with RFC1035. This regular expression describes a name where the first character is a lowercase letter, and all following characters are a dash, lowercase letter, or digit, except the last character, which isn't a dash.
	Name           *string         `pulumi:"name"`
	Project        *string         `pulumi:"project"`
	RequestId      *string         `pulumi:"requestId"`
	SslHealthCheck *SSLHealthCheck `pulumi:"sslHealthCheck"`
	TcpHealthCheck *TCPHealthCheck `pulumi:"tcpHealthCheck"`
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
	TimeoutSec *int `pulumi:"timeoutSec"`
	// Specifies the type of the healthCheck, either TCP, SSL, HTTP, HTTPS or HTTP2. Exactly one of the protocol-specific health check field must be specified, which must match type field.
	Type *HealthCheckType `pulumi:"type"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

// The set of arguments for constructing a HealthCheck resource.
type HealthCheckArgs struct {
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec pulumi.IntPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description     pulumi.StringPtrInput
	GrpcHealthCheck GRPCHealthCheckPtrInput
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrInput
	Http2HealthCheck HTTP2HealthCheckPtrInput
	HttpHealthCheck  HTTPHealthCheckTypePtrInput
	HttpsHealthCheck HTTPSHealthCheckTypePtrInput
	// Type of the resource.
	Kind pulumi.StringPtrInput
	// Configure logging on this health check.
	LogConfig HealthCheckLogConfigPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. For example, a name that is 1-63 characters long, matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`, and otherwise complies with RFC1035. This regular expression describes a name where the first character is a lowercase letter, and all following characters are a dash, lowercase letter, or digit, except the last character, which isn't a dash.
	Name           pulumi.StringPtrInput
	Project        pulumi.StringPtrInput
	RequestId      pulumi.StringPtrInput
	SslHealthCheck SSLHealthCheckPtrInput
	TcpHealthCheck TCPHealthCheckPtrInput
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrInput
	// Specifies the type of the healthCheck, either TCP, SSL, HTTP, HTTPS or HTTP2. Exactly one of the protocol-specific health check field must be specified, which must match type field.
	Type HealthCheckTypePtrInput
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (HealthCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckArgs)(nil)).Elem()
}

type HealthCheckInput interface {
	pulumi.Input

	ToHealthCheckOutput() HealthCheckOutput
	ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput
}

func (*HealthCheck) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthCheck)(nil)).Elem()
}

func (i *HealthCheck) ToHealthCheckOutput() HealthCheckOutput {
	return i.ToHealthCheckOutputWithContext(context.Background())
}

func (i *HealthCheck) ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthCheckOutput)
}

type HealthCheckOutput struct{ *pulumi.OutputState }

func (HealthCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthCheck)(nil)).Elem()
}

func (o HealthCheckOutput) ToHealthCheckOutput() HealthCheckOutput {
	return o
}

func (o HealthCheckOutput) ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HealthCheckInput)(nil)).Elem(), &HealthCheck{})
	pulumi.RegisterOutputType(HealthCheckOutput{})
}
