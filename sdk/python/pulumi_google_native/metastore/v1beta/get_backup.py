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
    'GetBackupResult',
    'AwaitableGetBackupResult',
    'get_backup',
    'get_backup_output',
]

@pulumi.output_type
class GetBackupResult:
    def __init__(__self__, create_time=None, description=None, end_time=None, name=None, restoring_services=None, service_revision=None, state=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if restoring_services and not isinstance(restoring_services, list):
            raise TypeError("Expected argument 'restoring_services' to be a list")
        pulumi.set(__self__, "restoring_services", restoring_services)
        if service_revision and not isinstance(service_revision, dict):
            raise TypeError("Expected argument 'service_revision' to be a dict")
        pulumi.set(__self__, "service_revision", service_revision)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the backup was started.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the backup.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        The time when the backup finished creating.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. The relative resource name of the backup, in the following form:projects/{project_number}/locations/{location_id}/services/{service_id}/backups/{backup_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="restoringServices")
    def restoring_services(self) -> Sequence[str]:
        """
        Services that are restoring from the backup.
        """
        return pulumi.get(self, "restoring_services")

    @property
    @pulumi.getter(name="serviceRevision")
    def service_revision(self) -> 'outputs.ServiceResponse':
        """
        The revision of the service at the time of backup.
        """
        return pulumi.get(self, "service_revision")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the backup.
        """
        return pulumi.get(self, "state")


class AwaitableGetBackupResult(GetBackupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupResult(
            create_time=self.create_time,
            description=self.description,
            end_time=self.end_time,
            name=self.name,
            restoring_services=self.restoring_services,
            service_revision=self.service_revision,
            state=self.state)


def get_backup(backup_id: Optional[str] = None,
               location: Optional[str] = None,
               project: Optional[str] = None,
               service_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupResult:
    """
    Gets details of a single backup.
    """
    __args__ = dict()
    __args__['backupId'] = backup_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['serviceId'] = service_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:metastore/v1beta:getBackup', __args__, opts=opts, typ=GetBackupResult).value

    return AwaitableGetBackupResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        end_time=__ret__.end_time,
        name=__ret__.name,
        restoring_services=__ret__.restoring_services,
        service_revision=__ret__.service_revision,
        state=__ret__.state)


@_utilities.lift_output_func(get_backup)
def get_backup_output(backup_id: Optional[pulumi.Input[str]] = None,
                      location: Optional[pulumi.Input[str]] = None,
                      project: Optional[pulumi.Input[Optional[str]]] = None,
                      service_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupResult]:
    """
    Gets details of a single backup.
    """
    ...
