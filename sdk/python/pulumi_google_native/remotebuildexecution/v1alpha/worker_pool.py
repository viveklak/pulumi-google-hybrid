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

__all__ = ['WorkerPoolArgs', 'WorkerPool']

@pulumi.input_type
class WorkerPoolArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 autoscale: Optional[pulumi.Input['GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleArgs']] = None,
                 channel: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 worker_config: Optional[pulumi.Input['GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs']] = None,
                 worker_count: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WorkerPool resource.
        :param pulumi.Input['GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleArgs'] autoscale: The autoscale policy to apply on a pool.
        :param pulumi.Input[str] channel: Channel specifies the release channel of the pool.
        :param pulumi.Input[str] name: WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
        :param pulumi.Input[str] parent: Resource name of the instance in which to create the new worker pool. Format: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]`.
        :param pulumi.Input[str] pool_id: ID of the created worker pool. A valid pool ID must: be 6-50 characters long, contain only lowercase letters, digits, hyphens and underscores, start with a lowercase letter, and end with a lowercase letter or a digit.
        :param pulumi.Input['GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs'] worker_config: Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
        :param pulumi.Input[str] worker_count: The desired number of workers in the worker pool. Must be a value between 0 and 15000.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if autoscale is not None:
            pulumi.set(__self__, "autoscale", autoscale)
        if channel is not None:
            pulumi.set(__self__, "channel", channel)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent is not None:
            pulumi.set(__self__, "parent", parent)
        if pool_id is not None:
            pulumi.set(__self__, "pool_id", pool_id)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if worker_config is not None:
            pulumi.set(__self__, "worker_config", worker_config)
        if worker_count is not None:
            pulumi.set(__self__, "worker_count", worker_count)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def autoscale(self) -> Optional[pulumi.Input['GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleArgs']]:
        """
        The autoscale policy to apply on a pool.
        """
        return pulumi.get(self, "autoscale")

    @autoscale.setter
    def autoscale(self, value: Optional[pulumi.Input['GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleArgs']]):
        pulumi.set(self, "autoscale", value)

    @property
    @pulumi.getter
    def channel(self) -> Optional[pulumi.Input[str]]:
        """
        Channel specifies the release channel of the pool.
        """
        return pulumi.get(self, "channel")

    @channel.setter
    def channel(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "channel", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name of the instance in which to create the new worker pool. Format: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]`.
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the created worker pool. A valid pool ID must: be 6-50 characters long, contain only lowercase letters, digits, hyphens and underscores, start with a lowercase letter, and end with a lowercase letter or a digit.
        """
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> Optional[pulumi.Input['GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs']]:
        """
        Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
        """
        return pulumi.get(self, "worker_config")

    @worker_config.setter
    def worker_config(self, value: Optional[pulumi.Input['GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs']]):
        pulumi.set(self, "worker_config", value)

    @property
    @pulumi.getter(name="workerCount")
    def worker_count(self) -> Optional[pulumi.Input[str]]:
        """
        The desired number of workers in the worker pool. Must be a value between 0 and 15000.
        """
        return pulumi.get(self, "worker_count")

    @worker_count.setter
    def worker_count(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "worker_count", value)


class WorkerPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscale: Optional[pulumi.Input[pulumi.InputType['GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleArgs']]] = None,
                 channel: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 worker_config: Optional[pulumi.Input[pulumi.InputType['GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs']]] = None,
                 worker_count: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new worker pool with a specified size and configuration. Returns a long running operation which contains a worker pool on completion. While the long running operation is in progress, any call to `GetWorkerPool` returns a worker pool in state `CREATING`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleArgs']] autoscale: The autoscale policy to apply on a pool.
        :param pulumi.Input[str] channel: Channel specifies the release channel of the pool.
        :param pulumi.Input[str] name: WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
        :param pulumi.Input[str] parent: Resource name of the instance in which to create the new worker pool. Format: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]`.
        :param pulumi.Input[str] pool_id: ID of the created worker pool. A valid pool ID must: be 6-50 characters long, contain only lowercase letters, digits, hyphens and underscores, start with a lowercase letter, and end with a lowercase letter or a digit.
        :param pulumi.Input[pulumi.InputType['GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs']] worker_config: Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
        :param pulumi.Input[str] worker_count: The desired number of workers in the worker pool. Must be a value between 0 and 15000.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkerPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new worker pool with a specified size and configuration. Returns a long running operation which contains a worker pool on completion. While the long running operation is in progress, any call to `GetWorkerPool` returns a worker pool in state `CREATING`.

        :param str resource_name: The name of the resource.
        :param WorkerPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkerPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscale: Optional[pulumi.Input[pulumi.InputType['GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleArgs']]] = None,
                 channel: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 worker_config: Optional[pulumi.Input[pulumi.InputType['GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs']]] = None,
                 worker_count: Optional[pulumi.Input[str]] = None,
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
            __props__ = WorkerPoolArgs.__new__(WorkerPoolArgs)

            __props__.__dict__["autoscale"] = autoscale
            __props__.__dict__["channel"] = channel
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["parent"] = parent
            __props__.__dict__["pool_id"] = pool_id
            __props__.__dict__["project"] = project
            __props__.__dict__["worker_config"] = worker_config
            __props__.__dict__["worker_count"] = worker_count
            __props__.__dict__["state"] = None
        super(WorkerPool, __self__).__init__(
            'google-native:remotebuildexecution/v1alpha:WorkerPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WorkerPool':
        """
        Get an existing WorkerPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WorkerPoolArgs.__new__(WorkerPoolArgs)

        __props__.__dict__["autoscale"] = None
        __props__.__dict__["channel"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["worker_config"] = None
        __props__.__dict__["worker_count"] = None
        return WorkerPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def autoscale(self) -> pulumi.Output['outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponse']:
        """
        The autoscale policy to apply on a pool.
        """
        return pulumi.get(self, "autoscale")

    @property
    @pulumi.getter
    def channel(self) -> pulumi.Output[str]:
        """
        Channel specifies the release channel of the pool.
        """
        return pulumi.get(self, "channel")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the worker pool.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> pulumi.Output['outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponse']:
        """
        Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
        """
        return pulumi.get(self, "worker_config")

    @property
    @pulumi.getter(name="workerCount")
    def worker_count(self) -> pulumi.Output[str]:
        """
        The desired number of workers in the worker pool. Must be a value between 0 and 15000.
        """
        return pulumi.get(self, "worker_count")

