// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDatacatalogV1BigQueryTableSpecResponse
    {
        /// <summary>
        /// The table source type.
        /// </summary>
        public readonly string TableSourceType;
        /// <summary>
        /// Specification of a BigQuery table. Populated only if the `table_source_type` is `BIGQUERY_TABLE`.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1TableSpecResponse TableSpec;
        /// <summary>
        /// Table view specification. Populated only if the `table_source_type` is `BIGQUERY_VIEW`.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1ViewSpecResponse ViewSpec;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1BigQueryTableSpecResponse(
            string tableSourceType,

            Outputs.GoogleCloudDatacatalogV1TableSpecResponse tableSpec,

            Outputs.GoogleCloudDatacatalogV1ViewSpecResponse viewSpec)
        {
            TableSourceType = tableSourceType;
            TableSpec = tableSpec;
            ViewSpec = viewSpec;
        }
    }
}
