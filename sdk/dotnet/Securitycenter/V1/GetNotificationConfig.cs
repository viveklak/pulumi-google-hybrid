// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Securitycenter.V1
{
    public static class GetNotificationConfig
    {
        /// <summary>
        /// Gets a notification config.
        /// </summary>
        public static Task<GetNotificationConfigResult> InvokeAsync(GetNotificationConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNotificationConfigResult>("google-native:securitycenter/v1:getNotificationConfig", args ?? new GetNotificationConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a notification config.
        /// </summary>
        public static Output<GetNotificationConfigResult> Invoke(GetNotificationConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNotificationConfigResult>("google-native:securitycenter/v1:getNotificationConfig", args ?? new GetNotificationConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotificationConfigArgs : Pulumi.InvokeArgs
    {
        [Input("notificationConfigId", required: true)]
        public string NotificationConfigId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetNotificationConfigArgs()
        {
        }
    }

    public sealed class GetNotificationConfigInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("notificationConfigId", required: true)]
        public Input<string> NotificationConfigId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public GetNotificationConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNotificationConfigResult
    {
        /// <summary>
        /// The description of the notification config (max of 1024 characters).
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The relative resource name of this notification config. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/notificationConfigs/notify_public_bucket".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
        /// </summary>
        public readonly string PubsubTopic;
        /// <summary>
        /// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// The config for triggering streaming-based notifications.
        /// </summary>
        public readonly Outputs.StreamingConfigResponse StreamingConfig;

        [OutputConstructor]
        private GetNotificationConfigResult(
            string description,

            string name,

            string pubsubTopic,

            string serviceAccount,

            Outputs.StreamingConfigResponse streamingConfig)
        {
            Description = description;
            Name = name;
            PubsubTopic = pubsubTopic;
            ServiceAccount = serviceAccount;
            StreamingConfig = streamingConfig;
        }
    }
}
