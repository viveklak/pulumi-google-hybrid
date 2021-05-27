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

__all__ = ['OrganizationInspectTemplateArgs', 'OrganizationInspectTemplate']

@pulumi.input_type
class OrganizationInspectTemplateArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 organization_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 inspect_config: Optional[pulumi.Input['GooglePrivacyDlpV2InspectConfigArgs']] = None,
                 template_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrganizationInspectTemplate resource.
        :param pulumi.Input[str] description: Short description (max 256 chars).
        :param pulumi.Input[str] display_name: Display name (max 256 chars).
        :param pulumi.Input['GooglePrivacyDlpV2InspectConfigArgs'] inspect_config: The core content of the template. Configuration of the scanning process.
        :param pulumi.Input[str] template_id: The template id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "organization_id", organization_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if inspect_config is not None:
            pulumi.set(__self__, "inspect_config", inspect_config)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Short description (max 256 chars).
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name (max 256 chars).
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="inspectConfig")
    def inspect_config(self) -> Optional[pulumi.Input['GooglePrivacyDlpV2InspectConfigArgs']]:
        """
        The core content of the template. Configuration of the scanning process.
        """
        return pulumi.get(self, "inspect_config")

    @inspect_config.setter
    def inspect_config(self, value: Optional[pulumi.Input['GooglePrivacyDlpV2InspectConfigArgs']]):
        pulumi.set(self, "inspect_config", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        The template id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)


class OrganizationInspectTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 inspect_config: Optional[pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2InspectConfigArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an InspectTemplate for re-using frequently used configuration for inspecting content, images, and storage. See https://cloud.google.com/dlp/docs/creating-templates to learn more.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Short description (max 256 chars).
        :param pulumi.Input[str] display_name: Display name (max 256 chars).
        :param pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2InspectConfigArgs']] inspect_config: The core content of the template. Configuration of the scanning process.
        :param pulumi.Input[str] template_id: The template id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationInspectTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an InspectTemplate for re-using frequently used configuration for inspecting content, images, and storage. See https://cloud.google.com/dlp/docs/creating-templates to learn more.

        :param str resource_name: The name of the resource.
        :param OrganizationInspectTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationInspectTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 inspect_config: Optional[pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2InspectConfigArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = OrganizationInspectTemplateArgs.__new__(OrganizationInspectTemplateArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["inspect_config"] = inspect_config
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["template_id"] = template_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        super(OrganizationInspectTemplate, __self__).__init__(
            'google-native:dlp/v2:OrganizationInspectTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OrganizationInspectTemplate':
        """
        Get an existing OrganizationInspectTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OrganizationInspectTemplateArgs.__new__(OrganizationInspectTemplateArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["inspect_config"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["update_time"] = None
        return OrganizationInspectTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation timestamp of an inspectTemplate.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Short description (max 256 chars).
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Display name (max 256 chars).
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="inspectConfig")
    def inspect_config(self) -> pulumi.Output['outputs.GooglePrivacyDlpV2InspectConfigResponse']:
        """
        The core content of the template. Configuration of the scanning process.
        """
        return pulumi.get(self, "inspect_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The template name. The template will have one of the following formats: `projects/PROJECT_ID/inspectTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/inspectTemplates/TEMPLATE_ID`;
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The last update timestamp of an inspectTemplate.
        """
        return pulumi.get(self, "update_time")

