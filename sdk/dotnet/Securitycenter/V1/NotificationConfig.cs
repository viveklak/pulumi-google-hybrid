// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Securitycenter.V1
{
    /// <summary>
    /// Creates a notification config.
    /// </summary>
    [GoogleNativeResourceType("google-native:securitycenter/v1:NotificationConfig")]
    public partial class NotificationConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the notification config (max of 1024 characters).
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The relative resource name of this notification config. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/notificationConfigs/notify_public_bucket".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
        /// </summary>
        [Output("pubsubTopic")]
        public Output<string> PubsubTopic { get; private set; } = null!;

        /// <summary>
        /// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// The config for triggering streaming-based notifications.
        /// </summary>
        [Output("streamingConfig")]
        public Output<Outputs.StreamingConfigResponse> StreamingConfig { get; private set; } = null!;


        /// <summary>
        /// Create a NotificationConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotificationConfig(string name, NotificationConfigArgs args, CustomResourceOptions? options = null)
            : base("google-native:securitycenter/v1:NotificationConfig", name, args ?? new NotificationConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NotificationConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:securitycenter/v1:NotificationConfig", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NotificationConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotificationConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NotificationConfig(name, id, options);
        }
    }

    public sealed class NotificationConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Unique identifier provided by the client within the parent scope. It must be between 1 and 128 characters, and contains alphanumeric characters, underscores or hyphens only.
        /// </summary>
        [Input("configId", required: true)]
        public Input<string> ConfigId { get; set; } = null!;

        /// <summary>
        /// The description of the notification config (max of 1024 characters).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The relative resource name of this notification config. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/notificationConfigs/notify_public_bucket".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
        /// </summary>
        [Input("pubsubTopic")]
        public Input<string>? PubsubTopic { get; set; }

        /// <summary>
        /// The config for triggering streaming-based notifications.
        /// </summary>
        [Input("streamingConfig")]
        public Input<Inputs.StreamingConfigArgs>? StreamingConfig { get; set; }

        public NotificationConfigArgs()
        {
        }
    }
}
