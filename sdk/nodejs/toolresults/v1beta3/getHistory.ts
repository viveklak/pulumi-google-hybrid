// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets a History. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the History does not exist
 */
export function getHistory(args: GetHistoryArgs, opts?: pulumi.InvokeOptions): Promise<GetHistoryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:toolresults/v1beta3:getHistory", {
        "historyId": args.historyId,
        "project": args.project,
    }, opts);
}

export interface GetHistoryArgs {
    historyId: string;
    project?: string;
}

export interface GetHistoryResult {
    /**
     * A short human-readable (plain text) name to display in the UI. Maximum of 100 characters. - In response: present if set during create. - In create request: optional
     */
    readonly displayName: string;
    /**
     * A unique identifier within a project for this History. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create request: never set
     */
    readonly historyId: string;
    /**
     * A name to uniquely identify a history within a project. Maximum of 200 characters. - In response always set - In create request: always set
     */
    readonly name: string;
    /**
     * The platform of the test history. - In response: always set. Returns the platform of the last execution if unknown.
     */
    readonly testPlatform: string;
}

export function getHistoryOutput(args: GetHistoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHistoryResult> {
    return pulumi.output(args).apply(a => getHistory(a, opts))
}

export interface GetHistoryOutputArgs {
    historyId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
