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
type ClusterNodePool struct {
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

// NewClusterNodePool registers a new resource with the given unique name, arguments, and options.
func NewClusterNodePool(ctx *pulumi.Context,
	name string, args *ClusterNodePoolArgs, opts ...pulumi.ResourceOption) (*ClusterNodePool, error) {
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
	var resource ClusterNodePool
	err := ctx.RegisterResource("google-native:container/v1:ClusterNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterNodePool gets an existing ClusterNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterNodePoolState, opts ...pulumi.ResourceOption) (*ClusterNodePool, error) {
	var resource ClusterNodePool
	err := ctx.ReadResource("google-native:container/v1:ClusterNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterNodePool resources.
type clusterNodePoolState struct {
	// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
	Autoscaling *NodePoolAutoscalingResponse `pulumi:"autoscaling"`
	// Which conditions caused the current node pool state.
	Conditions []StatusConditionResponse `pulumi:"conditions"`
	// The node configuration of the pool.
	Config *NodeConfigResponse `pulumi:"config"`
	// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
	InitialNodeCount *int `pulumi:"initialNodeCount"`
	// [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
	InstanceGroupUrls []string `pulumi:"instanceGroupUrls"`
	// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
	Locations []string `pulumi:"locations"`
	// NodeManagement configuration for this NodePool.
	Management *NodeManagementResponse `pulumi:"management"`
	// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
	MaxPodsConstraint *MaxPodsConstraintResponse `pulumi:"maxPodsConstraint"`
	// The name of the node pool.
	Name *string `pulumi:"name"`
	// [Output only] The pod CIDR block size per node in this node pool.
	PodIpv4CidrSize *int `pulumi:"podIpv4CidrSize"`
	// [Output only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output only] The status of the nodes in this pool instance.
	Status *string `pulumi:"status"`
	// Upgrade settings control disruption and speed of the upgrade.
	UpgradeSettings *UpgradeSettingsResponse `pulumi:"upgradeSettings"`
	// The version of the Kubernetes of this node.
	Version *string `pulumi:"version"`
}

type ClusterNodePoolState struct {
	// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
	Autoscaling NodePoolAutoscalingResponsePtrInput
	// Which conditions caused the current node pool state.
	Conditions StatusConditionResponseArrayInput
	// The node configuration of the pool.
	Config NodeConfigResponsePtrInput
	// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
	InitialNodeCount pulumi.IntPtrInput
	// [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
	InstanceGroupUrls pulumi.StringArrayInput
	// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
	Locations pulumi.StringArrayInput
	// NodeManagement configuration for this NodePool.
	Management NodeManagementResponsePtrInput
	// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
	MaxPodsConstraint MaxPodsConstraintResponsePtrInput
	// The name of the node pool.
	Name pulumi.StringPtrInput
	// [Output only] The pod CIDR block size per node in this node pool.
	PodIpv4CidrSize pulumi.IntPtrInput
	// [Output only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// [Output only] The status of the nodes in this pool instance.
	Status pulumi.StringPtrInput
	// Upgrade settings control disruption and speed of the upgrade.
	UpgradeSettings UpgradeSettingsResponsePtrInput
	// The version of the Kubernetes of this node.
	Version pulumi.StringPtrInput
}

func (ClusterNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterNodePoolState)(nil)).Elem()
}

type clusterNodePoolArgs struct {
	// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
	Autoscaling *NodePoolAutoscaling `pulumi:"autoscaling"`
	ClusterId   string               `pulumi:"clusterId"`
	// Which conditions caused the current node pool state.
	Conditions []StatusCondition `pulumi:"conditions"`
	// The node configuration of the pool.
	Config *NodeConfig `pulumi:"config"`
	// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
	InitialNodeCount *int `pulumi:"initialNodeCount"`
	// [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
	InstanceGroupUrls []string `pulumi:"instanceGroupUrls"`
	Location          string   `pulumi:"location"`
	// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
	Locations []string `pulumi:"locations"`
	// NodeManagement configuration for this NodePool.
	Management *NodeManagement `pulumi:"management"`
	// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
	MaxPodsConstraint *MaxPodsConstraint `pulumi:"maxPodsConstraint"`
	// The name of the node pool.
	Name *string `pulumi:"name"`
	// The parent (project, location, cluster id) where the node pool will be created. Specified in the format `projects/*/locations/*/clusters/*`.
	Parent *string `pulumi:"parent"`
	// [Output only] The pod CIDR block size per node in this node pool.
	PodIpv4CidrSize *int   `pulumi:"podIpv4CidrSize"`
	Project         string `pulumi:"project"`
	// [Output only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output only] The status of the nodes in this pool instance.
	Status *string `pulumi:"status"`
	// Upgrade settings control disruption and speed of the upgrade.
	UpgradeSettings *UpgradeSettings `pulumi:"upgradeSettings"`
	// The version of the Kubernetes of this node.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a ClusterNodePool resource.
type ClusterNodePoolArgs struct {
	// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
	Autoscaling NodePoolAutoscalingPtrInput
	ClusterId   pulumi.StringInput
	// Which conditions caused the current node pool state.
	Conditions StatusConditionArrayInput
	// The node configuration of the pool.
	Config NodeConfigPtrInput
	// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
	InitialNodeCount pulumi.IntPtrInput
	// [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
	InstanceGroupUrls pulumi.StringArrayInput
	Location          pulumi.StringInput
	// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
	Locations pulumi.StringArrayInput
	// NodeManagement configuration for this NodePool.
	Management NodeManagementPtrInput
	// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
	MaxPodsConstraint MaxPodsConstraintPtrInput
	// The name of the node pool.
	Name pulumi.StringPtrInput
	// The parent (project, location, cluster id) where the node pool will be created. Specified in the format `projects/*/locations/*/clusters/*`.
	Parent pulumi.StringPtrInput
	// [Output only] The pod CIDR block size per node in this node pool.
	PodIpv4CidrSize pulumi.IntPtrInput
	Project         pulumi.StringInput
	// [Output only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// [Output only] The status of the nodes in this pool instance.
	Status pulumi.StringPtrInput
	// Upgrade settings control disruption and speed of the upgrade.
	UpgradeSettings UpgradeSettingsPtrInput
	// The version of the Kubernetes of this node.
	Version pulumi.StringPtrInput
}

func (ClusterNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterNodePoolArgs)(nil)).Elem()
}

type ClusterNodePoolInput interface {
	pulumi.Input

	ToClusterNodePoolOutput() ClusterNodePoolOutput
	ToClusterNodePoolOutputWithContext(ctx context.Context) ClusterNodePoolOutput
}

func (*ClusterNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterNodePool)(nil))
}

func (i *ClusterNodePool) ToClusterNodePoolOutput() ClusterNodePoolOutput {
	return i.ToClusterNodePoolOutputWithContext(context.Background())
}

func (i *ClusterNodePool) ToClusterNodePoolOutputWithContext(ctx context.Context) ClusterNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterNodePoolOutput)
}

type ClusterNodePoolOutput struct {
	*pulumi.OutputState
}

func (ClusterNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterNodePool)(nil))
}

func (o ClusterNodePoolOutput) ToClusterNodePoolOutput() ClusterNodePoolOutput {
	return o
}

func (o ClusterNodePoolOutput) ToClusterNodePoolOutputWithContext(ctx context.Context) ClusterNodePoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterNodePoolOutput{})
}
