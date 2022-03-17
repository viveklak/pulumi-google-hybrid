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
    'GetWorkerPoolResult',
    'AwaitableGetWorkerPoolResult',
    'get_worker_pool',
    'get_worker_pool_output',
]

@pulumi.output_type
class GetWorkerPoolResult:
    def __init__(__self__, autoscale=None, channel=None, name=None, state=None, worker_config=None, worker_count=None):
        if autoscale and not isinstance(autoscale, dict):
            raise TypeError("Expected argument 'autoscale' to be a dict")
        pulumi.set(__self__, "autoscale", autoscale)
        if channel and not isinstance(channel, str):
            raise TypeError("Expected argument 'channel' to be a str")
        pulumi.set(__self__, "channel", channel)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if worker_config and not isinstance(worker_config, dict):
            raise TypeError("Expected argument 'worker_config' to be a dict")
        pulumi.set(__self__, "worker_config", worker_config)
        if worker_count and not isinstance(worker_count, str):
            raise TypeError("Expected argument 'worker_count' to be a str")
        pulumi.set(__self__, "worker_count", worker_count)

    @property
    @pulumi.getter
    def autoscale(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponse':
        """
        The autoscale policy to apply on a pool.
        """
        return pulumi.get(self, "autoscale")

    @property
    @pulumi.getter
    def channel(self) -> str:
        """
        Channel specifies the release channel of the pool.
        """
        return pulumi.get(self, "channel")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the worker pool.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponse':
        """
        Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
        """
        return pulumi.get(self, "worker_config")

    @property
    @pulumi.getter(name="workerCount")
    def worker_count(self) -> str:
        """
        The desired number of workers in the worker pool. Must be a value between 0 and 15000.
        """
        return pulumi.get(self, "worker_count")


class AwaitableGetWorkerPoolResult(GetWorkerPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkerPoolResult(
            autoscale=self.autoscale,
            channel=self.channel,
            name=self.name,
            state=self.state,
            worker_config=self.worker_config,
            worker_count=self.worker_count)


def get_worker_pool(instance_id: Optional[str] = None,
                    project: Optional[str] = None,
                    workerpool_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkerPoolResult:
    """
    Returns the specified worker pool.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['project'] = project
    __args__['workerpoolId'] = workerpool_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:remotebuildexecution/v1alpha:getWorkerPool', __args__, opts=opts, typ=GetWorkerPoolResult).value

    return AwaitableGetWorkerPoolResult(
        autoscale=__ret__.autoscale,
        channel=__ret__.channel,
        name=__ret__.name,
        state=__ret__.state,
        worker_config=__ret__.worker_config,
        worker_count=__ret__.worker_count)


@_utilities.lift_output_func(get_worker_pool)
def get_worker_pool_output(instance_id: Optional[pulumi.Input[str]] = None,
                           project: Optional[pulumi.Input[Optional[str]]] = None,
                           workerpool_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkerPoolResult]:
    """
    Returns the specified worker pool.
    """
    ...
