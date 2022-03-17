// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1Alpha1.Outputs
{

    /// <summary>
    /// MySQL table.
    /// </summary>
    [OutputType]
    public sealed class MysqlTableResponse
    {
        /// <summary>
        /// MySQL columns in the database. When unspecified as part of include/exclude lists, includes/excludes everything.
        /// </summary>
        public readonly ImmutableArray<Outputs.MysqlColumnResponse> MysqlColumns;
        /// <summary>
        /// Table name.
        /// </summary>
        public readonly string TableName;

        [OutputConstructor]
        private MysqlTableResponse(
            ImmutableArray<Outputs.MysqlColumnResponse> mysqlColumns,

            string tableName)
        {
            MysqlColumns = mysqlColumns;
            TableName = tableName;
        }
    }
}
