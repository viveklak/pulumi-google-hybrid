// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V2Beta.Inputs
{

    /// <summary>
    /// Location of the source in a Google Cloud Source Repository.
    /// </summary>
    public sealed class RepoSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Regex matching branches to build. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax
        /// </summary>
        [Input("branchName")]
        public Input<string>? BranchName { get; set; }

        /// <summary>
        /// Explicit commit SHA to build.
        /// </summary>
        [Input("commitSha")]
        public Input<string>? CommitSha { get; set; }

        /// <summary>
        /// Directory, relative to the source root, in which to run the build. This must be a relative path. If a step's `dir` is specified and is an absolute path, this value is ignored for that step's execution. eg. helloworld (no leading slash allowed)
        /// </summary>
        [Input("dir")]
        public Input<string>? Dir { get; set; }

        /// <summary>
        /// Only trigger a build if the revision regex does NOT match the revision regex.
        /// </summary>
        [Input("invertRegex")]
        public Input<bool>? InvertRegex { get; set; }

        /// <summary>
        /// ID of the project that owns the Cloud Source Repository. If omitted, the project ID requesting the build is assumed.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Name of the Cloud Source Repository.
        /// </summary>
        [Input("repoName")]
        public Input<string>? RepoName { get; set; }

        /// <summary>
        /// Regex matching tags to build. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax
        /// </summary>
        [Input("tagName")]
        public Input<string>? TagName { get; set; }

        public RepoSourceArgs()
        {
        }
    }
}
