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

__all__ = ['AgentEnvironmentSessionEntityTypeArgs', 'AgentEnvironmentSessionEntityType']

@pulumi.input_type
class AgentEnvironmentSessionEntityTypeArgs:
    def __init__(__self__, *,
                 agents_id: pulumi.Input[str],
                 entity_types_id: pulumi.Input[str],
                 environments_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 sessions_id: pulumi.Input[str],
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]]] = None,
                 entity_override_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AgentEnvironmentSessionEntityType resource.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]] entities: Required. The collection of entities to override or supplement the custom entity type.
        :param pulumi.Input[str] entity_override_mode: Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        :param pulumi.Input[str] name: Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
        """
        pulumi.set(__self__, "agents_id", agents_id)
        pulumi.set(__self__, "entity_types_id", entity_types_id)
        pulumi.set(__self__, "environments_id", environments_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        pulumi.set(__self__, "sessions_id", sessions_id)
        if entities is not None:
            pulumi.set(__self__, "entities", entities)
        if entity_override_mode is not None:
            pulumi.set(__self__, "entity_override_mode", entity_override_mode)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="agentsId")
    def agents_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "agents_id")

    @agents_id.setter
    def agents_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "agents_id", value)

    @property
    @pulumi.getter(name="entityTypesId")
    def entity_types_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "entity_types_id")

    @entity_types_id.setter
    def entity_types_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "entity_types_id", value)

    @property
    @pulumi.getter(name="environmentsId")
    def environments_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "environments_id")

    @environments_id.setter
    def environments_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environments_id", value)

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
    @pulumi.getter(name="sessionsId")
    def sessions_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "sessions_id")

    @sessions_id.setter
    def sessions_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "sessions_id", value)

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


class AgentEnvironmentSessionEntityType(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents_id: Optional[pulumi.Input[str]] = None,
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]]]] = None,
                 entity_override_mode: Optional[pulumi.Input[str]] = None,
                 entity_types_id: Optional[pulumi.Input[str]] = None,
                 environments_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 sessions_id: Optional[pulumi.Input[str]] = None,
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
                 args: AgentEnvironmentSessionEntityTypeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a session entity type.

        :param str resource_name: The name of the resource.
        :param AgentEnvironmentSessionEntityTypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AgentEnvironmentSessionEntityTypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents_id: Optional[pulumi.Input[str]] = None,
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs']]]]] = None,
                 entity_override_mode: Optional[pulumi.Input[str]] = None,
                 entity_types_id: Optional[pulumi.Input[str]] = None,
                 environments_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 sessions_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = AgentEnvironmentSessionEntityTypeArgs.__new__(AgentEnvironmentSessionEntityTypeArgs)

            if agents_id is None and not opts.urn:
                raise TypeError("Missing required property 'agents_id'")
            __props__.__dict__["agents_id"] = agents_id
            __props__.__dict__["entities"] = entities
            __props__.__dict__["entity_override_mode"] = entity_override_mode
            if entity_types_id is None and not opts.urn:
                raise TypeError("Missing required property 'entity_types_id'")
            __props__.__dict__["entity_types_id"] = entity_types_id
            if environments_id is None and not opts.urn:
                raise TypeError("Missing required property 'environments_id'")
            __props__.__dict__["environments_id"] = environments_id
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            __props__.__dict__["name"] = name
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            if sessions_id is None and not opts.urn:
                raise TypeError("Missing required property 'sessions_id'")
            __props__.__dict__["sessions_id"] = sessions_id
        super(AgentEnvironmentSessionEntityType, __self__).__init__(
            'gcp-native:dialogflow/v3beta1:AgentEnvironmentSessionEntityType',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentEnvironmentSessionEntityType':
        """
        Get an existing AgentEnvironmentSessionEntityType resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AgentEnvironmentSessionEntityTypeArgs.__new__(AgentEnvironmentSessionEntityTypeArgs)

        __props__.__dict__["entities"] = None
        __props__.__dict__["entity_override_mode"] = None
        __props__.__dict__["name"] = None
        return AgentEnvironmentSessionEntityType(resource_name, opts=opts, __props__=__props__)

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

