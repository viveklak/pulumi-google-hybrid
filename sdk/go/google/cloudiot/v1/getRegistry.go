// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a device registry configuration.
func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	var rv LookupRegistryResult
	err := ctx.Invoke("google-native:cloudiot/v1:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistryArgs struct {
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
	RegistryId string  `pulumi:"registryId"`
}

type LookupRegistryResult struct {
	// The credentials used to verify the device credentials. No more than 10 credentials can be bound to a single registry at a time. The verification process occurs at the time of device creation or update. If this field is empty, no verification is performed. Otherwise, the credentials of a newly created device or added credentials of an updated device should be signed with one of these registry credentials. Note, however, that existing devices will never be affected by modifications to this list of credentials: after a device has been successfully created in a registry, it should be able to connect even if its registry credentials are revoked, deleted, or modified.
	Credentials []RegistryCredentialResponse `pulumi:"credentials"`
	// The configuration for notification of telemetry events received from the device. All telemetry events that were successfully published by the device and acknowledged by Cloud IoT Core are guaranteed to be delivered to Cloud Pub/Sub. If multiple configurations match a message, only the first matching configuration is used. If you try to publish a device telemetry event using MQTT without specifying a Cloud Pub/Sub topic for the device's registry, the connection closes automatically. If you try to do so using an HTTP connection, an error is returned. Up to 10 configurations may be provided.
	EventNotificationConfigs []EventNotificationConfigResponse `pulumi:"eventNotificationConfigs"`
	// The DeviceService (HTTP) configuration for this device registry.
	HttpConfig HttpConfigResponse `pulumi:"httpConfig"`
	// **Beta Feature** The default logging verbosity for activity from devices in this registry. The verbosity level can be overridden by Device.log_level.
	LogLevel string `pulumi:"logLevel"`
	// The MQTT configuration for this device registry.
	MqttConfig MqttConfigResponse `pulumi:"mqttConfig"`
	// The resource path name. For example, `projects/example-project/locations/us-central1/registries/my-registry`.
	Name string `pulumi:"name"`
	// The configuration for notification of new states received from the device. State updates are guaranteed to be stored in the state history, but notifications to Cloud Pub/Sub are not guaranteed. For example, if permissions are misconfigured or the specified topic doesn't exist, no notification will be published but the state will still be stored in Cloud IoT Core.
	StateNotificationConfig StateNotificationConfigResponse `pulumi:"stateNotificationConfig"`
}

func LookupRegistryOutput(ctx *pulumi.Context, args LookupRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryResult, error) {
			args := v.(LookupRegistryArgs)
			r, err := LookupRegistry(ctx, &args, opts...)
			return *r, err
		}).(LookupRegistryResultOutput)
}

type LookupRegistryOutputArgs struct {
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
	RegistryId pulumi.StringInput    `pulumi:"registryId"`
}

func (LookupRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryArgs)(nil)).Elem()
}

type LookupRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryResult)(nil)).Elem()
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutput() LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutputWithContext(ctx context.Context) LookupRegistryResultOutput {
	return o
}

// The credentials used to verify the device credentials. No more than 10 credentials can be bound to a single registry at a time. The verification process occurs at the time of device creation or update. If this field is empty, no verification is performed. Otherwise, the credentials of a newly created device or added credentials of an updated device should be signed with one of these registry credentials. Note, however, that existing devices will never be affected by modifications to this list of credentials: after a device has been successfully created in a registry, it should be able to connect even if its registry credentials are revoked, deleted, or modified.
func (o LookupRegistryResultOutput) Credentials() RegistryCredentialResponseArrayOutput {
	return o.ApplyT(func(v LookupRegistryResult) []RegistryCredentialResponse { return v.Credentials }).(RegistryCredentialResponseArrayOutput)
}

// The configuration for notification of telemetry events received from the device. All telemetry events that were successfully published by the device and acknowledged by Cloud IoT Core are guaranteed to be delivered to Cloud Pub/Sub. If multiple configurations match a message, only the first matching configuration is used. If you try to publish a device telemetry event using MQTT without specifying a Cloud Pub/Sub topic for the device's registry, the connection closes automatically. If you try to do so using an HTTP connection, an error is returned. Up to 10 configurations may be provided.
func (o LookupRegistryResultOutput) EventNotificationConfigs() EventNotificationConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupRegistryResult) []EventNotificationConfigResponse { return v.EventNotificationConfigs }).(EventNotificationConfigResponseArrayOutput)
}

// The DeviceService (HTTP) configuration for this device registry.
func (o LookupRegistryResultOutput) HttpConfig() HttpConfigResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) HttpConfigResponse { return v.HttpConfig }).(HttpConfigResponseOutput)
}

// **Beta Feature** The default logging verbosity for activity from devices in this registry. The verbosity level can be overridden by Device.log_level.
func (o LookupRegistryResultOutput) LogLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.LogLevel }).(pulumi.StringOutput)
}

// The MQTT configuration for this device registry.
func (o LookupRegistryResultOutput) MqttConfig() MqttConfigResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) MqttConfigResponse { return v.MqttConfig }).(MqttConfigResponseOutput)
}

// The resource path name. For example, `projects/example-project/locations/us-central1/registries/my-registry`.
func (o LookupRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

// The configuration for notification of new states received from the device. State updates are guaranteed to be stored in the state history, but notifications to Cloud Pub/Sub are not guaranteed. For example, if permissions are misconfigured or the specified topic doesn't exist, no notification will be published but the state will still be stored in Cloud IoT Core.
func (o LookupRegistryResultOutput) StateNotificationConfig() StateNotificationConfigResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) StateNotificationConfigResponse { return v.StateNotificationConfig }).(StateNotificationConfigResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryResultOutput{})
}
