# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['OrganizationEnvgroupAttachmentArgs', 'OrganizationEnvgroupAttachment']

@pulumi.input_type
class OrganizationEnvgroupAttachmentArgs:
    def __init__(__self__, *,
                 envgroup_id: pulumi.Input[str],
                 organization_id: pulumi.Input[str],
                 environment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrganizationEnvgroupAttachment resource.
        :param pulumi.Input[str] environment: Required. ID of the attached environment.
        :param pulumi.Input[str] name: ID of the environment group attachment.
        """
        pulumi.set(__self__, "envgroup_id", envgroup_id)
        pulumi.set(__self__, "organization_id", organization_id)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="envgroupId")
    def envgroup_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "envgroup_id")

    @envgroup_id.setter
    def envgroup_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "envgroup_id", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        """
        Required. ID of the attached environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the environment group attachment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class OrganizationEnvgroupAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 envgroup_id: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new attachment of an environment to an environment group.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] environment: Required. ID of the attached environment.
        :param pulumi.Input[str] name: ID of the environment group attachment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationEnvgroupAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new attachment of an environment to an environment group.

        :param str resource_name: The name of the resource.
        :param OrganizationEnvgroupAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationEnvgroupAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 envgroup_id: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = OrganizationEnvgroupAttachmentArgs.__new__(OrganizationEnvgroupAttachmentArgs)

            if envgroup_id is None and not opts.urn:
                raise TypeError("Missing required property 'envgroup_id'")
            __props__.__dict__["envgroup_id"] = envgroup_id
            __props__.__dict__["environment"] = environment
            __props__.__dict__["name"] = name
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["created_at"] = None
        super(OrganizationEnvgroupAttachment, __self__).__init__(
            'google-native:apigee/v1:OrganizationEnvgroupAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OrganizationEnvgroupAttachment':
        """
        Get an existing OrganizationEnvgroupAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OrganizationEnvgroupAttachmentArgs.__new__(OrganizationEnvgroupAttachmentArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["environment"] = None
        __props__.__dict__["name"] = None
        return OrganizationEnvgroupAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The time at which the environment group attachment was created as milliseconds since epoch.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        """
        Required. ID of the attached environment.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        ID of the environment group attachment.
        """
        return pulumi.get(self, "name")

