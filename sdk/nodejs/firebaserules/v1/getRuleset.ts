// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Get a `Ruleset` by name including the full `Source` contents.
 */
export function getRuleset(args: GetRulesetArgs, opts?: pulumi.InvokeOptions): Promise<GetRulesetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:firebaserules/v1:getRuleset", {
        "project": args.project,
        "rulesetId": args.rulesetId,
    }, opts);
}

export interface GetRulesetArgs {
    project?: string;
    rulesetId: string;
}

export interface GetRulesetResult {
    /**
     * Time the `Ruleset` was created.
     */
    readonly createTime: string;
    /**
     * The metadata for this ruleset.
     */
    readonly metadata: outputs.firebaserules.v1.MetadataResponse;
    /**
     * Name of the `Ruleset`. The ruleset_id is auto generated by the service. Format: `projects/{project_id}/rulesets/{ruleset_id}`
     */
    readonly name: string;
    /**
     * `Source` for the `Ruleset`.
     */
    readonly source: outputs.firebaserules.v1.SourceResponse;
}

export function getRulesetOutput(args: GetRulesetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRulesetResult> {
    return pulumi.output(args).apply(a => getRuleset(a, opts))
}

export interface GetRulesetOutputArgs {
    project?: pulumi.Input<string>;
    rulesetId: pulumi.Input<string>;
}
