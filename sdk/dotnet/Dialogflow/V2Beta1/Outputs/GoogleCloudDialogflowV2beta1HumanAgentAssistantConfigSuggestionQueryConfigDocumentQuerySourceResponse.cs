// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Outputs
{

    /// <summary>
    /// Document source settings. Supported features: SMART_REPLY, SMART_COMPOSE.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionQueryConfigDocumentQuerySourceResponse
    {
        /// <summary>
        /// Knowledge documents to query from. Format: `projects//locations//knowledgeBases//documents/`. Currently, only one document is supported.
        /// </summary>
        public readonly ImmutableArray<string> Documents;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionQueryConfigDocumentQuerySourceResponse(ImmutableArray<string> documents)
        {
            Documents = documents;
        }
    }
}