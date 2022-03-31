// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Outputs
{

    /// <summary>
    /// A rich response message. Corresponds to the intent `Response` field in the Dialogflow console. For more information, see [Rich response messages](https://cloud.google.com/dialogflow/docs/intents-rich-messages).
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2IntentMessageResponse
    {
        /// <summary>
        /// The basic card response for Actions on Google.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageBasicCardResponse BasicCard;
        /// <summary>
        /// Browse carousel card for Actions on Google.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageBrowseCarouselCardResponse BrowseCarouselCard;
        /// <summary>
        /// The card response.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageCardResponse Card;
        /// <summary>
        /// The carousel card response for Actions on Google.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageCarouselSelectResponse CarouselSelect;
        /// <summary>
        /// The image response.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageImageResponse Image;
        /// <summary>
        /// The link out suggestion chip for Actions on Google.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageLinkOutSuggestionResponse LinkOutSuggestion;
        /// <summary>
        /// The list card response for Actions on Google.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageListSelectResponse ListSelect;
        /// <summary>
        /// The media content card for Actions on Google.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageMediaContentResponse MediaContent;
        /// <summary>
        /// A custom platform-specific response.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Payload;
        /// <summary>
        /// Optional. The platform that this message is intended for.
        /// </summary>
        public readonly string Platform;
        /// <summary>
        /// The quick replies response.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageQuickRepliesResponse QuickReplies;
        /// <summary>
        /// The voice and text-only responses for Actions on Google.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageSimpleResponsesResponse SimpleResponses;
        /// <summary>
        /// The suggestion chips for Actions on Google.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageSuggestionsResponse Suggestions;
        /// <summary>
        /// Table card for Actions on Google.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageTableCardResponse TableCard;
        /// <summary>
        /// The text response.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageTextResponse Text;

        [OutputConstructor]
        private GoogleCloudDialogflowV2IntentMessageResponse(
            Outputs.GoogleCloudDialogflowV2IntentMessageBasicCardResponse basicCard,

            Outputs.GoogleCloudDialogflowV2IntentMessageBrowseCarouselCardResponse browseCarouselCard,

            Outputs.GoogleCloudDialogflowV2IntentMessageCardResponse card,

            Outputs.GoogleCloudDialogflowV2IntentMessageCarouselSelectResponse carouselSelect,

            Outputs.GoogleCloudDialogflowV2IntentMessageImageResponse image,

            Outputs.GoogleCloudDialogflowV2IntentMessageLinkOutSuggestionResponse linkOutSuggestion,

            Outputs.GoogleCloudDialogflowV2IntentMessageListSelectResponse listSelect,

            Outputs.GoogleCloudDialogflowV2IntentMessageMediaContentResponse mediaContent,

            ImmutableDictionary<string, string> payload,

            string platform,

            Outputs.GoogleCloudDialogflowV2IntentMessageQuickRepliesResponse quickReplies,

            Outputs.GoogleCloudDialogflowV2IntentMessageSimpleResponsesResponse simpleResponses,

            Outputs.GoogleCloudDialogflowV2IntentMessageSuggestionsResponse suggestions,

            Outputs.GoogleCloudDialogflowV2IntentMessageTableCardResponse tableCard,

            Outputs.GoogleCloudDialogflowV2IntentMessageTextResponse text)
        {
            BasicCard = basicCard;
            BrowseCarouselCard = browseCarouselCard;
            Card = card;
            CarouselSelect = carouselSelect;
            Image = image;
            LinkOutSuggestion = linkOutSuggestion;
            ListSelect = listSelect;
            MediaContent = mediaContent;
            Payload = payload;
            Platform = platform;
            QuickReplies = quickReplies;
            SimpleResponses = simpleResponses;
            Suggestions = suggestions;
            TableCard = tableCard;
            Text = text;
        }
    }
}
