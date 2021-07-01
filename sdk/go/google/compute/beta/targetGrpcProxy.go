// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a TargetGrpcProxy in the specified project in the given scope using the parameters that are included in the request.
type TargetGrpcProxy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetGrpcProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetGrpcProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetGrpcProxy.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Type of the resource. Always compute#targetGrpcProxy for target grpc proxies.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL with id for the resource.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService. The protocol field in the BackendService must be set to GRPC.
	UrlMap pulumi.StringOutput `pulumi:"urlMap"`
	// If true, indicates that the BackendServices referenced by the urlMap may be accessed by gRPC applications without using a sidecar proxy. This will enable configuration checks on urlMap and its referenced BackendServices to not allow unsupported features. A gRPC application must use "xds:///" scheme in the target URI of the service it is connecting to. If false, indicates that the BackendServices referenced by the urlMap will be accessed by gRPC applications via a sidecar proxy. In this case, a gRPC application must not use "xds:///" scheme in the target URI of the service it is connecting to
	ValidateForProxyless pulumi.BoolOutput `pulumi:"validateForProxyless"`
}

// NewTargetGrpcProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetGrpcProxy(ctx *pulumi.Context,
	name string, args *TargetGrpcProxyArgs, opts ...pulumi.ResourceOption) (*TargetGrpcProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource TargetGrpcProxy
	err := ctx.RegisterResource("google-native:compute/beta:TargetGrpcProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetGrpcProxy gets an existing TargetGrpcProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetGrpcProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetGrpcProxyState, opts ...pulumi.ResourceOption) (*TargetGrpcProxy, error) {
	var resource TargetGrpcProxy
	err := ctx.ReadResource("google-native:compute/beta:TargetGrpcProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetGrpcProxy resources.
type targetGrpcProxyState struct {
}

type TargetGrpcProxyState struct {
}

func (TargetGrpcProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGrpcProxyState)(nil)).Elem()
}

type targetGrpcProxyArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      *string `pulumi:"name"`
	Project   string  `pulumi:"project"`
	RequestId *string `pulumi:"requestId"`
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService. The protocol field in the BackendService must be set to GRPC.
	UrlMap *string `pulumi:"urlMap"`
	// If true, indicates that the BackendServices referenced by the urlMap may be accessed by gRPC applications without using a sidecar proxy. This will enable configuration checks on urlMap and its referenced BackendServices to not allow unsupported features. A gRPC application must use "xds:///" scheme in the target URI of the service it is connecting to. If false, indicates that the BackendServices referenced by the urlMap will be accessed by gRPC applications via a sidecar proxy. In this case, a gRPC application must not use "xds:///" scheme in the target URI of the service it is connecting to
	ValidateForProxyless *bool `pulumi:"validateForProxyless"`
}

// The set of arguments for constructing a TargetGrpcProxy resource.
type TargetGrpcProxyArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService. The protocol field in the BackendService must be set to GRPC.
	UrlMap pulumi.StringPtrInput
	// If true, indicates that the BackendServices referenced by the urlMap may be accessed by gRPC applications without using a sidecar proxy. This will enable configuration checks on urlMap and its referenced BackendServices to not allow unsupported features. A gRPC application must use "xds:///" scheme in the target URI of the service it is connecting to. If false, indicates that the BackendServices referenced by the urlMap will be accessed by gRPC applications via a sidecar proxy. In this case, a gRPC application must not use "xds:///" scheme in the target URI of the service it is connecting to
	ValidateForProxyless pulumi.BoolPtrInput
}

func (TargetGrpcProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGrpcProxyArgs)(nil)).Elem()
}

type TargetGrpcProxyInput interface {
	pulumi.Input

	ToTargetGrpcProxyOutput() TargetGrpcProxyOutput
	ToTargetGrpcProxyOutputWithContext(ctx context.Context) TargetGrpcProxyOutput
}

func (*TargetGrpcProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetGrpcProxy)(nil))
}

func (i *TargetGrpcProxy) ToTargetGrpcProxyOutput() TargetGrpcProxyOutput {
	return i.ToTargetGrpcProxyOutputWithContext(context.Background())
}

func (i *TargetGrpcProxy) ToTargetGrpcProxyOutputWithContext(ctx context.Context) TargetGrpcProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGrpcProxyOutput)
}

type TargetGrpcProxyOutput struct {
	*pulumi.OutputState
}

func (TargetGrpcProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetGrpcProxy)(nil))
}

func (o TargetGrpcProxyOutput) ToTargetGrpcProxyOutput() TargetGrpcProxyOutput {
	return o
}

func (o TargetGrpcProxyOutput) ToTargetGrpcProxyOutputWithContext(ctx context.Context) TargetGrpcProxyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TargetGrpcProxyOutput{})
}
