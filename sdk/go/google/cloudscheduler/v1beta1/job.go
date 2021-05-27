// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a job.
type Job struct {
	pulumi.CustomResourceState

	// App Engine HTTP target.
	AppEngineHttpTarget AppEngineHttpTargetResponseOutput `pulumi:"appEngineHttpTarget"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * For PubSub targets, this field is ignored.
	AttemptDeadline pulumi.StringOutput `pulumi:"attemptDeadline"`
	// Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// HTTP target.
	HttpTarget HttpTargetResponseOutput `pulumi:"httpTarget"`
	// The time the last job attempt started.
	LastAttemptTime pulumi.StringOutput `pulumi:"lastAttemptTime"`
	// Immutable. This field is used to manage the legacy App Engine Cron jobs using the Cloud Scheduler API. If the field is set to true, the job will be considered a legacy job. Note that App Engine Cron jobs have fewer features than Cloud Scheduler jobs, e.g., are only limited to App Engine targets.
	LegacyAppEngineCron pulumi.BoolOutput `pulumi:"legacyAppEngineCron"`
	// Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Pub/Sub target.
	PubsubTarget PubsubTargetResponseOutput `pulumi:"pubsubTarget"`
	// Settings that determine the retry behavior.
	RetryConfig RetryConfigResponseOutput `pulumi:"retryConfig"`
	// Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
	ScheduleTime pulumi.StringOutput `pulumi:"scheduleTime"`
	// State of the job.
	State pulumi.StringOutput `pulumi:"state"`
	// The response from the target for the last attempted execution.
	Status StatusResponseOutput `pulumi:"status"`
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
	// The creation time of the job.
	UserUpdateTime pulumi.StringOutput `pulumi:"userUpdateTime"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Job
	err := ctx.RegisterResource("google-native:cloudscheduler/v1beta1:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("google-native:cloudscheduler/v1beta1:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// App Engine HTTP target.
	AppEngineHttpTarget *AppEngineHttpTargetResponse `pulumi:"appEngineHttpTarget"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * For PubSub targets, this field is ignored.
	AttemptDeadline *string `pulumi:"attemptDeadline"`
	// Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
	Description *string `pulumi:"description"`
	// HTTP target.
	HttpTarget *HttpTargetResponse `pulumi:"httpTarget"`
	// The time the last job attempt started.
	LastAttemptTime *string `pulumi:"lastAttemptTime"`
	// Immutable. This field is used to manage the legacy App Engine Cron jobs using the Cloud Scheduler API. If the field is set to true, the job will be considered a legacy job. Note that App Engine Cron jobs have fewer features than Cloud Scheduler jobs, e.g., are only limited to App Engine targets.
	LegacyAppEngineCron *bool `pulumi:"legacyAppEngineCron"`
	// Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name *string `pulumi:"name"`
	// Pub/Sub target.
	PubsubTarget *PubsubTargetResponse `pulumi:"pubsubTarget"`
	// Settings that determine the retry behavior.
	RetryConfig *RetryConfigResponse `pulumi:"retryConfig"`
	// Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
	Schedule *string `pulumi:"schedule"`
	// The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
	ScheduleTime *string `pulumi:"scheduleTime"`
	// State of the job.
	State *string `pulumi:"state"`
	// The response from the target for the last attempted execution.
	Status *StatusResponse `pulumi:"status"`
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
	TimeZone *string `pulumi:"timeZone"`
	// The creation time of the job.
	UserUpdateTime *string `pulumi:"userUpdateTime"`
}

type JobState struct {
	// App Engine HTTP target.
	AppEngineHttpTarget AppEngineHttpTargetResponsePtrInput
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * For PubSub targets, this field is ignored.
	AttemptDeadline pulumi.StringPtrInput
	// Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
	Description pulumi.StringPtrInput
	// HTTP target.
	HttpTarget HttpTargetResponsePtrInput
	// The time the last job attempt started.
	LastAttemptTime pulumi.StringPtrInput
	// Immutable. This field is used to manage the legacy App Engine Cron jobs using the Cloud Scheduler API. If the field is set to true, the job will be considered a legacy job. Note that App Engine Cron jobs have fewer features than Cloud Scheduler jobs, e.g., are only limited to App Engine targets.
	LegacyAppEngineCron pulumi.BoolPtrInput
	// Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name pulumi.StringPtrInput
	// Pub/Sub target.
	PubsubTarget PubsubTargetResponsePtrInput
	// Settings that determine the retry behavior.
	RetryConfig RetryConfigResponsePtrInput
	// Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
	Schedule pulumi.StringPtrInput
	// The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
	ScheduleTime pulumi.StringPtrInput
	// State of the job.
	State pulumi.StringPtrInput
	// The response from the target for the last attempted execution.
	Status StatusResponsePtrInput
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
	TimeZone pulumi.StringPtrInput
	// The creation time of the job.
	UserUpdateTime pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// App Engine HTTP target.
	AppEngineHttpTarget *AppEngineHttpTarget `pulumi:"appEngineHttpTarget"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * For PubSub targets, this field is ignored.
	AttemptDeadline *string `pulumi:"attemptDeadline"`
	// Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
	Description *string `pulumi:"description"`
	// HTTP target.
	HttpTarget *HttpTarget `pulumi:"httpTarget"`
	// The time the last job attempt started.
	LastAttemptTime *string `pulumi:"lastAttemptTime"`
	// Immutable. This field is used to manage the legacy App Engine Cron jobs using the Cloud Scheduler API. If the field is set to true, the job will be considered a legacy job. Note that App Engine Cron jobs have fewer features than Cloud Scheduler jobs, e.g., are only limited to App Engine targets.
	LegacyAppEngineCron *bool  `pulumi:"legacyAppEngineCron"`
	Location            string `pulumi:"location"`
	// Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// Pub/Sub target.
	PubsubTarget *PubsubTarget `pulumi:"pubsubTarget"`
	// Settings that determine the retry behavior.
	RetryConfig *RetryConfig `pulumi:"retryConfig"`
	// Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
	Schedule *string `pulumi:"schedule"`
	// The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
	ScheduleTime *string `pulumi:"scheduleTime"`
	// State of the job.
	State *string `pulumi:"state"`
	// The response from the target for the last attempted execution.
	Status *Status `pulumi:"status"`
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
	TimeZone *string `pulumi:"timeZone"`
	// The creation time of the job.
	UserUpdateTime *string `pulumi:"userUpdateTime"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// App Engine HTTP target.
	AppEngineHttpTarget AppEngineHttpTargetPtrInput
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * For PubSub targets, this field is ignored.
	AttemptDeadline pulumi.StringPtrInput
	// Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
	Description pulumi.StringPtrInput
	// HTTP target.
	HttpTarget HttpTargetPtrInput
	// The time the last job attempt started.
	LastAttemptTime pulumi.StringPtrInput
	// Immutable. This field is used to manage the legacy App Engine Cron jobs using the Cloud Scheduler API. If the field is set to true, the job will be considered a legacy job. Note that App Engine Cron jobs have fewer features than Cloud Scheduler jobs, e.g., are only limited to App Engine targets.
	LegacyAppEngineCron pulumi.BoolPtrInput
	Location            pulumi.StringInput
	// Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// Pub/Sub target.
	PubsubTarget PubsubTargetPtrInput
	// Settings that determine the retry behavior.
	RetryConfig RetryConfigPtrInput
	// Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
	Schedule pulumi.StringPtrInput
	// The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
	ScheduleTime pulumi.StringPtrInput
	// State of the job.
	State pulumi.StringPtrInput
	// The response from the target for the last attempted execution.
	Status StatusPtrInput
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
	TimeZone pulumi.StringPtrInput
	// The creation time of the job.
	UserUpdateTime pulumi.StringPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((*Job)(nil))
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct {
	*pulumi.OutputState
}

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Job)(nil))
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
