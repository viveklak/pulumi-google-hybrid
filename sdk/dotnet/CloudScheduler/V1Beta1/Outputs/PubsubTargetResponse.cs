// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudScheduler.V1Beta1.Outputs
{

    [OutputType]
    public sealed class PubsubTargetResponse
    {
        /// <summary>
        /// Attributes for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Attributes;
        /// <summary>
        /// The message payload for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// The name of the Cloud Pub/Sub topic to which messages will be published when a job is delivered. The topic name must be in the same format as required by PubSub's [PublishRequest.name](https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#publishrequest), for example `projects/PROJECT_ID/topics/TOPIC_ID`. The topic must be in the same project as the Cloud Scheduler job.
        /// </summary>
        public readonly string TopicName;

        [OutputConstructor]
        private PubsubTargetResponse(
            ImmutableDictionary<string, string> attributes,

            string data,

            string topicName)
        {
            Attributes = attributes;
            Data = data;
            TopicName = topicName;
        }
    }
}
