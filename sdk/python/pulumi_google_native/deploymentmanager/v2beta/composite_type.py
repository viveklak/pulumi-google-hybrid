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

__all__ = ['CompositeTypeArgs', 'CompositeType']

@pulumi.input_type
class CompositeTypeArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input['CompositeTypeLabelEntryArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['CompositeTypeStatus']] = None,
                 template_contents: Optional[pulumi.Input['TemplateContentsArgs']] = None):
        """
        The set of arguments for constructing a CompositeType resource.
        :param pulumi.Input[str] description: An optional textual description of the resource; provided by the client when the resource is created.
        :param pulumi.Input[Sequence[pulumi.Input['CompositeTypeLabelEntryArgs']]] labels: Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        :param pulumi.Input[str] name: Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
        :param pulumi.Input['TemplateContentsArgs'] template_contents: Files for the template type.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if template_contents is not None:
            pulumi.set(__self__, "template_contents", template_contents)

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
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CompositeTypeLabelEntryArgs']]]]:
        """
        Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CompositeTypeLabelEntryArgs']]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
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
    def status(self) -> Optional[pulumi.Input['CompositeTypeStatus']]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['CompositeTypeStatus']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="templateContents")
    def template_contents(self) -> Optional[pulumi.Input['TemplateContentsArgs']]:
        """
        Files for the template type.
        """
        return pulumi.get(self, "template_contents")

    @template_contents.setter
    def template_contents(self, value: Optional[pulumi.Input['TemplateContentsArgs']]):
        pulumi.set(self, "template_contents", value)


class CompositeType(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CompositeTypeLabelEntryArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['CompositeTypeStatus']] = None,
                 template_contents: Optional[pulumi.Input[pulumi.InputType['TemplateContentsArgs']]] = None,
                 __props__=None):
        """
        Creates a composite type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional textual description of the resource; provided by the client when the resource is created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CompositeTypeLabelEntryArgs']]]] labels: Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        :param pulumi.Input[str] name: Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
        :param pulumi.Input[pulumi.InputType['TemplateContentsArgs']] template_contents: Files for the template type.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CompositeTypeArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a composite type.

        :param str resource_name: The name of the resource.
        :param CompositeTypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CompositeTypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CompositeTypeLabelEntryArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['CompositeTypeStatus']] = None,
                 template_contents: Optional[pulumi.Input[pulumi.InputType['TemplateContentsArgs']]] = None,
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
            __props__ = CompositeTypeArgs.__new__(CompositeTypeArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["id"] = id
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["status"] = status
            __props__.__dict__["template_contents"] = template_contents
            __props__.__dict__["insert_time"] = None
            __props__.__dict__["operation"] = None
            __props__.__dict__["self_link"] = None
        super(CompositeType, __self__).__init__(
            'google-native:deploymentmanager/v2beta:CompositeType',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CompositeType':
        """
        Get an existing CompositeType resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CompositeTypeArgs.__new__(CompositeTypeArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["insert_time"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["operation"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["template_contents"] = None
        return CompositeType(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional textual description of the resource; provided by the client when the resource is created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="insertTime")
    def insert_time(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "insert_time")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Sequence['outputs.CompositeTypeLabelEntryResponse']]:
        """
        Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def operation(self) -> pulumi.Output['outputs.OperationResponse']:
        """
        The Operation that most recently ran, or is currently running, on this composite type.
        """
        return pulumi.get(self, "operation")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="templateContents")
    def template_contents(self) -> pulumi.Output['outputs.TemplateContentsResponse']:
        """
        Files for the template type.
        """
        return pulumi.get(self, "template_contents")

