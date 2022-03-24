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
    /// Schema in Avro JSON format.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatacatalogV1PhysicalSchemaAvroSchemaResponse
    {
        /// <summary>
        /// JSON source of the Avro schema.
        /// </summary>
        public readonly string Text;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1PhysicalSchemaAvroSchemaResponse(string text)
        {
            Text = text;
        }
    }
}