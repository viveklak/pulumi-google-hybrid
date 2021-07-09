// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a keystore or truststore.
 */
export function getKeystore(args: GetKeystoreArgs, opts?: pulumi.InvokeOptions): Promise<GetKeystoreResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:apigee/v1:getKeystore", {
        "environmentId": args.environmentId,
        "keystoreId": args.keystoreId,
        "organizationId": args.organizationId,
    }, opts);
}

export interface GetKeystoreArgs {
    environmentId: string;
    keystoreId: string;
    organizationId: string;
}

export interface GetKeystoreResult {
    /**
     * Aliases in this keystore.
     */
    readonly aliases: string[];
    /**
     * Resource ID for this keystore. Values must match the regular expression `[\w[:space:]-.]{1,255}`.
     */
    readonly name: string;
}
