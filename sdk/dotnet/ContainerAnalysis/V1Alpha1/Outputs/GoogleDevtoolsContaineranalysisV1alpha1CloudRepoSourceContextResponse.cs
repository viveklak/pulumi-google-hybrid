// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// A CloudRepoSourceContext denotes a particular revision in a Google Cloud Source Repo.
    /// </summary>
    [OutputType]
    public sealed class GoogleDevtoolsContaineranalysisV1alpha1CloudRepoSourceContextResponse
    {
        /// <summary>
        /// An alias, which may be a branch or tag.
        /// </summary>
        public readonly Outputs.GoogleDevtoolsContaineranalysisV1alpha1AliasContextResponse AliasContext;
        /// <summary>
        /// The ID of the repo.
        /// </summary>
        public readonly Outputs.GoogleDevtoolsContaineranalysisV1alpha1RepoIdResponse RepoId;
        /// <summary>
        /// A revision ID.
        /// </summary>
        public readonly string RevisionId;

        [OutputConstructor]
        private GoogleDevtoolsContaineranalysisV1alpha1CloudRepoSourceContextResponse(
            Outputs.GoogleDevtoolsContaineranalysisV1alpha1AliasContextResponse aliasContext,

            Outputs.GoogleDevtoolsContaineranalysisV1alpha1RepoIdResponse repoId,

            string revisionId)
        {
            AliasContext = aliasContext;
            RepoId = repoId;
            RevisionId = revisionId;
        }
    }
}
