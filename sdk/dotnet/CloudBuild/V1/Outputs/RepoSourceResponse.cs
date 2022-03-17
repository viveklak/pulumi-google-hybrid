// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Outputs
{

    /// <summary>
    /// Location of the source in a Google Cloud Source Repository.
    /// </summary>
    [OutputType]
    public sealed class RepoSourceResponse
    {
        /// <summary>
        /// Regex matching branches to build. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax
        /// </summary>
        public readonly string BranchName;
        /// <summary>
        /// Explicit commit SHA to build.
        /// </summary>
        public readonly string CommitSha;
        /// <summary>
        /// Directory, relative to the source root, in which to run the build. This must be a relative path. If a step's `dir` is specified and is an absolute path, this value is ignored for that step's execution.
        /// </summary>
        public readonly string Dir;
        /// <summary>
        /// Only trigger a build if the revision regex does NOT match the revision regex.
        /// </summary>
        public readonly bool InvertRegex;
        /// <summary>
        /// ID of the project that owns the Cloud Source Repository. If omitted, the project ID requesting the build is assumed.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Name of the Cloud Source Repository.
        /// </summary>
        public readonly string RepoName;
        /// <summary>
        /// Substitutions to use in a triggered build. Should only be used with RunBuildTrigger
        /// </summary>
        public readonly ImmutableDictionary<string, string> Substitutions;
        /// <summary>
        /// Regex matching tags to build. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax
        /// </summary>
        public readonly string TagName;

        [OutputConstructor]
        private RepoSourceResponse(
            string branchName,

            string commitSha,

            string dir,

            bool invertRegex,

            string project,

            string repoName,

            ImmutableDictionary<string, string> substitutions,

            string tagName)
        {
            BranchName = branchName;
            CommitSha = commitSha;
            Dir = dir;
            InvertRegex = invertRegex;
            Project = project;
            RepoName = repoName;
            Substitutions = substitutions;
            TagName = tagName;
        }
    }
}
