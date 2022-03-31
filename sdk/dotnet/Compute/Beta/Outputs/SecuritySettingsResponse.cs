// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// The authentication and authorization settings for a BackendService.
    /// </summary>
    [OutputType]
    public sealed class SecuritySettingsResponse
    {
        /// <summary>
        /// [Deprecated] Use clientTlsPolicy instead.
        /// </summary>
        public readonly string Authentication;
        /// <summary>
        /// Optional. A URL referring to a networksecurity.ClientTlsPolicy resource that describes how clients should authenticate with this service's backends. clientTlsPolicy only applies to a global BackendService with the loadBalancingScheme set to INTERNAL_SELF_MANAGED. If left blank, communications are not encrypted. Note: This field currently has no impact.
        /// </summary>
        public readonly string ClientTlsPolicy;
        /// <summary>
        /// Optional. A list of Subject Alternative Names (SANs) that the client verifies during a mutual TLS handshake with an server/endpoint for this BackendService. When the server presents its X.509 certificate to the client, the client inspects the certificate's subjectAltName field. If the field contains one of the specified values, the communication continues. Otherwise, it fails. This additional check enables the client to verify that the server is authorized to run the requested service. Note that the contents of the server certificate's subjectAltName field are configured by the Public Key Infrastructure which provisions server identities. Only applies to a global BackendService with loadBalancingScheme set to INTERNAL_SELF_MANAGED. Only applies when BackendService has an attached clientTlsPolicy with clientCertificate (mTLS mode). Note: This field currently has no impact.
        /// </summary>
        public readonly ImmutableArray<string> SubjectAltNames;

        [OutputConstructor]
        private SecuritySettingsResponse(
            string authentication,

            string clientTlsPolicy,

            ImmutableArray<string> subjectAltNames)
        {
            Authentication = authentication;
            ClientTlsPolicy = clientTlsPolicy;
            SubjectAltNames = subjectAltNames;
        }
    }
}
