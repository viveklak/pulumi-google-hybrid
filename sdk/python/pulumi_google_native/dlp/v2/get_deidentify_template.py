# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetDeidentifyTemplateResult',
    'AwaitableGetDeidentifyTemplateResult',
    'get_deidentify_template',
    'get_deidentify_template_output',
]

@pulumi.output_type
class GetDeidentifyTemplateResult:
    def __init__(__self__, create_time=None, deidentify_config=None, description=None, display_name=None, name=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if deidentify_config and not isinstance(deidentify_config, dict):
            raise TypeError("Expected argument 'deidentify_config' to be a dict")
        pulumi.set(__self__, "deidentify_config", deidentify_config)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation timestamp of an inspectTemplate.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deidentifyConfig")
    def deidentify_config(self) -> 'outputs.GooglePrivacyDlpV2DeidentifyConfigResponse':
        """
        The core content of the template.
        """
        return pulumi.get(self, "deidentify_config")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Short description (max 256 chars).
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name (max 256 chars).
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The template name. The template will have one of the following formats: `projects/PROJECT_ID/deidentifyTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/deidentifyTemplates/TEMPLATE_ID`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update timestamp of an inspectTemplate.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetDeidentifyTemplateResult(GetDeidentifyTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeidentifyTemplateResult(
            create_time=self.create_time,
            deidentify_config=self.deidentify_config,
            description=self.description,
            display_name=self.display_name,
            name=self.name,
            update_time=self.update_time)


def get_deidentify_template(deidentify_template_id: Optional[str] = None,
                            location: Optional[str] = None,
                            project: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeidentifyTemplateResult:
    """
    Gets a DeidentifyTemplate. See https://cloud.google.com/dlp/docs/creating-templates-deid to learn more.
    """
    __args__ = dict()
    __args__['deidentifyTemplateId'] = deidentify_template_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:dlp/v2:getDeidentifyTemplate', __args__, opts=opts, typ=GetDeidentifyTemplateResult).value

    return AwaitableGetDeidentifyTemplateResult(
        create_time=__ret__.create_time,
        deidentify_config=__ret__.deidentify_config,
        description=__ret__.description,
        display_name=__ret__.display_name,
        name=__ret__.name,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_deidentify_template)
def get_deidentify_template_output(deidentify_template_id: Optional[pulumi.Input[str]] = None,
                                   location: Optional[pulumi.Input[str]] = None,
                                   project: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeidentifyTemplateResult]:
    """
    Gets a DeidentifyTemplate. See https://cloud.google.com/dlp/docs/creating-templates-deid to learn more.
    """
    ...
