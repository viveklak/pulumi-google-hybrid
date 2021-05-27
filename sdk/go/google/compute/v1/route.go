// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Route resource in the specified project using the data included in the request.
type Route struct {
	pulumi.CustomResourceState

	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
	DestRange pulumi.StringOutput `pulumi:"destRange"`
	// [Output Only] Type of this resource. Always compute#routes for Route resources.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringOutput `pulumi:"name"`
	// Fully-qualified URL of the network that this route applies to.
	Network pulumi.StringOutput `pulumi:"network"`
	// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL:  projects/project/global/gateways/default-internet-gateway
	NextHopGateway pulumi.StringOutput `pulumi:"nextHopGateway"`
	// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs:
	// - 10.128.0.56
	// - https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// - regions/region/forwardingRules/forwardingRule
	NextHopIlb pulumi.StringOutput `pulumi:"nextHopIlb"`
	// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example:
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
	NextHopInstance pulumi.StringOutput `pulumi:"nextHopInstance"`
	// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
	NextHopIp pulumi.StringOutput `pulumi:"nextHopIp"`
	// The URL of the local network if it should handle matching packets.
	NextHopNetwork pulumi.StringOutput `pulumi:"nextHopNetwork"`
	// [Output Only] The network peering name that should handle matching packets, which should conform to RFC1035.
	NextHopPeering pulumi.StringOutput `pulumi:"nextHopPeering"`
	// The URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel pulumi.StringOutput `pulumi:"nextHopVpnTunnel"`
	// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// [Output Only] Server-defined fully-qualified URL for this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A list of instance tags to which this route applies.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// [Output Only] If potential misconfigurations are detected for this route, this field will be populated with warning messages.
	Warnings RouteWarningsItemResponseArrayOutput `pulumi:"warnings"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Route
	err := ctx.RegisterResource("google-native:compute/v1:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("google-native:compute/v1:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description *string `pulumi:"description"`
	// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
	DestRange *string `pulumi:"destRange"`
	// [Output Only] Type of this resource. Always compute#routes for Route resources.
	Kind *string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name *string `pulumi:"name"`
	// Fully-qualified URL of the network that this route applies to.
	Network *string `pulumi:"network"`
	// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL:  projects/project/global/gateways/default-internet-gateway
	NextHopGateway *string `pulumi:"nextHopGateway"`
	// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs:
	// - 10.128.0.56
	// - https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// - regions/region/forwardingRules/forwardingRule
	NextHopIlb *string `pulumi:"nextHopIlb"`
	// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example:
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
	NextHopInstance *string `pulumi:"nextHopInstance"`
	// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
	NextHopIp *string `pulumi:"nextHopIp"`
	// The URL of the local network if it should handle matching packets.
	NextHopNetwork *string `pulumi:"nextHopNetwork"`
	// [Output Only] The network peering name that should handle matching packets, which should conform to RFC1035.
	NextHopPeering *string `pulumi:"nextHopPeering"`
	// The URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel *string `pulumi:"nextHopVpnTunnel"`
	// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
	Priority *int `pulumi:"priority"`
	// [Output Only] Server-defined fully-qualified URL for this resource.
	SelfLink *string `pulumi:"selfLink"`
	// A list of instance tags to which this route applies.
	Tags []string `pulumi:"tags"`
	// [Output Only] If potential misconfigurations are detected for this route, this field will be populated with warning messages.
	Warnings []RouteWarningsItemResponse `pulumi:"warnings"`
}

type RouteState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringPtrInput
	// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
	DestRange pulumi.StringPtrInput
	// [Output Only] Type of this resource. Always compute#routes for Route resources.
	Kind pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringPtrInput
	// Fully-qualified URL of the network that this route applies to.
	Network pulumi.StringPtrInput
	// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL:  projects/project/global/gateways/default-internet-gateway
	NextHopGateway pulumi.StringPtrInput
	// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs:
	// - 10.128.0.56
	// - https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// - regions/region/forwardingRules/forwardingRule
	NextHopIlb pulumi.StringPtrInput
	// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example:
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
	NextHopInstance pulumi.StringPtrInput
	// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
	NextHopIp pulumi.StringPtrInput
	// The URL of the local network if it should handle matching packets.
	NextHopNetwork pulumi.StringPtrInput
	// [Output Only] The network peering name that should handle matching packets, which should conform to RFC1035.
	NextHopPeering pulumi.StringPtrInput
	// The URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel pulumi.StringPtrInput
	// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
	Priority pulumi.IntPtrInput
	// [Output Only] Server-defined fully-qualified URL for this resource.
	SelfLink pulumi.StringPtrInput
	// A list of instance tags to which this route applies.
	Tags pulumi.StringArrayInput
	// [Output Only] If potential misconfigurations are detected for this route, this field will be populated with warning messages.
	Warnings RouteWarningsItemResponseArrayInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description *string `pulumi:"description"`
	// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
	DestRange *string `pulumi:"destRange"`
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id *string `pulumi:"id"`
	// [Output Only] Type of this resource. Always compute#routes for Route resources.
	Kind *string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name *string `pulumi:"name"`
	// Fully-qualified URL of the network that this route applies to.
	Network *string `pulumi:"network"`
	// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL:  projects/project/global/gateways/default-internet-gateway
	NextHopGateway *string `pulumi:"nextHopGateway"`
	// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs:
	// - 10.128.0.56
	// - https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// - regions/region/forwardingRules/forwardingRule
	NextHopIlb *string `pulumi:"nextHopIlb"`
	// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example:
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
	NextHopInstance *string `pulumi:"nextHopInstance"`
	// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
	NextHopIp *string `pulumi:"nextHopIp"`
	// The URL of the local network if it should handle matching packets.
	NextHopNetwork *string `pulumi:"nextHopNetwork"`
	// [Output Only] The network peering name that should handle matching packets, which should conform to RFC1035.
	NextHopPeering *string `pulumi:"nextHopPeering"`
	// The URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel *string `pulumi:"nextHopVpnTunnel"`
	// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
	Priority  *int    `pulumi:"priority"`
	Project   string  `pulumi:"project"`
	RequestId *string `pulumi:"requestId"`
	// [Output Only] Server-defined fully-qualified URL for this resource.
	SelfLink *string `pulumi:"selfLink"`
	// A list of instance tags to which this route applies.
	Tags []string `pulumi:"tags"`
	// [Output Only] If potential misconfigurations are detected for this route, this field will be populated with warning messages.
	Warnings []RouteWarningsItem `pulumi:"warnings"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringPtrInput
	// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
	DestRange pulumi.StringPtrInput
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id pulumi.StringPtrInput
	// [Output Only] Type of this resource. Always compute#routes for Route resources.
	Kind pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringPtrInput
	// Fully-qualified URL of the network that this route applies to.
	Network pulumi.StringPtrInput
	// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL:  projects/project/global/gateways/default-internet-gateway
	NextHopGateway pulumi.StringPtrInput
	// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs:
	// - 10.128.0.56
	// - https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// - regions/region/forwardingRules/forwardingRule
	NextHopIlb pulumi.StringPtrInput
	// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example:
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
	NextHopInstance pulumi.StringPtrInput
	// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
	NextHopIp pulumi.StringPtrInput
	// The URL of the local network if it should handle matching packets.
	NextHopNetwork pulumi.StringPtrInput
	// [Output Only] The network peering name that should handle matching packets, which should conform to RFC1035.
	NextHopPeering pulumi.StringPtrInput
	// The URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel pulumi.StringPtrInput
	// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
	Priority  pulumi.IntPtrInput
	Project   pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// [Output Only] Server-defined fully-qualified URL for this resource.
	SelfLink pulumi.StringPtrInput
	// A list of instance tags to which this route applies.
	Tags pulumi.StringArrayInput
	// [Output Only] If potential misconfigurations are detected for this route, this field will be populated with warning messages.
	Warnings RouteWarningsItemArrayInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

type RouteOutput struct {
	*pulumi.OutputState
}

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
}
