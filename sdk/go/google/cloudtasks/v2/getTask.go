// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a task.
func LookupTask(ctx *pulumi.Context, args *LookupTaskArgs, opts ...pulumi.InvokeOption) (*LookupTaskResult, error) {
	var rv LookupTaskResult
	err := ctx.Invoke("google-native:cloudtasks/v2:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTaskArgs struct {
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
	QueueId      string  `pulumi:"queueId"`
	ResponseView *string `pulumi:"responseView"`
	TaskId       string  `pulumi:"taskId"`
}

type LookupTaskResult struct {
	// HTTP request that is sent to the App Engine app handler. An App Engine task is a task that has AppEngineHttpRequest set.
	AppEngineHttpRequest AppEngineHttpRequestResponse `pulumi:"appEngineHttpRequest"`
	// The time that the task was created. `create_time` will be truncated to the nearest second.
	CreateTime string `pulumi:"createTime"`
	// The number of attempts dispatched. This count includes attempts which have been dispatched but haven't received a response.
	DispatchCount int `pulumi:"dispatchCount"`
	// The deadline for requests sent to the worker. If the worker does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the task according to the RetryConfig. Note that when the request is cancelled, Cloud Tasks will stop listening for the response, but whether the worker stops processing depends on the worker. For example, if the worker is stuck, it may not react to cancelled requests. The default and maximum values depend on the type of request: * For HTTP tasks, the default is 10 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine tasks, 0 indicates that the request has the default deadline. The default deadline depends on the [scaling type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling) of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. Regardless of the task's `dispatch_deadline`, the app handler will not run for longer than than the service's timeout. We recommend setting the `dispatch_deadline` to at most a few seconds more than the app handler's timeout. For more information see [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts). `dispatch_deadline` will be truncated to the nearest millisecond. The deadline is an approximate deadline.
	DispatchDeadline string `pulumi:"dispatchDeadline"`
	// The status of the task's first attempt. Only dispatch_time will be set. The other Attempt information is not retained by Cloud Tasks.
	FirstAttempt AttemptResponse `pulumi:"firstAttempt"`
	// HTTP request that is sent to the worker. An HTTP task is a task that has HttpRequest set.
	HttpRequest HttpRequestResponse `pulumi:"httpRequest"`
	// The status of the task's last attempt.
	LastAttempt AttemptResponse `pulumi:"lastAttempt"`
	// Optionally caller-specified in CreateTask. The task name. The task name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the task's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters. * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name string `pulumi:"name"`
	// The number of attempts which have received a response.
	ResponseCount int `pulumi:"responseCount"`
	// The time when the task is scheduled to be attempted or retried. `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime string `pulumi:"scheduleTime"`
	// The view specifies which subset of the Task has been returned.
	View string `pulumi:"view"`
}

func LookupTaskOutput(ctx *pulumi.Context, args LookupTaskOutputArgs, opts ...pulumi.InvokeOption) LookupTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTaskResult, error) {
			args := v.(LookupTaskArgs)
			r, err := LookupTask(ctx, &args, opts...)
			return *r, err
		}).(LookupTaskResultOutput)
}

type LookupTaskOutputArgs struct {
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
	QueueId      pulumi.StringInput    `pulumi:"queueId"`
	ResponseView pulumi.StringPtrInput `pulumi:"responseView"`
	TaskId       pulumi.StringInput    `pulumi:"taskId"`
}

func (LookupTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskArgs)(nil)).Elem()
}

type LookupTaskResultOutput struct{ *pulumi.OutputState }

func (LookupTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskResult)(nil)).Elem()
}

func (o LookupTaskResultOutput) ToLookupTaskResultOutput() LookupTaskResultOutput {
	return o
}

func (o LookupTaskResultOutput) ToLookupTaskResultOutputWithContext(ctx context.Context) LookupTaskResultOutput {
	return o
}

// HTTP request that is sent to the App Engine app handler. An App Engine task is a task that has AppEngineHttpRequest set.
func (o LookupTaskResultOutput) AppEngineHttpRequest() AppEngineHttpRequestResponseOutput {
	return o.ApplyT(func(v LookupTaskResult) AppEngineHttpRequestResponse { return v.AppEngineHttpRequest }).(AppEngineHttpRequestResponseOutput)
}

// The time that the task was created. `create_time` will be truncated to the nearest second.
func (o LookupTaskResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The number of attempts dispatched. This count includes attempts which have been dispatched but haven't received a response.
func (o LookupTaskResultOutput) DispatchCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTaskResult) int { return v.DispatchCount }).(pulumi.IntOutput)
}

// The deadline for requests sent to the worker. If the worker does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the task according to the RetryConfig. Note that when the request is cancelled, Cloud Tasks will stop listening for the response, but whether the worker stops processing depends on the worker. For example, if the worker is stuck, it may not react to cancelled requests. The default and maximum values depend on the type of request: * For HTTP tasks, the default is 10 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine tasks, 0 indicates that the request has the default deadline. The default deadline depends on the [scaling type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling) of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. Regardless of the task's `dispatch_deadline`, the app handler will not run for longer than than the service's timeout. We recommend setting the `dispatch_deadline` to at most a few seconds more than the app handler's timeout. For more information see [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts). `dispatch_deadline` will be truncated to the nearest millisecond. The deadline is an approximate deadline.
func (o LookupTaskResultOutput) DispatchDeadline() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.DispatchDeadline }).(pulumi.StringOutput)
}

// The status of the task's first attempt. Only dispatch_time will be set. The other Attempt information is not retained by Cloud Tasks.
func (o LookupTaskResultOutput) FirstAttempt() AttemptResponseOutput {
	return o.ApplyT(func(v LookupTaskResult) AttemptResponse { return v.FirstAttempt }).(AttemptResponseOutput)
}

// HTTP request that is sent to the worker. An HTTP task is a task that has HttpRequest set.
func (o LookupTaskResultOutput) HttpRequest() HttpRequestResponseOutput {
	return o.ApplyT(func(v LookupTaskResult) HttpRequestResponse { return v.HttpRequest }).(HttpRequestResponseOutput)
}

// The status of the task's last attempt.
func (o LookupTaskResultOutput) LastAttempt() AttemptResponseOutput {
	return o.ApplyT(func(v LookupTaskResult) AttemptResponse { return v.LastAttempt }).(AttemptResponseOutput)
}

// Optionally caller-specified in CreateTask. The task name. The task name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the task's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters. * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
func (o LookupTaskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of attempts which have received a response.
func (o LookupTaskResultOutput) ResponseCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTaskResult) int { return v.ResponseCount }).(pulumi.IntOutput)
}

// The time when the task is scheduled to be attempted or retried. `schedule_time` will be truncated to the nearest microsecond.
func (o LookupTaskResultOutput) ScheduleTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.ScheduleTime }).(pulumi.StringOutput)
}

// The view specifies which subset of the Task has been returned.
func (o LookupTaskResultOutput) View() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.View }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTaskResultOutput{})
}
