// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new table in the specified instance. The table can be created with a full set of initial column families, specified in the request.
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:bigtableadmin/v2:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    /**
     * Map from cluster ID to per-cluster table state. If it could not be determined whether or not the table has data in a particular cluster (for example, if its zone is unavailable), then there will be an entry for the cluster with UNKNOWN `replication_status`. Views: `REPLICATION_VIEW`, `ENCRYPTION_VIEW`, `FULL`
     */
    public /*out*/ readonly clusterStates!: pulumi.Output<{[key: string]: string}>;
    /**
     * The column families configured for this table, mapped by column family ID. Views: `SCHEMA_VIEW`, `FULL`
     */
    public readonly columnFamilies!: pulumi.Output<{[key: string]: string}>;
    /**
     * Immutable. The granularity (i.e. `MILLIS`) at which timestamps are stored in this table. Timestamps not matching the granularity will be rejected. If unspecified at creation time, the value will be set to `MILLIS`. Views: `SCHEMA_VIEW`, `FULL`.
     */
    public readonly granularity!: pulumi.Output<string>;
    /**
     * The unique name of the table. Values are of the form `projects/{project}/instances/{instance}/tables/_a-zA-Z0-9*`. Views: `NAME_ONLY`, `SCHEMA_VIEW`, `REPLICATION_VIEW`, `FULL`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If this table was restored from another data source (e.g. a backup), this field will be populated with information about the restore.
     */
    public /*out*/ readonly restoreInfo!: pulumi.Output<outputs.bigtableadmin.v2.RestoreInfoResponse>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.tableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableId'");
            }
            resourceInputs["columnFamilies"] = args ? args.columnFamilies : undefined;
            resourceInputs["granularity"] = args ? args.granularity : undefined;
            resourceInputs["initialSplits"] = args ? args.initialSplits : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["tableId"] = args ? args.tableId : undefined;
            resourceInputs["clusterStates"] = undefined /*out*/;
            resourceInputs["restoreInfo"] = undefined /*out*/;
        } else {
            resourceInputs["clusterStates"] = undefined /*out*/;
            resourceInputs["columnFamilies"] = undefined /*out*/;
            resourceInputs["granularity"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["restoreInfo"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Table.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * The column families configured for this table, mapped by column family ID. Views: `SCHEMA_VIEW`, `FULL`
     */
    columnFamilies?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Immutable. The granularity (i.e. `MILLIS`) at which timestamps are stored in this table. Timestamps not matching the granularity will be rejected. If unspecified at creation time, the value will be set to `MILLIS`. Views: `SCHEMA_VIEW`, `FULL`.
     */
    granularity?: pulumi.Input<enums.bigtableadmin.v2.TableGranularity>;
    /**
     * The optional list of row keys that will be used to initially split the table into several tablets (tablets are similar to HBase regions). Given two split keys, `s1` and `s2`, three tablets will be created, spanning the key ranges: `[, s1), [s1, s2), [s2, )`. Example: * Row keys := `["a", "apple", "custom", "customer_1", "customer_2",` `"other", "zz"]` * initial_split_keys := `["apple", "customer_1", "customer_2", "other"]` * Key assignment: - Tablet 1 `[, apple) => {"a"}.` - Tablet 2 `[apple, customer_1) => {"apple", "custom"}.` - Tablet 3 `[customer_1, customer_2) => {"customer_1"}.` - Tablet 4 `[customer_2, other) => {"customer_2"}.` - Tablet 5 `[other, ) => {"other", "zz"}.`
     */
    initialSplits?: pulumi.Input<pulumi.Input<inputs.bigtableadmin.v2.SplitArgs>[]>;
    instanceId: pulumi.Input<string>;
    /**
     * The unique name of the table. Values are of the form `projects/{project}/instances/{instance}/tables/_a-zA-Z0-9*`. Views: `NAME_ONLY`, `SCHEMA_VIEW`, `REPLICATION_VIEW`, `FULL`
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The name by which the new table should be referred to within the parent instance, e.g., `foobar` rather than `{parent}/tables/foobar`. Maximum 50 characters.
     */
    tableId: pulumi.Input<string>;
}
