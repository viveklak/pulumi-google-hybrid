# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AnimationFadeFadeType',
    'ManifestType',
]


class AnimationFadeFadeType(str, Enum):
    """
    Required. Type of fade animation: `FADE_IN` or `FADE_OUT`.
    """
    FADE_TYPE_UNSPECIFIED = "FADE_TYPE_UNSPECIFIED"
    """
    The fade type is not specified.
    """
    FADE_IN = "FADE_IN"
    """
    Fade the overlay object into view.
    """
    FADE_OUT = "FADE_OUT"
    """
    Fade the overlay object out of view.
    """


class ManifestType(str, Enum):
    """
    Required. Type of the manifest, can be `HLS` or `DASH`.
    """
    MANIFEST_TYPE_UNSPECIFIED = "MANIFEST_TYPE_UNSPECIFIED"
    """
    The manifest type is not specified.
    """
    HLS = "HLS"
    """
    Create `HLS` manifest. The corresponding file extension is `.m3u8`.
    """
    DASH = "DASH"
    """
    Create `DASH` manifest. The corresponding file extension is `.mpd`.
    """
