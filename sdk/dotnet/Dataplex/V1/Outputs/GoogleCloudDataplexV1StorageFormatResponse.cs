// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Outputs
{

    /// <summary>
    /// Describes the format of the data within its storage location.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1StorageFormatResponse
    {
        /// <summary>
        /// Optional. The compression type associated with the stored data. If unspecified, the data is uncompressed.
        /// </summary>
        public readonly string CompressionFormat;
        /// <summary>
        /// Optional. Additional information about CSV formatted data.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1StorageFormatCsvOptionsResponse Csv;
        /// <summary>
        /// The data format associated with the stored data, which represents content type values. The value is inferred from mime type.
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// Optional. Additional information about CSV formatted data.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1StorageFormatJsonOptionsResponse Json;
        /// <summary>
        /// The mime type descriptor for the data. Must match the pattern {type}/{subtype}. Supported values: application/x-parquet application/x-avro application/x-orc application/x-tfrecord application/json application/{subtypes} text/csv text/ image/{image subtype} video/{video subtype} audio/{audio subtype}
        /// </summary>
        public readonly string MimeType;

        [OutputConstructor]
        private GoogleCloudDataplexV1StorageFormatResponse(
            string compressionFormat,

            Outputs.GoogleCloudDataplexV1StorageFormatCsvOptionsResponse csv,

            string format,

            Outputs.GoogleCloudDataplexV1StorageFormatJsonOptionsResponse json,

            string mimeType)
        {
            CompressionFormat = compressionFormat;
            Csv = csv;
            Format = format;
            Json = json;
            MimeType = mimeType;
        }
    }
}
