# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_google_native.tpu.v1 as __v1
    v1 = __v1
    import pulumi_google_native.tpu.v1alpha1 as __v1alpha1
    v1alpha1 = __v1alpha1
    import pulumi_google_native.tpu.v2alpha1 as __v2alpha1
    v2alpha1 = __v2alpha1
else:
    v1 = _utilities.lazy_import('pulumi_google_native.tpu.v1')
    v1alpha1 = _utilities.lazy_import('pulumi_google_native.tpu.v1alpha1')
    v2alpha1 = _utilities.lazy_import('pulumi_google_native.tpu.v2alpha1')

