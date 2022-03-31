// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta
{
    public static class GetAuthorizedCertificate
    {
        /// <summary>
        /// Gets the specified SSL certificate.
        /// </summary>
        public static Task<GetAuthorizedCertificateResult> InvokeAsync(GetAuthorizedCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizedCertificateResult>("google-native:appengine/v1beta:getAuthorizedCertificate", args ?? new GetAuthorizedCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified SSL certificate.
        /// </summary>
        public static Output<GetAuthorizedCertificateResult> Invoke(GetAuthorizedCertificateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAuthorizedCertificateResult>("google-native:appengine/v1beta:getAuthorizedCertificate", args ?? new GetAuthorizedCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthorizedCertificateArgs : Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        [Input("authorizedCertificateId", required: true)]
        public string AuthorizedCertificateId { get; set; } = null!;

        [Input("view")]
        public string? View { get; set; }

        public GetAuthorizedCertificateArgs()
        {
        }
    }

    public sealed class GetAuthorizedCertificateInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("authorizedCertificateId", required: true)]
        public Input<string> AuthorizedCertificateId { get; set; } = null!;

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetAuthorizedCertificateInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthorizedCertificateResult
    {
        /// <summary>
        /// The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
        /// </summary>
        public readonly Outputs.CertificateRawDataResponse CertificateRawData;
        /// <summary>
        /// The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.
        /// </summary>
        public readonly int DomainMappingsCount;
        /// <summary>
        /// Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.
        /// </summary>
        public readonly ImmutableArray<string> DomainNames;
        /// <summary>
        /// The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.
        /// </summary>
        public readonly Outputs.ManagedCertificateResponse ManagedCertificate;
        /// <summary>
        /// Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.
        /// </summary>
        public readonly ImmutableArray<string> VisibleDomainMappings;

        [OutputConstructor]
        private GetAuthorizedCertificateResult(
            Outputs.CertificateRawDataResponse certificateRawData,

            string displayName,

            int domainMappingsCount,

            ImmutableArray<string> domainNames,

            string expireTime,

            Outputs.ManagedCertificateResponse managedCertificate,

            string name,

            ImmutableArray<string> visibleDomainMappings)
        {
            CertificateRawData = certificateRawData;
            DisplayName = displayName;
            DomainMappingsCount = domainMappingsCount;
            DomainNames = domainNames;
            ExpireTime = expireTime;
            ManagedCertificate = managedCertificate;
            Name = name;
            VisibleDomainMappings = visibleDomainMappings;
        }
    }
}
