// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get task resource.
func LookupTask(ctx *pulumi.Context, args *LookupTaskArgs, opts ...pulumi.InvokeOption) (*LookupTaskResult, error) {
	var rv LookupTaskResult
	err := ctx.Invoke("google-native:dataplex/v1:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTaskArgs struct {
	LakeId   string  `pulumi:"lakeId"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	TaskId   string  `pulumi:"taskId"`
}

type LookupTaskResult struct {
	// The time when the task was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. Description of the task.
	Description string `pulumi:"description"`
	// Optional. User friendly display name.
	DisplayName string `pulumi:"displayName"`
	// Spec related to how a task is executed.
	ExecutionSpec GoogleCloudDataplexV1TaskExecutionSpecResponse `pulumi:"executionSpec"`
	// Optional. User-defined labels for the task.
	Labels map[string]string `pulumi:"labels"`
	// The relative resource name of the task, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/ tasks/{task_id}.
	Name string `pulumi:"name"`
	// Config related to running custom Spark tasks.
	Spark GoogleCloudDataplexV1TaskSparkTaskConfigResponse `pulumi:"spark"`
	// Current state of the task.
	State string `pulumi:"state"`
	// Spec related to how often and when a task should be triggered.
	TriggerSpec GoogleCloudDataplexV1TaskTriggerSpecResponse `pulumi:"triggerSpec"`
	// System generated globally unique ID for the task. This ID will be different if the task is deleted and re-created with the same name.
	Uid string `pulumi:"uid"`
	// The time when the task was last updated.
	UpdateTime string `pulumi:"updateTime"`
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
	LakeId   pulumi.StringInput    `pulumi:"lakeId"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	TaskId   pulumi.StringInput    `pulumi:"taskId"`
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

// The time when the task was created.
func (o LookupTaskResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Description of the task.
func (o LookupTaskResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. User friendly display name.
func (o LookupTaskResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Spec related to how a task is executed.
func (o LookupTaskResultOutput) ExecutionSpec() GoogleCloudDataplexV1TaskExecutionSpecResponseOutput {
	return o.ApplyT(func(v LookupTaskResult) GoogleCloudDataplexV1TaskExecutionSpecResponse { return v.ExecutionSpec }).(GoogleCloudDataplexV1TaskExecutionSpecResponseOutput)
}

// Optional. User-defined labels for the task.
func (o LookupTaskResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTaskResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The relative resource name of the task, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/ tasks/{task_id}.
func (o LookupTaskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Name }).(pulumi.StringOutput)
}

// Config related to running custom Spark tasks.
func (o LookupTaskResultOutput) Spark() GoogleCloudDataplexV1TaskSparkTaskConfigResponseOutput {
	return o.ApplyT(func(v LookupTaskResult) GoogleCloudDataplexV1TaskSparkTaskConfigResponse { return v.Spark }).(GoogleCloudDataplexV1TaskSparkTaskConfigResponseOutput)
}

// Current state of the task.
func (o LookupTaskResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.State }).(pulumi.StringOutput)
}

// Spec related to how often and when a task should be triggered.
func (o LookupTaskResultOutput) TriggerSpec() GoogleCloudDataplexV1TaskTriggerSpecResponseOutput {
	return o.ApplyT(func(v LookupTaskResult) GoogleCloudDataplexV1TaskTriggerSpecResponse { return v.TriggerSpec }).(GoogleCloudDataplexV1TaskTriggerSpecResponseOutput)
}

// System generated globally unique ID for the task. This ID will be different if the task is deleted and re-created with the same name.
func (o LookupTaskResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The time when the task was last updated.
func (o LookupTaskResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTaskResultOutput{})
}
