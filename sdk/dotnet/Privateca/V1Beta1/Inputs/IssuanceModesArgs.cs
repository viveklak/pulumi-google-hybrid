// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1Beta1.Inputs
{

    /// <summary>
    /// IssuanceModes specifies the allowed ways in which Certificates may be requested from this CertificateAuthority.
    /// </summary>
    public sealed class IssuanceModesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, allows callers to create Certificates by specifying a CertificateConfig.
        /// </summary>
        [Input("allowConfigBasedIssuance", required: true)]
        public Input<bool> AllowConfigBasedIssuance { get; set; } = null!;

        /// <summary>
        /// When true, allows callers to create Certificates by specifying a CSR.
        /// </summary>
        [Input("allowCsrBasedIssuance", required: true)]
        public Input<bool> AllowCsrBasedIssuance { get; set; } = null!;

        public IssuanceModesArgs()
        {
        }
    }
}
