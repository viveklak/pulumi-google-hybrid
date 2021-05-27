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

__all__ = ['RegionAutoscalerArgs', 'RegionAutoscaler']

@pulumi.input_type
class RegionAutoscalerArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 region: pulumi.Input[str],
                 autoscaling_policy: Optional[pulumi.Input['AutoscalingPolicyArgs']] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recommended_size: Optional[pulumi.Input[int]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 scaling_schedule_status: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 status_details: Optional[pulumi.Input[Sequence[pulumi.Input['AutoscalerStatusDetailsArgs']]]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RegionAutoscaler resource.
        :param pulumi.Input[str] region: [Output Only] URL of the region where the instance group resides (for autoscalers living in regional scope).
        :param pulumi.Input['AutoscalingPolicyArgs'] autoscaling_policy: The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
               
               If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output Only] Type of the resource. Always compute#autoscaler for autoscalers.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[int] recommended_size: [Output Only] Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] scaling_schedule_status: [Output Only] Status information of existing scaling schedules.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        :param pulumi.Input[str] self_link_with_id: [Output Only] Server-defined URL for this resource with the resource id.
        :param pulumi.Input[str] status: [Output Only] The status of the autoscaler configuration. Current set of possible values:  
               - PENDING: Autoscaler backend hasn't read new/updated configuration. 
               - DELETING: Configuration is being deleted. 
               - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field. 
               - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field.  New values might be added in the future.
        :param pulumi.Input[Sequence[pulumi.Input['AutoscalerStatusDetailsArgs']]] status_details: [Output Only] Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
        :param pulumi.Input[str] target: URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
        :param pulumi.Input[str] zone: [Output Only] URL of the zone where the instance group resides (for autoscalers living in zonal scope).
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "region", region)
        if autoscaling_policy is not None:
            pulumi.set(__self__, "autoscaling_policy", autoscaling_policy)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if recommended_size is not None:
            pulumi.set(__self__, "recommended_size", recommended_size)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if scaling_schedule_status is not None:
            pulumi.set(__self__, "scaling_schedule_status", scaling_schedule_status)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id is not None:
            pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if status_details is not None:
            pulumi.set(__self__, "status_details", status_details)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        [Output Only] URL of the region where the instance group resides (for autoscalers living in regional scope).
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> Optional[pulumi.Input['AutoscalingPolicyArgs']]:
        """
        The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.

        If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        """
        return pulumi.get(self, "autoscaling_policy")

    @autoscaling_policy.setter
    def autoscaling_policy(self, value: Optional[pulumi.Input['AutoscalingPolicyArgs']]):
        pulumi.set(self, "autoscaling_policy", value)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Type of the resource. Always compute#autoscaler for autoscalers.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="recommendedSize")
    def recommended_size(self) -> Optional[pulumi.Input[int]]:
        """
        [Output Only] Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
        """
        return pulumi.get(self, "recommended_size")

    @recommended_size.setter
    def recommended_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "recommended_size", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="scalingScheduleStatus")
    def scaling_schedule_status(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        [Output Only] Status information of existing scaling schedules.
        """
        return pulumi.get(self, "scaling_schedule_status")

    @scaling_schedule_status.setter
    def scaling_schedule_status(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "scaling_schedule_status", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @self_link_with_id.setter
    def self_link_with_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link_with_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The status of the autoscaler configuration. Current set of possible values:  
        - PENDING: Autoscaler backend hasn't read new/updated configuration. 
        - DELETING: Configuration is being deleted. 
        - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field. 
        - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field.  New values might be added in the future.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="statusDetails")
    def status_details(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutoscalerStatusDetailsArgs']]]]:
        """
        [Output Only] Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
        """
        return pulumi.get(self, "status_details")

    @status_details.setter
    def status_details(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutoscalerStatusDetailsArgs']]]]):
        pulumi.set(self, "status_details", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] URL of the zone where the instance group resides (for autoscalers living in zonal scope).
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class RegionAutoscaler(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['AutoscalingPolicyArgs']]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recommended_size: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 scaling_schedule_status: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 status_details: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoscalerStatusDetailsArgs']]]]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an autoscaler in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AutoscalingPolicyArgs']] autoscaling_policy: The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
               
               If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output Only] Type of the resource. Always compute#autoscaler for autoscalers.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[int] recommended_size: [Output Only] Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
        :param pulumi.Input[str] region: [Output Only] URL of the region where the instance group resides (for autoscalers living in regional scope).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] scaling_schedule_status: [Output Only] Status information of existing scaling schedules.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        :param pulumi.Input[str] self_link_with_id: [Output Only] Server-defined URL for this resource with the resource id.
        :param pulumi.Input[str] status: [Output Only] The status of the autoscaler configuration. Current set of possible values:  
               - PENDING: Autoscaler backend hasn't read new/updated configuration. 
               - DELETING: Configuration is being deleted. 
               - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field. 
               - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field.  New values might be added in the future.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoscalerStatusDetailsArgs']]]] status_details: [Output Only] Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
        :param pulumi.Input[str] target: URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
        :param pulumi.Input[str] zone: [Output Only] URL of the zone where the instance group resides (for autoscalers living in zonal scope).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegionAutoscalerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an autoscaler in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param RegionAutoscalerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegionAutoscalerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['AutoscalingPolicyArgs']]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recommended_size: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 scaling_schedule_status: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 status_details: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoscalerStatusDetailsArgs']]]]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
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
            __props__ = RegionAutoscalerArgs.__new__(RegionAutoscalerArgs)

            __props__.__dict__["autoscaling_policy"] = autoscaling_policy
            __props__.__dict__["creation_timestamp"] = creation_timestamp
            __props__.__dict__["description"] = description
            __props__.__dict__["id"] = id
            __props__.__dict__["kind"] = kind
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["recommended_size"] = recommended_size
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["scaling_schedule_status"] = scaling_schedule_status
            __props__.__dict__["self_link"] = self_link
            __props__.__dict__["self_link_with_id"] = self_link_with_id
            __props__.__dict__["status"] = status
            __props__.__dict__["status_details"] = status_details
            __props__.__dict__["target"] = target
            __props__.__dict__["zone"] = zone
        super(RegionAutoscaler, __self__).__init__(
            'google-native:compute/alpha:RegionAutoscaler',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RegionAutoscaler':
        """
        Get an existing RegionAutoscaler resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RegionAutoscalerArgs.__new__(RegionAutoscalerArgs)

        __props__.__dict__["autoscaling_policy"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["recommended_size"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["scaling_schedule_status"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["self_link_with_id"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["status_details"] = None
        __props__.__dict__["target"] = None
        __props__.__dict__["zone"] = None
        return RegionAutoscaler(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> pulumi.Output['outputs.AutoscalingPolicyResponse']:
        """
        The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.

        If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        """
        return pulumi.get(self, "autoscaling_policy")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        [Output Only] Type of the resource. Always compute#autoscaler for autoscalers.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="recommendedSize")
    def recommended_size(self) -> pulumi.Output[int]:
        """
        [Output Only] Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
        """
        return pulumi.get(self, "recommended_size")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        [Output Only] URL of the region where the instance group resides (for autoscalers living in regional scope).
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="scalingScheduleStatus")
    def scaling_schedule_status(self) -> pulumi.Output[Mapping[str, str]]:
        """
        [Output Only] Status information of existing scaling schedules.
        """
        return pulumi.get(self, "scaling_schedule_status")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> pulumi.Output[str]:
        """
        [Output Only] Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        [Output Only] The status of the autoscaler configuration. Current set of possible values:  
        - PENDING: Autoscaler backend hasn't read new/updated configuration. 
        - DELETING: Configuration is being deleted. 
        - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field. 
        - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field.  New values might be added in the future.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusDetails")
    def status_details(self) -> pulumi.Output[Sequence['outputs.AutoscalerStatusDetailsResponse']]:
        """
        [Output Only] Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
        """
        return pulumi.get(self, "status_details")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[str]:
        """
        URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        [Output Only] URL of the zone where the instance group resides (for autoscalers living in zonal scope).
        """
        return pulumi.get(self, "zone")

