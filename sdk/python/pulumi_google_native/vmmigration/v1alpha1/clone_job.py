# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = ['CloneJobArgs', 'CloneJob']

@pulumi.input_type
class CloneJobArgs:
    def __init__(__self__, *,
                 clone_job_id: pulumi.Input[str],
                 migrating_vm_id: pulumi.Input[str],
                 source_id: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CloneJob resource.
        :param pulumi.Input[str] clone_job_id: Required. The clone job identifier.
        :param pulumi.Input[str] request_id: A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        pulumi.set(__self__, "clone_job_id", clone_job_id)
        pulumi.set(__self__, "migrating_vm_id", migrating_vm_id)
        pulumi.set(__self__, "source_id", source_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)

    @property
    @pulumi.getter(name="cloneJobId")
    def clone_job_id(self) -> pulumi.Input[str]:
        """
        Required. The clone job identifier.
        """
        return pulumi.get(self, "clone_job_id")

    @clone_job_id.setter
    def clone_job_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "clone_job_id", value)

    @property
    @pulumi.getter(name="migratingVmId")
    def migrating_vm_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "migrating_vm_id")

    @migrating_vm_id.setter
    def migrating_vm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "migrating_vm_id", value)

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source_id")

    @source_id.setter
    def source_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)


class CloneJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clone_job_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 migrating_vm_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 source_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Initiates a Clone of a specific migrating VM.
        Auto-naming is currently not supported for this resource.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] clone_job_id: Required. The clone job identifier.
        :param pulumi.Input[str] request_id: A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloneJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Initiates a Clone of a specific migrating VM.
        Auto-naming is currently not supported for this resource.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param CloneJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloneJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clone_job_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 migrating_vm_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 source_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = CloneJobArgs.__new__(CloneJobArgs)

            if clone_job_id is None and not opts.urn:
                raise TypeError("Missing required property 'clone_job_id'")
            __props__.__dict__["clone_job_id"] = clone_job_id
            __props__.__dict__["location"] = location
            if migrating_vm_id is None and not opts.urn:
                raise TypeError("Missing required property 'migrating_vm_id'")
            __props__.__dict__["migrating_vm_id"] = migrating_vm_id
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            if source_id is None and not opts.urn:
                raise TypeError("Missing required property 'source_id'")
            __props__.__dict__["source_id"] = source_id
            __props__.__dict__["compute_engine_target_details"] = None
            __props__.__dict__["compute_engine_vm_details"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["error"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["state_time"] = None
            __props__.__dict__["target_details"] = None
        super(CloneJob, __self__).__init__(
            'google-native:vmmigration/v1alpha1:CloneJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CloneJob':
        """
        Get an existing CloneJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CloneJobArgs.__new__(CloneJobArgs)

        __props__.__dict__["compute_engine_target_details"] = None
        __props__.__dict__["compute_engine_vm_details"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["error"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["state_time"] = None
        __props__.__dict__["target_details"] = None
        return CloneJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="computeEngineTargetDetails")
    def compute_engine_target_details(self) -> pulumi.Output['outputs.ComputeEngineTargetDetailsResponse']:
        """
        Details of the target VM in Compute Engine.
        """
        return pulumi.get(self, "compute_engine_target_details")

    @property
    @pulumi.getter(name="computeEngineVmDetails")
    def compute_engine_vm_details(self) -> pulumi.Output['outputs.TargetVMDetailsResponse']:
        """
        Details of the VM in Compute Engine. Deprecated: Use compute_engine_target_details instead.
        """
        return pulumi.get(self, "compute_engine_vm_details")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time the clone job was created (as an API call, not when it was actually created in the target).
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def error(self) -> pulumi.Output['outputs.StatusResponse']:
        """
        Provides details for the errors that led to the Clone Job's state.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the clone.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the clone job.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateTime")
    def state_time(self) -> pulumi.Output[str]:
        """
        The time the state was last updated.
        """
        return pulumi.get(self, "state_time")

    @property
    @pulumi.getter(name="targetDetails")
    def target_details(self) -> pulumi.Output['outputs.TargetVMDetailsResponse']:
        """
        Details of the VM to create as the target of this clone job. Deprecated: Use compute_engine_target_details instead.
        """
        return pulumi.get(self, "target_details")

