// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a NodeTemplate resource in the specified project using the data included in the request.
type NodeTemplate struct {
	pulumi.CustomResourceState

	Accelerators AcceleratorConfigResponseArrayOutput `pulumi:"accelerators"`
	// CPU overcommit.
	CpuOvercommitType pulumi.StringOutput `pulumi:"cpuOvercommitType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput          `pulumi:"description"`
	Disks       LocalDiskResponseArrayOutput `pulumi:"disks"`
	// The type of the resource. Always compute#nodeTemplate for node templates.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels pulumi.StringMapOutput `pulumi:"nodeAffinityLabels"`
	// The node type to use for nodes group that are created from this template.
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties.
	//
	// This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityResponseOutput `pulumi:"nodeTypeFlexibility"`
	// The name of the region where the node template resides, such as us-central1.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// Sets the binding properties for the physical server. Valid values include:
	// - [Default] RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server
	// - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible
	//
	// See Sole-tenant node options for more information.
	ServerBinding ServerBindingResponseOutput `pulumi:"serverBinding"`
	// The status of the node template. One of the following values: CREATING, READY, and DELETING.
	Status pulumi.StringOutput `pulumi:"status"`
	// An optional, human-readable explanation of the status.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
}

// NewNodeTemplate registers a new resource with the given unique name, arguments, and options.
func NewNodeTemplate(ctx *pulumi.Context,
	name string, args *NodeTemplateArgs, opts ...pulumi.ResourceOption) (*NodeTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource NodeTemplate
	err := ctx.RegisterResource("google-native:compute/alpha:NodeTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeTemplate gets an existing NodeTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeTemplateState, opts ...pulumi.ResourceOption) (*NodeTemplate, error) {
	var resource NodeTemplate
	err := ctx.ReadResource("google-native:compute/alpha:NodeTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeTemplate resources.
type nodeTemplateState struct {
}

type NodeTemplateState struct {
}

func (NodeTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTemplateState)(nil)).Elem()
}

type nodeTemplateArgs struct {
	Accelerators []AcceleratorConfig `pulumi:"accelerators"`
	// CPU overcommit.
	CpuOvercommitType *string `pulumi:"cpuOvercommitType"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string     `pulumi:"description"`
	Disks       []LocalDisk `pulumi:"disks"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels map[string]string `pulumi:"nodeAffinityLabels"`
	// The node type to use for nodes group that are created from this template.
	NodeType *string `pulumi:"nodeType"`
	// The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties.
	//
	// This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
	NodeTypeFlexibility *NodeTemplateNodeTypeFlexibility `pulumi:"nodeTypeFlexibility"`
	Project             string                           `pulumi:"project"`
	Region              string                           `pulumi:"region"`
	RequestId           *string                          `pulumi:"requestId"`
	// Sets the binding properties for the physical server. Valid values include:
	// - [Default] RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server
	// - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible
	//
	// See Sole-tenant node options for more information.
	ServerBinding *ServerBinding `pulumi:"serverBinding"`
}

// The set of arguments for constructing a NodeTemplate resource.
type NodeTemplateArgs struct {
	Accelerators AcceleratorConfigArrayInput
	// CPU overcommit.
	CpuOvercommitType *NodeTemplateCpuOvercommitType
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	Disks       LocalDiskArrayInput
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels pulumi.StringMapInput
	// The node type to use for nodes group that are created from this template.
	NodeType pulumi.StringPtrInput
	// The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties.
	//
	// This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityPtrInput
	Project             pulumi.StringInput
	Region              pulumi.StringInput
	RequestId           pulumi.StringPtrInput
	// Sets the binding properties for the physical server. Valid values include:
	// - [Default] RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server
	// - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible
	//
	// See Sole-tenant node options for more information.
	ServerBinding ServerBindingPtrInput
}

func (NodeTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTemplateArgs)(nil)).Elem()
}

type NodeTemplateInput interface {
	pulumi.Input

	ToNodeTemplateOutput() NodeTemplateOutput
	ToNodeTemplateOutputWithContext(ctx context.Context) NodeTemplateOutput
}

func (*NodeTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTemplate)(nil))
}

func (i *NodeTemplate) ToNodeTemplateOutput() NodeTemplateOutput {
	return i.ToNodeTemplateOutputWithContext(context.Background())
}

func (i *NodeTemplate) ToNodeTemplateOutputWithContext(ctx context.Context) NodeTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTemplateOutput)
}

type NodeTemplateOutput struct {
	*pulumi.OutputState
}

func (NodeTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTemplate)(nil))
}

func (o NodeTemplateOutput) ToNodeTemplateOutput() NodeTemplateOutput {
	return o
}

func (o NodeTemplateOutput) ToNodeTemplateOutputWithContext(ctx context.Context) NodeTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NodeTemplateOutput{})
}
