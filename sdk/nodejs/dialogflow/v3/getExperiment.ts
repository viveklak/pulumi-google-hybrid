// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified Experiment.
 */
export function getExperiment(args: GetExperimentArgs, opts?: pulumi.InvokeOptions): Promise<GetExperimentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:dialogflow/v3:getExperiment", {
        "agentId": args.agentId,
        "environmentId": args.environmentId,
        "experimentId": args.experimentId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetExperimentArgs {
    agentId: string;
    environmentId: string;
    experimentId: string;
    location: string;
    project: string;
}

export interface GetExperimentResult {
    /**
     * Creation time of this experiment.
     */
    readonly createTime: string;
    /**
     * The definition of the experiment.
     */
    readonly definition: outputs.dialogflow.v3.GoogleCloudDialogflowCxV3ExperimentDefinitionResponse;
    /**
     * The human-readable description of the experiment.
     */
    readonly description: string;
    /**
     * The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
     */
    readonly displayName: string;
    /**
     * End time of this experiment.
     */
    readonly endTime: string;
    /**
     * Maximum number of days to run the experiment/rollout. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
     */
    readonly experimentLength: string;
    /**
     * Last update time of this experiment.
     */
    readonly lastUpdateTime: string;
    /**
     * The name of the experiment. Format: projects//locations//agents//environments//experiments/..
     */
    readonly name: string;
    /**
     * Inference result of the experiment.
     */
    readonly result: outputs.dialogflow.v3.GoogleCloudDialogflowCxV3ExperimentResultResponse;
    /**
     * Start time of this experiment.
     */
    readonly startTime: string;
    /**
     * The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
     */
    readonly state: string;
    /**
     * The history of updates to the experiment variants.
     */
    readonly variantsHistory: outputs.dialogflow.v3.GoogleCloudDialogflowCxV3VariantsHistoryResponse[];
}
