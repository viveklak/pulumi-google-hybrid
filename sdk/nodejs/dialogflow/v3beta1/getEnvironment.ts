// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified Environment.
 */
export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dialogflow/v3beta1:getEnvironment", {
        "agentId": args.agentId,
        "environmentId": args.environmentId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetEnvironmentArgs {
    agentId: string;
    environmentId: string;
    location: string;
    project?: string;
}

export interface GetEnvironmentResult {
    /**
     * The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    readonly description: string;
    /**
     * The human-readable name of the environment (unique in an agent). Limit of 64 characters.
     */
    readonly displayName: string;
    /**
     * The name of the environment. Format: `projects//locations//agents//environments/`.
     */
    readonly name: string;
    /**
     * The test cases config for continuous tests of this environment.
     */
    readonly testCasesConfig: outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1EnvironmentTestCasesConfigResponse;
    /**
     * Update time of this environment.
     */
    readonly updateTime: string;
    /**
     * A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
     */
    readonly versionConfigs: outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1EnvironmentVersionConfigResponse[];
}

export function getEnvironmentOutput(args: GetEnvironmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEnvironmentResult> {
    return pulumi.output(args).apply(a => getEnvironment(a, opts))
}

export interface GetEnvironmentOutputArgs {
    agentId: pulumi.Input<string>;
    environmentId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
