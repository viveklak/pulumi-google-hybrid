// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns a function with the given name from the requested project.
 */
export function getFunction(args: GetFunctionArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:cloudfunctions/v2beta:getFunction", {
        "functionId": args.functionId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetFunctionArgs {
    functionId: string;
    location: string;
    project?: string;
}

export interface GetFunctionResult {
    /**
     * Describes the Build step of the function that builds a container from the given source.
     */
    readonly buildConfig: outputs.cloudfunctions.v2beta.BuildConfigResponse;
    /**
     * User-provided description of a function.
     */
    readonly description: string;
    /**
     * Describe whether the function is gen1 or gen2.
     */
    readonly environment: string;
    /**
     * An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
     */
    readonly eventTrigger: outputs.cloudfunctions.v2beta.EventTriggerResponse;
    /**
     * Labels associated with this Cloud Function.
     */
    readonly labels: {[key: string]: string};
    /**
     * A user-defined name of the function. Function names must be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`
     */
    readonly name: string;
    /**
     * Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
     */
    readonly serviceConfig: outputs.cloudfunctions.v2beta.ServiceConfigResponse;
    /**
     * State of the function.
     */
    readonly state: string;
    /**
     * State Messages for this Cloud Function.
     */
    readonly stateMessages: outputs.cloudfunctions.v2beta.GoogleCloudFunctionsV2betaStateMessageResponse[];
    /**
     * The last update timestamp of a Cloud Function.
     */
    readonly updateTime: string;
}

export function getFunctionOutput(args: GetFunctionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionResult> {
    return pulumi.output(args).apply(a => getFunction(a, opts))
}

export interface GetFunctionOutputArgs {
    functionId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
