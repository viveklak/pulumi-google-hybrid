# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_google_native.metastore.v1alpha as __v1alpha
    v1alpha = __v1alpha
    import pulumi_google_native.metastore.v1beta as __v1beta
    v1beta = __v1beta
else:
    v1alpha = _utilities.lazy_import('pulumi_google_native.metastore.v1alpha')
    v1beta = _utilities.lazy_import('pulumi_google_native.metastore.v1beta')

