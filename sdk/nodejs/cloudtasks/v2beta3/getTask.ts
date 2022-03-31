// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a task.
 */
export function getTask(args: GetTaskArgs, opts?: pulumi.InvokeOptions): Promise<GetTaskResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:cloudtasks/v2beta3:getTask", {
        "location": args.location,
        "project": args.project,
        "queueId": args.queueId,
        "responseView": args.responseView,
        "taskId": args.taskId,
    }, opts);
}

export interface GetTaskArgs {
    location: string;
    project?: string;
    queueId: string;
    responseView?: string;
    taskId: string;
}

export interface GetTaskResult {
    /**
     * HTTP request that is sent to the App Engine app handler. An App Engine task is a task that has AppEngineHttpRequest set.
     */
    readonly appEngineHttpRequest: outputs.cloudtasks.v2beta3.AppEngineHttpRequestResponse;
    /**
     * The time that the task was created. `create_time` will be truncated to the nearest second.
     */
    readonly createTime: string;
    /**
     * The number of attempts dispatched. This count includes attempts which have been dispatched but haven't received a response.
     */
    readonly dispatchCount: number;
    /**
     * The deadline for requests sent to the worker. If the worker does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the task according to the RetryConfig. Note that when the request is cancelled, Cloud Tasks will stop listening for the response, but whether the worker stops processing depends on the worker. For example, if the worker is stuck, it may not react to cancelled requests. The default and maximum values depend on the type of request: * For HTTP tasks, the default is 10 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine tasks, 0 indicates that the request has the default deadline. The default deadline depends on the [scaling type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling) of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. Regardless of the task's `dispatch_deadline`, the app handler will not run for longer than than the service's timeout. We recommend setting the `dispatch_deadline` to at most a few seconds more than the app handler's timeout. For more information see [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts). `dispatch_deadline` will be truncated to the nearest millisecond. The deadline is an approximate deadline.
     */
    readonly dispatchDeadline: string;
    /**
     * The status of the task's first attempt. Only dispatch_time will be set. The other Attempt information is not retained by Cloud Tasks.
     */
    readonly firstAttempt: outputs.cloudtasks.v2beta3.AttemptResponse;
    /**
     * HTTP request that is sent to the task's target. An HTTP task is a task that has HttpRequest set.
     */
    readonly httpRequest: outputs.cloudtasks.v2beta3.HttpRequestResponse;
    /**
     * The status of the task's last attempt.
     */
    readonly lastAttempt: outputs.cloudtasks.v2beta3.AttemptResponse;
    /**
     * Optionally caller-specified in CreateTask. The task name. The task name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the task's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters. * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
     */
    readonly name: string;
    /**
     * Pull Message contained in a task in a PULL queue type. This payload type cannot be explicitly set through Cloud Tasks API. Its purpose, currently is to provide backward compatibility with App Engine Task Queue [pull](https://cloud.google.com/appengine/docs/standard/java/taskqueue/pull/) queues to provide a way to inspect contents of pull tasks through the CloudTasks.GetTask.
     */
    readonly pullMessage: outputs.cloudtasks.v2beta3.PullMessageResponse;
    /**
     * The number of attempts which have received a response.
     */
    readonly responseCount: number;
    /**
     * The time when the task is scheduled to be attempted. For App Engine queues, this is when the task will be attempted or retried. `schedule_time` will be truncated to the nearest microsecond.
     */
    readonly scheduleTime: string;
    /**
     * The view specifies which subset of the Task has been returned.
     */
    readonly view: string;
}

export function getTaskOutput(args: GetTaskOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTaskResult> {
    return pulumi.output(args).apply(a => getTask(a, opts))
}

export interface GetTaskOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    queueId: pulumi.Input<string>;
    responseView?: pulumi.Input<string>;
    taskId: pulumi.Input<string>;
}
