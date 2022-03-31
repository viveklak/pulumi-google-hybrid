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

__all__ = ['DeviceArgs', 'Device']

@pulumi.input_type
class DeviceArgs:
    def __init__(__self__, *,
                 registry_id: pulumi.Input[str],
                 blocked: Optional[pulumi.Input[bool]] = None,
                 config: Optional[pulumi.Input['DeviceConfigArgs']] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input['DeviceCredentialArgs']]]] = None,
                 gateway_config: Optional[pulumi.Input['GatewayConfigArgs']] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input['DeviceLogLevel']] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Device resource.
        :param pulumi.Input[bool] blocked: If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
        :param pulumi.Input['DeviceConfigArgs'] config: The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
        :param pulumi.Input[Sequence[pulumi.Input['DeviceCredentialArgs']]] credentials: The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
        :param pulumi.Input['GatewayConfigArgs'] gateway_config: Gateway-related configuration and state.
        :param pulumi.Input[str] id: The user-defined device identifier. The device ID must be unique within a device registry.
        :param pulumi.Input['DeviceLogLevel'] log_level: **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
        :param pulumi.Input[str] name: The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
        """
        pulumi.set(__self__, "registry_id", registry_id)
        if blocked is not None:
            pulumi.set(__self__, "blocked", blocked)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if gateway_config is not None:
            pulumi.set(__self__, "gateway_config", gateway_config)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log_level is not None:
            pulumi.set(__self__, "log_level", log_level)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter
    def blocked(self) -> Optional[pulumi.Input[bool]]:
        """
        If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
        """
        return pulumi.get(self, "blocked")

    @blocked.setter
    def blocked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "blocked", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['DeviceConfigArgs']]:
        """
        The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['DeviceConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeviceCredentialArgs']]]]:
        """
        The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeviceCredentialArgs']]]]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter(name="gatewayConfig")
    def gateway_config(self) -> Optional[pulumi.Input['GatewayConfigArgs']]:
        """
        Gateway-related configuration and state.
        """
        return pulumi.get(self, "gateway_config")

    @gateway_config.setter
    def gateway_config(self, value: Optional[pulumi.Input['GatewayConfigArgs']]):
        pulumi.set(self, "gateway_config", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The user-defined device identifier. The device ID must be unique within a device registry.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> Optional[pulumi.Input['DeviceLogLevel']]:
        """
        **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
        """
        return pulumi.get(self, "log_level")

    @log_level.setter
    def log_level(self, value: Optional[pulumi.Input['DeviceLogLevel']]):
        pulumi.set(self, "log_level", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Device(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocked: Optional[pulumi.Input[bool]] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['DeviceConfigArgs']]] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DeviceCredentialArgs']]]]] = None,
                 gateway_config: Optional[pulumi.Input[pulumi.InputType['GatewayConfigArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input['DeviceLogLevel']] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a device in a device registry.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] blocked: If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
        :param pulumi.Input[pulumi.InputType['DeviceConfigArgs']] config: The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DeviceCredentialArgs']]]] credentials: The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
        :param pulumi.Input[pulumi.InputType['GatewayConfigArgs']] gateway_config: Gateway-related configuration and state.
        :param pulumi.Input[str] id: The user-defined device identifier. The device ID must be unique within a device registry.
        :param pulumi.Input['DeviceLogLevel'] log_level: **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
        :param pulumi.Input[str] name: The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a device in a device registry.

        :param str resource_name: The name of the resource.
        :param DeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocked: Optional[pulumi.Input[bool]] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['DeviceConfigArgs']]] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DeviceCredentialArgs']]]]] = None,
                 gateway_config: Optional[pulumi.Input[pulumi.InputType['GatewayConfigArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input['DeviceLogLevel']] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = DeviceArgs.__new__(DeviceArgs)

            __props__.__dict__["blocked"] = blocked
            __props__.__dict__["config"] = config
            __props__.__dict__["credentials"] = credentials
            __props__.__dict__["gateway_config"] = gateway_config
            __props__.__dict__["id"] = id
            __props__.__dict__["location"] = location
            __props__.__dict__["log_level"] = log_level
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if registry_id is None and not opts.urn:
                raise TypeError("Missing required property 'registry_id'")
            __props__.__dict__["registry_id"] = registry_id
            __props__.__dict__["last_config_ack_time"] = None
            __props__.__dict__["last_config_send_time"] = None
            __props__.__dict__["last_error_status"] = None
            __props__.__dict__["last_error_time"] = None
            __props__.__dict__["last_event_time"] = None
            __props__.__dict__["last_heartbeat_time"] = None
            __props__.__dict__["last_state_time"] = None
            __props__.__dict__["num_id"] = None
            __props__.__dict__["state"] = None
        super(Device, __self__).__init__(
            'google-native:cloudiot/v1:Device',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Device':
        """
        Get an existing Device resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DeviceArgs.__new__(DeviceArgs)

        __props__.__dict__["blocked"] = None
        __props__.__dict__["config"] = None
        __props__.__dict__["credentials"] = None
        __props__.__dict__["gateway_config"] = None
        __props__.__dict__["last_config_ack_time"] = None
        __props__.__dict__["last_config_send_time"] = None
        __props__.__dict__["last_error_status"] = None
        __props__.__dict__["last_error_time"] = None
        __props__.__dict__["last_event_time"] = None
        __props__.__dict__["last_heartbeat_time"] = None
        __props__.__dict__["last_state_time"] = None
        __props__.__dict__["log_level"] = None
        __props__.__dict__["metadata"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["num_id"] = None
        __props__.__dict__["state"] = None
        return Device(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def blocked(self) -> pulumi.Output[bool]:
        """
        If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
        """
        return pulumi.get(self, "blocked")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.DeviceConfigResponse']:
        """
        The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Sequence['outputs.DeviceCredentialResponse']]:
        """
        The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter(name="gatewayConfig")
    def gateway_config(self) -> pulumi.Output['outputs.GatewayConfigResponse']:
        """
        Gateway-related configuration and state.
        """
        return pulumi.get(self, "gateway_config")

    @property
    @pulumi.getter(name="lastConfigAckTime")
    def last_config_ack_time(self) -> pulumi.Output[str]:
        """
        [Output only] The last time a cloud-to-device config version acknowledgment was received from the device. This field is only for configurations sent through MQTT.
        """
        return pulumi.get(self, "last_config_ack_time")

    @property
    @pulumi.getter(name="lastConfigSendTime")
    def last_config_send_time(self) -> pulumi.Output[str]:
        """
        [Output only] The last time a cloud-to-device config version was sent to the device.
        """
        return pulumi.get(self, "last_config_send_time")

    @property
    @pulumi.getter(name="lastErrorStatus")
    def last_error_status(self) -> pulumi.Output['outputs.StatusResponse']:
        """
        [Output only] The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub. 'last_error_time' is the timestamp of this field. If no errors have occurred, this field has an empty message and the status code 0 == OK. Otherwise, this field is expected to have a status code other than OK.
        """
        return pulumi.get(self, "last_error_status")

    @property
    @pulumi.getter(name="lastErrorTime")
    def last_error_time(self) -> pulumi.Output[str]:
        """
        [Output only] The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub. This field is the timestamp of 'last_error_status'.
        """
        return pulumi.get(self, "last_error_time")

    @property
    @pulumi.getter(name="lastEventTime")
    def last_event_time(self) -> pulumi.Output[str]:
        """
        [Output only] The last time a telemetry event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
        """
        return pulumi.get(self, "last_event_time")

    @property
    @pulumi.getter(name="lastHeartbeatTime")
    def last_heartbeat_time(self) -> pulumi.Output[str]:
        """
        [Output only] The last time an MQTT `PINGREQ` was received. This field applies only to devices connecting through MQTT. MQTT clients usually only send `PINGREQ` messages if the connection is idle, and no other messages have been sent. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
        """
        return pulumi.get(self, "last_heartbeat_time")

    @property
    @pulumi.getter(name="lastStateTime")
    def last_state_time(self) -> pulumi.Output[str]:
        """
        [Output only] The last time a state event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
        """
        return pulumi.get(self, "last_state_time")

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> pulumi.Output[str]:
        """
        **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
        """
        return pulumi.get(self, "log_level")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numId")
    def num_id(self) -> pulumi.Output[str]:
        """
        [Output only] A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally unique.
        """
        return pulumi.get(self, "num_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output['outputs.DeviceStateResponse']:
        """
        [Output only] The state most recently received from the device. If no state has been reported, this field is not present.
        """
        return pulumi.get(self, "state")

