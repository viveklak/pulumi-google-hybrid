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
    'GetServiceAttachmentResult',
    'AwaitableGetServiceAttachmentResult',
    'get_service_attachment',
    'get_service_attachment_output',
]

@pulumi.output_type
class GetServiceAttachmentResult:
    def __init__(__self__, connected_endpoints=None, connection_preference=None, consumer_accept_lists=None, consumer_reject_lists=None, creation_timestamp=None, description=None, domain_names=None, enable_proxy_protocol=None, fingerprint=None, kind=None, name=None, nat_subnets=None, producer_forwarding_rule=None, psc_service_attachment_id=None, region=None, self_link=None, target_service=None):
        if connected_endpoints and not isinstance(connected_endpoints, list):
            raise TypeError("Expected argument 'connected_endpoints' to be a list")
        pulumi.set(__self__, "connected_endpoints", connected_endpoints)
        if connection_preference and not isinstance(connection_preference, str):
            raise TypeError("Expected argument 'connection_preference' to be a str")
        pulumi.set(__self__, "connection_preference", connection_preference)
        if consumer_accept_lists and not isinstance(consumer_accept_lists, list):
            raise TypeError("Expected argument 'consumer_accept_lists' to be a list")
        pulumi.set(__self__, "consumer_accept_lists", consumer_accept_lists)
        if consumer_reject_lists and not isinstance(consumer_reject_lists, list):
            raise TypeError("Expected argument 'consumer_reject_lists' to be a list")
        pulumi.set(__self__, "consumer_reject_lists", consumer_reject_lists)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if domain_names and not isinstance(domain_names, list):
            raise TypeError("Expected argument 'domain_names' to be a list")
        pulumi.set(__self__, "domain_names", domain_names)
        if enable_proxy_protocol and not isinstance(enable_proxy_protocol, bool):
            raise TypeError("Expected argument 'enable_proxy_protocol' to be a bool")
        pulumi.set(__self__, "enable_proxy_protocol", enable_proxy_protocol)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nat_subnets and not isinstance(nat_subnets, list):
            raise TypeError("Expected argument 'nat_subnets' to be a list")
        pulumi.set(__self__, "nat_subnets", nat_subnets)
        if producer_forwarding_rule and not isinstance(producer_forwarding_rule, str):
            raise TypeError("Expected argument 'producer_forwarding_rule' to be a str")
        pulumi.set(__self__, "producer_forwarding_rule", producer_forwarding_rule)
        if psc_service_attachment_id and not isinstance(psc_service_attachment_id, dict):
            raise TypeError("Expected argument 'psc_service_attachment_id' to be a dict")
        pulumi.set(__self__, "psc_service_attachment_id", psc_service_attachment_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if target_service and not isinstance(target_service, str):
            raise TypeError("Expected argument 'target_service' to be a str")
        pulumi.set(__self__, "target_service", target_service)

    @property
    @pulumi.getter(name="connectedEndpoints")
    def connected_endpoints(self) -> Sequence['outputs.ServiceAttachmentConnectedEndpointResponse']:
        """
        An array of connections for all the consumers connected to this service attachment.
        """
        return pulumi.get(self, "connected_endpoints")

    @property
    @pulumi.getter(name="connectionPreference")
    def connection_preference(self) -> str:
        """
        The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
        """
        return pulumi.get(self, "connection_preference")

    @property
    @pulumi.getter(name="consumerAcceptLists")
    def consumer_accept_lists(self) -> Sequence['outputs.ServiceAttachmentConsumerProjectLimitResponse']:
        """
        Projects that are allowed to connect to this service attachment.
        """
        return pulumi.get(self, "consumer_accept_lists")

    @property
    @pulumi.getter(name="consumerRejectLists")
    def consumer_reject_lists(self) -> Sequence[str]:
        """
        Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
        """
        return pulumi.get(self, "consumer_reject_lists")

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
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Sequence[str]:
        """
        If specified, the domain name will be used during the integration between the PSC connected endpoints and the Cloud DNS. For example, this is a valid domain name: "p.mycompany.com.". Current max number of domain names supported is 1.
        """
        return pulumi.get(self, "domain_names")

    @property
    @pulumi.getter(name="enableProxyProtocol")
    def enable_proxy_protocol(self) -> bool:
        """
        If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
        """
        return pulumi.get(self, "enable_proxy_protocol")

    @property
    @pulumi.getter
    def fingerprint(self) -> str:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ServiceAttachment. An up-to-date fingerprint must be provided in order to patch/update the ServiceAttachment; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the ServiceAttachment.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#serviceAttachment for service attachments.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="natSubnets")
    def nat_subnets(self) -> Sequence[str]:
        """
        An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
        """
        return pulumi.get(self, "nat_subnets")

    @property
    @pulumi.getter(name="producerForwardingRule")
    def producer_forwarding_rule(self) -> str:
        """
        The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
        """
        return pulumi.get(self, "producer_forwarding_rule")

    @property
    @pulumi.getter(name="pscServiceAttachmentId")
    def psc_service_attachment_id(self) -> 'outputs.Uint128Response':
        """
        An 128-bit global unique ID of the PSC service attachment.
        """
        return pulumi.get(self, "psc_service_attachment_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the service attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="targetService")
    def target_service(self) -> str:
        """
        The URL of a service serving the endpoint identified by this service attachment.
        """
        return pulumi.get(self, "target_service")


class AwaitableGetServiceAttachmentResult(GetServiceAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceAttachmentResult(
            connected_endpoints=self.connected_endpoints,
            connection_preference=self.connection_preference,
            consumer_accept_lists=self.consumer_accept_lists,
            consumer_reject_lists=self.consumer_reject_lists,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            domain_names=self.domain_names,
            enable_proxy_protocol=self.enable_proxy_protocol,
            fingerprint=self.fingerprint,
            kind=self.kind,
            name=self.name,
            nat_subnets=self.nat_subnets,
            producer_forwarding_rule=self.producer_forwarding_rule,
            psc_service_attachment_id=self.psc_service_attachment_id,
            region=self.region,
            self_link=self.self_link,
            target_service=self.target_service)


def get_service_attachment(project: Optional[str] = None,
                           region: Optional[str] = None,
                           service_attachment: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceAttachmentResult:
    """
    Returns the specified ServiceAttachment resource in the given scope.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['region'] = region
    __args__['serviceAttachment'] = service_attachment
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getServiceAttachment', __args__, opts=opts, typ=GetServiceAttachmentResult).value

    return AwaitableGetServiceAttachmentResult(
        connected_endpoints=__ret__.connected_endpoints,
        connection_preference=__ret__.connection_preference,
        consumer_accept_lists=__ret__.consumer_accept_lists,
        consumer_reject_lists=__ret__.consumer_reject_lists,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        domain_names=__ret__.domain_names,
        enable_proxy_protocol=__ret__.enable_proxy_protocol,
        fingerprint=__ret__.fingerprint,
        kind=__ret__.kind,
        name=__ret__.name,
        nat_subnets=__ret__.nat_subnets,
        producer_forwarding_rule=__ret__.producer_forwarding_rule,
        psc_service_attachment_id=__ret__.psc_service_attachment_id,
        region=__ret__.region,
        self_link=__ret__.self_link,
        target_service=__ret__.target_service)


@_utilities.lift_output_func(get_service_attachment)
def get_service_attachment_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                                  region: Optional[pulumi.Input[str]] = None,
                                  service_attachment: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceAttachmentResult]:
    """
    Returns the specified ServiceAttachment resource in the given scope.
    """
    ...
