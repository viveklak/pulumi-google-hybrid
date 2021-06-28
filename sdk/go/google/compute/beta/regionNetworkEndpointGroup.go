// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a network endpoint group in the specified project using the parameters that are included in the request.
type RegionNetworkEndpointGroup struct {
	pulumi.CustomResourceState

	// Metadata defined as annotations on the network endpoint group.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine NetworkEndpointGroupAppEngineResponseOutput `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction NetworkEndpointGroupCloudFunctionResponseOutput `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun NetworkEndpointGroupCloudRunResponseOutput `pulumi:"cloudRun"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort pulumi.IntOutput `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network pulumi.StringOutput `pulumi:"network"`
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, or SERVERLESS.
	NetworkEndpointType pulumi.StringOutput `pulumi:"networkEndpointType"`
	// The URL of the region where the network endpoint group is located.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// [Output only] Number of network endpoints in the network endpoint group.
	Size pulumi.IntOutput `pulumi:"size"`
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	// The URL of the zone where the network endpoint group is located.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewRegionNetworkEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewRegionNetworkEndpointGroup(ctx *pulumi.Context,
	name string, args *RegionNetworkEndpointGroupArgs, opts ...pulumi.ResourceOption) (*RegionNetworkEndpointGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionNetworkEndpointGroup
	err := ctx.RegisterResource("google-native:compute/beta:RegionNetworkEndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionNetworkEndpointGroup gets an existing RegionNetworkEndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionNetworkEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionNetworkEndpointGroupState, opts ...pulumi.ResourceOption) (*RegionNetworkEndpointGroup, error) {
	var resource RegionNetworkEndpointGroup
	err := ctx.ReadResource("google-native:compute/beta:RegionNetworkEndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionNetworkEndpointGroup resources.
type regionNetworkEndpointGroupState struct {
	// Metadata defined as annotations on the network endpoint group.
	Annotations map[string]string `pulumi:"annotations"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine *NetworkEndpointGroupAppEngineResponse `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction *NetworkEndpointGroupCloudFunctionResponse `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun *NetworkEndpointGroupCloudRunResponse `pulumi:"cloudRun"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort *int `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
	Kind *string `pulumi:"kind"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network *string `pulumi:"network"`
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, or SERVERLESS.
	NetworkEndpointType *string `pulumi:"networkEndpointType"`
	// The URL of the region where the network endpoint group is located.
	Region *string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output only] Number of network endpoints in the network endpoint group.
	Size *int `pulumi:"size"`
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork *string `pulumi:"subnetwork"`
	// The URL of the zone where the network endpoint group is located.
	Zone *string `pulumi:"zone"`
}

type RegionNetworkEndpointGroupState struct {
	// Metadata defined as annotations on the network endpoint group.
	Annotations pulumi.StringMapInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine NetworkEndpointGroupAppEngineResponsePtrInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction NetworkEndpointGroupCloudFunctionResponsePtrInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun NetworkEndpointGroupCloudRunResponsePtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort pulumi.IntPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
	Kind pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network pulumi.StringPtrInput
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, or SERVERLESS.
	NetworkEndpointType pulumi.StringPtrInput
	// The URL of the region where the network endpoint group is located.
	Region pulumi.StringPtrInput
	// Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// [Output only] Number of network endpoints in the network endpoint group.
	Size pulumi.IntPtrInput
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork pulumi.StringPtrInput
	// The URL of the zone where the network endpoint group is located.
	Zone pulumi.StringPtrInput
}

func (RegionNetworkEndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNetworkEndpointGroupState)(nil)).Elem()
}

type regionNetworkEndpointGroupArgs struct {
	// Metadata defined as annotations on the network endpoint group.
	Annotations map[string]string `pulumi:"annotations"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine *NetworkEndpointGroupAppEngine `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction *NetworkEndpointGroupCloudFunction `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun *NetworkEndpointGroupCloudRun `pulumi:"cloudRun"`
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort *int `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network *string `pulumi:"network"`
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, or SERVERLESS.
	NetworkEndpointType *string `pulumi:"networkEndpointType"`
	Project             string  `pulumi:"project"`
	Region              string  `pulumi:"region"`
	RequestId           *string `pulumi:"requestId"`
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork *string `pulumi:"subnetwork"`
}

// The set of arguments for constructing a RegionNetworkEndpointGroup resource.
type RegionNetworkEndpointGroupArgs struct {
	// Metadata defined as annotations on the network endpoint group.
	Annotations pulumi.StringMapInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine NetworkEndpointGroupAppEnginePtrInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction NetworkEndpointGroupCloudFunctionPtrInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun NetworkEndpointGroupCloudRunPtrInput
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort pulumi.IntPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network pulumi.StringPtrInput
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, or SERVERLESS.
	NetworkEndpointType *RegionNetworkEndpointGroupNetworkEndpointType
	Project             pulumi.StringInput
	Region              pulumi.StringInput
	RequestId           pulumi.StringPtrInput
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork pulumi.StringPtrInput
}

func (RegionNetworkEndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNetworkEndpointGroupArgs)(nil)).Elem()
}

type RegionNetworkEndpointGroupInput interface {
	pulumi.Input

	ToRegionNetworkEndpointGroupOutput() RegionNetworkEndpointGroupOutput
	ToRegionNetworkEndpointGroupOutputWithContext(ctx context.Context) RegionNetworkEndpointGroupOutput
}

func (*RegionNetworkEndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionNetworkEndpointGroup)(nil))
}

func (i *RegionNetworkEndpointGroup) ToRegionNetworkEndpointGroupOutput() RegionNetworkEndpointGroupOutput {
	return i.ToRegionNetworkEndpointGroupOutputWithContext(context.Background())
}

func (i *RegionNetworkEndpointGroup) ToRegionNetworkEndpointGroupOutputWithContext(ctx context.Context) RegionNetworkEndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionNetworkEndpointGroupOutput)
}

type RegionNetworkEndpointGroupOutput struct {
	*pulumi.OutputState
}

func (RegionNetworkEndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionNetworkEndpointGroup)(nil))
}

func (o RegionNetworkEndpointGroupOutput) ToRegionNetworkEndpointGroupOutput() RegionNetworkEndpointGroupOutput {
	return o
}

func (o RegionNetworkEndpointGroupOutput) ToRegionNetworkEndpointGroupOutputWithContext(ctx context.Context) RegionNetworkEndpointGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionNetworkEndpointGroupOutput{})
}
