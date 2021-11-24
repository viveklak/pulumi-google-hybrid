// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a notification subscription for a given bucket.
 * Auto-naming is currently not supported for this resource.
 */
export class Notification extends pulumi.CustomResource {
    /**
     * Get an existing Notification resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Notification {
        return new Notification(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:storage/v1:Notification';

    /**
     * Returns true if the given object is an instance of Notification.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Notification {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Notification.__pulumiType;
    }

    /**
     * An optional list of additional attributes to attach to each Cloud PubSub message published for this notification subscription.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string}>;
    /**
     * HTTP 1.1 Entity tag for this subscription notification.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * If present, only send notifications about listed event types. If empty, sent notifications for all event types.
     */
    public readonly eventTypes!: pulumi.Output<string[]>;
    /**
     * The kind of item this is. For notifications, this is always storage#notification.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * If present, only apply this notification configuration to object names that begin with this prefix.
     */
    public readonly objectNamePrefix!: pulumi.Output<string>;
    /**
     * The desired content of the Payload.
     */
    public readonly payloadFormat!: pulumi.Output<string>;
    /**
     * The canonical URL of this notification.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * The Cloud PubSub topic to which this subscription publishes. Formatted as: '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
     */
    public readonly topic!: pulumi.Output<string>;

    /**
     * Create a Notification resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["customAttributes"] = args ? args.customAttributes : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["eventTypes"] = args ? args.eventTypes : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["objectNamePrefix"] = args ? args.objectNamePrefix : undefined;
            resourceInputs["payloadFormat"] = args ? args.payloadFormat : undefined;
            resourceInputs["provisionalUserProject"] = args ? args.provisionalUserProject : undefined;
            resourceInputs["selfLink"] = args ? args.selfLink : undefined;
            resourceInputs["topic"] = args ? args.topic : undefined;
            resourceInputs["userProject"] = args ? args.userProject : undefined;
        } else {
            resourceInputs["customAttributes"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["eventTypes"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["objectNamePrefix"] = undefined /*out*/;
            resourceInputs["payloadFormat"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["topic"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Notification.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Notification resource.
 */
export interface NotificationArgs {
    bucket: pulumi.Input<string>;
    /**
     * An optional list of additional attributes to attach to each Cloud PubSub message published for this notification subscription.
     */
    customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * HTTP 1.1 Entity tag for this subscription notification.
     */
    etag?: pulumi.Input<string>;
    /**
     * If present, only send notifications about listed event types. If empty, sent notifications for all event types.
     */
    eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the notification.
     */
    id?: pulumi.Input<string>;
    /**
     * The kind of item this is. For notifications, this is always storage#notification.
     */
    kind?: pulumi.Input<string>;
    /**
     * If present, only apply this notification configuration to object names that begin with this prefix.
     */
    objectNamePrefix?: pulumi.Input<string>;
    /**
     * The desired content of the Payload.
     */
    payloadFormat?: pulumi.Input<string>;
    provisionalUserProject?: pulumi.Input<string>;
    /**
     * The canonical URL of this notification.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The Cloud PubSub topic to which this subscription publishes. Formatted as: '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
     */
    topic?: pulumi.Input<string>;
    userProject?: pulumi.Input<string>;
}
