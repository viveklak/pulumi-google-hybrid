// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets metadata for a given Secret.
 */
export function getSecret(args: GetSecretArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:secretmanager/v1:getSecret", {
        "project": args.project,
        "secretId": args.secretId,
    }, opts);
}

export interface GetSecretArgs {
    project: string;
    secretId: string;
}

export interface GetSecretResult {
    /**
     * The time at which the Secret was created.
     */
    readonly createTime: string;
    /**
     * Optional. Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
     */
    readonly expireTime: string;
    /**
     * The labels assigned to this Secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `\p{Ll}\p{Lo}{0,62}` Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}` No more than 64 labels can be assigned to a given resource.
     */
    readonly labels: {[key: string]: string};
    /**
     * The resource name of the Secret in the format `projects/*&#47;secrets/*`.
     */
    readonly name: string;
    /**
     * Immutable. The replication policy of the secret data attached to the Secret. The replication policy cannot be changed after the Secret has been created.
     */
    readonly replication: outputs.secretmanager.v1.ReplicationResponse;
    /**
     * Optional. Rotation policy attached to the Secret. May be excluded if there is no rotation policy.
     */
    readonly rotation: outputs.secretmanager.v1.RotationResponse;
    /**
     * Optional. A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
     */
    readonly topics: outputs.secretmanager.v1.TopicResponse[];
    /**
     * Input only. The TTL for the Secret.
     */
    readonly ttl: string;
}
