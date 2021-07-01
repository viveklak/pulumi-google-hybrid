# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'InstanceTier',
    'NetworkConfigModesItem',
    'NfsExportOptionsAccessMode',
    'NfsExportOptionsSquashMode',
]


class InstanceTier(str, Enum):
    """
    The service tier of the instance.
    """
    TIER_UNSPECIFIED = "TIER_UNSPECIFIED"
    """Not set."""
    STANDARD = "STANDARD"
    """STANDARD tier. BASIC_HDD is the preferred term for this tier."""
    PREMIUM = "PREMIUM"
    """PREMIUM tier. BASIC_SSD is the preferred term for this tier."""
    BASIC_HDD = "BASIC_HDD"
    """BASIC instances offer a maximum capacity of 63.9 TB. BASIC_HDD is an alias for STANDARD Tier, offering economical performance backed by HDD."""
    BASIC_SSD = "BASIC_SSD"
    """BASIC instances offer a maximum capacity of 63.9 TB. BASIC_SSD is an alias for PREMIUM Tier, and offers improved performance backed by SSD."""
    HIGH_SCALE_SSD = "HIGH_SCALE_SSD"
    """HIGH_SCALE instances offer expanded capacity and performance scaling capabilities."""


class NetworkConfigModesItem(str, Enum):
    ADDRESS_MODE_UNSPECIFIED = "ADDRESS_MODE_UNSPECIFIED"
    """Internet protocol not set."""
    MODE_IPV4 = "MODE_IPV4"
    """Use the IPv4 internet protocol."""


class NfsExportOptionsAccessMode(str, Enum):
    """
    Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
    """
    ACCESS_MODE_UNSPECIFIED = "ACCESS_MODE_UNSPECIFIED"
    """AccessMode not set."""
    READ_ONLY = "READ_ONLY"
    """The client can only read the file share."""
    READ_WRITE = "READ_WRITE"
    """The client can read and write the file share (default)."""


class NfsExportOptionsSquashMode(str, Enum):
    """
    Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH.
    """
    SQUASH_MODE_UNSPECIFIED = "SQUASH_MODE_UNSPECIFIED"
    """SquashMode not set."""
    NO_ROOT_SQUASH = "NO_ROOT_SQUASH"
    """The Root user has root access to the file share (default)."""
    ROOT_SQUASH = "ROOT_SQUASH"
    """The Root user has squashed access to the anonymous uid/gid."""
