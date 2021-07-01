// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a device in a device registry.
 */
export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudiot/v1:Device';

    /**
     * Returns true if the given object is an instance of Device.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Device {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Device.__pulumiType;
    }

    /**
     * If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
     */
    public readonly blocked!: pulumi.Output<boolean>;
    /**
     * The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
     */
    public readonly config!: pulumi.Output<outputs.cloudiot.v1.DeviceConfigResponse>;
    /**
     * The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
     */
    public readonly credentials!: pulumi.Output<outputs.cloudiot.v1.DeviceCredentialResponse[]>;
    /**
     * Gateway-related configuration and state.
     */
    public readonly gatewayConfig!: pulumi.Output<outputs.cloudiot.v1.GatewayConfigResponse>;
    /**
     * [Output only] The last time a cloud-to-device config version acknowledgment was received from the device. This field is only for configurations sent through MQTT.
     */
    public /*out*/ readonly lastConfigAckTime!: pulumi.Output<string>;
    /**
     * [Output only] The last time a cloud-to-device config version was sent to the device.
     */
    public /*out*/ readonly lastConfigSendTime!: pulumi.Output<string>;
    /**
     * [Output only] The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub. 'last_error_time' is the timestamp of this field. If no errors have occurred, this field has an empty message and the status code 0 == OK. Otherwise, this field is expected to have a status code other than OK.
     */
    public /*out*/ readonly lastErrorStatus!: pulumi.Output<outputs.cloudiot.v1.StatusResponse>;
    /**
     * [Output only] The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub. This field is the timestamp of 'last_error_status'.
     */
    public /*out*/ readonly lastErrorTime!: pulumi.Output<string>;
    /**
     * [Output only] The last time a telemetry event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
     */
    public /*out*/ readonly lastEventTime!: pulumi.Output<string>;
    /**
     * [Output only] The last time an MQTT `PINGREQ` was received. This field applies only to devices connecting through MQTT. MQTT clients usually only send `PINGREQ` messages if the connection is idle, and no other messages have been sent. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
     */
    public /*out*/ readonly lastHeartbeatTime!: pulumi.Output<string>;
    /**
     * [Output only] The last time a state event was received. Timestamps are periodically collected and written to storage; they may be stale by a few minutes.
     */
    public /*out*/ readonly lastStateTime!: pulumi.Output<string>;
    /**
     * **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
     */
    public readonly logLevel!: pulumi.Output<string>;
    /**
     * The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * [Output only] A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally unique.
     */
    public /*out*/ readonly numId!: pulumi.Output<string>;
    /**
     * [Output only] The state most recently received from the device. If no state has been reported, this field is not present.
     */
    public /*out*/ readonly state!: pulumi.Output<outputs.cloudiot.v1.DeviceStateResponse>;

    /**
     * Create a Device resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.registryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryId'");
            }
            inputs["blocked"] = args ? args.blocked : undefined;
            inputs["config"] = args ? args.config : undefined;
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["gatewayConfig"] = args ? args.gatewayConfig : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["logLevel"] = args ? args.logLevel : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["registryId"] = args ? args.registryId : undefined;
            inputs["lastConfigAckTime"] = undefined /*out*/;
            inputs["lastConfigSendTime"] = undefined /*out*/;
            inputs["lastErrorStatus"] = undefined /*out*/;
            inputs["lastErrorTime"] = undefined /*out*/;
            inputs["lastEventTime"] = undefined /*out*/;
            inputs["lastHeartbeatTime"] = undefined /*out*/;
            inputs["lastStateTime"] = undefined /*out*/;
            inputs["numId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        } else {
            inputs["blocked"] = undefined /*out*/;
            inputs["config"] = undefined /*out*/;
            inputs["credentials"] = undefined /*out*/;
            inputs["gatewayConfig"] = undefined /*out*/;
            inputs["lastConfigAckTime"] = undefined /*out*/;
            inputs["lastConfigSendTime"] = undefined /*out*/;
            inputs["lastErrorStatus"] = undefined /*out*/;
            inputs["lastErrorTime"] = undefined /*out*/;
            inputs["lastEventTime"] = undefined /*out*/;
            inputs["lastHeartbeatTime"] = undefined /*out*/;
            inputs["lastStateTime"] = undefined /*out*/;
            inputs["logLevel"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["numId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Device.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Device resource.
 */
export interface DeviceArgs {
    /**
     * If a device is blocked, connections or requests from this device will fail. Can be used to temporarily prevent the device from connecting if, for example, the sensor is generating bad data and needs maintenance.
     */
    blocked?: pulumi.Input<boolean>;
    /**
     * The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. If not present on creation, the configuration will be initialized with an empty payload and version value of `1`. To update this field after creation, use the `DeviceManager.ModifyCloudToDeviceConfig` method.
     */
    config?: pulumi.Input<inputs.cloudiot.v1.DeviceConfigArgs>;
    /**
     * The credentials used to authenticate this device. To allow credential rotation without interruption, multiple device credentials can be bound to this device. No more than 3 credentials can be bound to a single device at a time. When new credentials are added to a device, they are verified against the registry credentials. For details, see the description of the `DeviceRegistry.credentials` field.
     */
    credentials?: pulumi.Input<pulumi.Input<inputs.cloudiot.v1.DeviceCredentialArgs>[]>;
    /**
     * Gateway-related configuration and state.
     */
    gatewayConfig?: pulumi.Input<inputs.cloudiot.v1.GatewayConfigArgs>;
    /**
     * The user-defined device identifier. The device ID must be unique within a device registry.
     */
    id?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    /**
     * **Beta Feature** The logging verbosity for device activity. If unspecified, DeviceRegistry.log_level will be used.
     */
    logLevel?: pulumi.Input<enums.cloudiot.v1.DeviceLogLevel>;
    /**
     * The metadata key-value pairs assigned to the device. This metadata is not interpreted or indexed by Cloud IoT Core. It can be used to add contextual information for the device. Keys must conform to the regular expression a-zA-Z+ and be less than 128 bytes in length. Values are free-form strings. Each value must be less than or equal to 32 KB in size. The total size of all keys and values must be less than 256 KB, and the maximum number of key-value pairs is 500.
     */
    metadata?: pulumi.Input<{[key: string]: string}>;
    /**
     * The resource path name. For example, `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`. When `name` is populated as a response from the service, it always ends in the device numeric ID.
     */
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    registryId: pulumi.Input<string>;
}
