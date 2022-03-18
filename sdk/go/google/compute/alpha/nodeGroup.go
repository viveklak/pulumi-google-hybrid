// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a NodeGroup resource in the specified project using the data included in the request.
type NodeGroup struct {
	pulumi.CustomResourceState

	// Specifies how autoscaling should behave.
	AutoscalingPolicy NodeGroupAutoscalingPolicyResponseOutput `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The type of the resource. Always compute#nodeGroup for node group.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
	LocationHint pulumi.StringOutput `pulumi:"locationHint"`
	// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see Maintenance policies.
	MaintenancePolicy pulumi.StringOutput                      `pulumi:"maintenancePolicy"`
	MaintenanceWindow NodeGroupMaintenanceWindowResponseOutput `pulumi:"maintenanceWindow"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the node template to create the node group from.
	NodeTemplate pulumi.StringOutput `pulumi:"nodeTemplate"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// Share-settings for the node group
	ShareSettings ShareSettingsResponseOutput `pulumi:"shareSettings"`
	// The total number of nodes in the node group.
	Size   pulumi.IntOutput    `pulumi:"size"`
	Status pulumi.StringOutput `pulumi:"status"`
	// The name of the zone where the node group resides, such as us-central1-a.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewNodeGroup registers a new resource with the given unique name, arguments, and options.
func NewNodeGroup(ctx *pulumi.Context,
	name string, args *NodeGroupArgs, opts ...pulumi.ResourceOption) (*NodeGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InitialNodeCount == nil {
		return nil, errors.New("invalid value for required argument 'InitialNodeCount'")
	}
	var resource NodeGroup
	err := ctx.RegisterResource("google-native:compute/alpha:NodeGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeGroup gets an existing NodeGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeGroupState, opts ...pulumi.ResourceOption) (*NodeGroup, error) {
	var resource NodeGroup
	err := ctx.ReadResource("google-native:compute/alpha:NodeGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeGroup resources.
type nodeGroupState struct {
}

type NodeGroupState struct {
}

func (NodeGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeGroupState)(nil)).Elem()
}

type nodeGroupArgs struct {
	// Specifies how autoscaling should behave.
	AutoscalingPolicy *NodeGroupAutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Initial count of nodes in the node group.
	InitialNodeCount string `pulumi:"initialNodeCount"`
	// An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
	LocationHint *string `pulumi:"locationHint"`
	// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see Maintenance policies.
	MaintenancePolicy *NodeGroupMaintenancePolicy `pulumi:"maintenancePolicy"`
	MaintenanceWindow *NodeGroupMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// URL of the node template to create the node group from.
	NodeTemplate *string `pulumi:"nodeTemplate"`
	Project      *string `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Share-settings for the node group
	ShareSettings *ShareSettings `pulumi:"shareSettings"`
	Zone          *string        `pulumi:"zone"`
}

// The set of arguments for constructing a NodeGroup resource.
type NodeGroupArgs struct {
	// Specifies how autoscaling should behave.
	AutoscalingPolicy NodeGroupAutoscalingPolicyPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Initial count of nodes in the node group.
	InitialNodeCount pulumi.StringInput
	// An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
	LocationHint pulumi.StringPtrInput
	// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see Maintenance policies.
	MaintenancePolicy NodeGroupMaintenancePolicyPtrInput
	MaintenanceWindow NodeGroupMaintenanceWindowPtrInput
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// URL of the node template to create the node group from.
	NodeTemplate pulumi.StringPtrInput
	Project      pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Share-settings for the node group
	ShareSettings ShareSettingsPtrInput
	Zone          pulumi.StringPtrInput
}

func (NodeGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeGroupArgs)(nil)).Elem()
}

type NodeGroupInput interface {
	pulumi.Input

	ToNodeGroupOutput() NodeGroupOutput
	ToNodeGroupOutputWithContext(ctx context.Context) NodeGroupOutput
}

func (*NodeGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroup)(nil)).Elem()
}

func (i *NodeGroup) ToNodeGroupOutput() NodeGroupOutput {
	return i.ToNodeGroupOutputWithContext(context.Background())
}

func (i *NodeGroup) ToNodeGroupOutputWithContext(ctx context.Context) NodeGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupOutput)
}

type NodeGroupOutput struct{ *pulumi.OutputState }

func (NodeGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroup)(nil)).Elem()
}

func (o NodeGroupOutput) ToNodeGroupOutput() NodeGroupOutput {
	return o
}

func (o NodeGroupOutput) ToNodeGroupOutputWithContext(ctx context.Context) NodeGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeGroupInput)(nil)).Elem(), &NodeGroup{})
	pulumi.RegisterOutputType(NodeGroupOutput{})
}
