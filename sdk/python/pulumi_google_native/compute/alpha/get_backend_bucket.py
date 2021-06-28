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
    'GetBackendBucketResult',
    'AwaitableGetBackendBucketResult',
    'get_backend_bucket',
]

@pulumi.output_type
class GetBackendBucketResult:
    def __init__(__self__, bucket_name=None, cdn_policy=None, creation_timestamp=None, custom_response_headers=None, description=None, edge_security_policy=None, enable_cdn=None, kind=None, name=None, self_link=None, self_link_with_id=None):
        if bucket_name and not isinstance(bucket_name, str):
            raise TypeError("Expected argument 'bucket_name' to be a str")
        pulumi.set(__self__, "bucket_name", bucket_name)
        if cdn_policy and not isinstance(cdn_policy, dict):
            raise TypeError("Expected argument 'cdn_policy' to be a dict")
        pulumi.set(__self__, "cdn_policy", cdn_policy)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if custom_response_headers and not isinstance(custom_response_headers, list):
            raise TypeError("Expected argument 'custom_response_headers' to be a list")
        pulumi.set(__self__, "custom_response_headers", custom_response_headers)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if edge_security_policy and not isinstance(edge_security_policy, str):
            raise TypeError("Expected argument 'edge_security_policy' to be a str")
        pulumi.set(__self__, "edge_security_policy", edge_security_policy)
        if enable_cdn and not isinstance(enable_cdn, bool):
            raise TypeError("Expected argument 'enable_cdn' to be a bool")
        pulumi.set(__self__, "enable_cdn", enable_cdn)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id and not isinstance(self_link_with_id, str):
            raise TypeError("Expected argument 'self_link_with_id' to be a str")
        pulumi.set(__self__, "self_link_with_id", self_link_with_id)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> str:
        """
        Cloud Storage bucket name.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="cdnPolicy")
    def cdn_policy(self) -> 'outputs.BackendBucketCdnPolicyResponse':
        """
        Cloud CDN configuration for this BackendBucket.
        """
        return pulumi.get(self, "cdn_policy")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="customResponseHeaders")
    def custom_response_headers(self) -> Sequence[str]:
        """
        Headers that the HTTP/S load balancer should add to proxied responses.
        """
        return pulumi.get(self, "custom_response_headers")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional textual description of the resource; provided by the client when the resource is created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="edgeSecurityPolicy")
    def edge_security_policy(self) -> str:
        """
        The resource URL for the edge security policy associated with this backend bucket.
        """
        return pulumi.get(self, "edge_security_policy")

    @property
    @pulumi.getter(name="enableCdn")
    def enable_cdn(self) -> bool:
        """
        If true, enable Cloud CDN for this BackendBucket.
        """
        return pulumi.get(self, "enable_cdn")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource.
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
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> str:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")


class AwaitableGetBackendBucketResult(GetBackendBucketResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackendBucketResult(
            bucket_name=self.bucket_name,
            cdn_policy=self.cdn_policy,
            creation_timestamp=self.creation_timestamp,
            custom_response_headers=self.custom_response_headers,
            description=self.description,
            edge_security_policy=self.edge_security_policy,
            enable_cdn=self.enable_cdn,
            kind=self.kind,
            name=self.name,
            self_link=self.self_link,
            self_link_with_id=self.self_link_with_id)


def get_backend_bucket(backend_bucket: Optional[str] = None,
                       project: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackendBucketResult:
    """
    Returns the specified BackendBucket resource. Gets a list of available backend buckets by making a list() request.
    """
    __args__ = dict()
    __args__['backendBucket'] = backend_bucket
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/alpha:getBackendBucket', __args__, opts=opts, typ=GetBackendBucketResult).value

    return AwaitableGetBackendBucketResult(
        bucket_name=__ret__.bucket_name,
        cdn_policy=__ret__.cdn_policy,
        creation_timestamp=__ret__.creation_timestamp,
        custom_response_headers=__ret__.custom_response_headers,
        description=__ret__.description,
        edge_security_policy=__ret__.edge_security_policy,
        enable_cdn=__ret__.enable_cdn,
        kind=__ret__.kind,
        name=__ret__.name,
        self_link=__ret__.self_link,
        self_link_with_id=__ret__.self_link_with_id)
