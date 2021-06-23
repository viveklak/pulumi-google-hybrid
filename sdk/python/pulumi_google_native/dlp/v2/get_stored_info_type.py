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
    'GetStoredInfoTypeResult',
    'AwaitableGetStoredInfoTypeResult',
    'get_stored_info_type',
]

@pulumi.output_type
class GetStoredInfoTypeResult:
    def __init__(__self__, current_version=None, name=None, pending_versions=None):
        if current_version and not isinstance(current_version, dict):
            raise TypeError("Expected argument 'current_version' to be a dict")
        pulumi.set(__self__, "current_version", current_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pending_versions and not isinstance(pending_versions, list):
            raise TypeError("Expected argument 'pending_versions' to be a list")
        pulumi.set(__self__, "pending_versions", pending_versions)

    @property
    @pulumi.getter(name="currentVersion")
    def current_version(self) -> 'outputs.GooglePrivacyDlpV2StoredInfoTypeVersionResponse':
        """
        Current version of the stored info type.
        """
        return pulumi.get(self, "current_version")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pendingVersions")
    def pending_versions(self) -> Sequence['outputs.GooglePrivacyDlpV2StoredInfoTypeVersionResponse']:
        """
        Pending versions of the stored info type. Empty if no versions are pending.
        """
        return pulumi.get(self, "pending_versions")


class AwaitableGetStoredInfoTypeResult(GetStoredInfoTypeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStoredInfoTypeResult(
            current_version=self.current_version,
            name=self.name,
            pending_versions=self.pending_versions)


def get_stored_info_type(location: Optional[str] = None,
                         project: Optional[str] = None,
                         stored_info_type_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStoredInfoTypeResult:
    """
    Gets a stored infoType. See https://cloud.google.com/dlp/docs/creating-stored-infotypes to learn more.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['storedInfoTypeId'] = stored_info_type_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:dlp/v2:getStoredInfoType', __args__, opts=opts, typ=GetStoredInfoTypeResult).value

    return AwaitableGetStoredInfoTypeResult(
        current_version=__ret__.current_version,
        name=__ret__.name,
        pending_versions=__ret__.pending_versions)