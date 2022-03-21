// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified flow.
 */
export function getFlow(args: GetFlowArgs, opts?: pulumi.InvokeOptions): Promise<GetFlowResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dialogflow/v3beta1:getFlow", {
        "agentId": args.agentId,
        "flowId": args.flowId,
        "languageCode": args.languageCode,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetFlowArgs {
    agentId: string;
    flowId: string;
    languageCode?: string;
    location: string;
    project?: string;
}

export interface GetFlowResult {
    /**
     * The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    readonly description: string;
    /**
     * The human-readable name of the flow.
     */
    readonly displayName: string;
    /**
     * A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
     */
    readonly eventHandlers: outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1EventHandlerResponse[];
    /**
     * The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
     */
    readonly name: string;
    /**
     * NLU related settings of the flow.
     */
    readonly nluSettings: outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1NluSettingsResponse;
    /**
     * A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
     */
    readonly transitionRouteGroups: string[];
    /**
     * A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
     */
    readonly transitionRoutes: outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1TransitionRouteResponse[];
}

export function getFlowOutput(args: GetFlowOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlowResult> {
    return pulumi.output(args).apply(a => getFlow(a, opts))
}

export interface GetFlowOutputArgs {
    agentId: pulumi.Input<string>;
    flowId: pulumi.Input<string>;
    languageCode?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
