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
    /// A KeyUsage describes key usage values that may appear in an X.509 certificate.
    /// </summary>
    [OutputType]
    public sealed class KeyUsageResponse
    {
        /// <summary>
        /// Describes high-level ways in which a key may be used.
        /// </summary>
        public readonly Outputs.KeyUsageOptionsResponse BaseKeyUsage;
        /// <summary>
        /// Detailed scenarios in which a key may be used.
        /// </summary>
        public readonly Outputs.ExtendedKeyUsageOptionsResponse ExtendedKeyUsage;
        /// <summary>
        /// Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message.
        /// </summary>
        public readonly ImmutableArray<Outputs.ObjectIdResponse> UnknownExtendedKeyUsages;

        [OutputConstructor]
        private KeyUsageResponse(
            Outputs.KeyUsageOptionsResponse baseKeyUsage,

            Outputs.ExtendedKeyUsageOptionsResponse extendedKeyUsage,

            ImmutableArray<Outputs.ObjectIdResponse> unknownExtendedKeyUsages)
        {
            BaseKeyUsage = baseKeyUsage;
            ExtendedKeyUsage = extendedKeyUsage;
            UnknownExtendedKeyUsages = unknownExtendedKeyUsages;
        }
    }
}
