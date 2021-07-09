// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified page.
 */
export function getPage(args: GetPageArgs, opts?: pulumi.InvokeOptions): Promise<GetPageResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:dialogflow/v3:getPage", {
        "agentId": args.agentId,
        "flowId": args.flowId,
        "languageCode": args.languageCode,
        "location": args.location,
        "pageId": args.pageId,
        "project": args.project,
    }, opts);
}

export interface GetPageArgs {
    agentId: string;
    flowId: string;
    languageCode?: string;
    location: string;
    pageId: string;
    project: string;
}

export interface GetPageResult {
    /**
     * The human-readable name of the page, unique within the agent.
     */
    readonly displayName: string;
    /**
     * The fulfillment to call when the session is entering the page.
     */
    readonly entryFulfillment: outputs.dialogflow.v3.GoogleCloudDialogflowCxV3FulfillmentResponse;
    /**
     * Handlers associated with the page to handle events such as webhook errors, no match or no input.
     */
    readonly eventHandlers: outputs.dialogflow.v3.GoogleCloudDialogflowCxV3EventHandlerResponse[];
    /**
     * The form associated with the page, used for collecting parameters relevant to the page.
     */
    readonly form: outputs.dialogflow.v3.GoogleCloudDialogflowCxV3FormResponse;
    /**
     * The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
     */
    readonly name: string;
    /**
     * Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
     */
    readonly transitionRouteGroups: string[];
    /**
     * A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
     */
    readonly transitionRoutes: outputs.dialogflow.v3.GoogleCloudDialogflowCxV3TransitionRouteResponse[];
}
