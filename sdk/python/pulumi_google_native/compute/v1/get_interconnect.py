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
    'GetInterconnectResult',
    'AwaitableGetInterconnectResult',
    'get_interconnect',
    'get_interconnect_output',
]

@pulumi.output_type
class GetInterconnectResult:
    def __init__(__self__, admin_enabled=None, circuit_infos=None, creation_timestamp=None, customer_name=None, description=None, expected_outages=None, google_ip_address=None, google_reference_id=None, interconnect_attachments=None, interconnect_type=None, kind=None, link_type=None, location=None, name=None, noc_contact_email=None, operational_status=None, peer_ip_address=None, provisioned_link_count=None, requested_link_count=None, satisfies_pzs=None, self_link=None, state=None):
        if admin_enabled and not isinstance(admin_enabled, bool):
            raise TypeError("Expected argument 'admin_enabled' to be a bool")
        pulumi.set(__self__, "admin_enabled", admin_enabled)
        if circuit_infos and not isinstance(circuit_infos, list):
            raise TypeError("Expected argument 'circuit_infos' to be a list")
        pulumi.set(__self__, "circuit_infos", circuit_infos)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if customer_name and not isinstance(customer_name, str):
            raise TypeError("Expected argument 'customer_name' to be a str")
        pulumi.set(__self__, "customer_name", customer_name)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if expected_outages and not isinstance(expected_outages, list):
            raise TypeError("Expected argument 'expected_outages' to be a list")
        pulumi.set(__self__, "expected_outages", expected_outages)
        if google_ip_address and not isinstance(google_ip_address, str):
            raise TypeError("Expected argument 'google_ip_address' to be a str")
        pulumi.set(__self__, "google_ip_address", google_ip_address)
        if google_reference_id and not isinstance(google_reference_id, str):
            raise TypeError("Expected argument 'google_reference_id' to be a str")
        pulumi.set(__self__, "google_reference_id", google_reference_id)
        if interconnect_attachments and not isinstance(interconnect_attachments, list):
            raise TypeError("Expected argument 'interconnect_attachments' to be a list")
        pulumi.set(__self__, "interconnect_attachments", interconnect_attachments)
        if interconnect_type and not isinstance(interconnect_type, str):
            raise TypeError("Expected argument 'interconnect_type' to be a str")
        pulumi.set(__self__, "interconnect_type", interconnect_type)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if link_type and not isinstance(link_type, str):
            raise TypeError("Expected argument 'link_type' to be a str")
        pulumi.set(__self__, "link_type", link_type)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if noc_contact_email and not isinstance(noc_contact_email, str):
            raise TypeError("Expected argument 'noc_contact_email' to be a str")
        pulumi.set(__self__, "noc_contact_email", noc_contact_email)
        if operational_status and not isinstance(operational_status, str):
            raise TypeError("Expected argument 'operational_status' to be a str")
        pulumi.set(__self__, "operational_status", operational_status)
        if peer_ip_address and not isinstance(peer_ip_address, str):
            raise TypeError("Expected argument 'peer_ip_address' to be a str")
        pulumi.set(__self__, "peer_ip_address", peer_ip_address)
        if provisioned_link_count and not isinstance(provisioned_link_count, int):
            raise TypeError("Expected argument 'provisioned_link_count' to be a int")
        pulumi.set(__self__, "provisioned_link_count", provisioned_link_count)
        if requested_link_count and not isinstance(requested_link_count, int):
            raise TypeError("Expected argument 'requested_link_count' to be a int")
        pulumi.set(__self__, "requested_link_count", requested_link_count)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="adminEnabled")
    def admin_enabled(self) -> bool:
        """
        Administrative status of the interconnect. When this is set to true, the Interconnect is functional and can carry traffic. When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it. By default, the status is set to true.
        """
        return pulumi.get(self, "admin_enabled")

    @property
    @pulumi.getter(name="circuitInfos")
    def circuit_infos(self) -> Sequence['outputs.InterconnectCircuitInfoResponse']:
        """
        A list of CircuitInfo objects, that describe the individual circuits in this LAG.
        """
        return pulumi.get(self, "circuit_infos")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="customerName")
    def customer_name(self) -> str:
        """
        Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect.
        """
        return pulumi.get(self, "customer_name")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expectedOutages")
    def expected_outages(self) -> Sequence['outputs.InterconnectOutageNotificationResponse']:
        """
        A list of outages expected for this Interconnect.
        """
        return pulumi.get(self, "expected_outages")

    @property
    @pulumi.getter(name="googleIpAddress")
    def google_ip_address(self) -> str:
        """
        IP address configured on the Google side of the Interconnect link. This can be used only for ping tests.
        """
        return pulumi.get(self, "google_ip_address")

    @property
    @pulumi.getter(name="googleReferenceId")
    def google_reference_id(self) -> str:
        """
        Google reference ID to be used when raising support tickets with Google or otherwise to debug backend connectivity issues.
        """
        return pulumi.get(self, "google_reference_id")

    @property
    @pulumi.getter(name="interconnectAttachments")
    def interconnect_attachments(self) -> Sequence[str]:
        """
        A list of the URLs of all InterconnectAttachments configured to use this Interconnect.
        """
        return pulumi.get(self, "interconnect_attachments")

    @property
    @pulumi.getter(name="interconnectType")
    def interconnect_type(self) -> str:
        """
        Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner. - DEDICATED: A dedicated physical interconnection with the customer. Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.
        """
        return pulumi.get(self, "interconnect_type")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#interconnect for interconnects.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="linkType")
    def link_type(self) -> str:
        """
        Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics. Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle.
        """
        return pulumi.get(self, "link_type")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        URL of the InterconnectLocation object that represents where this connection is to be provisioned.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nocContactEmail")
    def noc_contact_email(self) -> str:
        """
        Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect. If specified, this will be used for notifications in addition to all other forms described, such as Stackdriver logs alerting and Cloud Notifications.
        """
        return pulumi.get(self, "noc_contact_email")

    @property
    @pulumi.getter(name="operationalStatus")
    def operational_status(self) -> str:
        """
        The current status of this Interconnect's functionality, which can take one of the following values: - OS_ACTIVE: A valid Interconnect, which is turned up and is ready to use. Attachments may be provisioned on this Interconnect. - OS_UNPROVISIONED: An Interconnect that has not completed turnup. No attachments may be provisioned on this Interconnect. - OS_UNDER_MAINTENANCE: An Interconnect that is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect. 
        """
        return pulumi.get(self, "operational_status")

    @property
    @pulumi.getter(name="peerIpAddress")
    def peer_ip_address(self) -> str:
        """
        IP address configured on the customer side of the Interconnect link. The customer should configure this IP address during turnup when prompted by Google NOC. This can be used only for ping tests.
        """
        return pulumi.get(self, "peer_ip_address")

    @property
    @pulumi.getter(name="provisionedLinkCount")
    def provisioned_link_count(self) -> int:
        """
        Number of links actually provisioned in this interconnect.
        """
        return pulumi.get(self, "provisioned_link_count")

    @property
    @pulumi.getter(name="requestedLinkCount")
    def requested_link_count(self) -> int:
        """
        Target number of physical links in the link bundle, as requested by the customer.
        """
        return pulumi.get(self, "requested_link_count")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        Set to true if the resource satisfies the zone separation organization policy constraints and false otherwise. Defaults to false if the field is not present.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of Interconnect functionality, which can take one of the following values: - ACTIVE: The Interconnect is valid, turned up and ready to use. Attachments may be provisioned on this Interconnect. - UNPROVISIONED: The Interconnect has not completed turnup. No attachments may be provisioned on this Interconnect. - UNDER_MAINTENANCE: The Interconnect is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect. 
        """
        return pulumi.get(self, "state")


class AwaitableGetInterconnectResult(GetInterconnectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInterconnectResult(
            admin_enabled=self.admin_enabled,
            circuit_infos=self.circuit_infos,
            creation_timestamp=self.creation_timestamp,
            customer_name=self.customer_name,
            description=self.description,
            expected_outages=self.expected_outages,
            google_ip_address=self.google_ip_address,
            google_reference_id=self.google_reference_id,
            interconnect_attachments=self.interconnect_attachments,
            interconnect_type=self.interconnect_type,
            kind=self.kind,
            link_type=self.link_type,
            location=self.location,
            name=self.name,
            noc_contact_email=self.noc_contact_email,
            operational_status=self.operational_status,
            peer_ip_address=self.peer_ip_address,
            provisioned_link_count=self.provisioned_link_count,
            requested_link_count=self.requested_link_count,
            satisfies_pzs=self.satisfies_pzs,
            self_link=self.self_link,
            state=self.state)


def get_interconnect(interconnect: Optional[str] = None,
                     project: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInterconnectResult:
    """
    Returns the specified interconnect. Get a list of available interconnects by making a list() request.
    """
    __args__ = dict()
    __args__['interconnect'] = interconnect
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getInterconnect', __args__, opts=opts, typ=GetInterconnectResult).value

    return AwaitableGetInterconnectResult(
        admin_enabled=__ret__.admin_enabled,
        circuit_infos=__ret__.circuit_infos,
        creation_timestamp=__ret__.creation_timestamp,
        customer_name=__ret__.customer_name,
        description=__ret__.description,
        expected_outages=__ret__.expected_outages,
        google_ip_address=__ret__.google_ip_address,
        google_reference_id=__ret__.google_reference_id,
        interconnect_attachments=__ret__.interconnect_attachments,
        interconnect_type=__ret__.interconnect_type,
        kind=__ret__.kind,
        link_type=__ret__.link_type,
        location=__ret__.location,
        name=__ret__.name,
        noc_contact_email=__ret__.noc_contact_email,
        operational_status=__ret__.operational_status,
        peer_ip_address=__ret__.peer_ip_address,
        provisioned_link_count=__ret__.provisioned_link_count,
        requested_link_count=__ret__.requested_link_count,
        satisfies_pzs=__ret__.satisfies_pzs,
        self_link=__ret__.self_link,
        state=__ret__.state)


@_utilities.lift_output_func(get_interconnect)
def get_interconnect_output(interconnect: Optional[pulumi.Input[str]] = None,
                            project: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInterconnectResult]:
    """
    Returns the specified interconnect. Get a list of available interconnects by making a list() request.
    """
    ...
