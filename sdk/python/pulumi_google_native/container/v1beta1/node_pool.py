# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['NodePoolInitArgs', 'NodePool']

@pulumi.input_type
class NodePoolInitArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 autoscaling: Optional[pulumi.Input['NodePoolAutoscalingArgs']] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['StatusConditionArgs']]]] = None,
                 config: Optional[pulumi.Input['NodeConfigArgs']] = None,
                 initial_node_count: Optional[pulumi.Input[int]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 management: Optional[pulumi.Input['NodeManagementArgs']] = None,
                 max_pods_constraint: Optional[pulumi.Input['MaxPodsConstraintArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input['NodeNetworkConfigArgs']] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 placement_policy: Optional[pulumi.Input['PlacementPolicyArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 upgrade_settings: Optional[pulumi.Input['UpgradeSettingsArgs']] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NodePool resource.
        :param pulumi.Input[str] cluster_id: Deprecated. The name of the cluster. This field has been deprecated and replaced by the parent field.
        :param pulumi.Input['NodePoolAutoscalingArgs'] autoscaling: Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
        :param pulumi.Input[Sequence[pulumi.Input['StatusConditionArgs']]] conditions: Which conditions caused the current node pool state.
        :param pulumi.Input['NodeConfigArgs'] config: The node configuration of the pool.
        :param pulumi.Input[int] initial_node_count: The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
        :param pulumi.Input['NodeManagementArgs'] management: NodeManagement configuration for this NodePool.
        :param pulumi.Input['MaxPodsConstraintArgs'] max_pods_constraint: The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        :param pulumi.Input[str] name: The name of the node pool.
        :param pulumi.Input['NodeNetworkConfigArgs'] network_config: Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
        :param pulumi.Input[str] parent: The parent (project, location, cluster name) where the node pool will be created. Specified in the format `projects/*/locations/*/clusters/*`.
        :param pulumi.Input['PlacementPolicyArgs'] placement_policy: Specifies the node placement policy.
        :param pulumi.Input[str] project: Deprecated. The Google Developers Console [project ID or project number](https://developers.google.com/console/help/new/#projectnumber). This field has been deprecated and replaced by the parent field.
        :param pulumi.Input['UpgradeSettingsArgs'] upgrade_settings: Upgrade settings control disruption and speed of the upgrade.
        :param pulumi.Input[str] version: The version of the Kubernetes of this node.
        :param pulumi.Input[str] zone: Deprecated. The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/zones#available) in which the cluster resides. This field has been deprecated and replaced by the parent field.
        """
        if cluster_id is not None:
            warnings.warn("""Required. Deprecated. The name of the cluster. This field has been deprecated and replaced by the parent field.""", DeprecationWarning)
            pulumi.log.warn("""cluster_id is deprecated: Required. Deprecated. The name of the cluster. This field has been deprecated and replaced by the parent field.""")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if autoscaling is not None:
            pulumi.set(__self__, "autoscaling", autoscaling)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if initial_node_count is not None:
            pulumi.set(__self__, "initial_node_count", initial_node_count)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if locations is not None:
            pulumi.set(__self__, "locations", locations)
        if management is not None:
            pulumi.set(__self__, "management", management)
        if max_pods_constraint is not None:
            pulumi.set(__self__, "max_pods_constraint", max_pods_constraint)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_config is not None:
            pulumi.set(__self__, "network_config", network_config)
        if parent is not None:
            pulumi.set(__self__, "parent", parent)
        if placement_policy is not None:
            pulumi.set(__self__, "placement_policy", placement_policy)
        if project is not None:
            warnings.warn("""Required. Deprecated. The Google Developers Console [project ID or project number](https://developers.google.com/console/help/new/#projectnumber). This field has been deprecated and replaced by the parent field.""", DeprecationWarning)
            pulumi.log.warn("""project is deprecated: Required. Deprecated. The Google Developers Console [project ID or project number](https://developers.google.com/console/help/new/#projectnumber). This field has been deprecated and replaced by the parent field.""")
        if project is not None:
            pulumi.set(__self__, "project", project)
        if upgrade_settings is not None:
            pulumi.set(__self__, "upgrade_settings", upgrade_settings)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if zone is not None:
            warnings.warn("""Required. Deprecated. The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/zones#available) in which the cluster resides. This field has been deprecated and replaced by the parent field.""", DeprecationWarning)
            pulumi.log.warn("""zone is deprecated: Required. Deprecated. The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/zones#available) in which the cluster resides. This field has been deprecated and replaced by the parent field.""")
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Deprecated. The name of the cluster. This field has been deprecated and replaced by the parent field.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def autoscaling(self) -> Optional[pulumi.Input['NodePoolAutoscalingArgs']]:
        """
        Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
        """
        return pulumi.get(self, "autoscaling")

    @autoscaling.setter
    def autoscaling(self, value: Optional[pulumi.Input['NodePoolAutoscalingArgs']]):
        pulumi.set(self, "autoscaling", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StatusConditionArgs']]]]:
        """
        Which conditions caused the current node pool state.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StatusConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['NodeConfigArgs']]:
        """
        The node configuration of the pool.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['NodeConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="initialNodeCount")
    def initial_node_count(self) -> Optional[pulumi.Input[int]]:
        """
        The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
        """
        return pulumi.get(self, "initial_node_count")

    @initial_node_count.setter
    def initial_node_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "initial_node_count", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def locations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
        """
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter
    def management(self) -> Optional[pulumi.Input['NodeManagementArgs']]:
        """
        NodeManagement configuration for this NodePool.
        """
        return pulumi.get(self, "management")

    @management.setter
    def management(self, value: Optional[pulumi.Input['NodeManagementArgs']]):
        pulumi.set(self, "management", value)

    @property
    @pulumi.getter(name="maxPodsConstraint")
    def max_pods_constraint(self) -> Optional[pulumi.Input['MaxPodsConstraintArgs']]:
        """
        The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        """
        return pulumi.get(self, "max_pods_constraint")

    @max_pods_constraint.setter
    def max_pods_constraint(self, value: Optional[pulumi.Input['MaxPodsConstraintArgs']]):
        pulumi.set(self, "max_pods_constraint", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the node pool.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> Optional[pulumi.Input['NodeNetworkConfigArgs']]:
        """
        Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
        """
        return pulumi.get(self, "network_config")

    @network_config.setter
    def network_config(self, value: Optional[pulumi.Input['NodeNetworkConfigArgs']]):
        pulumi.set(self, "network_config", value)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        The parent (project, location, cluster name) where the node pool will be created. Specified in the format `projects/*/locations/*/clusters/*`.
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter(name="placementPolicy")
    def placement_policy(self) -> Optional[pulumi.Input['PlacementPolicyArgs']]:
        """
        Specifies the node placement policy.
        """
        return pulumi.get(self, "placement_policy")

    @placement_policy.setter
    def placement_policy(self, value: Optional[pulumi.Input['PlacementPolicyArgs']]):
        pulumi.set(self, "placement_policy", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        Deprecated. The Google Developers Console [project ID or project number](https://developers.google.com/console/help/new/#projectnumber). This field has been deprecated and replaced by the parent field.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="upgradeSettings")
    def upgrade_settings(self) -> Optional[pulumi.Input['UpgradeSettingsArgs']]:
        """
        Upgrade settings control disruption and speed of the upgrade.
        """
        return pulumi.get(self, "upgrade_settings")

    @upgrade_settings.setter
    def upgrade_settings(self, value: Optional[pulumi.Input['UpgradeSettingsArgs']]):
        pulumi.set(self, "upgrade_settings", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the Kubernetes of this node.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Deprecated. The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/zones#available) in which the cluster resides. This field has been deprecated and replaced by the parent field.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class NodePool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling: Optional[pulumi.Input[pulumi.InputType['NodePoolAutoscalingArgs']]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StatusConditionArgs']]]]] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['NodeConfigArgs']]] = None,
                 initial_node_count: Optional[pulumi.Input[int]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 management: Optional[pulumi.Input[pulumi.InputType['NodeManagementArgs']]] = None,
                 max_pods_constraint: Optional[pulumi.Input[pulumi.InputType['MaxPodsConstraintArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['NodeNetworkConfigArgs']]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 placement_policy: Optional[pulumi.Input[pulumi.InputType['PlacementPolicyArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 upgrade_settings: Optional[pulumi.Input[pulumi.InputType['UpgradeSettingsArgs']]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a node pool for a cluster.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NodePoolAutoscalingArgs']] autoscaling: Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
        :param pulumi.Input[str] cluster_id: Deprecated. The name of the cluster. This field has been deprecated and replaced by the parent field.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StatusConditionArgs']]]] conditions: Which conditions caused the current node pool state.
        :param pulumi.Input[pulumi.InputType['NodeConfigArgs']] config: The node configuration of the pool.
        :param pulumi.Input[int] initial_node_count: The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
        :param pulumi.Input[pulumi.InputType['NodeManagementArgs']] management: NodeManagement configuration for this NodePool.
        :param pulumi.Input[pulumi.InputType['MaxPodsConstraintArgs']] max_pods_constraint: The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        :param pulumi.Input[str] name: The name of the node pool.
        :param pulumi.Input[pulumi.InputType['NodeNetworkConfigArgs']] network_config: Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
        :param pulumi.Input[str] parent: The parent (project, location, cluster name) where the node pool will be created. Specified in the format `projects/*/locations/*/clusters/*`.
        :param pulumi.Input[pulumi.InputType['PlacementPolicyArgs']] placement_policy: Specifies the node placement policy.
        :param pulumi.Input[str] project: Deprecated. The Google Developers Console [project ID or project number](https://developers.google.com/console/help/new/#projectnumber). This field has been deprecated and replaced by the parent field.
        :param pulumi.Input[pulumi.InputType['UpgradeSettingsArgs']] upgrade_settings: Upgrade settings control disruption and speed of the upgrade.
        :param pulumi.Input[str] version: The version of the Kubernetes of this node.
        :param pulumi.Input[str] zone: Deprecated. The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/zones#available) in which the cluster resides. This field has been deprecated and replaced by the parent field.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NodePoolInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a node pool for a cluster.

        :param str resource_name: The name of the resource.
        :param NodePoolInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NodePoolInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling: Optional[pulumi.Input[pulumi.InputType['NodePoolAutoscalingArgs']]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StatusConditionArgs']]]]] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['NodeConfigArgs']]] = None,
                 initial_node_count: Optional[pulumi.Input[int]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 management: Optional[pulumi.Input[pulumi.InputType['NodeManagementArgs']]] = None,
                 max_pods_constraint: Optional[pulumi.Input[pulumi.InputType['MaxPodsConstraintArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['NodeNetworkConfigArgs']]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 placement_policy: Optional[pulumi.Input[pulumi.InputType['PlacementPolicyArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 upgrade_settings: Optional[pulumi.Input[pulumi.InputType['UpgradeSettingsArgs']]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NodePoolInitArgs.__new__(NodePoolInitArgs)

            __props__.__dict__["autoscaling"] = autoscaling
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            if cluster_id is not None and not opts.urn:
                warnings.warn("""Required. Deprecated. The name of the cluster. This field has been deprecated and replaced by the parent field.""", DeprecationWarning)
                pulumi.log.warn("""cluster_id is deprecated: Required. Deprecated. The name of the cluster. This field has been deprecated and replaced by the parent field.""")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["conditions"] = conditions
            __props__.__dict__["config"] = config
            __props__.__dict__["initial_node_count"] = initial_node_count
            __props__.__dict__["location"] = location
            __props__.__dict__["locations"] = locations
            __props__.__dict__["management"] = management
            __props__.__dict__["max_pods_constraint"] = max_pods_constraint
            __props__.__dict__["name"] = name
            __props__.__dict__["network_config"] = network_config
            __props__.__dict__["parent"] = parent
            __props__.__dict__["placement_policy"] = placement_policy
            if project is not None and not opts.urn:
                warnings.warn("""Required. Deprecated. The Google Developers Console [project ID or project number](https://developers.google.com/console/help/new/#projectnumber). This field has been deprecated and replaced by the parent field.""", DeprecationWarning)
                pulumi.log.warn("""project is deprecated: Required. Deprecated. The Google Developers Console [project ID or project number](https://developers.google.com/console/help/new/#projectnumber). This field has been deprecated and replaced by the parent field.""")
            __props__.__dict__["project"] = project
            __props__.__dict__["upgrade_settings"] = upgrade_settings
            __props__.__dict__["version"] = version
            if zone is not None and not opts.urn:
                warnings.warn("""Required. Deprecated. The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/zones#available) in which the cluster resides. This field has been deprecated and replaced by the parent field.""", DeprecationWarning)
                pulumi.log.warn("""zone is deprecated: Required. Deprecated. The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/zones#available) in which the cluster resides. This field has been deprecated and replaced by the parent field.""")
            __props__.__dict__["zone"] = zone
            __props__.__dict__["instance_group_urls"] = None
            __props__.__dict__["pod_ipv4_cidr_size"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["status_message"] = None
        super(NodePool, __self__).__init__(
            'google-native:container/v1beta1:NodePool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NodePool':
        """
        Get an existing NodePool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NodePoolInitArgs.__new__(NodePoolInitArgs)

        __props__.__dict__["autoscaling"] = None
        __props__.__dict__["conditions"] = None
        __props__.__dict__["config"] = None
        __props__.__dict__["initial_node_count"] = None
        __props__.__dict__["instance_group_urls"] = None
        __props__.__dict__["locations"] = None
        __props__.__dict__["management"] = None
        __props__.__dict__["max_pods_constraint"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_config"] = None
        __props__.__dict__["placement_policy"] = None
        __props__.__dict__["pod_ipv4_cidr_size"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["status_message"] = None
        __props__.__dict__["upgrade_settings"] = None
        __props__.__dict__["version"] = None
        return NodePool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def autoscaling(self) -> pulumi.Output['outputs.NodePoolAutoscalingResponse']:
        """
        Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
        """
        return pulumi.get(self, "autoscaling")

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Output[Sequence['outputs.StatusConditionResponse']]:
        """
        Which conditions caused the current node pool state.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.NodeConfigResponse']:
        """
        The node configuration of the pool.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="initialNodeCount")
    def initial_node_count(self) -> pulumi.Output[int]:
        """
        The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
        """
        return pulumi.get(self, "initial_node_count")

    @property
    @pulumi.getter(name="instanceGroupUrls")
    def instance_group_urls(self) -> pulumi.Output[Sequence[str]]:
        """
        [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
        """
        return pulumi.get(self, "instance_group_urls")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def management(self) -> pulumi.Output['outputs.NodeManagementResponse']:
        """
        NodeManagement configuration for this NodePool.
        """
        return pulumi.get(self, "management")

    @property
    @pulumi.getter(name="maxPodsConstraint")
    def max_pods_constraint(self) -> pulumi.Output['outputs.MaxPodsConstraintResponse']:
        """
        The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        """
        return pulumi.get(self, "max_pods_constraint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the node pool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> pulumi.Output['outputs.NodeNetworkConfigResponse']:
        """
        Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
        """
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter(name="placementPolicy")
    def placement_policy(self) -> pulumi.Output['outputs.PlacementPolicyResponse']:
        """
        Specifies the node placement policy.
        """
        return pulumi.get(self, "placement_policy")

    @property
    @pulumi.getter(name="podIpv4CidrSize")
    def pod_ipv4_cidr_size(self) -> pulumi.Output[int]:
        """
        [Output only] The pod CIDR block size per node in this node pool.
        """
        return pulumi.get(self, "pod_ipv4_cidr_size")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        [Output only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        [Output only] The status of the nodes in this pool instance.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> pulumi.Output[str]:
        """
        [Output only] Deprecated. Use conditions instead. Additional information about the current status of this node pool instance, if available.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter(name="upgradeSettings")
    def upgrade_settings(self) -> pulumi.Output['outputs.UpgradeSettingsResponse']:
        """
        Upgrade settings control disruption and speed of the upgrade.
        """
        return pulumi.get(self, "upgrade_settings")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The version of the Kubernetes of this node.
        """
        return pulumi.get(self, "version")

