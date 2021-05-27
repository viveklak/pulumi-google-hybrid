// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a device in a device registry.
type RegistryDevice struct {
	pulumi.CustomResourceState

	// If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
	Blocked pulumi.BoolOutput `pulumi:"blocked"`
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
	Config DeviceConfigResponseOutput `pulumi:"config"`
	// The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
	Credentials DeviceCredentialResponseArrayOutput `pulumi:"credentials"`
	// Gateway-related configuration and state.
	GatewayConfig GatewayConfigResponseOutput `pulumi:"gatewayConfig"`
	// [Output only] The last time a cloud-to-device config version acknowledgment was received from the device. This field is only for configurations sent through MQTT.
	LastConfigAckTime pulumi.StringOutput `pulumi:"lastConfigAckTime"`
	// [Output only] The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime pulumi.StringOutput `pulumi:"lastConfigSendTime"`
	// [Output only] The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub. 'last_error_time' is the timestamp of this field. If no errors have occurred, this field has an empty message and the status code 0 == OK. Otherwise, this field is expected to have a status code other than OK.
	LastErrorStatus StatusResponseOutput `pulumi:"lastErrorStatus"`
	// [Output only] The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub. This field is the timestamp of 'last_error_status'.
	LastErrorTime pulumi.StringOutput `pulumi:"lastErrorTime"`
	// [Output only] The last time a telemetry event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastEventTime pulumi.StringOutput `pulumi:"lastEventTime"`
	// [Output only] The last time an MQTT `PINGREQ` was received. This field applies only to devices connecting through MQTT. MQTT clients usually only send `PINGREQ` messages if the connection is idle, and no other messages have been sent. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastHeartbeatTime pulumi.StringOutput `pulumi:"lastHeartbeatTime"`
	// [Output only] The last time a state event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastStateTime pulumi.StringOutput `pulumi:"lastStateTime"`
	// **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
	LogLevel pulumi.StringOutput `pulumi:"logLevel"`
	// The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
	Name pulumi.StringOutput `pulumi:"name"`
	// [Output only] A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally unique.
	NumId pulumi.StringOutput `pulumi:"numId"`
	// [Output only] The state most recently received from the device. If no state has been reported, this field is not present.
	State DeviceStateResponseOutput `pulumi:"state"`
}

// NewRegistryDevice registers a new resource with the given unique name, arguments, and options.
func NewRegistryDevice(ctx *pulumi.Context,
	name string, args *RegistryDeviceArgs, opts ...pulumi.ResourceOption) (*RegistryDevice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	var resource RegistryDevice
	err := ctx.RegisterResource("google-native:cloudiot/v1:RegistryDevice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryDevice gets an existing RegistryDevice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryDeviceState, opts ...pulumi.ResourceOption) (*RegistryDevice, error) {
	var resource RegistryDevice
	err := ctx.ReadResource("google-native:cloudiot/v1:RegistryDevice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryDevice resources.
type registryDeviceState struct {
	// If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
	Blocked *bool `pulumi:"blocked"`
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
	Config *DeviceConfigResponse `pulumi:"config"`
	// The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
	Credentials []DeviceCredentialResponse `pulumi:"credentials"`
	// Gateway-related configuration and state.
	GatewayConfig *GatewayConfigResponse `pulumi:"gatewayConfig"`
	// [Output only] The last time a cloud-to-device config version acknowledgment was received from the device. This field is only for configurations sent through MQTT.
	LastConfigAckTime *string `pulumi:"lastConfigAckTime"`
	// [Output only] The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime *string `pulumi:"lastConfigSendTime"`
	// [Output only] The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub. 'last_error_time' is the timestamp of this field. If no errors have occurred, this field has an empty message and the status code 0 == OK. Otherwise, this field is expected to have a status code other than OK.
	LastErrorStatus *StatusResponse `pulumi:"lastErrorStatus"`
	// [Output only] The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub. This field is the timestamp of 'last_error_status'.
	LastErrorTime *string `pulumi:"lastErrorTime"`
	// [Output only] The last time a telemetry event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastEventTime *string `pulumi:"lastEventTime"`
	// [Output only] The last time an MQTT `PINGREQ` was received. This field applies only to devices connecting through MQTT. MQTT clients usually only send `PINGREQ` messages if the connection is idle, and no other messages have been sent. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastHeartbeatTime *string `pulumi:"lastHeartbeatTime"`
	// [Output only] The last time a state event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastStateTime *string `pulumi:"lastStateTime"`
	// **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
	LogLevel *string `pulumi:"logLevel"`
	// The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
	Metadata map[string]string `pulumi:"metadata"`
	// The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
	Name *string `pulumi:"name"`
	// [Output only] A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally unique.
	NumId *string `pulumi:"numId"`
	// [Output only] The state most recently received from the device. If no state has been reported, this field is not present.
	State *DeviceStateResponse `pulumi:"state"`
}

type RegistryDeviceState struct {
	// If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
	Blocked pulumi.BoolPtrInput
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
	Config DeviceConfigResponsePtrInput
	// The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
	Credentials DeviceCredentialResponseArrayInput
	// Gateway-related configuration and state.
	GatewayConfig GatewayConfigResponsePtrInput
	// [Output only] The last time a cloud-to-device config version acknowledgment was received from the device. This field is only for configurations sent through MQTT.
	LastConfigAckTime pulumi.StringPtrInput
	// [Output only] The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime pulumi.StringPtrInput
	// [Output only] The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub. 'last_error_time' is the timestamp of this field. If no errors have occurred, this field has an empty message and the status code 0 == OK. Otherwise, this field is expected to have a status code other than OK.
	LastErrorStatus StatusResponsePtrInput
	// [Output only] The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub. This field is the timestamp of 'last_error_status'.
	LastErrorTime pulumi.StringPtrInput
	// [Output only] The last time a telemetry event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastEventTime pulumi.StringPtrInput
	// [Output only] The last time an MQTT `PINGREQ` was received. This field applies only to devices connecting through MQTT. MQTT clients usually only send `PINGREQ` messages if the connection is idle, and no other messages have been sent. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastHeartbeatTime pulumi.StringPtrInput
	// [Output only] The last time a state event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastStateTime pulumi.StringPtrInput
	// **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
	LogLevel pulumi.StringPtrInput
	// The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
	Metadata pulumi.StringMapInput
	// The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
	Name pulumi.StringPtrInput
	// [Output only] A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally unique.
	NumId pulumi.StringPtrInput
	// [Output only] The state most recently received from the device. If no state has been reported, this field is not present.
	State DeviceStateResponsePtrInput
}

func (RegistryDeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryDeviceState)(nil)).Elem()
}

type registryDeviceArgs struct {
	// If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
	Blocked *bool `pulumi:"blocked"`
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
	Config *DeviceConfig `pulumi:"config"`
	// The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
	Credentials []DeviceCredential `pulumi:"credentials"`
	// Gateway-related configuration and state.
	GatewayConfig *GatewayConfig `pulumi:"gatewayConfig"`
	// The user-defined device identifier. The device ID must be unique within a device registry.
	Id *string `pulumi:"id"`
	// [Output only] The last time a cloud-to-device config version acknowledgment was received from the device. This field is only for configurations sent through MQTT.
	LastConfigAckTime *string `pulumi:"lastConfigAckTime"`
	// [Output only] The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime *string `pulumi:"lastConfigSendTime"`
	// [Output only] The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub. 'last_error_time' is the timestamp of this field. If no errors have occurred, this field has an empty message and the status code 0 == OK. Otherwise, this field is expected to have a status code other than OK.
	LastErrorStatus *Status `pulumi:"lastErrorStatus"`
	// [Output only] The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub. This field is the timestamp of 'last_error_status'.
	LastErrorTime *string `pulumi:"lastErrorTime"`
	// [Output only] The last time a telemetry event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastEventTime *string `pulumi:"lastEventTime"`
	// [Output only] The last time an MQTT `PINGREQ` was received. This field applies only to devices connecting through MQTT. MQTT clients usually only send `PINGREQ` messages if the connection is idle, and no other messages have been sent. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastHeartbeatTime *string `pulumi:"lastHeartbeatTime"`
	// [Output only] The last time a state event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastStateTime *string `pulumi:"lastStateTime"`
	Location      string  `pulumi:"location"`
	// **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
	LogLevel *string `pulumi:"logLevel"`
	// The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
	Metadata map[string]string `pulumi:"metadata"`
	// The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
	Name *string `pulumi:"name"`
	// [Output only] A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally unique.
	NumId      *string `pulumi:"numId"`
	Project    string  `pulumi:"project"`
	RegistryId string  `pulumi:"registryId"`
	// [Output only] The state most recently received from the device. If no state has been reported, this field is not present.
	State *DeviceState `pulumi:"state"`
}

// The set of arguments for constructing a RegistryDevice resource.
type RegistryDeviceArgs struct {
	// If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
	Blocked pulumi.BoolPtrInput
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
	Config DeviceConfigPtrInput
	// The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
	Credentials DeviceCredentialArrayInput
	// Gateway-related configuration and state.
	GatewayConfig GatewayConfigPtrInput
	// The user-defined device identifier. The device ID must be unique within a device registry.
	Id pulumi.StringPtrInput
	// [Output only] The last time a cloud-to-device config version acknowledgment was received from the device. This field is only for configurations sent through MQTT.
	LastConfigAckTime pulumi.StringPtrInput
	// [Output only] The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime pulumi.StringPtrInput
	// [Output only] The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub. 'last_error_time' is the timestamp of this field. If no errors have occurred, this field has an empty message and the status code 0 == OK. Otherwise, this field is expected to have a status code other than OK.
	LastErrorStatus StatusPtrInput
	// [Output only] The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub. This field is the timestamp of 'last_error_status'.
	LastErrorTime pulumi.StringPtrInput
	// [Output only] The last time a telemetry event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastEventTime pulumi.StringPtrInput
	// [Output only] The last time an MQTT `PINGREQ` was received. This field applies only to devices connecting through MQTT. MQTT clients usually only send `PINGREQ` messages if the connection is idle, and no other messages have been sent. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastHeartbeatTime pulumi.StringPtrInput
	// [Output only] The last time a state event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
	LastStateTime pulumi.StringPtrInput
	Location      pulumi.StringInput
	// **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
	LogLevel pulumi.StringPtrInput
	// The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
	Metadata pulumi.StringMapInput
	// The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
	Name pulumi.StringPtrInput
	// [Output only] A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally unique.
	NumId      pulumi.StringPtrInput
	Project    pulumi.StringInput
	RegistryId pulumi.StringInput
	// [Output only] The state most recently received from the device. If no state has been reported, this field is not present.
	State DeviceStatePtrInput
}

func (RegistryDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryDeviceArgs)(nil)).Elem()
}

type RegistryDeviceInput interface {
	pulumi.Input

	ToRegistryDeviceOutput() RegistryDeviceOutput
	ToRegistryDeviceOutputWithContext(ctx context.Context) RegistryDeviceOutput
}

func (*RegistryDevice) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryDevice)(nil))
}

func (i *RegistryDevice) ToRegistryDeviceOutput() RegistryDeviceOutput {
	return i.ToRegistryDeviceOutputWithContext(context.Background())
}

func (i *RegistryDevice) ToRegistryDeviceOutputWithContext(ctx context.Context) RegistryDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryDeviceOutput)
}

type RegistryDeviceOutput struct {
	*pulumi.OutputState
}

func (RegistryDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryDevice)(nil))
}

func (o RegistryDeviceOutput) ToRegistryDeviceOutput() RegistryDeviceOutput {
	return o
}

func (o RegistryDeviceOutput) ToRegistryDeviceOutputWithContext(ctx context.Context) RegistryDeviceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegistryDeviceOutput{})
}
