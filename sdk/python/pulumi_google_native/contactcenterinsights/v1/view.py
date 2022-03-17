# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['ViewArgs', 'View']

@pulumi.input_type
class ViewArgs:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a View resource.
        :param pulumi.Input[str] display_name: The human-readable display name of the view.
        :param pulumi.Input[str] name: Immutable. The resource name of the view. Format: projects/{project}/locations/{location}/views/{view}
        :param pulumi.Input[str] value: String with specific view properties.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable display name of the view.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The resource name of the view. Format: projects/{project}/locations/{location}/views/{view}
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        String with specific view properties.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class View(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a view.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The human-readable display name of the view.
        :param pulumi.Input[str] name: Immutable. The resource name of the view. Format: projects/{project}/locations/{location}/views/{view}
        :param pulumi.Input[str] value: String with specific view properties.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ViewArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a view.

        :param str resource_name: The name of the resource.
        :param ViewArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ViewArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
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
            __props__ = ViewArgs.__new__(ViewArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["value"] = value
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        super(View, __self__).__init__(
            'google-native:contactcenterinsights/v1:View',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'View':
        """
        Get an existing View resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ViewArgs.__new__(ViewArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["value"] = None
        return View(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time at which this view was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The human-readable display name of the view.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Immutable. The resource name of the view. Format: projects/{project}/locations/{location}/views/{view}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The most recent time at which the view was updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        String with specific view properties.
        """
        return pulumi.get(self, "value")

