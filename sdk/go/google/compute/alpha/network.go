// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a network in the specified project using the data included in the request.
type Network struct {
	pulumi.CustomResourceState
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Network
	err := ctx.RegisterResource("google-cloud:compute/alpha:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetwork gets an existing Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("google-cloud:compute/alpha:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network resources.
type networkState struct {
}

type NetworkState struct {
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	// Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
	IPv4Range *string `pulumi:"IPv4Range"`
	// Must be set to create a VPC network. If not set, a legacy network is created.
	//
	// When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode.
	//
	// An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges.
	//
	// For custom mode VPC networks, you can add subnets using the subnetworks insert method.
	AutoCreateSubnetworks *bool `pulumi:"autoCreateSubnetworks"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description *string `pulumi:"description"`
	// [Output Only] URL of the firewall policy the network is associated with.
	FirewallPolicy *string `pulumi:"firewallPolicy"`
	// [Output Only] The gateway address for default routing out of the network, selected by GCP.
	GatewayIPv4 *string `pulumi:"gatewayIPv4"`
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id *string `pulumi:"id"`
	// [Output Only] Type of the resource. Always compute#network for networks.
	Kind *string `pulumi:"kind"`
	// Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes.
	Mtu *int `pulumi:"mtu"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name    *string `pulumi:"name"`
	Network string  `pulumi:"network"`
	// [Output Only] A list of network peerings for the resource.
	Peerings []NetworkPeering `pulumi:"peerings"`
	Project  string           `pulumi:"project"`
	// The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
	RoutingConfig *NetworkRoutingConfig `pulumi:"routingConfig"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// [Output Only] Server-defined fully-qualified URLs for all subnetworks in this VPC network.
	Subnetworks []string `pulumi:"subnetworks"`
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
	IPv4Range pulumi.StringPtrInput
	// Must be set to create a VPC network. If not set, a legacy network is created.
	//
	// When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode.
	//
	// An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges.
	//
	// For custom mode VPC networks, you can add subnets using the subnetworks insert method.
	AutoCreateSubnetworks pulumi.BoolPtrInput
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringPtrInput
	// [Output Only] URL of the firewall policy the network is associated with.
	FirewallPolicy pulumi.StringPtrInput
	// [Output Only] The gateway address for default routing out of the network, selected by GCP.
	GatewayIPv4 pulumi.StringPtrInput
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id pulumi.StringPtrInput
	// [Output Only] Type of the resource. Always compute#network for networks.
	Kind pulumi.StringPtrInput
	// Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes.
	Mtu pulumi.IntPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name    pulumi.StringPtrInput
	Network pulumi.StringInput
	// [Output Only] A list of network peerings for the resource.
	Peerings NetworkPeeringArrayInput
	Project  pulumi.StringInput
	// The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
	RoutingConfig NetworkRoutingConfigPtrInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	// [Output Only] Server-defined fully-qualified URLs for all subnetworks in this VPC network.
	Subnetworks pulumi.StringArrayInput
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkArgs)(nil)).Elem()
}

type NetworkInput interface {
	pulumi.Input

	ToNetworkOutput() NetworkOutput
	ToNetworkOutputWithContext(ctx context.Context) NetworkOutput
}

func (*Network) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil))
}

func (i *Network) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i *Network) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
}

type NetworkOutput struct {
	*pulumi.OutputState
}

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil))
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkOutput{})
}
