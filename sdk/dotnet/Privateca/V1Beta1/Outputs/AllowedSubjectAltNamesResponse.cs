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
    /// AllowedSubjectAltNames specifies the allowed values for SubjectAltNames by the CertificateAuthority when issuing Certificates.
    /// </summary>
    [OutputType]
    public sealed class AllowedSubjectAltNamesResponse
    {
        /// <summary>
        /// Optional. Specifies if to allow custom X509Extension values.
        /// </summary>
        public readonly bool AllowCustomSans;
        /// <summary>
        /// Optional. Specifies if glob patterns used for allowed_dns_names allow wildcard certificates. If this is set, certificate requests with wildcard domains will be permitted to match a glob pattern specified in allowed_dns_names. Otherwise, certificate requests with wildcard domains will be permitted only if allowed_dns_names contains a literal wildcard.
        /// </summary>
        public readonly bool AllowGlobbingDnsWildcards;
        /// <summary>
        /// Optional. Contains valid, fully-qualified host names. Glob patterns are also supported. To allow an explicit wildcard certificate, escape with backlash (i.e. `\*`). E.g. for globbed entries: `*bar.com` will allow `foo.bar.com`, but not `*.bar.com`, unless the allow_globbing_dns_wildcards field is set. E.g. for wildcard entries: `\*.bar.com` will allow `*.bar.com`, but not `foo.bar.com`.
        /// </summary>
        public readonly ImmutableArray<string> AllowedDnsNames;
        /// <summary>
        /// Optional. Contains valid RFC 2822 E-mail addresses. Glob patterns are also supported.
        /// </summary>
        public readonly ImmutableArray<string> AllowedEmailAddresses;
        /// <summary>
        /// Optional. Contains valid 32-bit IPv4 addresses and subnet ranges or RFC 4291 IPv6 addresses and subnet ranges. Subnet ranges are specified using the '/' notation (e.g. 10.0.0.0/8, 2001:700:300:1800::/64). Glob patterns are supported only for ip address entries (i.e. not for subnet ranges).
        /// </summary>
        public readonly ImmutableArray<string> AllowedIps;
        /// <summary>
        /// Optional. Contains valid RFC 3986 URIs. Glob patterns are also supported. To match across path seperators (i.e. '/') use the double star glob pattern (i.e. '**').
        /// </summary>
        public readonly ImmutableArray<string> AllowedUris;

        [OutputConstructor]
        private AllowedSubjectAltNamesResponse(
            bool allowCustomSans,

            bool allowGlobbingDnsWildcards,

            ImmutableArray<string> allowedDnsNames,

            ImmutableArray<string> allowedEmailAddresses,

            ImmutableArray<string> allowedIps,

            ImmutableArray<string> allowedUris)
        {
            AllowCustomSans = allowCustomSans;
            AllowGlobbingDnsWildcards = allowGlobbingDnsWildcards;
            AllowedDnsNames = allowedDnsNames;
            AllowedEmailAddresses = allowedEmailAddresses;
            AllowedIps = allowedIps;
            AllowedUris = allowedUris;
        }
    }
}
