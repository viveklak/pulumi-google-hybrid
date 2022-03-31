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
    'GetEnvironmentResult',
    'AwaitableGetEnvironmentResult',
    'get_environment',
    'get_environment_output',
]

@pulumi.output_type
class GetEnvironmentResult:
    def __init__(__self__, agent_version=None, description=None, fulfillment=None, name=None, state=None, text_to_speech_settings=None, update_time=None):
        if agent_version and not isinstance(agent_version, str):
            raise TypeError("Expected argument 'agent_version' to be a str")
        pulumi.set(__self__, "agent_version", agent_version)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if fulfillment and not isinstance(fulfillment, dict):
            raise TypeError("Expected argument 'fulfillment' to be a dict")
        pulumi.set(__self__, "fulfillment", fulfillment)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if text_to_speech_settings and not isinstance(text_to_speech_settings, dict):
            raise TypeError("Expected argument 'text_to_speech_settings' to be a dict")
        pulumi.set(__self__, "text_to_speech_settings", text_to_speech_settings)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="agentVersion")
    def agent_version(self) -> str:
        """
        Optional. The agent version loaded into this environment. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
        """
        return pulumi.get(self, "agent_version")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. The developer-provided description for this environment. The maximum length is 500 characters. If exceeded, the request is rejected.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fulfillment(self) -> 'outputs.GoogleCloudDialogflowV2FulfillmentResponse':
        """
        Optional. The fulfillment settings to use for this environment.
        """
        return pulumi.get(self, "fulfillment")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The unique identifier of this agent environment. Supported formats: - `projects//agent/environments/` - `projects//locations//agent/environments/` The environment ID for the default environment is `-`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="textToSpeechSettings")
    def text_to_speech_settings(self) -> 'outputs.GoogleCloudDialogflowV2TextToSpeechSettingsResponse':
        """
        Optional. Text to speech settings for this environment.
        """
        return pulumi.get(self, "text_to_speech_settings")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update time of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetEnvironmentResult(GetEnvironmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentResult(
            agent_version=self.agent_version,
            description=self.description,
            fulfillment=self.fulfillment,
            name=self.name,
            state=self.state,
            text_to_speech_settings=self.text_to_speech_settings,
            update_time=self.update_time)


def get_environment(environment_id: Optional[str] = None,
                    location: Optional[str] = None,
                    project: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvironmentResult:
    """
    Retrieves the specified agent environment.
    """
    __args__ = dict()
    __args__['environmentId'] = environment_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:dialogflow/v2:getEnvironment', __args__, opts=opts, typ=GetEnvironmentResult).value

    return AwaitableGetEnvironmentResult(
        agent_version=__ret__.agent_version,
        description=__ret__.description,
        fulfillment=__ret__.fulfillment,
        name=__ret__.name,
        state=__ret__.state,
        text_to_speech_settings=__ret__.text_to_speech_settings,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_environment)
def get_environment_output(environment_id: Optional[pulumi.Input[str]] = None,
                           location: Optional[pulumi.Input[str]] = None,
                           project: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnvironmentResult]:
    """
    Retrieves the specified agent environment.
    """
    ...
