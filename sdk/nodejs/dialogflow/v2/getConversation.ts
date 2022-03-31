// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specific conversation.
 */
export function getConversation(args: GetConversationArgs, opts?: pulumi.InvokeOptions): Promise<GetConversationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dialogflow/v2:getConversation", {
        "conversationId": args.conversationId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetConversationArgs {
    conversationId: string;
    location: string;
    project?: string;
}

export interface GetConversationResult {
    /**
     * The Conversation Profile to be used to configure this Conversation. This field cannot be updated. Format: `projects//locations//conversationProfiles/`.
     */
    readonly conversationProfile: string;
    /**
     * The stage of a conversation. It indicates whether the virtual agent or a human agent is handling the conversation. If the conversation is created with the conversation profile that has Dialogflow config set, defaults to ConversationStage.VIRTUAL_AGENT_STAGE; Otherwise, defaults to ConversationStage.HUMAN_ASSIST_STAGE. If the conversation is created with the conversation profile that has Dialogflow config set but explicitly sets conversation_stage to ConversationStage.HUMAN_ASSIST_STAGE, it skips ConversationStage.VIRTUAL_AGENT_STAGE stage and directly goes to ConversationStage.HUMAN_ASSIST_STAGE.
     */
    readonly conversationStage: string;
    /**
     * The time the conversation was finished.
     */
    readonly endTime: string;
    /**
     * The current state of the Conversation.
     */
    readonly lifecycleState: string;
    /**
     * The unique identifier of this conversation. Format: `projects//locations//conversations/`.
     */
    readonly name: string;
    /**
     * It will not be empty if the conversation is to be connected over telephony.
     */
    readonly phoneNumber: outputs.dialogflow.v2.GoogleCloudDialogflowV2ConversationPhoneNumberResponse;
    /**
     * The time the conversation was started.
     */
    readonly startTime: string;
}

export function getConversationOutput(args: GetConversationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConversationResult> {
    return pulumi.output(args).apply(a => getConversation(a, opts))
}

export interface GetConversationOutputArgs {
    conversationId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
