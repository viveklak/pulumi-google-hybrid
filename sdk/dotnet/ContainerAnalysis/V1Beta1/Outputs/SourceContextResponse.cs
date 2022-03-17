// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Outputs
{

    /// <summary>
    /// A SourceContext is a reference to a tree of files. A SourceContext together with a path point to a unique revision of a single file or directory.
    /// </summary>
    [OutputType]
    public sealed class SourceContextResponse
    {
        /// <summary>
        /// A SourceContext referring to a revision in a Google Cloud Source Repo.
        /// </summary>
        public readonly Outputs.CloudRepoSourceContextResponse CloudRepo;
        /// <summary>
        /// A SourceContext referring to a Gerrit project.
        /// </summary>
        public readonly Outputs.GerritSourceContextResponse Gerrit;
        /// <summary>
        /// A SourceContext referring to any third party Git repo (e.g., GitHub).
        /// </summary>
        public readonly Outputs.GitSourceContextResponse Git;
        /// <summary>
        /// Labels with user defined metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;

        [OutputConstructor]
        private SourceContextResponse(
            Outputs.CloudRepoSourceContextResponse cloudRepo,

            Outputs.GerritSourceContextResponse gerrit,

            Outputs.GitSourceContextResponse git,

            ImmutableDictionary<string, string> labels)
        {
            CloudRepo = cloudRepo;
            Gerrit = gerrit;
            Git = git;
            Labels = labels;
        }
    }
}
