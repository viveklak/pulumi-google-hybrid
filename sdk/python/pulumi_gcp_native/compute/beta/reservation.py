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

__all__ = ['ReservationArgs', 'Reservation']

@pulumi.input_type
class ReservationArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 reservation: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 commitment: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 satisfies_pzs: Optional[pulumi.Input[bool]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 specific_reservation: Optional[pulumi.Input['AllocationSpecificSKUReservationArgs']] = None,
                 specific_reservation_required: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Reservation resource.
        :param pulumi.Input[str] zone: Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
        :param pulumi.Input[str] commitment: [Output Only] Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output Only] Type of the resource. Always compute#reservations for reservations.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[bool] satisfies_pzs: [Output Only] Reserved for future use.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined fully-qualified URL for this resource.
        :param pulumi.Input['AllocationSpecificSKUReservationArgs'] specific_reservation: Reservation for instances with specific machine shapes.
        :param pulumi.Input[bool] specific_reservation_required: Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
        :param pulumi.Input[str] status: [Output Only] The status of the reservation.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "reservation", reservation)
        pulumi.set(__self__, "zone", zone)
        if commitment is not None:
            pulumi.set(__self__, "commitment", commitment)
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
        if satisfies_pzs is not None:
            pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if specific_reservation is not None:
            pulumi.set(__self__, "specific_reservation", specific_reservation)
        if specific_reservation_required is not None:
            pulumi.set(__self__, "specific_reservation_required", specific_reservation_required)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def reservation(self) -> pulumi.Input[str]:
        return pulumi.get(self, "reservation")

    @reservation.setter
    def reservation(self, value: pulumi.Input[str]):
        pulumi.set(self, "reservation", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter
    def commitment(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
        """
        return pulumi.get(self, "commitment")

    @commitment.setter
    def commitment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commitment", value)

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
        [Output Only] Type of the resource. Always compute#reservations for reservations.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> Optional[pulumi.Input[bool]]:
        """
        [Output Only] Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @satisfies_pzs.setter
    def satisfies_pzs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "satisfies_pzs", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Server-defined fully-qualified URL for this resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="specificReservation")
    def specific_reservation(self) -> Optional[pulumi.Input['AllocationSpecificSKUReservationArgs']]:
        """
        Reservation for instances with specific machine shapes.
        """
        return pulumi.get(self, "specific_reservation")

    @specific_reservation.setter
    def specific_reservation(self, value: Optional[pulumi.Input['AllocationSpecificSKUReservationArgs']]):
        pulumi.set(self, "specific_reservation", value)

    @property
    @pulumi.getter(name="specificReservationRequired")
    def specific_reservation_required(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
        """
        return pulumi.get(self, "specific_reservation_required")

    @specific_reservation_required.setter
    def specific_reservation_required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "specific_reservation_required", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The status of the reservation.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Reservation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 commitment: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reservation: Optional[pulumi.Input[str]] = None,
                 satisfies_pzs: Optional[pulumi.Input[bool]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 specific_reservation: Optional[pulumi.Input[pulumi.InputType['AllocationSpecificSKUReservationArgs']]] = None,
                 specific_reservation_required: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new reservation. For more information, read Reserving zonal resources.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] commitment: [Output Only] Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output Only] Type of the resource. Always compute#reservations for reservations.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[bool] satisfies_pzs: [Output Only] Reserved for future use.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined fully-qualified URL for this resource.
        :param pulumi.Input[pulumi.InputType['AllocationSpecificSKUReservationArgs']] specific_reservation: Reservation for instances with specific machine shapes.
        :param pulumi.Input[bool] specific_reservation_required: Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
        :param pulumi.Input[str] status: [Output Only] The status of the reservation.
        :param pulumi.Input[str] zone: Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReservationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new reservation. For more information, read Reserving zonal resources.

        :param str resource_name: The name of the resource.
        :param ReservationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReservationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 commitment: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reservation: Optional[pulumi.Input[str]] = None,
                 satisfies_pzs: Optional[pulumi.Input[bool]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 specific_reservation: Optional[pulumi.Input[pulumi.InputType['AllocationSpecificSKUReservationArgs']]] = None,
                 specific_reservation_required: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None,
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
            __props__ = ReservationArgs.__new__(ReservationArgs)

            __props__.__dict__["commitment"] = commitment
            __props__.__dict__["creation_timestamp"] = creation_timestamp
            __props__.__dict__["description"] = description
            __props__.__dict__["id"] = id
            __props__.__dict__["kind"] = kind
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if reservation is None and not opts.urn:
                raise TypeError("Missing required property 'reservation'")
            __props__.__dict__["reservation"] = reservation
            __props__.__dict__["satisfies_pzs"] = satisfies_pzs
            __props__.__dict__["self_link"] = self_link
            __props__.__dict__["specific_reservation"] = specific_reservation
            __props__.__dict__["specific_reservation_required"] = specific_reservation_required
            __props__.__dict__["status"] = status
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
        super(Reservation, __self__).__init__(
            'gcp-native:compute/beta:Reservation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Reservation':
        """
        Get an existing Reservation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ReservationArgs.__new__(ReservationArgs)

        __props__.__dict__["commitment"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["satisfies_pzs"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["specific_reservation"] = None
        __props__.__dict__["specific_reservation_required"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["zone"] = None
        return Reservation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def commitment(self) -> pulumi.Output[str]:
        """
        [Output Only] Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
        """
        return pulumi.get(self, "commitment")

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
        [Output Only] Type of the resource. Always compute#reservations for reservations.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> pulumi.Output[bool]:
        """
        [Output Only] Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        [Output Only] Server-defined fully-qualified URL for this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="specificReservation")
    def specific_reservation(self) -> pulumi.Output['outputs.AllocationSpecificSKUReservationResponse']:
        """
        Reservation for instances with specific machine shapes.
        """
        return pulumi.get(self, "specific_reservation")

    @property
    @pulumi.getter(name="specificReservationRequired")
    def specific_reservation_required(self) -> pulumi.Output[bool]:
        """
        Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
        """
        return pulumi.get(self, "specific_reservation_required")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        [Output Only] The status of the reservation.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
        """
        return pulumi.get(self, "zone")

