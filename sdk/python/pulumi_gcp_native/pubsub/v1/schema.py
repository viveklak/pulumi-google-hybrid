# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['SchemaArgs', 'Schema']

@pulumi.input_type
class SchemaArgs:
    def __init__(__self__, *,
                 projects_id: pulumi.Input[str],
                 schemas_id: pulumi.Input[str],
                 definition: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Schema resource.
        :param pulumi.Input[str] definition: The definition of the schema. This should contain a string representing the full definition of the schema that is a valid schema definition of the type specified in `type`.
        :param pulumi.Input[str] name: Required. Name of the schema. Format is `projects/{project}/schemas/{schema}`.
        :param pulumi.Input[str] type: The type of the schema definition.
        """
        pulumi.set(__self__, "projects_id", projects_id)
        pulumi.set(__self__, "schemas_id", schemas_id)
        if definition is not None:
            pulumi.set(__self__, "definition", definition)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter(name="schemasId")
    def schemas_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "schemas_id")

    @schemas_id.setter
    def schemas_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "schemas_id", value)

    @property
    @pulumi.getter
    def definition(self) -> Optional[pulumi.Input[str]]:
        """
        The definition of the schema. This should contain a string representing the full definition of the schema that is a valid schema definition of the type specified in `type`.
        """
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. Name of the schema. Format is `projects/{project}/schemas/{schema}`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the schema definition.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Schema(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 schemas_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a schema.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] definition: The definition of the schema. This should contain a string representing the full definition of the schema that is a valid schema definition of the type specified in `type`.
        :param pulumi.Input[str] name: Required. Name of the schema. Format is `projects/{project}/schemas/{schema}`.
        :param pulumi.Input[str] type: The type of the schema definition.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SchemaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a schema.

        :param str resource_name: The name of the resource.
        :param SchemaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SchemaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 schemas_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = SchemaArgs.__new__(SchemaArgs)

            __props__.__dict__["definition"] = definition
            __props__.__dict__["name"] = name
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            if schemas_id is None and not opts.urn:
                raise TypeError("Missing required property 'schemas_id'")
            __props__.__dict__["schemas_id"] = schemas_id
            __props__.__dict__["type"] = type
        super(Schema, __self__).__init__(
            'gcp-native:pubsub/v1:Schema',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Schema':
        """
        Get an existing Schema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SchemaArgs.__new__(SchemaArgs)

        __props__.__dict__["definition"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["type"] = None
        return Schema(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output[str]:
        """
        The definition of the schema. This should contain a string representing the full definition of the schema that is a valid schema definition of the type specified in `type`.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Required. Name of the schema. Format is `projects/{project}/schemas/{schema}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the schema definition.
        """
        return pulumi.get(self, "type")

