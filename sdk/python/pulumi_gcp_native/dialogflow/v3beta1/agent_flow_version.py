# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = ['AgentFlowVersionArgs', 'AgentFlowVersion']

@pulumi.input_type
class AgentFlowVersionArgs:
    def __init__(__self__, *,
                 agents_id: pulumi.Input[str],
                 flows_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 versions_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AgentFlowVersion resource.
        :param pulumi.Input[str] description: The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the version. Limit of 64 characters.
        :param pulumi.Input[str] name: Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
        """
        pulumi.set(__self__, "agents_id", agents_id)
        pulumi.set(__self__, "flows_id", flows_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        pulumi.set(__self__, "versions_id", versions_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
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
    @pulumi.getter(name="versionsId")
    def versions_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "versions_id")

    @versions_id.setter
    def versions_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "versions_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The human-readable name of the version. Limit of 64 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class AgentFlowVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 flows_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 versions_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a Version in the specified Flow.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the version. Limit of 64 characters.
        :param pulumi.Input[str] name: Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AgentFlowVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Version in the specified Flow.

        :param str resource_name: The name of the resource.
        :param AgentFlowVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AgentFlowVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
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
                 flows_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 versions_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = AgentFlowVersionArgs.__new__(AgentFlowVersionArgs)

            if agents_id is None and not opts.urn:
                raise TypeError("Missing required property 'agents_id'")
            __props__.__dict__["agents_id"] = agents_id
            __props__.__dict__["description"] = description
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
            if versions_id is None and not opts.urn:
                raise TypeError("Missing required property 'versions_id'")
            __props__.__dict__["versions_id"] = versions_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["nlu_settings"] = None
            __props__.__dict__["state"] = None
        super(AgentFlowVersion, __self__).__init__(
            'gcp-native:dialogflow/v3beta1:AgentFlowVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentFlowVersion':
        """
        Get an existing AgentFlowVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AgentFlowVersionArgs.__new__(AgentFlowVersionArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["nlu_settings"] = None
        __props__.__dict__["state"] = None
        return AgentFlowVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the version.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Required. The human-readable name of the version. Limit of 64 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nluSettings")
    def nlu_settings(self) -> pulumi.Output['outputs.GoogleCloudDialogflowCxV3beta1NluSettingsResponse']:
        """
        The NLU settings of the flow at version creation.
        """
        return pulumi.get(self, "nlu_settings")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of this version. This field is read-only and cannot be set by create and update methods.
        """
        return pulumi.get(self, "state")

