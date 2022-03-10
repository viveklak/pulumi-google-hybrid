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
    'GetInstanceResult',
    'AwaitableGetInstanceResult',
    'get_instance',
    'get_instance_output',
]

@pulumi.output_type
class GetInstanceResult:
    def __init__(__self__, create_time=None, description=None, etag=None, file_shares=None, kms_key_name=None, labels=None, max_share_count=None, name=None, networks=None, satisfies_pzs=None, state=None, status_message=None, suspension_reasons=None, tier=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if file_shares and not isinstance(file_shares, list):
            raise TypeError("Expected argument 'file_shares' to be a list")
        pulumi.set(__self__, "file_shares", file_shares)
        if kms_key_name and not isinstance(kms_key_name, str):
            raise TypeError("Expected argument 'kms_key_name' to be a str")
        pulumi.set(__self__, "kms_key_name", kms_key_name)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if max_share_count and not isinstance(max_share_count, str):
            raise TypeError("Expected argument 'max_share_count' to be a str")
        pulumi.set(__self__, "max_share_count", max_share_count)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status_message and not isinstance(status_message, str):
            raise TypeError("Expected argument 'status_message' to be a str")
        pulumi.set(__self__, "status_message", status_message)
        if suspension_reasons and not isinstance(suspension_reasons, list):
            raise TypeError("Expected argument 'suspension_reasons' to be a list")
        pulumi.set(__self__, "suspension_reasons", suspension_reasons)
        if tier and not isinstance(tier, str):
            raise TypeError("Expected argument 'tier' to be a str")
        pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the instance was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the instance (2048 characters or less).
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="fileShares")
    def file_shares(self) -> Sequence['outputs.FileShareConfigResponse']:
        """
        File system shares on the instance. For this version, only a single file share is supported.
        """
        return pulumi.get(self, "file_shares")

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> str:
        """
        KMS key name used for data encryption.
        """
        return pulumi.get(self, "kms_key_name")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Resource labels to represent user provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="maxShareCount")
    def max_share_count(self) -> str:
        """
        The max number of shares allowed.
        """
        return pulumi.get(self, "max_share_count")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the instance, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.NetworkConfigResponse']:
        """
        VPC networks to which the instance is connected. For this version, only a single network is supported.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The instance state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> str:
        """
        Additional information about the instance state, if available.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter(name="suspensionReasons")
    def suspension_reasons(self) -> Sequence[str]:
        """
        field indicates all the reasons the instance is in "SUSPENDED" state.
        """
        return pulumi.get(self, "suspension_reasons")

    @property
    @pulumi.getter
    def tier(self) -> str:
        """
        The service tier of the instance.
        """
        return pulumi.get(self, "tier")


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            create_time=self.create_time,
            description=self.description,
            etag=self.etag,
            file_shares=self.file_shares,
            kms_key_name=self.kms_key_name,
            labels=self.labels,
            max_share_count=self.max_share_count,
            name=self.name,
            networks=self.networks,
            satisfies_pzs=self.satisfies_pzs,
            state=self.state,
            status_message=self.status_message,
            suspension_reasons=self.suspension_reasons,
            tier=self.tier)


def get_instance(instance_id: Optional[str] = None,
                 location: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceResult:
    """
    Gets the details of a specific instance.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:file/v1beta1:getInstance', __args__, opts=opts, typ=GetInstanceResult).value

    return AwaitableGetInstanceResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        etag=__ret__.etag,
        file_shares=__ret__.file_shares,
        kms_key_name=__ret__.kms_key_name,
        labels=__ret__.labels,
        max_share_count=__ret__.max_share_count,
        name=__ret__.name,
        networks=__ret__.networks,
        satisfies_pzs=__ret__.satisfies_pzs,
        state=__ret__.state,
        status_message=__ret__.status_message,
        suspension_reasons=__ret__.suspension_reasons,
        tier=__ret__.tier)


@_utilities.lift_output_func(get_instance)
def get_instance_output(instance_id: Optional[pulumi.Input[str]] = None,
                        location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceResult]:
    """
    Gets the details of a specific instance.
    """
    ...