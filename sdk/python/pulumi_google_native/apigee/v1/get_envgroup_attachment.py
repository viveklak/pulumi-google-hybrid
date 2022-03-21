# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetEnvgroupAttachmentResult',
    'AwaitableGetEnvgroupAttachmentResult',
    'get_envgroup_attachment',
    'get_envgroup_attachment_output',
]

@pulumi.output_type
class GetEnvgroupAttachmentResult:
    def __init__(__self__, created_at=None, environment=None, environment_group_id=None, name=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
        if environment_group_id and not isinstance(environment_group_id, str):
            raise TypeError("Expected argument 'environment_group_id' to be a str")
        pulumi.set(__self__, "environment_group_id", environment_group_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The time at which the environment group attachment was created as milliseconds since epoch.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def environment(self) -> str:
        """
        ID of the attached environment.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="environmentGroupId")
    def environment_group_id(self) -> str:
        """
        ID of the environment group.
        """
        return pulumi.get(self, "environment_group_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        ID of the environment group attachment.
        """
        return pulumi.get(self, "name")


class AwaitableGetEnvgroupAttachmentResult(GetEnvgroupAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvgroupAttachmentResult(
            created_at=self.created_at,
            environment=self.environment,
            environment_group_id=self.environment_group_id,
            name=self.name)


def get_envgroup_attachment(attachment_id: Optional[str] = None,
                            envgroup_id: Optional[str] = None,
                            organization_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvgroupAttachmentResult:
    """
    Gets an environment group attachment.
    """
    __args__ = dict()
    __args__['attachmentId'] = attachment_id
    __args__['envgroupId'] = envgroup_id
    __args__['organizationId'] = organization_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:apigee/v1:getEnvgroupAttachment', __args__, opts=opts, typ=GetEnvgroupAttachmentResult).value

    return AwaitableGetEnvgroupAttachmentResult(
        created_at=__ret__.created_at,
        environment=__ret__.environment,
        environment_group_id=__ret__.environment_group_id,
        name=__ret__.name)


@_utilities.lift_output_func(get_envgroup_attachment)
def get_envgroup_attachment_output(attachment_id: Optional[pulumi.Input[str]] = None,
                                   envgroup_id: Optional[pulumi.Input[str]] = None,
                                   organization_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnvgroupAttachmentResult]:
    """
    Gets an environment group attachment.
    """
    ...
