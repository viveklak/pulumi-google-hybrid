// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.ArtifactRegistry.V1Beta1
{
    public static class GetRepository
    {
        /// <summary>
        /// Gets a repository.
        /// </summary>
        public static Task<GetRepositoryResult> InvokeAsync(GetRepositoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryResult>("google-native:artifactregistry/v1beta1:getRepository", args ?? new GetRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a repository.
        /// </summary>
        public static Output<GetRepositoryResult> Invoke(GetRepositoryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRepositoryResult>("google-native:artifactregistry/v1beta1:getRepository", args ?? new GetRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("repositoryId", required: true)]
        public string RepositoryId { get; set; } = null!;

        public GetRepositoryArgs()
        {
        }
    }

    public sealed class GetRepositoryInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        public GetRepositoryInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRepositoryResult
    {
        /// <summary>
        /// The time when the repository was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The user-provided description of the repository.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The format of packages that are stored in the repository.
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// The Cloud KMS resource name of the customer managed encryption key that’s used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
        /// </summary>
        public readonly string KmsKeyName;
        /// <summary>
        /// Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The time when the repository was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetRepositoryResult(
            string createTime,

            string description,

            string format,

            string kmsKeyName,

            ImmutableDictionary<string, string> labels,

            string name,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Format = format;
            KmsKeyName = kmsKeyName;
            Labels = labels;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}
