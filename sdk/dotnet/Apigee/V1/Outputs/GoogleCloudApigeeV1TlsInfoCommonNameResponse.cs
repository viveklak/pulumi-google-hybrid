// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudApigeeV1TlsInfoCommonNameResponse
    {
        /// <summary>
        /// The TLS Common Name string of the certificate.
        /// </summary>
        public readonly string Value;
        /// <summary>
        /// Indicates whether the cert should be matched against as a wildcard cert.
        /// </summary>
        public readonly bool WildcardMatch;

        [OutputConstructor]
        private GoogleCloudApigeeV1TlsInfoCommonNameResponse(
            string value,

            bool wildcardMatch)
        {
            Value = value;
            WildcardMatch = wildcardMatch;
        }
    }
}
