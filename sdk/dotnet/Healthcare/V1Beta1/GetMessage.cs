// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1
{
    public static class GetMessage
    {
        /// <summary>
        /// Gets an HL7v2 message.
        /// </summary>
        public static Task<GetMessageResult> InvokeAsync(GetMessageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMessageResult>("google-native:healthcare/v1beta1:getMessage", args ?? new GetMessageArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an HL7v2 message.
        /// </summary>
        public static Output<GetMessageResult> Invoke(GetMessageInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMessageResult>("google-native:healthcare/v1beta1:getMessage", args ?? new GetMessageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMessageArgs : Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("hl7V2StoreId", required: true)]
        public string Hl7V2StoreId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("messageId", required: true)]
        public string MessageId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("view")]
        public string? View { get; set; }

        public GetMessageArgs()
        {
        }
    }

    public sealed class GetMessageInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("hl7V2StoreId", required: true)]
        public Input<string> Hl7V2StoreId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("messageId", required: true)]
        public Input<string> MessageId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetMessageInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMessageResult
    {
        /// <summary>
        /// The datetime when the message was created. Set by the server.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Raw message bytes.
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The message type for this message. MSH-9.1.
        /// </summary>
        public readonly string MessageType;
        /// <summary>
        /// Resource name of the Message, of the form `projects/{project_id}/datasets/{dataset_id}/hl7V2Stores/{hl7_v2_store_id}/messages/{message_id}`. Assigned by the server.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The parsed version of the raw message data.
        /// </summary>
        public readonly Outputs.ParsedDataResponse ParsedData;
        /// <summary>
        /// All patient IDs listed in the PID-2, PID-3, and PID-4 segments of this message.
        /// </summary>
        public readonly ImmutableArray<Outputs.PatientIdResponse> PatientIds;
        /// <summary>
        /// The parsed version of the raw message data schematized according to this store's schemas and type definitions.
        /// </summary>
        public readonly Outputs.SchematizedDataResponse SchematizedData;
        /// <summary>
        /// The hospital that this message came from. MSH-4.
        /// </summary>
        public readonly string SendFacility;
        /// <summary>
        /// The datetime the sending application sent this message. MSH-7.
        /// </summary>
        public readonly string SendTime;

        [OutputConstructor]
        private GetMessageResult(
            string createTime,

            string data,

            ImmutableDictionary<string, string> labels,

            string messageType,

            string name,

            Outputs.ParsedDataResponse parsedData,

            ImmutableArray<Outputs.PatientIdResponse> patientIds,

            Outputs.SchematizedDataResponse schematizedData,

            string sendFacility,

            string sendTime)
        {
            CreateTime = createTime;
            Data = data;
            Labels = labels;
            MessageType = messageType;
            Name = name;
            ParsedData = parsedData;
            PatientIds = patientIds;
            SchematizedData = schematizedData;
            SendFacility = sendFacility;
            SendTime = sendTime;
        }
    }
}
