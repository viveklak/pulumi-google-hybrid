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
    'GetAuthorizedCertificateResult',
    'AwaitableGetAuthorizedCertificateResult',
    'get_authorized_certificate',
]

@pulumi.output_type
class GetAuthorizedCertificateResult:
    def __init__(__self__, certificate_raw_data=None, display_name=None, domain_mappings_count=None, domain_names=None, expire_time=None, managed_certificate=None, name=None, visible_domain_mappings=None):
        if certificate_raw_data and not isinstance(certificate_raw_data, dict):
            raise TypeError("Expected argument 'certificate_raw_data' to be a dict")
        pulumi.set(__self__, "certificate_raw_data", certificate_raw_data)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if domain_mappings_count and not isinstance(domain_mappings_count, int):
            raise TypeError("Expected argument 'domain_mappings_count' to be a int")
        pulumi.set(__self__, "domain_mappings_count", domain_mappings_count)
        if domain_names and not isinstance(domain_names, list):
            raise TypeError("Expected argument 'domain_names' to be a list")
        pulumi.set(__self__, "domain_names", domain_names)
        if expire_time and not isinstance(expire_time, str):
            raise TypeError("Expected argument 'expire_time' to be a str")
        pulumi.set(__self__, "expire_time", expire_time)
        if managed_certificate and not isinstance(managed_certificate, dict):
            raise TypeError("Expected argument 'managed_certificate' to be a dict")
        pulumi.set(__self__, "managed_certificate", managed_certificate)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if visible_domain_mappings and not isinstance(visible_domain_mappings, list):
            raise TypeError("Expected argument 'visible_domain_mappings' to be a list")
        pulumi.set(__self__, "visible_domain_mappings", visible_domain_mappings)

    @property
    @pulumi.getter(name="certificateRawData")
    def certificate_raw_data(self) -> 'outputs.CertificateRawDataResponse':
        """
        The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
        """
        return pulumi.get(self, "certificate_raw_data")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="domainMappingsCount")
    def domain_mappings_count(self) -> int:
        """
        Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.
        """
        return pulumi.get(self, "domain_mappings_count")

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Sequence[str]:
        """
        Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.
        """
        return pulumi.get(self, "domain_names")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="managedCertificate")
    def managed_certificate(self) -> 'outputs.ManagedCertificateResponse':
        """
        Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.
        """
        return pulumi.get(self, "managed_certificate")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="visibleDomainMappings")
    def visible_domain_mappings(self) -> Sequence[str]:
        """
        The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.
        """
        return pulumi.get(self, "visible_domain_mappings")


class AwaitableGetAuthorizedCertificateResult(GetAuthorizedCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthorizedCertificateResult(
            certificate_raw_data=self.certificate_raw_data,
            display_name=self.display_name,
            domain_mappings_count=self.domain_mappings_count,
            domain_names=self.domain_names,
            expire_time=self.expire_time,
            managed_certificate=self.managed_certificate,
            name=self.name,
            visible_domain_mappings=self.visible_domain_mappings)


def get_authorized_certificate(app_id: Optional[str] = None,
                               authorized_certificate_id: Optional[str] = None,
                               view: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthorizedCertificateResult:
    """
    Gets the specified SSL certificate.
    """
    __args__ = dict()
    __args__['appId'] = app_id
    __args__['authorizedCertificateId'] = authorized_certificate_id
    __args__['view'] = view
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:appengine/v1:getAuthorizedCertificate', __args__, opts=opts, typ=GetAuthorizedCertificateResult).value

    return AwaitableGetAuthorizedCertificateResult(
        certificate_raw_data=__ret__.certificate_raw_data,
        display_name=__ret__.display_name,
        domain_mappings_count=__ret__.domain_mappings_count,
        domain_names=__ret__.domain_names,
        expire_time=__ret__.expire_time,
        managed_certificate=__ret__.managed_certificate,
        name=__ret__.name,
        visible_domain_mappings=__ret__.visible_domain_mappings)
