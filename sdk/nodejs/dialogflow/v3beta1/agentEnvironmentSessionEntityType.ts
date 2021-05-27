// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a session entity type.
 */
export class AgentEnvironmentSessionEntityType extends pulumi.CustomResource {
    /**
     * Get an existing AgentEnvironmentSessionEntityType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AgentEnvironmentSessionEntityType {
        return new AgentEnvironmentSessionEntityType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3beta1:AgentEnvironmentSessionEntityType';

    /**
     * Returns true if the given object is an instance of AgentEnvironmentSessionEntityType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentEnvironmentSessionEntityType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentEnvironmentSessionEntityType.__pulumiType;
    }

    /**
     * Required. The collection of entities to override or supplement the custom entity type.
     */
    public readonly entities!: pulumi.Output<outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponse[]>;
    /**
     * Required. Indicates whether the additional data should override or supplement the custom entity type definition.
     */
    public readonly entityOverrideMode!: pulumi.Output<string>;
    /**
     * Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a AgentEnvironmentSessionEntityType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentEnvironmentSessionEntityTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.sessionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sessionId'");
            }
            inputs["agentId"] = args ? args.agentId : undefined;
            inputs["entities"] = args ? args.entities : undefined;
            inputs["entityOverrideMode"] = args ? args.entityOverrideMode : undefined;
            inputs["environmentId"] = args ? args.environmentId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sessionId"] = args ? args.sessionId : undefined;
        } else {
            inputs["entities"] = undefined /*out*/;
            inputs["entityOverrideMode"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AgentEnvironmentSessionEntityType.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AgentEnvironmentSessionEntityType resource.
 */
export interface AgentEnvironmentSessionEntityTypeArgs {
    readonly agentId: pulumi.Input<string>;
    /**
     * Required. The collection of entities to override or supplement the custom entity type.
     */
    readonly entities?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs>[]>;
    /**
     * Required. Indicates whether the additional data should override or supplement the custom entity type definition.
     */
    readonly entityOverrideMode?: pulumi.Input<string>;
    readonly environmentId: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    /**
     * Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    readonly sessionId: pulumi.Input<string>;
}
