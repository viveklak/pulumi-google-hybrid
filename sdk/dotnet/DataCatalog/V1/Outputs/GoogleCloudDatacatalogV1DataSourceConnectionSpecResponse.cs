// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Outputs
{

    /// <summary>
    /// Specification that applies to a data source connection. Valid only for entries with the `DATA_SOURCE_CONNECTION` type.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatacatalogV1DataSourceConnectionSpecResponse
    {
        /// <summary>
        /// Fields specific to BigQuery connections.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1BigQueryConnectionSpecResponse BigqueryConnectionSpec;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1DataSourceConnectionSpecResponse(Outputs.GoogleCloudDatacatalogV1BigQueryConnectionSpecResponse bigqueryConnectionSpec)
        {
            BigqueryConnectionSpec = bigqueryConnectionSpec;
        }
    }
}
