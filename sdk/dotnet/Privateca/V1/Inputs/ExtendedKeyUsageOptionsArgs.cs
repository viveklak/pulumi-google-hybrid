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
    /// KeyUsage.ExtendedKeyUsageOptions has fields that correspond to certain common OIDs that could be specified as an extended key usage value.
    /// </summary>
    public sealed class ExtendedKeyUsageOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.
        /// </summary>
        [Input("clientAuth")]
        public Input<bool>? ClientAuth { get; set; }

        /// <summary>
        /// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".
        /// </summary>
        [Input("codeSigning")]
        public Input<bool>? CodeSigning { get; set; }

        /// <summary>
        /// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".
        /// </summary>
        [Input("emailProtection")]
        public Input<bool>? EmailProtection { get; set; }

        /// <summary>
        /// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".
        /// </summary>
        [Input("ocspSigning")]
        public Input<bool>? OcspSigning { get; set; }

        /// <summary>
        /// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.
        /// </summary>
        [Input("serverAuth")]
        public Input<bool>? ServerAuth { get; set; }

        /// <summary>
        /// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".
        /// </summary>
        [Input("timeStamping")]
        public Input<bool>? TimeStamping { get; set; }

        public ExtendedKeyUsageOptionsArgs()
        {
        }
    }
}
