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
    'GetIndexResult',
    'AwaitableGetIndexResult',
    'get_index',
    'get_index_output',
]

@pulumi.output_type
class GetIndexResult:
    def __init__(__self__, ancestor=None, index_id=None, kind=None, project=None, properties=None, state=None):
        if ancestor and not isinstance(ancestor, str):
            raise TypeError("Expected argument 'ancestor' to be a str")
        pulumi.set(__self__, "ancestor", ancestor)
        if index_id and not isinstance(index_id, str):
            raise TypeError("Expected argument 'index_id' to be a str")
        pulumi.set(__self__, "index_id", index_id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if properties and not isinstance(properties, list):
            raise TypeError("Expected argument 'properties' to be a list")
        pulumi.set(__self__, "properties", properties)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def ancestor(self) -> str:
        """
        The index's ancestor mode. Must not be ANCESTOR_MODE_UNSPECIFIED.
        """
        return pulumi.get(self, "ancestor")

    @property
    @pulumi.getter(name="indexId")
    def index_id(self) -> str:
        """
        The resource ID of the index.
        """
        return pulumi.get(self, "index_id")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The entity kind to which this index applies.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        Project ID.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def properties(self) -> Sequence['outputs.GoogleDatastoreAdminV1IndexedPropertyResponse']:
        """
        An ordered sequence of property names and their index attributes.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the index.
        """
        return pulumi.get(self, "state")


class AwaitableGetIndexResult(GetIndexResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIndexResult(
            ancestor=self.ancestor,
            index_id=self.index_id,
            kind=self.kind,
            project=self.project,
            properties=self.properties,
            state=self.state)


def get_index(index_id: Optional[str] = None,
              project: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIndexResult:
    """
    Gets an index.
    """
    __args__ = dict()
    __args__['indexId'] = index_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:datastore/v1:getIndex', __args__, opts=opts, typ=GetIndexResult).value

    return AwaitableGetIndexResult(
        ancestor=__ret__.ancestor,
        index_id=__ret__.index_id,
        kind=__ret__.kind,
        project=__ret__.project,
        properties=__ret__.properties,
        state=__ret__.state)


@_utilities.lift_output_func(get_index)
def get_index_output(index_id: Optional[pulumi.Input[str]] = None,
                     project: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIndexResult]:
    """
    Gets an index.
    """
    ...
