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

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 node_config: pulumi.Input['NodeConfigArgs'],
                 node_count: pulumi.Input[int],
                 authorized_network: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 instance_messages: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceMessageArgs']]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_policy: Optional[pulumi.Input['GoogleCloudMemcacheV1MaintenancePolicyArgs']] = None,
                 memcache_version: Optional[pulumi.Input['InstanceMemcacheVersion']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input['MemcacheParametersArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] instance_id: Required. The logical name of the Memcached instance in the user project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-40 characters. * Must end with a number or a letter. * Must be unique within the user project / location. If any of the above are not met, the API raises an invalid argument error.
        :param pulumi.Input['NodeConfigArgs'] node_config: Configuration for Memcached nodes.
        :param pulumi.Input[int] node_count: Number of nodes in the Memcached instance.
        :param pulumi.Input[str] authorized_network: The full name of the Google Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. If left unspecified, the `default` network will be used.
        :param pulumi.Input[str] display_name: User provided name for the instance, which is only used for display purposes. Cannot be more than 80 characters.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceMessageArgs']]] instance_messages: List of messages that describe the current state of the Memcached instance.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        :param pulumi.Input['GoogleCloudMemcacheV1MaintenancePolicyArgs'] maintenance_policy: The maintenance policy for the instance. If not provided, the maintenance event will be performed based on Memorystore internal rollout schedule.
        :param pulumi.Input['InstanceMemcacheVersion'] memcache_version: The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
        :param pulumi.Input[str] name: Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.
        :param pulumi.Input['MemcacheParametersArgs'] parameters: User defined parameters to apply to the memcached process on each node.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: Zones in which Memcached nodes should be provisioned. Memcached nodes will be equally distributed across these zones. If not provided, the service will by default create nodes in all zones in the region for the instance.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "node_config", node_config)
        pulumi.set(__self__, "node_count", node_count)
        if authorized_network is not None:
            pulumi.set(__self__, "authorized_network", authorized_network)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if instance_messages is not None:
            pulumi.set(__self__, "instance_messages", instance_messages)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if maintenance_policy is not None:
            pulumi.set(__self__, "maintenance_policy", maintenance_policy)
        if memcache_version is not None:
            pulumi.set(__self__, "memcache_version", memcache_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if zones is not None:
            pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Required. The logical name of the Memcached instance in the user project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-40 characters. * Must end with a number or a letter. * Must be unique within the user project / location. If any of the above are not met, the API raises an invalid argument error.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> pulumi.Input['NodeConfigArgs']:
        """
        Configuration for Memcached nodes.
        """
        return pulumi.get(self, "node_config")

    @node_config.setter
    def node_config(self, value: pulumi.Input['NodeConfigArgs']):
        pulumi.set(self, "node_config", value)

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Input[int]:
        """
        Number of nodes in the Memcached instance.
        """
        return pulumi.get(self, "node_count")

    @node_count.setter
    def node_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "node_count", value)

    @property
    @pulumi.getter(name="authorizedNetwork")
    def authorized_network(self) -> Optional[pulumi.Input[str]]:
        """
        The full name of the Google Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. If left unspecified, the `default` network will be used.
        """
        return pulumi.get(self, "authorized_network")

    @authorized_network.setter
    def authorized_network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorized_network", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        User provided name for the instance, which is only used for display purposes. Cannot be more than 80 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="instanceMessages")
    def instance_messages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceMessageArgs']]]]:
        """
        List of messages that describe the current state of the Memcached instance.
        """
        return pulumi.get(self, "instance_messages")

    @instance_messages.setter
    def instance_messages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceMessageArgs']]]]):
        pulumi.set(self, "instance_messages", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> Optional[pulumi.Input['GoogleCloudMemcacheV1MaintenancePolicyArgs']]:
        """
        The maintenance policy for the instance. If not provided, the maintenance event will be performed based on Memorystore internal rollout schedule.
        """
        return pulumi.get(self, "maintenance_policy")

    @maintenance_policy.setter
    def maintenance_policy(self, value: Optional[pulumi.Input['GoogleCloudMemcacheV1MaintenancePolicyArgs']]):
        pulumi.set(self, "maintenance_policy", value)

    @property
    @pulumi.getter(name="memcacheVersion")
    def memcache_version(self) -> Optional[pulumi.Input['InstanceMemcacheVersion']]:
        """
        The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
        """
        return pulumi.get(self, "memcache_version")

    @memcache_version.setter
    def memcache_version(self, value: Optional[pulumi.Input['InstanceMemcacheVersion']]):
        pulumi.set(self, "memcache_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input['MemcacheParametersArgs']]:
        """
        User defined parameters to apply to the memcached process on each node.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input['MemcacheParametersArgs']]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Zones in which Memcached nodes should be provisioned. Memcached nodes will be equally distributed across these zones. If not provided, the service will by default create nodes in all zones in the region for the instance.
        """
        return pulumi.get(self, "zones")

    @zones.setter
    def zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "zones", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized_network: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_messages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceMessageArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_policy: Optional[pulumi.Input[pulumi.InputType['GoogleCloudMemcacheV1MaintenancePolicyArgs']]] = None,
                 memcache_version: Optional[pulumi.Input['InstanceMemcacheVersion']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_config: Optional[pulumi.Input[pulumi.InputType['NodeConfigArgs']]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 parameters: Optional[pulumi.Input[pulumi.InputType['MemcacheParametersArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Creates a new Instance in a given location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorized_network: The full name of the Google Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. If left unspecified, the `default` network will be used.
        :param pulumi.Input[str] display_name: User provided name for the instance, which is only used for display purposes. Cannot be more than 80 characters.
        :param pulumi.Input[str] instance_id: Required. The logical name of the Memcached instance in the user project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-40 characters. * Must end with a number or a letter. * Must be unique within the user project / location. If any of the above are not met, the API raises an invalid argument error.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceMessageArgs']]]] instance_messages: List of messages that describe the current state of the Memcached instance.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        :param pulumi.Input[pulumi.InputType['GoogleCloudMemcacheV1MaintenancePolicyArgs']] maintenance_policy: The maintenance policy for the instance. If not provided, the maintenance event will be performed based on Memorystore internal rollout schedule.
        :param pulumi.Input['InstanceMemcacheVersion'] memcache_version: The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
        :param pulumi.Input[str] name: Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.
        :param pulumi.Input[pulumi.InputType['NodeConfigArgs']] node_config: Configuration for Memcached nodes.
        :param pulumi.Input[int] node_count: Number of nodes in the Memcached instance.
        :param pulumi.Input[pulumi.InputType['MemcacheParametersArgs']] parameters: User defined parameters to apply to the memcached process on each node.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: Zones in which Memcached nodes should be provisioned. Memcached nodes will be equally distributed across these zones. If not provided, the service will by default create nodes in all zones in the region for the instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Instance in a given location.

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized_network: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_messages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceMessageArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_policy: Optional[pulumi.Input[pulumi.InputType['GoogleCloudMemcacheV1MaintenancePolicyArgs']]] = None,
                 memcache_version: Optional[pulumi.Input['InstanceMemcacheVersion']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_config: Optional[pulumi.Input[pulumi.InputType['NodeConfigArgs']]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 parameters: Optional[pulumi.Input[pulumi.InputType['MemcacheParametersArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["authorized_network"] = authorized_network
            __props__.__dict__["display_name"] = display_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["instance_messages"] = instance_messages
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["maintenance_policy"] = maintenance_policy
            __props__.__dict__["memcache_version"] = memcache_version
            __props__.__dict__["name"] = name
            if node_config is None and not opts.urn:
                raise TypeError("Missing required property 'node_config'")
            __props__.__dict__["node_config"] = node_config
            if node_count is None and not opts.urn:
                raise TypeError("Missing required property 'node_count'")
            __props__.__dict__["node_count"] = node_count
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["project"] = project
            __props__.__dict__["zones"] = zones
            __props__.__dict__["create_time"] = None
            __props__.__dict__["discovery_endpoint"] = None
            __props__.__dict__["maintenance_schedule"] = None
            __props__.__dict__["memcache_full_version"] = None
            __props__.__dict__["memcache_nodes"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
        super(Instance, __self__).__init__(
            'google-native:memcache/v1:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InstanceArgs.__new__(InstanceArgs)

        __props__.__dict__["authorized_network"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["discovery_endpoint"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["instance_messages"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["maintenance_policy"] = None
        __props__.__dict__["maintenance_schedule"] = None
        __props__.__dict__["memcache_full_version"] = None
        __props__.__dict__["memcache_nodes"] = None
        __props__.__dict__["memcache_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["node_config"] = None
        __props__.__dict__["node_count"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["zones"] = None
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizedNetwork")
    def authorized_network(self) -> pulumi.Output[str]:
        """
        The full name of the Google Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. If left unspecified, the `default` network will be used.
        """
        return pulumi.get(self, "authorized_network")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time the instance was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="discoveryEndpoint")
    def discovery_endpoint(self) -> pulumi.Output[str]:
        """
        Endpoint for the Discovery API.
        """
        return pulumi.get(self, "discovery_endpoint")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        User provided name for the instance, which is only used for display purposes. Cannot be more than 80 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="instanceMessages")
    def instance_messages(self) -> pulumi.Output[Sequence['outputs.InstanceMessageResponse']]:
        """
        List of messages that describe the current state of the Memcached instance.
        """
        return pulumi.get(self, "instance_messages")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> pulumi.Output['outputs.GoogleCloudMemcacheV1MaintenancePolicyResponse']:
        """
        The maintenance policy for the instance. If not provided, the maintenance event will be performed based on Memorystore internal rollout schedule.
        """
        return pulumi.get(self, "maintenance_policy")

    @property
    @pulumi.getter(name="maintenanceSchedule")
    def maintenance_schedule(self) -> pulumi.Output['outputs.MaintenanceScheduleResponse']:
        """
        Published maintenance schedule.
        """
        return pulumi.get(self, "maintenance_schedule")

    @property
    @pulumi.getter(name="memcacheFullVersion")
    def memcache_full_version(self) -> pulumi.Output[str]:
        """
        The full version of memcached server running on this instance. System automatically determines the full memcached version for an instance based on the input MemcacheVersion. The full version format will be "memcached-1.5.16".
        """
        return pulumi.get(self, "memcache_full_version")

    @property
    @pulumi.getter(name="memcacheNodes")
    def memcache_nodes(self) -> pulumi.Output[Sequence['outputs.NodeResponse']]:
        """
        List of Memcached nodes. Refer to Node message for more details.
        """
        return pulumi.get(self, "memcache_nodes")

    @property
    @pulumi.getter(name="memcacheVersion")
    def memcache_version(self) -> pulumi.Output[str]:
        """
        The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
        """
        return pulumi.get(self, "memcache_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> pulumi.Output['outputs.NodeConfigResponse']:
        """
        Configuration for Memcached nodes.
        """
        return pulumi.get(self, "node_config")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Output[int]:
        """
        Number of nodes in the Memcached instance.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output['outputs.MemcacheParametersResponse']:
        """
        User defined parameters to apply to the memcached process on each node.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of this Memcached instance.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time the instance was updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Sequence[str]]:
        """
        Zones in which Memcached nodes should be provisioned. Memcached nodes will be equally distributed across these zones. If not provided, the service will by default create nodes in all zones in the region for the instance.
        """
        return pulumi.get(self, "zones")

