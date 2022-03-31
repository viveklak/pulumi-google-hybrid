// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Commitment for a particular license resource.
    /// </summary>
    [OutputType]
    public sealed class LicenseResourceCommitmentResponse
    {
        /// <summary>
        /// The number of licenses purchased.
        /// </summary>
        public readonly string Amount;
        /// <summary>
        /// Specifies the core range of the instance for which this license applies.
        /// </summary>
        public readonly string CoresPerLicense;
        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        public readonly string License;

        [OutputConstructor]
        private LicenseResourceCommitmentResponse(
            string amount,

            string coresPerLicense,

            string license)
        {
            Amount = amount;
            CoresPerLicense = coresPerLicense;
            License = license;
        }
    }
}
