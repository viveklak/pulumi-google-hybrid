# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetPeeringResult',
    'AwaitableGetPeeringResult',
    'get_peering',
    'get_peering_output',
]

@pulumi.output_type
class GetPeeringResult:
    def __init__(__self__, authorized_network=None, create_time=None, domain_resource=None, labels=None, name=None, state=None, status_message=None, update_time=None):
        if authorized_network and not isinstance(authorized_network, str):
            raise TypeError("Expected argument 'authorized_network' to be a str")
        pulumi.set(__self__, "authorized_network", authorized_network)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if domain_resource and not isinstance(domain_resource, str):
            raise TypeError("Expected argument 'domain_resource' to be a str")
        pulumi.set(__self__, "domain_resource", domain_resource)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status_message and not isinstance(status_message, str):
            raise TypeError("Expected argument 'status_message' to be a str")
        pulumi.set(__self__, "status_message", status_message)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="authorizedNetwork")
    def authorized_network(self) -> str:
        """
        The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
        """
        return pulumi.get(self, "authorized_network")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time the instance was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="domainResource")
    def domain_resource(self) -> str:
        """
        Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form: `projects/{project_id}/locations/global/domains/{domain_name}`
        """
        return pulumi.get(self, "domain_resource")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. Resource labels to represent user provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique name of the peering in this scope including projects and location using the form: `projects/{project_id}/locations/global/peerings/{peering_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of this Peering.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> str:
        """
        Additional information about the current status of this peering, if available.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Last update time.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetPeeringResult(GetPeeringResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPeeringResult(
            authorized_network=self.authorized_network,
            create_time=self.create_time,
            domain_resource=self.domain_resource,
            labels=self.labels,
            name=self.name,
            state=self.state,
            status_message=self.status_message,
            update_time=self.update_time)


def get_peering(peering_id: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPeeringResult:
    """
    Gets details of a single Peering.
    """
    __args__ = dict()
    __args__['peeringId'] = peering_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:managedidentities/v1beta1:getPeering', __args__, opts=opts, typ=GetPeeringResult).value

    return AwaitableGetPeeringResult(
        authorized_network=__ret__.authorized_network,
        create_time=__ret__.create_time,
        domain_resource=__ret__.domain_resource,
        labels=__ret__.labels,
        name=__ret__.name,
        state=__ret__.state,
        status_message=__ret__.status_message,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_peering)
def get_peering_output(peering_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPeeringResult]:
    """
    Gets details of a single Peering.
    """
    ...
