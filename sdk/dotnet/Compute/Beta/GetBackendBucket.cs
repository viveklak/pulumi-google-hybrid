// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetBackendBucket
    {
        /// <summary>
        /// Returns the specified BackendBucket resource. Gets a list of available backend buckets by making a list() request.
        /// </summary>
        public static Task<GetBackendBucketResult> InvokeAsync(GetBackendBucketArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackendBucketResult>("google-native:compute/beta:getBackendBucket", args ?? new GetBackendBucketArgs(), options.WithVersion());
    }


    public sealed class GetBackendBucketArgs : Pulumi.InvokeArgs
    {
        [Input("backendBucket", required: true)]
        public string BackendBucket { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetBackendBucketArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackendBucketResult
    {
        /// <summary>
        /// Cloud Storage bucket name.
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// Cloud CDN configuration for this BackendBucket.
        /// </summary>
        public readonly Outputs.BackendBucketCdnPolicyResponse CdnPolicy;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied responses.
        /// </summary>
        public readonly ImmutableArray<string> CustomResponseHeaders;
        /// <summary>
        /// An optional textual description of the resource; provided by the client when the resource is created.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// If true, enable Cloud CDN for this BackendBucket.
        /// </summary>
        public readonly bool EnableCdn;
        /// <summary>
        /// Type of the resource.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;

        [OutputConstructor]
        private GetBackendBucketResult(
            string bucketName,

            Outputs.BackendBucketCdnPolicyResponse cdnPolicy,

            string creationTimestamp,

            ImmutableArray<string> customResponseHeaders,

            string description,

            bool enableCdn,

            string kind,

            string name,

            string selfLink)
        {
            BucketName = bucketName;
            CdnPolicy = cdnPolicy;
            CreationTimestamp = creationTimestamp;
            CustomResponseHeaders = customResponseHeaders;
            Description = description;
            EnableCdn = enableCdn;
            Kind = kind;
            Name = name;
            SelfLink = selfLink;
        }
    }
}
