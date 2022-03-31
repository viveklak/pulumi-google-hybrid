# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['PerfSampleSeriesArgs', 'PerfSampleSeries']

@pulumi.input_type
class PerfSampleSeriesArgs:
    def __init__(__self__, *,
                 execution_id: pulumi.Input[str],
                 history_id: pulumi.Input[str],
                 step_id: pulumi.Input[str],
                 basic_perf_sample_series: Optional[pulumi.Input['BasicPerfSampleSeriesArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PerfSampleSeries resource.
        :param pulumi.Input['BasicPerfSampleSeriesArgs'] basic_perf_sample_series: Basic series represented by a line chart
        """
        pulumi.set(__self__, "execution_id", execution_id)
        pulumi.set(__self__, "history_id", history_id)
        pulumi.set(__self__, "step_id", step_id)
        if basic_perf_sample_series is not None:
            pulumi.set(__self__, "basic_perf_sample_series", basic_perf_sample_series)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="executionId")
    def execution_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "execution_id")

    @execution_id.setter
    def execution_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "execution_id", value)

    @property
    @pulumi.getter(name="historyId")
    def history_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "history_id")

    @history_id.setter
    def history_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "history_id", value)

    @property
    @pulumi.getter(name="stepId")
    def step_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "step_id")

    @step_id.setter
    def step_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "step_id", value)

    @property
    @pulumi.getter(name="basicPerfSampleSeries")
    def basic_perf_sample_series(self) -> Optional[pulumi.Input['BasicPerfSampleSeriesArgs']]:
        """
        Basic series represented by a line chart
        """
        return pulumi.get(self, "basic_perf_sample_series")

    @basic_perf_sample_series.setter
    def basic_perf_sample_series(self, value: Optional[pulumi.Input['BasicPerfSampleSeriesArgs']]):
        pulumi.set(self, "basic_perf_sample_series", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class PerfSampleSeries(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 basic_perf_sample_series: Optional[pulumi.Input[pulumi.InputType['BasicPerfSampleSeriesArgs']]] = None,
                 execution_id: Optional[pulumi.Input[str]] = None,
                 history_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 step_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a PerfSampleSeries. May return any of the following error code(s): - ALREADY_EXISTS - PerfMetricSummary already exists for the given Step - NOT_FOUND - The containing Step does not exist
        Auto-naming is currently not supported for this resource.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BasicPerfSampleSeriesArgs']] basic_perf_sample_series: Basic series represented by a line chart
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PerfSampleSeriesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a PerfSampleSeries. May return any of the following error code(s): - ALREADY_EXISTS - PerfMetricSummary already exists for the given Step - NOT_FOUND - The containing Step does not exist
        Auto-naming is currently not supported for this resource.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param PerfSampleSeriesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PerfSampleSeriesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 basic_perf_sample_series: Optional[pulumi.Input[pulumi.InputType['BasicPerfSampleSeriesArgs']]] = None,
                 execution_id: Optional[pulumi.Input[str]] = None,
                 history_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 step_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = PerfSampleSeriesArgs.__new__(PerfSampleSeriesArgs)

            __props__.__dict__["basic_perf_sample_series"] = basic_perf_sample_series
            if execution_id is None and not opts.urn:
                raise TypeError("Missing required property 'execution_id'")
            __props__.__dict__["execution_id"] = execution_id
            if history_id is None and not opts.urn:
                raise TypeError("Missing required property 'history_id'")
            __props__.__dict__["history_id"] = history_id
            __props__.__dict__["project"] = project
            if step_id is None and not opts.urn:
                raise TypeError("Missing required property 'step_id'")
            __props__.__dict__["step_id"] = step_id
            __props__.__dict__["sample_series_id"] = None
        super(PerfSampleSeries, __self__).__init__(
            'google-native:toolresults/v1beta3:PerfSampleSeries',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PerfSampleSeries':
        """
        Get an existing PerfSampleSeries resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PerfSampleSeriesArgs.__new__(PerfSampleSeriesArgs)

        __props__.__dict__["basic_perf_sample_series"] = None
        __props__.__dict__["execution_id"] = None
        __props__.__dict__["history_id"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["sample_series_id"] = None
        __props__.__dict__["step_id"] = None
        return PerfSampleSeries(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="basicPerfSampleSeries")
    def basic_perf_sample_series(self) -> pulumi.Output['outputs.BasicPerfSampleSeriesResponse']:
        """
        Basic series represented by a line chart
        """
        return pulumi.get(self, "basic_perf_sample_series")

    @property
    @pulumi.getter(name="executionId")
    def execution_id(self) -> pulumi.Output[str]:
        """
        A tool results execution ID. 
        """
        return pulumi.get(self, "execution_id")

    @property
    @pulumi.getter(name="historyId")
    def history_id(self) -> pulumi.Output[str]:
        """
        A tool results history ID. 
        """
        return pulumi.get(self, "history_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The cloud project 
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="sampleSeriesId")
    def sample_series_id(self) -> pulumi.Output[str]:
        """
        A sample series id 
        """
        return pulumi.get(self, "sample_series_id")

    @property
    @pulumi.getter(name="stepId")
    def step_id(self) -> pulumi.Output[str]:
        """
        A tool results step ID. 
        """
        return pulumi.get(self, "step_id")

