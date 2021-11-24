// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new data transfer configuration.
 */
export class TransferConfig extends pulumi.CustomResource {
    /**
     * Get an existing TransferConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TransferConfig {
        return new TransferConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:bigquerydatatransfer/v1:TransferConfig';

    /**
     * Returns true if the given object is an instance of TransferConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransferConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransferConfig.__pulumiType;
    }

    /**
     * The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
     */
    public readonly dataRefreshWindowDays!: pulumi.Output<number>;
    /**
     * Data source id. Cannot be changed once data transfer is created.
     */
    public readonly dataSourceId!: pulumi.Output<string>;
    /**
     * Region in which BigQuery dataset is located.
     */
    public /*out*/ readonly datasetRegion!: pulumi.Output<string>;
    /**
     * The BigQuery target dataset id.
     */
    public readonly destinationDatasetId!: pulumi.Output<string>;
    /**
     * Is this config disabled. When set to true, no runs are scheduled for a given transfer.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * User specified display name for the data transfer.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
     */
    public readonly emailPreferences!: pulumi.Output<outputs.bigquerydatatransfer.v1.EmailPreferencesResponse>;
    /**
     * The resource name of the transfer config. Transfer config names have the form `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`. Where `config_id` is usually a uuid, even though it is not guaranteed or required. The name is ignored when creating a transfer config.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Next time when data transfer will run.
     */
    public /*out*/ readonly nextRunTime!: pulumi.Output<string>;
    /**
     * Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish. The format for specifying a pubsub topic is: `projects/{project}/topics/{topic}`
     */
    public readonly notificationPubsubTopic!: pulumi.Output<string>;
    /**
     * Information about the user whose credentials are used to transfer data. Populated only for `transferConfigs.get` requests. In case the user information is not available, this field will not be populated.
     */
    public /*out*/ readonly ownerInfo!: pulumi.Output<outputs.bigquerydatatransfer.v1.UserInfoResponse>;
    /**
     * Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
     */
    public readonly params!: pulumi.Output<{[key: string]: string}>;
    /**
     * Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the default value for the data source will be used. The specified times are in UTC. Examples of valid format: `1st,3rd monday of month 15:30`, `every wed,fri of jan,jun 13:15`, and `first sunday of quarter 00:00`. See more explanation about the format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format NOTE: the granularity should be at least 8 hours, or less frequent.
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * Options customizing the data transfer schedule.
     */
    public readonly scheduleOptions!: pulumi.Output<outputs.bigquerydatatransfer.v1.ScheduleOptionsResponse>;
    /**
     * State of the most recently updated transfer run.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Data transfer modification time. Ignored by server on input.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a TransferConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TransferConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["authorizationCode"] = args ? args.authorizationCode : undefined;
            resourceInputs["dataRefreshWindowDays"] = args ? args.dataRefreshWindowDays : undefined;
            resourceInputs["dataSourceId"] = args ? args.dataSourceId : undefined;
            resourceInputs["destinationDatasetId"] = args ? args.destinationDatasetId : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["emailPreferences"] = args ? args.emailPreferences : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationPubsubTopic"] = args ? args.notificationPubsubTopic : undefined;
            resourceInputs["params"] = args ? args.params : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["scheduleOptions"] = args ? args.scheduleOptions : undefined;
            resourceInputs["serviceAccountName"] = args ? args.serviceAccountName : undefined;
            resourceInputs["versionInfo"] = args ? args.versionInfo : undefined;
            resourceInputs["datasetRegion"] = undefined /*out*/;
            resourceInputs["nextRunTime"] = undefined /*out*/;
            resourceInputs["ownerInfo"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["dataRefreshWindowDays"] = undefined /*out*/;
            resourceInputs["dataSourceId"] = undefined /*out*/;
            resourceInputs["datasetRegion"] = undefined /*out*/;
            resourceInputs["destinationDatasetId"] = undefined /*out*/;
            resourceInputs["disabled"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["emailPreferences"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nextRunTime"] = undefined /*out*/;
            resourceInputs["notificationPubsubTopic"] = undefined /*out*/;
            resourceInputs["ownerInfo"] = undefined /*out*/;
            resourceInputs["params"] = undefined /*out*/;
            resourceInputs["schedule"] = undefined /*out*/;
            resourceInputs["scheduleOptions"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TransferConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TransferConfig resource.
 */
export interface TransferConfigArgs {
    authorizationCode?: pulumi.Input<string>;
    /**
     * The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
     */
    dataRefreshWindowDays?: pulumi.Input<number>;
    /**
     * Data source id. Cannot be changed once data transfer is created.
     */
    dataSourceId?: pulumi.Input<string>;
    /**
     * The BigQuery target dataset id.
     */
    destinationDatasetId?: pulumi.Input<string>;
    /**
     * Is this config disabled. When set to true, no runs are scheduled for a given transfer.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * User specified display name for the data transfer.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
     */
    emailPreferences?: pulumi.Input<inputs.bigquerydatatransfer.v1.EmailPreferencesArgs>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the transfer config. Transfer config names have the form `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`. Where `config_id` is usually a uuid, even though it is not guaranteed or required. The name is ignored when creating a transfer config.
     */
    name?: pulumi.Input<string>;
    /**
     * Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish. The format for specifying a pubsub topic is: `projects/{project}/topics/{topic}`
     */
    notificationPubsubTopic?: pulumi.Input<string>;
    /**
     * Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
     */
    params?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    project?: pulumi.Input<string>;
    /**
     * Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the default value for the data source will be used. The specified times are in UTC. Examples of valid format: `1st,3rd monday of month 15:30`, `every wed,fri of jan,jun 13:15`, and `first sunday of quarter 00:00`. See more explanation about the format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format NOTE: the granularity should be at least 8 hours, or less frequent.
     */
    schedule?: pulumi.Input<string>;
    /**
     * Options customizing the data transfer schedule.
     */
    scheduleOptions?: pulumi.Input<inputs.bigquerydatatransfer.v1.ScheduleOptionsArgs>;
    serviceAccountName?: pulumi.Input<string>;
    versionInfo?: pulumi.Input<string>;
}
