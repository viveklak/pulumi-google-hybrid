# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['ProvisioningConfigArgs', 'ProvisioningConfig']

@pulumi.input_type
class ProvisioningConfigArgs:
    def __init__(__self__, *,
                 email: Optional[pulumi.Input[str]] = None,
                 handover_service_account: Optional[pulumi.Input[str]] = None,
                 instances: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceConfigArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkConfigArgs']]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ticket_id: Optional[pulumi.Input[str]] = None,
                 volumes: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeConfigArgs']]]] = None):
        """
        The set of arguments for constructing a ProvisioningConfig resource.
        :param pulumi.Input[str] email: Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
        :param pulumi.Input[str] handover_service_account: A service account to enable customers to access instance credentials upon handover.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceConfigArgs']]] instances: Instances to be created.
        :param pulumi.Input[str] location: Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkConfigArgs']]] networks: Networks to be created.
        :param pulumi.Input[str] ticket_id: A generated buganizer id to track provisioning request.
        :param pulumi.Input[Sequence[pulumi.Input['VolumeConfigArgs']]] volumes: Volumes to be created.
        """
        if email is not None:
            warnings.warn("""Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.""", DeprecationWarning)
            pulumi.log.warn("""email is deprecated: Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.""")
        if email is not None:
            pulumi.set(__self__, "email", email)
        if handover_service_account is not None:
            pulumi.set(__self__, "handover_service_account", handover_service_account)
        if instances is not None:
            pulumi.set(__self__, "instances", instances)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if ticket_id is not None:
            pulumi.set(__self__, "ticket_id", ticket_id)
        if volumes is not None:
            pulumi.set(__self__, "volumes", volumes)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="handoverServiceAccount")
    def handover_service_account(self) -> Optional[pulumi.Input[str]]:
        """
        A service account to enable customers to access instance credentials upon handover.
        """
        return pulumi.get(self, "handover_service_account")

    @handover_service_account.setter
    def handover_service_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "handover_service_account", value)

    @property
    @pulumi.getter
    def instances(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceConfigArgs']]]]:
        """
        Instances to be created.
        """
        return pulumi.get(self, "instances")

    @instances.setter
    def instances(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceConfigArgs']]]]):
        pulumi.set(self, "instances", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkConfigArgs']]]]:
        """
        Networks to be created.
        """
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkConfigArgs']]]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="ticketId")
    def ticket_id(self) -> Optional[pulumi.Input[str]]:
        """
        A generated buganizer id to track provisioning request.
        """
        return pulumi.get(self, "ticket_id")

    @ticket_id.setter
    def ticket_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ticket_id", value)

    @property
    @pulumi.getter
    def volumes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VolumeConfigArgs']]]]:
        """
        Volumes to be created.
        """
        return pulumi.get(self, "volumes")

    @volumes.setter
    def volumes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeConfigArgs']]]]):
        pulumi.set(self, "volumes", value)


class ProvisioningConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 handover_service_account: Optional[pulumi.Input[str]] = None,
                 instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceConfigArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConfigArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ticket_id: Optional[pulumi.Input[str]] = None,
                 volumes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeConfigArgs']]]]] = None,
                 __props__=None):
        """
        Create new ProvisioningConfig.
        Auto-naming is currently not supported for this resource.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
        :param pulumi.Input[str] handover_service_account: A service account to enable customers to access instance credentials upon handover.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceConfigArgs']]]] instances: Instances to be created.
        :param pulumi.Input[str] location: Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConfigArgs']]]] networks: Networks to be created.
        :param pulumi.Input[str] ticket_id: A generated buganizer id to track provisioning request.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeConfigArgs']]]] volumes: Volumes to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProvisioningConfigArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create new ProvisioningConfig.
        Auto-naming is currently not supported for this resource.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param ProvisioningConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProvisioningConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 handover_service_account: Optional[pulumi.Input[str]] = None,
                 instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceConfigArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConfigArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ticket_id: Optional[pulumi.Input[str]] = None,
                 volumes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeConfigArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProvisioningConfigArgs.__new__(ProvisioningConfigArgs)

            if email is not None and not opts.urn:
                warnings.warn("""Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.""", DeprecationWarning)
                pulumi.log.warn("""email is deprecated: Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.""")
            __props__.__dict__["email"] = email
            __props__.__dict__["handover_service_account"] = handover_service_account
            __props__.__dict__["instances"] = instances
            __props__.__dict__["location"] = location
            __props__.__dict__["networks"] = networks
            __props__.__dict__["project"] = project
            __props__.__dict__["ticket_id"] = ticket_id
            __props__.__dict__["volumes"] = volumes
            __props__.__dict__["cloud_console_uri"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
        super(ProvisioningConfig, __self__).__init__(
            'google-native:baremetalsolution/v2:ProvisioningConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ProvisioningConfig':
        """
        Get an existing ProvisioningConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProvisioningConfigArgs.__new__(ProvisioningConfigArgs)

        __props__.__dict__["cloud_console_uri"] = None
        __props__.__dict__["email"] = None
        __props__.__dict__["handover_service_account"] = None
        __props__.__dict__["instances"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["networks"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["ticket_id"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["volumes"] = None
        return ProvisioningConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudConsoleUri")
    def cloud_console_uri(self) -> pulumi.Output[str]:
        """
        URI to Cloud Console UI view of this provisioning config.
        """
        return pulumi.get(self, "cloud_console_uri")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="handoverServiceAccount")
    def handover_service_account(self) -> pulumi.Output[str]:
        """
        A service account to enable customers to access instance credentials upon handover.
        """
        return pulumi.get(self, "handover_service_account")

    @property
    @pulumi.getter
    def instances(self) -> pulumi.Output[Sequence['outputs.InstanceConfigResponse']]:
        """
        Instances to be created.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the provisioning config.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> pulumi.Output[Sequence['outputs.NetworkConfigResponse']]:
        """
        Networks to be created.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of ProvisioningConfig.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="ticketId")
    def ticket_id(self) -> pulumi.Output[str]:
        """
        A generated buganizer id to track provisioning request.
        """
        return pulumi.get(self, "ticket_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last update timestamp.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def volumes(self) -> pulumi.Output[Sequence['outputs.VolumeConfigResponse']]:
        """
        Volumes to be created.
        """
        return pulumi.get(self, "volumes")

