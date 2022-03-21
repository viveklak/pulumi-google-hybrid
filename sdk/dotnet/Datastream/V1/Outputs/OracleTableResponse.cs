// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Outputs
{

    /// <summary>
    /// Oracle table.
    /// </summary>
    [OutputType]
    public sealed class OracleTableResponse
    {
        /// <summary>
        /// Oracle columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.
        /// </summary>
        public readonly ImmutableArray<Outputs.OracleColumnResponse> OracleColumns;
        /// <summary>
        /// Table name.
        /// </summary>
        public readonly string Table;

        [OutputConstructor]
        private OracleTableResponse(
            ImmutableArray<Outputs.OracleColumnResponse> oracleColumns,

            string table)
        {
            OracleColumns = oracleColumns;
            Table = table;
        }
    }
}
