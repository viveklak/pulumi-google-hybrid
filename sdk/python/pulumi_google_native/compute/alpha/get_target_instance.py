# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetTargetInstanceResult',
    'AwaitableGetTargetInstanceResult',
    'get_target_instance',
    'get_target_instance_output',
]

@pulumi.output_type
class GetTargetInstanceResult:
    def __init__(__self__, creation_timestamp=None, description=None, instance=None, kind=None, name=None, nat_policy=None, network=None, self_link=None, self_link_with_id=None, zone=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if instance and not isinstance(instance, str):
            raise TypeError("Expected argument 'instance' to be a str")
        pulumi.set(__self__, "instance", instance)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nat_policy and not isinstance(nat_policy, str):
            raise TypeError("Expected argument 'nat_policy' to be a str")
        pulumi.set(__self__, "nat_policy", nat_policy)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id and not isinstance(self_link_with_id, str):
            raise TypeError("Expected argument 'self_link_with_id' to be a str")
        pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

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
    def instance(self) -> str:
        """
        A URL to the virtual machine instance that handles traffic for this target instance. When creating a target instance, you can provide the fully-qualified URL or a valid partial URL to the desired virtual machine. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance - zones/zone/instances/instance 
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The type of the resource. Always compute#targetInstance for target instances.
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
    @pulumi.getter(name="natPolicy")
    def nat_policy(self) -> str:
        """
        NAT option controlling how IPs are NAT'ed to the instance. Currently only NO_NAT (default value) is supported.
        """
        return pulumi.get(self, "nat_policy")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
        """
        return pulumi.get(self, "network")

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

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        URL of the zone where the target instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "zone")


class AwaitableGetTargetInstanceResult(GetTargetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTargetInstanceResult(
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            instance=self.instance,
            kind=self.kind,
            name=self.name,
            nat_policy=self.nat_policy,
            network=self.network,
            self_link=self.self_link,
            self_link_with_id=self.self_link_with_id,
            zone=self.zone)


def get_target_instance(project: Optional[str] = None,
                        target_instance: Optional[str] = None,
                        zone: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTargetInstanceResult:
    """
    Returns the specified TargetInstance resource. Gets a list of available target instances by making a list() request.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['targetInstance'] = target_instance
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/alpha:getTargetInstance', __args__, opts=opts, typ=GetTargetInstanceResult).value

    return AwaitableGetTargetInstanceResult(
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        instance=__ret__.instance,
        kind=__ret__.kind,
        name=__ret__.name,
        nat_policy=__ret__.nat_policy,
        network=__ret__.network,
        self_link=__ret__.self_link,
        self_link_with_id=__ret__.self_link_with_id,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_target_instance)
def get_target_instance_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                               target_instance: Optional[pulumi.Input[str]] = None,
                               zone: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTargetInstanceResult]:
    """
    Returns the specified TargetInstance resource. Gets a list of available target instances by making a list() request.
    """
    ...
