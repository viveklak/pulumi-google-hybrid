// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1
{
    /// <summary>
    /// Creates a new HL7v2 store within the parent dataset.
    /// </summary>
    [GoogleNativeResourceType("google-native:healthcare/v1:Hl7V2Store")]
    public partial class Hl7V2Store : Pulumi.CustomResource
    {
        /// <summary>
        /// User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Resource name of the HL7v2 store, of the form `projects/{project_id}/datasets/{dataset_id}/hl7V2Stores/{hl7v2_store_id}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of notification configs. Each configuration uses a filter to determine whether to publish a message (both Ingest &amp; Create) on the corresponding notification destination. Only the message name is sent as part of the notification. Supplied by the client.
        /// </summary>
        [Output("notificationConfigs")]
        public Output<ImmutableArray<Outputs.Hl7V2NotificationConfigResponse>> NotificationConfigs { get; private set; } = null!;

        /// <summary>
        /// The configuration for the parser. It determines how the server parses the messages.
        /// </summary>
        [Output("parserConfig")]
        public Output<Outputs.ParserConfigResponse> ParserConfig { get; private set; } = null!;

        /// <summary>
        /// Determines whether to reject duplicate messages. A duplicate message is a message with the same raw bytes as a message that has already been ingested/created in this HL7v2 store. The default value is false, meaning that the store accepts the duplicate messages and it also returns the same ACK message in the IngestMessageResponse as has been returned previously. Note that only one resource is created in the store. When this field is set to true, CreateMessage/IngestMessage requests with a duplicate message will be rejected by the store, and IngestMessageErrorDetail returns a NACK message upon rejection.
        /// </summary>
        [Output("rejectDuplicateMessage")]
        public Output<bool> RejectDuplicateMessage { get; private set; } = null!;


        /// <summary>
        /// Create a Hl7V2Store resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Hl7V2Store(string name, Hl7V2StoreArgs args, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1:Hl7V2Store", name, args ?? new Hl7V2StoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Hl7V2Store(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1:Hl7V2Store", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Hl7V2Store resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Hl7V2Store Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Hl7V2Store(name, id, options);
        }
    }

    public sealed class Hl7V2StoreArgs : Pulumi.ResourceArgs
    {
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// The ID of the HL7v2 store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
        /// </summary>
        [Input("hl7V2StoreId")]
        public Input<string>? Hl7V2StoreId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Resource name of the HL7v2 store, of the form `projects/{project_id}/datasets/{dataset_id}/hl7V2Stores/{hl7v2_store_id}`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationConfigs")]
        private InputList<Inputs.Hl7V2NotificationConfigArgs>? _notificationConfigs;

        /// <summary>
        /// A list of notification configs. Each configuration uses a filter to determine whether to publish a message (both Ingest &amp; Create) on the corresponding notification destination. Only the message name is sent as part of the notification. Supplied by the client.
        /// </summary>
        public InputList<Inputs.Hl7V2NotificationConfigArgs> NotificationConfigs
        {
            get => _notificationConfigs ?? (_notificationConfigs = new InputList<Inputs.Hl7V2NotificationConfigArgs>());
            set => _notificationConfigs = value;
        }

        /// <summary>
        /// The configuration for the parser. It determines how the server parses the messages.
        /// </summary>
        [Input("parserConfig")]
        public Input<Inputs.ParserConfigArgs>? ParserConfig { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Determines whether to reject duplicate messages. A duplicate message is a message with the same raw bytes as a message that has already been ingested/created in this HL7v2 store. The default value is false, meaning that the store accepts the duplicate messages and it also returns the same ACK message in the IngestMessageResponse as has been returned previously. Note that only one resource is created in the store. When this field is set to true, CreateMessage/IngestMessage requests with a duplicate message will be rejected by the store, and IngestMessageErrorDetail returns a NACK message upon rejection.
        /// </summary>
        [Input("rejectDuplicateMessage")]
        public Input<bool>? RejectDuplicateMessage { get; set; }

        public Hl7V2StoreArgs()
        {
        }
    }
}
