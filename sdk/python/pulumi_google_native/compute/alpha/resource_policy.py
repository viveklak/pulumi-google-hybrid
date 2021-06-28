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

__all__ = ['ResourcePolicyArgs', 'ResourcePolicy']

@pulumi.input_type
class ResourcePolicyArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 region: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 group_placement_policy: Optional[pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs']] = None,
                 instance_schedule_policy: Optional[pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 snapshot_schedule_policy: Optional[pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs']] = None,
                 vm_maintenance_policy: Optional[pulumi.Input['ResourcePolicyVmMaintenancePolicyArgs']] = None):
        """
        The set of arguments for constructing a ResourcePolicy resource.
        :param pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs'] group_placement_policy: Resource policy for instances for placement configuration.
        :param pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs'] instance_schedule_policy: Resource policy for scheduling instance operations.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs'] snapshot_schedule_policy: Resource policy for persistent disks for creating snapshots.
        :param pulumi.Input['ResourcePolicyVmMaintenancePolicyArgs'] vm_maintenance_policy: Resource policy applicable to VMs for infrastructure maintenance.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "region", region)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if group_placement_policy is not None:
            pulumi.set(__self__, "group_placement_policy", group_placement_policy)
        if instance_schedule_policy is not None:
            pulumi.set(__self__, "instance_schedule_policy", instance_schedule_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if snapshot_schedule_policy is not None:
            pulumi.set(__self__, "snapshot_schedule_policy", snapshot_schedule_policy)
        if vm_maintenance_policy is not None:
            pulumi.set(__self__, "vm_maintenance_policy", vm_maintenance_policy)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="groupPlacementPolicy")
    def group_placement_policy(self) -> Optional[pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs']]:
        """
        Resource policy for instances for placement configuration.
        """
        return pulumi.get(self, "group_placement_policy")

    @group_placement_policy.setter
    def group_placement_policy(self, value: Optional[pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs']]):
        pulumi.set(self, "group_placement_policy", value)

    @property
    @pulumi.getter(name="instanceSchedulePolicy")
    def instance_schedule_policy(self) -> Optional[pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs']]:
        """
        Resource policy for scheduling instance operations.
        """
        return pulumi.get(self, "instance_schedule_policy")

    @instance_schedule_policy.setter
    def instance_schedule_policy(self, value: Optional[pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs']]):
        pulumi.set(self, "instance_schedule_policy", value)

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
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="snapshotSchedulePolicy")
    def snapshot_schedule_policy(self) -> Optional[pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs']]:
        """
        Resource policy for persistent disks for creating snapshots.
        """
        return pulumi.get(self, "snapshot_schedule_policy")

    @snapshot_schedule_policy.setter
    def snapshot_schedule_policy(self, value: Optional[pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs']]):
        pulumi.set(self, "snapshot_schedule_policy", value)

    @property
    @pulumi.getter(name="vmMaintenancePolicy")
    def vm_maintenance_policy(self) -> Optional[pulumi.Input['ResourcePolicyVmMaintenancePolicyArgs']]:
        """
        Resource policy applicable to VMs for infrastructure maintenance.
        """
        return pulumi.get(self, "vm_maintenance_policy")

    @vm_maintenance_policy.setter
    def vm_maintenance_policy(self, value: Optional[pulumi.Input['ResourcePolicyVmMaintenancePolicyArgs']]):
        pulumi.set(self, "vm_maintenance_policy", value)


class ResourcePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_placement_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyGroupPlacementPolicyArgs']]] = None,
                 instance_schedule_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyInstanceSchedulePolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 snapshot_schedule_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicySnapshotSchedulePolicyArgs']]] = None,
                 vm_maintenance_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyVmMaintenancePolicyArgs']]] = None,
                 __props__=None):
        """
        Creates a new resource policy.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ResourcePolicyGroupPlacementPolicyArgs']] group_placement_policy: Resource policy for instances for placement configuration.
        :param pulumi.Input[pulumi.InputType['ResourcePolicyInstanceSchedulePolicyArgs']] instance_schedule_policy: Resource policy for scheduling instance operations.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[pulumi.InputType['ResourcePolicySnapshotSchedulePolicyArgs']] snapshot_schedule_policy: Resource policy for persistent disks for creating snapshots.
        :param pulumi.Input[pulumi.InputType['ResourcePolicyVmMaintenancePolicyArgs']] vm_maintenance_policy: Resource policy applicable to VMs for infrastructure maintenance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourcePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new resource policy.

        :param str resource_name: The name of the resource.
        :param ResourcePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourcePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_placement_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyGroupPlacementPolicyArgs']]] = None,
                 instance_schedule_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyInstanceSchedulePolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 snapshot_schedule_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicySnapshotSchedulePolicyArgs']]] = None,
                 vm_maintenance_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyVmMaintenancePolicyArgs']]] = None,
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
            __props__ = ResourcePolicyArgs.__new__(ResourcePolicyArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["group_placement_policy"] = group_placement_policy
            __props__.__dict__["instance_schedule_policy"] = instance_schedule_policy
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["snapshot_schedule_policy"] = snapshot_schedule_policy
            __props__.__dict__["vm_maintenance_policy"] = vm_maintenance_policy
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["resource_status"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["self_link_with_id"] = None
            __props__.__dict__["status"] = None
        super(ResourcePolicy, __self__).__init__(
            'google-native:compute/alpha:ResourcePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ResourcePolicy':
        """
        Get an existing ResourcePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ResourcePolicyArgs.__new__(ResourcePolicyArgs)

        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["group_placement_policy"] = None
        __props__.__dict__["instance_schedule_policy"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["resource_status"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["self_link_with_id"] = None
        __props__.__dict__["snapshot_schedule_policy"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["vm_maintenance_policy"] = None
        return ResourcePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="groupPlacementPolicy")
    def group_placement_policy(self) -> pulumi.Output['outputs.ResourcePolicyGroupPlacementPolicyResponse']:
        """
        Resource policy for instances for placement configuration.
        """
        return pulumi.get(self, "group_placement_policy")

    @property
    @pulumi.getter(name="instanceSchedulePolicy")
    def instance_schedule_policy(self) -> pulumi.Output['outputs.ResourcePolicyInstanceSchedulePolicyResponse']:
        """
        Resource policy for scheduling instance operations.
        """
        return pulumi.get(self, "instance_schedule_policy")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Type of the resource. Always compute#resource_policies for resource policies.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="resourceStatus")
    def resource_status(self) -> pulumi.Output['outputs.ResourcePolicyResourceStatusResponse']:
        """
        The system status of the resource policy.
        """
        return pulumi.get(self, "resource_status")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined fully-qualified URL for this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> pulumi.Output[str]:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter(name="snapshotSchedulePolicy")
    def snapshot_schedule_policy(self) -> pulumi.Output['outputs.ResourcePolicySnapshotSchedulePolicyResponse']:
        """
        Resource policy for persistent disks for creating snapshots.
        """
        return pulumi.get(self, "snapshot_schedule_policy")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of resource policy creation.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vmMaintenancePolicy")
    def vm_maintenance_policy(self) -> pulumi.Output['outputs.ResourcePolicyVmMaintenancePolicyResponse']:
        """
        Resource policy applicable to VMs for infrastructure maintenance.
        """
        return pulumi.get(self, "vm_maintenance_policy")

