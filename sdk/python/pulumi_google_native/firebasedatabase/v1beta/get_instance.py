# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetInstanceResult',
    'AwaitableGetInstanceResult',
    'get_instance',
    'get_instance_output',
]

@pulumi.output_type
class GetInstanceResult:
    def __init__(__self__, database_url=None, name=None, project=None, state=None, type=None):
        if database_url and not isinstance(database_url, str):
            raise TypeError("Expected argument 'database_url' to be a str")
        pulumi.set(__self__, "database_url", database_url)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="databaseUrl")
    def database_url(self) -> str:
        """
        Immutable. The globally unique hostname of the database.
        """
        return pulumi.get(self, "database_url")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The fully qualified resource name of the database instance, in the form: `projects/{project-number}/locations/{location-id}/instances/{database-id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The resource name of the project this instance belongs to. For example: `projects/{project-number}`.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The database's lifecycle state. Read-only.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The database instance type. On creation only USER_DATABASE is allowed, which is also the default when omitted.
        """
        return pulumi.get(self, "type")


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            database_url=self.database_url,
            name=self.name,
            project=self.project,
            state=self.state,
            type=self.type)


def get_instance(instance_id: Optional[str] = None,
                 location: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceResult:
    """
    Gets the DatabaseInstance identified by the specified resource name.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:firebasedatabase/v1beta:getInstance', __args__, opts=opts, typ=GetInstanceResult).value

    return AwaitableGetInstanceResult(
        database_url=__ret__.database_url,
        name=__ret__.name,
        project=__ret__.project,
        state=__ret__.state,
        type=__ret__.type)


@_utilities.lift_output_func(get_instance)
def get_instance_output(instance_id: Optional[pulumi.Input[str]] = None,
                        location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceResult]:
    """
    Gets the DatabaseInstance identified by the specified resource name.
    """
    ...
