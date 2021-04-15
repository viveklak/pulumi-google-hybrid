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

__all__ = ['AgentFlowArgs', 'AgentFlow']

@pulumi.input_type
class AgentFlowArgs:
    def __init__(__self__, *,
                 agents_id: pulumi.Input[str],
                 flows_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 event_handlers: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1EventHandlerArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nlu_settings: Optional[pulumi.Input['GoogleCloudDialogflowCxV3beta1NluSettingsArgs']] = None,
                 transition_route_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 transition_routes: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1TransitionRouteArgs']]]] = None):
        """
        The set of arguments for constructing a AgentFlow resource.
        :param pulumi.Input[str] description: The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the flow.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1EventHandlerArgs']]] event_handlers: A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
        :param pulumi.Input[str] name: The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
        :param pulumi.Input['GoogleCloudDialogflowCxV3beta1NluSettingsArgs'] nlu_settings: NLU related settings of the flow.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] transition_route_groups: A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1TransitionRouteArgs']]] transition_routes: A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
        """
        pulumi.set(__self__, "agents_id", agents_id)
        pulumi.set(__self__, "flows_id", flows_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if event_handlers is not None:
            pulumi.set(__self__, "event_handlers", event_handlers)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nlu_settings is not None:
            pulumi.set(__self__, "nlu_settings", nlu_settings)
        if transition_route_groups is not None:
            pulumi.set(__self__, "transition_route_groups", transition_route_groups)
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
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The human-readable name of the flow.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="eventHandlers")
    def event_handlers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1EventHandlerArgs']]]]:
        """
        A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
        """
        return pulumi.get(self, "event_handlers")

    @event_handlers.setter
    def event_handlers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1EventHandlerArgs']]]]):
        pulumi.set(self, "event_handlers", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nluSettings")
    def nlu_settings(self) -> Optional[pulumi.Input['GoogleCloudDialogflowCxV3beta1NluSettingsArgs']]:
        """
        NLU related settings of the flow.
        """
        return pulumi.get(self, "nlu_settings")

    @nlu_settings.setter
    def nlu_settings(self, value: Optional[pulumi.Input['GoogleCloudDialogflowCxV3beta1NluSettingsArgs']]):
        pulumi.set(self, "nlu_settings", value)

    @property
    @pulumi.getter(name="transitionRouteGroups")
    def transition_route_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
        """
        return pulumi.get(self, "transition_route_groups")

    @transition_route_groups.setter
    def transition_route_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "transition_route_groups", value)

    @property
    @pulumi.getter(name="transitionRoutes")
    def transition_routes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1TransitionRouteArgs']]]]:
        """
        A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
        """
        return pulumi.get(self, "transition_routes")

    @transition_routes.setter
    def transition_routes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1TransitionRouteArgs']]]]):
        pulumi.set(self, "transition_routes", value)


class AgentFlow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 event_handlers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1EventHandlerArgs']]]]] = None,
                 flows_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nlu_settings: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1NluSettingsArgs']]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 transition_route_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 transition_routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1TransitionRouteArgs']]]]] = None,
                 __props__=None):
        """
        Creates a flow in the specified agent.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the flow.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1EventHandlerArgs']]]] event_handlers: A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
        :param pulumi.Input[str] name: The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1NluSettingsArgs']] nlu_settings: NLU related settings of the flow.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] transition_route_groups: A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1TransitionRouteArgs']]]] transition_routes: A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AgentFlowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a flow in the specified agent.

        :param str resource_name: The name of the resource.
        :param AgentFlowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AgentFlowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 event_handlers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1EventHandlerArgs']]]]] = None,
                 flows_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nlu_settings: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1NluSettingsArgs']]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 transition_route_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 transition_routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1TransitionRouteArgs']]]]] = None,
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
            __props__ = AgentFlowArgs.__new__(AgentFlowArgs)

            if agents_id is None and not opts.urn:
                raise TypeError("Missing required property 'agents_id'")
            __props__.__dict__["agents_id"] = agents_id
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["event_handlers"] = event_handlers
            if flows_id is None and not opts.urn:
                raise TypeError("Missing required property 'flows_id'")
            __props__.__dict__["flows_id"] = flows_id
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            __props__.__dict__["name"] = name
            __props__.__dict__["nlu_settings"] = nlu_settings
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            __props__.__dict__["transition_route_groups"] = transition_route_groups
            __props__.__dict__["transition_routes"] = transition_routes
        super(AgentFlow, __self__).__init__(
            'gcp-native:dialogflow/v3beta1:AgentFlow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentFlow':
        """
        Get an existing AgentFlow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AgentFlowArgs.__new__(AgentFlowArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["event_handlers"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["nlu_settings"] = None
        __props__.__dict__["transition_route_groups"] = None
        __props__.__dict__["transition_routes"] = None
        return AgentFlow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Required. The human-readable name of the flow.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="eventHandlers")
    def event_handlers(self) -> pulumi.Output[Sequence['outputs.GoogleCloudDialogflowCxV3beta1EventHandlerResponse']]:
        """
        A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
        """
        return pulumi.get(self, "event_handlers")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nluSettings")
    def nlu_settings(self) -> pulumi.Output['outputs.GoogleCloudDialogflowCxV3beta1NluSettingsResponse']:
        """
        NLU related settings of the flow.
        """
        return pulumi.get(self, "nlu_settings")

    @property
    @pulumi.getter(name="transitionRouteGroups")
    def transition_route_groups(self) -> pulumi.Output[Sequence[str]]:
        """
        A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
        """
        return pulumi.get(self, "transition_route_groups")

    @property
    @pulumi.getter(name="transitionRoutes")
    def transition_routes(self) -> pulumi.Output[Sequence['outputs.GoogleCloudDialogflowCxV3beta1TransitionRouteResponse']]:
        """
        A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
        """
        return pulumi.get(self, "transition_routes")

