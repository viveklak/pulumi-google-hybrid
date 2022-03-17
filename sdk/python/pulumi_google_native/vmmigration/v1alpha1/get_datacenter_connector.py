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
    'GetDatacenterConnectorResult',
    'AwaitableGetDatacenterConnectorResult',
    'get_datacenter_connector',
    'get_datacenter_connector_output',
]

@pulumi.output_type
class GetDatacenterConnectorResult:
    def __init__(__self__, appliance_infrastructure_version=None, appliance_software_version=None, available_versions=None, bucket=None, create_time=None, error=None, name=None, registration_id=None, service_account=None, state=None, state_time=None, update_time=None, upgrade_status=None, version=None):
        if appliance_infrastructure_version and not isinstance(appliance_infrastructure_version, str):
            raise TypeError("Expected argument 'appliance_infrastructure_version' to be a str")
        pulumi.set(__self__, "appliance_infrastructure_version", appliance_infrastructure_version)
        if appliance_software_version and not isinstance(appliance_software_version, str):
            raise TypeError("Expected argument 'appliance_software_version' to be a str")
        pulumi.set(__self__, "appliance_software_version", appliance_software_version)
        if available_versions and not isinstance(available_versions, dict):
            raise TypeError("Expected argument 'available_versions' to be a dict")
        pulumi.set(__self__, "available_versions", available_versions)
        if bucket and not isinstance(bucket, str):
            raise TypeError("Expected argument 'bucket' to be a str")
        pulumi.set(__self__, "bucket", bucket)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if error and not isinstance(error, dict):
            raise TypeError("Expected argument 'error' to be a dict")
        pulumi.set(__self__, "error", error)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if registration_id and not isinstance(registration_id, str):
            raise TypeError("Expected argument 'registration_id' to be a str")
        pulumi.set(__self__, "registration_id", registration_id)
        if service_account and not isinstance(service_account, str):
            raise TypeError("Expected argument 'service_account' to be a str")
        pulumi.set(__self__, "service_account", service_account)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if state_time and not isinstance(state_time, str):
            raise TypeError("Expected argument 'state_time' to be a str")
        pulumi.set(__self__, "state_time", state_time)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if upgrade_status and not isinstance(upgrade_status, dict):
            raise TypeError("Expected argument 'upgrade_status' to be a dict")
        pulumi.set(__self__, "upgrade_status", upgrade_status)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="applianceInfrastructureVersion")
    def appliance_infrastructure_version(self) -> str:
        """
        Appliance OVA version. This is the OVA which is manually installed by the user and contains the infrastructure for the automatically updatable components on the appliance.
        """
        return pulumi.get(self, "appliance_infrastructure_version")

    @property
    @pulumi.getter(name="applianceSoftwareVersion")
    def appliance_software_version(self) -> str:
        """
        Appliance last installed update bundle version. This is the version of the automatically updatable components on the appliance.
        """
        return pulumi.get(self, "appliance_software_version")

    @property
    @pulumi.getter(name="availableVersions")
    def available_versions(self) -> 'outputs.AvailableUpdatesResponse':
        """
        The available versions for updating this appliance.
        """
        return pulumi.get(self, "available_versions")

    @property
    @pulumi.getter
    def bucket(self) -> str:
        """
        The communication channel between the datacenter connector and GCP.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time the connector was created (as an API call, not when it was actually installed).
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def error(self) -> 'outputs.StatusResponse':
        """
        Provides details on the state of the Datacenter Connector in case of an error.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The connector's name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="registrationId")
    def registration_id(self) -> str:
        """
        Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
        """
        return pulumi.get(self, "registration_id")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        """
        The service account to use in the connector when communicating with the cloud.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the DatacenterConnector, as determined by the health checks.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateTime")
    def state_time(self) -> str:
        """
        The time the state was last set.
        """
        return pulumi.get(self, "state_time")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last time the connector was updated with an API call.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="upgradeStatus")
    def upgrade_status(self) -> 'outputs.UpgradeStatusResponse':
        """
        The status of the current / last upgradeAppliance operation.
        """
        return pulumi.get(self, "upgrade_status")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
        """
        return pulumi.get(self, "version")


class AwaitableGetDatacenterConnectorResult(GetDatacenterConnectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatacenterConnectorResult(
            appliance_infrastructure_version=self.appliance_infrastructure_version,
            appliance_software_version=self.appliance_software_version,
            available_versions=self.available_versions,
            bucket=self.bucket,
            create_time=self.create_time,
            error=self.error,
            name=self.name,
            registration_id=self.registration_id,
            service_account=self.service_account,
            state=self.state,
            state_time=self.state_time,
            update_time=self.update_time,
            upgrade_status=self.upgrade_status,
            version=self.version)


def get_datacenter_connector(datacenter_connector_id: Optional[str] = None,
                             location: Optional[str] = None,
                             project: Optional[str] = None,
                             source_id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatacenterConnectorResult:
    """
    Gets details of a single DatacenterConnector.
    """
    __args__ = dict()
    __args__['datacenterConnectorId'] = datacenter_connector_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['sourceId'] = source_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:vmmigration/v1alpha1:getDatacenterConnector', __args__, opts=opts, typ=GetDatacenterConnectorResult).value

    return AwaitableGetDatacenterConnectorResult(
        appliance_infrastructure_version=__ret__.appliance_infrastructure_version,
        appliance_software_version=__ret__.appliance_software_version,
        available_versions=__ret__.available_versions,
        bucket=__ret__.bucket,
        create_time=__ret__.create_time,
        error=__ret__.error,
        name=__ret__.name,
        registration_id=__ret__.registration_id,
        service_account=__ret__.service_account,
        state=__ret__.state,
        state_time=__ret__.state_time,
        update_time=__ret__.update_time,
        upgrade_status=__ret__.upgrade_status,
        version=__ret__.version)


@_utilities.lift_output_func(get_datacenter_connector)
def get_datacenter_connector_output(datacenter_connector_id: Optional[pulumi.Input[str]] = None,
                                    location: Optional[pulumi.Input[str]] = None,
                                    project: Optional[pulumi.Input[Optional[str]]] = None,
                                    source_id: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatacenterConnectorResult]:
    """
    Gets details of a single DatacenterConnector.
    """
    ...
