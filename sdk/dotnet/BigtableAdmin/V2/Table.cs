// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2
{
    /// <summary>
    /// Creates a new table in the specified instance. The table can be created with a full set of initial column families, specified in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:bigtableadmin/v2:Table")]
    public partial class Table : Pulumi.CustomResource
    {
        /// <summary>
        /// Map from cluster ID to per-cluster table state. If it could not be determined whether or not the table has data in a particular cluster (for example, if its zone is unavailable), then there will be an entry for the cluster with UNKNOWN `replication_status`. Views: `REPLICATION_VIEW`, `ENCRYPTION_VIEW`, `FULL`
        /// </summary>
        [Output("clusterStates")]
        public Output<ImmutableDictionary<string, string>> ClusterStates { get; private set; } = null!;

        /// <summary>
        /// The column families configured for this table, mapped by column family ID. Views: `SCHEMA_VIEW`, `FULL`
        /// </summary>
        [Output("columnFamilies")]
        public Output<ImmutableDictionary<string, string>> ColumnFamilies { get; private set; } = null!;

        /// <summary>
        /// Immutable. The granularity (i.e. `MILLIS`) at which timestamps are stored in this table. Timestamps not matching the granularity will be rejected. If unspecified at creation time, the value will be set to `MILLIS`. Views: `SCHEMA_VIEW`, `FULL`.
        /// </summary>
        [Output("granularity")]
        public Output<string> Granularity { get; private set; } = null!;

        /// <summary>
        /// The unique name of the table. Values are of the form `projects/{project}/instances/{instance}/tables/_a-zA-Z0-9*`. Views: `NAME_ONLY`, `SCHEMA_VIEW`, `REPLICATION_VIEW`, `FULL`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If this table was restored from another data source (e.g. a backup), this field will be populated with information about the restore.
        /// </summary>
        [Output("restoreInfo")]
        public Output<Outputs.RestoreInfoResponse> RestoreInfo { get; private set; } = null!;


        /// <summary>
        /// Create a Table resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Table(string name, TableArgs args, CustomResourceOptions? options = null)
            : base("google-native:bigtableadmin/v2:Table", name, args ?? new TableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Table(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:bigtableadmin/v2:Table", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Table resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Table Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Table(name, id, options);
        }
    }

    public sealed class TableArgs : Pulumi.ResourceArgs
    {
        [Input("columnFamilies")]
        private InputMap<string>? _columnFamilies;

        /// <summary>
        /// The column families configured for this table, mapped by column family ID. Views: `SCHEMA_VIEW`, `FULL`
        /// </summary>
        public InputMap<string> ColumnFamilies
        {
            get => _columnFamilies ?? (_columnFamilies = new InputMap<string>());
            set => _columnFamilies = value;
        }

        /// <summary>
        /// Immutable. The granularity (i.e. `MILLIS`) at which timestamps are stored in this table. Timestamps not matching the granularity will be rejected. If unspecified at creation time, the value will be set to `MILLIS`. Views: `SCHEMA_VIEW`, `FULL`.
        /// </summary>
        [Input("granularity")]
        public Input<Pulumi.GoogleNative.BigtableAdmin.V2.TableGranularity>? Granularity { get; set; }

        [Input("initialSplits")]
        private InputList<Inputs.SplitArgs>? _initialSplits;

        /// <summary>
        /// The optional list of row keys that will be used to initially split the table into several tablets (tablets are similar to HBase regions). Given two split keys, `s1` and `s2`, three tablets will be created, spanning the key ranges: `[, s1), [s1, s2), [s2, )`. Example: * Row keys := `["a", "apple", "custom", "customer_1", "customer_2",` `"other", "zz"]` * initial_split_keys := `["apple", "customer_1", "customer_2", "other"]` * Key assignment: - Tablet 1 `[, apple) =&gt; {"a"}.` - Tablet 2 `[apple, customer_1) =&gt; {"apple", "custom"}.` - Tablet 3 `[customer_1, customer_2) =&gt; {"customer_1"}.` - Tablet 4 `[customer_2, other) =&gt; {"customer_2"}.` - Tablet 5 `[other, ) =&gt; {"other", "zz"}.`
        /// </summary>
        public InputList<Inputs.SplitArgs> InitialSplits
        {
            get => _initialSplits ?? (_initialSplits = new InputList<Inputs.SplitArgs>());
            set => _initialSplits = value;
        }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The unique name of the table. Values are of the form `projects/{project}/instances/{instance}/tables/_a-zA-Z0-9*`. Views: `NAME_ONLY`, `SCHEMA_VIEW`, `REPLICATION_VIEW`, `FULL`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name by which the new table should be referred to within the parent instance, e.g., `foobar` rather than `{parent}/tables/foobar`. Maximum 50 characters.
        /// </summary>
        [Input("tableId", required: true)]
        public Input<string> TableId { get; set; } = null!;

        public TableArgs()
        {
        }
    }
}
