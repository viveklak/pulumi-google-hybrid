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

__all__ = ['StudyTrialArgs', 'StudyTrial']

@pulumi.input_type
class StudyTrialArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 project: pulumi.Input[str],
                 study_id: pulumi.Input[str],
                 final_measurement: Optional[pulumi.Input['GoogleCloudMlV1__MeasurementArgs']] = None,
                 measurements: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudMlV1__MeasurementArgs']]]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudMlV1_Trial_ParameterArgs']]]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StudyTrial resource.
        :param pulumi.Input['GoogleCloudMlV1__MeasurementArgs'] final_measurement: The final measurement containing the objective value.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudMlV1__MeasurementArgs']]] measurements: A list of measurements that are strictly lexicographically ordered by their induced tuples (steps, elapsed_time). These are used for early stopping computations.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudMlV1_Trial_ParameterArgs']]] parameters: The parameters of the trial.
        :param pulumi.Input[str] state: The detailed state of a trial.
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "study_id", study_id)
        if final_measurement is not None:
            pulumi.set(__self__, "final_measurement", final_measurement)
        if measurements is not None:
            pulumi.set(__self__, "measurements", measurements)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if state is not None:
            pulumi.set(__self__, "state", state)

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
    @pulumi.getter(name="studyId")
    def study_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "study_id")

    @study_id.setter
    def study_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "study_id", value)

    @property
    @pulumi.getter(name="finalMeasurement")
    def final_measurement(self) -> Optional[pulumi.Input['GoogleCloudMlV1__MeasurementArgs']]:
        """
        The final measurement containing the objective value.
        """
        return pulumi.get(self, "final_measurement")

    @final_measurement.setter
    def final_measurement(self, value: Optional[pulumi.Input['GoogleCloudMlV1__MeasurementArgs']]):
        pulumi.set(self, "final_measurement", value)

    @property
    @pulumi.getter
    def measurements(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudMlV1__MeasurementArgs']]]]:
        """
        A list of measurements that are strictly lexicographically ordered by their induced tuples (steps, elapsed_time). These are used for early stopping computations.
        """
        return pulumi.get(self, "measurements")

    @measurements.setter
    def measurements(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudMlV1__MeasurementArgs']]]]):
        pulumi.set(self, "measurements", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudMlV1_Trial_ParameterArgs']]]]:
        """
        The parameters of the trial.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudMlV1_Trial_ParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The detailed state of a trial.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


class StudyTrial(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 final_measurement: Optional[pulumi.Input[pulumi.InputType['GoogleCloudMlV1__MeasurementArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 measurements: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudMlV1__MeasurementArgs']]]]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudMlV1_Trial_ParameterArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 study_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Adds a user provided trial to a study.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GoogleCloudMlV1__MeasurementArgs']] final_measurement: The final measurement containing the objective value.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudMlV1__MeasurementArgs']]]] measurements: A list of measurements that are strictly lexicographically ordered by their induced tuples (steps, elapsed_time). These are used for early stopping computations.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudMlV1_Trial_ParameterArgs']]]] parameters: The parameters of the trial.
        :param pulumi.Input[str] state: The detailed state of a trial.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StudyTrialArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Adds a user provided trial to a study.

        :param str resource_name: The name of the resource.
        :param StudyTrialArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StudyTrialArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 final_measurement: Optional[pulumi.Input[pulumi.InputType['GoogleCloudMlV1__MeasurementArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 measurements: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudMlV1__MeasurementArgs']]]]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudMlV1_Trial_ParameterArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 study_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = StudyTrialArgs.__new__(StudyTrialArgs)

            __props__.__dict__["final_measurement"] = final_measurement
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["measurements"] = measurements
            __props__.__dict__["parameters"] = parameters
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["state"] = state
            if study_id is None and not opts.urn:
                raise TypeError("Missing required property 'study_id'")
            __props__.__dict__["study_id"] = study_id
            __props__.__dict__["client_id"] = None
            __props__.__dict__["end_time"] = None
            __props__.__dict__["infeasible_reason"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["start_time"] = None
            __props__.__dict__["trial_infeasible"] = None
        super(StudyTrial, __self__).__init__(
            'google-native:ml/v1:StudyTrial',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'StudyTrial':
        """
        Get an existing StudyTrial resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = StudyTrialArgs.__new__(StudyTrialArgs)

        __props__.__dict__["client_id"] = None
        __props__.__dict__["end_time"] = None
        __props__.__dict__["final_measurement"] = None
        __props__.__dict__["infeasible_reason"] = None
        __props__.__dict__["measurements"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["start_time"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["trial_infeasible"] = None
        return StudyTrial(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        The identifier of the client that originally requested this trial.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[str]:
        """
        Time at which the trial's status changed to COMPLETED.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="finalMeasurement")
    def final_measurement(self) -> pulumi.Output['outputs.GoogleCloudMlV1__MeasurementResponse']:
        """
        The final measurement containing the objective value.
        """
        return pulumi.get(self, "final_measurement")

    @property
    @pulumi.getter(name="infeasibleReason")
    def infeasible_reason(self) -> pulumi.Output[str]:
        """
        A human readable string describing why the trial is infeasible. This should only be set if trial_infeasible is true.
        """
        return pulumi.get(self, "infeasible_reason")

    @property
    @pulumi.getter
    def measurements(self) -> pulumi.Output[Sequence['outputs.GoogleCloudMlV1__MeasurementResponse']]:
        """
        A list of measurements that are strictly lexicographically ordered by their induced tuples (steps, elapsed_time). These are used for early stopping computations.
        """
        return pulumi.get(self, "measurements")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the trial assigned by the service.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Sequence['outputs.GoogleCloudMlV1_Trial_ParameterResponse']]:
        """
        The parameters of the trial.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[str]:
        """
        Time at which the trial was started.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The detailed state of a trial.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="trialInfeasible")
    def trial_infeasible(self) -> pulumi.Output[bool]:
        """
        If true, the parameters in this trial are not attempted again.
        """
        return pulumi.get(self, "trial_infeasible")

