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

__all__ = ['TaskArgs', 'Task']

@pulumi.input_type
class TaskArgs:
    def __init__(__self__, *,
                 queue_id: pulumi.Input[str],
                 app_engine_http_request: Optional[pulumi.Input['AppEngineHttpRequestArgs']] = None,
                 dispatch_deadline: Optional[pulumi.Input[str]] = None,
                 http_request: Optional[pulumi.Input['HttpRequestArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 response_view: Optional[pulumi.Input['TaskResponseView']] = None,
                 schedule_time: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Task resource.
        :param pulumi.Input['AppEngineHttpRequestArgs'] app_engine_http_request: HTTP request that is sent to the App Engine app handler. An App Engine task is a task that has AppEngineHttpRequest set.
        :param pulumi.Input[str] dispatch_deadline: The deadline for requests sent to the worker. If the worker does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the task according to the RetryConfig. Note that when the request is cancelled, Cloud Tasks will stop listening for the response, but whether the worker stops processing depends on the worker. For example, if the worker is stuck, it may not react to cancelled requests. The default and maximum values depend on the type of request: * For HTTP tasks, the default is 10 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine tasks, 0 indicates that the request has the default deadline. The default deadline depends on the [scaling type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling) of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. Regardless of the task's `dispatch_deadline`, the app handler will not run for longer than than the service's timeout. We recommend setting the `dispatch_deadline` to at most a few seconds more than the app handler's timeout. For more information see [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts). `dispatch_deadline` will be truncated to the nearest millisecond. The deadline is an approximate deadline.
        :param pulumi.Input['HttpRequestArgs'] http_request: HTTP request that is sent to the worker. An HTTP task is a task that has HttpRequest set.
        :param pulumi.Input[str] name: Optionally caller-specified in CreateTask. The task name. The task name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the task's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters. * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
        :param pulumi.Input['TaskResponseView'] response_view: The response_view specifies which subset of the Task will be returned. By default response_view is BASIC; not all information is retrieved by default because some data, such as payloads, might be desirable to return only when needed because of its large size or because of the sensitivity of data that it contains. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Task resource.
        :param pulumi.Input[str] schedule_time: The time when the task is scheduled to be attempted or retried. `schedule_time` will be truncated to the nearest microsecond.
        """
        pulumi.set(__self__, "queue_id", queue_id)
        if app_engine_http_request is not None:
            pulumi.set(__self__, "app_engine_http_request", app_engine_http_request)
        if dispatch_deadline is not None:
            pulumi.set(__self__, "dispatch_deadline", dispatch_deadline)
        if http_request is not None:
            pulumi.set(__self__, "http_request", http_request)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if response_view is not None:
            pulumi.set(__self__, "response_view", response_view)
        if schedule_time is not None:
            pulumi.set(__self__, "schedule_time", schedule_time)

    @property
    @pulumi.getter(name="queueId")
    def queue_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "queue_id")

    @queue_id.setter
    def queue_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "queue_id", value)

    @property
    @pulumi.getter(name="appEngineHttpRequest")
    def app_engine_http_request(self) -> Optional[pulumi.Input['AppEngineHttpRequestArgs']]:
        """
        HTTP request that is sent to the App Engine app handler. An App Engine task is a task that has AppEngineHttpRequest set.
        """
        return pulumi.get(self, "app_engine_http_request")

    @app_engine_http_request.setter
    def app_engine_http_request(self, value: Optional[pulumi.Input['AppEngineHttpRequestArgs']]):
        pulumi.set(self, "app_engine_http_request", value)

    @property
    @pulumi.getter(name="dispatchDeadline")
    def dispatch_deadline(self) -> Optional[pulumi.Input[str]]:
        """
        The deadline for requests sent to the worker. If the worker does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the task according to the RetryConfig. Note that when the request is cancelled, Cloud Tasks will stop listening for the response, but whether the worker stops processing depends on the worker. For example, if the worker is stuck, it may not react to cancelled requests. The default and maximum values depend on the type of request: * For HTTP tasks, the default is 10 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine tasks, 0 indicates that the request has the default deadline. The default deadline depends on the [scaling type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling) of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. Regardless of the task's `dispatch_deadline`, the app handler will not run for longer than than the service's timeout. We recommend setting the `dispatch_deadline` to at most a few seconds more than the app handler's timeout. For more information see [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts). `dispatch_deadline` will be truncated to the nearest millisecond. The deadline is an approximate deadline.
        """
        return pulumi.get(self, "dispatch_deadline")

    @dispatch_deadline.setter
    def dispatch_deadline(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dispatch_deadline", value)

    @property
    @pulumi.getter(name="httpRequest")
    def http_request(self) -> Optional[pulumi.Input['HttpRequestArgs']]:
        """
        HTTP request that is sent to the worker. An HTTP task is a task that has HttpRequest set.
        """
        return pulumi.get(self, "http_request")

    @http_request.setter
    def http_request(self, value: Optional[pulumi.Input['HttpRequestArgs']]):
        pulumi.set(self, "http_request", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Optionally caller-specified in CreateTask. The task name. The task name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the task's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters. * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="responseView")
    def response_view(self) -> Optional[pulumi.Input['TaskResponseView']]:
        """
        The response_view specifies which subset of the Task will be returned. By default response_view is BASIC; not all information is retrieved by default because some data, such as payloads, might be desirable to return only when needed because of its large size or because of the sensitivity of data that it contains. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Task resource.
        """
        return pulumi.get(self, "response_view")

    @response_view.setter
    def response_view(self, value: Optional[pulumi.Input['TaskResponseView']]):
        pulumi.set(self, "response_view", value)

    @property
    @pulumi.getter(name="scheduleTime")
    def schedule_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the task is scheduled to be attempted or retried. `schedule_time` will be truncated to the nearest microsecond.
        """
        return pulumi.get(self, "schedule_time")

    @schedule_time.setter
    def schedule_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_time", value)


class Task(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_engine_http_request: Optional[pulumi.Input[pulumi.InputType['AppEngineHttpRequestArgs']]] = None,
                 dispatch_deadline: Optional[pulumi.Input[str]] = None,
                 http_request: Optional[pulumi.Input[pulumi.InputType['HttpRequestArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 queue_id: Optional[pulumi.Input[str]] = None,
                 response_view: Optional[pulumi.Input['TaskResponseView']] = None,
                 schedule_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a task and adds it to a queue. Tasks cannot be updated after creation; there is no UpdateTask command. * The maximum task size is 100KB.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AppEngineHttpRequestArgs']] app_engine_http_request: HTTP request that is sent to the App Engine app handler. An App Engine task is a task that has AppEngineHttpRequest set.
        :param pulumi.Input[str] dispatch_deadline: The deadline for requests sent to the worker. If the worker does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the task according to the RetryConfig. Note that when the request is cancelled, Cloud Tasks will stop listening for the response, but whether the worker stops processing depends on the worker. For example, if the worker is stuck, it may not react to cancelled requests. The default and maximum values depend on the type of request: * For HTTP tasks, the default is 10 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine tasks, 0 indicates that the request has the default deadline. The default deadline depends on the [scaling type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling) of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. Regardless of the task's `dispatch_deadline`, the app handler will not run for longer than than the service's timeout. We recommend setting the `dispatch_deadline` to at most a few seconds more than the app handler's timeout. For more information see [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts). `dispatch_deadline` will be truncated to the nearest millisecond. The deadline is an approximate deadline.
        :param pulumi.Input[pulumi.InputType['HttpRequestArgs']] http_request: HTTP request that is sent to the worker. An HTTP task is a task that has HttpRequest set.
        :param pulumi.Input[str] name: Optionally caller-specified in CreateTask. The task name. The task name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the task's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters. * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
        :param pulumi.Input['TaskResponseView'] response_view: The response_view specifies which subset of the Task will be returned. By default response_view is BASIC; not all information is retrieved by default because some data, such as payloads, might be desirable to return only when needed because of its large size or because of the sensitivity of data that it contains. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Task resource.
        :param pulumi.Input[str] schedule_time: The time when the task is scheduled to be attempted or retried. `schedule_time` will be truncated to the nearest microsecond.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TaskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a task and adds it to a queue. Tasks cannot be updated after creation; there is no UpdateTask command. * The maximum task size is 100KB.

        :param str resource_name: The name of the resource.
        :param TaskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TaskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_engine_http_request: Optional[pulumi.Input[pulumi.InputType['AppEngineHttpRequestArgs']]] = None,
                 dispatch_deadline: Optional[pulumi.Input[str]] = None,
                 http_request: Optional[pulumi.Input[pulumi.InputType['HttpRequestArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 queue_id: Optional[pulumi.Input[str]] = None,
                 response_view: Optional[pulumi.Input['TaskResponseView']] = None,
                 schedule_time: Optional[pulumi.Input[str]] = None,
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
            __props__ = TaskArgs.__new__(TaskArgs)

            __props__.__dict__["app_engine_http_request"] = app_engine_http_request
            __props__.__dict__["dispatch_deadline"] = dispatch_deadline
            __props__.__dict__["http_request"] = http_request
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if queue_id is None and not opts.urn:
                raise TypeError("Missing required property 'queue_id'")
            __props__.__dict__["queue_id"] = queue_id
            __props__.__dict__["response_view"] = response_view
            __props__.__dict__["schedule_time"] = schedule_time
            __props__.__dict__["create_time"] = None
            __props__.__dict__["dispatch_count"] = None
            __props__.__dict__["first_attempt"] = None
            __props__.__dict__["last_attempt"] = None
            __props__.__dict__["response_count"] = None
            __props__.__dict__["view"] = None
        super(Task, __self__).__init__(
            'google-native:cloudtasks/v2:Task',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Task':
        """
        Get an existing Task resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TaskArgs.__new__(TaskArgs)

        __props__.__dict__["app_engine_http_request"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["dispatch_count"] = None
        __props__.__dict__["dispatch_deadline"] = None
        __props__.__dict__["first_attempt"] = None
        __props__.__dict__["http_request"] = None
        __props__.__dict__["last_attempt"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["response_count"] = None
        __props__.__dict__["schedule_time"] = None
        __props__.__dict__["view"] = None
        return Task(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appEngineHttpRequest")
    def app_engine_http_request(self) -> pulumi.Output['outputs.AppEngineHttpRequestResponse']:
        """
        HTTP request that is sent to the App Engine app handler. An App Engine task is a task that has AppEngineHttpRequest set.
        """
        return pulumi.get(self, "app_engine_http_request")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time that the task was created. `create_time` will be truncated to the nearest second.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dispatchCount")
    def dispatch_count(self) -> pulumi.Output[int]:
        """
        The number of attempts dispatched. This count includes attempts which have been dispatched but haven't received a response.
        """
        return pulumi.get(self, "dispatch_count")

    @property
    @pulumi.getter(name="dispatchDeadline")
    def dispatch_deadline(self) -> pulumi.Output[str]:
        """
        The deadline for requests sent to the worker. If the worker does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the task according to the RetryConfig. Note that when the request is cancelled, Cloud Tasks will stop listening for the response, but whether the worker stops processing depends on the worker. For example, if the worker is stuck, it may not react to cancelled requests. The default and maximum values depend on the type of request: * For HTTP tasks, the default is 10 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine tasks, 0 indicates that the request has the default deadline. The default deadline depends on the [scaling type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling) of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. Regardless of the task's `dispatch_deadline`, the app handler will not run for longer than than the service's timeout. We recommend setting the `dispatch_deadline` to at most a few seconds more than the app handler's timeout. For more information see [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts). `dispatch_deadline` will be truncated to the nearest millisecond. The deadline is an approximate deadline.
        """
        return pulumi.get(self, "dispatch_deadline")

    @property
    @pulumi.getter(name="firstAttempt")
    def first_attempt(self) -> pulumi.Output['outputs.AttemptResponse']:
        """
        The status of the task's first attempt. Only dispatch_time will be set. The other Attempt information is not retained by Cloud Tasks.
        """
        return pulumi.get(self, "first_attempt")

    @property
    @pulumi.getter(name="httpRequest")
    def http_request(self) -> pulumi.Output['outputs.HttpRequestResponse']:
        """
        HTTP request that is sent to the worker. An HTTP task is a task that has HttpRequest set.
        """
        return pulumi.get(self, "http_request")

    @property
    @pulumi.getter(name="lastAttempt")
    def last_attempt(self) -> pulumi.Output['outputs.AttemptResponse']:
        """
        The status of the task's last attempt.
        """
        return pulumi.get(self, "last_attempt")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Optionally caller-specified in CreateTask. The task name. The task name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the task's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters. * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="responseCount")
    def response_count(self) -> pulumi.Output[int]:
        """
        The number of attempts which have received a response.
        """
        return pulumi.get(self, "response_count")

    @property
    @pulumi.getter(name="scheduleTime")
    def schedule_time(self) -> pulumi.Output[str]:
        """
        The time when the task is scheduled to be attempted or retried. `schedule_time` will be truncated to the nearest microsecond.
        """
        return pulumi.get(self, "schedule_time")

    @property
    @pulumi.getter
    def view(self) -> pulumi.Output[str]:
        """
        The view specifies which subset of the Task has been returned.
        """
        return pulumi.get(self, "view")

