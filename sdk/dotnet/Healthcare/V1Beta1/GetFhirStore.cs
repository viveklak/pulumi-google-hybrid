// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1
{
    public static class GetFhirStore
    {
        /// <summary>
        /// Gets the configuration of the specified FHIR store.
        /// </summary>
        public static Task<GetFhirStoreResult> InvokeAsync(GetFhirStoreArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFhirStoreResult>("google-native:healthcare/v1beta1:getFhirStore", args ?? new GetFhirStoreArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the configuration of the specified FHIR store.
        /// </summary>
        public static Output<GetFhirStoreResult> Invoke(GetFhirStoreInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFhirStoreResult>("google-native:healthcare/v1beta1:getFhirStore", args ?? new GetFhirStoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFhirStoreArgs : Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("fhirStoreId", required: true)]
        public string FhirStoreId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetFhirStoreArgs()
        {
        }
    }

    public sealed class GetFhirStoreInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("fhirStoreId", required: true)]
        public Input<string> FhirStoreId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetFhirStoreInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFhirStoreResult
    {
        /// <summary>
        /// If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
        /// </summary>
        public readonly bool DefaultSearchHandlingStrict;
        /// <summary>
        /// Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
        /// </summary>
        public readonly bool DisableReferentialIntegrity;
        /// <summary>
        /// Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
        /// </summary>
        public readonly bool DisableResourceVersioning;
        /// <summary>
        /// Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
        /// </summary>
        public readonly bool EnableUpdateCreate;
        /// <summary>
        /// User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Resource name of the FHIR store, of the form `projects/{project_id}/datasets/{dataset_id}/fhirStores/{fhir_store_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource".
        /// </summary>
        public readonly Outputs.NotificationConfigResponse NotificationConfig;
        /// <summary>
        /// Configuration for how FHIR resource can be searched.
        /// </summary>
        public readonly Outputs.SearchConfigResponse SearchConfig;
        /// <summary>
        /// A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
        /// </summary>
        public readonly ImmutableArray<Outputs.StreamConfigResponse> StreamConfigs;
        /// <summary>
        /// Configuration for how to validate incoming FHIR resources against configured profiles.
        /// </summary>
        public readonly Outputs.ValidationConfigResponse ValidationConfig;
        /// <summary>
        /// Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetFhirStoreResult(
            bool defaultSearchHandlingStrict,

            bool disableReferentialIntegrity,

            bool disableResourceVersioning,

            bool enableUpdateCreate,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.NotificationConfigResponse notificationConfig,

            Outputs.SearchConfigResponse searchConfig,

            ImmutableArray<Outputs.StreamConfigResponse> streamConfigs,

            Outputs.ValidationConfigResponse validationConfig,

            string version)
        {
            DefaultSearchHandlingStrict = defaultSearchHandlingStrict;
            DisableReferentialIntegrity = disableReferentialIntegrity;
            DisableResourceVersioning = disableResourceVersioning;
            EnableUpdateCreate = enableUpdateCreate;
            Labels = labels;
            Name = name;
            NotificationConfig = notificationConfig;
            SearchConfig = searchConfig;
            StreamConfigs = streamConfigs;
            ValidationConfig = validationConfig;
            Version = version;
        }
    }
}
