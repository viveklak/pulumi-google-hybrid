// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.CloudIoT.V1
{
    public static class GetRegistry
    {
        /// <summary>
        /// Gets a device registry configuration.
        /// </summary>
        public static Task<GetRegistryResult> InvokeAsync(GetRegistryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistryResult>("google-native:cloudiot/v1:getRegistry", args ?? new GetRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a device registry configuration.
        /// </summary>
        public static Output<GetRegistryResult> Invoke(GetRegistryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegistryResult>("google-native:cloudiot/v1:getRegistry", args ?? new GetRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistryArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("registryId", required: true)]
        public string RegistryId { get; set; } = null!;

        public GetRegistryArgs()
        {
        }
    }

    public sealed class GetRegistryInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        public GetRegistryInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistryResult
    {
        /// <summary>
        /// The credentials used to verify the device credentials. No more than 10 credentials can be bound to a single registry at a time. The verification process occurs at the time of device creation or update. If this field is empty, no verification is performed. Otherwise, the credentials of a newly created device or added credentials of an updated device should be signed with one of these registry credentials. Note, however, that existing devices will never be affected by modifications to this list of credentials: after a device has been successfully created in a registry, it should be able to connect even if its registry credentials are revoked, deleted, or modified.
        /// </summary>
        public readonly ImmutableArray<Outputs.RegistryCredentialResponse> Credentials;
        /// <summary>
        /// The configuration for notification of telemetry events received from the device. All telemetry events that were successfully published by the device and acknowledged by Cloud IoT Core are guaranteed to be delivered to Cloud Pub/Sub. If multiple configurations match a message, only the first matching configuration is used. If you try to publish a device telemetry event using MQTT without specifying a Cloud Pub/Sub topic for the device's registry, the connection closes automatically. If you try to do so using an HTTP connection, an error is returned. Up to 10 configurations may be provided.
        /// </summary>
        public readonly ImmutableArray<Outputs.EventNotificationConfigResponse> EventNotificationConfigs;
        /// <summary>
        /// The DeviceService (HTTP) configuration for this device registry.
        /// </summary>
        public readonly Outputs.HttpConfigResponse HttpConfig;
        /// <summary>
        /// **Beta Feature** The default logging verbosity for activity from devices in this registry. The verbosity level can be overridden by Device.log_level.
        /// </summary>
        public readonly string LogLevel;
        /// <summary>
        /// The MQTT configuration for this device registry.
        /// </summary>
        public readonly Outputs.MqttConfigResponse MqttConfig;
        /// <summary>
        /// The resource path name. For example, `projects/example-project/locations/us-central1/registries/my-registry`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The configuration for notification of new states received from the device. State updates are guaranteed to be stored in the state history, but notifications to Cloud Pub/Sub are not guaranteed. For example, if permissions are misconfigured or the specified topic doesn't exist, no notification will be published but the state will still be stored in Cloud IoT Core.
        /// </summary>
        public readonly Outputs.StateNotificationConfigResponse StateNotificationConfig;

        [OutputConstructor]
        private GetRegistryResult(
            ImmutableArray<Outputs.RegistryCredentialResponse> credentials,

            ImmutableArray<Outputs.EventNotificationConfigResponse> eventNotificationConfigs,

            Outputs.HttpConfigResponse httpConfig,

            string logLevel,

            Outputs.MqttConfigResponse mqttConfig,

            string name,

            Outputs.StateNotificationConfigResponse stateNotificationConfig)
        {
            Credentials = credentials;
            EventNotificationConfigs = eventNotificationConfigs;
            HttpConfig = httpConfig;
            LogLevel = logLevel;
            MqttConfig = mqttConfig;
            Name = name;
            StateNotificationConfig = stateNotificationConfig;
        }
    }
}
