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

__all__ = ['AgentEnvironmentExperimentArgs', 'AgentEnvironmentExperiment']

@pulumi.input_type
class AgentEnvironmentExperimentArgs:
    def __init__(__self__, *,
                 agent_id: pulumi.Input[str],
                 environment_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 project: pulumi.Input[str],
                 create_time: Optional[pulumi.Input[str]] = None,
                 definition: Optional[pulumi.Input['GoogleCloudDialogflowCxV3beta1ExperimentDefinitionArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 experiment_length: Optional[pulumi.Input[str]] = None,
                 last_update_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 result: Optional[pulumi.Input['GoogleCloudDialogflowCxV3beta1ExperimentResultArgs']] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 variants_history: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs']]]] = None):
        """
        The set of arguments for constructing a AgentEnvironmentExperiment resource.
        :param pulumi.Input[str] create_time: Creation time of this experiment.
        :param pulumi.Input['GoogleCloudDialogflowCxV3beta1ExperimentDefinitionArgs'] definition: The definition of the experiment.
        :param pulumi.Input[str] description: The human-readable description of the experiment.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
        :param pulumi.Input[str] end_time: End time of this experiment.
        :param pulumi.Input[str] experiment_length: Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
        :param pulumi.Input[str] last_update_time: Last update time of this experiment.
        :param pulumi.Input[str] name: The name of the experiment. Format: projects//locations//agents//environments//experiments/..
        :param pulumi.Input['GoogleCloudDialogflowCxV3beta1ExperimentResultArgs'] result: Inference result of the experiment.
        :param pulumi.Input[str] start_time: Start time of this experiment.
        :param pulumi.Input[str] state: The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs']]] variants_history: The history of updates to the experiment variants.
        """
        pulumi.set(__self__, "agent_id", agent_id)
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "project", project)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if definition is not None:
            pulumi.set(__self__, "definition", definition)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if experiment_length is not None:
            pulumi.set(__self__, "experiment_length", experiment_length)
        if last_update_time is not None:
            pulumi.set(__self__, "last_update_time", last_update_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if result is not None:
            pulumi.set(__self__, "result", result)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if variants_history is not None:
            pulumi.set(__self__, "variants_history", variants_history)

    @property
    @pulumi.getter(name="agentId")
    def agent_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "agent_id")

    @agent_id.setter
    def agent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "agent_id", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_id", value)

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
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of this experiment.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def definition(self) -> Optional[pulumi.Input['GoogleCloudDialogflowCxV3beta1ExperimentDefinitionArgs']]:
        """
        The definition of the experiment.
        """
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: Optional[pulumi.Input['GoogleCloudDialogflowCxV3beta1ExperimentDefinitionArgs']]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable description of the experiment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        End time of this experiment.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter(name="experimentLength")
    def experiment_length(self) -> Optional[pulumi.Input[str]]:
        """
        Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
        """
        return pulumi.get(self, "experiment_length")

    @experiment_length.setter
    def experiment_length(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "experiment_length", value)

    @property
    @pulumi.getter(name="lastUpdateTime")
    def last_update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last update time of this experiment.
        """
        return pulumi.get(self, "last_update_time")

    @last_update_time.setter
    def last_update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_update_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the experiment. Format: projects//locations//agents//environments//experiments/..
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def result(self) -> Optional[pulumi.Input['GoogleCloudDialogflowCxV3beta1ExperimentResultArgs']]:
        """
        Inference result of the experiment.
        """
        return pulumi.get(self, "result")

    @result.setter
    def result(self, value: Optional[pulumi.Input['GoogleCloudDialogflowCxV3beta1ExperimentResultArgs']]):
        pulumi.set(self, "result", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[str]]:
        """
        Start time of this experiment.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="variantsHistory")
    def variants_history(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs']]]]:
        """
        The history of updates to the experiment variants.
        """
        return pulumi.get(self, "variants_history")

    @variants_history.setter
    def variants_history(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs']]]]):
        pulumi.set(self, "variants_history", value)


class AgentEnvironmentExperiment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_id: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1ExperimentDefinitionArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 experiment_length: Optional[pulumi.Input[str]] = None,
                 last_update_time: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 result: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1ExperimentResultArgs']]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 variants_history: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs']]]]] = None,
                 __props__=None):
        """
        Creates an Experiment in the specified Environment.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Creation time of this experiment.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1ExperimentDefinitionArgs']] definition: The definition of the experiment.
        :param pulumi.Input[str] description: The human-readable description of the experiment.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
        :param pulumi.Input[str] end_time: End time of this experiment.
        :param pulumi.Input[str] experiment_length: Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
        :param pulumi.Input[str] last_update_time: Last update time of this experiment.
        :param pulumi.Input[str] name: The name of the experiment. Format: projects//locations//agents//environments//experiments/..
        :param pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1ExperimentResultArgs']] result: Inference result of the experiment.
        :param pulumi.Input[str] start_time: Start time of this experiment.
        :param pulumi.Input[str] state: The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs']]]] variants_history: The history of updates to the experiment variants.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AgentEnvironmentExperimentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an Experiment in the specified Environment.

        :param str resource_name: The name of the resource.
        :param AgentEnvironmentExperimentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AgentEnvironmentExperimentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_id: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1ExperimentDefinitionArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 experiment_length: Optional[pulumi.Input[str]] = None,
                 last_update_time: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 result: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1ExperimentResultArgs']]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 variants_history: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs']]]]] = None,
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
            __props__ = AgentEnvironmentExperimentArgs.__new__(AgentEnvironmentExperimentArgs)

            if agent_id is None and not opts.urn:
                raise TypeError("Missing required property 'agent_id'")
            __props__.__dict__["agent_id"] = agent_id
            __props__.__dict__["create_time"] = create_time
            __props__.__dict__["definition"] = definition
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["end_time"] = end_time
            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            __props__.__dict__["experiment_length"] = experiment_length
            __props__.__dict__["last_update_time"] = last_update_time
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["result"] = result
            __props__.__dict__["start_time"] = start_time
            __props__.__dict__["state"] = state
            __props__.__dict__["variants_history"] = variants_history
        super(AgentEnvironmentExperiment, __self__).__init__(
            'google-native:dialogflow/v3beta1:AgentEnvironmentExperiment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentEnvironmentExperiment':
        """
        Get an existing AgentEnvironmentExperiment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AgentEnvironmentExperimentArgs.__new__(AgentEnvironmentExperimentArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["definition"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["end_time"] = None
        __props__.__dict__["experiment_length"] = None
        __props__.__dict__["last_update_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["result"] = None
        __props__.__dict__["start_time"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["variants_history"] = None
        return AgentEnvironmentExperiment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time of this experiment.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output['outputs.GoogleCloudDialogflowCxV3beta1ExperimentDefinitionResponse']:
        """
        The definition of the experiment.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The human-readable description of the experiment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[str]:
        """
        End time of this experiment.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="experimentLength")
    def experiment_length(self) -> pulumi.Output[str]:
        """
        Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
        """
        return pulumi.get(self, "experiment_length")

    @property
    @pulumi.getter(name="lastUpdateTime")
    def last_update_time(self) -> pulumi.Output[str]:
        """
        Last update time of this experiment.
        """
        return pulumi.get(self, "last_update_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the experiment. Format: projects//locations//agents//environments//experiments/..
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output['outputs.GoogleCloudDialogflowCxV3beta1ExperimentResultResponse']:
        """
        Inference result of the experiment.
        """
        return pulumi.get(self, "result")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[str]:
        """
        Start time of this experiment.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="variantsHistory")
    def variants_history(self) -> pulumi.Output[Sequence['outputs.GoogleCloudDialogflowCxV3beta1VariantsHistoryResponse']]:
        """
        The history of updates to the experiment variants.
        """
        return pulumi.get(self, "variants_history")

