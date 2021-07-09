// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an agent in the specified location.
 */
export class Agent extends pulumi.CustomResource {
    /**
     * Get an existing Agent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Agent {
        return new Agent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3beta1:Agent';

    /**
     * Returns true if the given object is an instance of Agent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Agent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Agent.__pulumiType;
    }

    /**
     * The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted [Web Demo](https://cloud.google.com/dialogflow/docs/integrations/web-demo) integration.
     */
    public readonly avatarUri!: pulumi.Output<string>;
    /**
     * Immutable. The default language of the agent as a language tag. See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes. This field cannot be set by the Agents.UpdateAgent method.
     */
    public readonly defaultLanguageCode!: pulumi.Output<string>;
    /**
     * The description of the agent. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The human-readable name of the agent, unique within the location.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Indicates if automatic spell correction is enabled in detect intent requests.
     */
    public readonly enableSpellCorrection!: pulumi.Output<boolean>;
    /**
     * Indicates if stackdriver logging is enabled for the agent.
     */
    public readonly enableStackdriverLogging!: pulumi.Output<boolean>;
    /**
     * The unique identifier of the agent. Required for the Agents.UpdateAgent method. Agents.CreateAgent populates the name automatically. Format: `projects//locations//agents/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of the SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
     */
    public readonly securitySettings!: pulumi.Output<string>;
    /**
     * Speech recognition related settings.
     */
    public readonly speechToTextSettings!: pulumi.Output<outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1SpeechToTextSettingsResponse>;
    /**
     * Immutable. Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: `projects//locations//agents//flows/`.
     */
    public readonly startFlow!: pulumi.Output<string>;
    /**
     * The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
     */
    public readonly timeZone!: pulumi.Output<string>;

    /**
     * Create a Agent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.timeZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeZone'");
            }
            inputs["avatarUri"] = args ? args.avatarUri : undefined;
            inputs["defaultLanguageCode"] = args ? args.defaultLanguageCode : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enableSpellCorrection"] = args ? args.enableSpellCorrection : undefined;
            inputs["enableStackdriverLogging"] = args ? args.enableStackdriverLogging : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["securitySettings"] = args ? args.securitySettings : undefined;
            inputs["speechToTextSettings"] = args ? args.speechToTextSettings : undefined;
            inputs["startFlow"] = args ? args.startFlow : undefined;
            inputs["timeZone"] = args ? args.timeZone : undefined;
        } else {
            inputs["avatarUri"] = undefined /*out*/;
            inputs["defaultLanguageCode"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["enableSpellCorrection"] = undefined /*out*/;
            inputs["enableStackdriverLogging"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["securitySettings"] = undefined /*out*/;
            inputs["speechToTextSettings"] = undefined /*out*/;
            inputs["startFlow"] = undefined /*out*/;
            inputs["timeZone"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Agent.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Agent resource.
 */
export interface AgentArgs {
    /**
     * The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted [Web Demo](https://cloud.google.com/dialogflow/docs/integrations/web-demo) integration.
     */
    avatarUri?: pulumi.Input<string>;
    /**
     * Immutable. The default language of the agent as a language tag. See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes. This field cannot be set by the Agents.UpdateAgent method.
     */
    defaultLanguageCode?: pulumi.Input<string>;
    /**
     * The description of the agent. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name of the agent, unique within the location.
     */
    displayName: pulumi.Input<string>;
    /**
     * Indicates if automatic spell correction is enabled in detect intent requests.
     */
    enableSpellCorrection?: pulumi.Input<boolean>;
    /**
     * Indicates if stackdriver logging is enabled for the agent.
     */
    enableStackdriverLogging?: pulumi.Input<boolean>;
    location: pulumi.Input<string>;
    /**
     * The unique identifier of the agent. Required for the Agents.UpdateAgent method. Agents.CreateAgent populates the name automatically. Format: `projects//locations//agents/`.
     */
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    /**
     * Name of the SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
     */
    securitySettings?: pulumi.Input<string>;
    /**
     * Speech recognition related settings.
     */
    speechToTextSettings?: pulumi.Input<inputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1SpeechToTextSettingsArgs>;
    /**
     * Immutable. Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: `projects//locations//agents//flows/`.
     */
    startFlow?: pulumi.Input<string>;
    /**
     * The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
     */
    timeZone: pulumi.Input<string>;
}
