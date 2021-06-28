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
    'GetReservationResult',
    'AwaitableGetReservationResult',
    'get_reservation',
]

@pulumi.output_type
class GetReservationResult:
    def __init__(__self__, commitment=None, creation_timestamp=None, description=None, kind=None, name=None, satisfies_pzs=None, self_link=None, self_link_with_id=None, share_settings=None, specific_reservation=None, specific_reservation_required=None, status=None, zone=None):
        if commitment and not isinstance(commitment, str):
            raise TypeError("Expected argument 'commitment' to be a str")
        pulumi.set(__self__, "commitment", commitment)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id and not isinstance(self_link_with_id, str):
            raise TypeError("Expected argument 'self_link_with_id' to be a str")
        pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if share_settings and not isinstance(share_settings, dict):
            raise TypeError("Expected argument 'share_settings' to be a dict")
        pulumi.set(__self__, "share_settings", share_settings)
        if specific_reservation and not isinstance(specific_reservation, dict):
            raise TypeError("Expected argument 'specific_reservation' to be a dict")
        pulumi.set(__self__, "specific_reservation", specific_reservation)
        if specific_reservation_required and not isinstance(specific_reservation_required, bool):
            raise TypeError("Expected argument 'specific_reservation_required' to be a bool")
        pulumi.set(__self__, "specific_reservation_required", specific_reservation_required)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def commitment(self) -> str:
        """
        Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
        """
        return pulumi.get(self, "commitment")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#reservations for reservations.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined fully-qualified URL for this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> str:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter(name="shareSettings")
    def share_settings(self) -> 'outputs.AllocationShareSettingsResponse':
        """
        Share-settings for shared-reservation
        """
        return pulumi.get(self, "share_settings")

    @property
    @pulumi.getter(name="specificReservation")
    def specific_reservation(self) -> 'outputs.AllocationSpecificSKUReservationResponse':
        """
        Reservation for instances with specific machine shapes.
        """
        return pulumi.get(self, "specific_reservation")

    @property
    @pulumi.getter(name="specificReservationRequired")
    def specific_reservation_required(self) -> bool:
        """
        Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
        """
        return pulumi.get(self, "specific_reservation_required")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the reservation.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
        """
        return pulumi.get(self, "zone")


class AwaitableGetReservationResult(GetReservationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReservationResult(
            commitment=self.commitment,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            kind=self.kind,
            name=self.name,
            satisfies_pzs=self.satisfies_pzs,
            self_link=self.self_link,
            self_link_with_id=self.self_link_with_id,
            share_settings=self.share_settings,
            specific_reservation=self.specific_reservation,
            specific_reservation_required=self.specific_reservation_required,
            status=self.status,
            zone=self.zone)


def get_reservation(project: Optional[str] = None,
                    reservation: Optional[str] = None,
                    zone: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReservationResult:
    """
    Retrieves information about the specified reservation.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['reservation'] = reservation
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/alpha:getReservation', __args__, opts=opts, typ=GetReservationResult).value

    return AwaitableGetReservationResult(
        commitment=__ret__.commitment,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        kind=__ret__.kind,
        name=__ret__.name,
        satisfies_pzs=__ret__.satisfies_pzs,
        self_link=__ret__.self_link,
        self_link_with_id=__ret__.self_link_with_id,
        share_settings=__ret__.share_settings,
        specific_reservation=__ret__.specific_reservation,
        specific_reservation_required=__ret__.specific_reservation_required,
        status=__ret__.status,
        zone=__ret__.zone)
