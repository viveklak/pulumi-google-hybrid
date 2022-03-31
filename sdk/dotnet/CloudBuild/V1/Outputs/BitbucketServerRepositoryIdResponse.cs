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
    /// BitbucketServerRepositoryId identifies a specific repository hosted on a Bitbucket Server.
    /// </summary>
    [OutputType]
    public sealed class BitbucketServerRepositoryIdResponse
    {
        /// <summary>
        /// Identifier for the project storing the repository.
        /// </summary>
        public readonly string ProjectKey;
        /// <summary>
        /// Identifier for the repository.
        /// </summary>
        public readonly string RepoSlug;
        /// <summary>
        /// The ID of the webhook that was created for receiving events from this repo. We only create and manage a single webhook for each repo.
        /// </summary>
        public readonly int WebhookId;

        [OutputConstructor]
        private BitbucketServerRepositoryIdResponse(
            string projectKey,

            string repoSlug,

            int webhookId)
        {
            ProjectKey = projectKey;
            RepoSlug = repoSlug;
            WebhookId = webhookId;
        }
    }
}
