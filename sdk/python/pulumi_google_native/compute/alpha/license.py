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

__all__ = ['LicenseArgs', 'License']

@pulumi.input_type
class LicenseArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 license_code: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 resource_requirements: Optional[pulumi.Input['LicenseResourceRequirementsArgs']] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_id: Optional[pulumi.Input[str]] = None,
                 transferable: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a License resource.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional textual description of the resource; provided by the client when the resource is created.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output Only] Type of resource. Always compute#license for licenses.
        :param pulumi.Input[str] license_code: [Output Only] The unique code used to attach this license to images, snapshots, and disks.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        :param pulumi.Input[str] self_link_with_id: [Output Only] Server-defined URL for this resource with the resource id.
        :param pulumi.Input[bool] transferable: If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
        """
        pulumi.set(__self__, "project", project)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if license_code is not None:
            pulumi.set(__self__, "license_code", license_code)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if resource_requirements is not None:
            pulumi.set(__self__, "resource_requirements", resource_requirements)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id is not None:
            pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if transferable is not None:
            pulumi.set(__self__, "transferable", transferable)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

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
        An optional textual description of the resource; provided by the client when the resource is created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
        [Output Only] Type of resource. Always compute#license for licenses.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="licenseCode")
    def license_code(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The unique code used to attach this license to images, snapshots, and disks.
        """
        return pulumi.get(self, "license_code")

    @license_code.setter
    def license_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "license_code", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
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
    @pulumi.getter(name="resourceRequirements")
    def resource_requirements(self) -> Optional[pulumi.Input['LicenseResourceRequirementsArgs']]:
        return pulumi.get(self, "resource_requirements")

    @resource_requirements.setter
    def resource_requirements(self, value: Optional[pulumi.Input['LicenseResourceRequirementsArgs']]):
        pulumi.set(self, "resource_requirements", value)

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
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @self_link_with_id.setter
    def self_link_with_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link_with_id", value)

    @property
    @pulumi.getter
    def transferable(self) -> Optional[pulumi.Input[bool]]:
        """
        If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
        """
        return pulumi.get(self, "transferable")

    @transferable.setter
    def transferable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "transferable", value)


class License(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 license_code: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 resource_requirements: Optional[pulumi.Input[pulumi.InputType['LicenseResourceRequirementsArgs']]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_id: Optional[pulumi.Input[str]] = None,
                 transferable: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a License resource in the specified project.  Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional textual description of the resource; provided by the client when the resource is created.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output Only] Type of resource. Always compute#license for licenses.
        :param pulumi.Input[str] license_code: [Output Only] The unique code used to attach this license to images, snapshots, and disks.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        :param pulumi.Input[str] self_link_with_id: [Output Only] Server-defined URL for this resource with the resource id.
        :param pulumi.Input[bool] transferable: If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LicenseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a License resource in the specified project.  Caution This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.

        :param str resource_name: The name of the resource.
        :param LicenseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LicenseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 license_code: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 resource_requirements: Optional[pulumi.Input[pulumi.InputType['LicenseResourceRequirementsArgs']]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_id: Optional[pulumi.Input[str]] = None,
                 transferable: Optional[pulumi.Input[bool]] = None,
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
            __props__ = LicenseArgs.__new__(LicenseArgs)

            __props__.__dict__["creation_timestamp"] = creation_timestamp
            __props__.__dict__["description"] = description
            __props__.__dict__["id"] = id
            __props__.__dict__["kind"] = kind
            __props__.__dict__["license_code"] = license_code
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["resource_requirements"] = resource_requirements
            __props__.__dict__["self_link"] = self_link
            __props__.__dict__["self_link_with_id"] = self_link_with_id
            __props__.__dict__["transferable"] = transferable
        super(License, __self__).__init__(
            'google-native:compute/alpha:License',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'License':
        """
        Get an existing License resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LicenseArgs.__new__(LicenseArgs)

        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["license_code"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["resource_requirements"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["self_link_with_id"] = None
        __props__.__dict__["transferable"] = None
        return License(resource_name, opts=opts, __props__=__props__)

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
        An optional textual description of the resource; provided by the client when the resource is created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        [Output Only] Type of resource. Always compute#license for licenses.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="licenseCode")
    def license_code(self) -> pulumi.Output[str]:
        """
        [Output Only] The unique code used to attach this license to images, snapshots, and disks.
        """
        return pulumi.get(self, "license_code")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceRequirements")
    def resource_requirements(self) -> pulumi.Output['outputs.LicenseResourceRequirementsResponse']:
        return pulumi.get(self, "resource_requirements")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> pulumi.Output[str]:
        """
        [Output Only] Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter
    def transferable(self) -> pulumi.Output[bool]:
        """
        If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
        """
        return pulumi.get(self, "transferable")

