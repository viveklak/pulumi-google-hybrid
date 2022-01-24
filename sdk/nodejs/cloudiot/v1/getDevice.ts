// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details about a device.
 */
export function getDevice(args: GetDeviceArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:cloudiot/v1:getDevice", {
        "deviceId": args.deviceId,
        "fieldMask": args.fieldMask,
        "location": args.location,
        "project": args.project,
        "registryId": args.registryId,
    }, opts);
}

export interface GetDeviceArgs {
    deviceId: string;
    fieldMask?: string;
    location: string;
    project?: string;
    registryId: string;
}

export interface GetDeviceResult {
    /**
     * If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
     */
    readonly blocked: boolean;
    /**
     * The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
     */
    readonly config: outputs.cloudiot.v1.DeviceConfigResponse;
    /**
     * The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
     */
    readonly credentials: outputs.cloudiot.v1.DeviceCredentialResponse[];
    /**
     * Gateway-related configuration and state.
     */
    readonly gatewayConfig: outputs.cloudiot.v1.GatewayConfigResponse;
    /**
     * [Output only] The last time a cloud-to-device config version acknowledgment was received from the device. This field is only for configurations sent through MQTT.
     */
    readonly lastConfigAckTime: string;
    /**
     * [Output only] The last time a cloud-to-device config version was sent to the device.
     */
    readonly lastConfigSendTime: string;
    /**
     * [Output only] The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub. 'last_error_time' is the timestamp of this field. If no errors have occurred, this field has an empty message and the status code 0 == OK. Otherwise, this field is expected to have a status code other than OK.
     */
    readonly lastErrorStatus: outputs.cloudiot.v1.StatusResponse;
    /**
     * [Output only] The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub. This field is the timestamp of 'last_error_status'.
     */
    readonly lastErrorTime: string;
    /**
     * [Output only] The last time a telemetry event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
     */
    readonly lastEventTime: string;
    /**
     * [Output only] The last time an MQTT `PINGREQ` was received. This field applies only to devices connecting through MQTT. MQTT clients usually only send `PINGREQ` messages if the connection is idle, and no other messages have been sent. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
     */
    readonly lastHeartbeatTime: string;
    /**
     * [Output only] The last time a state event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
     */
    readonly lastStateTime: string;
    /**
     * **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
     */
    readonly logLevel: string;
    /**
     * The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
     */
    readonly metadata: {[key: string]: string};
    /**
     * The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
     */
    readonly name: string;
    /**
     * [Output only] A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally unique.
     */
    readonly numId: string;
    /**
     * [Output only] The state most recently received from the device. If no state has been reported, this field is not present.
     */
    readonly state: outputs.cloudiot.v1.DeviceStateResponse;
}

export function getDeviceOutput(args: GetDeviceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeviceResult> {
    return pulumi.output(args).apply(a => getDevice(a, opts))
}

export interface GetDeviceOutputArgs {
    deviceId: pulumi.Input<string>;
    fieldMask?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    registryId: pulumi.Input<string>;
}
