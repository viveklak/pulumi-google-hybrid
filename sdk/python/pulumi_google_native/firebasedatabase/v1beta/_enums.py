# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'InstanceState',
    'InstanceType',
]


class InstanceState(str, Enum):
    """
    The database's lifecycle state. Read-only.
    """
    LIFECYCLE_STATE_UNSPECIFIED = "LIFECYCLE_STATE_UNSPECIFIED"
    """
    Unspecified state, likely the result of an error on the backend. This is only used for distinguishing unset values.
    """
    ACTIVE = "ACTIVE"
    """
    The normal and active state.
    """
    DISABLED = "DISABLED"
    """
    The database is in a disabled state. It can be re-enabled later.
    """
    DELETED = "DELETED"
    """
    The database is in a deleted state.
    """


class InstanceType(str, Enum):
    """
    The database instance type. On creation only USER_DATABASE is allowed, which is also the default when omitted.
    """
    DATABASE_INSTANCE_TYPE_UNSPECIFIED = "DATABASE_INSTANCE_TYPE_UNSPECIFIED"
    """
    Unknown state, likely the result of an error on the backend. This is only used for distinguishing unset values.
    """
    DEFAULT_DATABASE = "DEFAULT_DATABASE"
    """
    The default database that is provisioned when a project is created.
    """
    USER_DATABASE = "USER_DATABASE"
    """
    A database that the user created.
    """
