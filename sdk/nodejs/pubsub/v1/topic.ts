// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates the given topic with the given name. See the [resource name rules] (https://cloud.google.com/pubsub/docs/admin#resource_names).
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:pubsub/v1:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this topic. The expected format is `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*`.
     */
    public readonly kmsKeyName!: pulumi.Output<string>;
    /**
     * See [Creating and managing labels] (https://cloud.google.com/pubsub/docs/labels).
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Indicates the minimum duration to retain a message after it is published to the topic. If this field is set, messages published to the topic in the last `message_retention_duration` are always available to subscribers. For instance, it allows any attached subscription to [seek to a timestamp](https://cloud.google.com/pubsub/docs/replay-overview#seek_to_a_time) that is up to `message_retention_duration` in the past. If this field is not set, message retention is controlled by settings on individual subscriptions. Cannot be more than 31 days or less than 10 minutes.
     */
    public readonly messageRetentionDuration!: pulumi.Output<string>;
    /**
     * Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be stored. If not present, then no constraints are in effect.
     */
    public readonly messageStoragePolicy!: pulumi.Output<outputs.pubsub.v1.MessageStoragePolicyResponse>;
    /**
     * The name of the topic. It must have the format `"projects/{project}/topics/{topic}"`. `{topic}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Reserved for future use. This field is set only in responses from the server; it is ignored if it is set in any requests.
     */
    public readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * Settings for validating messages published against a schema.
     */
    public readonly schemaSettings!: pulumi.Output<outputs.pubsub.v1.SchemaSettingsResponse>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.topicId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicId'");
            }
            resourceInputs["kmsKeyName"] = args ? args.kmsKeyName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["messageRetentionDuration"] = args ? args.messageRetentionDuration : undefined;
            resourceInputs["messageStoragePolicy"] = args ? args.messageStoragePolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["satisfiesPzs"] = args ? args.satisfiesPzs : undefined;
            resourceInputs["schemaSettings"] = args ? args.schemaSettings : undefined;
            resourceInputs["topicId"] = args ? args.topicId : undefined;
        } else {
            resourceInputs["kmsKeyName"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["messageRetentionDuration"] = undefined /*out*/;
            resourceInputs["messageStoragePolicy"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["schemaSettings"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Topic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this topic. The expected format is `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*`.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * See [Creating and managing labels] (https://cloud.google.com/pubsub/docs/labels).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the minimum duration to retain a message after it is published to the topic. If this field is set, messages published to the topic in the last `message_retention_duration` are always available to subscribers. For instance, it allows any attached subscription to [seek to a timestamp](https://cloud.google.com/pubsub/docs/replay-overview#seek_to_a_time) that is up to `message_retention_duration` in the past. If this field is not set, message retention is controlled by settings on individual subscriptions. Cannot be more than 31 days or less than 10 minutes.
     */
    messageRetentionDuration?: pulumi.Input<string>;
    /**
     * Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be stored. If not present, then no constraints are in effect.
     */
    messageStoragePolicy?: pulumi.Input<inputs.pubsub.v1.MessageStoragePolicyArgs>;
    /**
     * The name of the topic. It must have the format `"projects/{project}/topics/{topic}"`. `{topic}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Reserved for future use. This field is set only in responses from the server; it is ignored if it is set in any requests.
     */
    satisfiesPzs?: pulumi.Input<boolean>;
    /**
     * Settings for validating messages published against a schema.
     */
    schemaSettings?: pulumi.Input<inputs.pubsub.v1.SchemaSettingsArgs>;
    topicId: pulumi.Input<string>;
}
