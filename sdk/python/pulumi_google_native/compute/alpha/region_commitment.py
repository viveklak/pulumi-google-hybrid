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

__all__ = ['RegionCommitmentArgs', 'RegionCommitment']

@pulumi.input_type
class RegionCommitmentArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 region: pulumi.Input[str],
                 category: Optional[pulumi.Input['RegionCommitmentCategory']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 license_resource: Optional[pulumi.Input['LicenseResourceCommitmentArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['RegionCommitmentPlan']] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 reservations: Optional[pulumi.Input[Sequence[pulumi.Input['ReservationArgs']]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceCommitmentArgs']]]] = None,
                 type: Optional[pulumi.Input['RegionCommitmentType']] = None):
        """
        The set of arguments for constructing a RegionCommitment resource.
        :param pulumi.Input['RegionCommitmentCategory'] category: The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input['LicenseResourceCommitmentArgs'] license_resource: The license specification required as part of a license commitment.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input['RegionCommitmentPlan'] plan: The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        :param pulumi.Input[Sequence[pulumi.Input['ReservationArgs']]] reservations: List of reservations in this commitment.
        :param pulumi.Input[Sequence[pulumi.Input['ResourceCommitmentArgs']]] resources: A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        :param pulumi.Input['RegionCommitmentType'] type: The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "region", region)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if license_resource is not None:
            pulumi.set(__self__, "license_resource", license_resource)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if reservations is not None:
            pulumi.set(__self__, "reservations", reservations)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if type is not None:
            pulumi.set(__self__, "type", type)

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
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input['RegionCommitmentCategory']]:
        """
        The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input['RegionCommitmentCategory']]):
        pulumi.set(self, "category", value)

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
    @pulumi.getter(name="licenseResource")
    def license_resource(self) -> Optional[pulumi.Input['LicenseResourceCommitmentArgs']]:
        """
        The license specification required as part of a license commitment.
        """
        return pulumi.get(self, "license_resource")

    @license_resource.setter
    def license_resource(self, value: Optional[pulumi.Input['LicenseResourceCommitmentArgs']]):
        pulumi.set(self, "license_resource", value)

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
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['RegionCommitmentPlan']]:
        """
        The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['RegionCommitmentPlan']]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def reservations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReservationArgs']]]]:
        """
        List of reservations in this commitment.
        """
        return pulumi.get(self, "reservations")

    @reservations.setter
    def reservations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReservationArgs']]]]):
        pulumi.set(self, "reservations", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResourceCommitmentArgs']]]]:
        """
        A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceCommitmentArgs']]]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['RegionCommitmentType']]:
        """
        The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['RegionCommitmentType']]):
        pulumi.set(self, "type", value)


class RegionCommitment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category: Optional[pulumi.Input['RegionCommitmentCategory']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 license_resource: Optional[pulumi.Input[pulumi.InputType['LicenseResourceCommitmentArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['RegionCommitmentPlan']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 reservations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReservationArgs']]]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceCommitmentArgs']]]]] = None,
                 type: Optional[pulumi.Input['RegionCommitmentType']] = None,
                 __props__=None):
        """
        Creates a commitment in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['RegionCommitmentCategory'] category: The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[pulumi.InputType['LicenseResourceCommitmentArgs']] license_resource: The license specification required as part of a license commitment.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input['RegionCommitmentPlan'] plan: The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReservationArgs']]]] reservations: List of reservations in this commitment.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceCommitmentArgs']]]] resources: A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        :param pulumi.Input['RegionCommitmentType'] type: The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegionCommitmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a commitment in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param RegionCommitmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegionCommitmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category: Optional[pulumi.Input['RegionCommitmentCategory']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 license_resource: Optional[pulumi.Input[pulumi.InputType['LicenseResourceCommitmentArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['RegionCommitmentPlan']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 reservations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReservationArgs']]]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceCommitmentArgs']]]]] = None,
                 type: Optional[pulumi.Input['RegionCommitmentType']] = None,
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
            __props__ = RegionCommitmentArgs.__new__(RegionCommitmentArgs)

            __props__.__dict__["category"] = category
            __props__.__dict__["description"] = description
            __props__.__dict__["license_resource"] = license_resource
            __props__.__dict__["name"] = name
            __props__.__dict__["plan"] = plan
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["reservations"] = reservations
            __props__.__dict__["resources"] = resources
            __props__.__dict__["type"] = type
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["end_timestamp"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["self_link_with_id"] = None
            __props__.__dict__["start_timestamp"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["status_message"] = None
        super(RegionCommitment, __self__).__init__(
            'google-native:compute/alpha:RegionCommitment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RegionCommitment':
        """
        Get an existing RegionCommitment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RegionCommitmentArgs.__new__(RegionCommitmentArgs)

        __props__.__dict__["category"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["end_timestamp"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["license_resource"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["plan"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["reservations"] = None
        __props__.__dict__["resources"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["self_link_with_id"] = None
        __props__.__dict__["start_timestamp"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["status_message"] = None
        __props__.__dict__["type"] = None
        return RegionCommitment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        """
        The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
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
    @pulumi.getter(name="endTimestamp")
    def end_timestamp(self) -> pulumi.Output[str]:
        """
        Commitment end time in RFC3339 text format.
        """
        return pulumi.get(self, "end_timestamp")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Type of the resource. Always compute#commitment for commitments.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="licenseResource")
    def license_resource(self) -> pulumi.Output['outputs.LicenseResourceCommitmentResponse']:
        """
        The license specification required as part of a license commitment.
        """
        return pulumi.get(self, "license_resource")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[str]:
        """
        The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        URL of the region where this commitment may be used.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def reservations(self) -> pulumi.Output[Sequence['outputs.ReservationResponse']]:
        """
        List of reservations in this commitment.
        """
        return pulumi.get(self, "reservations")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[Sequence['outputs.ResourceCommitmentResponse']]:
        """
        A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> pulumi.Output[str]:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter(name="startTimestamp")
    def start_timestamp(self) -> pulumi.Output[str]:
        """
        Commitment start time in RFC3339 text format.
        """
        return pulumi.get(self, "start_timestamp")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> pulumi.Output[str]:
        """
        An optional, human-readable explanation of the status.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        """
        return pulumi.get(self, "type")

