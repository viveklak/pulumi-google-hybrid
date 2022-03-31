// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Inputs
{

    /// <summary>
    /// IssuanceModes specifies the allowed ways in which Certificates may be requested from this CaPool.
    /// </summary>
    public sealed class IssuanceModesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. When true, allows callers to create Certificates by specifying a CertificateConfig.
        /// </summary>
        [Input("allowConfigBasedIssuance")]
        public Input<bool>? AllowConfigBasedIssuance { get; set; }

        /// <summary>
        /// Optional. When true, allows callers to create Certificates by specifying a CSR.
        /// </summary>
        [Input("allowCsrBasedIssuance")]
        public Input<bool>? AllowCsrBasedIssuance { get; set; }

        public IssuanceModesArgs()
        {
        }
    }
}
