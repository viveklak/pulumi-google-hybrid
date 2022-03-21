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
    /// GitFileSource describes a file within a (possibly remote) code repository.
    /// </summary>
    public sealed class GitFileSourceArgs : Pulumi.ResourceArgs
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
        /// The path of the file, with the repo root as the root of the path.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// See RepoType above.
        /// </summary>
        [Input("repoType")]
        public Input<Pulumi.GoogleNative.CloudBuild.V1.GitFileSourceRepoType>? RepoType { get; set; }

        /// <summary>
        /// The branch, tag, arbitrary ref, or SHA version of the repo to use when resolving the filename (optional). This field respects the same syntax/resolution as described here: https://git-scm.com/docs/gitrevisions If unspecified, the revision from which the trigger invocation originated is assumed to be the revision from which to read the specified path.
        /// </summary>
        [Input("revision")]
        public Input<string>? Revision { get; set; }

        /// <summary>
        /// The URI of the repo (optional). If unspecified, the repo from which the trigger invocation originated is assumed to be the repo from which to read the specified path.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public GitFileSourceArgs()
        {
        }
    }
}
