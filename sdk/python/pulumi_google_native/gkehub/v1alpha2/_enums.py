# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuditLogConfigLogType',
    'MembershipInfrastructureType',
]


class AuditLogConfigLogType(str, Enum):
    """
    The log type that this config enables.
    """
    LOG_TYPE_UNSPECIFIED = "LOG_TYPE_UNSPECIFIED"
    """
    Default case. Should never be this.
    """
    ADMIN_READ = "ADMIN_READ"
    """
    Admin reads. Example: CloudIAM getIamPolicy
    """
    DATA_WRITE = "DATA_WRITE"
    """
    Data writes. Example: CloudSQL Users create
    """
    DATA_READ = "DATA_READ"
    """
    Data reads. Example: CloudSQL Users list
    """


class MembershipInfrastructureType(str, Enum):
    """
    Optional. The infrastructure type this Membership is running on.
    """
    INFRASTRUCTURE_TYPE_UNSPECIFIED = "INFRASTRUCTURE_TYPE_UNSPECIFIED"
    """
    No type was specified. Some Hub functionality may require a type be specified, and will not support Memberships with this value.
    """
    ON_PREM = "ON_PREM"
    """
    Private infrastructure that is owned or operated by customer. This includes GKE distributions such as GKE-OnPrem and GKE-OnBareMetal.
    """
    MULTI_CLOUD = "MULTI_CLOUD"
    """
    Public cloud infrastructure.
    """
