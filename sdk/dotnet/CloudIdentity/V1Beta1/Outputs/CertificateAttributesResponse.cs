// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1Beta1.Outputs
{

    /// <summary>
    /// Stores information about a certificate.
    /// </summary>
    [OutputType]
    public sealed class CertificateAttributesResponse
    {
        /// <summary>
        /// The X.509 extension for CertificateTemplate.
        /// </summary>
        public readonly Outputs.CertificateTemplateResponse CertificateTemplate;
        /// <summary>
        /// The encoded certificate fingerprint.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// The name of the issuer of this certificate.
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// Serial number of the certificate, Example: "123456789".
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// The subject name of this certificate.
        /// </summary>
        public readonly string Subject;
        /// <summary>
        /// The certificate thumbprint.
        /// </summary>
        public readonly string Thumbprint;
        /// <summary>
        /// Validation state of this certificate.
        /// </summary>
        public readonly string ValidationState;
        /// <summary>
        /// Certificate not valid at or after this timestamp.
        /// </summary>
        public readonly string ValidityExpirationTime;
        /// <summary>
        /// Certificate not valid before this timestamp.
        /// </summary>
        public readonly string ValidityStartTime;

        [OutputConstructor]
        private CertificateAttributesResponse(
            Outputs.CertificateTemplateResponse certificateTemplate,

            string fingerprint,

            string issuer,

            string serialNumber,

            string subject,

            string thumbprint,

            string validationState,

            string validityExpirationTime,

            string validityStartTime)
        {
            CertificateTemplate = certificateTemplate;
            Fingerprint = fingerprint;
            Issuer = issuer;
            SerialNumber = serialNumber;
            Subject = subject;
            Thumbprint = thumbprint;
            ValidationState = validationState;
            ValidityExpirationTime = validityExpirationTime;
            ValidityStartTime = validityStartTime;
        }
    }
}
