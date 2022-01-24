// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Privateca.V1Beta1
{
    public static class GetCertificate
    {
        /// <summary>
        /// Returns a Certificate.
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("google-native:privateca/v1beta1:getCertificate", args ?? new GetCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a Certificate.
        /// </summary>
        public static Output<GetCertificateResult> Invoke(GetCertificateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCertificateResult>("google-native:privateca/v1beta1:getCertificate", args ?? new GetCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateArgs : Pulumi.InvokeArgs
    {
        [Input("certificateAuthorityId", required: true)]
        public string CertificateAuthorityId { get; set; } = null!;

        [Input("certificateId", required: true)]
        public string CertificateId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetCertificateArgs()
        {
        }
    }

    public sealed class GetCertificateInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("certificateAuthorityId", required: true)]
        public Input<string> CertificateAuthorityId { get; set; } = null!;

        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetCertificateInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        /// <summary>
        /// A structured description of the issued X.509 certificate.
        /// </summary>
        public readonly Outputs.CertificateDescriptionResponse CertificateDescription;
        /// <summary>
        /// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
        /// </summary>
        public readonly Outputs.CertificateConfigResponse Config;
        /// <summary>
        /// The time at which this Certificate was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Labels with user-defined metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
        /// </summary>
        public readonly string Lifetime;
        /// <summary>
        /// The resource path for this Certificate in the format `projects/*/locations/*/certificateAuthorities/*/certificates/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The pem-encoded, signed X.509 certificate.
        /// </summary>
        public readonly string PemCertificate;
        /// <summary>
        /// The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
        /// </summary>
        public readonly ImmutableArray<string> PemCertificateChain;
        /// <summary>
        /// Immutable. A pem-encoded X.509 certificate signing request (CSR).
        /// </summary>
        public readonly string PemCsr;
        /// <summary>
        /// Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
        /// </summary>
        public readonly Outputs.RevocationDetailsResponse RevocationDetails;
        /// <summary>
        /// The time at which this Certificate was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetCertificateResult(
            Outputs.CertificateDescriptionResponse certificateDescription,

            Outputs.CertificateConfigResponse config,

            string createTime,

            ImmutableDictionary<string, string> labels,

            string lifetime,

            string name,

            string pemCertificate,

            ImmutableArray<string> pemCertificateChain,

            string pemCsr,

            Outputs.RevocationDetailsResponse revocationDetails,

            string updateTime)
        {
            CertificateDescription = certificateDescription;
            Config = config;
            CreateTime = createTime;
            Labels = labels;
            Lifetime = lifetime;
            Name = name;
            PemCertificate = pemCertificate;
            PemCertificateChain = pemCertificateChain;
            PemCsr = pemCsr;
            RevocationDetails = revocationDetails;
            UpdateTime = updateTime;
        }
    }
}
