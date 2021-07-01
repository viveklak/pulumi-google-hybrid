// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a target pool in the specified project and region using the data included in the request.
type TargetPool struct {
	pulumi.CustomResourceState

	// The server-defined URL for the resource. This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool, and its failoverRatio field is properly set to a value between [0, 1].
	//
	// backupPool and failoverRatio together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below failoverRatio, traffic arriving at the load-balanced IP will be directed to the backup pool.
	//
	// In case where failoverRatio and backupPool are not set, or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
	BackupPool pulumi.StringOutput `pulumi:"backupPool"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool (i.e., not as a backup pool to some other target pool). The value of the field must be in [0, 1].
	//
	// If set, backupPool must also be set. They together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below this number, traffic arriving at the load-balanced IP will be directed to the backup pool.
	//
	// In case where failoverRatio is not set or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
	FailoverRatio pulumi.Float64Output `pulumi:"failoverRatio"`
	// The URL of the HttpHealthCheck resource. A member instance in this pool is considered healthy if and only if the health checks pass. Only legacy HttpHealthChecks are supported. Only one health check may be specified.
	HealthChecks pulumi.StringArrayOutput `pulumi:"healthChecks"`
	// A list of resource URLs to the virtual machine instances serving this pool. They must live in zones contained in the same region as this pool.
	Instances pulumi.StringArrayOutput `pulumi:"instances"`
	// Type of the resource. Always compute#targetPool for target pools.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the region where the target pool resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// Session affinity option, must be one of the following values:
	// NONE: Connections from the same client IP may go to any instance in the pool.
	// CLIENT_IP: Connections from the same client IP will go to the same instance in the pool while that instance remains healthy.
	// CLIENT_IP_PROTO: Connections from the same client IP with the same IP protocol will go to the same instance in the pool while that instance remains healthy.
	SessionAffinity pulumi.StringOutput `pulumi:"sessionAffinity"`
}

// NewTargetPool registers a new resource with the given unique name, arguments, and options.
func NewTargetPool(ctx *pulumi.Context,
	name string, args *TargetPoolArgs, opts ...pulumi.ResourceOption) (*TargetPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource TargetPool
	err := ctx.RegisterResource("google-native:compute/alpha:TargetPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetPool gets an existing TargetPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetPoolState, opts ...pulumi.ResourceOption) (*TargetPool, error) {
	var resource TargetPool
	err := ctx.ReadResource("google-native:compute/alpha:TargetPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetPool resources.
type targetPoolState struct {
}

type TargetPoolState struct {
}

func (TargetPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetPoolState)(nil)).Elem()
}

type targetPoolArgs struct {
	// The server-defined URL for the resource. This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool, and its failoverRatio field is properly set to a value between [0, 1].
	//
	// backupPool and failoverRatio together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below failoverRatio, traffic arriving at the load-balanced IP will be directed to the backup pool.
	//
	// In case where failoverRatio and backupPool are not set, or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
	BackupPool *string `pulumi:"backupPool"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool (i.e., not as a backup pool to some other target pool). The value of the field must be in [0, 1].
	//
	// If set, backupPool must also be set. They together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below this number, traffic arriving at the load-balanced IP will be directed to the backup pool.
	//
	// In case where failoverRatio is not set or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
	FailoverRatio *float64 `pulumi:"failoverRatio"`
	// The URL of the HttpHealthCheck resource. A member instance in this pool is considered healthy if and only if the health checks pass. Only legacy HttpHealthChecks are supported. Only one health check may be specified.
	HealthChecks []string `pulumi:"healthChecks"`
	// A list of resource URLs to the virtual machine instances serving this pool. They must live in zones contained in the same region as this pool.
	Instances []string `pulumi:"instances"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      *string `pulumi:"name"`
	Project   string  `pulumi:"project"`
	Region    string  `pulumi:"region"`
	RequestId *string `pulumi:"requestId"`
	// Session affinity option, must be one of the following values:
	// NONE: Connections from the same client IP may go to any instance in the pool.
	// CLIENT_IP: Connections from the same client IP will go to the same instance in the pool while that instance remains healthy.
	// CLIENT_IP_PROTO: Connections from the same client IP with the same IP protocol will go to the same instance in the pool while that instance remains healthy.
	SessionAffinity *string `pulumi:"sessionAffinity"`
}

// The set of arguments for constructing a TargetPool resource.
type TargetPoolArgs struct {
	// The server-defined URL for the resource. This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool, and its failoverRatio field is properly set to a value between [0, 1].
	//
	// backupPool and failoverRatio together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below failoverRatio, traffic arriving at the load-balanced IP will be directed to the backup pool.
	//
	// In case where failoverRatio and backupPool are not set, or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
	BackupPool pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool (i.e., not as a backup pool to some other target pool). The value of the field must be in [0, 1].
	//
	// If set, backupPool must also be set. They together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below this number, traffic arriving at the load-balanced IP will be directed to the backup pool.
	//
	// In case where failoverRatio is not set or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
	FailoverRatio pulumi.Float64PtrInput
	// The URL of the HttpHealthCheck resource. A member instance in this pool is considered healthy if and only if the health checks pass. Only legacy HttpHealthChecks are supported. Only one health check may be specified.
	HealthChecks pulumi.StringArrayInput
	// A list of resource URLs to the virtual machine instances serving this pool. They must live in zones contained in the same region as this pool.
	Instances pulumi.StringArrayInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringInput
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// Session affinity option, must be one of the following values:
	// NONE: Connections from the same client IP may go to any instance in the pool.
	// CLIENT_IP: Connections from the same client IP will go to the same instance in the pool while that instance remains healthy.
	// CLIENT_IP_PROTO: Connections from the same client IP with the same IP protocol will go to the same instance in the pool while that instance remains healthy.
	SessionAffinity *TargetPoolSessionAffinity
}

func (TargetPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetPoolArgs)(nil)).Elem()
}

type TargetPoolInput interface {
	pulumi.Input

	ToTargetPoolOutput() TargetPoolOutput
	ToTargetPoolOutputWithContext(ctx context.Context) TargetPoolOutput
}

func (*TargetPool) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetPool)(nil))
}

func (i *TargetPool) ToTargetPoolOutput() TargetPoolOutput {
	return i.ToTargetPoolOutputWithContext(context.Background())
}

func (i *TargetPool) ToTargetPoolOutputWithContext(ctx context.Context) TargetPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPoolOutput)
}

type TargetPoolOutput struct {
	*pulumi.OutputState
}

func (TargetPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetPool)(nil))
}

func (o TargetPoolOutput) ToTargetPoolOutput() TargetPoolOutput {
	return o
}

func (o TargetPoolOutput) ToTargetPoolOutputWithContext(ctx context.Context) TargetPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TargetPoolOutput{})
}
