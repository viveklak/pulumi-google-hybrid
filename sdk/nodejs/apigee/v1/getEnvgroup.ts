// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets an environment group.
 */
export function getEnvgroup(args: GetEnvgroupArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvgroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:apigee/v1:getEnvgroup", {
        "envgroupId": args.envgroupId,
        "organizationId": args.organizationId,
    }, opts);
}

export interface GetEnvgroupArgs {
    envgroupId: string;
    organizationId: string;
}

export interface GetEnvgroupResult {
    /**
     * The time at which the environment group was created as milliseconds since epoch.
     */
    readonly createdAt: string;
    /**
     * Host names for this environment group.
     */
    readonly hostnames: string[];
    /**
     * The time at which the environment group was last updated as milliseconds since epoch.
     */
    readonly lastModifiedAt: string;
    /**
     * ID of the environment group.
     */
    readonly name: string;
    /**
     * State of the environment group. Values other than ACTIVE means the resource is not ready to use.
     */
    readonly state: string;
}
