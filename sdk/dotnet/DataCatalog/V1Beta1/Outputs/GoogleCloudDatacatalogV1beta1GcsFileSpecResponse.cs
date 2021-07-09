// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1Beta1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDatacatalogV1beta1GcsFileSpecResponse
    {
        /// <summary>
        /// The full file path. Example: `gs://bucket_name/a/b.txt`.
        /// </summary>
        public readonly string FilePath;
        /// <summary>
        /// Timestamps about the Cloud Storage file.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse GcsTimestamps;
        /// <summary>
        /// The size of the file, in bytes.
        /// </summary>
        public readonly string SizeBytes;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1beta1GcsFileSpecResponse(
            string filePath,

            Outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse gcsTimestamps,

            string sizeBytes)
        {
            FilePath = filePath;
            GcsTimestamps = gcsTimestamps;
            SizeBytes = sizeBytes;
        }
    }
}
