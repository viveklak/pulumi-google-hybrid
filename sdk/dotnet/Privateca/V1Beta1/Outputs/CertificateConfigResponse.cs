// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1Beta1.Outputs
{

    /// <summary>
    /// A CertificateConfig describes an X.509 certificate or CSR that is to be created, as an alternative to using ASN.1.
    /// </summary>
    [OutputType]
    public sealed class CertificateConfigResponse
    {
        /// <summary>
        /// Optional. The public key that corresponds to this config. This is, for example, used when issuing Certificates, but not when creating a self-signed CertificateAuthority or CertificateAuthority CSR.
        /// </summary>
        public readonly Outputs.PublicKeyResponse PublicKey;
        /// <summary>
        /// Describes how some of the technical fields in a certificate should be populated.
        /// </summary>
        public readonly Outputs.ReusableConfigWrapperResponse ReusableConfig;
        /// <summary>
        /// Specifies some of the values in a certificate that are related to the subject.
        /// </summary>
        public readonly Outputs.SubjectConfigResponse SubjectConfig;

        [OutputConstructor]
        private CertificateConfigResponse(
            Outputs.PublicKeyResponse publicKey,

            Outputs.ReusableConfigWrapperResponse reusableConfig,

            Outputs.SubjectConfigResponse subjectConfig)
        {
            PublicKey = publicKey;
            ReusableConfig = reusableConfig;
            SubjectConfig = subjectConfig;
        }
    }
}
