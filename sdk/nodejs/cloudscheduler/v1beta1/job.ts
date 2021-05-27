// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a job.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudscheduler/v1beta1:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * App Engine HTTP target.
     */
    public readonly appEngineHttpTarget!: pulumi.Output<outputs.cloudscheduler.v1beta1.AppEngineHttpTargetResponse>;
    /**
     * The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * For PubSub targets, this field is ignored.
     */
    public readonly attemptDeadline!: pulumi.Output<string>;
    /**
     * Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * HTTP target.
     */
    public readonly httpTarget!: pulumi.Output<outputs.cloudscheduler.v1beta1.HttpTargetResponse>;
    /**
     * The time the last job attempt started.
     */
    public readonly lastAttemptTime!: pulumi.Output<string>;
    /**
     * Immutable. This field is used to manage the legacy App Engine Cron jobs using the Cloud Scheduler API. If the field is set to true, the job will be considered a legacy job. Note that App Engine Cron jobs have fewer features than Cloud Scheduler jobs, e.g., are only limited to App Engine targets.
     */
    public readonly legacyAppEngineCron!: pulumi.Output<boolean>;
    /**
     * Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Pub/Sub target.
     */
    public readonly pubsubTarget!: pulumi.Output<outputs.cloudscheduler.v1beta1.PubsubTargetResponse>;
    /**
     * Settings that determine the retry behavior.
     */
    public readonly retryConfig!: pulumi.Output<outputs.cloudscheduler.v1beta1.RetryConfigResponse>;
    /**
     * Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
     */
    public readonly scheduleTime!: pulumi.Output<string>;
    /**
     * State of the job.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The response from the target for the last attempted execution.
     */
    public readonly status!: pulumi.Output<outputs.cloudscheduler.v1beta1.StatusResponse>;
    /**
     * Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
     */
    public readonly timeZone!: pulumi.Output<string>;
    /**
     * The creation time of the job.
     */
    public readonly userUpdateTime!: pulumi.Output<string>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["appEngineHttpTarget"] = args ? args.appEngineHttpTarget : undefined;
            inputs["attemptDeadline"] = args ? args.attemptDeadline : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["httpTarget"] = args ? args.httpTarget : undefined;
            inputs["lastAttemptTime"] = args ? args.lastAttemptTime : undefined;
            inputs["legacyAppEngineCron"] = args ? args.legacyAppEngineCron : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pubsubTarget"] = args ? args.pubsubTarget : undefined;
            inputs["retryConfig"] = args ? args.retryConfig : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["scheduleTime"] = args ? args.scheduleTime : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["timeZone"] = args ? args.timeZone : undefined;
            inputs["userUpdateTime"] = args ? args.userUpdateTime : undefined;
        } else {
            inputs["appEngineHttpTarget"] = undefined /*out*/;
            inputs["attemptDeadline"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["httpTarget"] = undefined /*out*/;
            inputs["lastAttemptTime"] = undefined /*out*/;
            inputs["legacyAppEngineCron"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["pubsubTarget"] = undefined /*out*/;
            inputs["retryConfig"] = undefined /*out*/;
            inputs["schedule"] = undefined /*out*/;
            inputs["scheduleTime"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["timeZone"] = undefined /*out*/;
            inputs["userUpdateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * App Engine HTTP target.
     */
    readonly appEngineHttpTarget?: pulumi.Input<inputs.cloudscheduler.v1beta1.AppEngineHttpTargetArgs>;
    /**
     * The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * For PubSub targets, this field is ignored.
     */
    readonly attemptDeadline?: pulumi.Input<string>;
    /**
     * Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * HTTP target.
     */
    readonly httpTarget?: pulumi.Input<inputs.cloudscheduler.v1beta1.HttpTargetArgs>;
    /**
     * The time the last job attempt started.
     */
    readonly lastAttemptTime?: pulumi.Input<string>;
    /**
     * Immutable. This field is used to manage the legacy App Engine Cron jobs using the Cloud Scheduler API. If the field is set to true, the job will be considered a legacy job. Note that App Engine Cron jobs have fewer features than Cloud Scheduler jobs, e.g., are only limited to App Engine targets.
     */
    readonly legacyAppEngineCron?: pulumi.Input<boolean>;
    readonly location: pulumi.Input<string>;
    /**
     * Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * Pub/Sub target.
     */
    readonly pubsubTarget?: pulumi.Input<inputs.cloudscheduler.v1beta1.PubsubTargetArgs>;
    /**
     * Settings that determine the retry behavior.
     */
    readonly retryConfig?: pulumi.Input<inputs.cloudscheduler.v1beta1.RetryConfigArgs>;
    /**
     * Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
     */
    readonly scheduleTime?: pulumi.Input<string>;
    /**
     * State of the job.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The response from the target for the last attempted execution.
     */
    readonly status?: pulumi.Input<inputs.cloudscheduler.v1beta1.StatusArgs>;
    /**
     * Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
     */
    readonly timeZone?: pulumi.Input<string>;
    /**
     * The creation time of the job.
     */
    readonly userUpdateTime?: pulumi.Input<string>;
}
