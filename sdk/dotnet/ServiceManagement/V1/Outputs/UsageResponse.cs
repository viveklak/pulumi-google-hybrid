// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Outputs
{

    /// <summary>
    /// Configuration controlling usage of a service.
    /// </summary>
    [OutputType]
    public sealed class UsageResponse
    {
        /// <summary>
        /// The full resource name of a channel used for sending notifications to the service producer. Google Service Management currently only supports [Google Cloud Pub/Sub](https://cloud.google.com/pubsub) as a notification channel. To use Google Cloud Pub/Sub as the channel, this must be the name of a Cloud Pub/Sub topic that uses the Cloud Pub/Sub topic name format documented in https://cloud.google.com/pubsub/docs/overview.
        /// </summary>
        public readonly string ProducerNotificationChannel;
        /// <summary>
        /// Requirements that must be satisfied before a consumer project can use the service. Each requirement is of the form /; for example 'serviceusage.googleapis.com/billing-enabled'. For Google APIs, a Terms of Service requirement must be included here. Google Cloud APIs must include "serviceusage.googleapis.com/tos/cloud". Other Google APIs should include "serviceusage.googleapis.com/tos/universal". Additional ToS can be included based on the business needs.
        /// </summary>
        public readonly ImmutableArray<string> Requirements;
        /// <summary>
        /// A list of usage rules that apply to individual API methods. **NOTE:** All service configuration rules follow "last one wins" order.
        /// </summary>
        public readonly ImmutableArray<Outputs.UsageRuleResponse> Rules;

        [OutputConstructor]
        private UsageResponse(
            string producerNotificationChannel,

            ImmutableArray<string> requirements,

            ImmutableArray<Outputs.UsageRuleResponse> rules)
        {
            ProducerNotificationChannel = producerNotificationChannel;
            Requirements = requirements;
            Rules = rules;
        }
    }
}
