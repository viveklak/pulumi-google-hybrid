# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ConversationConversationStage',
    'DocumentKnowledgeTypesItem',
    'EntityTypeAutoExpansionMode',
    'EntityTypeKind',
    'GoogleCloudDialogflowV2ArticleSuggestionModelMetadataTrainingModelType',
    'GoogleCloudDialogflowV2FulfillmentFeatureType',
    'GoogleCloudDialogflowV2IntentMessageBrowseCarouselCardBrowseCarouselCardItemOpenUrlActionUrlTypeHint',
    'GoogleCloudDialogflowV2IntentMessageBrowseCarouselCardImageDisplayOptions',
    'GoogleCloudDialogflowV2IntentMessageColumnPropertiesHorizontalAlignment',
    'GoogleCloudDialogflowV2IntentMessageMediaContentMediaType',
    'GoogleCloudDialogflowV2IntentMessagePlatform',
    'GoogleCloudDialogflowV2IntentTrainingPhraseType',
    'GoogleCloudDialogflowV2NotificationConfigMessageFormat',
    'GoogleCloudDialogflowV2SmartReplyModelMetadataTrainingModelType',
    'GoogleCloudDialogflowV2SpeechToTextConfigSpeechModelVariant',
    'GoogleCloudDialogflowV2SuggestionFeatureType',
    'GoogleCloudDialogflowV2TextToSpeechSettingsOutputAudioEncoding',
    'IntentDefaultResponsePlatformsItem',
    'IntentWebhookState',
    'ParticipantRole',
    'SessionEntityTypeEntityOverrideMode',
]


class ConversationConversationStage(str, Enum):
    """
    The stage of a conversation. It indicates whether the virtual agent or a human agent is handling the conversation. If the conversation is created with the conversation profile that has Dialogflow config set, defaults to ConversationStage.VIRTUAL_AGENT_STAGE; Otherwise, defaults to ConversationStage.HUMAN_ASSIST_STAGE. If the conversation is created with the conversation profile that has Dialogflow config set but explicitly sets conversation_stage to ConversationStage.HUMAN_ASSIST_STAGE, it skips ConversationStage.VIRTUAL_AGENT_STAGE stage and directly goes to ConversationStage.HUMAN_ASSIST_STAGE.
    """
    CONVERSATION_STAGE_UNSPECIFIED = "CONVERSATION_STAGE_UNSPECIFIED"
    """
    Unknown. Should never be used after a conversation is successfully created.
    """
    VIRTUAL_AGENT_STAGE = "VIRTUAL_AGENT_STAGE"
    """
    The conversation should return virtual agent responses into the conversation.
    """
    HUMAN_ASSIST_STAGE = "HUMAN_ASSIST_STAGE"
    """
    The conversation should not provide responses, just listen and provide suggestions.
    """


class DocumentKnowledgeTypesItem(str, Enum):
    KNOWLEDGE_TYPE_UNSPECIFIED = "KNOWLEDGE_TYPE_UNSPECIFIED"
    """
    The type is unspecified or arbitrary.
    """
    FAQ = "FAQ"
    """
    The document content contains question and answer pairs as either HTML or CSV. Typical FAQ HTML formats are parsed accurately, but unusual formats may fail to be parsed. CSV must have questions in the first column and answers in the second, with no header. Because of this explicit format, they are always parsed accurately.
    """
    EXTRACTIVE_QA = "EXTRACTIVE_QA"
    """
    Documents for which unstructured text is extracted and used for question answering.
    """
    ARTICLE_SUGGESTION = "ARTICLE_SUGGESTION"
    """
    The entire document content as a whole can be used for query results. Only for Contact Center Solutions on Dialogflow.
    """
    AGENT_FACING_SMART_REPLY = "AGENT_FACING_SMART_REPLY"
    """
    The document contains agent-facing Smart Reply entries.
    """


class EntityTypeAutoExpansionMode(str, Enum):
    """
    Optional. Indicates whether the entity type can be automatically expanded.
    """
    AUTO_EXPANSION_MODE_UNSPECIFIED = "AUTO_EXPANSION_MODE_UNSPECIFIED"
    """
    Auto expansion disabled for the entity.
    """
    AUTO_EXPANSION_MODE_DEFAULT = "AUTO_EXPANSION_MODE_DEFAULT"
    """
    Allows an agent to recognize values that have not been explicitly listed in the entity.
    """


class EntityTypeKind(str, Enum):
    """
    Required. Indicates the kind of entity type.
    """
    KIND_UNSPECIFIED = "KIND_UNSPECIFIED"
    """
    Not specified. This value should be never used.
    """
    KIND_MAP = "KIND_MAP"
    """
    Map entity types allow mapping of a group of synonyms to a reference value.
    """
    KIND_LIST = "KIND_LIST"
    """
    List entity types contain a set of entries that do not map to reference values. However, list entity types can contain references to other entity types (with or without aliases).
    """
    KIND_REGEXP = "KIND_REGEXP"
    """
    Regexp entity types allow to specify regular expressions in entries values.
    """


class GoogleCloudDialogflowV2ArticleSuggestionModelMetadataTrainingModelType(str, Enum):
    """
    Optional. Type of the article suggestion model. If not provided, model_type is used.
    """
    MODEL_TYPE_UNSPECIFIED = "MODEL_TYPE_UNSPECIFIED"
    """
    ModelType unspecified.
    """
    SMART_REPLY_DUAL_ENCODER_MODEL = "SMART_REPLY_DUAL_ENCODER_MODEL"
    """
    ModelType smart reply dual encoder model.
    """
    SMART_REPLY_BERT_MODEL = "SMART_REPLY_BERT_MODEL"
    """
    ModelType smart reply bert model.
    """


class GoogleCloudDialogflowV2FulfillmentFeatureType(str, Enum):
    """
    The type of the feature that enabled for fulfillment.
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    """
    Feature type not specified.
    """
    SMALLTALK = "SMALLTALK"
    """
    Fulfillment is enabled for SmallTalk.
    """


class GoogleCloudDialogflowV2IntentMessageBrowseCarouselCardBrowseCarouselCardItemOpenUrlActionUrlTypeHint(str, Enum):
    """
    Optional. Specifies the type of viewer that is used when opening the URL. Defaults to opening via web browser.
    """
    URL_TYPE_HINT_UNSPECIFIED = "URL_TYPE_HINT_UNSPECIFIED"
    """
    Unspecified
    """
    AMP_ACTION = "AMP_ACTION"
    """
    Url would be an amp action
    """
    AMP_CONTENT = "AMP_CONTENT"
    """
    URL that points directly to AMP content, or to a canonical URL which refers to AMP content via .
    """


class GoogleCloudDialogflowV2IntentMessageBrowseCarouselCardImageDisplayOptions(str, Enum):
    """
    Optional. Settings for displaying the image. Applies to every image in items.
    """
    IMAGE_DISPLAY_OPTIONS_UNSPECIFIED = "IMAGE_DISPLAY_OPTIONS_UNSPECIFIED"
    """
    Fill the gaps between the image and the image container with gray bars.
    """
    GRAY = "GRAY"
    """
    Fill the gaps between the image and the image container with gray bars.
    """
    WHITE = "WHITE"
    """
    Fill the gaps between the image and the image container with white bars.
    """
    CROPPED = "CROPPED"
    """
    Image is scaled such that the image width and height match or exceed the container dimensions. This may crop the top and bottom of the image if the scaled image height is greater than the container height, or crop the left and right of the image if the scaled image width is greater than the container width. This is similar to "Zoom Mode" on a widescreen TV when playing a 4:3 video.
    """
    BLURRED_BACKGROUND = "BLURRED_BACKGROUND"
    """
    Pad the gaps between image and image frame with a blurred copy of the same image.
    """


class GoogleCloudDialogflowV2IntentMessageColumnPropertiesHorizontalAlignment(str, Enum):
    """
    Optional. Defines text alignment for all cells in this column.
    """
    HORIZONTAL_ALIGNMENT_UNSPECIFIED = "HORIZONTAL_ALIGNMENT_UNSPECIFIED"
    """
    Text is aligned to the leading edge of the column.
    """
    LEADING = "LEADING"
    """
    Text is aligned to the leading edge of the column.
    """
    CENTER = "CENTER"
    """
    Text is centered in the column.
    """
    TRAILING = "TRAILING"
    """
    Text is aligned to the trailing edge of the column.
    """


class GoogleCloudDialogflowV2IntentMessageMediaContentMediaType(str, Enum):
    """
    Optional. What type of media is the content (ie "audio").
    """
    RESPONSE_MEDIA_TYPE_UNSPECIFIED = "RESPONSE_MEDIA_TYPE_UNSPECIFIED"
    """
    Unspecified.
    """
    AUDIO = "AUDIO"
    """
    Response media type is audio.
    """


class GoogleCloudDialogflowV2IntentMessagePlatform(str, Enum):
    """
    Optional. The platform that this message is intended for.
    """
    PLATFORM_UNSPECIFIED = "PLATFORM_UNSPECIFIED"
    """
    Default platform.
    """
    FACEBOOK = "FACEBOOK"
    """
    Facebook.
    """
    SLACK = "SLACK"
    """
    Slack.
    """
    TELEGRAM = "TELEGRAM"
    """
    Telegram.
    """
    KIK = "KIK"
    """
    Kik.
    """
    SKYPE = "SKYPE"
    """
    Skype.
    """
    LINE = "LINE"
    """
    Line.
    """
    VIBER = "VIBER"
    """
    Viber.
    """
    ACTIONS_ON_GOOGLE = "ACTIONS_ON_GOOGLE"
    """
    Google Assistant See [Dialogflow webhook format](https://developers.google.com/assistant/actions/build/json/dialogflow-webhook-json)
    """
    GOOGLE_HANGOUTS = "GOOGLE_HANGOUTS"
    """
    Google Hangouts.
    """


class GoogleCloudDialogflowV2IntentTrainingPhraseType(str, Enum):
    """
    Required. The type of the training phrase.
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    """
    Not specified. This value should never be used.
    """
    EXAMPLE = "EXAMPLE"
    """
    Examples do not contain @-prefixed entity type names, but example parts can be annotated with entity types.
    """
    TEMPLATE = "TEMPLATE"
    """
    Templates are not annotated with entity types, but they can contain @-prefixed entity type names as substrings. Template mode has been deprecated. Example mode is the only supported way to create new training phrases. If you have existing training phrases that you've created in template mode, those will continue to work.
    """


class GoogleCloudDialogflowV2NotificationConfigMessageFormat(str, Enum):
    """
    Format of message.
    """
    MESSAGE_FORMAT_UNSPECIFIED = "MESSAGE_FORMAT_UNSPECIFIED"
    """
    If it is unspecified, PROTO will be used.
    """
    PROTO = "PROTO"
    """
    Pub/Sub message will be serialized proto.
    """
    JSON = "JSON"
    """
    Pub/Sub message will be json.
    """


class GoogleCloudDialogflowV2SmartReplyModelMetadataTrainingModelType(str, Enum):
    """
    Optional. Type of the smart reply model. If not provided, model_type is used.
    """
    MODEL_TYPE_UNSPECIFIED = "MODEL_TYPE_UNSPECIFIED"
    """
    ModelType unspecified.
    """
    SMART_REPLY_DUAL_ENCODER_MODEL = "SMART_REPLY_DUAL_ENCODER_MODEL"
    """
    ModelType smart reply dual encoder model.
    """
    SMART_REPLY_BERT_MODEL = "SMART_REPLY_BERT_MODEL"
    """
    ModelType smart reply bert model.
    """


class GoogleCloudDialogflowV2SpeechToTextConfigSpeechModelVariant(str, Enum):
    """
    The speech model used in speech to text. `SPEECH_MODEL_VARIANT_UNSPECIFIED`, `USE_BEST_AVAILABLE` will be treated as `USE_ENHANCED`. It can be overridden in AnalyzeContentRequest and StreamingAnalyzeContentRequest request. If enhanced model variant is specified and an enhanced version of the specified model for the language does not exist, then it would emit an error.
    """
    SPEECH_MODEL_VARIANT_UNSPECIFIED = "SPEECH_MODEL_VARIANT_UNSPECIFIED"
    """
    No model variant specified. In this case Dialogflow defaults to USE_BEST_AVAILABLE.
    """
    USE_BEST_AVAILABLE = "USE_BEST_AVAILABLE"
    """
    Use the best available variant of the Speech model that the caller is eligible for. Please see the [Dialogflow docs](https://cloud.google.com/dialogflow/docs/data-logging) for how to make your project eligible for enhanced models.
    """
    USE_STANDARD = "USE_STANDARD"
    """
    Use standard model variant even if an enhanced model is available. See the [Cloud Speech documentation](https://cloud.google.com/speech-to-text/docs/enhanced-models) for details about enhanced models.
    """
    USE_ENHANCED = "USE_ENHANCED"
    """
    Use an enhanced model variant: * If an enhanced variant does not exist for the given model and request language, Dialogflow falls back to the standard variant. The [Cloud Speech documentation](https://cloud.google.com/speech-to-text/docs/enhanced-models) describes which models have enhanced variants. * If the API caller isn't eligible for enhanced models, Dialogflow returns an error. Please see the [Dialogflow docs](https://cloud.google.com/dialogflow/docs/data-logging) for how to make your project eligible.
    """


class GoogleCloudDialogflowV2SuggestionFeatureType(str, Enum):
    """
    Type of Human Agent Assistant API feature to request.
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    """
    Unspecified feature type.
    """
    ARTICLE_SUGGESTION = "ARTICLE_SUGGESTION"
    """
    Run article suggestion model.
    """
    FAQ = "FAQ"
    """
    Run FAQ model.
    """
    SMART_REPLY = "SMART_REPLY"
    """
    Run smart reply model.
    """


class GoogleCloudDialogflowV2TextToSpeechSettingsOutputAudioEncoding(str, Enum):
    """
    Required. Audio encoding of the synthesized audio content.
    """
    OUTPUT_AUDIO_ENCODING_UNSPECIFIED = "OUTPUT_AUDIO_ENCODING_UNSPECIFIED"
    """
    Not specified.
    """
    OUTPUT_AUDIO_ENCODING_LINEAR16 = "OUTPUT_AUDIO_ENCODING_LINEAR_16"
    """
    Uncompressed 16-bit signed little-endian samples (Linear PCM). Audio content returned as LINEAR16 also contains a WAV header.
    """
    OUTPUT_AUDIO_ENCODING_MP3 = "OUTPUT_AUDIO_ENCODING_MP3"
    """
    MP3 audio at 32kbps.
    """
    OUTPUT_AUDIO_ENCODING_MP364_KBPS = "OUTPUT_AUDIO_ENCODING_MP3_64_KBPS"
    """
    MP3 audio at 64kbps.
    """
    OUTPUT_AUDIO_ENCODING_OGG_OPUS = "OUTPUT_AUDIO_ENCODING_OGG_OPUS"
    """
    Opus encoded audio wrapped in an ogg container. The result will be a file which can be played natively on Android, and in browsers (at least Chrome and Firefox). The quality of the encoding is considerably higher than MP3 while using approximately the same bitrate.
    """
    OUTPUT_AUDIO_ENCODING_MULAW = "OUTPUT_AUDIO_ENCODING_MULAW"
    """
    8-bit samples that compand 14-bit audio samples using G.711 PCMU/mu-law.
    """


class IntentDefaultResponsePlatformsItem(str, Enum):
    PLATFORM_UNSPECIFIED = "PLATFORM_UNSPECIFIED"
    """
    Default platform.
    """
    FACEBOOK = "FACEBOOK"
    """
    Facebook.
    """
    SLACK = "SLACK"
    """
    Slack.
    """
    TELEGRAM = "TELEGRAM"
    """
    Telegram.
    """
    KIK = "KIK"
    """
    Kik.
    """
    SKYPE = "SKYPE"
    """
    Skype.
    """
    LINE = "LINE"
    """
    Line.
    """
    VIBER = "VIBER"
    """
    Viber.
    """
    ACTIONS_ON_GOOGLE = "ACTIONS_ON_GOOGLE"
    """
    Google Assistant See [Dialogflow webhook format](https://developers.google.com/assistant/actions/build/json/dialogflow-webhook-json)
    """
    GOOGLE_HANGOUTS = "GOOGLE_HANGOUTS"
    """
    Google Hangouts.
    """


class IntentWebhookState(str, Enum):
    """
    Optional. Indicates whether webhooks are enabled for the intent.
    """
    WEBHOOK_STATE_UNSPECIFIED = "WEBHOOK_STATE_UNSPECIFIED"
    """
    Webhook is disabled in the agent and in the intent.
    """
    WEBHOOK_STATE_ENABLED = "WEBHOOK_STATE_ENABLED"
    """
    Webhook is enabled in the agent and in the intent.
    """
    WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING = "WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING"
    """
    Webhook is enabled in the agent and in the intent. Also, each slot filling prompt is forwarded to the webhook.
    """


class ParticipantRole(str, Enum):
    """
    Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
    """
    ROLE_UNSPECIFIED = "ROLE_UNSPECIFIED"
    """
    Participant role not set.
    """
    HUMAN_AGENT = "HUMAN_AGENT"
    """
    Participant is a human agent.
    """
    AUTOMATED_AGENT = "AUTOMATED_AGENT"
    """
    Participant is an automated agent, such as a Dialogflow agent.
    """
    END_USER = "END_USER"
    """
    Participant is an end user that has called or chatted with Dialogflow services.
    """


class SessionEntityTypeEntityOverrideMode(str, Enum):
    """
    Required. Indicates whether the additional data should override or supplement the custom entity type definition.
    """
    ENTITY_OVERRIDE_MODE_UNSPECIFIED = "ENTITY_OVERRIDE_MODE_UNSPECIFIED"
    """
    Not specified. This value should be never used.
    """
    ENTITY_OVERRIDE_MODE_OVERRIDE = "ENTITY_OVERRIDE_MODE_OVERRIDE"
    """
    The collection of session entities overrides the collection of entities in the corresponding custom entity type.
    """
    ENTITY_OVERRIDE_MODE_SUPPLEMENT = "ENTITY_OVERRIDE_MODE_SUPPLEMENT"
    """
    The collection of session entities extends the collection of entities in the corresponding custom entity type. Note: Even in this override mode calls to `ListSessionEntityTypes`, `GetSessionEntityType`, `CreateSessionEntityType` and `UpdateSessionEntityType` only return the additional entities added in this session entity type. If you want to get the supplemented list, please call EntityTypes.GetEntityType on the custom entity type and merge.
    """
