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

__all__ = ['AgentSessionEntityTypeArgs', 'AgentSessionEntityType']

@pulumi.input_type
class AgentSessionEntityTypeArgs:
    def __init__(__self__, *,
                 agent_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 project: pulumi.Input[str],
                 session_id: pulumi.Input[str],
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]]] = None,
                 entity_override_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AgentSessionEntityType resource.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]] entities: Required. The collection of entities to override or supplement the custom entity type.
        :param pulumi.Input[str] entity_override_mode: Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        :param pulumi.Input[str] name: Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
        """
        pulumi.set(__self__, "agent_id", agent_id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "session_id", session_id)
        if entities is not None:
            pulumi.set(__self__, "entities", entities)
        if entity_override_mode is not None:
            pulumi.set(__self__, "entity_override_mode", entity_override_mode)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="agentId")
    def agent_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "agent_id")

    @agent_id.setter
    def agent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "agent_id", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="sessionId")
    def session_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "session_id")

    @session_id.setter
    def session_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "session_id", value)

    @property
    @pulumi.getter
    def entities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]]]:
        """
        Required. The collection of entities to override or supplement the custom entity type.
        """
        return pulumi.get(self, "entities")

    @entities.setter
    def entities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]]]):
        pulumi.set(self, "entities", value)

    @property
    @pulumi.getter(name="entityOverrideMode")
    def entity_override_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        """
        return pulumi.get(self, "entity_override_mode")

    @entity_override_mode.setter
    def entity_override_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_override_mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class AgentSessionEntityType(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_id: Optional[pulumi.Input[str]] = None,
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]]]] = None,
                 entity_override_mode: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 session_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a session entity type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]]] entities: Required. The collection of entities to override or supplement the custom entity type.
        :param pulumi.Input[str] entity_override_mode: Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        :param pulumi.Input[str] name: Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AgentSessionEntityTypeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a session entity type.

        :param str resource_name: The name of the resource.
        :param AgentSessionEntityTypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AgentSessionEntityTypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_id: Optional[pulumi.Input[str]] = None,
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]]]] = None,
                 entity_override_mode: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 session_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = AgentSessionEntityTypeArgs.__new__(AgentSessionEntityTypeArgs)

            if agent_id is None and not opts.urn:
                raise TypeError("Missing required property 'agent_id'")
            __props__.__dict__["agent_id"] = agent_id
            __props__.__dict__["entities"] = entities
            __props__.__dict__["entity_override_mode"] = entity_override_mode
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if session_id is None and not opts.urn:
                raise TypeError("Missing required property 'session_id'")
            __props__.__dict__["session_id"] = session_id
        super(AgentSessionEntityType, __self__).__init__(
            'google-native:dialogflow/v3beta1:AgentSessionEntityType',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentSessionEntityType':
        """
        Get an existing AgentSessionEntityType resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AgentSessionEntityTypeArgs.__new__(AgentSessionEntityTypeArgs)

        __props__.__dict__["entities"] = None
        __props__.__dict__["entity_override_mode"] = None
        __props__.__dict__["name"] = None
        return AgentSessionEntityType(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def entities(self) -> pulumi.Output[Sequence['outputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponse']]:
        """
        Required. The collection of entities to override or supplement the custom entity type.
        """
        return pulumi.get(self, "entities")

    @property
    @pulumi.getter(name="entityOverrideMode")
    def entity_override_mode(self) -> pulumi.Output[str]:
        """
        Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        """
        return pulumi.get(self, "entity_override_mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
        """
        return pulumi.get(self, "name")

