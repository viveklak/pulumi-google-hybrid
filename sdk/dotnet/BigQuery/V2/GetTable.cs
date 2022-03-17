// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2
{
    public static class GetTable
    {
        /// <summary>
        /// Gets the specified table resource by table ID. This method does not return the data in the table, it only returns the table resource, which describes the structure of this table.
        /// </summary>
        public static Task<GetTableResult> InvokeAsync(GetTableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTableResult>("google-native:bigquery/v2:getTable", args ?? new GetTableArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified table resource by table ID. This method does not return the data in the table, it only returns the table resource, which describes the structure of this table.
        /// </summary>
        public static Output<GetTableResult> Invoke(GetTableInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTableResult>("google-native:bigquery/v2:getTable", args ?? new GetTableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTableArgs : Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("selectedFields")]
        public string? SelectedFields { get; set; }

        [Input("tableId", required: true)]
        public string TableId { get; set; } = null!;

        public GetTableArgs()
        {
        }
    }

    public sealed class GetTableInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("selectedFields")]
        public Input<string>? SelectedFields { get; set; }

        [Input("tableId", required: true)]
        public Input<string> TableId { get; set; } = null!;

        public GetTableInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTableResult
    {
        /// <summary>
        /// [Beta] Clustering specification for the table. Must be specified with partitioning, data in the table will be first partitioned and subsequently clustered.
        /// </summary>
        public readonly Outputs.ClusteringResponse Clustering;
        /// <summary>
        /// The time when this table was created, in milliseconds since the epoch.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The default collation of the table.
        /// </summary>
        public readonly string DefaultCollation;
        /// <summary>
        /// [Optional] A user-friendly description of this table.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Custom encryption configuration (e.g., Cloud KMS keys).
        /// </summary>
        public readonly Outputs.EncryptionConfigurationResponse EncryptionConfiguration;
        /// <summary>
        /// A hash of the table metadata. Used to ensure there were no concurrent modifications to the resource when attempting an update. Not guaranteed to change when the table contents or the fields numRows, numBytes, numLongTermBytes or lastModifiedTime change.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// [Optional] The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
        /// </summary>
        public readonly string ExpirationTime;
        /// <summary>
        /// [Optional] Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
        /// </summary>
        public readonly Outputs.ExternalDataConfigurationResponse ExternalDataConfiguration;
        /// <summary>
        /// [Optional] A descriptive name for this table.
        /// </summary>
        public readonly string FriendlyName;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The time when this table was last modified, in milliseconds since the epoch.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The geographic location where the table resides. This value is inherited from the dataset.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// [Optional] Materialized view definition.
        /// </summary>
        public readonly Outputs.MaterializedViewDefinitionResponse MaterializedView;
        /// <summary>
        /// [Output-only, Beta] Present iff this table represents a ML model. Describes the training information for the model, and it is required to run 'PREDICT' queries.
        /// </summary>
        public readonly Outputs.ModelDefinitionResponse Model;
        /// <summary>
        /// The size of this table in bytes, excluding any data in the streaming buffer.
        /// </summary>
        public readonly string NumBytes;
        /// <summary>
        /// The number of bytes in the table that are considered "long-term storage".
        /// </summary>
        public readonly string NumLongTermBytes;
        /// <summary>
        /// [TrustedTester] The physical size of this table in bytes, excluding any data in the streaming buffer. This includes compression and storage used for time travel.
        /// </summary>
        public readonly string NumPhysicalBytes;
        /// <summary>
        /// The number of rows of data in this table, excluding any data in the streaming buffer.
        /// </summary>
        public readonly string NumRows;
        /// <summary>
        /// [TrustedTester] Range partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        /// </summary>
        public readonly Outputs.RangePartitioningResponse RangePartitioning;
        /// <summary>
        /// [Optional] If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
        /// </summary>
        public readonly bool RequirePartitionFilter;
        /// <summary>
        /// [Optional] Describes the schema of this table.
        /// </summary>
        public readonly Outputs.TableSchemaResponse Schema;
        /// <summary>
        /// A URL that can be used to access this resource again.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Snapshot definition.
        /// </summary>
        public readonly Outputs.SnapshotDefinitionResponse SnapshotDefinition;
        /// <summary>
        /// Contains information regarding this table's streaming buffer, if one is present. This field will be absent if the table is not being streamed to or if there is no data in the streaming buffer.
        /// </summary>
        public readonly Outputs.StreamingbufferResponse StreamingBuffer;
        /// <summary>
        /// [Required] Reference describing the ID of this table.
        /// </summary>
        public readonly Outputs.TableReferenceResponse TableReference;
        /// <summary>
        /// Time-based partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        /// </summary>
        public readonly Outputs.TimePartitioningResponse TimePartitioning;
        /// <summary>
        /// Describes the table type. The following values are supported: TABLE: A normal BigQuery table. VIEW: A virtual table defined by a SQL query. SNAPSHOT: An immutable, read-only table that is a copy of another table. [TrustedTester] MATERIALIZED_VIEW: SQL query whose result is persisted. EXTERNAL: A table that references data stored in an external storage system, such as Google Cloud Storage. The default value is TABLE.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// [Optional] The view definition.
        /// </summary>
        public readonly Outputs.ViewDefinitionResponse View;

        [OutputConstructor]
        private GetTableResult(
            Outputs.ClusteringResponse clustering,

            string creationTime,

            string defaultCollation,

            string description,

            Outputs.EncryptionConfigurationResponse encryptionConfiguration,

            string etag,

            string expirationTime,

            Outputs.ExternalDataConfigurationResponse externalDataConfiguration,

            string friendlyName,

            string kind,

            ImmutableDictionary<string, string> labels,

            string lastModifiedTime,

            string location,

            Outputs.MaterializedViewDefinitionResponse materializedView,

            Outputs.ModelDefinitionResponse model,

            string numBytes,

            string numLongTermBytes,

            string numPhysicalBytes,

            string numRows,

            Outputs.RangePartitioningResponse rangePartitioning,

            bool requirePartitionFilter,

            Outputs.TableSchemaResponse schema,

            string selfLink,

            Outputs.SnapshotDefinitionResponse snapshotDefinition,

            Outputs.StreamingbufferResponse streamingBuffer,

            Outputs.TableReferenceResponse tableReference,

            Outputs.TimePartitioningResponse timePartitioning,

            string type,

            Outputs.ViewDefinitionResponse view)
        {
            Clustering = clustering;
            CreationTime = creationTime;
            DefaultCollation = defaultCollation;
            Description = description;
            EncryptionConfiguration = encryptionConfiguration;
            Etag = etag;
            ExpirationTime = expirationTime;
            ExternalDataConfiguration = externalDataConfiguration;
            FriendlyName = friendlyName;
            Kind = kind;
            Labels = labels;
            LastModifiedTime = lastModifiedTime;
            Location = location;
            MaterializedView = materializedView;
            Model = model;
            NumBytes = numBytes;
            NumLongTermBytes = numLongTermBytes;
            NumPhysicalBytes = numPhysicalBytes;
            NumRows = numRows;
            RangePartitioning = rangePartitioning;
            RequirePartitionFilter = requirePartitionFilter;
            Schema = schema;
            SelfLink = selfLink;
            SnapshotDefinition = snapshotDefinition;
            StreamingBuffer = streamingBuffer;
            TableReference = tableReference;
            TimePartitioning = timePartitioning;
            Type = type;
            View = view;
        }
    }
}
