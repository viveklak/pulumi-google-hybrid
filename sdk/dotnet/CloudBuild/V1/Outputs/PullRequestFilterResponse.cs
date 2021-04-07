// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.CloudBuild.V1.Outputs
{

    [OutputType]
    public sealed class PullRequestFilterResponse
    {
        /// <summary>
        /// Regex of branches to match. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax
        /// </summary>
        public readonly string Branch;
        /// <summary>
        /// Configure builds to run whether a repository owner or collaborator need to comment `/gcbrun`.
        /// </summary>
        public readonly string CommentControl;
        /// <summary>
        /// If true, branches that do NOT match the git_ref will trigger a build.
        /// </summary>
        public readonly bool InvertRegex;

        [OutputConstructor]
        private PullRequestFilterResponse(
            string branch,

            string commentControl,

            bool invertRegex)
        {
            Branch = branch;
            CommentControl = commentControl;
            InvertRegex = invertRegex;
        }
    }
}