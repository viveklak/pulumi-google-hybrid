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
    /// Describes an Elliptic Curve key that may be used in a Certificate issued from a CaPool.
    /// </summary>
    [OutputType]
    public sealed class EcKeyTypeResponse
    {
        /// <summary>
        /// Optional. A signature algorithm that must be used. If this is omitted, any EC-based signature algorithm will be allowed.
        /// </summary>
        public readonly string SignatureAlgorithm;

        [OutputConstructor]
        private EcKeyTypeResponse(string signatureAlgorithm)
        {
            SignatureAlgorithm = signatureAlgorithm;
        }
    }
}
