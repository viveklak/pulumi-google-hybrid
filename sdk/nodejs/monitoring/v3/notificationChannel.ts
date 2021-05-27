// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new notification channel, representing a single notification endpoint such as an email address, SMS number, or PagerDuty service.
 */
export class NotificationChannel extends pulumi.CustomResource {
    /**
     * Get an existing NotificationChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NotificationChannel {
        return new NotificationChannel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:monitoring/v3:NotificationChannel';

    /**
     * Returns true if the given object is an instance of NotificationChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationChannel.__pulumiType;
    }

    /**
     * Record of the creation of this channel.
     */
    public readonly creationRecord!: pulumi.Output<outputs.monitoring.v3.MutationRecordResponse>;
    /**
     * An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the NotificationChannelDescriptor.labels of the NotificationChannelDescriptor corresponding to the type field.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Records of the modification of this channel.
     */
    public readonly mutationRecords!: pulumi.Output<outputs.monitoring.v3.MutationRecordResponse[]>;
    /**
     * The full REST resource name for this channel. The format is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] The [CHANNEL_ID] is automatically assigned by the server on creation.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
     */
    public readonly userLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
     */
    public readonly verificationStatus!: pulumi.Output<string>;

    /**
     * Create a NotificationChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationChannelArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["creationRecord"] = args ? args.creationRecord : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["mutationRecords"] = args ? args.mutationRecords : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["userLabels"] = args ? args.userLabels : undefined;
            inputs["verificationStatus"] = args ? args.verificationStatus : undefined;
        } else {
            inputs["creationRecord"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["enabled"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["mutationRecords"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["userLabels"] = undefined /*out*/;
            inputs["verificationStatus"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NotificationChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NotificationChannel resource.
 */
export interface NotificationChannelArgs {
    /**
     * Record of the creation of this channel.
     */
    readonly creationRecord?: pulumi.Input<inputs.monitoring.v3.MutationRecordArgs>;
    /**
     * An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the NotificationChannelDescriptor.labels of the NotificationChannelDescriptor corresponding to the type field.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Records of the modification of this channel.
     */
    readonly mutationRecords?: pulumi.Input<pulumi.Input<inputs.monitoring.v3.MutationRecordArgs>[]>;
    /**
     * The full REST resource name for this channel. The format is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] The [CHANNEL_ID] is automatically assigned by the server on creation.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
     */
    readonly userLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
     */
    readonly verificationStatus?: pulumi.Input<string>;
}
