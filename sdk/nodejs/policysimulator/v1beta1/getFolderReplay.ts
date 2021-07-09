// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets the specified Replay. Each `Replay` is available for at least 7 days.
 */
export function getFolderReplay(args: GetFolderReplayArgs, opts?: pulumi.InvokeOptions): Promise<GetFolderReplayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:policysimulator/v1beta1:getFolderReplay", {
        "folderId": args.folderId,
        "location": args.location,
        "replayId": args.replayId,
    }, opts);
}

export interface GetFolderReplayArgs {
    folderId: string;
    location: string;
    replayId: string;
}

export interface GetFolderReplayResult {
    /**
     * The configuration used for the `Replay`.
     */
    readonly config: outputs.policysimulator.v1beta1.GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse;
    /**
     * The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
     */
    readonly name: string;
    /**
     * Summary statistics about the replayed log entries.
     */
    readonly resultsSummary: outputs.policysimulator.v1beta1.GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse;
    /**
     * The current state of the `Replay`.
     */
    readonly state: string;
}
