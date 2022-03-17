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
    def __init__(__self__, create_time=None, name=None, preferred_member_key=None, roles=None, type=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if preferred_member_key and not isinstance(preferred_member_key, dict):
            raise TypeError("Expected argument 'preferred_member_key' to be a dict")
        pulumi.set(__self__, "preferred_member_key", preferred_member_key)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the `Membership` was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Membership`. Shall be of the form `groups/{group}/memberships/{membership}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="preferredMemberKey")
    def preferred_member_key(self) -> 'outputs.EntityKeyResponse':
        """
        Immutable. The `EntityKey` of the member.
        """
        return pulumi.get(self, "preferred_member_key")

    @property
    @pulumi.getter
    def roles(self) -> Sequence['outputs.MembershipRoleResponse']:
        """
        The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the membership.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the `Membership` was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetMembershipResult(GetMembershipResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMembershipResult(
            create_time=self.create_time,
            name=self.name,
            preferred_member_key=self.preferred_member_key,
            roles=self.roles,
            type=self.type,
            update_time=self.update_time)


def get_membership(group_id: Optional[str] = None,
                   membership_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMembershipResult:
    """
    Retrieves a `Membership`.
    """
    __args__ = dict()
    __args__['groupId'] = group_id
    __args__['membershipId'] = membership_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:cloudidentity/v1:getMembership', __args__, opts=opts, typ=GetMembershipResult).value

    return AwaitableGetMembershipResult(
        create_time=__ret__.create_time,
        name=__ret__.name,
        preferred_member_key=__ret__.preferred_member_key,
        roles=__ret__.roles,
        type=__ret__.type,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_membership)
def get_membership_output(group_id: Optional[pulumi.Input[str]] = None,
                          membership_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMembershipResult]:
    """
    Retrieves a `Membership`.
    """
    ...
