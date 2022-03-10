// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CertificateManager.V1.Inputs
{

    /// <summary>
    /// Configuration and state of a Managed Certificate. Certificate Manager provisions and renews Managed Certificates automatically, for as long as it's authorized to do so.
    /// </summary>
    public sealed class ManagedCertificateArgs : Pulumi.ResourceArgs
    {
        [Input("dnsAuthorizations")]
        private InputList<string>? _dnsAuthorizations;

        /// <summary>
        /// Immutable. Authorizations that will be used for performing domain authorization.
        /// </summary>
        public InputList<string> DnsAuthorizations
        {
            get => _dnsAuthorizations ?? (_dnsAuthorizations = new InputList<string>());
            set => _dnsAuthorizations = value;
        }

        [Input("domains")]
        private InputList<string>? _domains;

        /// <summary>
        /// Immutable. The domains for which a managed SSL certificate will be generated. Wildcard domains are only supported with DNS challenge resolution.
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        /// <summary>
        /// Information about issues with provisioning a Managed Certificate.
        /// </summary>
        [Input("provisioningIssue")]
        public Input<Inputs.ProvisioningIssueArgs>? ProvisioningIssue { get; set; }

        public ManagedCertificateArgs()
        {
        }
    }
}