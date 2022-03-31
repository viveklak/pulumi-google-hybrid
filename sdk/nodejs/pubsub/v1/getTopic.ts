// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets the configuration of a topic.
 */
export function getTopic(args: GetTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:pubsub/v1:getTopic", {
        "project": args.project,
        "topicId": args.topicId,
    }, opts);
}

export interface GetTopicArgs {
    project?: string;
    topicId: string;
}

export interface GetTopicResult {
    /**
     * The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this topic. The expected format is `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*`.
     */
    readonly kmsKeyName: string;
    /**
     * See [Creating and managing labels] (https://cloud.google.com/pubsub/docs/labels).
     */
    readonly labels: {[key: string]: string};
    /**
     * Indicates the minimum duration to retain a message after it is published to the topic. If this field is set, messages published to the topic in the last `message_retention_duration` are always available to subscribers. For instance, it allows any attached subscription to [seek to a timestamp](https://cloud.google.com/pubsub/docs/replay-overview#seek_to_a_time) that is up to `message_retention_duration` in the past. If this field is not set, message retention is controlled by settings on individual subscriptions. Cannot be more than 31 days or less than 10 minutes.
     */
    readonly messageRetentionDuration: string;
    /**
     * Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be stored. If not present, then no constraints are in effect.
     */
    readonly messageStoragePolicy: outputs.pubsub.v1.MessageStoragePolicyResponse;
    /**
     * The name of the topic. It must have the format `"projects/{project}/topics/{topic}"`. `{topic}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
     */
    readonly name: string;
    /**
     * Reserved for future use. This field is set only in responses from the server; it is ignored if it is set in any requests.
     */
    readonly satisfiesPzs: boolean;
    /**
     * Settings for validating messages published against a schema.
     */
    readonly schemaSettings: outputs.pubsub.v1.SchemaSettingsResponse;
}

export function getTopicOutput(args: GetTopicOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicResult> {
    return pulumi.output(args).apply(a => getTopic(a, opts))
}

export interface GetTopicOutputArgs {
    project?: pulumi.Input<string>;
    topicId: pulumi.Input<string>;
}
