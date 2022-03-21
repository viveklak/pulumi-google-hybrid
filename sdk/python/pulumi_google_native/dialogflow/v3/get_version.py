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
    'GetVersionResult',
    'AwaitableGetVersionResult',
    'get_version',
    'get_version_output',
]

@pulumi.output_type
class GetVersionResult:
    def __init__(__self__, create_time=None, description=None, display_name=None, name=None, nlu_settings=None, state=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nlu_settings and not isinstance(nlu_settings, dict):
            raise TypeError("Expected argument 'nlu_settings' to be a dict")
        pulumi.set(__self__, "nlu_settings", nlu_settings)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Create time of the version.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The human-readable name of the version. Limit of 64 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nluSettings")
    def nlu_settings(self) -> 'outputs.GoogleCloudDialogflowCxV3NluSettingsResponse':
        """
        The NLU settings of the flow at version creation.
        """
        return pulumi.get(self, "nlu_settings")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of this version. This field is read-only and cannot be set by create and update methods.
        """
        return pulumi.get(self, "state")


class AwaitableGetVersionResult(GetVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVersionResult(
            create_time=self.create_time,
            description=self.description,
            display_name=self.display_name,
            name=self.name,
            nlu_settings=self.nlu_settings,
            state=self.state)


def get_version(agent_id: Optional[str] = None,
                flow_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                version_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVersionResult:
    """
    Retrieves the specified Version.
    """
    __args__ = dict()
    __args__['agentId'] = agent_id
    __args__['flowId'] = flow_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['versionId'] = version_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:dialogflow/v3:getVersion', __args__, opts=opts, typ=GetVersionResult).value

    return AwaitableGetVersionResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        display_name=__ret__.display_name,
        name=__ret__.name,
        nlu_settings=__ret__.nlu_settings,
        state=__ret__.state)


@_utilities.lift_output_func(get_version)
def get_version_output(agent_id: Optional[pulumi.Input[str]] = None,
                       flow_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       version_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVersionResult]:
    """
    Retrieves the specified Version.
    """
    ...
