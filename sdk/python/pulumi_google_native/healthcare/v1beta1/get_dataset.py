# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetDatasetResult',
    'AwaitableGetDatasetResult',
    'get_dataset',
    'get_dataset_output',
]

@pulumi.output_type
class GetDatasetResult:
    def __init__(__self__, name=None, time_zone=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if time_zone and not isinstance(time_zone, str):
            raise TypeError("Expected argument 'time_zone' to be a str")
        pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> str:
        """
        The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7 messages, where no explicit timezone is specified.
        """
        return pulumi.get(self, "time_zone")


class AwaitableGetDatasetResult(GetDatasetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatasetResult(
            name=self.name,
            time_zone=self.time_zone)


def get_dataset(dataset_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatasetResult:
    """
    Gets any metadata associated with a dataset.
    """
    __args__ = dict()
    __args__['datasetId'] = dataset_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:healthcare/v1beta1:getDataset', __args__, opts=opts, typ=GetDatasetResult).value

    return AwaitableGetDatasetResult(
        name=__ret__.name,
        time_zone=__ret__.time_zone)


@_utilities.lift_output_func(get_dataset)
def get_dataset_output(dataset_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatasetResult]:
    """
    Gets any metadata associated with a dataset.
    """
    ...
