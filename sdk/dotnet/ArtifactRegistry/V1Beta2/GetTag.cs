// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ArtifactRegistry.V1Beta2
{
    public static class GetTag
    {
        /// <summary>
        /// Gets a tag.
        /// </summary>
        public static Task<GetTagResult> InvokeAsync(GetTagArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagResult>("google-native:artifactregistry/v1beta2:getTag", args ?? new GetTagArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a tag.
        /// </summary>
        public static Output<GetTagResult> Invoke(GetTagInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTagResult>("google-native:artifactregistry/v1beta2:getTag", args ?? new GetTagInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("packageId", required: true)]
        public string PackageId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("repositoryId", required: true)]
        public string RepositoryId { get; set; } = null!;

        [Input("tagId", required: true)]
        public string TagId { get; set; } = null!;

        public GetTagArgs()
        {
        }
    }

    public sealed class GetTagInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("packageId", required: true)]
        public Input<string> PackageId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        [Input("tagId", required: true)]
        public Input<string> TagId { get; set; } = null!;

        public GetTagInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTagResult
    {
        /// <summary>
        /// The name of the tag, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/tags/tag1". If the package part contains slashes, the slashes are escaped. The tag part can only have characters in [a-zA-Z0-9\-._~:@], anything else must be URL encoded.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the version the tag refers to, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/sha256:5243811" If the package or version ID parts contain slashes, the slashes are escaped.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetTagResult(
            string name,

            string version)
        {
            Name = name;
            Version = version;
        }
    }
}
