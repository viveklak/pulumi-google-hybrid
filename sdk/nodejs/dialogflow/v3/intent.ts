// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an intent in the specified agent.
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
    public static readonly __pulumiType = 'google-native:dialogflow/v3:Intent';

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
     * Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Required. The human-readable name of the intent, unique within the agent.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
     */
    public readonly isFallback!: pulumi.Output<boolean>;
    /**
     * The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys." is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys.head * sys.contextual The above labels do not require value. "sys.head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The collection of parameters associated with the intent.
     */
    public readonly parameters!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3IntentParameterResponse[]>;
    /**
     * The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The collection of training phrases the agent is trained on to identify the intent.
     */
    public readonly trainingPhrases!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3IntentTrainingPhraseResponse[]>;

    /**
     * Create a Intent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["agentId"] = args ? args.agentId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["isFallback"] = args ? args.isFallback : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["languageCode"] = args ? args.languageCode : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["trainingPhrases"] = args ? args.trainingPhrases : undefined;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["isFallback"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["priority"] = undefined /*out*/;
            inputs["trainingPhrases"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Intent.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Intent resource.
 */
export interface IntentArgs {
    agentId: pulumi.Input<string>;
    /**
     * Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Required. The human-readable name of the intent, unique within the agent.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
     */
    isFallback?: pulumi.Input<boolean>;
    /**
     * The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys." is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys.head * sys.contextual The above labels do not require value. "sys.head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    languageCode?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    /**
     * The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
     */
    name?: pulumi.Input<string>;
    /**
     * The collection of parameters associated with the intent.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3IntentParameterArgs>[]>;
    /**
     * The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
     */
    priority?: pulumi.Input<number>;
    project: pulumi.Input<string>;
    /**
     * The collection of training phrases the agent is trained on to identify the intent.
     */
    trainingPhrases?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs>[]>;
}