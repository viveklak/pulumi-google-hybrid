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
    'GetDomainResult',
    'AwaitableGetDomainResult',
    'get_domain',
    'get_domain_output',
]

@pulumi.output_type
class GetDomainResult:
    def __init__(__self__, domain_name=None, domain_redirect=None, provisioning=None, site=None, status=None, update_time=None):
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        pulumi.set(__self__, "domain_name", domain_name)
        if domain_redirect and not isinstance(domain_redirect, dict):
            raise TypeError("Expected argument 'domain_redirect' to be a dict")
        pulumi.set(__self__, "domain_redirect", domain_redirect)
        if provisioning and not isinstance(provisioning, dict):
            raise TypeError("Expected argument 'provisioning' to be a dict")
        pulumi.set(__self__, "provisioning", provisioning)
        if site and not isinstance(site, str):
            raise TypeError("Expected argument 'site' to be a str")
        pulumi.set(__self__, "site", site)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> str:
        """
        The domain name of the association.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="domainRedirect")
    def domain_redirect(self) -> 'outputs.DomainRedirectResponse':
        """
        If set, the domain should redirect with the provided parameters.
        """
        return pulumi.get(self, "domain_redirect")

    @property
    @pulumi.getter
    def provisioning(self) -> 'outputs.DomainProvisioningResponse':
        """
        Information about the provisioning of certificates and the health of the DNS resolution for the domain.
        """
        return pulumi.get(self, "provisioning")

    @property
    @pulumi.getter
    def site(self) -> str:
        """
        The site name of the association.
        """
        return pulumi.get(self, "site")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Additional status of the domain association.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time at which the domain was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetDomainResult(GetDomainResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainResult(
            domain_name=self.domain_name,
            domain_redirect=self.domain_redirect,
            provisioning=self.provisioning,
            site=self.site,
            status=self.status,
            update_time=self.update_time)


def get_domain(domain_id: Optional[str] = None,
               project: Optional[str] = None,
               site_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainResult:
    """
    Gets a domain mapping on the specified site.
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    __args__['project'] = project
    __args__['siteId'] = site_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:firebasehosting/v1beta1:getDomain', __args__, opts=opts, typ=GetDomainResult).value

    return AwaitableGetDomainResult(
        domain_name=__ret__.domain_name,
        domain_redirect=__ret__.domain_redirect,
        provisioning=__ret__.provisioning,
        site=__ret__.site,
        status=__ret__.status,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_domain)
def get_domain_output(domain_id: Optional[pulumi.Input[str]] = None,
                      project: Optional[pulumi.Input[Optional[str]]] = None,
                      site_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainResult]:
    """
    Gets a domain mapping on the specified site.
    """
    ...
