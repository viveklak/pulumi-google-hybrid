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
    'GetDeveloperResult',
    'AwaitableGetDeveloperResult',
    'get_developer',
    'get_developer_output',
]

@pulumi.output_type
class GetDeveloperResult:
    def __init__(__self__, access_type=None, app_family=None, apps=None, attributes=None, companies=None, created_at=None, developer_id=None, email=None, first_name=None, last_modified_at=None, last_name=None, organization_name=None, status=None, user_name=None):
        if access_type and not isinstance(access_type, str):
            raise TypeError("Expected argument 'access_type' to be a str")
        pulumi.set(__self__, "access_type", access_type)
        if app_family and not isinstance(app_family, str):
            raise TypeError("Expected argument 'app_family' to be a str")
        pulumi.set(__self__, "app_family", app_family)
        if apps and not isinstance(apps, list):
            raise TypeError("Expected argument 'apps' to be a list")
        pulumi.set(__self__, "apps", apps)
        if attributes and not isinstance(attributes, list):
            raise TypeError("Expected argument 'attributes' to be a list")
        pulumi.set(__self__, "attributes", attributes)
        if companies and not isinstance(companies, list):
            raise TypeError("Expected argument 'companies' to be a list")
        pulumi.set(__self__, "companies", companies)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if developer_id and not isinstance(developer_id, str):
            raise TypeError("Expected argument 'developer_id' to be a str")
        pulumi.set(__self__, "developer_id", developer_id)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if first_name and not isinstance(first_name, str):
            raise TypeError("Expected argument 'first_name' to be a str")
        pulumi.set(__self__, "first_name", first_name)
        if last_modified_at and not isinstance(last_modified_at, str):
            raise TypeError("Expected argument 'last_modified_at' to be a str")
        pulumi.set(__self__, "last_modified_at", last_modified_at)
        if last_name and not isinstance(last_name, str):
            raise TypeError("Expected argument 'last_name' to be a str")
        pulumi.set(__self__, "last_name", last_name)
        if organization_name and not isinstance(organization_name, str):
            raise TypeError("Expected argument 'organization_name' to be a str")
        pulumi.set(__self__, "organization_name", organization_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> str:
        """
        Access type.
        """
        return pulumi.get(self, "access_type")

    @property
    @pulumi.getter(name="appFamily")
    def app_family(self) -> str:
        """
        Developer app family.
        """
        return pulumi.get(self, "app_family")

    @property
    @pulumi.getter
    def apps(self) -> Sequence[str]:
        """
        List of apps associated with the developer.
        """
        return pulumi.get(self, "apps")

    @property
    @pulumi.getter
    def attributes(self) -> Sequence['outputs.GoogleCloudApigeeV1AttributeResponse']:
        """
        Optional. Developer attributes (name/value pairs). The custom attribute limit is 18.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def companies(self) -> Sequence[str]:
        """
        List of companies associated with the developer.
        """
        return pulumi.get(self, "companies")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Time at which the developer was created in milliseconds since epoch.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="developerId")
    def developer_id(self) -> str:
        """
        ID of the developer. **Note**: IDs are generated internally by Apigee and are not guaranteed to stay the same over time.
        """
        return pulumi.get(self, "developer_id")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> str:
        """
        First name of the developer.
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> str:
        """
        Time at which the developer was last modified in milliseconds since epoch.
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> str:
        """
        Last name of the developer.
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter(name="organizationName")
    def organization_name(self) -> str:
        """
        Name of the Apigee organization in which the developer resides.
        """
        return pulumi.get(self, "organization_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the developer. Valid values are `active` and `inactive`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        User name of the developer. Not used by Apigee hybrid.
        """
        return pulumi.get(self, "user_name")


class AwaitableGetDeveloperResult(GetDeveloperResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeveloperResult(
            access_type=self.access_type,
            app_family=self.app_family,
            apps=self.apps,
            attributes=self.attributes,
            companies=self.companies,
            created_at=self.created_at,
            developer_id=self.developer_id,
            email=self.email,
            first_name=self.first_name,
            last_modified_at=self.last_modified_at,
            last_name=self.last_name,
            organization_name=self.organization_name,
            status=self.status,
            user_name=self.user_name)


def get_developer(action: Optional[str] = None,
                  developer_id: Optional[str] = None,
                  organization_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeveloperResult:
    """
    Returns the developer details, including the developer's name, email address, apps, and other information. **Note**: The response includes only the first 100 developer apps.
    """
    __args__ = dict()
    __args__['action'] = action
    __args__['developerId'] = developer_id
    __args__['organizationId'] = organization_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:apigee/v1:getDeveloper', __args__, opts=opts, typ=GetDeveloperResult).value

    return AwaitableGetDeveloperResult(
        access_type=__ret__.access_type,
        app_family=__ret__.app_family,
        apps=__ret__.apps,
        attributes=__ret__.attributes,
        companies=__ret__.companies,
        created_at=__ret__.created_at,
        developer_id=__ret__.developer_id,
        email=__ret__.email,
        first_name=__ret__.first_name,
        last_modified_at=__ret__.last_modified_at,
        last_name=__ret__.last_name,
        organization_name=__ret__.organization_name,
        status=__ret__.status,
        user_name=__ret__.user_name)


@_utilities.lift_output_func(get_developer)
def get_developer_output(action: Optional[pulumi.Input[Optional[str]]] = None,
                         developer_id: Optional[pulumi.Input[str]] = None,
                         organization_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeveloperResult]:
    """
    Returns the developer details, including the developer's name, email address, apps, and other information. **Note**: The response includes only the first 100 developer apps.
    """
    ...
