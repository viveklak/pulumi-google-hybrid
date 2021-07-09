// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1
{
    /// <summary>
    /// Creates an entity type in the specified agent.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v3beta1:EntityType")]
    public partial class EntityType : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether the entity type can be automatically expanded.
        /// </summary>
        [Output("autoExpansionMode")]
        public Output<string> AutoExpansionMode { get; private set; } = null!;

        /// <summary>
        /// The human-readable name of the entity type, unique within the agent.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Enables fuzzy entity extraction during classification.
        /// </summary>
        [Output("enableFuzzyExtraction")]
        public Output<bool> EnableFuzzyExtraction { get; private set; } = null!;

        /// <summary>
        /// The collection of entity entries associated with the entity type.
        /// </summary>
        [Output("entities")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponse>> Entities { get; private set; } = null!;

        /// <summary>
        /// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
        /// </summary>
        [Output("excludedPhrases")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponse>> ExcludedPhrases { get; private set; } = null!;

        /// <summary>
        /// Indicates the kind of entity type.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name during logging.
        /// </summary>
        [Output("redact")]
        public Output<bool> Redact { get; private set; } = null!;


        /// <summary>
        /// Create a EntityType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EntityType(string name, EntityTypeArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3beta1:EntityType", name, args ?? new EntityTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EntityType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3beta1:EntityType", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing EntityType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EntityType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EntityType(name, id, options);
        }
    }

    public sealed class EntityTypeArgs : Pulumi.ResourceArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        /// <summary>
        /// Indicates whether the entity type can be automatically expanded.
        /// </summary>
        [Input("autoExpansionMode")]
        public Input<Pulumi.GoogleNative.Dialogflow.V3Beta1.EntityTypeAutoExpansionMode>? AutoExpansionMode { get; set; }

        /// <summary>
        /// The human-readable name of the entity type, unique within the agent.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Enables fuzzy entity extraction during classification.
        /// </summary>
        [Input("enableFuzzyExtraction")]
        public Input<bool>? EnableFuzzyExtraction { get; set; }

        [Input("entities")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs>? _entities;

        /// <summary>
        /// The collection of entity entries associated with the entity type.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs> Entities
        {
            get => _entities ?? (_entities = new InputList<Inputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs>());
            set => _entities = value;
        }

        [Input("excludedPhrases")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseArgs>? _excludedPhrases;

        /// <summary>
        /// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseArgs> ExcludedPhrases
        {
            get => _excludedPhrases ?? (_excludedPhrases = new InputList<Inputs.GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseArgs>());
            set => _excludedPhrases = value;
        }

        /// <summary>
        /// Indicates the kind of entity type.
        /// </summary>
        [Input("kind", required: true)]
        public Input<Pulumi.GoogleNative.Dialogflow.V3Beta1.EntityTypeKind> Kind { get; set; } = null!;

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name during logging.
        /// </summary>
        [Input("redact")]
        public Input<bool>? Redact { get; set; }

        public EntityTypeArgs()
        {
        }
    }
}
