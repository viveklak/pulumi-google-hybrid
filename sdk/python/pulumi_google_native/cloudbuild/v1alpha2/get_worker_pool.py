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
]

@pulumi.output_type
class GetWorkerPoolResult:
    def __init__(__self__, create_time=None, delete_time=None, name=None, network_config=None, region=None, state=None, update_time=None, worker_config=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if delete_time and not isinstance(delete_time, str):
            raise TypeError("Expected argument 'delete_time' to be a str")
        pulumi.set(__self__, "delete_time", delete_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_config and not isinstance(network_config, dict):
            raise TypeError("Expected argument 'network_config' to be a dict")
        pulumi.set(__self__, "network_config", network_config)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if worker_config and not isinstance(worker_config, dict):
            raise TypeError("Expected argument 'worker_config' to be a dict")
        pulumi.set(__self__, "worker_config", worker_config)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time at which the request to create the `WorkerPool` was received.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> str:
        """
        Time at which the request to delete the `WorkerPool` was received.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the `WorkerPool`. Format of the name is `projects/{project_id}/workerPools/{worker_pool_id}`, where the value of {worker_pool_id} is provided in the CreateWorkerPool request.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> 'outputs.NetworkConfigResponse':
        """
        Network configuration for the `WorkerPool`.
        """
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Immutable. The region where the `WorkerPool` runs. Only "us-central1" is currently supported. Note that `region` cannot be changed once the `WorkerPool` is created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        WorkerPool state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Time at which the request to update the `WorkerPool` was received.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> 'outputs.WorkerConfigResponse':
        """
        Worker configuration for the `WorkerPool`.
        """
        return pulumi.get(self, "worker_config")


class AwaitableGetWorkerPoolResult(GetWorkerPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkerPoolResult(
            create_time=self.create_time,
            delete_time=self.delete_time,
            name=self.name,
            network_config=self.network_config,
            region=self.region,
            state=self.state,
            update_time=self.update_time,
            worker_config=self.worker_config)


def get_worker_pool(project: Optional[str] = None,
                    worker_pool_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkerPoolResult:
    """
    Returns details of a `WorkerPool`.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['workerPoolId'] = worker_pool_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:cloudbuild/v1alpha2:getWorkerPool', __args__, opts=opts, typ=GetWorkerPoolResult).value

    return AwaitableGetWorkerPoolResult(
        create_time=__ret__.create_time,
        delete_time=__ret__.delete_time,
        name=__ret__.name,
        network_config=__ret__.network_config,
        region=__ret__.region,
        state=__ret__.state,
        update_time=__ret__.update_time,
        worker_config=__ret__.worker_config)
