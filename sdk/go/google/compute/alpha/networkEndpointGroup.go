// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a network endpoint group in the specified project using the parameters that are included in the request.
type NetworkEndpointGroup struct {
	pulumi.CustomResourceState
}

// NewNetworkEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkEndpointGroup(ctx *pulumi.Context,
	name string, args *NetworkEndpointGroupArgs, opts ...pulumi.ResourceOption) (*NetworkEndpointGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkEndpointGroup == nil {
		return nil, errors.New("invalid value for required argument 'NetworkEndpointGroup'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource NetworkEndpointGroup
	err := ctx.RegisterResource("google-cloud:compute/alpha:NetworkEndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkEndpointGroup gets an existing NetworkEndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkEndpointGroupState, opts ...pulumi.ResourceOption) (*NetworkEndpointGroup, error) {
	var resource NetworkEndpointGroup
	err := ctx.ReadResource("google-cloud:compute/alpha:NetworkEndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkEndpointGroup resources.
type networkEndpointGroupState struct {
}

type NetworkEndpointGroupState struct {
}

func (NetworkEndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkEndpointGroupState)(nil)).Elem()
}

type networkEndpointGroupArgs struct {
	// Metadata defined as annotations on the network endpoint group.
	Annotations map[string]string `pulumi:"annotations"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine *NetworkEndpointGroupAppEngine `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction *NetworkEndpointGroupCloudFunction `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun *NetworkEndpointGroupCloudRun `pulumi:"cloudRun"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort *int `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id *string `pulumi:"id"`
	// [Output Only] Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
	Kind *string `pulumi:"kind"`
	// This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
	LoadBalancer *NetworkEndpointGroupLbNetworkEndpointGroup `pulumi:"loadBalancer"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network              *string `pulumi:"network"`
	NetworkEndpointGroup string  `pulumi:"networkEndpointGroup"`
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, or SERVERLESS.
	NetworkEndpointType *string `pulumi:"networkEndpointType"`
	Project             string  `pulumi:"project"`
	// The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
	PscTargetService *string `pulumi:"pscTargetService"`
	// [Output Only] The URL of the region where the network endpoint group is located.
	Region *string `pulumi:"region"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine cloudFunction or serverlessDeployment may be set.
	ServerlessDeployment *NetworkEndpointGroupServerlessDeployment `pulumi:"serverlessDeployment"`
	// [Output only] Number of network endpoints in the network endpoint group.
	Size *int `pulumi:"size"`
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork *string `pulumi:"subnetwork"`
	// Specify the type of this network endpoint group. Only LOAD_BALANCING is valid for now.
	Type *string `pulumi:"type"`
	// [Output Only] The URL of the zone where the network endpoint group is located.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a NetworkEndpointGroup resource.
type NetworkEndpointGroupArgs struct {
	// Metadata defined as annotations on the network endpoint group.
	Annotations pulumi.StringMapInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine NetworkEndpointGroupAppEnginePtrInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction NetworkEndpointGroupCloudFunctionPtrInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun NetworkEndpointGroupCloudRunPtrInput
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort pulumi.IntPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id pulumi.StringPtrInput
	// [Output Only] Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
	Kind pulumi.StringPtrInput
	// This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
	LoadBalancer NetworkEndpointGroupLbNetworkEndpointGroupPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network              pulumi.StringPtrInput
	NetworkEndpointGroup pulumi.StringInput
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, or SERVERLESS.
	NetworkEndpointType pulumi.StringPtrInput
	Project             pulumi.StringInput
	// The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
	PscTargetService pulumi.StringPtrInput
	// [Output Only] The URL of the region where the network endpoint group is located.
	Region pulumi.StringPtrInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine cloudFunction or serverlessDeployment may be set.
	ServerlessDeployment NetworkEndpointGroupServerlessDeploymentPtrInput
	// [Output only] Number of network endpoints in the network endpoint group.
	Size pulumi.IntPtrInput
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork pulumi.StringPtrInput
	// Specify the type of this network endpoint group. Only LOAD_BALANCING is valid for now.
	Type pulumi.StringPtrInput
	// [Output Only] The URL of the zone where the network endpoint group is located.
	Zone pulumi.StringInput
}

func (NetworkEndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkEndpointGroupArgs)(nil)).Elem()
}

type NetworkEndpointGroupInput interface {
	pulumi.Input

	ToNetworkEndpointGroupOutput() NetworkEndpointGroupOutput
	ToNetworkEndpointGroupOutputWithContext(ctx context.Context) NetworkEndpointGroupOutput
}

func (*NetworkEndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkEndpointGroup)(nil))
}

func (i *NetworkEndpointGroup) ToNetworkEndpointGroupOutput() NetworkEndpointGroupOutput {
	return i.ToNetworkEndpointGroupOutputWithContext(context.Background())
}

func (i *NetworkEndpointGroup) ToNetworkEndpointGroupOutputWithContext(ctx context.Context) NetworkEndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkEndpointGroupOutput)
}

type NetworkEndpointGroupOutput struct {
	*pulumi.OutputState
}

func (NetworkEndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkEndpointGroup)(nil))
}

func (o NetworkEndpointGroupOutput) ToNetworkEndpointGroupOutput() NetworkEndpointGroupOutput {
	return o
}

func (o NetworkEndpointGroupOutput) ToNetworkEndpointGroupOutputWithContext(ctx context.Context) NetworkEndpointGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkEndpointGroupOutput{})
}
