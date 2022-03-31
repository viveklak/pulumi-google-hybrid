// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified intent.
 */
export function getIntent(args: GetIntentArgs, opts?: pulumi.InvokeOptions): Promise<GetIntentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dialogflow/v2:getIntent", {
        "intentId": args.intentId,
        "intentView": args.intentView,
        "languageCode": args.languageCode,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetIntentArgs {
    intentId: string;
    intentView?: string;
    languageCode?: string;
    location: string;
    project?: string;
}

export interface GetIntentResult {
    /**
     * Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
     */
    readonly action: string;
    /**
     * Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
     */
    readonly defaultResponsePlatforms: string[];
    /**
     * The name of this intent.
     */
    readonly displayName: string;
    /**
     * Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
     */
    readonly endInteraction: boolean;
    /**
     * Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
     */
    readonly events: string[];
    /**
     * Read-only. Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only in the output.
     */
    readonly followupIntentInfo: outputs.dialogflow.v2.GoogleCloudDialogflowV2IntentFollowupIntentInfoResponse[];
    /**
     * Optional. The list of context names required for this intent to be triggered. Format: `projects//agent/sessions/-/contexts/`.
     */
    readonly inputContextNames: string[];
    /**
     * Optional. Indicates whether this is a fallback intent.
     */
    readonly isFallback: boolean;
    /**
     * Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
     */
    readonly liveAgentHandoff: boolean;
    /**
     * Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
     */
    readonly messages: outputs.dialogflow.v2.GoogleCloudDialogflowV2IntentMessageResponse[];
    /**
     * Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
     */
    readonly mlDisabled: boolean;
    /**
     * Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Format: `projects//agent/intents/`.
     */
    readonly name: string;
    /**
     * Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
     */
    readonly outputContexts: outputs.dialogflow.v2.GoogleCloudDialogflowV2ContextResponse[];
    /**
     * Optional. The collection of parameters associated with the intent.
     */
    readonly parameters: outputs.dialogflow.v2.GoogleCloudDialogflowV2IntentParameterResponse[];
    /**
     * Read-only after creation. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
     */
    readonly parentFollowupIntentName: string;
    /**
     * Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
     */
    readonly priority: number;
    /**
     * Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
     */
    readonly resetContexts: boolean;
    /**
     * Read-only. The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents chain for this intent. We populate this field only in the output. Format: `projects//agent/intents/`.
     */
    readonly rootFollowupIntentName: string;
    /**
     * Optional. The collection of examples that the agent is trained on.
     */
    readonly trainingPhrases: outputs.dialogflow.v2.GoogleCloudDialogflowV2IntentTrainingPhraseResponse[];
    /**
     * Optional. Indicates whether webhooks are enabled for the intent.
     */
    readonly webhookState: string;
}

export function getIntentOutput(args: GetIntentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIntentResult> {
    return pulumi.output(args).apply(a => getIntent(a, opts))
}

export interface GetIntentOutputArgs {
    intentId: pulumi.Input<string>;
    intentView?: pulumi.Input<string>;
    languageCode?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
