// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified context.
 */
export function getContext(args: GetContextArgs, opts?: pulumi.InvokeOptions): Promise<GetContextResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dialogflow/v2:getContext", {
        "contextId": args.contextId,
        "environmentId": args.environmentId,
        "location": args.location,
        "project": args.project,
        "sessionId": args.sessionId,
        "userId": args.userId,
    }, opts);
}

export interface GetContextArgs {
    contextId: string;
    environmentId: string;
    location: string;
    project?: string;
    sessionId: string;
    userId: string;
}

export interface GetContextResult {
    /**
     * Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
     */
    readonly lifespanCount: number;
    /**
     * The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
     */
    readonly name: string;
    /**
     * Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
     */
    readonly parameters: {[key: string]: string};
}

export function getContextOutput(args: GetContextOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetContextResult> {
    return pulumi.output(args).apply(a => getContext(a, opts))
}

export interface GetContextOutputArgs {
    contextId: pulumi.Input<string>;
    environmentId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    sessionId: pulumi.Input<string>;
    userId: pulumi.Input<string>;
}
