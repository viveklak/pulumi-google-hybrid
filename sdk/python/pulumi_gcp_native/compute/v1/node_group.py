# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['NodeGroupArgs', 'NodeGroup']

@pulumi.input_type
class NodeGroupArgs:
    def __init__(__self__, *,
                 initial_node_count: pulumi.Input[str],
                 node_group: pulumi.Input[str],
                 project: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 autoscaling_policy: Optional[pulumi.Input['NodeGroupAutoscalingPolicyArgs']] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location_hint: Optional[pulumi.Input[str]] = None,
                 maintenance_policy: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input['NodeGroupMaintenanceWindowArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_template: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NodeGroup resource.
        :param pulumi.Input[str] zone: [Output Only] The name of the zone where the node group resides, such as us-central1-a.
        :param pulumi.Input['NodeGroupAutoscalingPolicyArgs'] autoscaling_policy: Specifies how autoscaling should behave.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output Only] The type of the resource. Always compute#nodeGroup for node group.
        :param pulumi.Input[str] location_hint: An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
        :param pulumi.Input[str] maintenance_policy: Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see  Maintenance policies.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] node_template: URL of the node template to create the node group from.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        :param pulumi.Input[int] size: [Output Only] The total number of nodes in the node group.
        """
        pulumi.set(__self__, "initial_node_count", initial_node_count)
        pulumi.set(__self__, "node_group", node_group)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "zone", zone)
        if autoscaling_policy is not None:
            pulumi.set(__self__, "autoscaling_policy", autoscaling_policy)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if location_hint is not None:
            pulumi.set(__self__, "location_hint", location_hint)
        if maintenance_policy is not None:
            pulumi.set(__self__, "maintenance_policy", maintenance_policy)
        if maintenance_window is not None:
            pulumi.set(__self__, "maintenance_window", maintenance_window)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_template is not None:
            pulumi.set(__self__, "node_template", node_template)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="initialNodeCount")
    def initial_node_count(self) -> pulumi.Input[str]:
        return pulumi.get(self, "initial_node_count")

    @initial_node_count.setter
    def initial_node_count(self, value: pulumi.Input[str]):
        pulumi.set(self, "initial_node_count", value)

    @property
    @pulumi.getter(name="nodeGroup")
    def node_group(self) -> pulumi.Input[str]:
        return pulumi.get(self, "node_group")

    @node_group.setter
    def node_group(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_group", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        [Output Only] The name of the zone where the node group resides, such as us-central1-a.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> Optional[pulumi.Input['NodeGroupAutoscalingPolicyArgs']]:
        """
        Specifies how autoscaling should behave.
        """
        return pulumi.get(self, "autoscaling_policy")

    @autoscaling_policy.setter
    def autoscaling_policy(self, value: Optional[pulumi.Input['NodeGroupAutoscalingPolicyArgs']]):
        pulumi.set(self, "autoscaling_policy", value)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The type of the resource. Always compute#nodeGroup for node group.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="locationHint")
    def location_hint(self) -> Optional[pulumi.Input[str]]:
        """
        An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
        """
        return pulumi.get(self, "location_hint")

    @location_hint.setter
    def location_hint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location_hint", value)

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see  Maintenance policies.
        """
        return pulumi.get(self, "maintenance_policy")

    @maintenance_policy.setter
    def maintenance_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "maintenance_policy", value)

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> Optional[pulumi.Input['NodeGroupMaintenanceWindowArgs']]:
        return pulumi.get(self, "maintenance_window")

    @maintenance_window.setter
    def maintenance_window(self, value: Optional[pulumi.Input['NodeGroupMaintenanceWindowArgs']]):
        pulumi.set(self, "maintenance_window", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeTemplate")
    def node_template(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the node template to create the node group from.
        """
        return pulumi.get(self, "node_template")

    @node_template.setter
    def node_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_template", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        [Output Only] The total number of nodes in the node group.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class NodeGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['NodeGroupAutoscalingPolicyArgs']]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 initial_node_count: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location_hint: Optional[pulumi.Input[str]] = None,
                 maintenance_policy: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input[pulumi.InputType['NodeGroupMaintenanceWindowArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_group: Optional[pulumi.Input[str]] = None,
                 node_template: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a NodeGroup resource in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NodeGroupAutoscalingPolicyArgs']] autoscaling_policy: Specifies how autoscaling should behave.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output Only] The type of the resource. Always compute#nodeGroup for node group.
        :param pulumi.Input[str] location_hint: An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
        :param pulumi.Input[str] maintenance_policy: Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see  Maintenance policies.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] node_template: URL of the node template to create the node group from.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        :param pulumi.Input[int] size: [Output Only] The total number of nodes in the node group.
        :param pulumi.Input[str] zone: [Output Only] The name of the zone where the node group resides, such as us-central1-a.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NodeGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a NodeGroup resource in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param NodeGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NodeGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['NodeGroupAutoscalingPolicyArgs']]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 initial_node_count: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location_hint: Optional[pulumi.Input[str]] = None,
                 maintenance_policy: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input[pulumi.InputType['NodeGroupMaintenanceWindowArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_group: Optional[pulumi.Input[str]] = None,
                 node_template: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
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
            __props__ = NodeGroupArgs.__new__(NodeGroupArgs)

            __props__.__dict__["autoscaling_policy"] = autoscaling_policy
            __props__.__dict__["creation_timestamp"] = creation_timestamp
            __props__.__dict__["description"] = description
            __props__.__dict__["fingerprint"] = fingerprint
            __props__.__dict__["id"] = id
            if initial_node_count is None and not opts.urn:
                raise TypeError("Missing required property 'initial_node_count'")
            __props__.__dict__["initial_node_count"] = initial_node_count
            __props__.__dict__["kind"] = kind
            __props__.__dict__["location_hint"] = location_hint
            __props__.__dict__["maintenance_policy"] = maintenance_policy
            __props__.__dict__["maintenance_window"] = maintenance_window
            __props__.__dict__["name"] = name
            if node_group is None and not opts.urn:
                raise TypeError("Missing required property 'node_group'")
            __props__.__dict__["node_group"] = node_group
            __props__.__dict__["node_template"] = node_template
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["self_link"] = self_link
            __props__.__dict__["size"] = size
            __props__.__dict__["status"] = status
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
        super(NodeGroup, __self__).__init__(
            'gcp-native:compute/v1:NodeGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NodeGroup':
        """
        Get an existing NodeGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NodeGroupArgs.__new__(NodeGroupArgs)

        __props__.__dict__["autoscaling_policy"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["fingerprint"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["location_hint"] = None
        __props__.__dict__["maintenance_policy"] = None
        __props__.__dict__["maintenance_window"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["node_template"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["size"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["zone"] = None
        return NodeGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> pulumi.Output['outputs.NodeGroupAutoscalingPolicyResponse']:
        """
        Specifies how autoscaling should behave.
        """
        return pulumi.get(self, "autoscaling_policy")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        [Output Only] The type of the resource. Always compute#nodeGroup for node group.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="locationHint")
    def location_hint(self) -> pulumi.Output[str]:
        """
        An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
        """
        return pulumi.get(self, "location_hint")

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> pulumi.Output[str]:
        """
        Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see  Maintenance policies.
        """
        return pulumi.get(self, "maintenance_policy")

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> pulumi.Output['outputs.NodeGroupMaintenanceWindowResponse']:
        return pulumi.get(self, "maintenance_window")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeTemplate")
    def node_template(self) -> pulumi.Output[str]:
        """
        URL of the node template to create the node group from.
        """
        return pulumi.get(self, "node_template")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        [Output Only] The total number of nodes in the node group.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        [Output Only] The name of the zone where the node group resides, such as us-central1-a.
        """
        return pulumi.get(self, "zone")

