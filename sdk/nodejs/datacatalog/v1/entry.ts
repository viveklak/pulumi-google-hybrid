// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an entry. You can create entries only with 'FILESET', 'CLUSTER', 'DATA_STREAM', or custom types. Data Catalog automatically creates entries with other types during metadata ingestion from integrated systems. You must enable the Data Catalog API in the project identified by the `parent` parameter. For more information, see [Data Catalog resource project](https://cloud.google.com/data-catalog/docs/concepts/resource-project). An entry group can have a maximum of 100,000 entries.
 * Auto-naming is currently not supported for this resource.
 */
export class Entry extends pulumi.CustomResource {
    /**
     * Get an existing Entry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Entry {
        return new Entry(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datacatalog/v1:Entry';

    /**
     * Returns true if the given object is an instance of Entry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Entry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Entry.__pulumiType;
    }

    /**
     * Specification for a group of BigQuery tables with the `[prefix]YYYYMMDD` name pattern. For more information, see [Introduction to partitioned tables] (https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding).
     */
    public readonly bigqueryDateShardedSpec!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1BigQueryDateShardedSpecResponse>;
    /**
     * Specification that applies to a BigQuery table. Valid only for entries with the `TABLE` type.
     */
    public readonly bigqueryTableSpec!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1BigQueryTableSpecResponse>;
    /**
     * Business Context of the entry.
     */
    public readonly businessContext!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1BusinessContextResponse>;
    /**
     * Physical location of the entry.
     */
    public /*out*/ readonly dataSource!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1DataSourceResponse>;
    /**
     * Specification that applies to a data source connection. Valid only for entries with the `DATA_SOURCE_CONNECTION` type.
     */
    public readonly dataSourceConnectionSpec!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1DataSourceConnectionSpecResponse>;
    /**
     * Specification that applies to a table resource. Valid only for entries with the `TABLE` type.
     */
    public readonly databaseTableSpec!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1DatabaseTableSpecResponse>;
    /**
     * Entry description that can consist of several sentences or paragraphs that describe entry contents. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). The maximum size is 2000 bytes when encoded in UTF-8. Default value is an empty string.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Display name of an entry. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum size is 200 bytes when encoded in UTF-8. Default value is an empty string.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Fully qualified name (FQN) of the resource. Set automatically for entries representing resources from synced systems. Settable only during creation and read-only afterwards. Can be used for search and lookup of the entries. FQNs take two forms: * For non-regionalized resources: `{SYSTEM}:{PROJECT}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` * For regionalized resources: `{SYSTEM}:{PROJECT}.{LOCATION_ID}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` Example for a DPMS table: `dataproc_metastore:{PROJECT_ID}.{LOCATION_ID}.{INSTANCE_ID}.{DATABASE_ID}.{TABLE_ID}`
     */
    public readonly fullyQualifiedName!: pulumi.Output<string>;
    /**
     * Specification that applies to a Cloud Storage fileset. Valid only for entries with the `FILESET` type.
     */
    public readonly gcsFilesetSpec!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1GcsFilesetSpecResponse>;
    /**
     * Indicates the entry's source system that Data Catalog integrates with, such as BigQuery, Pub/Sub, or Dataproc Metastore.
     */
    public /*out*/ readonly integratedSystem!: pulumi.Output<string>;
    /**
     * Cloud labels attached to the entry. In Data Catalog, you can create and modify labels attached only to custom entries. Synced entries have unmodifiable labels that come from the source system.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [Full Resource Name] (https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: `//bigquery.googleapis.com/projects/{PROJECT_ID}/datasets/{DATASET_ID}/tables/{TABLE_ID}` Output only when the entry is one of the types in the `EntryType` enum. For entries with a `user_specified_type`, this field is optional and defaults to an empty string. The resource string must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), periods (.), colons (:), slashes (/), dashes (-), and hashes (#). The maximum size is 200 bytes when encoded in UTF-8.
     */
    public readonly linkedResource!: pulumi.Output<string>;
    /**
     * The resource name of an entry in URL format. Note: The entry itself and its child resources might not be stored in the location specified in its name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Additional information related to the entry. Private to the current user.
     */
    public /*out*/ readonly personalDetails!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1PersonalDetailsResponse>;
    /**
     * Specification that applies to a user-defined function or procedure. Valid only for entries with the `ROUTINE` type.
     */
    public readonly routineSpec!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1RoutineSpecResponse>;
    /**
     * Schema of the entry. An entry might not have any schema attached to it.
     */
    public readonly schema!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1SchemaResponse>;
    /**
     * Timestamps from the underlying resource, not from the Data Catalog entry. Output only when the entry has a type listed in the `EntryType` enum. For entries with `user_specified_type`, this field is optional and defaults to an empty timestamp.
     */
    public readonly sourceSystemTimestamps!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1SystemTimestampsResponse>;
    /**
     * The type of the entry. Only used for entries with types listed in the `EntryType` enum. Currently, only `FILESET` enum value is allowed. All other entries created in Data Catalog must use the `user_specified_type`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Resource usage statistics.
     */
    public /*out*/ readonly usageSignal!: pulumi.Output<outputs.datacatalog.v1.GoogleCloudDatacatalogV1UsageSignalResponse>;
    /**
     * Indicates the entry's source system that Data Catalog doesn't automatically integrate with. The `user_specified_system` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
     */
    public readonly userSpecifiedSystem!: pulumi.Output<string>;
    /**
     * Custom entry type that doesn't match any of the values allowed for input and listed in the `EntryType` enum. When creating an entry, first check the type values in the enum. If there are no appropriate types for the new entry, provide a custom value, for example, `my_special_type`. The `user_specified_type` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
     */
    public readonly userSpecifiedType!: pulumi.Output<string>;

    /**
     * Create a Entry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntryArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.entryGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entryGroupId'");
            }
            if ((!args || args.entryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entryId'");
            }
            resourceInputs["bigqueryDateShardedSpec"] = args ? args.bigqueryDateShardedSpec : undefined;
            resourceInputs["bigqueryTableSpec"] = args ? args.bigqueryTableSpec : undefined;
            resourceInputs["businessContext"] = args ? args.businessContext : undefined;
            resourceInputs["dataSourceConnectionSpec"] = args ? args.dataSourceConnectionSpec : undefined;
            resourceInputs["databaseTableSpec"] = args ? args.databaseTableSpec : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["entryGroupId"] = args ? args.entryGroupId : undefined;
            resourceInputs["entryId"] = args ? args.entryId : undefined;
            resourceInputs["fullyQualifiedName"] = args ? args.fullyQualifiedName : undefined;
            resourceInputs["gcsFilesetSpec"] = args ? args.gcsFilesetSpec : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["linkedResource"] = args ? args.linkedResource : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["routineSpec"] = args ? args.routineSpec : undefined;
            resourceInputs["schema"] = args ? args.schema : undefined;
            resourceInputs["sourceSystemTimestamps"] = args ? args.sourceSystemTimestamps : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userSpecifiedSystem"] = args ? args.userSpecifiedSystem : undefined;
            resourceInputs["userSpecifiedType"] = args ? args.userSpecifiedType : undefined;
            resourceInputs["dataSource"] = undefined /*out*/;
            resourceInputs["integratedSystem"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["personalDetails"] = undefined /*out*/;
            resourceInputs["usageSignal"] = undefined /*out*/;
        } else {
            resourceInputs["bigqueryDateShardedSpec"] = undefined /*out*/;
            resourceInputs["bigqueryTableSpec"] = undefined /*out*/;
            resourceInputs["businessContext"] = undefined /*out*/;
            resourceInputs["dataSource"] = undefined /*out*/;
            resourceInputs["dataSourceConnectionSpec"] = undefined /*out*/;
            resourceInputs["databaseTableSpec"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["fullyQualifiedName"] = undefined /*out*/;
            resourceInputs["gcsFilesetSpec"] = undefined /*out*/;
            resourceInputs["integratedSystem"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["linkedResource"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["personalDetails"] = undefined /*out*/;
            resourceInputs["routineSpec"] = undefined /*out*/;
            resourceInputs["schema"] = undefined /*out*/;
            resourceInputs["sourceSystemTimestamps"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["usageSignal"] = undefined /*out*/;
            resourceInputs["userSpecifiedSystem"] = undefined /*out*/;
            resourceInputs["userSpecifiedType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Entry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Entry resource.
 */
export interface EntryArgs {
    /**
     * Specification for a group of BigQuery tables with the `[prefix]YYYYMMDD` name pattern. For more information, see [Introduction to partitioned tables] (https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding).
     */
    bigqueryDateShardedSpec?: pulumi.Input<inputs.datacatalog.v1.GoogleCloudDatacatalogV1BigQueryDateShardedSpecArgs>;
    /**
     * Specification that applies to a BigQuery table. Valid only for entries with the `TABLE` type.
     */
    bigqueryTableSpec?: pulumi.Input<inputs.datacatalog.v1.GoogleCloudDatacatalogV1BigQueryTableSpecArgs>;
    /**
     * Business Context of the entry.
     */
    businessContext?: pulumi.Input<inputs.datacatalog.v1.GoogleCloudDatacatalogV1BusinessContextArgs>;
    /**
     * Specification that applies to a data source connection. Valid only for entries with the `DATA_SOURCE_CONNECTION` type.
     */
    dataSourceConnectionSpec?: pulumi.Input<inputs.datacatalog.v1.GoogleCloudDatacatalogV1DataSourceConnectionSpecArgs>;
    /**
     * Specification that applies to a table resource. Valid only for entries with the `TABLE` type.
     */
    databaseTableSpec?: pulumi.Input<inputs.datacatalog.v1.GoogleCloudDatacatalogV1DatabaseTableSpecArgs>;
    /**
     * Entry description that can consist of several sentences or paragraphs that describe entry contents. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). The maximum size is 2000 bytes when encoded in UTF-8. Default value is an empty string.
     */
    description?: pulumi.Input<string>;
    /**
     * Display name of an entry. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum size is 200 bytes when encoded in UTF-8. Default value is an empty string.
     */
    displayName?: pulumi.Input<string>;
    entryGroupId: pulumi.Input<string>;
    entryId: pulumi.Input<string>;
    /**
     * Fully qualified name (FQN) of the resource. Set automatically for entries representing resources from synced systems. Settable only during creation and read-only afterwards. Can be used for search and lookup of the entries. FQNs take two forms: * For non-regionalized resources: `{SYSTEM}:{PROJECT}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` * For regionalized resources: `{SYSTEM}:{PROJECT}.{LOCATION_ID}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` Example for a DPMS table: `dataproc_metastore:{PROJECT_ID}.{LOCATION_ID}.{INSTANCE_ID}.{DATABASE_ID}.{TABLE_ID}`
     */
    fullyQualifiedName?: pulumi.Input<string>;
    /**
     * Specification that applies to a Cloud Storage fileset. Valid only for entries with the `FILESET` type.
     */
    gcsFilesetSpec?: pulumi.Input<inputs.datacatalog.v1.GoogleCloudDatacatalogV1GcsFilesetSpecArgs>;
    /**
     * Cloud labels attached to the entry. In Data Catalog, you can create and modify labels attached only to custom entries. Synced entries have unmodifiable labels that come from the source system.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [Full Resource Name] (https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: `//bigquery.googleapis.com/projects/{PROJECT_ID}/datasets/{DATASET_ID}/tables/{TABLE_ID}` Output only when the entry is one of the types in the `EntryType` enum. For entries with a `user_specified_type`, this field is optional and defaults to an empty string. The resource string must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), periods (.), colons (:), slashes (/), dashes (-), and hashes (#). The maximum size is 200 bytes when encoded in UTF-8.
     */
    linkedResource?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Specification that applies to a user-defined function or procedure. Valid only for entries with the `ROUTINE` type.
     */
    routineSpec?: pulumi.Input<inputs.datacatalog.v1.GoogleCloudDatacatalogV1RoutineSpecArgs>;
    /**
     * Schema of the entry. An entry might not have any schema attached to it.
     */
    schema?: pulumi.Input<inputs.datacatalog.v1.GoogleCloudDatacatalogV1SchemaArgs>;
    /**
     * Timestamps from the underlying resource, not from the Data Catalog entry. Output only when the entry has a type listed in the `EntryType` enum. For entries with `user_specified_type`, this field is optional and defaults to an empty timestamp.
     */
    sourceSystemTimestamps?: pulumi.Input<inputs.datacatalog.v1.GoogleCloudDatacatalogV1SystemTimestampsArgs>;
    /**
     * The type of the entry. Only used for entries with types listed in the `EntryType` enum. Currently, only `FILESET` enum value is allowed. All other entries created in Data Catalog must use the `user_specified_type`.
     */
    type?: pulumi.Input<enums.datacatalog.v1.EntryType>;
    /**
     * Indicates the entry's source system that Data Catalog doesn't automatically integrate with. The `user_specified_system` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
     */
    userSpecifiedSystem?: pulumi.Input<string>;
    /**
     * Custom entry type that doesn't match any of the values allowed for input and listed in the `EntryType` enum. When creating an entry, first check the type values in the enum. If there are no appropriate types for the new entry, provide a custom value, for example, `my_special_type`. The `user_specified_type` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
     */
    userSpecifiedType?: pulumi.Input<string>;
}
