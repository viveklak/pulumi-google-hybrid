# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_google_native.jobs.v3 as __v3
    v3 = __v3
    import pulumi_google_native.jobs.v4 as __v4
    v4 = __v4
else:
    v3 = _utilities.lazy_import('pulumi_google_native.jobs.v3')
    v4 = _utilities.lazy_import('pulumi_google_native.jobs.v4')

