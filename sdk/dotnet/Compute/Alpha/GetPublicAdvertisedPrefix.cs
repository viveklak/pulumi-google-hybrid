// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetPublicAdvertisedPrefix
    {
        /// <summary>
        /// Returns the specified PublicAdvertisedPrefix resource.
        /// </summary>
        public static Task<GetPublicAdvertisedPrefixResult> InvokeAsync(GetPublicAdvertisedPrefixArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPublicAdvertisedPrefixResult>("google-native:compute/alpha:getPublicAdvertisedPrefix", args ?? new GetPublicAdvertisedPrefixArgs(), options.WithVersion());
    }


    public sealed class GetPublicAdvertisedPrefixArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("publicAdvertisedPrefix", required: true)]
        public string PublicAdvertisedPrefix { get; set; } = null!;

        public GetPublicAdvertisedPrefixArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPublicAdvertisedPrefixResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The IPv4 address to be used for reverse DNS verification.
        /// </summary>
        public readonly string DnsVerificationIp;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicAdvertisedPrefix. An up-to-date fingerprint must be provided in order to update the PublicAdvertisedPrefix, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a PublicAdvertisedPrefix.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// The IPv4 address range, in CIDR format, represented by this public advertised prefix.
        /// </summary>
        public readonly string IpCidrRange;
        /// <summary>
        /// Type of the resource. Always compute#publicAdvertisedPrefix for public advertised prefixes.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of public delegated prefixes that exist for this public advertised prefix.
        /// </summary>
        public readonly ImmutableArray<Outputs.PublicAdvertisedPrefixPublicDelegatedPrefixResponse> PublicDelegatedPrefixs;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL with id for the resource.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// The shared secret to be used for reverse DNS verification.
        /// </summary>
        public readonly string SharedSecret;
        /// <summary>
        /// The status of the public advertised prefix.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetPublicAdvertisedPrefixResult(
            string creationTimestamp,

            string description,

            string dnsVerificationIp,

            string fingerprint,

            string ipCidrRange,

            string kind,

            string name,

            ImmutableArray<Outputs.PublicAdvertisedPrefixPublicDelegatedPrefixResponse> publicDelegatedPrefixs,

            string selfLink,

            string selfLinkWithId,

            string sharedSecret,

            string status)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            DnsVerificationIp = dnsVerificationIp;
            Fingerprint = fingerprint;
            IpCidrRange = ipCidrRange;
            Kind = kind;
            Name = name;
            PublicDelegatedPrefixs = publicDelegatedPrefixs;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            SharedSecret = sharedSecret;
            Status = status;
        }
    }
}
