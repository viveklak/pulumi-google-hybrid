// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Inputs
{

    /// <summary>
    /// GitRepoSource describes a repo and ref of a code repository.
    /// </summary>
    public sealed class GitRepoSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full resource name of the bitbucket server config. Format: `projects/{project}/locations/{location}/bitbucketServerConfigs/{id}`.
        /// </summary>
        [Input("bitbucketServerConfig")]
        public Input<string>? BitbucketServerConfig { get; set; }

        /// <summary>
        /// The full resource name of the github enterprise config. Format: `projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}`. `projects/{project}/githubEnterpriseConfigs/{id}`.
        /// </summary>
        [Input("githubEnterpriseConfig")]
        public Input<string>? GithubEnterpriseConfig { get; set; }

        /// <summary>
        /// The branch or tag to use. Must start with "refs/" (required).
        /// </summary>
        [Input("ref")]
        public Input<string>? Ref { get; set; }

        /// <summary>
        /// See RepoType below.
        /// </summary>
        [Input("repoType")]
        public Input<Pulumi.GoogleNative.CloudBuild.V1.GitRepoSourceRepoType>? RepoType { get; set; }

        /// <summary>
        /// The URI of the repo (required).
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public GitRepoSourceArgs()
        {
        }
    }
}
