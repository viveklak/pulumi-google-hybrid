// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1
{
    public static class GetKnowledgeBase
    {
        /// <summary>
        /// Retrieves the specified knowledge base. Note: The `projects.agent.knowledgeBases` resource is deprecated; only use `projects.knowledgeBases`.
        /// </summary>
        public static Task<GetKnowledgeBaseResult> InvokeAsync(GetKnowledgeBaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKnowledgeBaseResult>("google-native:dialogflow/v2beta1:getKnowledgeBase", args ?? new GetKnowledgeBaseArgs(), options.WithVersion());
    }


    public sealed class GetKnowledgeBaseArgs : Pulumi.InvokeArgs
    {
        [Input("knowledgeBaseId", required: true)]
        public string KnowledgeBaseId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetKnowledgeBaseArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKnowledgeBaseResult
    {
        /// <summary>
        /// The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, this is populated for all non en-us languages. If not populated, the default language en-us applies.
        /// </summary>
        public readonly string LanguageCode;
        /// <summary>
        /// The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetKnowledgeBaseResult(
            string displayName,

            string languageCode,

            string name)
        {
            DisplayName = displayName;
            LanguageCode = languageCode;
            Name = name;
        }
    }
}
