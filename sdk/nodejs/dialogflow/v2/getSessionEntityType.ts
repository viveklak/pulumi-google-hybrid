// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified session entity type. This method doesn't work with Google Assistant integration. Contact Dialogflow support if you need to use session entities with Google Assistant integration.
 */
export function getSessionEntityType(args: GetSessionEntityTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetSessionEntityTypeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dialogflow/v2:getSessionEntityType", {
        "entityTypeId": args.entityTypeId,
        "environmentId": args.environmentId,
        "location": args.location,
        "project": args.project,
        "sessionId": args.sessionId,
        "userId": args.userId,
    }, opts);
}

export interface GetSessionEntityTypeArgs {
    entityTypeId: string;
    environmentId: string;
    location: string;
    project?: string;
    sessionId: string;
    userId: string;
}

export interface GetSessionEntityTypeResult {
    /**
     * The collection of entities associated with this session entity type.
     */
    readonly entities: outputs.dialogflow.v2.GoogleCloudDialogflowV2EntityTypeEntityResponse[];
    /**
     * Indicates whether the additional data should override or supplement the custom entity type definition.
     */
    readonly entityOverrideMode: string;
    /**
     * The unique identifier of this session entity type. Format: `projects//agent/sessions//entityTypes/`, or `projects//agent/environments//users//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
     */
    readonly name: string;
}

export function getSessionEntityTypeOutput(args: GetSessionEntityTypeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSessionEntityTypeResult> {
    return pulumi.output(args).apply(a => getSessionEntityType(a, opts))
}

export interface GetSessionEntityTypeOutputArgs {
    entityTypeId: pulumi.Input<string>;
    environmentId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    sessionId: pulumi.Input<string>;
    userId: pulumi.Input<string>;
}
