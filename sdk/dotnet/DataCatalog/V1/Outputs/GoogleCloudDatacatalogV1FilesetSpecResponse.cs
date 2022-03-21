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
    /// Specification that applies to a fileset. Valid only for entries with the 'FILESET' type.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatacatalogV1FilesetSpecResponse
    {
        /// <summary>
        /// Fields specific to a Dataplex fileset and present only in the Dataplex fileset entries.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1DataplexFilesetSpecResponse DataplexFileset;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1FilesetSpecResponse(Outputs.GoogleCloudDatacatalogV1DataplexFilesetSpecResponse dataplexFileset)
        {
            DataplexFileset = dataplexFileset;
        }
    }
}
