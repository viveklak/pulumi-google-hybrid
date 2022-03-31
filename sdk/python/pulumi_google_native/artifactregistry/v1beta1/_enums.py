# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'RepositoryFormat',
]


class RepositoryFormat(str, Enum):
    """
    The format of packages that are stored in the repository.
    """
    FORMAT_UNSPECIFIED = "FORMAT_UNSPECIFIED"
    """
    Unspecified package format.
    """
    DOCKER = "DOCKER"
    """
    Docker package format.
    """
    MAVEN = "MAVEN"
    """
    Maven package format.
    """
    NPM = "NPM"
    """
    NPM package format.
    """
    APT = "APT"
    """
    APT package format.
    """
    YUM = "YUM"
    """
    YUM package format.
    """
    PYTHON = "PYTHON"
    """
    Python package format.
    """
