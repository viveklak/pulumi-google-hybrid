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

__all__ = ['ConfigWaiterArgs', 'ConfigWaiter']

@pulumi.input_type
class ConfigWaiterArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[str],
                 project: pulumi.Input[str],
                 create_time: Optional[pulumi.Input[str]] = None,
                 done: Optional[pulumi.Input[bool]] = None,
                 error: Optional[pulumi.Input['StatusArgs']] = None,
                 failure: Optional[pulumi.Input['EndConditionArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 success: Optional[pulumi.Input['EndConditionArgs']] = None,
                 timeout: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConfigWaiter resource.
        :param pulumi.Input[str] create_time: The instant at which this Waiter resource was created. Adding the value of `timeout` to this instant yields the timeout deadline for the waiter.
        :param pulumi.Input[bool] done: If the value is `false`, it means the waiter is still waiting for one of its conditions to be met. If true, the waiter has finished. If the waiter finished due to a timeout or failure, `error` will be set.
        :param pulumi.Input['StatusArgs'] error: If the waiter ended due to a failure or timeout, this value will be set.
        :param pulumi.Input['EndConditionArgs'] failure: [Optional] The failure condition of this waiter. If this condition is met, `done` will be set to `true` and the `error` code will be set to `ABORTED`. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated. This value is optional; if no failure condition is set, the only failure scenario will be a timeout.
        :param pulumi.Input[str] name: The name of the Waiter resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/waiters/[WAITER_NAME] The `[PROJECT_ID]` must be a valid Google Cloud project ID, the `[CONFIG_NAME]` must be a valid RuntimeConfig resource, the `[WAITER_NAME]` must match RFC 1035 segment specification, and the length of `[WAITER_NAME]` must be less than 64 bytes. After you create a Waiter resource, you cannot change the resource name.
        :param pulumi.Input['EndConditionArgs'] success: [Required] The success condition. If this condition is met, `done` will be set to `true` and the `error` value will remain unset. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated.
        :param pulumi.Input[str] timeout: [Required] Specifies the timeout of the waiter in seconds, beginning from the instant that `waiters().create` method is called. If this time elapses before the success or failure conditions are met, the waiter fails and sets the `error` code to `DEADLINE_EXCEEDED`.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "project", project)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if done is not None:
            pulumi.set(__self__, "done", done)
        if error is not None:
            pulumi.set(__self__, "error", error)
        if failure is not None:
            pulumi.set(__self__, "failure", failure)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if success is not None:
            pulumi.set(__self__, "success", success)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_id", value)

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
        The instant at which this Waiter resource was created. Adding the value of `timeout` to this instant yields the timeout deadline for the waiter.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def done(self) -> Optional[pulumi.Input[bool]]:
        """
        If the value is `false`, it means the waiter is still waiting for one of its conditions to be met. If true, the waiter has finished. If the waiter finished due to a timeout or failure, `error` will be set.
        """
        return pulumi.get(self, "done")

    @done.setter
    def done(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "done", value)

    @property
    @pulumi.getter
    def error(self) -> Optional[pulumi.Input['StatusArgs']]:
        """
        If the waiter ended due to a failure or timeout, this value will be set.
        """
        return pulumi.get(self, "error")

    @error.setter
    def error(self, value: Optional[pulumi.Input['StatusArgs']]):
        pulumi.set(self, "error", value)

    @property
    @pulumi.getter
    def failure(self) -> Optional[pulumi.Input['EndConditionArgs']]:
        """
        [Optional] The failure condition of this waiter. If this condition is met, `done` will be set to `true` and the `error` code will be set to `ABORTED`. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated. This value is optional; if no failure condition is set, the only failure scenario will be a timeout.
        """
        return pulumi.get(self, "failure")

    @failure.setter
    def failure(self, value: Optional[pulumi.Input['EndConditionArgs']]):
        pulumi.set(self, "failure", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Waiter resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/waiters/[WAITER_NAME] The `[PROJECT_ID]` must be a valid Google Cloud project ID, the `[CONFIG_NAME]` must be a valid RuntimeConfig resource, the `[WAITER_NAME]` must match RFC 1035 segment specification, and the length of `[WAITER_NAME]` must be less than 64 bytes. After you create a Waiter resource, you cannot change the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def success(self) -> Optional[pulumi.Input['EndConditionArgs']]:
        """
        [Required] The success condition. If this condition is met, `done` will be set to `true` and the `error` value will remain unset. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated.
        """
        return pulumi.get(self, "success")

    @success.setter
    def success(self, value: Optional[pulumi.Input['EndConditionArgs']]):
        pulumi.set(self, "success", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[str]]:
        """
        [Required] Specifies the timeout of the waiter in seconds, beginning from the instant that `waiters().create` method is called. If this time elapses before the success or failure conditions are met, the waiter fails and sets the `error` code to `DEADLINE_EXCEEDED`.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout", value)


class ConfigWaiter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 done: Optional[pulumi.Input[bool]] = None,
                 error: Optional[pulumi.Input[pulumi.InputType['StatusArgs']]] = None,
                 failure: Optional[pulumi.Input[pulumi.InputType['EndConditionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 success: Optional[pulumi.Input[pulumi.InputType['EndConditionArgs']]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a Waiter resource. This operation returns a long-running Operation resource which can be polled for completion. However, a waiter with the given name will exist (and can be retrieved) prior to the operation completing. If the operation fails, the failed Waiter resource will still exist and must be deleted prior to subsequent creation attempts.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The instant at which this Waiter resource was created. Adding the value of `timeout` to this instant yields the timeout deadline for the waiter.
        :param pulumi.Input[bool] done: If the value is `false`, it means the waiter is still waiting for one of its conditions to be met. If true, the waiter has finished. If the waiter finished due to a timeout or failure, `error` will be set.
        :param pulumi.Input[pulumi.InputType['StatusArgs']] error: If the waiter ended due to a failure or timeout, this value will be set.
        :param pulumi.Input[pulumi.InputType['EndConditionArgs']] failure: [Optional] The failure condition of this waiter. If this condition is met, `done` will be set to `true` and the `error` code will be set to `ABORTED`. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated. This value is optional; if no failure condition is set, the only failure scenario will be a timeout.
        :param pulumi.Input[str] name: The name of the Waiter resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/waiters/[WAITER_NAME] The `[PROJECT_ID]` must be a valid Google Cloud project ID, the `[CONFIG_NAME]` must be a valid RuntimeConfig resource, the `[WAITER_NAME]` must match RFC 1035 segment specification, and the length of `[WAITER_NAME]` must be less than 64 bytes. After you create a Waiter resource, you cannot change the resource name.
        :param pulumi.Input[pulumi.InputType['EndConditionArgs']] success: [Required] The success condition. If this condition is met, `done` will be set to `true` and the `error` value will remain unset. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated.
        :param pulumi.Input[str] timeout: [Required] Specifies the timeout of the waiter in seconds, beginning from the instant that `waiters().create` method is called. If this time elapses before the success or failure conditions are met, the waiter fails and sets the `error` code to `DEADLINE_EXCEEDED`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigWaiterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Waiter resource. This operation returns a long-running Operation resource which can be polled for completion. However, a waiter with the given name will exist (and can be retrieved) prior to the operation completing. If the operation fails, the failed Waiter resource will still exist and must be deleted prior to subsequent creation attempts.

        :param str resource_name: The name of the resource.
        :param ConfigWaiterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigWaiterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 done: Optional[pulumi.Input[bool]] = None,
                 error: Optional[pulumi.Input[pulumi.InputType['StatusArgs']]] = None,
                 failure: Optional[pulumi.Input[pulumi.InputType['EndConditionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 success: Optional[pulumi.Input[pulumi.InputType['EndConditionArgs']]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
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
            __props__ = ConfigWaiterArgs.__new__(ConfigWaiterArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            __props__.__dict__["create_time"] = create_time
            __props__.__dict__["done"] = done
            __props__.__dict__["error"] = error
            __props__.__dict__["failure"] = failure
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["success"] = success
            __props__.__dict__["timeout"] = timeout
        super(ConfigWaiter, __self__).__init__(
            'google-native:runtimeconfig/v1beta1:ConfigWaiter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConfigWaiter':
        """
        Get an existing ConfigWaiter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConfigWaiterArgs.__new__(ConfigWaiterArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["done"] = None
        __props__.__dict__["error"] = None
        __props__.__dict__["failure"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["success"] = None
        __props__.__dict__["timeout"] = None
        return ConfigWaiter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The instant at which this Waiter resource was created. Adding the value of `timeout` to this instant yields the timeout deadline for the waiter.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def done(self) -> pulumi.Output[bool]:
        """
        If the value is `false`, it means the waiter is still waiting for one of its conditions to be met. If true, the waiter has finished. If the waiter finished due to a timeout or failure, `error` will be set.
        """
        return pulumi.get(self, "done")

    @property
    @pulumi.getter
    def error(self) -> pulumi.Output['outputs.StatusResponse']:
        """
        If the waiter ended due to a failure or timeout, this value will be set.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def failure(self) -> pulumi.Output['outputs.EndConditionResponse']:
        """
        [Optional] The failure condition of this waiter. If this condition is met, `done` will be set to `true` and the `error` code will be set to `ABORTED`. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated. This value is optional; if no failure condition is set, the only failure scenario will be a timeout.
        """
        return pulumi.get(self, "failure")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Waiter resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/waiters/[WAITER_NAME] The `[PROJECT_ID]` must be a valid Google Cloud project ID, the `[CONFIG_NAME]` must be a valid RuntimeConfig resource, the `[WAITER_NAME]` must match RFC 1035 segment specification, and the length of `[WAITER_NAME]` must be less than 64 bytes. After you create a Waiter resource, you cannot change the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def success(self) -> pulumi.Output['outputs.EndConditionResponse']:
        """
        [Required] The success condition. If this condition is met, `done` will be set to `true` and the `error` value will remain unset. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated.
        """
        return pulumi.get(self, "success")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[str]:
        """
        [Required] Specifies the timeout of the waiter in seconds, beginning from the instant that `waiters().create` method is called. If this time elapses before the success or failure conditions are met, the waiter fails and sets the `error` code to `DEADLINE_EXCEEDED`.
        """
        return pulumi.get(self, "timeout")

