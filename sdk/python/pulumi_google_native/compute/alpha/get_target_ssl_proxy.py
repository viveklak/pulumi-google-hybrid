# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetTargetSslProxyResult',
    'AwaitableGetTargetSslProxyResult',
    'get_target_ssl_proxy',
    'get_target_ssl_proxy_output',
]

@pulumi.output_type
class GetTargetSslProxyResult:
    def __init__(__self__, certificate_map=None, creation_timestamp=None, description=None, kind=None, name=None, proxy_header=None, self_link=None, service=None, ssl_certificates=None, ssl_policy=None):
        if certificate_map and not isinstance(certificate_map, str):
            raise TypeError("Expected argument 'certificate_map' to be a str")
        pulumi.set(__self__, "certificate_map", certificate_map)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if proxy_header and not isinstance(proxy_header, str):
            raise TypeError("Expected argument 'proxy_header' to be a str")
        pulumi.set(__self__, "proxy_header", proxy_header)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if service and not isinstance(service, str):
            raise TypeError("Expected argument 'service' to be a str")
        pulumi.set(__self__, "service", service)
        if ssl_certificates and not isinstance(ssl_certificates, list):
            raise TypeError("Expected argument 'ssl_certificates' to be a list")
        pulumi.set(__self__, "ssl_certificates", ssl_certificates)
        if ssl_policy and not isinstance(ssl_policy, str):
            raise TypeError("Expected argument 'ssl_policy' to be a str")
        pulumi.set(__self__, "ssl_policy", ssl_policy)

    @property
    @pulumi.getter(name="certificateMap")
    def certificate_map(self) -> str:
        """
        URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored.
        """
        return pulumi.get(self, "certificate_map")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#targetSslProxy for target SSL proxies.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="proxyHeader")
    def proxy_header(self) -> str:
        """
        Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
        """
        return pulumi.get(self, "proxy_header")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def service(self) -> str:
        """
        URL to the BackendService resource.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter(name="sslCertificates")
    def ssl_certificates(self) -> Sequence[str]:
        """
        URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
        """
        return pulumi.get(self, "ssl_certificates")

    @property
    @pulumi.getter(name="sslPolicy")
    def ssl_policy(self) -> str:
        """
        URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
        """
        return pulumi.get(self, "ssl_policy")


class AwaitableGetTargetSslProxyResult(GetTargetSslProxyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTargetSslProxyResult(
            certificate_map=self.certificate_map,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            kind=self.kind,
            name=self.name,
            proxy_header=self.proxy_header,
            self_link=self.self_link,
            service=self.service,
            ssl_certificates=self.ssl_certificates,
            ssl_policy=self.ssl_policy)


def get_target_ssl_proxy(project: Optional[str] = None,
                         target_ssl_proxy: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTargetSslProxyResult:
    """
    Returns the specified TargetSslProxy resource. Gets a list of available target SSL proxies by making a list() request.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['targetSslProxy'] = target_ssl_proxy
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/alpha:getTargetSslProxy', __args__, opts=opts, typ=GetTargetSslProxyResult).value

    return AwaitableGetTargetSslProxyResult(
        certificate_map=__ret__.certificate_map,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        kind=__ret__.kind,
        name=__ret__.name,
        proxy_header=__ret__.proxy_header,
        self_link=__ret__.self_link,
        service=__ret__.service,
        ssl_certificates=__ret__.ssl_certificates,
        ssl_policy=__ret__.ssl_policy)


@_utilities.lift_output_func(get_target_ssl_proxy)
def get_target_ssl_proxy_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                                target_ssl_proxy: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTargetSslProxyResult]:
    """
    Returns the specified TargetSslProxy resource. Gets a list of available target SSL proxies by making a list() request.
    """
    ...
