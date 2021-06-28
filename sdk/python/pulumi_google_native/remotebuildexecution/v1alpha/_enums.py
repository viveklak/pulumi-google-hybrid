# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeaturePolicy',
    'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyLinuxIsolation',
]


class GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeaturePolicy(str, Enum):
    """
    The policy of the feature.
    """
    POLICY_UNSPECIFIED = "POLICY_UNSPECIFIED"
    ALLOWED = "ALLOWED"
    FORBIDDEN = "FORBIDDEN"
    RESTRICTED = "RESTRICTED"


class GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyLinuxIsolation(str, Enum):
    """
    linux_isolation allows overriding the docker runtime used for containers started on Linux.
    """
    LINUX_ISOLATION_UNSPECIFIED = "LINUX_ISOLATION_UNSPECIFIED"
    GVISOR = "GVISOR"
    OFF = "OFF"
