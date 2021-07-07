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
    'GetPerfSampleSeriesResult',
    'AwaitableGetPerfSampleSeriesResult',
    'get_perf_sample_series',
]

@pulumi.output_type
class GetPerfSampleSeriesResult:
    def __init__(__self__, basic_perf_sample_series=None, execution_id=None, history_id=None, project=None, sample_series_id=None, step_id=None):
        if basic_perf_sample_series and not isinstance(basic_perf_sample_series, dict):
            raise TypeError("Expected argument 'basic_perf_sample_series' to be a dict")
        pulumi.set(__self__, "basic_perf_sample_series", basic_perf_sample_series)
        if execution_id and not isinstance(execution_id, str):
            raise TypeError("Expected argument 'execution_id' to be a str")
        pulumi.set(__self__, "execution_id", execution_id)
        if history_id and not isinstance(history_id, str):
            raise TypeError("Expected argument 'history_id' to be a str")
        pulumi.set(__self__, "history_id", history_id)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if sample_series_id and not isinstance(sample_series_id, str):
            raise TypeError("Expected argument 'sample_series_id' to be a str")
        pulumi.set(__self__, "sample_series_id", sample_series_id)
        if step_id and not isinstance(step_id, str):
            raise TypeError("Expected argument 'step_id' to be a str")
        pulumi.set(__self__, "step_id", step_id)

    @property
    @pulumi.getter(name="basicPerfSampleSeries")
    def basic_perf_sample_series(self) -> 'outputs.BasicPerfSampleSeriesResponse':
        """
        Basic series represented by a line chart
        """
        return pulumi.get(self, "basic_perf_sample_series")

    @property
    @pulumi.getter(name="executionId")
    def execution_id(self) -> str:
        """
        A tool results execution ID. 
        """
        return pulumi.get(self, "execution_id")

    @property
    @pulumi.getter(name="historyId")
    def history_id(self) -> str:
        """
        A tool results history ID. 
        """
        return pulumi.get(self, "history_id")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The cloud project 
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="sampleSeriesId")
    def sample_series_id(self) -> str:
        """
        A sample series id 
        """
        return pulumi.get(self, "sample_series_id")

    @property
    @pulumi.getter(name="stepId")
    def step_id(self) -> str:
        """
        A tool results step ID. 
        """
        return pulumi.get(self, "step_id")


class AwaitableGetPerfSampleSeriesResult(GetPerfSampleSeriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPerfSampleSeriesResult(
            basic_perf_sample_series=self.basic_perf_sample_series,
            execution_id=self.execution_id,
            history_id=self.history_id,
            project=self.project,
            sample_series_id=self.sample_series_id,
            step_id=self.step_id)


def get_perf_sample_series(execution_id: Optional[str] = None,
                           history_id: Optional[str] = None,
                           project: Optional[str] = None,
                           sample_series_id: Optional[str] = None,
                           step_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPerfSampleSeriesResult:
    """
    Gets a PerfSampleSeries. May return any of the following error code(s): - NOT_FOUND - The specified PerfSampleSeries does not exist
    """
    __args__ = dict()
    __args__['executionId'] = execution_id
    __args__['historyId'] = history_id
    __args__['project'] = project
    __args__['sampleSeriesId'] = sample_series_id
    __args__['stepId'] = step_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:toolresults/v1beta3:getPerfSampleSeries', __args__, opts=opts, typ=GetPerfSampleSeriesResult).value

    return AwaitableGetPerfSampleSeriesResult(
        basic_perf_sample_series=__ret__.basic_perf_sample_series,
        execution_id=__ret__.execution_id,
        history_id=__ret__.history_id,
        project=__ret__.project,
        sample_series_id=__ret__.sample_series_id,
        step_id=__ret__.step_id)
