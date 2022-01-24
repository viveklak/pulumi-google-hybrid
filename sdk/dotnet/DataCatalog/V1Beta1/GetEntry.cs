// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.DataCatalog.V1Beta1
{
    public static class GetEntry
    {
        /// <summary>
        /// Gets an entry.
        /// </summary>
        public static Task<GetEntryResult> InvokeAsync(GetEntryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEntryResult>("google-native:datacatalog/v1beta1:getEntry", args ?? new GetEntryArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an entry.
        /// </summary>
        public static Output<GetEntryResult> Invoke(GetEntryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEntryResult>("google-native:datacatalog/v1beta1:getEntry", args ?? new GetEntryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEntryArgs : Pulumi.InvokeArgs
    {
        [Input("entryGroupId", required: true)]
        public string EntryGroupId { get; set; } = null!;

        [Input("entryId", required: true)]
        public string EntryId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetEntryArgs()
        {
        }
    }

    public sealed class GetEntryInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("entryGroupId", required: true)]
        public Input<string> EntryGroupId { get; set; } = null!;

        [Input("entryId", required: true)]
        public Input<string> EntryId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetEntryInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEntryResult
    {
        /// <summary>
        /// Specification for a group of BigQuery tables with name pattern `[prefix]YYYYMMDD`. Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecResponse BigqueryDateShardedSpec;
        /// <summary>
        /// Specification that applies to a BigQuery table. This is only valid on entries of type `TABLE`.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1beta1BigQueryTableSpecResponse BigqueryTableSpec;
        /// <summary>
        /// Entry description, which can consist of several sentences or paragraphs that describe entry contents. Default value is an empty string.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Display information such as title and description. A short name to identify the entry, for example, "Analytics Data - Jan 2011". Default value is an empty string.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1beta1GcsFilesetSpecResponse GcsFilesetSpec;
        /// <summary>
        /// This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
        /// </summary>
        public readonly string IntegratedSystem;
        /// <summary>
        /// The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [full name of the resource](https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty string.
        /// </summary>
        public readonly string LinkedResource;
        /// <summary>
        /// The Data Catalog resource name of the entry in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id}/entries/{entry_id} Note that this Entry and its child resources may not actually be stored in the location in this name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Schema of the entry. An entry might not have any schema attached to it.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1beta1SchemaResponse Schema;
        /// <summary>
        /// Timestamps about the underlying resource, not about this Data Catalog entry. Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty timestamp.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse SourceSystemTimestamps;
        /// <summary>
        /// The type of the entry. Only used for Entries with types in the EntryType enum.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Statistics on the usage level of the resource.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1beta1UsageSignalResponse UsageSignal;
        /// <summary>
        /// This field indicates the entry's source system that Data Catalog does not integrate with. `user_specified_system` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        /// </summary>
        public readonly string UserSpecifiedSystem;
        /// <summary>
        /// Entry type if it does not fit any of the input-allowed values listed in `EntryType` enum above. When creating an entry, users should check the enum values first, if nothing matches the entry to be created, then provide a custom value, for example "my_special_type". `user_specified_type` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long. Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use `user_specified_type`.
        /// </summary>
        public readonly string UserSpecifiedType;

        [OutputConstructor]
        private GetEntryResult(
            Outputs.GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecResponse bigqueryDateShardedSpec,

            Outputs.GoogleCloudDatacatalogV1beta1BigQueryTableSpecResponse bigqueryTableSpec,

            string description,

            string displayName,

            Outputs.GoogleCloudDatacatalogV1beta1GcsFilesetSpecResponse gcsFilesetSpec,

            string integratedSystem,

            string linkedResource,

            string name,

            Outputs.GoogleCloudDatacatalogV1beta1SchemaResponse schema,

            Outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse sourceSystemTimestamps,

            string type,

            Outputs.GoogleCloudDatacatalogV1beta1UsageSignalResponse usageSignal,

            string userSpecifiedSystem,

            string userSpecifiedType)
        {
            BigqueryDateShardedSpec = bigqueryDateShardedSpec;
            BigqueryTableSpec = bigqueryTableSpec;
            Description = description;
            DisplayName = displayName;
            GcsFilesetSpec = gcsFilesetSpec;
            IntegratedSystem = integratedSystem;
            LinkedResource = linkedResource;
            Name = name;
            Schema = schema;
            SourceSystemTimestamps = sourceSystemTimestamps;
            Type = type;
            UsageSignal = usageSignal;
            UserSpecifiedSystem = userSpecifiedSystem;
            UserSpecifiedType = userSpecifiedType;
        }
    }
}
