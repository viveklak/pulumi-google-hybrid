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

__all__ = ['AgentFlowTransitionRouteGroupArgs', 'AgentFlowTransitionRouteGroup']

@pulumi.input_type
class AgentFlowTransitionRouteGroupArgs:
    def __init__(__self__, *,
                 agents_id: pulumi.Input[str],
                 flows_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 transition_route_groups_id: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 transition_routes: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3TransitionRouteArgs']]]] = None):
        """
        The set of arguments for constructing a AgentFlowTransitionRouteGroup resource.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
        :param pulumi.Input[str] name: The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3TransitionRouteArgs']]] transition_routes: Transition routes associated with the TransitionRouteGroup.
        """
        pulumi.set(__self__, "agents_id", agents_id)
        pulumi.set(__self__, "flows_id", flows_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        pulumi.set(__self__, "transition_route_groups_id", transition_route_groups_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if transition_routes is not None:
            pulumi.set(__self__, "transition_routes", transition_routes)

    @property
    @pulumi.getter(name="agentsId")
    def agents_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "agents_id")

    @agents_id.setter
    def agents_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "agents_id", value)

    @property
    @pulumi.getter(name="flowsId")
    def flows_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "flows_id")

    @flows_id.setter
    def flows_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "flows_id", value)

    @property
    @pulumi.getter(name="locationsId")
    def locations_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "locations_id")

    @locations_id.setter
    def locations_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "locations_id", value)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter(name="transitionRouteGroupsId")
    def transition_route_groups_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "transition_route_groups_id")

    @transition_route_groups_id.setter
    def transition_route_groups_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transition_route_groups_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="transitionRoutes")
    def transition_routes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3TransitionRouteArgs']]]]:
        """
        Transition routes associated with the TransitionRouteGroup.
        """
        return pulumi.get(self, "transition_routes")

    @transition_routes.setter
    def transition_routes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3TransitionRouteArgs']]]]):
        pulumi.set(self, "transition_routes", value)


class AgentFlowTransitionRouteGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 flows_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 transition_route_groups_id: Optional[pulumi.Input[str]] = None,
                 transition_routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3TransitionRouteArgs']]]]] = None,
                 __props__=None):
        """
        Creates an TransitionRouteGroup in the specified flow.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
        :param pulumi.Input[str] name: The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3TransitionRouteArgs']]]] transition_routes: Transition routes associated with the TransitionRouteGroup.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AgentFlowTransitionRouteGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an TransitionRouteGroup in the specified flow.

        :param str resource_name: The name of the resource.
        :param AgentFlowTransitionRouteGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AgentFlowTransitionRouteGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 flows_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 transition_route_groups_id: Optional[pulumi.Input[str]] = None,
                 transition_routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3TransitionRouteArgs']]]]] = None,
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
            __props__ = AgentFlowTransitionRouteGroupArgs.__new__(AgentFlowTransitionRouteGroupArgs)

            if agents_id is None and not opts.urn:
                raise TypeError("Missing required property 'agents_id'")
            __props__.__dict__["agents_id"] = agents_id
            __props__.__dict__["display_name"] = display_name
            if flows_id is None and not opts.urn:
                raise TypeError("Missing required property 'flows_id'")
            __props__.__dict__["flows_id"] = flows_id
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            __props__.__dict__["name"] = name
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            if transition_route_groups_id is None and not opts.urn:
                raise TypeError("Missing required property 'transition_route_groups_id'")
            __props__.__dict__["transition_route_groups_id"] = transition_route_groups_id
            __props__.__dict__["transition_routes"] = transition_routes
        super(AgentFlowTransitionRouteGroup, __self__).__init__(
            'gcp-native:dialogflow/v3:AgentFlowTransitionRouteGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentFlowTransitionRouteGroup':
        """
        Get an existing AgentFlowTransitionRouteGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AgentFlowTransitionRouteGroupArgs.__new__(AgentFlowTransitionRouteGroupArgs)

        __props__.__dict__["display_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["transition_routes"] = None
        return AgentFlowTransitionRouteGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Required. The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="transitionRoutes")
    def transition_routes(self) -> pulumi.Output[Sequence['outputs.GoogleCloudDialogflowCxV3TransitionRouteResponse']]:
        """
        Transition routes associated with the TransitionRouteGroup.
        """
        return pulumi.get(self, "transition_routes")

