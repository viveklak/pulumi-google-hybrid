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
    /// Describe JSON data format.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1ZoneDiscoverySpecJsonOptionsResponse
    {
        /// <summary>
        /// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
        /// </summary>
        public readonly bool DisableTypeInference;
        /// <summary>
        /// Optional. The character encoding of the data. The default is UTF-8.
        /// </summary>
        public readonly string Encoding;

        [OutputConstructor]
        private GoogleCloudDataplexV1ZoneDiscoverySpecJsonOptionsResponse(
            bool disableTypeInference,

            string encoding)
        {
            DisableTypeInference = disableTypeInference;
            Encoding = encoding;
        }
    }
}
