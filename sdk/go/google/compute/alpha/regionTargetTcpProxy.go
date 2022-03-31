// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a TargetTcpProxy resource in the specified project and region using the data included in the request.
type RegionTargetTcpProxy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of the resource. Always compute#targetTcpProxy for target TCP proxies.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
	ProxyBind pulumi.BoolOutput `pulumi:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
	ProxyHeader pulumi.StringOutput `pulumi:"proxyHeader"`
	// URL of the region where the regional TCP proxy resides. This field is not applicable to global TCP proxy.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// URL to the BackendService resource.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewRegionTargetTcpProxy registers a new resource with the given unique name, arguments, and options.
func NewRegionTargetTcpProxy(ctx *pulumi.Context,
	name string, args *RegionTargetTcpProxyArgs, opts ...pulumi.ResourceOption) (*RegionTargetTcpProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionTargetTcpProxy
	err := ctx.RegisterResource("google-native:compute/alpha:RegionTargetTcpProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionTargetTcpProxy gets an existing RegionTargetTcpProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionTargetTcpProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionTargetTcpProxyState, opts ...pulumi.ResourceOption) (*RegionTargetTcpProxy, error) {
	var resource RegionTargetTcpProxy
	err := ctx.ReadResource("google-native:compute/alpha:RegionTargetTcpProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionTargetTcpProxy resources.
type regionTargetTcpProxyState struct {
}

type RegionTargetTcpProxyState struct {
}

func (RegionTargetTcpProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionTargetTcpProxyState)(nil)).Elem()
}

type regionTargetTcpProxyArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
	ProxyBind *bool `pulumi:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
	ProxyHeader *RegionTargetTcpProxyProxyHeader `pulumi:"proxyHeader"`
	Region      string                           `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// URL to the BackendService resource.
	Service *string `pulumi:"service"`
}

// The set of arguments for constructing a RegionTargetTcpProxy resource.
type RegionTargetTcpProxyArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
	ProxyBind pulumi.BoolPtrInput
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
	ProxyHeader RegionTargetTcpProxyProxyHeaderPtrInput
	Region      pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// URL to the BackendService resource.
	Service pulumi.StringPtrInput
}

func (RegionTargetTcpProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionTargetTcpProxyArgs)(nil)).Elem()
}

type RegionTargetTcpProxyInput interface {
	pulumi.Input

	ToRegionTargetTcpProxyOutput() RegionTargetTcpProxyOutput
	ToRegionTargetTcpProxyOutputWithContext(ctx context.Context) RegionTargetTcpProxyOutput
}

func (*RegionTargetTcpProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionTargetTcpProxy)(nil)).Elem()
}

func (i *RegionTargetTcpProxy) ToRegionTargetTcpProxyOutput() RegionTargetTcpProxyOutput {
	return i.ToRegionTargetTcpProxyOutputWithContext(context.Background())
}

func (i *RegionTargetTcpProxy) ToRegionTargetTcpProxyOutputWithContext(ctx context.Context) RegionTargetTcpProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionTargetTcpProxyOutput)
}

type RegionTargetTcpProxyOutput struct{ *pulumi.OutputState }

func (RegionTargetTcpProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionTargetTcpProxy)(nil)).Elem()
}

func (o RegionTargetTcpProxyOutput) ToRegionTargetTcpProxyOutput() RegionTargetTcpProxyOutput {
	return o
}

func (o RegionTargetTcpProxyOutput) ToRegionTargetTcpProxyOutputWithContext(ctx context.Context) RegionTargetTcpProxyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionTargetTcpProxyInput)(nil)).Elem(), &RegionTargetTcpProxy{})
	pulumi.RegisterOutputType(RegionTargetTcpProxyOutput{})
}
