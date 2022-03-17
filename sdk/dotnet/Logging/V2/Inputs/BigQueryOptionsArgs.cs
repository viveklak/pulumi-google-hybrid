// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2.Inputs
{

    /// <summary>
    /// Options that change functionality of a sink exporting data to BigQuery.
    /// </summary>
    public sealed class BigQueryOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Whether to use BigQuery's partition tables (https://cloud.google.com/bigquery/docs/partitioned-tables). By default, Cloud Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned tables the date suffix is no longer present and special query syntax (https://cloud.google.com/bigquery/docs/querying-partitioned-tables) has to be used instead. In both cases, tables are sharded based on UTC timezone.
        /// </summary>
        [Input("usePartitionedTables")]
        public Input<bool>? UsePartitionedTables { get; set; }

        public BigQueryOptionsArgs()
        {
        }
    }
}
