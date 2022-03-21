// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudScheduler.V1.Inputs
{

    /// <summary>
    /// Pub/Sub target. The job will be delivered by publishing a message to the given Pub/Sub topic.
    /// </summary>
    public sealed class PubsubTargetArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<string>? _attributes;

        /// <summary>
        /// Attributes for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.
        /// </summary>
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        /// <summary>
        /// The message payload for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// The name of the Cloud Pub/Sub topic to which messages will be published when a job is delivered. The topic name must be in the same format as required by Pub/Sub's [PublishRequest.name](https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#publishrequest), for example `projects/PROJECT_ID/topics/TOPIC_ID`. The topic must be in the same project as the Cloud Scheduler job.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public PubsubTargetArgs()
        {
        }
    }
}
