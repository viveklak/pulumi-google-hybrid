// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1
{
    public static class GetEntityType
    {
        /// <summary>
        /// Retrieves the specified entity type.
        /// </summary>
        public static Task<GetEntityTypeResult> InvokeAsync(GetEntityTypeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEntityTypeResult>("google-native:dialogflow/v3beta1:getEntityType", args ?? new GetEntityTypeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified entity type.
        /// </summary>
        public static Output<GetEntityTypeResult> Invoke(GetEntityTypeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEntityTypeResult>("google-native:dialogflow/v3beta1:getEntityType", args ?? new GetEntityTypeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEntityTypeArgs : Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public string AgentId { get; set; } = null!;

        [Input("entityTypeId", required: true)]
        public string EntityTypeId { get; set; } = null!;

        [Input("languageCode")]
        public string? LanguageCode { get; set; }

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetEntityTypeArgs()
        {
        }
    }

    public sealed class GetEntityTypeInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        [Input("entityTypeId", required: true)]
        public Input<string> EntityTypeId { get; set; } = null!;

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetEntityTypeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEntityTypeResult
    {
        /// <summary>
        /// Indicates whether the entity type can be automatically expanded.
        /// </summary>
        public readonly string AutoExpansionMode;
        /// <summary>
        /// The human-readable name of the entity type, unique within the agent.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Enables fuzzy entity extraction during classification.
        /// </summary>
        public readonly bool EnableFuzzyExtraction;
        /// <summary>
        /// The collection of entity entries associated with the entity type.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponse> Entities;
        /// <summary>
        /// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponse> ExcludedPhrases;
        /// <summary>
        /// Indicates the kind of entity type.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name during logging.
        /// </summary>
        public readonly bool Redact;

        [OutputConstructor]
        private GetEntityTypeResult(
            string autoExpansionMode,

            string displayName,

            bool enableFuzzyExtraction,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponse> entities,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponse> excludedPhrases,

            string kind,

            string name,

            bool redact)
        {
            AutoExpansionMode = autoExpansionMode;
            DisplayName = displayName;
            EnableFuzzyExtraction = enableFuzzyExtraction;
            Entities = entities;
            ExcludedPhrases = excludedPhrases;
            Kind = kind;
            Name = name;
            Redact = redact;
        }
    }
}
