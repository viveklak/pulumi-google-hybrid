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
    'GetRouteResult',
    'AwaitableGetRouteResult',
    'get_route',
]

@pulumi.output_type
class GetRouteResult:
    def __init__(__self__, creation_timestamp=None, description=None, dest_range=None, kind=None, name=None, network=None, next_hop_gateway=None, next_hop_ilb=None, next_hop_instance=None, next_hop_interconnect_attachment=None, next_hop_ip=None, next_hop_network=None, next_hop_peering=None, next_hop_vpn_tunnel=None, priority=None, self_link=None, tags=None, warnings=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dest_range and not isinstance(dest_range, str):
            raise TypeError("Expected argument 'dest_range' to be a str")
        pulumi.set(__self__, "dest_range", dest_range)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if next_hop_gateway and not isinstance(next_hop_gateway, str):
            raise TypeError("Expected argument 'next_hop_gateway' to be a str")
        pulumi.set(__self__, "next_hop_gateway", next_hop_gateway)
        if next_hop_ilb and not isinstance(next_hop_ilb, str):
            raise TypeError("Expected argument 'next_hop_ilb' to be a str")
        pulumi.set(__self__, "next_hop_ilb", next_hop_ilb)
        if next_hop_instance and not isinstance(next_hop_instance, str):
            raise TypeError("Expected argument 'next_hop_instance' to be a str")
        pulumi.set(__self__, "next_hop_instance", next_hop_instance)
        if next_hop_interconnect_attachment and not isinstance(next_hop_interconnect_attachment, str):
            raise TypeError("Expected argument 'next_hop_interconnect_attachment' to be a str")
        pulumi.set(__self__, "next_hop_interconnect_attachment", next_hop_interconnect_attachment)
        if next_hop_ip and not isinstance(next_hop_ip, str):
            raise TypeError("Expected argument 'next_hop_ip' to be a str")
        pulumi.set(__self__, "next_hop_ip", next_hop_ip)
        if next_hop_network and not isinstance(next_hop_network, str):
            raise TypeError("Expected argument 'next_hop_network' to be a str")
        pulumi.set(__self__, "next_hop_network", next_hop_network)
        if next_hop_peering and not isinstance(next_hop_peering, str):
            raise TypeError("Expected argument 'next_hop_peering' to be a str")
        pulumi.set(__self__, "next_hop_peering", next_hop_peering)
        if next_hop_vpn_tunnel and not isinstance(next_hop_vpn_tunnel, str):
            raise TypeError("Expected argument 'next_hop_vpn_tunnel' to be a str")
        pulumi.set(__self__, "next_hop_vpn_tunnel", next_hop_vpn_tunnel)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if warnings and not isinstance(warnings, list):
            raise TypeError("Expected argument 'warnings' to be a list")
        pulumi.set(__self__, "warnings", warnings)

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
        An optional description of this resource. Provide this field when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destRange")
    def dest_range(self) -> str:
        """
        The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
        """
        return pulumi.get(self, "dest_range")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of this resource. Always compute#routes for Route resources.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        Fully-qualified URL of the network that this route applies to.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="nextHopGateway")
    def next_hop_gateway(self) -> str:
        """
        The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL:  projects/project/global/gateways/default-internet-gateway
        """
        return pulumi.get(self, "next_hop_gateway")

    @property
    @pulumi.getter(name="nextHopIlb")
    def next_hop_ilb(self) -> str:
        """
        The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs:  
        - 10.128.0.56 
        - https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule 
        - regions/region/forwardingRules/forwardingRule
        """
        return pulumi.get(self, "next_hop_ilb")

    @property
    @pulumi.getter(name="nextHopInstance")
    def next_hop_instance(self) -> str:
        """
        The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example:
        https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
        """
        return pulumi.get(self, "next_hop_instance")

    @property
    @pulumi.getter(name="nextHopInterconnectAttachment")
    def next_hop_interconnect_attachment(self) -> str:
        """
        The URL to an InterconnectAttachment which is the next hop for the route. This field will only be populated for the dynamic routes generated by Cloud Router with a linked interconnectAttachment.
        """
        return pulumi.get(self, "next_hop_interconnect_attachment")

    @property
    @pulumi.getter(name="nextHopIp")
    def next_hop_ip(self) -> str:
        """
        The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
        """
        return pulumi.get(self, "next_hop_ip")

    @property
    @pulumi.getter(name="nextHopNetwork")
    def next_hop_network(self) -> str:
        """
        The URL of the local network if it should handle matching packets.
        """
        return pulumi.get(self, "next_hop_network")

    @property
    @pulumi.getter(name="nextHopPeering")
    def next_hop_peering(self) -> str:
        """
        The network peering name that should handle matching packets, which should conform to RFC1035.
        """
        return pulumi.get(self, "next_hop_peering")

    @property
    @pulumi.getter(name="nextHopVpnTunnel")
    def next_hop_vpn_tunnel(self) -> str:
        """
        The URL to a VpnTunnel that should handle matching packets.
        """
        return pulumi.get(self, "next_hop_vpn_tunnel")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined fully-qualified URL for this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        A list of instance tags to which this route applies.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def warnings(self) -> Sequence['outputs.RouteWarningsItemResponse']:
        """
        If potential misconfigurations are detected for this route, this field will be populated with warning messages.
        """
        return pulumi.get(self, "warnings")


class AwaitableGetRouteResult(GetRouteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteResult(
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            dest_range=self.dest_range,
            kind=self.kind,
            name=self.name,
            network=self.network,
            next_hop_gateway=self.next_hop_gateway,
            next_hop_ilb=self.next_hop_ilb,
            next_hop_instance=self.next_hop_instance,
            next_hop_interconnect_attachment=self.next_hop_interconnect_attachment,
            next_hop_ip=self.next_hop_ip,
            next_hop_network=self.next_hop_network,
            next_hop_peering=self.next_hop_peering,
            next_hop_vpn_tunnel=self.next_hop_vpn_tunnel,
            priority=self.priority,
            self_link=self.self_link,
            tags=self.tags,
            warnings=self.warnings)


def get_route(project: Optional[str] = None,
              route: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteResult:
    """
    Returns the specified Route resource. Gets a list of available routes by making a list() request.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['route'] = route
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getRoute', __args__, opts=opts, typ=GetRouteResult).value

    return AwaitableGetRouteResult(
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        dest_range=__ret__.dest_range,
        kind=__ret__.kind,
        name=__ret__.name,
        network=__ret__.network,
        next_hop_gateway=__ret__.next_hop_gateway,
        next_hop_ilb=__ret__.next_hop_ilb,
        next_hop_instance=__ret__.next_hop_instance,
        next_hop_interconnect_attachment=__ret__.next_hop_interconnect_attachment,
        next_hop_ip=__ret__.next_hop_ip,
        next_hop_network=__ret__.next_hop_network,
        next_hop_peering=__ret__.next_hop_peering,
        next_hop_vpn_tunnel=__ret__.next_hop_vpn_tunnel,
        priority=__ret__.priority,
        self_link=__ret__.self_link,
        tags=__ret__.tags,
        warnings=__ret__.warnings)
