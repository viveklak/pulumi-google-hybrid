// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Cloudkms.V1.Outputs
{

    /// <summary>
    /// A Certificate represents an X.509 certificate used to authenticate HTTPS connections to EKM replicas.
    /// </summary>
    [OutputType]
    public sealed class CertificateResponse
    {
        /// <summary>
        /// The issuer distinguished name in RFC 2253 format. Only present if parsed is true.
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// The certificate is not valid after this time. Only present if parsed is true.
        /// </summary>
        public readonly string NotAfterTime;
        /// <summary>
        /// The certificate is not valid before this time. Only present if parsed is true.
        /// </summary>
        public readonly string NotBeforeTime;
        /// <summary>
        /// True if the certificate was parsed successfully.
        /// </summary>
        public readonly bool Parsed;
        /// <summary>
        /// The raw certificate bytes in DER format.
        /// </summary>
        public readonly string RawDer;
        /// <summary>
        /// The certificate serial number as a hex string. Only present if parsed is true.
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// The SHA-256 certificate fingerprint as a hex string. Only present if parsed is true.
        /// </summary>
        public readonly string Sha256Fingerprint;
        /// <summary>
        /// The subject distinguished name in RFC 2253 format. Only present if parsed is true.
        /// </summary>
        public readonly string Subject;
        /// <summary>
        /// The subject Alternative DNS names. Only present if parsed is true.
        /// </summary>
        public readonly ImmutableArray<string> SubjectAlternativeDnsNames;

        [OutputConstructor]
        private CertificateResponse(
            string issuer,

            string notAfterTime,

            string notBeforeTime,

            bool parsed,

            string rawDer,

            string serialNumber,

            string sha256Fingerprint,

            string subject,

            ImmutableArray<string> subjectAlternativeDnsNames)
        {
            Issuer = issuer;
            NotAfterTime = notAfterTime;
            NotBeforeTime = notBeforeTime;
            Parsed = parsed;
            RawDer = rawDer;
            SerialNumber = serialNumber;
            Sha256Fingerprint = sha256Fingerprint;
            Subject = subject;
            SubjectAlternativeDnsNames = subjectAlternativeDnsNames;
        }
    }
}
