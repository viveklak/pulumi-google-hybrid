# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_google_native.dns.v1 as __v1
    v1 = __v1
    import pulumi_google_native.dns.v1beta2 as __v1beta2
    v1beta2 = __v1beta2
    import pulumi_google_native.dns.v2 as __v2
    v2 = __v2
else:
    v1 = _utilities.lazy_import('pulumi_google_native.dns.v1')
    v1beta2 = _utilities.lazy_import('pulumi_google_native.dns.v1beta2')
    v2 = _utilities.lazy_import('pulumi_google_native.dns.v2')

