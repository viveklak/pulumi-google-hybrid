// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets the specified table resource by table ID. This method does not return the data in the table, it only returns the table resource, which describes the structure of this table.
 */
export function getTable(args: GetTableArgs, opts?: pulumi.InvokeOptions): Promise<GetTableResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:bigquery/v2:getTable", {
        "datasetId": args.datasetId,
        "project": args.project,
        "selectedFields": args.selectedFields,
        "tableId": args.tableId,
    }, opts);
}

export interface GetTableArgs {
    datasetId: string;
    project?: string;
    selectedFields?: string;
    tableId: string;
}

export interface GetTableResult {
    /**
     * Clone definition.
     */
    readonly cloneDefinition: outputs.bigquery.v2.CloneDefinitionResponse;
    /**
     * [Beta] Clustering specification for the table. Must be specified with partitioning, data in the table will be first partitioned and subsequently clustered.
     */
    readonly clustering: outputs.bigquery.v2.ClusteringResponse;
    /**
     * The time when this table was created, in milliseconds since the epoch.
     */
    readonly creationTime: string;
    /**
     * The default collation of the table.
     */
    readonly defaultCollation: string;
    /**
     * [Optional] A user-friendly description of this table.
     */
    readonly description: string;
    /**
     * Custom encryption configuration (e.g., Cloud KMS keys).
     */
    readonly encryptionConfiguration: outputs.bigquery.v2.EncryptionConfigurationResponse;
    /**
     * A hash of the table metadata. Used to ensure there were no concurrent modifications to the resource when attempting an update. Not guaranteed to change when the table contents or the fields numRows, numBytes, numLongTermBytes or lastModifiedTime change.
     */
    readonly etag: string;
    /**
     * [Optional] The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
     */
    readonly expirationTime: string;
    /**
     * [Optional] Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
     */
    readonly externalDataConfiguration: outputs.bigquery.v2.ExternalDataConfigurationResponse;
    /**
     * [Optional] A descriptive name for this table.
     */
    readonly friendlyName: string;
    /**
     * The type of the resource.
     */
    readonly kind: string;
    /**
     * The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
     */
    readonly labels: {[key: string]: string};
    /**
     * The time when this table was last modified, in milliseconds since the epoch.
     */
    readonly lastModifiedTime: string;
    /**
     * The geographic location where the table resides. This value is inherited from the dataset.
     */
    readonly location: string;
    /**
     * [Optional] Materialized view definition.
     */
    readonly materializedView: outputs.bigquery.v2.MaterializedViewDefinitionResponse;
    /**
     * [Output-only, Beta] Present iff this table represents a ML model. Describes the training information for the model, and it is required to run 'PREDICT' queries.
     */
    readonly model: outputs.bigquery.v2.ModelDefinitionResponse;
    /**
     * The size of this table in bytes, excluding any data in the streaming buffer.
     */
    readonly numBytes: string;
    /**
     * The number of bytes in the table that are considered "long-term storage".
     */
    readonly numLongTermBytes: string;
    /**
     * [TrustedTester] The physical size of this table in bytes, excluding any data in the streaming buffer. This includes compression and storage used for time travel.
     */
    readonly numPhysicalBytes: string;
    /**
     * The number of rows of data in this table, excluding any data in the streaming buffer.
     */
    readonly numRows: string;
    /**
     * [TrustedTester] Range partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
     */
    readonly rangePartitioning: outputs.bigquery.v2.RangePartitioningResponse;
    /**
     * [Optional] If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
     */
    readonly requirePartitionFilter: boolean;
    /**
     * [Optional] Describes the schema of this table.
     */
    readonly schema: outputs.bigquery.v2.TableSchemaResponse;
    /**
     * A URL that can be used to access this resource again.
     */
    readonly selfLink: string;
    /**
     * Snapshot definition.
     */
    readonly snapshotDefinition: outputs.bigquery.v2.SnapshotDefinitionResponse;
    /**
     * Contains information regarding this table's streaming buffer, if one is present. This field will be absent if the table is not being streamed to or if there is no data in the streaming buffer.
     */
    readonly streamingBuffer: outputs.bigquery.v2.StreamingbufferResponse;
    /**
     * [Required] Reference describing the ID of this table.
     */
    readonly tableReference: outputs.bigquery.v2.TableReferenceResponse;
    /**
     * Time-based partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
     */
    readonly timePartitioning: outputs.bigquery.v2.TimePartitioningResponse;
    /**
     * Describes the table type. The following values are supported: TABLE: A normal BigQuery table. VIEW: A virtual table defined by a SQL query. SNAPSHOT: An immutable, read-only table that is a copy of another table. [TrustedTester] MATERIALIZED_VIEW: SQL query whose result is persisted. EXTERNAL: A table that references data stored in an external storage system, such as Google Cloud Storage. The default value is TABLE.
     */
    readonly type: string;
    /**
     * [Optional] The view definition.
     */
    readonly view: outputs.bigquery.v2.ViewDefinitionResponse;
}

export function getTableOutput(args: GetTableOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTableResult> {
    return pulumi.output(args).apply(a => getTable(a, opts))
}

export interface GetTableOutputArgs {
    datasetId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    selectedFields?: pulumi.Input<string>;
    tableId: pulumi.Input<string>;
}
