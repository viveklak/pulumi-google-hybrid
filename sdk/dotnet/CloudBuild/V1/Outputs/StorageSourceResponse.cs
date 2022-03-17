// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Outputs
{

    /// <summary>
    /// Location of the source in an archive file in Google Cloud Storage.
    /// </summary>
    [OutputType]
    public sealed class StorageSourceResponse
    {
        /// <summary>
        /// Google Cloud Storage bucket containing the source (see [Bucket Name Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)).
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// Google Cloud Storage generation for the object. If the generation is omitted, the latest generation will be used.
        /// </summary>
        public readonly string Generation;
        /// <summary>
        /// Google Cloud Storage object containing the source. This object must be a zipped (`.zip`) or gzipped archive file (`.tar.gz`) containing source to build.
        /// </summary>
        public readonly string Object;

        [OutputConstructor]
        private StorageSourceResponse(
            string bucket,

            string generation,

            string @object)
        {
            Bucket = bucket;
            Generation = generation;
            Object = @object;
        }
    }
}
