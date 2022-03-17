// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Outputs
{

    /// <summary>
    /// Describes a set of X.509 extensions that may be part of some certificate issuance controls.
    /// </summary>
    [OutputType]
    public sealed class CertificateExtensionConstraintsResponse
    {
        /// <summary>
        /// Optional. A set of ObjectIds identifying custom X.509 extensions. Will be combined with known_extensions to determine the full set of X.509 extensions.
        /// </summary>
        public readonly ImmutableArray<Outputs.ObjectIdResponse> AdditionalExtensions;
        /// <summary>
        /// Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.
        /// </summary>
        public readonly ImmutableArray<string> KnownExtensions;

        [OutputConstructor]
        private CertificateExtensionConstraintsResponse(
            ImmutableArray<Outputs.ObjectIdResponse> additionalExtensions,

            ImmutableArray<string> knownExtensions)
        {
            AdditionalExtensions = additionalExtensions;
            KnownExtensions = knownExtensions;
        }
    }
}
