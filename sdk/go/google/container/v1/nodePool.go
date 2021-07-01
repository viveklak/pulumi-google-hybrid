// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a node pool for a cluster.
type NodePool struct {
	pulumi.CustomResourceState

	// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
	Autoscaling NodePoolAutoscalingResponseOutput `pulumi:"autoscaling"`
	// Which conditions caused the current node pool state.
	Conditions StatusConditionResponseArrayOutput `pulumi:"conditions"`
	// The node configuration of the pool.
	Config NodeConfigResponseOutput `pulumi:"config"`
	// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
	InitialNodeCount pulumi.IntOutput `pulumi:"initialNodeCount"`
	// [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
	InstanceGroupUrls pulumi.StringArrayOutput `pulumi:"instanceGroupUrls"`
	// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
	Locations pulumi.StringArrayOutput `pulumi:"locations"`
	// NodeManagement configuration for this NodePool.
	Management NodeManagementResponseOutput `pulumi:"management"`
	// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
	MaxPodsConstraint MaxPodsConstraintResponseOutput `pulumi:"maxPodsConstraint"`
	// The name of the node pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// [Output only] The pod CIDR block size per node in this node pool.
	PodIpv4CidrSize pulumi.IntOutput `pulumi:"podIpv4CidrSize"`
	// [Output only] Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// [Output only] The status of the nodes in this pool instance.
	Status pulumi.StringOutput `pulumi:"status"`
	// Upgrade settings control disruption and speed of the upgrade.
	UpgradeSettings UpgradeSettingsResponseOutput `pulumi:"upgradeSettings"`
	// The version of the Kubernetes of this node.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewNodePool registers a new resource with the given unique name, arguments, and options.
func NewNodePool(ctx *pulumi.Context,
	name string, args *NodePoolArgs, opts ...pulumi.ResourceOption) (*NodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource NodePool
	err := ctx.RegisterResource("google-native:container/v1:NodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodePool gets an existing NodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodePoolState, opts ...pulumi.ResourceOption) (*NodePool, error) {
	var resource NodePool
	err := ctx.ReadResource("google-native:container/v1:NodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodePool resources.
type nodePoolState struct {
}

type NodePoolState struct {
}

func (NodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodePoolState)(nil)).Elem()
}

type nodePoolArgs struct {
	// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
	Autoscaling *NodePoolAutoscaling `pulumi:"autoscaling"`
	ClusterId   string               `pulumi:"clusterId"`
	// Which conditions caused the current node pool state.
	Conditions []StatusCondition `pulumi:"conditions"`
	// The node configuration of the pool.
	Config *NodeConfig `pulumi:"config"`
	// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
	InitialNodeCount *int   `pulumi:"initialNodeCount"`
	Location         string `pulumi:"location"`
	// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
	Locations []string `pulumi:"locations"`
	// NodeManagement configuration for this NodePool.
	Management *NodeManagement `pulumi:"management"`
	// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
	MaxPodsConstraint *MaxPodsConstraint `pulumi:"maxPodsConstraint"`
	// The name of the node pool.
	Name *string `pulumi:"name"`
	// The parent (project, location, cluster id) where the node pool will be created. Specified in the format `projects/*/locations/*/clusters/*`.
	Parent  *string `pulumi:"parent"`
	Project string  `pulumi:"project"`
	// Upgrade settings control disruption and speed of the upgrade.
	UpgradeSettings *UpgradeSettings `pulumi:"upgradeSettings"`
	// The version of the Kubernetes of this node.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a NodePool resource.
type NodePoolArgs struct {
	// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
	Autoscaling NodePoolAutoscalingPtrInput
	ClusterId   pulumi.StringInput
	// Which conditions caused the current node pool state.
	Conditions StatusConditionArrayInput
	// The node configuration of the pool.
	Config NodeConfigPtrInput
	// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
	InitialNodeCount pulumi.IntPtrInput
	Location         pulumi.StringInput
	// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
	Locations pulumi.StringArrayInput
	// NodeManagement configuration for this NodePool.
	Management NodeManagementPtrInput
	// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
	MaxPodsConstraint MaxPodsConstraintPtrInput
	// The name of the node pool.
	Name pulumi.StringPtrInput
	// The parent (project, location, cluster id) where the node pool will be created. Specified in the format `projects/*/locations/*/clusters/*`.
	Parent  pulumi.StringPtrInput
	Project pulumi.StringInput
	// Upgrade settings control disruption and speed of the upgrade.
	UpgradeSettings UpgradeSettingsPtrInput
	// The version of the Kubernetes of this node.
	Version pulumi.StringPtrInput
}

func (NodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodePoolArgs)(nil)).Elem()
}

type NodePoolInput interface {
	pulumi.Input

	ToNodePoolOutput() NodePoolOutput
	ToNodePoolOutputWithContext(ctx context.Context) NodePoolOutput
}

func (*NodePool) ElementType() reflect.Type {
	return reflect.TypeOf((*NodePool)(nil))
}

func (i *NodePool) ToNodePoolOutput() NodePoolOutput {
	return i.ToNodePoolOutputWithContext(context.Background())
}

func (i *NodePool) ToNodePoolOutputWithContext(ctx context.Context) NodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolOutput)
}

type NodePoolOutput struct {
	*pulumi.OutputState
}

func (NodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodePool)(nil))
}

func (o NodePoolOutput) ToNodePoolOutput() NodePoolOutput {
	return o
}

func (o NodePoolOutput) ToNodePoolOutputWithContext(ctx context.Context) NodePoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NodePoolOutput{})
}
