// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an intent in the specified agent.
 */
export class AgentIntent extends pulumi.CustomResource {
    /**
     * Get an existing AgentIntent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AgentIntent {
        return new AgentIntent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2beta1:AgentIntent';

    /**
     * Returns true if the given object is an instance of AgentIntent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentIntent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentIntent.__pulumiType;
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
     * Required. The name of this intent.
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
     * Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only in the output.
     */
    public /*out*/ readonly followupIntentInfo!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1IntentFollowupIntentInfoResponse[]>;
    /**
     * Optional. The list of context names required for this intent to be triggered. Formats: - `projects//agent/sessions/-/contexts/` - `projects//locations//agent/sessions/-/contexts/`
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
    public readonly messages!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1IntentMessageResponse[]>;
    /**
     * Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
     */
    public readonly mlDisabled!: pulumi.Output<boolean>;
    /**
     * Optional. Indicates whether Machine Learning is enabled for the intent. Note: If `ml_enabled` setting is set to false, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off. DEPRECATED! Please use `ml_disabled` field instead. NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false, then the default value is determined as follows: - Before April 15th, 2018 the default is: ml_enabled = false / ml_disabled = true. - After April 15th, 2018 the default is: ml_enabled = true / ml_disabled = false.
     */
    public readonly mlEnabled!: pulumi.Output<boolean>;
    /**
     * Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Supported formats: - `projects//agent/intents/` - `projects//locations//agent/intents/`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
     */
    public readonly outputContexts!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1ContextResponse[]>;
    /**
     * Optional. The collection of parameters associated with the intent.
     */
    public readonly parameters!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1IntentParameterResponse[]>;
    /**
     * Optional. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
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
     * The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents chain for this intent. Format: `projects//agent/intents/`.
     */
    public /*out*/ readonly rootFollowupIntentName!: pulumi.Output<string>;
    /**
     * Optional. The collection of examples that the agent is trained on.
     */
    public readonly trainingPhrases!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1IntentTrainingPhraseResponse[]>;
    /**
     * Optional. Indicates whether webhooks are enabled for the intent.
     */
    public readonly webhookState!: pulumi.Output<string>;

    /**
     * Create a AgentIntent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentIntentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["defaultResponsePlatforms"] = args ? args.defaultResponsePlatforms : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["endInteraction"] = args ? args.endInteraction : undefined;
            inputs["events"] = args ? args.events : undefined;
            inputs["inputContextNames"] = args ? args.inputContextNames : undefined;
            inputs["intentView"] = args ? args.intentView : undefined;
            inputs["isFallback"] = args ? args.isFallback : undefined;
            inputs["languageCode"] = args ? args.languageCode : undefined;
            inputs["liveAgentHandoff"] = args ? args.liveAgentHandoff : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["messages"] = args ? args.messages : undefined;
            inputs["mlDisabled"] = args ? args.mlDisabled : undefined;
            inputs["mlEnabled"] = args ? args.mlEnabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["outputContexts"] = args ? args.outputContexts : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["parentFollowupIntentName"] = args ? args.parentFollowupIntentName : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["resetContexts"] = args ? args.resetContexts : undefined;
            inputs["trainingPhrases"] = args ? args.trainingPhrases : undefined;
            inputs["webhookState"] = args ? args.webhookState : undefined;
            inputs["followupIntentInfo"] = undefined /*out*/;
            inputs["rootFollowupIntentName"] = undefined /*out*/;
        } else {
            inputs["action"] = undefined /*out*/;
            inputs["defaultResponsePlatforms"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["endInteraction"] = undefined /*out*/;
            inputs["events"] = undefined /*out*/;
            inputs["followupIntentInfo"] = undefined /*out*/;
            inputs["inputContextNames"] = undefined /*out*/;
            inputs["isFallback"] = undefined /*out*/;
            inputs["liveAgentHandoff"] = undefined /*out*/;
            inputs["messages"] = undefined /*out*/;
            inputs["mlDisabled"] = undefined /*out*/;
            inputs["mlEnabled"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["outputContexts"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["parentFollowupIntentName"] = undefined /*out*/;
            inputs["priority"] = undefined /*out*/;
            inputs["resetContexts"] = undefined /*out*/;
            inputs["rootFollowupIntentName"] = undefined /*out*/;
            inputs["trainingPhrases"] = undefined /*out*/;
            inputs["webhookState"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AgentIntent.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AgentIntent resource.
 */
export interface AgentIntentArgs {
    /**
     * Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
     */
    readonly action?: pulumi.Input<string>;
    /**
     * Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
     */
    readonly defaultResponsePlatforms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. The name of this intent.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
     */
    readonly endInteraction?: pulumi.Input<boolean>;
    /**
     * Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
     */
    readonly events?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. The list of context names required for this intent to be triggered. Formats: - `projects//agent/sessions/-/contexts/` - `projects//locations//agent/sessions/-/contexts/`
     */
    readonly inputContextNames?: pulumi.Input<pulumi.Input<string>[]>;
    readonly intentView?: pulumi.Input<string>;
    /**
     * Optional. Indicates whether this is a fallback intent.
     */
    readonly isFallback?: pulumi.Input<boolean>;
    readonly languageCode?: pulumi.Input<string>;
    /**
     * Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
     */
    readonly liveAgentHandoff?: pulumi.Input<boolean>;
    readonly location: pulumi.Input<string>;
    /**
     * Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
     */
    readonly messages?: pulumi.Input<pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1IntentMessageArgs>[]>;
    /**
     * Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
     */
    readonly mlDisabled?: pulumi.Input<boolean>;
    /**
     * Optional. Indicates whether Machine Learning is enabled for the intent. Note: If `ml_enabled` setting is set to false, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off. DEPRECATED! Please use `ml_disabled` field instead. NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false, then the default value is determined as follows: - Before April 15th, 2018 the default is: ml_enabled = false / ml_disabled = true. - After April 15th, 2018 the default is: ml_enabled = true / ml_disabled = false.
     */
    readonly mlEnabled?: pulumi.Input<boolean>;
    /**
     * Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Supported formats: - `projects//agent/intents/` - `projects//locations//agent/intents/`
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
     */
    readonly outputContexts?: pulumi.Input<pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1ContextArgs>[]>;
    /**
     * Optional. The collection of parameters associated with the intent.
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1IntentParameterArgs>[]>;
    /**
     * Optional. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
     */
    readonly parentFollowupIntentName?: pulumi.Input<string>;
    /**
     * Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
     */
    readonly priority?: pulumi.Input<number>;
    readonly project: pulumi.Input<string>;
    /**
     * Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
     */
    readonly resetContexts?: pulumi.Input<boolean>;
    /**
     * Optional. The collection of examples that the agent is trained on.
     */
    readonly trainingPhrases?: pulumi.Input<pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1IntentTrainingPhraseArgs>[]>;
    /**
     * Optional. Indicates whether webhooks are enabled for the intent.
     */
    readonly webhookState?: pulumi.Input<string>;
}
