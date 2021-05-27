// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Alpha
{
    /// <summary>
    /// Uploads the specified SSL certificate.
    /// </summary>
    [GoogleNativeResourceType("google-native:appengine/v1alpha:AppAuthorizedCertificate")]
    public partial class AppAuthorizedCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
        /// </summary>
        [Output("certificateRawData")]
        public Output<Outputs.CertificateRawDataResponse> CertificateRawData { get; private set; } = null!;

        /// <summary>
        /// The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
        /// </summary>
        [Output("domainMappingsCount")]
        public Output<int> DomainMappingsCount { get; private set; } = null!;

        /// <summary>
        /// Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.@OutputOnly
        /// </summary>
        [Output("domainNames")]
        public Output<ImmutableArray<string>> DomainNames { get; private set; } = null!;

        /// <summary>
        /// The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.@OutputOnly
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.@OutputOnly
        /// </summary>
        [Output("managedCertificate")]
        public Output<Outputs.ManagedCertificateResponse> ManagedCertificate { get; private set; } = null!;

        /// <summary>
        /// Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.@OutputOnly
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
        /// </summary>
        [Output("visibleDomainMappings")]
        public Output<ImmutableArray<string>> VisibleDomainMappings { get; private set; } = null!;


        /// <summary>
        /// Create a AppAuthorizedCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppAuthorizedCertificate(string name, AppAuthorizedCertificateArgs args, CustomResourceOptions? options = null)
            : base("google-native:appengine/v1alpha:AppAuthorizedCertificate", name, args ?? new AppAuthorizedCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppAuthorizedCertificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:appengine/v1alpha:AppAuthorizedCertificate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AppAuthorizedCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppAuthorizedCertificate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppAuthorizedCertificate(name, id, options);
        }
    }

    public sealed class AppAuthorizedCertificateArgs : Pulumi.ResourceArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
        /// </summary>
        [Input("certificateRawData")]
        public Input<Inputs.CertificateRawDataArgs>? CertificateRawData { get; set; }

        /// <summary>
        /// The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
        /// </summary>
        [Input("domainMappingsCount")]
        public Input<int>? DomainMappingsCount { get; set; }

        [Input("domainNames")]
        private InputList<string>? _domainNames;

        /// <summary>
        /// Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.@OutputOnly
        /// </summary>
        public InputList<string> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<string>());
            set => _domainNames = value;
        }

        /// <summary>
        /// The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.@OutputOnly
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// Relative name of the certificate. This is a unique value autogenerated on AuthorizedCertificate resource creation. Example: 12345.@OutputOnly
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.@OutputOnly
        /// </summary>
        [Input("managedCertificate")]
        public Input<Inputs.ManagedCertificateArgs>? ManagedCertificate { get; set; }

        /// <summary>
        /// Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.@OutputOnly
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("visibleDomainMappings")]
        private InputList<string>? _visibleDomainMappings;

        /// <summary>
        /// The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.@OutputOnly
        /// </summary>
        public InputList<string> VisibleDomainMappings
        {
            get => _visibleDomainMappings ?? (_visibleDomainMappings = new InputList<string>());
            set => _visibleDomainMappings = value;
        }

        public AppAuthorizedCertificateArgs()
        {
        }
    }
}
