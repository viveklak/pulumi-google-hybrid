// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Dialogflow.V2Beta1
{
    /// <summary>
    /// Creates a knowledge base. Note: The `projects.agent.knowledgeBases` resource is deprecated; only use `projects.knowledgeBases`.
    /// </summary>
    [GcpNativeResourceType("gcp-native:dialogflow/v2beta1:AgentKnowledgeBase")]
    public partial class AgentKnowledgeBase : Pulumi.CustomResource
    {
        /// <summary>
        /// Required. The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, this is populated for all non en-us languages. If not populated, the default language en-us applies.
        /// </summary>
        [Output("languageCode")]
        public Output<string> LanguageCode { get; private set; } = null!;

        /// <summary>
        /// The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a AgentKnowledgeBase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AgentKnowledgeBase(string name, AgentKnowledgeBaseArgs args, CustomResourceOptions? options = null)
            : base("gcp-native:dialogflow/v2beta1:AgentKnowledgeBase", name, args ?? new AgentKnowledgeBaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AgentKnowledgeBase(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("gcp-native:dialogflow/v2beta1:AgentKnowledgeBase", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AgentKnowledgeBase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AgentKnowledgeBase Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AgentKnowledgeBase(name, id, options);
        }
    }

    public sealed class AgentKnowledgeBaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("knowledgeBasesId", required: true)]
        public Input<string> KnowledgeBasesId { get; set; } = null!;

        /// <summary>
        /// Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, this is populated for all non en-us languages. If not populated, the default language en-us applies.
        /// </summary>
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        public AgentKnowledgeBaseArgs()
        {
        }
    }
}