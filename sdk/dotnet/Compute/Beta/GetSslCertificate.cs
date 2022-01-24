// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetSslCertificate
    {
        /// <summary>
        /// Returns the specified SslCertificate resource. Gets a list of available SSL certificates by making a list() request.
        /// </summary>
        public static Task<GetSslCertificateResult> InvokeAsync(GetSslCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSslCertificateResult>("google-native:compute/beta:getSslCertificate", args ?? new GetSslCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified SslCertificate resource. Gets a list of available SSL certificates by making a list() request.
        /// </summary>
        public static Output<GetSslCertificateResult> Invoke(GetSslCertificateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSslCertificateResult>("google-native:compute/beta:getSslCertificate", args ?? new GetSslCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSslCertificateArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("sslCertificate", required: true)]
        public string SslCertificate { get; set; } = null!;

        public GetSslCertificateArgs()
        {
        }
    }

    public sealed class GetSslCertificateInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sslCertificate", required: true)]
        public Input<string> SslCertificate { get; set; } = null!;

        public GetSslCertificateInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSslCertificateResult
    {
        /// <summary>
        /// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
        /// </summary>
        public readonly string Certificate;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Expire time of the certificate. RFC3339
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// Type of the resource. Always compute#sslCertificate for SSL certificates.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Configuration and status of a managed SSL certificate.
        /// </summary>
        public readonly Outputs.SslCertificateManagedSslCertificateResponse Managed;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
        /// </summary>
        public readonly string PrivateKey;
        /// <summary>
        /// URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// [Output only] Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Configuration and status of a self-managed SSL certificate.
        /// </summary>
        public readonly Outputs.SslCertificateSelfManagedSslCertificateResponse SelfManaged;
        /// <summary>
        /// Domains associated with the certificate via Subject Alternative Name.
        /// </summary>
        public readonly ImmutableArray<string> SubjectAlternativeNames;
        /// <summary>
        /// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSslCertificateResult(
            string certificate,

            string creationTimestamp,

            string description,

            string expireTime,

            string kind,

            Outputs.SslCertificateManagedSslCertificateResponse managed,

            string name,

            string privateKey,

            string region,

            string selfLink,

            Outputs.SslCertificateSelfManagedSslCertificateResponse selfManaged,

            ImmutableArray<string> subjectAlternativeNames,

            string type)
        {
            Certificate = certificate;
            CreationTimestamp = creationTimestamp;
            Description = description;
            ExpireTime = expireTime;
            Kind = kind;
            Managed = managed;
            Name = name;
            PrivateKey = privateKey;
            Region = region;
            SelfLink = selfLink;
            SelfManaged = selfManaged;
            SubjectAlternativeNames = subjectAlternativeNames;
            Type = type;
        }
    }
}
