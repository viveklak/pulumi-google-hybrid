// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1.Outputs
{

    /// <summary>
    /// SSL configuration information.
    /// </summary>
    [OutputType]
    public sealed class SslConfigResponse
    {
        /// <summary>
        /// Input only. The x509 PEM-encoded certificate of the CA that signed the source database server's certificate. The replica will use this certificate to verify it's connecting to the right host.
        /// </summary>
        public readonly string CaCertificate;
        /// <summary>
        /// Input only. The x509 PEM-encoded certificate that will be used by the replica to authenticate against the source database server.If this field is used then the 'client_key' field is mandatory.
        /// </summary>
        public readonly string ClientCertificate;
        /// <summary>
        /// Input only. The unencrypted PKCS#1 or PKCS#8 PEM-encoded private key associated with the Client Certificate. If this field is used then the 'client_certificate' field is mandatory.
        /// </summary>
        public readonly string ClientKey;
        /// <summary>
        /// The ssl config type according to 'client_key', 'client_certificate' and 'ca_certificate'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SslConfigResponse(
            string caCertificate,

            string clientCertificate,

            string clientKey,

            string type)
        {
            CaCertificate = caCertificate;
            ClientCertificate = clientCertificate;
            ClientKey = clientKey;
            Type = type;
        }
    }
}
