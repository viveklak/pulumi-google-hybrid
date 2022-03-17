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
    'GetProvisioningConfigResult',
    'AwaitableGetProvisioningConfigResult',
    'get_provisioning_config',
    'get_provisioning_config_output',
]

@pulumi.output_type
class GetProvisioningConfigResult:
    def __init__(__self__, cloud_console_uri=None, email=None, handover_service_account=None, instances=None, location=None, name=None, networks=None, state=None, ticket_id=None, update_time=None, volumes=None):
        if cloud_console_uri and not isinstance(cloud_console_uri, str):
            raise TypeError("Expected argument 'cloud_console_uri' to be a str")
        pulumi.set(__self__, "cloud_console_uri", cloud_console_uri)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if handover_service_account and not isinstance(handover_service_account, str):
            raise TypeError("Expected argument 'handover_service_account' to be a str")
        pulumi.set(__self__, "handover_service_account", handover_service_account)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if ticket_id and not isinstance(ticket_id, str):
            raise TypeError("Expected argument 'ticket_id' to be a str")
        pulumi.set(__self__, "ticket_id", ticket_id)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if volumes and not isinstance(volumes, list):
            raise TypeError("Expected argument 'volumes' to be a list")
        pulumi.set(__self__, "volumes", volumes)

    @property
    @pulumi.getter(name="cloudConsoleUri")
    def cloud_console_uri(self) -> str:
        """
        URI to Cloud Console UI view of this provisioning config.
        """
        return pulumi.get(self, "cloud_console_uri")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        Optional. Email provided to send a confirmation with provisioning config to.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="handoverServiceAccount")
    def handover_service_account(self) -> str:
        """
        A service account to enable customers to access instance credentials upon handover.
        """
        return pulumi.get(self, "handover_service_account")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.InstanceConfigResponse']:
        """
        Instances to be created.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the provisioning config.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.NetworkConfigResponse']:
        """
        Networks to be created.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of ProvisioningConfig.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="ticketId")
    def ticket_id(self) -> str:
        """
        A generated buganizer id to track provisioning request.
        """
        return pulumi.get(self, "ticket_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Last update timestamp.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def volumes(self) -> Sequence['outputs.VolumeConfigResponse']:
        """
        Volumes to be created.
        """
        return pulumi.get(self, "volumes")


class AwaitableGetProvisioningConfigResult(GetProvisioningConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProvisioningConfigResult(
            cloud_console_uri=self.cloud_console_uri,
            email=self.email,
            handover_service_account=self.handover_service_account,
            instances=self.instances,
            location=self.location,
            name=self.name,
            networks=self.networks,
            state=self.state,
            ticket_id=self.ticket_id,
            update_time=self.update_time,
            volumes=self.volumes)


def get_provisioning_config(location: Optional[str] = None,
                            project: Optional[str] = None,
                            provisioning_config_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProvisioningConfigResult:
    """
    Get ProvisioningConfig by name.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['provisioningConfigId'] = provisioning_config_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:baremetalsolution/v2:getProvisioningConfig', __args__, opts=opts, typ=GetProvisioningConfigResult).value

    return AwaitableGetProvisioningConfigResult(
        cloud_console_uri=__ret__.cloud_console_uri,
        email=__ret__.email,
        handover_service_account=__ret__.handover_service_account,
        instances=__ret__.instances,
        location=__ret__.location,
        name=__ret__.name,
        networks=__ret__.networks,
        state=__ret__.state,
        ticket_id=__ret__.ticket_id,
        update_time=__ret__.update_time,
        volumes=__ret__.volumes)


@_utilities.lift_output_func(get_provisioning_config)
def get_provisioning_config_output(location: Optional[pulumi.Input[str]] = None,
                                   project: Optional[pulumi.Input[Optional[str]]] = None,
                                   provisioning_config_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProvisioningConfigResult]:
    """
    Get ProvisioningConfig by name.
    """
    ...
