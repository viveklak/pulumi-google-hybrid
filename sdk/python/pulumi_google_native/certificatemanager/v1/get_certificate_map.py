# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetCertificateMapResult',
    'AwaitableGetCertificateMapResult',
    'get_certificate_map',
    'get_certificate_map_output',
]

@pulumi.output_type
class GetCertificateMapResult:
    def __init__(__self__, create_time=None, description=None, gclb_targets=None, labels=None, name=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if gclb_targets and not isinstance(gclb_targets, list):
            raise TypeError("Expected argument 'gclb_targets' to be a list")
        pulumi.set(__self__, "gclb_targets", gclb_targets)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation timestamp of a Certificate Map.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        One or more paragraphs of text description of a certificate map.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="gclbTargets")
    def gclb_targets(self) -> Sequence['outputs.GclbTargetResponse']:
        """
        A list of GCLB targets which use this Certificate Map.
        """
        return pulumi.get(self, "gclb_targets")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Set of labels associated with a Certificate Map.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A user-defined name of the Certificate Map. Certificate Map names must be unique globally and match pattern `projects/*/locations/*/certificateMaps/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The update timestamp of a Certificate Map.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetCertificateMapResult(GetCertificateMapResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificateMapResult(
            create_time=self.create_time,
            description=self.description,
            gclb_targets=self.gclb_targets,
            labels=self.labels,
            name=self.name,
            update_time=self.update_time)


def get_certificate_map(certificate_map_id: Optional[str] = None,
                        location: Optional[str] = None,
                        project: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCertificateMapResult:
    """
    Gets details of a single CertificateMap.
    """
    __args__ = dict()
    __args__['certificateMapId'] = certificate_map_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:certificatemanager/v1:getCertificateMap', __args__, opts=opts, typ=GetCertificateMapResult).value

    return AwaitableGetCertificateMapResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        gclb_targets=__ret__.gclb_targets,
        labels=__ret__.labels,
        name=__ret__.name,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_certificate_map)
def get_certificate_map_output(certificate_map_id: Optional[pulumi.Input[str]] = None,
                               location: Optional[pulumi.Input[str]] = None,
                               project: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCertificateMapResult]:
    """
    Gets details of a single CertificateMap.
    """
    ...
