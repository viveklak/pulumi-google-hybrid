// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// RepoSource describes the location of the source in a Google Cloud Source Repository.
    /// </summary>
    public sealed class RepoSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the branch to build.
        /// </summary>
        [Input("branchName")]
        public Input<string>? BranchName { get; set; }

        /// <summary>
        /// Explicit commit SHA to build.
        /// </summary>
        [Input("commitSha")]
        public Input<string>? CommitSha { get; set; }

        /// <summary>
        /// ID of the project that owns the repo.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Name of the repo.
        /// </summary>
        [Input("repoName")]
        public Input<string>? RepoName { get; set; }

        /// <summary>
        /// Name of the tag to build.
        /// </summary>
        [Input("tagName")]
        public Input<string>? TagName { get; set; }

        public RepoSourceArgs()
        {
        }
    }
}
