# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetIntentResult',
    'AwaitableGetIntentResult',
    'get_intent',
]

@pulumi.output_type
class GetIntentResult:
    def __init__(__self__, action=None, default_response_platforms=None, display_name=None, end_interaction=None, events=None, followup_intent_info=None, input_context_names=None, is_fallback=None, live_agent_handoff=None, messages=None, ml_disabled=None, name=None, output_contexts=None, parameters=None, parent_followup_intent_name=None, priority=None, reset_contexts=None, root_followup_intent_name=None, training_phrases=None, webhook_state=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if default_response_platforms and not isinstance(default_response_platforms, list):
            raise TypeError("Expected argument 'default_response_platforms' to be a list")
        pulumi.set(__self__, "default_response_platforms", default_response_platforms)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if end_interaction and not isinstance(end_interaction, bool):
            raise TypeError("Expected argument 'end_interaction' to be a bool")
        pulumi.set(__self__, "end_interaction", end_interaction)
        if events and not isinstance(events, list):
            raise TypeError("Expected argument 'events' to be a list")
        pulumi.set(__self__, "events", events)
        if followup_intent_info and not isinstance(followup_intent_info, list):
            raise TypeError("Expected argument 'followup_intent_info' to be a list")
        pulumi.set(__self__, "followup_intent_info", followup_intent_info)
        if input_context_names and not isinstance(input_context_names, list):
            raise TypeError("Expected argument 'input_context_names' to be a list")
        pulumi.set(__self__, "input_context_names", input_context_names)
        if is_fallback and not isinstance(is_fallback, bool):
            raise TypeError("Expected argument 'is_fallback' to be a bool")
        pulumi.set(__self__, "is_fallback", is_fallback)
        if live_agent_handoff and not isinstance(live_agent_handoff, bool):
            raise TypeError("Expected argument 'live_agent_handoff' to be a bool")
        pulumi.set(__self__, "live_agent_handoff", live_agent_handoff)
        if messages and not isinstance(messages, list):
            raise TypeError("Expected argument 'messages' to be a list")
        pulumi.set(__self__, "messages", messages)
        if ml_disabled and not isinstance(ml_disabled, bool):
            raise TypeError("Expected argument 'ml_disabled' to be a bool")
        pulumi.set(__self__, "ml_disabled", ml_disabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if output_contexts and not isinstance(output_contexts, list):
            raise TypeError("Expected argument 'output_contexts' to be a list")
        pulumi.set(__self__, "output_contexts", output_contexts)
        if parameters and not isinstance(parameters, list):
            raise TypeError("Expected argument 'parameters' to be a list")
        pulumi.set(__self__, "parameters", parameters)
        if parent_followup_intent_name and not isinstance(parent_followup_intent_name, str):
            raise TypeError("Expected argument 'parent_followup_intent_name' to be a str")
        pulumi.set(__self__, "parent_followup_intent_name", parent_followup_intent_name)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if reset_contexts and not isinstance(reset_contexts, bool):
            raise TypeError("Expected argument 'reset_contexts' to be a bool")
        pulumi.set(__self__, "reset_contexts", reset_contexts)
        if root_followup_intent_name and not isinstance(root_followup_intent_name, str):
            raise TypeError("Expected argument 'root_followup_intent_name' to be a str")
        pulumi.set(__self__, "root_followup_intent_name", root_followup_intent_name)
        if training_phrases and not isinstance(training_phrases, list):
            raise TypeError("Expected argument 'training_phrases' to be a list")
        pulumi.set(__self__, "training_phrases", training_phrases)
        if webhook_state and not isinstance(webhook_state, str):
            raise TypeError("Expected argument 'webhook_state' to be a str")
        pulumi.set(__self__, "webhook_state", webhook_state)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="defaultResponsePlatforms")
    def default_response_platforms(self) -> Sequence[str]:
        """
        Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
        """
        return pulumi.get(self, "default_response_platforms")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The name of this intent.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="endInteraction")
    def end_interaction(self) -> bool:
        """
        Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
        """
        return pulumi.get(self, "end_interaction")

    @property
    @pulumi.getter
    def events(self) -> Sequence[str]:
        """
        Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
        """
        return pulumi.get(self, "events")

    @property
    @pulumi.getter(name="followupIntentInfo")
    def followup_intent_info(self) -> Sequence['outputs.GoogleCloudDialogflowV2IntentFollowupIntentInfoResponse']:
        """
        Read-only. Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only in the output.
        """
        return pulumi.get(self, "followup_intent_info")

    @property
    @pulumi.getter(name="inputContextNames")
    def input_context_names(self) -> Sequence[str]:
        """
        Optional. The list of context names required for this intent to be triggered. Format: `projects//agent/sessions/-/contexts/`.
        """
        return pulumi.get(self, "input_context_names")

    @property
    @pulumi.getter(name="isFallback")
    def is_fallback(self) -> bool:
        """
        Optional. Indicates whether this is a fallback intent.
        """
        return pulumi.get(self, "is_fallback")

    @property
    @pulumi.getter(name="liveAgentHandoff")
    def live_agent_handoff(self) -> bool:
        """
        Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
        """
        return pulumi.get(self, "live_agent_handoff")

    @property
    @pulumi.getter
    def messages(self) -> Sequence['outputs.GoogleCloudDialogflowV2IntentMessageResponse']:
        """
        Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
        """
        return pulumi.get(self, "messages")

    @property
    @pulumi.getter(name="mlDisabled")
    def ml_disabled(self) -> bool:
        """
        Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
        """
        return pulumi.get(self, "ml_disabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Format: `projects//agent/intents/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputContexts")
    def output_contexts(self) -> Sequence['outputs.GoogleCloudDialogflowV2ContextResponse']:
        """
        Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
        """
        return pulumi.get(self, "output_contexts")

    @property
    @pulumi.getter
    def parameters(self) -> Sequence['outputs.GoogleCloudDialogflowV2IntentParameterResponse']:
        """
        Optional. The collection of parameters associated with the intent.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="parentFollowupIntentName")
    def parent_followup_intent_name(self) -> str:
        """
        Read-only after creation. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
        """
        return pulumi.get(self, "parent_followup_intent_name")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="resetContexts")
    def reset_contexts(self) -> bool:
        """
        Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
        """
        return pulumi.get(self, "reset_contexts")

    @property
    @pulumi.getter(name="rootFollowupIntentName")
    def root_followup_intent_name(self) -> str:
        """
        Read-only. The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents chain for this intent. We populate this field only in the output. Format: `projects//agent/intents/`.
        """
        return pulumi.get(self, "root_followup_intent_name")

    @property
    @pulumi.getter(name="trainingPhrases")
    def training_phrases(self) -> Sequence['outputs.GoogleCloudDialogflowV2IntentTrainingPhraseResponse']:
        """
        Optional. The collection of examples that the agent is trained on.
        """
        return pulumi.get(self, "training_phrases")

    @property
    @pulumi.getter(name="webhookState")
    def webhook_state(self) -> str:
        """
        Optional. Indicates whether webhooks are enabled for the intent.
        """
        return pulumi.get(self, "webhook_state")


class AwaitableGetIntentResult(GetIntentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIntentResult(
            action=self.action,
            default_response_platforms=self.default_response_platforms,
            display_name=self.display_name,
            end_interaction=self.end_interaction,
            events=self.events,
            followup_intent_info=self.followup_intent_info,
            input_context_names=self.input_context_names,
            is_fallback=self.is_fallback,
            live_agent_handoff=self.live_agent_handoff,
            messages=self.messages,
            ml_disabled=self.ml_disabled,
            name=self.name,
            output_contexts=self.output_contexts,
            parameters=self.parameters,
            parent_followup_intent_name=self.parent_followup_intent_name,
            priority=self.priority,
            reset_contexts=self.reset_contexts,
            root_followup_intent_name=self.root_followup_intent_name,
            training_phrases=self.training_phrases,
            webhook_state=self.webhook_state)


def get_intent(intent_id: Optional[str] = None,
               intent_view: Optional[str] = None,
               language_code: Optional[str] = None,
               location: Optional[str] = None,
               project: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntentResult:
    """
    Retrieves the specified intent.
    """
    __args__ = dict()
    __args__['intentId'] = intent_id
    __args__['intentView'] = intent_view
    __args__['languageCode'] = language_code
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:dialogflow/v2:getIntent', __args__, opts=opts, typ=GetIntentResult).value

    return AwaitableGetIntentResult(
        action=__ret__.action,
        default_response_platforms=__ret__.default_response_platforms,
        display_name=__ret__.display_name,
        end_interaction=__ret__.end_interaction,
        events=__ret__.events,
        followup_intent_info=__ret__.followup_intent_info,
        input_context_names=__ret__.input_context_names,
        is_fallback=__ret__.is_fallback,
        live_agent_handoff=__ret__.live_agent_handoff,
        messages=__ret__.messages,
        ml_disabled=__ret__.ml_disabled,
        name=__ret__.name,
        output_contexts=__ret__.output_contexts,
        parameters=__ret__.parameters,
        parent_followup_intent_name=__ret__.parent_followup_intent_name,
        priority=__ret__.priority,
        reset_contexts=__ret__.reset_contexts,
        root_followup_intent_name=__ret__.root_followup_intent_name,
        training_phrases=__ret__.training_phrases,
        webhook_state=__ret__.webhook_state)
