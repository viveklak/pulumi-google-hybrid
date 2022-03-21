# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetHmacKeyResult',
    'AwaitableGetHmacKeyResult',
    'get_hmac_key',
    'get_hmac_key_output',
]

@pulumi.output_type
class GetHmacKeyResult:
    def __init__(__self__, access_id=None, etag=None, kind=None, project=None, self_link=None, service_account_email=None, state=None, time_created=None, updated=None):
        if access_id and not isinstance(access_id, str):
            raise TypeError("Expected argument 'access_id' to be a str")
        pulumi.set(__self__, "access_id", access_id)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if service_account_email and not isinstance(service_account_email, str):
            raise TypeError("Expected argument 'service_account_email' to be a str")
        pulumi.set(__self__, "service_account_email", service_account_email)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if time_created and not isinstance(time_created, str):
            raise TypeError("Expected argument 'time_created' to be a str")
        pulumi.set(__self__, "time_created", time_created)
        if updated and not isinstance(updated, str):
            raise TypeError("Expected argument 'updated' to be a str")
        pulumi.set(__self__, "updated", updated)

    @property
    @pulumi.getter(name="accessId")
    def access_id(self) -> str:
        """
        The ID of the HMAC Key.
        """
        return pulumi.get(self, "access_id")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        HTTP 1.1 Entity tag for the HMAC key.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The kind of item this is. For HMAC Key metadata, this is always storage#hmacKeyMetadata.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        Project ID owning the service account to which the key authenticates.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The link to this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> str:
        """
        The email address of the key's associated service account.
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the key. Can be one of ACTIVE, INACTIVE, or DELETED.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The creation time of the HMAC key in RFC 3339 format.
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter
    def updated(self) -> str:
        """
        The last modification time of the HMAC key metadata in RFC 3339 format.
        """
        return pulumi.get(self, "updated")


class AwaitableGetHmacKeyResult(GetHmacKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHmacKeyResult(
            access_id=self.access_id,
            etag=self.etag,
            kind=self.kind,
            project=self.project,
            self_link=self.self_link,
            service_account_email=self.service_account_email,
            state=self.state,
            time_created=self.time_created,
            updated=self.updated)


def get_hmac_key(access_id: Optional[str] = None,
                 project: Optional[str] = None,
                 user_project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHmacKeyResult:
    """
    Retrieves an HMAC key's metadata
    """
    __args__ = dict()
    __args__['accessId'] = access_id
    __args__['project'] = project
    __args__['userProject'] = user_project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:storage/v1:getHmacKey', __args__, opts=opts, typ=GetHmacKeyResult).value

    return AwaitableGetHmacKeyResult(
        access_id=__ret__.access_id,
        etag=__ret__.etag,
        kind=__ret__.kind,
        project=__ret__.project,
        self_link=__ret__.self_link,
        service_account_email=__ret__.service_account_email,
        state=__ret__.state,
        time_created=__ret__.time_created,
        updated=__ret__.updated)


@_utilities.lift_output_func(get_hmac_key)
def get_hmac_key_output(access_id: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        user_project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHmacKeyResult]:
    """
    Retrieves an HMAC key's metadata
    """
    ...
