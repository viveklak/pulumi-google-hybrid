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

__all__ = ['JobArgs', 'Job']

@pulumi.input_type
class JobArgs:
    def __init__(__self__, *,
                 job_id: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 configuration: Optional[pulumi.Input['JobConfigurationArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 job_reference: Optional[pulumi.Input['JobReferenceArgs']] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 statistics: Optional[pulumi.Input['JobStatisticsArgs']] = None,
                 status: Optional[pulumi.Input['JobStatusArgs']] = None,
                 user_email: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Job resource.
        :param pulumi.Input['JobConfigurationArgs'] configuration: [Required] Describes the job configuration.
        :param pulumi.Input[str] etag: [Output-only] A hash of this resource.
        :param pulumi.Input[str] id: [Output-only] Opaque ID field of the job
        :param pulumi.Input['JobReferenceArgs'] job_reference: [Optional] Reference describing the unique-per-user name of the job.
        :param pulumi.Input[str] kind: [Output-only] The type of the resource.
        :param pulumi.Input[str] self_link: [Output-only] A URL that can be used to access this resource again.
        :param pulumi.Input['JobStatisticsArgs'] statistics: [Output-only] Information about the job, including starting time and ending time of the job.
        :param pulumi.Input['JobStatusArgs'] status: [Output-only] The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
        :param pulumi.Input[str] user_email: [Output-only] Email address of the user who ran the job.
        """
        pulumi.set(__self__, "job_id", job_id)
        pulumi.set(__self__, "project_id", project_id)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if job_reference is not None:
            pulumi.set(__self__, "job_reference", job_reference)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if statistics is not None:
            pulumi.set(__self__, "statistics", statistics)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if user_email is not None:
            pulumi.set(__self__, "user_email", user_email)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "job_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input['JobConfigurationArgs']]:
        """
        [Required] Describes the job configuration.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input['JobConfigurationArgs']]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        [Output-only] A hash of this resource.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        [Output-only] Opaque ID field of the job
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="jobReference")
    def job_reference(self) -> Optional[pulumi.Input['JobReferenceArgs']]:
        """
        [Optional] Reference describing the unique-per-user name of the job.
        """
        return pulumi.get(self, "job_reference")

    @job_reference.setter
    def job_reference(self, value: Optional[pulumi.Input['JobReferenceArgs']]):
        pulumi.set(self, "job_reference", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        [Output-only] The type of the resource.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        [Output-only] A URL that can be used to access this resource again.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter
    def statistics(self) -> Optional[pulumi.Input['JobStatisticsArgs']]:
        """
        [Output-only] Information about the job, including starting time and ending time of the job.
        """
        return pulumi.get(self, "statistics")

    @statistics.setter
    def statistics(self, value: Optional[pulumi.Input['JobStatisticsArgs']]):
        pulumi.set(self, "statistics", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['JobStatusArgs']]:
        """
        [Output-only] The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['JobStatusArgs']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def user_email(self) -> Optional[pulumi.Input[str]]:
        """
        [Output-only] Email address of the user who ran the job.
        """
        return pulumi.get(self, "user_email")

    @user_email.setter
    def user_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_email", value)


class Job(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['JobConfigurationArgs']]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 job_reference: Optional[pulumi.Input[pulumi.InputType['JobReferenceArgs']]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 statistics: Optional[pulumi.Input[pulumi.InputType['JobStatisticsArgs']]] = None,
                 status: Optional[pulumi.Input[pulumi.InputType['JobStatusArgs']]] = None,
                 user_email: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Starts a new asynchronous job. Requires the Can View project role.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['JobConfigurationArgs']] configuration: [Required] Describes the job configuration.
        :param pulumi.Input[str] etag: [Output-only] A hash of this resource.
        :param pulumi.Input[str] id: [Output-only] Opaque ID field of the job
        :param pulumi.Input[pulumi.InputType['JobReferenceArgs']] job_reference: [Optional] Reference describing the unique-per-user name of the job.
        :param pulumi.Input[str] kind: [Output-only] The type of the resource.
        :param pulumi.Input[str] self_link: [Output-only] A URL that can be used to access this resource again.
        :param pulumi.Input[pulumi.InputType['JobStatisticsArgs']] statistics: [Output-only] Information about the job, including starting time and ending time of the job.
        :param pulumi.Input[pulumi.InputType['JobStatusArgs']] status: [Output-only] The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
        :param pulumi.Input[str] user_email: [Output-only] Email address of the user who ran the job.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Starts a new asynchronous job. Requires the Can View project role.

        :param str resource_name: The name of the resource.
        :param JobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['JobConfigurationArgs']]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 job_reference: Optional[pulumi.Input[pulumi.InputType['JobReferenceArgs']]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 statistics: Optional[pulumi.Input[pulumi.InputType['JobStatisticsArgs']]] = None,
                 status: Optional[pulumi.Input[pulumi.InputType['JobStatusArgs']]] = None,
                 user_email: Optional[pulumi.Input[str]] = None,
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
            __props__ = JobArgs.__new__(JobArgs)

            __props__.__dict__["configuration"] = configuration
            __props__.__dict__["etag"] = etag
            __props__.__dict__["id"] = id
            if job_id is None and not opts.urn:
                raise TypeError("Missing required property 'job_id'")
            __props__.__dict__["job_id"] = job_id
            __props__.__dict__["job_reference"] = job_reference
            __props__.__dict__["kind"] = kind
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["self_link"] = self_link
            __props__.__dict__["statistics"] = statistics
            __props__.__dict__["status"] = status
            __props__.__dict__["user_email"] = user_email
        super(Job, __self__).__init__(
            'gcp-native:bigquery/v2:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Job':
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = JobArgs.__new__(JobArgs)

        __props__.__dict__["configuration"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["job_reference"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["statistics"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["user_email"] = None
        return Job(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output['outputs.JobConfigurationResponse']:
        """
        [Required] Describes the job configuration.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        [Output-only] A hash of this resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="jobReference")
    def job_reference(self) -> pulumi.Output['outputs.JobReferenceResponse']:
        """
        [Optional] Reference describing the unique-per-user name of the job.
        """
        return pulumi.get(self, "job_reference")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        [Output-only] The type of the resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        [Output-only] A URL that can be used to access this resource again.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def statistics(self) -> pulumi.Output['outputs.JobStatisticsResponse']:
        """
        [Output-only] Information about the job, including starting time and ending time of the job.
        """
        return pulumi.get(self, "statistics")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.JobStatusResponse']:
        """
        [Output-only] The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def user_email(self) -> pulumi.Output[str]:
        """
        [Output-only] Email address of the user who ran the job.
        """
        return pulumi.get(self, "user_email")

