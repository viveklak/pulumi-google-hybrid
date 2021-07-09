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
    'GetAttestorResult',
    'AwaitableGetAttestorResult',
    'get_attestor',
]

@pulumi.output_type
class GetAttestorResult:
    def __init__(__self__, description=None, name=None, update_time=None, user_owned_drydock_note=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if user_owned_drydock_note and not isinstance(user_owned_drydock_note, dict):
            raise TypeError("Expected argument 'user_owned_drydock_note' to be a dict")
        pulumi.set(__self__, "user_owned_drydock_note", user_owned_drydock_note)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Time when the attestor was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="userOwnedDrydockNote")
    def user_owned_drydock_note(self) -> 'outputs.UserOwnedDrydockNoteResponse':
        """
        A Drydock ATTESTATION_AUTHORITY Note, created by the user.
        """
        return pulumi.get(self, "user_owned_drydock_note")


class AwaitableGetAttestorResult(GetAttestorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAttestorResult(
            description=self.description,
            name=self.name,
            update_time=self.update_time,
            user_owned_drydock_note=self.user_owned_drydock_note)


def get_attestor(attestor_id: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAttestorResult:
    """
    Gets an attestor. Returns NOT_FOUND if the attestor does not exist.
    """
    __args__ = dict()
    __args__['attestorId'] = attestor_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:binaryauthorization/v1beta1:getAttestor', __args__, opts=opts, typ=GetAttestorResult).value

    return AwaitableGetAttestorResult(
        description=__ret__.description,
        name=__ret__.name,
        update_time=__ret__.update_time,
        user_owned_drydock_note=__ret__.user_owned_drydock_note)
