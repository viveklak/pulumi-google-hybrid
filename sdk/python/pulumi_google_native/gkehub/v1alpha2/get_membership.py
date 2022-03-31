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
    'GetMembershipResult',
    'AwaitableGetMembershipResult',
    'get_membership',
    'get_membership_output',
]

@pulumi.output_type
class GetMembershipResult:
    def __init__(__self__, authority=None, create_time=None, delete_time=None, description=None, endpoint=None, external_id=None, infrastructure_type=None, labels=None, last_connection_time=None, name=None, state=None, unique_id=None, update_time=None):
        if authority and not isinstance(authority, dict):
            raise TypeError("Expected argument 'authority' to be a dict")
        pulumi.set(__self__, "authority", authority)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if delete_time and not isinstance(delete_time, str):
            raise TypeError("Expected argument 'delete_time' to be a str")
        pulumi.set(__self__, "delete_time", delete_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if endpoint and not isinstance(endpoint, dict):
            raise TypeError("Expected argument 'endpoint' to be a dict")
        pulumi.set(__self__, "endpoint", endpoint)
        if external_id and not isinstance(external_id, str):
            raise TypeError("Expected argument 'external_id' to be a str")
        pulumi.set(__self__, "external_id", external_id)
        if infrastructure_type and not isinstance(infrastructure_type, str):
            raise TypeError("Expected argument 'infrastructure_type' to be a str")
        pulumi.set(__self__, "infrastructure_type", infrastructure_type)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if last_connection_time and not isinstance(last_connection_time, str):
            raise TypeError("Expected argument 'last_connection_time' to be a str")
        pulumi.set(__self__, "last_connection_time", last_connection_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, dict):
            raise TypeError("Expected argument 'state' to be a dict")
        pulumi.set(__self__, "state", state)
        if unique_id and not isinstance(unique_id, str):
            raise TypeError("Expected argument 'unique_id' to be a str")
        pulumi.set(__self__, "unique_id", unique_id)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def authority(self) -> 'outputs.AuthorityResponse':
        """
        Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
        """
        return pulumi.get(self, "authority")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        When the Membership was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> str:
        """
        When the Membership was deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of this membership, limited to 63 characters. Must match the regex: `a-zA-Z0-9*` This field is present for legacy purposes.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def endpoint(self) -> 'outputs.MembershipEndpointResponse':
        """
        Optional. Endpoint information to reach this member.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> str:
        """
        Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. For GKE clusters, external_id is managed by the Hub API and updates will be ignored. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter(name="infrastructureType")
    def infrastructure_type(self) -> str:
        """
        Optional. The infrastructure type this Membership is running on.
        """
        return pulumi.get(self, "infrastructure_type")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. GCP labels for this membership.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastConnectionTime")
    def last_connection_time(self) -> str:
        """
        For clusters using Connect, the timestamp of the most recent connection established with Google Cloud. This time is updated every several minutes, not continuously. For clusters that do not use GKE Connect, or that have never connected successfully, this field will be unset.
        """
        return pulumi.get(self, "last_connection_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The full, unique name of this Membership resource in the format `projects/*/locations/*/memberships/{membership_id}`, set during creation. `membership_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> 'outputs.MembershipStateResponse':
        """
        State of the Membership resource.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> str:
        """
        Google-generated UUID for this resource. This is unique across all Membership resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id.
        """
        return pulumi.get(self, "unique_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        When the Membership was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetMembershipResult(GetMembershipResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMembershipResult(
            authority=self.authority,
            create_time=self.create_time,
            delete_time=self.delete_time,
            description=self.description,
            endpoint=self.endpoint,
            external_id=self.external_id,
            infrastructure_type=self.infrastructure_type,
            labels=self.labels,
            last_connection_time=self.last_connection_time,
            name=self.name,
            state=self.state,
            unique_id=self.unique_id,
            update_time=self.update_time)


def get_membership(location: Optional[str] = None,
                   membership_id: Optional[str] = None,
                   project: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMembershipResult:
    """
    Gets the details of a Membership.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['membershipId'] = membership_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:gkehub/v1alpha2:getMembership', __args__, opts=opts, typ=GetMembershipResult).value

    return AwaitableGetMembershipResult(
        authority=__ret__.authority,
        create_time=__ret__.create_time,
        delete_time=__ret__.delete_time,
        description=__ret__.description,
        endpoint=__ret__.endpoint,
        external_id=__ret__.external_id,
        infrastructure_type=__ret__.infrastructure_type,
        labels=__ret__.labels,
        last_connection_time=__ret__.last_connection_time,
        name=__ret__.name,
        state=__ret__.state,
        unique_id=__ret__.unique_id,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_membership)
def get_membership_output(location: Optional[pulumi.Input[str]] = None,
                          membership_id: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMembershipResult]:
    """
    Gets the details of a Membership.
    """
    ...
