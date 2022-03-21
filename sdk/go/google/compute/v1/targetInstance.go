// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a TargetInstance resource in the specified project and zone using the data included in the request.
type TargetInstance struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// A URL to the virtual machine instance that handles traffic for this target instance. When creating a target instance, you can provide the fully-qualified URL or a valid partial URL to the desired virtual machine. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance - zones/zone/instances/instance
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The type of the resource. Always compute#targetInstance for target instances.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// NAT option controlling how IPs are NAT'ed to the instance. Currently only NO_NAT (default value) is supported.
	NatPolicy pulumi.StringOutput `pulumi:"natPolicy"`
	// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	Network pulumi.StringOutput `pulumi:"network"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// URL of the zone where the target instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewTargetInstance registers a new resource with the given unique name, arguments, and options.
func NewTargetInstance(ctx *pulumi.Context,
	name string, args *TargetInstanceArgs, opts ...pulumi.ResourceOption) (*TargetInstance, error) {
	if args == nil {
		args = &TargetInstanceArgs{}
	}

	var resource TargetInstance
	err := ctx.RegisterResource("google-native:compute/v1:TargetInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetInstance gets an existing TargetInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetInstanceState, opts ...pulumi.ResourceOption) (*TargetInstance, error) {
	var resource TargetInstance
	err := ctx.ReadResource("google-native:compute/v1:TargetInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetInstance resources.
type targetInstanceState struct {
}

type TargetInstanceState struct {
}

func (TargetInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetInstanceState)(nil)).Elem()
}

type targetInstanceArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// A URL to the virtual machine instance that handles traffic for this target instance. When creating a target instance, you can provide the fully-qualified URL or a valid partial URL to the desired virtual machine. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance - zones/zone/instances/instance
	Instance *string `pulumi:"instance"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// NAT option controlling how IPs are NAT'ed to the instance. Currently only NO_NAT (default value) is supported.
	NatPolicy *TargetInstanceNatPolicy `pulumi:"natPolicy"`
	// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	Network *string `pulumi:"network"`
	Project *string `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	Zone      *string `pulumi:"zone"`
}

// The set of arguments for constructing a TargetInstance resource.
type TargetInstanceArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// A URL to the virtual machine instance that handles traffic for this target instance. When creating a target instance, you can provide the fully-qualified URL or a valid partial URL to the desired virtual machine. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance - zones/zone/instances/instance
	Instance pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// NAT option controlling how IPs are NAT'ed to the instance. Currently only NO_NAT (default value) is supported.
	NatPolicy TargetInstanceNatPolicyPtrInput
	// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	Network pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	Zone      pulumi.StringPtrInput
}

func (TargetInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetInstanceArgs)(nil)).Elem()
}

type TargetInstanceInput interface {
	pulumi.Input

	ToTargetInstanceOutput() TargetInstanceOutput
	ToTargetInstanceOutputWithContext(ctx context.Context) TargetInstanceOutput
}

func (*TargetInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetInstance)(nil)).Elem()
}

func (i *TargetInstance) ToTargetInstanceOutput() TargetInstanceOutput {
	return i.ToTargetInstanceOutputWithContext(context.Background())
}

func (i *TargetInstance) ToTargetInstanceOutputWithContext(ctx context.Context) TargetInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetInstanceOutput)
}

type TargetInstanceOutput struct{ *pulumi.OutputState }

func (TargetInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetInstance)(nil)).Elem()
}

func (o TargetInstanceOutput) ToTargetInstanceOutput() TargetInstanceOutput {
	return o
}

func (o TargetInstanceOutput) ToTargetInstanceOutputWithContext(ctx context.Context) TargetInstanceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetInstanceInput)(nil)).Elem(), &TargetInstance{})
	pulumi.RegisterOutputType(TargetInstanceOutput{})
}
