// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Outputs
{

    /// <summary>
    /// A storage location within Google cloud storage (GCS).
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudStorageResponse
    {
        /// <summary>
        /// The path to a directory in GCS that will eventually contain the results for this test. The requesting user must have write access on the bucket in the supplied path.
        /// </summary>
        public readonly string GcsPath;

        [OutputConstructor]
        private GoogleCloudStorageResponse(string gcsPath)
        {
            GcsPath = gcsPath;
        }
    }
}
