// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a ExternalVpnGateway in the specified project using the data included in the request.
type ExternalVpnGateway struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// List of interfaces for this external VPN gateway.
	Interfaces ExternalVpnGatewayInterfaceResponseArrayOutput `pulumi:"interfaces"`
	// Type of the resource. Always compute#externalVpnGateway for externalVpnGateways.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this ExternalVpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve an ExternalVpnGateway.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates the user-supplied redundancy type of this external VPN gateway.
	RedundancyType pulumi.StringOutput `pulumi:"redundancyType"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewExternalVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewExternalVpnGateway(ctx *pulumi.Context,
	name string, args *ExternalVpnGatewayArgs, opts ...pulumi.ResourceOption) (*ExternalVpnGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource ExternalVpnGateway
	err := ctx.RegisterResource("google-native:compute/alpha:ExternalVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalVpnGateway gets an existing ExternalVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalVpnGatewayState, opts ...pulumi.ResourceOption) (*ExternalVpnGateway, error) {
	var resource ExternalVpnGateway
	err := ctx.ReadResource("google-native:compute/alpha:ExternalVpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalVpnGateway resources.
type externalVpnGatewayState struct {
}

type ExternalVpnGatewayState struct {
}

func (ExternalVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalVpnGatewayState)(nil)).Elem()
}

type externalVpnGatewayArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// List of interfaces for this external VPN gateway.
	Interfaces []ExternalVpnGatewayInterface `pulumi:"interfaces"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// Indicates the user-supplied redundancy type of this external VPN gateway.
	RedundancyType *string `pulumi:"redundancyType"`
	RequestId      *string `pulumi:"requestId"`
}

// The set of arguments for constructing a ExternalVpnGateway resource.
type ExternalVpnGatewayArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// List of interfaces for this external VPN gateway.
	Interfaces ExternalVpnGatewayInterfaceArrayInput
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// Indicates the user-supplied redundancy type of this external VPN gateway.
	RedundancyType *ExternalVpnGatewayRedundancyType
	RequestId      pulumi.StringPtrInput
}

func (ExternalVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalVpnGatewayArgs)(nil)).Elem()
}

type ExternalVpnGatewayInput interface {
	pulumi.Input

	ToExternalVpnGatewayOutput() ExternalVpnGatewayOutput
	ToExternalVpnGatewayOutputWithContext(ctx context.Context) ExternalVpnGatewayOutput
}

func (*ExternalVpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalVpnGateway)(nil))
}

func (i *ExternalVpnGateway) ToExternalVpnGatewayOutput() ExternalVpnGatewayOutput {
	return i.ToExternalVpnGatewayOutputWithContext(context.Background())
}

func (i *ExternalVpnGateway) ToExternalVpnGatewayOutputWithContext(ctx context.Context) ExternalVpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalVpnGatewayOutput)
}

type ExternalVpnGatewayOutput struct {
	*pulumi.OutputState
}

func (ExternalVpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalVpnGateway)(nil))
}

func (o ExternalVpnGatewayOutput) ToExternalVpnGatewayOutput() ExternalVpnGatewayOutput {
	return o
}

func (o ExternalVpnGatewayOutput) ToExternalVpnGatewayOutputWithContext(ctx context.Context) ExternalVpnGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExternalVpnGatewayOutput{})
}
