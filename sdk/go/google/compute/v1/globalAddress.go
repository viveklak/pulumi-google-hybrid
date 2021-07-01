// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an address resource in the specified project by using the data included in the request.
type GlobalAddress struct {
	pulumi.CustomResourceState

	// The static IP address represented by this resource.
	Address pulumi.StringOutput `pulumi:"address"`
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType pulumi.StringOutput `pulumi:"addressType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion pulumi.StringOutput `pulumi:"ipVersion"`
	// Type of the resource. Always compute#address for addresses.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network pulumi.StringOutput `pulumi:"network"`
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Global forwarding rules can only be Premium Tier. Regional forwarding rules can be either Premium or Standard Tier. Standard Tier addresses applied to regional forwarding rules can be used with any external load balancer. Regional forwarding rules in Premium Tier can only be used with a network load balancer.
	//
	// If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier pulumi.StringOutput `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength pulumi.IntOutput `pulumi:"prefixLength"`
	// The purpose of this resource, which can be one of the following values:
	// - `GCE_ENDPOINT` for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// - `DNS_RESOLVER` for a DNS resolver address in a subnetwork
	// - `VPC_PEERING` for addresses that are reserved for VPC peer networks.
	// - `NAT_AUTO` for addresses that are external IP addresses automatically reserved for Cloud NAT.
	// - `IPSEC_INTERCONNECT` for addresses created from a private IP range that are reserved for a VLAN attachment in an IPsec-encrypted Cloud Interconnect configuration. These addresses are regional resources.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// The URL of the region where a regional address resides. For regional addresses, you must specify the region as a path parameter in the HTTP request URL. This field is not applicable to global addresses.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The status of the address, which can be one of RESERVING, RESERVED, or IN_USE. An address that is RESERVING is currently in the process of being reserved. A RESERVED address is currently reserved and available to use. An IN_USE address is currently being used by another resource and is not available.
	Status pulumi.StringOutput `pulumi:"status"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	// The URLs of the resources that are using this address.
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewGlobalAddress registers a new resource with the given unique name, arguments, and options.
func NewGlobalAddress(ctx *pulumi.Context,
	name string, args *GlobalAddressArgs, opts ...pulumi.ResourceOption) (*GlobalAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource GlobalAddress
	err := ctx.RegisterResource("google-native:compute/v1:GlobalAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalAddress gets an existing GlobalAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalAddressState, opts ...pulumi.ResourceOption) (*GlobalAddress, error) {
	var resource GlobalAddress
	err := ctx.ReadResource("google-native:compute/v1:GlobalAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalAddress resources.
type globalAddressState struct {
}

type GlobalAddressState struct {
}

func (GlobalAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalAddressState)(nil)).Elem()
}

type globalAddressArgs struct {
	// The static IP address represented by this resource.
	Address *string `pulumi:"address"`
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType *string `pulumi:"addressType"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description *string `pulumi:"description"`
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion *string `pulumi:"ipVersion"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name *string `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network *string `pulumi:"network"`
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Global forwarding rules can only be Premium Tier. Regional forwarding rules can be either Premium or Standard Tier. Standard Tier addresses applied to regional forwarding rules can be used with any external load balancer. Regional forwarding rules in Premium Tier can only be used with a network load balancer.
	//
	// If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier *string `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength *int   `pulumi:"prefixLength"`
	Project      string `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values:
	// - `GCE_ENDPOINT` for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// - `DNS_RESOLVER` for a DNS resolver address in a subnetwork
	// - `VPC_PEERING` for addresses that are reserved for VPC peer networks.
	// - `NAT_AUTO` for addresses that are external IP addresses automatically reserved for Cloud NAT.
	// - `IPSEC_INTERCONNECT` for addresses created from a private IP range that are reserved for a VLAN attachment in an IPsec-encrypted Cloud Interconnect configuration. These addresses are regional resources.
	Purpose   *string `pulumi:"purpose"`
	RequestId *string `pulumi:"requestId"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork *string `pulumi:"subnetwork"`
}

// The set of arguments for constructing a GlobalAddress resource.
type GlobalAddressArgs struct {
	// The static IP address represented by this resource.
	Address pulumi.StringPtrInput
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType *GlobalAddressAddressType
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringPtrInput
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion *GlobalAddressIpVersion
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringPtrInput
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network pulumi.StringPtrInput
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Global forwarding rules can only be Premium Tier. Regional forwarding rules can be either Premium or Standard Tier. Standard Tier addresses applied to regional forwarding rules can be used with any external load balancer. Regional forwarding rules in Premium Tier can only be used with a network load balancer.
	//
	// If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier *GlobalAddressNetworkTier
	// The prefix length if the resource represents an IP range.
	PrefixLength pulumi.IntPtrInput
	Project      pulumi.StringInput
	// The purpose of this resource, which can be one of the following values:
	// - `GCE_ENDPOINT` for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// - `DNS_RESOLVER` for a DNS resolver address in a subnetwork
	// - `VPC_PEERING` for addresses that are reserved for VPC peer networks.
	// - `NAT_AUTO` for addresses that are external IP addresses automatically reserved for Cloud NAT.
	// - `IPSEC_INTERCONNECT` for addresses created from a private IP range that are reserved for a VLAN attachment in an IPsec-encrypted Cloud Interconnect configuration. These addresses are regional resources.
	Purpose   *GlobalAddressPurpose
	RequestId pulumi.StringPtrInput
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork pulumi.StringPtrInput
}

func (GlobalAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalAddressArgs)(nil)).Elem()
}

type GlobalAddressInput interface {
	pulumi.Input

	ToGlobalAddressOutput() GlobalAddressOutput
	ToGlobalAddressOutputWithContext(ctx context.Context) GlobalAddressOutput
}

func (*GlobalAddress) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalAddress)(nil))
}

func (i *GlobalAddress) ToGlobalAddressOutput() GlobalAddressOutput {
	return i.ToGlobalAddressOutputWithContext(context.Background())
}

func (i *GlobalAddress) ToGlobalAddressOutputWithContext(ctx context.Context) GlobalAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalAddressOutput)
}

type GlobalAddressOutput struct {
	*pulumi.OutputState
}

func (GlobalAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalAddress)(nil))
}

func (o GlobalAddressOutput) ToGlobalAddressOutput() GlobalAddressOutput {
	return o
}

func (o GlobalAddressOutput) ToGlobalAddressOutputWithContext(ctx context.Context) GlobalAddressOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GlobalAddressOutput{})
}
