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
    'GetTransitionRouteGroupResult',
    'AwaitableGetTransitionRouteGroupResult',
    'get_transition_route_group',
    'get_transition_route_group_output',
]

@pulumi.output_type
class GetTransitionRouteGroupResult:
    def __init__(__self__, display_name=None, name=None, transition_routes=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if transition_routes and not isinstance(transition_routes, list):
            raise TypeError("Expected argument 'transition_routes' to be a list")
        pulumi.set(__self__, "transition_routes", transition_routes)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="transitionRoutes")
    def transition_routes(self) -> Sequence['outputs.GoogleCloudDialogflowCxV3beta1TransitionRouteResponse']:
        """
        Transition routes associated with the TransitionRouteGroup.
        """
        return pulumi.get(self, "transition_routes")


class AwaitableGetTransitionRouteGroupResult(GetTransitionRouteGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransitionRouteGroupResult(
            display_name=self.display_name,
            name=self.name,
            transition_routes=self.transition_routes)


def get_transition_route_group(agent_id: Optional[str] = None,
                               flow_id: Optional[str] = None,
                               language_code: Optional[str] = None,
                               location: Optional[str] = None,
                               project: Optional[str] = None,
                               transition_route_group_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransitionRouteGroupResult:
    """
    Retrieves the specified TransitionRouteGroup.
    """
    __args__ = dict()
    __args__['agentId'] = agent_id
    __args__['flowId'] = flow_id
    __args__['languageCode'] = language_code
    __args__['location'] = location
    __args__['project'] = project
    __args__['transitionRouteGroupId'] = transition_route_group_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:dialogflow/v3beta1:getTransitionRouteGroup', __args__, opts=opts, typ=GetTransitionRouteGroupResult).value

    return AwaitableGetTransitionRouteGroupResult(
        display_name=__ret__.display_name,
        name=__ret__.name,
        transition_routes=__ret__.transition_routes)


@_utilities.lift_output_func(get_transition_route_group)
def get_transition_route_group_output(agent_id: Optional[pulumi.Input[str]] = None,
                                      flow_id: Optional[pulumi.Input[str]] = None,
                                      language_code: Optional[pulumi.Input[Optional[str]]] = None,
                                      location: Optional[pulumi.Input[str]] = None,
                                      project: Optional[pulumi.Input[Optional[str]]] = None,
                                      transition_route_group_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransitionRouteGroupResult]:
    """
    Retrieves the specified TransitionRouteGroup.
    """
    ...
