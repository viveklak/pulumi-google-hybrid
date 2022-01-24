// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an entity type in the specified agent. Note: You should always train a flow prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/cx/docs/concept/training).
 */
export class EntityType extends pulumi.CustomResource {
    /**
     * Get an existing EntityType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EntityType {
        return new EntityType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3:EntityType';

    /**
     * Returns true if the given object is an instance of EntityType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EntityType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EntityType.__pulumiType;
    }

    /**
     * Indicates whether the entity type can be automatically expanded.
     */
    public readonly autoExpansionMode!: pulumi.Output<string>;
    /**
     * The human-readable name of the entity type, unique within the agent.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Enables fuzzy entity extraction during classification.
     */
    public readonly enableFuzzyExtraction!: pulumi.Output<boolean>;
    /**
     * The collection of entity entries associated with the entity type.
     */
    public readonly entities!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3EntityTypeEntityResponse[]>;
    /**
     * Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
     */
    public readonly excludedPhrases!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3EntityTypeExcludedPhraseResponse[]>;
    /**
     * Indicates the kind of entity type.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
     */
    public readonly redact!: pulumi.Output<boolean>;

    /**
     * Create a EntityType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntityTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            resourceInputs["agentId"] = args ? args.agentId : undefined;
            resourceInputs["autoExpansionMode"] = args ? args.autoExpansionMode : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enableFuzzyExtraction"] = args ? args.enableFuzzyExtraction : undefined;
            resourceInputs["entities"] = args ? args.entities : undefined;
            resourceInputs["excludedPhrases"] = args ? args.excludedPhrases : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["languageCode"] = args ? args.languageCode : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["redact"] = args ? args.redact : undefined;
        } else {
            resourceInputs["autoExpansionMode"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["enableFuzzyExtraction"] = undefined /*out*/;
            resourceInputs["entities"] = undefined /*out*/;
            resourceInputs["excludedPhrases"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["redact"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EntityType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EntityType resource.
 */
export interface EntityTypeArgs {
    agentId: pulumi.Input<string>;
    /**
     * Indicates whether the entity type can be automatically expanded.
     */
    autoExpansionMode?: pulumi.Input<enums.dialogflow.v3.EntityTypeAutoExpansionMode>;
    /**
     * The human-readable name of the entity type, unique within the agent.
     */
    displayName: pulumi.Input<string>;
    /**
     * Enables fuzzy entity extraction during classification.
     */
    enableFuzzyExtraction?: pulumi.Input<boolean>;
    /**
     * The collection of entity entries associated with the entity type.
     */
    entities?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3EntityTypeEntityArgs>[]>;
    /**
     * Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
     */
    excludedPhrases?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3EntityTypeExcludedPhraseArgs>[]>;
    /**
     * Indicates the kind of entity type.
     */
    kind: pulumi.Input<enums.dialogflow.v3.EntityTypeKind>;
    languageCode?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
     */
    redact?: pulumi.Input<boolean>;
}
