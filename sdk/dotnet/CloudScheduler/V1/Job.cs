// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudScheduler.V1
{
    /// <summary>
    /// Creates a job.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudscheduler/v1:Job")]
    public partial class Job : Pulumi.CustomResource
    {
        /// <summary>
        /// App Engine HTTP target.
        /// </summary>
        [Output("appEngineHttpTarget")]
        public Output<Outputs.AppEngineHttpTargetResponse> AppEngineHttpTarget { get; private set; } = null!;

        /// <summary>
        /// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours 15 seconds.
        /// </summary>
        [Output("attemptDeadline")]
        public Output<string> AttemptDeadline { get; private set; } = null!;

        /// <summary>
        /// Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// HTTP target.
        /// </summary>
        [Output("httpTarget")]
        public Output<Outputs.HttpTargetResponse> HttpTarget { get; private set; } = null!;

        /// <summary>
        /// The time the last job attempt started.
        /// </summary>
        [Output("lastAttemptTime")]
        public Output<string> LastAttemptTime { get; private set; } = null!;

        /// <summary>
        /// Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Pub/Sub target.
        /// </summary>
        [Output("pubsubTarget")]
        public Output<Outputs.PubsubTargetResponse> PubsubTarget { get; private set; } = null!;

        /// <summary>
        /// Settings that determine the retry behavior.
        /// </summary>
        [Output("retryConfig")]
        public Output<Outputs.RetryConfigResponse> RetryConfig { get; private set; } = null!;

        /// <summary>
        /// Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](https://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count &gt; 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
        /// </summary>
        [Output("schedule")]
        public Output<string> Schedule { get; private set; } = null!;

        /// <summary>
        /// The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
        /// </summary>
        [Output("scheduleTime")]
        public Output<string> ScheduleTime { get; private set; } = null!;

        /// <summary>
        /// State of the job.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The response from the target for the last attempted execution.
        /// </summary>
        [Output("status")]
        public Output<Outputs.StatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
        /// </summary>
        [Output("timeZone")]
        public Output<string> TimeZone { get; private set; } = null!;

        /// <summary>
        /// The creation time of the job.
        /// </summary>
        [Output("userUpdateTime")]
        public Output<string> UserUpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:cloudscheduler/v1:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudscheduler/v1:Job", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Job(name, id, options);
        }
    }

    public sealed class JobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// App Engine HTTP target.
        /// </summary>
        [Input("appEngineHttpTarget")]
        public Input<Inputs.AppEngineHttpTargetArgs>? AppEngineHttpTarget { get; set; }

        /// <summary>
        /// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours 15 seconds.
        /// </summary>
        [Input("attemptDeadline")]
        public Input<string>? AttemptDeadline { get; set; }

        /// <summary>
        /// Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// HTTP target.
        /// </summary>
        [Input("httpTarget")]
        public Input<Inputs.HttpTargetArgs>? HttpTarget { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Pub/Sub target.
        /// </summary>
        [Input("pubsubTarget")]
        public Input<Inputs.PubsubTargetArgs>? PubsubTarget { get; set; }

        /// <summary>
        /// Settings that determine the retry behavior.
        /// </summary>
        [Input("retryConfig")]
        public Input<Inputs.RetryConfigArgs>? RetryConfig { get; set; }

        /// <summary>
        /// Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](https://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count &gt; 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public JobArgs()
        {
        }
    }
}
