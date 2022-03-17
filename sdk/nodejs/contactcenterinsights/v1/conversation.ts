// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a conversation.
 */
export class Conversation extends pulumi.CustomResource {
    /**
     * Get an existing Conversation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Conversation {
        return new Conversation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:contactcenterinsights/v1:Conversation';

    /**
     * Returns true if the given object is an instance of Conversation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Conversation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Conversation.__pulumiType;
    }

    /**
     * An opaque, user-specified string representing the human agent who handled the conversation.
     */
    public readonly agentId!: pulumi.Output<string>;
    /**
     * Call-specific metadata.
     */
    public readonly callMetadata!: pulumi.Output<outputs.contactcenterinsights.v1.GoogleCloudContactcenterinsightsV1ConversationCallMetadataResponse>;
    /**
     * The time at which the conversation was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The source of the audio and transcription for the conversation.
     */
    public readonly dataSource!: pulumi.Output<outputs.contactcenterinsights.v1.GoogleCloudContactcenterinsightsV1ConversationDataSourceResponse>;
    /**
     * All the matched Dialogflow intents in the call. The key corresponds to a Dialogflow intent, format: projects/{project}/agent/{agent}/intents/{intent}
     */
    public /*out*/ readonly dialogflowIntents!: pulumi.Output<{[key: string]: string}>;
    /**
     * The duration of the conversation.
     */
    public /*out*/ readonly duration!: pulumi.Output<string>;
    /**
     * The time at which this conversation should expire. After this time, the conversation data and any associated analyses will be deleted.
     */
    public readonly expireTime!: pulumi.Output<string>;
    /**
     * A map for the user to specify any custom fields. A maximum of 20 labels per conversation is allowed, with a maximum of 256 characters per entry.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * A user-specified language code for the conversation.
     */
    public readonly languageCode!: pulumi.Output<string>;
    /**
     * The conversation's latest analysis, if one exists.
     */
    public /*out*/ readonly latestAnalysis!: pulumi.Output<outputs.contactcenterinsights.v1.GoogleCloudContactcenterinsightsV1AnalysisResponse>;
    /**
     * Immutable. The conversation medium, if unspecified will default to PHONE_CALL.
     */
    public readonly medium!: pulumi.Output<string>;
    /**
     * Immutable. The resource name of the conversation. Format: projects/{project}/locations/{location}/conversations/{conversation}
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Obfuscated user ID which the customer sent to us.
     */
    public readonly obfuscatedUserId!: pulumi.Output<string>;
    /**
     * The annotations that were generated during the customer and agent interaction.
     */
    public /*out*/ readonly runtimeAnnotations!: pulumi.Output<outputs.contactcenterinsights.v1.GoogleCloudContactcenterinsightsV1RuntimeAnnotationResponse[]>;
    /**
     * The time at which the conversation started.
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * The conversation transcript.
     */
    public /*out*/ readonly transcript!: pulumi.Output<outputs.contactcenterinsights.v1.GoogleCloudContactcenterinsightsV1ConversationTranscriptResponse>;
    /**
     * Input only. The TTL for this resource. If specified, then this TTL will be used to calculate the expire time.
     */
    public readonly ttl!: pulumi.Output<string>;
    /**
     * The number of turns in the conversation.
     */
    public /*out*/ readonly turnCount!: pulumi.Output<number>;
    /**
     * The most recent time at which the conversation was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Conversation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ConversationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["agentId"] = args ? args.agentId : undefined;
            resourceInputs["callMetadata"] = args ? args.callMetadata : undefined;
            resourceInputs["conversationId"] = args ? args.conversationId : undefined;
            resourceInputs["dataSource"] = args ? args.dataSource : undefined;
            resourceInputs["expireTime"] = args ? args.expireTime : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["languageCode"] = args ? args.languageCode : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["medium"] = args ? args.medium : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["obfuscatedUserId"] = args ? args.obfuscatedUserId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dialogflowIntents"] = undefined /*out*/;
            resourceInputs["duration"] = undefined /*out*/;
            resourceInputs["latestAnalysis"] = undefined /*out*/;
            resourceInputs["runtimeAnnotations"] = undefined /*out*/;
            resourceInputs["transcript"] = undefined /*out*/;
            resourceInputs["turnCount"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["agentId"] = undefined /*out*/;
            resourceInputs["callMetadata"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dataSource"] = undefined /*out*/;
            resourceInputs["dialogflowIntents"] = undefined /*out*/;
            resourceInputs["duration"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["languageCode"] = undefined /*out*/;
            resourceInputs["latestAnalysis"] = undefined /*out*/;
            resourceInputs["medium"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["obfuscatedUserId"] = undefined /*out*/;
            resourceInputs["runtimeAnnotations"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["transcript"] = undefined /*out*/;
            resourceInputs["ttl"] = undefined /*out*/;
            resourceInputs["turnCount"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Conversation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Conversation resource.
 */
export interface ConversationArgs {
    /**
     * An opaque, user-specified string representing the human agent who handled the conversation.
     */
    agentId?: pulumi.Input<string>;
    /**
     * Call-specific metadata.
     */
    callMetadata?: pulumi.Input<inputs.contactcenterinsights.v1.GoogleCloudContactcenterinsightsV1ConversationCallMetadataArgs>;
    /**
     * A unique ID for the new conversation. This ID will become the final component of the conversation's resource name. If no ID is specified, a server-generated ID will be used. This value should be 4-64 characters and must match the regular expression `^[a-z0-9-]{4,64}$`. Valid characters are `a-z-`
     */
    conversationId?: pulumi.Input<string>;
    /**
     * The source of the audio and transcription for the conversation.
     */
    dataSource?: pulumi.Input<inputs.contactcenterinsights.v1.GoogleCloudContactcenterinsightsV1ConversationDataSourceArgs>;
    /**
     * The time at which this conversation should expire. After this time, the conversation data and any associated analyses will be deleted.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * A map for the user to specify any custom fields. A maximum of 20 labels per conversation is allowed, with a maximum of 256 characters per entry.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A user-specified language code for the conversation.
     */
    languageCode?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Immutable. The conversation medium, if unspecified will default to PHONE_CALL.
     */
    medium?: pulumi.Input<enums.contactcenterinsights.v1.ConversationMedium>;
    /**
     * Immutable. The resource name of the conversation. Format: projects/{project}/locations/{location}/conversations/{conversation}
     */
    name?: pulumi.Input<string>;
    /**
     * Obfuscated user ID which the customer sent to us.
     */
    obfuscatedUserId?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The time at which the conversation started.
     */
    startTime?: pulumi.Input<string>;
    /**
     * Input only. The TTL for this resource. If specified, then this TTL will be used to calculate the expire time.
     */
    ttl?: pulumi.Input<string>;
}
