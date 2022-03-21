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
    'GetPacketMirroringResult',
    'AwaitableGetPacketMirroringResult',
    'get_packet_mirroring',
    'get_packet_mirroring_output',
]

@pulumi.output_type
class GetPacketMirroringResult:
    def __init__(__self__, collector_ilb=None, creation_timestamp=None, description=None, enable=None, filter=None, kind=None, mirrored_resources=None, name=None, network=None, priority=None, region=None, self_link=None, self_link_with_id=None):
        if collector_ilb and not isinstance(collector_ilb, dict):
            raise TypeError("Expected argument 'collector_ilb' to be a dict")
        pulumi.set(__self__, "collector_ilb", collector_ilb)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enable and not isinstance(enable, str):
            raise TypeError("Expected argument 'enable' to be a str")
        pulumi.set(__self__, "enable", enable)
        if filter and not isinstance(filter, dict):
            raise TypeError("Expected argument 'filter' to be a dict")
        pulumi.set(__self__, "filter", filter)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if mirrored_resources and not isinstance(mirrored_resources, dict):
            raise TypeError("Expected argument 'mirrored_resources' to be a dict")
        pulumi.set(__self__, "mirrored_resources", mirrored_resources)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, dict):
            raise TypeError("Expected argument 'network' to be a dict")
        pulumi.set(__self__, "network", network)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id and not isinstance(self_link_with_id, str):
            raise TypeError("Expected argument 'self_link_with_id' to be a str")
        pulumi.set(__self__, "self_link_with_id", self_link_with_id)

    @property
    @pulumi.getter(name="collectorIlb")
    def collector_ilb(self) -> 'outputs.PacketMirroringForwardingRuleInfoResponse':
        """
        The Forwarding Rule resource of type loadBalancingScheme=INTERNAL that will be used as collector for mirrored traffic. The specified forwarding rule must have isMirroringCollector set to true.
        """
        return pulumi.get(self, "collector_ilb")

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
    def enable(self) -> str:
        """
        Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network. The default is TRUE.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter
    def filter(self) -> 'outputs.PacketMirroringFilterResponse':
        """
        Filter for mirrored traffic. If unspecified, all traffic is mirrored.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#packetMirroring for packet mirrorings.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="mirroredResources")
    def mirrored_resources(self) -> 'outputs.PacketMirroringMirroredResourceInfoResponse':
        """
        PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
        """
        return pulumi.get(self, "mirrored_resources")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> 'outputs.PacketMirroringNetworkInfoResponse':
        """
        Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URI of the region where the packetMirroring resides.
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
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> str:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")


class AwaitableGetPacketMirroringResult(GetPacketMirroringResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPacketMirroringResult(
            collector_ilb=self.collector_ilb,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            enable=self.enable,
            filter=self.filter,
            kind=self.kind,
            mirrored_resources=self.mirrored_resources,
            name=self.name,
            network=self.network,
            priority=self.priority,
            region=self.region,
            self_link=self.self_link,
            self_link_with_id=self.self_link_with_id)


def get_packet_mirroring(packet_mirroring: Optional[str] = None,
                         project: Optional[str] = None,
                         region: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPacketMirroringResult:
    """
    Returns the specified PacketMirroring resource.
    """
    __args__ = dict()
    __args__['packetMirroring'] = packet_mirroring
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/alpha:getPacketMirroring', __args__, opts=opts, typ=GetPacketMirroringResult).value

    return AwaitableGetPacketMirroringResult(
        collector_ilb=__ret__.collector_ilb,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        enable=__ret__.enable,
        filter=__ret__.filter,
        kind=__ret__.kind,
        mirrored_resources=__ret__.mirrored_resources,
        name=__ret__.name,
        network=__ret__.network,
        priority=__ret__.priority,
        region=__ret__.region,
        self_link=__ret__.self_link,
        self_link_with_id=__ret__.self_link_with_id)


@_utilities.lift_output_func(get_packet_mirroring)
def get_packet_mirroring_output(packet_mirroring: Optional[pulumi.Input[str]] = None,
                                project: Optional[pulumi.Input[Optional[str]]] = None,
                                region: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPacketMirroringResult]:
    """
    Returns the specified PacketMirroring resource.
    """
    ...
