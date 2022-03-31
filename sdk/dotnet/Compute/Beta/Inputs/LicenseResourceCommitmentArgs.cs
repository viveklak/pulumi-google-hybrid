// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Commitment for a particular license resource.
    /// </summary>
    public sealed class LicenseResourceCommitmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of licenses purchased.
        /// </summary>
        [Input("amount")]
        public Input<string>? Amount { get; set; }

        /// <summary>
        /// Specifies the core range of the instance for which this license applies.
        /// </summary>
        [Input("coresPerLicense")]
        public Input<string>? CoresPerLicense { get; set; }

        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        [Input("license")]
        public Input<string>? License { get; set; }

        public LicenseResourceCommitmentArgs()
        {
        }
    }
}
