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

__all__ = ['UserDataMappingArgs', 'UserDataMapping']

@pulumi.input_type
class UserDataMappingArgs:
    def __init__(__self__, *,
                 consent_store_id: pulumi.Input[str],
                 data_id: pulumi.Input[str],
                 dataset_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 project: pulumi.Input[str],
                 user_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 resource_attributes: Optional[pulumi.Input[Sequence[pulumi.Input['AttributeArgs']]]] = None):
        """
        The set of arguments for constructing a UserDataMapping resource.
        :param pulumi.Input[str] data_id: A unique identifier for the mapped resource.
        :param pulumi.Input[str] user_id: User's UUID provided by the client.
        :param pulumi.Input[str] name: Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
        :param pulumi.Input[Sequence[pulumi.Input['AttributeArgs']]] resource_attributes: Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
        """
        pulumi.set(__self__, "consent_store_id", consent_store_id)
        pulumi.set(__self__, "data_id", data_id)
        pulumi.set(__self__, "dataset_id", dataset_id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "user_id", user_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_attributes is not None:
            pulumi.set(__self__, "resource_attributes", resource_attributes)

    @property
    @pulumi.getter(name="consentStoreId")
    def consent_store_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "consent_store_id")

    @consent_store_id.setter
    def consent_store_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "consent_store_id", value)

    @property
    @pulumi.getter(name="dataId")
    def data_id(self) -> pulumi.Input[str]:
        """
        A unique identifier for the mapped resource.
        """
        return pulumi.get(self, "data_id")

    @data_id.setter
    def data_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_id", value)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "dataset_id")

    @dataset_id.setter
    def dataset_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset_id", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        User's UUID provided by the client.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceAttributes")
    def resource_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AttributeArgs']]]]:
        """
        Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
        """
        return pulumi.get(self, "resource_attributes")

    @resource_attributes.setter
    def resource_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AttributeArgs']]]]):
        pulumi.set(self, "resource_attributes", value)


class UserDataMapping(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consent_store_id: Optional[pulumi.Input[str]] = None,
                 data_id: Optional[pulumi.Input[str]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 resource_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AttributeArgs']]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new User data mapping in the parent consent store.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_id: A unique identifier for the mapped resource.
        :param pulumi.Input[str] name: Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AttributeArgs']]]] resource_attributes: Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
        :param pulumi.Input[str] user_id: User's UUID provided by the client.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserDataMappingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new User data mapping in the parent consent store.

        :param str resource_name: The name of the resource.
        :param UserDataMappingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserDataMappingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consent_store_id: Optional[pulumi.Input[str]] = None,
                 data_id: Optional[pulumi.Input[str]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 resource_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AttributeArgs']]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = UserDataMappingArgs.__new__(UserDataMappingArgs)

            if consent_store_id is None and not opts.urn:
                raise TypeError("Missing required property 'consent_store_id'")
            __props__.__dict__["consent_store_id"] = consent_store_id
            if data_id is None and not opts.urn:
                raise TypeError("Missing required property 'data_id'")
            __props__.__dict__["data_id"] = data_id
            if dataset_id is None and not opts.urn:
                raise TypeError("Missing required property 'dataset_id'")
            __props__.__dict__["dataset_id"] = dataset_id
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["resource_attributes"] = resource_attributes
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["archive_time"] = None
            __props__.__dict__["archived"] = None
        super(UserDataMapping, __self__).__init__(
            'google-native:healthcare/v1beta1:UserDataMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'UserDataMapping':
        """
        Get an existing UserDataMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UserDataMappingArgs.__new__(UserDataMappingArgs)

        __props__.__dict__["archive_time"] = None
        __props__.__dict__["archived"] = None
        __props__.__dict__["data_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["resource_attributes"] = None
        __props__.__dict__["user_id"] = None
        return UserDataMapping(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="archiveTime")
    def archive_time(self) -> pulumi.Output[str]:
        """
        Indicates the time when this mapping was archived.
        """
        return pulumi.get(self, "archive_time")

    @property
    @pulumi.getter
    def archived(self) -> pulumi.Output[bool]:
        """
        Indicates whether this mapping is archived.
        """
        return pulumi.get(self, "archived")

    @property
    @pulumi.getter(name="dataId")
    def data_id(self) -> pulumi.Output[str]:
        """
        A unique identifier for the mapped resource.
        """
        return pulumi.get(self, "data_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceAttributes")
    def resource_attributes(self) -> pulumi.Output[Sequence['outputs.AttributeResponse']]:
        """
        Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
        """
        return pulumi.get(self, "resource_attributes")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        User's UUID provided by the client.
        """
        return pulumi.get(self, "user_id")

