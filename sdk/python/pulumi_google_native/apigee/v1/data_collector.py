# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = ['DataCollectorArgs', 'DataCollector']

@pulumi.input_type
class DataCollectorArgs:
    def __init__(__self__, *,
                 organization_id: pulumi.Input[str],
                 data_collector_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['DataCollectorType']] = None):
        """
        The set of arguments for constructing a DataCollector resource.
        :param pulumi.Input[str] data_collector_id: ID of the data collector. Overrides any ID in the data collector resource. Must be a string beginning with `dc_` that contains only letters, numbers, and underscores.
        :param pulumi.Input[str] description: A description of the data collector.
        :param pulumi.Input[str] name: ID of the data collector. Must begin with `dc_`.
        :param pulumi.Input['DataCollectorType'] type: Immutable. The type of data this data collector will collect.
        """
        pulumi.set(__self__, "organization_id", organization_id)
        if data_collector_id is not None:
            pulumi.set(__self__, "data_collector_id", data_collector_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="dataCollectorId")
    def data_collector_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the data collector. Overrides any ID in the data collector resource. Must be a string beginning with `dc_` that contains only letters, numbers, and underscores.
        """
        return pulumi.get(self, "data_collector_id")

    @data_collector_id.setter
    def data_collector_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_collector_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the data collector.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the data collector. Must begin with `dc_`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['DataCollectorType']]:
        """
        Immutable. The type of data this data collector will collect.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['DataCollectorType']]):
        pulumi.set(self, "type", value)


class DataCollector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_collector_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['DataCollectorType']] = None,
                 __props__=None):
        """
        Creates a new data collector.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_collector_id: ID of the data collector. Overrides any ID in the data collector resource. Must be a string beginning with `dc_` that contains only letters, numbers, and underscores.
        :param pulumi.Input[str] description: A description of the data collector.
        :param pulumi.Input[str] name: ID of the data collector. Must begin with `dc_`.
        :param pulumi.Input['DataCollectorType'] type: Immutable. The type of data this data collector will collect.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataCollectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new data collector.

        :param str resource_name: The name of the resource.
        :param DataCollectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataCollectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_collector_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['DataCollectorType']] = None,
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
            __props__ = DataCollectorArgs.__new__(DataCollectorArgs)

            __props__.__dict__["data_collector_id"] = data_collector_id
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["type"] = type
            __props__.__dict__["created_at"] = None
            __props__.__dict__["last_modified_at"] = None
        super(DataCollector, __self__).__init__(
            'google-native:apigee/v1:DataCollector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataCollector':
        """
        Get an existing DataCollector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DataCollectorArgs.__new__(DataCollectorArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["last_modified_at"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["type"] = None
        return DataCollector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The time at which the data collector was created in milliseconds since the epoch.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A description of the data collector.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> pulumi.Output[str]:
        """
        The time at which the Data Collector was last updated in milliseconds since the epoch.
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        ID of the data collector. Must begin with `dc_`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Immutable. The type of data this data collector will collect.
        """
        return pulumi.get(self, "type")

