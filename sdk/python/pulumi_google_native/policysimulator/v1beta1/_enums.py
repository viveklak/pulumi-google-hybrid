# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource',
]


class GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource(str, Enum):
    """
    The logs to use as input for the Replay.
    """
    LOG_SOURCE_UNSPECIFIED = "LOG_SOURCE_UNSPECIFIED"
    """An unspecified log source. If the log source is unspecified, the Replay defaults to using `RECENT_ACCESSES`."""
    RECENT_ACCESSES = "RECENT_ACCESSES"
    """All access logs from the last 90 days. These logs may not include logs from the most recent 7 days."""
