// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an intent in the specified agent. Note: You should always train an agent prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/es/docs/training).
 * Auto-naming is currently not supported for this resource.
 */
export class Intent extends pulumi.CustomResource {
    /**
     * Get an existing Intent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Intent {
        return new Intent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2:Intent';

    /**
     * Returns true if the given object is an instance of Intent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Intent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Intent.__pulumiType;
    }

    /**
     * Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
     */
    public readonly defaultResponsePlatforms!: pulumi.Output<string[]>;
    /**
     * The name of this intent.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
     */
    public readonly endInteraction!: pulumi.Output<boolean>;
    /**
     * Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
     */
    public readonly events!: pulumi.Output<string[]>;
    /**
     * Read-only. Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only in the output.
     */
    public /*out*/ readonly followupIntentInfo!: pulumi.Output<outputs.dialogflow.v2.GoogleCloudDialogflowV2IntentFollowupIntentInfoResponse[]>;
    /**
     * Optional. The list of context names required for this intent to be triggered. Format: `projects//agent/sessions/-/contexts/`.
     */
    public readonly inputContextNames!: pulumi.Output<string[]>;
    /**
     * Optional. Indicates whether this is a fallback intent.
     */
    public readonly isFallback!: pulumi.Output<boolean>;
    /**
     * Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
     */
    public readonly liveAgentHandoff!: pulumi.Output<boolean>;
    /**
     * Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
     */
    public readonly messages!: pulumi.Output<outputs.dialogflow.v2.GoogleCloudDialogflowV2IntentMessageResponse[]>;
    /**
     * Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
     */
    public readonly mlDisabled!: pulumi.Output<boolean>;
    /**
     * Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Format: `projects//agent/intents/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
     */
    public readonly outputContexts!: pulumi.Output<outputs.dialogflow.v2.GoogleCloudDialogflowV2ContextResponse[]>;
    /**
     * Optional. The collection of parameters associated with the intent.
     */
    public readonly parameters!: pulumi.Output<outputs.dialogflow.v2.GoogleCloudDialogflowV2IntentParameterResponse[]>;
    /**
     * Read-only after creation. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
     */
    public readonly parentFollowupIntentName!: pulumi.Output<string>;
    /**
     * Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
     */
    public readonly resetContexts!: pulumi.Output<boolean>;
    /**
     * Read-only. The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents chain for this intent. We populate this field only in the output. Format: `projects//agent/intents/`.
     */
    public /*out*/ readonly rootFollowupIntentName!: pulumi.Output<string>;
    /**
     * Optional. The collection of examples that the agent is trained on.
     */
    public readonly trainingPhrases!: pulumi.Output<outputs.dialogflow.v2.GoogleCloudDialogflowV2IntentTrainingPhraseResponse[]>;
    /**
     * Optional. Indicates whether webhooks are enabled for the intent.
     */
    public readonly webhookState!: pulumi.Output<string>;

    /**
     * Create a Intent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["defaultResponsePlatforms"] = args ? args.defaultResponsePlatforms : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["endInteraction"] = args ? args.endInteraction : undefined;
            resourceInputs["events"] = args ? args.events : undefined;
            resourceInputs["inputContextNames"] = args ? args.inputContextNames : undefined;
            resourceInputs["intentView"] = args ? args.intentView : undefined;
            resourceInputs["isFallback"] = args ? args.isFallback : undefined;
            resourceInputs["languageCode"] = args ? args.languageCode : undefined;
            resourceInputs["liveAgentHandoff"] = args ? args.liveAgentHandoff : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["messages"] = args ? args.messages : undefined;
            resourceInputs["mlDisabled"] = args ? args.mlDisabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outputContexts"] = args ? args.outputContexts : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["parentFollowupIntentName"] = args ? args.parentFollowupIntentName : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["resetContexts"] = args ? args.resetContexts : undefined;
            resourceInputs["trainingPhrases"] = args ? args.trainingPhrases : undefined;
            resourceInputs["webhookState"] = args ? args.webhookState : undefined;
            resourceInputs["followupIntentInfo"] = undefined /*out*/;
            resourceInputs["rootFollowupIntentName"] = undefined /*out*/;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["defaultResponsePlatforms"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["endInteraction"] = undefined /*out*/;
            resourceInputs["events"] = undefined /*out*/;
            resourceInputs["followupIntentInfo"] = undefined /*out*/;
            resourceInputs["inputContextNames"] = undefined /*out*/;
            resourceInputs["isFallback"] = undefined /*out*/;
            resourceInputs["liveAgentHandoff"] = undefined /*out*/;
            resourceInputs["messages"] = undefined /*out*/;
            resourceInputs["mlDisabled"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["outputContexts"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["parentFollowupIntentName"] = undefined /*out*/;
            resourceInputs["priority"] = undefined /*out*/;
            resourceInputs["resetContexts"] = undefined /*out*/;
            resourceInputs["rootFollowupIntentName"] = undefined /*out*/;
            resourceInputs["trainingPhrases"] = undefined /*out*/;
            resourceInputs["webhookState"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Intent.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Intent resource.
 */
export interface IntentArgs {
    /**
     * Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
     */
    action?: pulumi.Input<string>;
    /**
     * Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
     */
    defaultResponsePlatforms?: pulumi.Input<pulumi.Input<enums.dialogflow.v2.IntentDefaultResponsePlatformsItem>[]>;
    /**
     * The name of this intent.
     */
    displayName: pulumi.Input<string>;
    /**
     * Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
     */
    endInteraction?: pulumi.Input<boolean>;
    /**
     * Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
     */
    events?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. The list of context names required for this intent to be triggered. Format: `projects//agent/sessions/-/contexts/`.
     */
    inputContextNames?: pulumi.Input<pulumi.Input<string>[]>;
    intentView?: pulumi.Input<string>;
    /**
     * Optional. Indicates whether this is a fallback intent.
     */
    isFallback?: pulumi.Input<boolean>;
    languageCode?: pulumi.Input<string>;
    /**
     * Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
     */
    liveAgentHandoff?: pulumi.Input<boolean>;
    location?: pulumi.Input<string>;
    /**
     * Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
     */
    messages?: pulumi.Input<pulumi.Input<inputs.dialogflow.v2.GoogleCloudDialogflowV2IntentMessageArgs>[]>;
    /**
     * Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
     */
    mlDisabled?: pulumi.Input<boolean>;
    /**
     * Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Format: `projects//agent/intents/`.
     */
    name?: pulumi.Input<string>;
    /**
     * Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
     */
    outputContexts?: pulumi.Input<pulumi.Input<inputs.dialogflow.v2.GoogleCloudDialogflowV2ContextArgs>[]>;
    /**
     * Optional. The collection of parameters associated with the intent.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.dialogflow.v2.GoogleCloudDialogflowV2IntentParameterArgs>[]>;
    /**
     * Read-only after creation. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
     */
    parentFollowupIntentName?: pulumi.Input<string>;
    /**
     * Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
     */
    priority?: pulumi.Input<number>;
    project?: pulumi.Input<string>;
    /**
     * Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
     */
    resetContexts?: pulumi.Input<boolean>;
    /**
     * Optional. The collection of examples that the agent is trained on.
     */
    trainingPhrases?: pulumi.Input<pulumi.Input<inputs.dialogflow.v2.GoogleCloudDialogflowV2IntentTrainingPhraseArgs>[]>;
    /**
     * Optional. Indicates whether webhooks are enabled for the intent.
     */
    webhookState?: pulumi.Input<enums.dialogflow.v2.IntentWebhookState>;
}
