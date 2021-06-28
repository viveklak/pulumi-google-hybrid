// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetPublicDelegatedPrefix
    {
        /// <summary>
        /// Returns the specified PublicDelegatedPrefix resource in the given region.
        /// </summary>
        public static Task<GetPublicDelegatedPrefixResult> InvokeAsync(GetPublicDelegatedPrefixArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPublicDelegatedPrefixResult>("google-native:compute/alpha:getPublicDelegatedPrefix", args ?? new GetPublicDelegatedPrefixArgs(), options.WithVersion());
    }


    public sealed class GetPublicDelegatedPrefixArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("publicDelegatedPrefix", required: true)]
        public string PublicDelegatedPrefix { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetPublicDelegatedPrefixArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPublicDelegatedPrefixResult
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
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
        /// </summary>
        public readonly string IpCidrRange;
        /// <summary>
        /// If true, the prefix will be live migrated.
        /// </summary>
        public readonly bool IsLiveMigration;
        /// <summary>
        /// Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
        /// </summary>
        public readonly string ParentPrefix;
        /// <summary>
        /// The list of sub public delegated prefixes that exist for this public delegated prefix.
        /// </summary>
        public readonly ImmutableArray<Outputs.PublicDelegatedPrefixPublicDelegatedSubPrefixResponse> PublicDelegatedSubPrefixs;
        /// <summary>
        /// URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL with id for the resource.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// The status of the public delegated prefix.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetPublicDelegatedPrefixResult(
            string creationTimestamp,

            string description,

            string fingerprint,

            string ipCidrRange,

            bool isLiveMigration,

            string kind,

            string name,

            string parentPrefix,

            ImmutableArray<Outputs.PublicDelegatedPrefixPublicDelegatedSubPrefixResponse> publicDelegatedSubPrefixs,

            string region,

            string selfLink,

            string selfLinkWithId,

            string status)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Fingerprint = fingerprint;
            IpCidrRange = ipCidrRange;
            IsLiveMigration = isLiveMigration;
            Kind = kind;
            Name = name;
            ParentPrefix = parentPrefix;
            PublicDelegatedSubPrefixs = publicDelegatedSubPrefixs;
            Region = region;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            Status = status;
        }
    }
}
