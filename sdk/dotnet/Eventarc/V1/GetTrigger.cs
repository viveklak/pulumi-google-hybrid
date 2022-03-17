// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1
{
    public static class GetTrigger
    {
        /// <summary>
        /// Get a single trigger.
        /// </summary>
        public static Task<GetTriggerResult> InvokeAsync(GetTriggerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTriggerResult>("google-native:eventarc/v1:getTrigger", args ?? new GetTriggerArgs(), options.WithDefaults());

        /// <summary>
        /// Get a single trigger.
        /// </summary>
        public static Output<GetTriggerResult> Invoke(GetTriggerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTriggerResult>("google-native:eventarc/v1:getTrigger", args ?? new GetTriggerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTriggerArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("triggerId", required: true)]
        public string TriggerId { get; set; } = null!;

        public GetTriggerArgs()
        {
        }
    }

    public sealed class GetTriggerInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("triggerId", required: true)]
        public Input<string> TriggerId { get; set; } = null!;

        public GetTriggerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTriggerResult
    {
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Destination specifies where the events should be sent to.
        /// </summary>
        public readonly Outputs.DestinationResponse Destination;
        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and might be sent only on create requests to ensure that the client has an up-to-date value before proceeding.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// null The list of filters that applies to event attributes. Only events that match all the provided filters are sent to the destination.
        /// </summary>
        public readonly ImmutableArray<Outputs.EventFilterResponse> EventFilters;
        /// <summary>
        /// Optional. User labels attached to the triggers that can be used to group resources.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource name of the trigger. Must be unique within the location of the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have the `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. To create Audit Log triggers, the service account should also have the `roles/eventarc.eventReceiver` IAM role.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Optional. To deliver messages, Eventarc might use other GCP products as a transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        /// </summary>
        public readonly Outputs.TransportResponse Transport;
        /// <summary>
        /// Server-assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// The last-modified time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetTriggerResult(
            string createTime,

            Outputs.DestinationResponse destination,

            string etag,

            ImmutableArray<Outputs.EventFilterResponse> eventFilters,

            ImmutableDictionary<string, string> labels,

            string name,

            string serviceAccount,

            Outputs.TransportResponse transport,

            string uid,

            string updateTime)
        {
            CreateTime = createTime;
            Destination = destination;
            Etag = etag;
            EventFilters = eventFilters;
            Labels = labels;
            Name = name;
            ServiceAccount = serviceAccount;
            Transport = transport;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}
