# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetTargetPoolResult',
    'AwaitableGetTargetPoolResult',
    'get_target_pool',
]

@pulumi.output_type
class GetTargetPoolResult:
    def __init__(__self__, backup_pool=None, creation_timestamp=None, description=None, failover_ratio=None, health_checks=None, instances=None, kind=None, name=None, region=None, self_link=None, session_affinity=None):
        if backup_pool and not isinstance(backup_pool, str):
            raise TypeError("Expected argument 'backup_pool' to be a str")
        pulumi.set(__self__, "backup_pool", backup_pool)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if failover_ratio and not isinstance(failover_ratio, float):
            raise TypeError("Expected argument 'failover_ratio' to be a float")
        pulumi.set(__self__, "failover_ratio", failover_ratio)
        if health_checks and not isinstance(health_checks, list):
            raise TypeError("Expected argument 'health_checks' to be a list")
        pulumi.set(__self__, "health_checks", health_checks)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if session_affinity and not isinstance(session_affinity, str):
            raise TypeError("Expected argument 'session_affinity' to be a str")
        pulumi.set(__self__, "session_affinity", session_affinity)

    @property
    @pulumi.getter(name="backupPool")
    def backup_pool(self) -> str:
        """
        The server-defined URL for the resource. This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool, and its failoverRatio field is properly set to a value between [0, 1].

        backupPool and failoverRatio together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below failoverRatio, traffic arriving at the load-balanced IP will be directed to the backup pool.

        In case where failoverRatio and backupPool are not set, or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
        """
        return pulumi.get(self, "backup_pool")

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
    @pulumi.getter(name="failoverRatio")
    def failover_ratio(self) -> float:
        """
        This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool (i.e., not as a backup pool to some other target pool). The value of the field must be in [0, 1].

        If set, backupPool must also be set. They together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below this number, traffic arriving at the load-balanced IP will be directed to the backup pool.

        In case where failoverRatio is not set or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
        """
        return pulumi.get(self, "failover_ratio")

    @property
    @pulumi.getter(name="healthChecks")
    def health_checks(self) -> Sequence[str]:
        """
        The URL of the HttpHealthCheck resource. A member instance in this pool is considered healthy if and only if the health checks pass. Only legacy HttpHealthChecks are supported. Only one health check may be specified.
        """
        return pulumi.get(self, "health_checks")

    @property
    @pulumi.getter
    def instances(self) -> Sequence[str]:
        """
        A list of resource URLs to the virtual machine instances serving this pool. They must live in zones contained in the same region as this pool.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#targetPool for target pools.
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
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the target pool resides.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="sessionAffinity")
    def session_affinity(self) -> str:
        """
        Session affinity option, must be one of the following values:
        NONE: Connections from the same client IP may go to any instance in the pool.
        CLIENT_IP: Connections from the same client IP will go to the same instance in the pool while that instance remains healthy.
        CLIENT_IP_PROTO: Connections from the same client IP with the same IP protocol will go to the same instance in the pool while that instance remains healthy.
        """
        return pulumi.get(self, "session_affinity")


class AwaitableGetTargetPoolResult(GetTargetPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTargetPoolResult(
            backup_pool=self.backup_pool,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            failover_ratio=self.failover_ratio,
            health_checks=self.health_checks,
            instances=self.instances,
            kind=self.kind,
            name=self.name,
            region=self.region,
            self_link=self.self_link,
            session_affinity=self.session_affinity)


def get_target_pool(project: Optional[str] = None,
                    region: Optional[str] = None,
                    target_pool: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTargetPoolResult:
    """
    Returns the specified target pool. Gets a list of available target pools by making a list() request.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['region'] = region
    __args__['targetPool'] = target_pool
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getTargetPool', __args__, opts=opts, typ=GetTargetPoolResult).value

    return AwaitableGetTargetPoolResult(
        backup_pool=__ret__.backup_pool,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        failover_ratio=__ret__.failover_ratio,
        health_checks=__ret__.health_checks,
        instances=__ret__.instances,
        kind=__ret__.kind,
        name=__ret__.name,
        region=__ret__.region,
        self_link=__ret__.self_link,
        session_affinity=__ret__.session_affinity)
