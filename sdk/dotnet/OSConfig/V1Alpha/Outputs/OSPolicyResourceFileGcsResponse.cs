// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Alpha.Outputs
{

    /// <summary>
    /// Specifies a file available as a Cloud Storage Object.
    /// </summary>
    [OutputType]
    public sealed class OSPolicyResourceFileGcsResponse
    {
        /// <summary>
        /// Bucket of the Cloud Storage object.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// Generation number of the Cloud Storage object.
        /// </summary>
        public readonly string Generation;
        /// <summary>
        /// Name of the Cloud Storage object.
        /// </summary>
        public readonly string Object;

        [OutputConstructor]
        private OSPolicyResourceFileGcsResponse(
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
