// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VPN gateway in the specified project and region using the data included in the request.
type VpnGateway struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of resource. Always compute#vpnGateway for VPN gateways.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this VpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve an VpnGateway.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	Network pulumi.StringOutput `pulumi:"network"`
	// URL of the region where the VPN gateway resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A list of interfaces on this VPN gateway.
	VpnInterfaces VpnGatewayVpnGatewayInterfaceResponseArrayOutput `pulumi:"vpnInterfaces"`
}

// NewVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewVpnGateway(ctx *pulumi.Context,
	name string, args *VpnGatewayArgs, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource VpnGateway
	err := ctx.RegisterResource("google-native:compute/beta:VpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnGateway gets an existing VpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnGatewayState, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	var resource VpnGateway
	err := ctx.ReadResource("google-native:compute/beta:VpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnGateway resources.
type vpnGatewayState struct {
}

type VpnGatewayState struct {
}

func (VpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayState)(nil)).Elem()
}

type vpnGatewayArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	Network   *string `pulumi:"network"`
	Project   string  `pulumi:"project"`
	Region    string  `pulumi:"region"`
	RequestId *string `pulumi:"requestId"`
	// A list of interfaces on this VPN gateway.
	VpnInterfaces []VpnGatewayVpnGatewayInterface `pulumi:"vpnInterfaces"`
}

// The set of arguments for constructing a VpnGateway resource.
type VpnGatewayArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	Network   pulumi.StringPtrInput
	Project   pulumi.StringInput
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// A list of interfaces on this VPN gateway.
	VpnInterfaces VpnGatewayVpnGatewayInterfaceArrayInput
}

func (VpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayArgs)(nil)).Elem()
}

type VpnGatewayInput interface {
	pulumi.Input

	ToVpnGatewayOutput() VpnGatewayOutput
	ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput
}

func (*VpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGateway)(nil))
}

func (i *VpnGateway) ToVpnGatewayOutput() VpnGatewayOutput {
	return i.ToVpnGatewayOutputWithContext(context.Background())
}

func (i *VpnGateway) ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayOutput)
}

type VpnGatewayOutput struct {
	*pulumi.OutputState
}

func (VpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGateway)(nil))
}

func (o VpnGatewayOutput) ToVpnGatewayOutput() VpnGatewayOutput {
	return o
}

func (o VpnGatewayOutput) ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VpnGatewayOutput{})
}
