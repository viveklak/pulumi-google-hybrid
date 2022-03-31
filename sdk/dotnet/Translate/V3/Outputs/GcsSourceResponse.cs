// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Translate.V3.Outputs
{

    /// <summary>
    /// The Google Cloud Storage location for the input content.
    /// </summary>
    [OutputType]
    public sealed class GcsSourceResponse
    {
        /// <summary>
        /// Source data URI. For example, `gs://my_bucket/my_object`.
        /// </summary>
        public readonly string InputUri;

        [OutputConstructor]
        private GcsSourceResponse(string inputUri)
        {
            InputUri = inputUri;
        }
    }
}
