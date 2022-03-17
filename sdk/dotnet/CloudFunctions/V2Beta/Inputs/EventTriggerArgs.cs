// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V2Beta.Inputs
{

    /// <summary>
    /// Describes EventTrigger, used to request events to be sent from another service.
    /// </summary>
    public sealed class EventTriggerArgs : Pulumi.ResourceArgs
    {
        [Input("eventFilters")]
        private InputList<Inputs.EventFilterArgs>? _eventFilters;

        /// <summary>
        /// Criteria used to filter events.
        /// </summary>
        public InputList<Inputs.EventFilterArgs> EventFilters
        {
            get => _eventFilters ?? (_eventFilters = new InputList<Inputs.EventFilterArgs>());
            set => _eventFilters = value;
        }

        /// <summary>
        /// The type of event to observe. For example: `google.cloud.audit.log.v1.written` or `google.cloud.pubsub.topic.v1.messagePublished`.
        /// </summary>
        [Input("eventType", required: true)]
        public Input<string> EventType { get; set; } = null!;

        /// <summary>
        /// Optional. The name of a Pub/Sub topic in the same project that will be used as the transport topic for the event delivery. Format: `projects/{project}/topics/{topic}`. This is only valid for events of type `google.cloud.pubsub.topic.v1.messagePublished`. The topic provided here will not be deleted at function deletion.
        /// </summary>
        [Input("pubsubTopic")]
        public Input<string>? PubsubTopic { get; set; }

        /// <summary>
        /// Optional. If unset, then defaults to ignoring failures (i.e. not retrying them).
        /// </summary>
        [Input("retryPolicy")]
        public Input<Pulumi.GoogleNative.CloudFunctions.V2Beta.EventTriggerRetryPolicy>? RetryPolicy { get; set; }

        /// <summary>
        /// Optional. The email of the trigger's service account. The service account must have permission to invoke Cloud Run services, the permission is `run.routes.invoke`. If empty, defaults to the Compute Engine default service account: `{project_number}-compute@developer.gserviceaccount.com`.
        /// </summary>
        [Input("serviceAccountEmail")]
        public Input<string>? ServiceAccountEmail { get; set; }

        /// <summary>
        /// The region that the trigger will be in. The trigger will only receive events originating in this region. It can be the same region as the function, a different region or multi-region, or the global region. If not provided, defaults to the same region as the function.
        /// </summary>
        [Input("triggerRegion")]
        public Input<string>? TriggerRegion { get; set; }

        public EventTriggerArgs()
        {
        }
    }
}
