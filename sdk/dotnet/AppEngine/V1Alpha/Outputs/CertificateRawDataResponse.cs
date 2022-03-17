// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Alpha.Outputs
{

    /// <summary>
    /// An SSL certificate obtained from a certificate authority.
    /// </summary>
    [OutputType]
    public sealed class CertificateRawDataResponse
    {
        /// <summary>
        /// Unencrypted PEM encoded RSA private key. This field is set once on certificate creation and then encrypted. The key size must be 2048 bits or fewer. Must include the header and footer. Example: -----BEGIN RSA PRIVATE KEY----- -----END RSA PRIVATE KEY----- @InputOnly
        /// </summary>
        public readonly string PrivateKey;
        /// <summary>
        /// PEM encoded x.509 public key certificate. This field is set once on certificate creation. Must include the header and footer. Example: -----BEGIN CERTIFICATE----- -----END CERTIFICATE----- 
        /// </summary>
        public readonly string PublicCertificate;

        [OutputConstructor]
        private CertificateRawDataResponse(
            string privateKey,

            string publicCertificate)
        {
            PrivateKey = privateKey;
            PublicCertificate = publicCertificate;
        }
    }
}
