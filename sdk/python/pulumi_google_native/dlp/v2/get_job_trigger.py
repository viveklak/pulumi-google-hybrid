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
    'GetJobTriggerResult',
    'AwaitableGetJobTriggerResult',
    'get_job_trigger',
    'get_job_trigger_output',
]

@pulumi.output_type
class GetJobTriggerResult:
    def __init__(__self__, create_time=None, description=None, display_name=None, errors=None, inspect_job=None, last_run_time=None, name=None, status=None, triggers=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if errors and not isinstance(errors, list):
            raise TypeError("Expected argument 'errors' to be a list")
        pulumi.set(__self__, "errors", errors)
        if inspect_job and not isinstance(inspect_job, dict):
            raise TypeError("Expected argument 'inspect_job' to be a dict")
        pulumi.set(__self__, "inspect_job", inspect_job)
        if last_run_time and not isinstance(last_run_time, str):
            raise TypeError("Expected argument 'last_run_time' to be a str")
        pulumi.set(__self__, "last_run_time", last_run_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if triggers and not isinstance(triggers, list):
            raise TypeError("Expected argument 'triggers' to be a list")
        pulumi.set(__self__, "triggers", triggers)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation timestamp of a triggeredJob.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        User provided description (max 256 chars)
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name (max 100 chars)
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def errors(self) -> Sequence['outputs.GooglePrivacyDlpV2ErrorResponse']:
        """
        A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter(name="inspectJob")
    def inspect_job(self) -> 'outputs.GooglePrivacyDlpV2InspectJobConfigResponse':
        """
        For inspect jobs, a snapshot of the configuration.
        """
        return pulumi.get(self, "inspect_job")

    @property
    @pulumi.getter(name="lastRunTime")
    def last_run_time(self) -> str:
        """
        The timestamp of the last time this trigger executed.
        """
        return pulumi.get(self, "last_run_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        A status for this trigger.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def triggers(self) -> Sequence['outputs.GooglePrivacyDlpV2TriggerResponse']:
        """
        A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
        """
        return pulumi.get(self, "triggers")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update timestamp of a triggeredJob.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetJobTriggerResult(GetJobTriggerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJobTriggerResult(
            create_time=self.create_time,
            description=self.description,
            display_name=self.display_name,
            errors=self.errors,
            inspect_job=self.inspect_job,
            last_run_time=self.last_run_time,
            name=self.name,
            status=self.status,
            triggers=self.triggers,
            update_time=self.update_time)


def get_job_trigger(job_trigger_id: Optional[str] = None,
                    location: Optional[str] = None,
                    project: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJobTriggerResult:
    """
    Gets a job trigger. See https://cloud.google.com/dlp/docs/creating-job-triggers to learn more.
    """
    __args__ = dict()
    __args__['jobTriggerId'] = job_trigger_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:dlp/v2:getJobTrigger', __args__, opts=opts, typ=GetJobTriggerResult).value

    return AwaitableGetJobTriggerResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        display_name=__ret__.display_name,
        errors=__ret__.errors,
        inspect_job=__ret__.inspect_job,
        last_run_time=__ret__.last_run_time,
        name=__ret__.name,
        status=__ret__.status,
        triggers=__ret__.triggers,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_job_trigger)
def get_job_trigger_output(job_trigger_id: Optional[pulumi.Input[str]] = None,
                           location: Optional[pulumi.Input[str]] = None,
                           project: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetJobTriggerResult]:
    """
    Gets a job trigger. See https://cloud.google.com/dlp/docs/creating-job-triggers to learn more.
    """
    ...
