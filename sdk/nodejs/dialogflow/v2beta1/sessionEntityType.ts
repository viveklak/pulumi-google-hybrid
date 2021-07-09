// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a session entity type. If the specified session entity type already exists, overrides the session entity type. This method doesn't work with Google Assistant integration. Contact Dialogflow support if you need to use session entities with Google Assistant integration.
 */
export class SessionEntityType extends pulumi.CustomResource {
    /**
     * Get an existing SessionEntityType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SessionEntityType {
        return new SessionEntityType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2beta1:SessionEntityType';

    /**
     * Returns true if the given object is an instance of SessionEntityType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SessionEntityType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SessionEntityType.__pulumiType;
    }

    /**
     * The collection of entities associated with this session entity type.
     */
    public readonly entities!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1EntityTypeEntityResponse[]>;
    /**
     * Indicates whether the additional data should override or supplement the custom entity type definition.
     */
    public readonly entityOverrideMode!: pulumi.Output<string>;
    /**
     * The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a SessionEntityType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SessionEntityTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.entities === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entities'");
            }
            if ((!args || args.entityOverrideMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entityOverrideMode'");
            }
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.sessionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sessionId'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            inputs["entities"] = args ? args.entities : undefined;
            inputs["entityOverrideMode"] = args ? args.entityOverrideMode : undefined;
            inputs["environmentId"] = args ? args.environmentId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sessionId"] = args ? args.sessionId : undefined;
            inputs["userId"] = args ? args.userId : undefined;
        } else {
            inputs["entities"] = undefined /*out*/;
            inputs["entityOverrideMode"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SessionEntityType.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SessionEntityType resource.
 */
export interface SessionEntityTypeArgs {
    /**
     * The collection of entities associated with this session entity type.
     */
    entities: pulumi.Input<pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1EntityTypeEntityArgs>[]>;
    /**
     * Indicates whether the additional data should override or supplement the custom entity type definition.
     */
    entityOverrideMode: pulumi.Input<enums.dialogflow.v2beta1.SessionEntityTypeEntityOverrideMode>;
    environmentId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    /**
     * The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
     */
    name: pulumi.Input<string>;
    project: pulumi.Input<string>;
    sessionId: pulumi.Input<string>;
    userId: pulumi.Input<string>;
}
