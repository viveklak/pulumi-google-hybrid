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
    /// Cloud Storage bucket profile.
    /// </summary>
    [OutputType]
    public sealed class GcsProfileResponse
    {
        /// <summary>
        /// The Cloud Storage bucket name.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// The root path inside the Cloud Storage bucket.
        /// </summary>
        public readonly string RootPath;

        [OutputConstructor]
        private GcsProfileResponse(
            string bucket,

            string rootPath)
        {
            Bucket = bucket;
            RootPath = rootPath;
        }
    }
}
