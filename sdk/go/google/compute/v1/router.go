// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Router resource in the specified project and region using the data included in the request.
type Router struct {
	pulumi.CustomResourceState

	// BGP information specific to this router.
	Bgp RouterBgpResponseOutput `pulumi:"bgp"`
	// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
	BgpPeers RouterBgpPeerResponseArrayOutput `pulumi:"bgpPeers"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Field to indicate if a router is dedicated to use with encrypted Interconnect Attachment (IPsec-encrypted Cloud Interconnect feature).
	// Not currently available in all Interconnect locations.
	EncryptedInterconnectRouter pulumi.BoolOutput `pulumi:"encryptedInterconnectRouter"`
	// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
	Interfaces RouterInterfaceResponseArrayOutput `pulumi:"interfaces"`
	// Type of resource. Always compute#router for routers.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of NAT services created in this router.
	Nats RouterNatResponseArrayOutput `pulumi:"nats"`
	// URI of the network to which this router belongs.
	Network pulumi.StringOutput `pulumi:"network"`
	// URI of the region where the router resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewRouter registers a new resource with the given unique name, arguments, and options.
func NewRouter(ctx *pulumi.Context,
	name string, args *RouterArgs, opts ...pulumi.ResourceOption) (*Router, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource Router
	err := ctx.RegisterResource("google-native:compute/v1:Router", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouter gets an existing Router resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterState, opts ...pulumi.ResourceOption) (*Router, error) {
	var resource Router
	err := ctx.ReadResource("google-native:compute/v1:Router", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Router resources.
type routerState struct {
	// BGP information specific to this router.
	Bgp *RouterBgpResponse `pulumi:"bgp"`
	// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
	BgpPeers []RouterBgpPeerResponse `pulumi:"bgpPeers"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Field to indicate if a router is dedicated to use with encrypted Interconnect Attachment (IPsec-encrypted Cloud Interconnect feature).
	// Not currently available in all Interconnect locations.
	EncryptedInterconnectRouter *bool `pulumi:"encryptedInterconnectRouter"`
	// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
	Interfaces []RouterInterfaceResponse `pulumi:"interfaces"`
	// Type of resource. Always compute#router for routers.
	Kind *string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// A list of NAT services created in this router.
	Nats []RouterNatResponse `pulumi:"nats"`
	// URI of the network to which this router belongs.
	Network *string `pulumi:"network"`
	// URI of the region where the router resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region *string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
}

type RouterState struct {
	// BGP information specific to this router.
	Bgp RouterBgpResponsePtrInput
	// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
	BgpPeers RouterBgpPeerResponseArrayInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Field to indicate if a router is dedicated to use with encrypted Interconnect Attachment (IPsec-encrypted Cloud Interconnect feature).
	// Not currently available in all Interconnect locations.
	EncryptedInterconnectRouter pulumi.BoolPtrInput
	// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
	Interfaces RouterInterfaceResponseArrayInput
	// Type of resource. Always compute#router for routers.
	Kind pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// A list of NAT services created in this router.
	Nats RouterNatResponseArrayInput
	// URI of the network to which this router belongs.
	Network pulumi.StringPtrInput
	// URI of the region where the router resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringPtrInput
	// Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
}

func (RouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerState)(nil)).Elem()
}

type routerArgs struct {
	// BGP information specific to this router.
	Bgp *RouterBgp `pulumi:"bgp"`
	// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
	BgpPeers []RouterBgpPeer `pulumi:"bgpPeers"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Field to indicate if a router is dedicated to use with encrypted Interconnect Attachment (IPsec-encrypted Cloud Interconnect feature).
	// Not currently available in all Interconnect locations.
	EncryptedInterconnectRouter *bool `pulumi:"encryptedInterconnectRouter"`
	// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
	Interfaces []RouterInterface `pulumi:"interfaces"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// A list of NAT services created in this router.
	Nats []RouterNat `pulumi:"nats"`
	// URI of the network to which this router belongs.
	Network   *string `pulumi:"network"`
	Project   string  `pulumi:"project"`
	Region    string  `pulumi:"region"`
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a Router resource.
type RouterArgs struct {
	// BGP information specific to this router.
	Bgp RouterBgpPtrInput
	// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
	BgpPeers RouterBgpPeerArrayInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Field to indicate if a router is dedicated to use with encrypted Interconnect Attachment (IPsec-encrypted Cloud Interconnect feature).
	// Not currently available in all Interconnect locations.
	EncryptedInterconnectRouter pulumi.BoolPtrInput
	// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
	Interfaces RouterInterfaceArrayInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// A list of NAT services created in this router.
	Nats RouterNatArrayInput
	// URI of the network to which this router belongs.
	Network   pulumi.StringPtrInput
	Project   pulumi.StringInput
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
}

func (RouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerArgs)(nil)).Elem()
}

type RouterInput interface {
	pulumi.Input

	ToRouterOutput() RouterOutput
	ToRouterOutputWithContext(ctx context.Context) RouterOutput
}

func (*Router) ElementType() reflect.Type {
	return reflect.TypeOf((*Router)(nil))
}

func (i *Router) ToRouterOutput() RouterOutput {
	return i.ToRouterOutputWithContext(context.Background())
}

func (i *Router) ToRouterOutputWithContext(ctx context.Context) RouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterOutput)
}

type RouterOutput struct {
	*pulumi.OutputState
}

func (RouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Router)(nil))
}

func (o RouterOutput) ToRouterOutput() RouterOutput {
	return o
}

func (o RouterOutput) ToRouterOutputWithContext(ctx context.Context) RouterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouterOutput{})
}
