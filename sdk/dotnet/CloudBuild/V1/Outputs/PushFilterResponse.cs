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
    /// Push contains filter properties for matching GitHub git pushes.
    /// </summary>
    [OutputType]
    public sealed class PushFilterResponse
    {
        /// <summary>
        /// Regexes matching branches to build. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax
        /// </summary>
        public readonly string Branch;
        /// <summary>
        /// When true, only trigger a build if the revision regex does NOT match the git_ref regex.
        /// </summary>
        public readonly bool InvertRegex;
        /// <summary>
        /// Regexes matching tags to build. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax
        /// </summary>
        public readonly string Tag;

        [OutputConstructor]
        private PushFilterResponse(
            string branch,

            bool invertRegex,

            string tag)
        {
            Branch = branch;
            InvertRegex = invertRegex;
            Tag = tag;
        }
    }
}
