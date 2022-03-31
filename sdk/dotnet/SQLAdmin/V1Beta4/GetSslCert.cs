// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4
{
    public static class GetSslCert
    {
        /// <summary>
        /// Retrieves a particular SSL certificate. Does not include the private key (required for usage). The private key must be saved from the response to initial creation.
        /// </summary>
        public static Task<GetSslCertResult> InvokeAsync(GetSslCertArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSslCertResult>("google-native:sqladmin/v1beta4:getSslCert", args ?? new GetSslCertArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a particular SSL certificate. Does not include the private key (required for usage). The private key must be saved from the response to initial creation.
        /// </summary>
        public static Output<GetSslCertResult> Invoke(GetSslCertInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSslCertResult>("google-native:sqladmin/v1beta4:getSslCert", args ?? new GetSslCertInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSslCertArgs : Pulumi.InvokeArgs
    {
        [Input("instance", required: true)]
        public string Instance { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sha1Fingerprint", required: true)]
        public string Sha1Fingerprint { get; set; } = null!;

        public GetSslCertArgs()
        {
        }
    }

    public sealed class GetSslCertInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sha1Fingerprint", required: true)]
        public Input<string> Sha1Fingerprint { get; set; } = null!;

        public GetSslCertInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSslCertResult
    {
        /// <summary>
        /// PEM representation.
        /// </summary>
        public readonly string Cert;
        /// <summary>
        /// Serial number, as extracted from the certificate.
        /// </summary>
        public readonly string CertSerialNumber;
        /// <summary>
        /// User supplied name. Constrained to [a-zA-Z.-_ ]+.
        /// </summary>
        public readonly string CommonName;
        /// <summary>
        /// The time when the certificate was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The time when the certificate expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        public readonly string ExpirationTime;
        /// <summary>
        /// Name of the database instance.
        /// </summary>
        public readonly string Instance;
        /// <summary>
        /// This is always `sql#sslCert`.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The URI of this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Sha1 Fingerprint.
        /// </summary>
        public readonly string Sha1Fingerprint;

        [OutputConstructor]
        private GetSslCertResult(
            string cert,

            string certSerialNumber,

            string commonName,

            string createTime,

            string expirationTime,

            string instance,

            string kind,

            string selfLink,

            string sha1Fingerprint)
        {
            Cert = cert;
            CertSerialNumber = certSerialNumber;
            CommonName = commonName;
            CreateTime = createTime;
            ExpirationTime = expirationTime;
            Instance = instance;
            Kind = kind;
            SelfLink = selfLink;
            Sha1Fingerprint = sha1Fingerprint;
        }
    }
}
