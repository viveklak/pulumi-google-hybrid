// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a job.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("google-native:ml/v1:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	JobId   string  `pulumi:"jobId"`
	Project *string `pulumi:"project"`
}

type LookupJobResult struct {
	// When the job was created.
	CreateTime string `pulumi:"createTime"`
	// When the job processing was completed.
	EndTime string `pulumi:"endTime"`
	// The details of a failure or a cancellation.
	ErrorMessage string `pulumi:"errorMessage"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a job from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform job updates in order to avoid race conditions: An `etag` is returned in the response to `GetJob`, and systems are expected to put that etag in the request to `UpdateJob` to ensure that their change will be applied to the same version of the job.
	Etag string `pulumi:"etag"`
	// The user-specified id of the job.
	JobId string `pulumi:"jobId"`
	// It's only effect when the job is in QUEUED state. If it's positive, it indicates the job's position in the job scheduler. It's 0 when the job is already scheduled.
	JobPosition string `pulumi:"jobPosition"`
	// Optional. One or more labels that you can add, to organize your jobs. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels.
	Labels map[string]string `pulumi:"labels"`
	// Input parameters to create a prediction job.
	PredictionInput GoogleCloudMlV1__PredictionInputResponse `pulumi:"predictionInput"`
	// The current prediction job result.
	PredictionOutput GoogleCloudMlV1__PredictionOutputResponse `pulumi:"predictionOutput"`
	// When the job processing was started.
	StartTime string `pulumi:"startTime"`
	// The detailed state of a job.
	State string `pulumi:"state"`
	// Input parameters to create a training job.
	TrainingInput GoogleCloudMlV1__TrainingInputResponse `pulumi:"trainingInput"`
	// The current training job result.
	TrainingOutput GoogleCloudMlV1__TrainingOutputResponse `pulumi:"trainingOutput"`
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			return *r, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	JobId   pulumi.StringInput    `pulumi:"jobId"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

// When the job was created.
func (o LookupJobResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// When the job processing was completed.
func (o LookupJobResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The details of a failure or a cancellation.
func (o LookupJobResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a job from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform job updates in order to avoid race conditions: An `etag` is returned in the response to `GetJob`, and systems are expected to put that etag in the request to `UpdateJob` to ensure that their change will be applied to the same version of the job.
func (o LookupJobResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The user-specified id of the job.
func (o LookupJobResultOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.JobId }).(pulumi.StringOutput)
}

// It's only effect when the job is in QUEUED state. If it's positive, it indicates the job's position in the job scheduler. It's 0 when the job is already scheduled.
func (o LookupJobResultOutput) JobPosition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.JobPosition }).(pulumi.StringOutput)
}

// Optional. One or more labels that you can add, to organize your jobs. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels.
func (o LookupJobResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Input parameters to create a prediction job.
func (o LookupJobResultOutput) PredictionInput() GoogleCloudMlV1__PredictionInputResponseOutput {
	return o.ApplyT(func(v LookupJobResult) GoogleCloudMlV1__PredictionInputResponse { return v.PredictionInput }).(GoogleCloudMlV1__PredictionInputResponseOutput)
}

// The current prediction job result.
func (o LookupJobResultOutput) PredictionOutput() GoogleCloudMlV1__PredictionOutputResponseOutput {
	return o.ApplyT(func(v LookupJobResult) GoogleCloudMlV1__PredictionOutputResponse { return v.PredictionOutput }).(GoogleCloudMlV1__PredictionOutputResponseOutput)
}

// When the job processing was started.
func (o LookupJobResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// The detailed state of a job.
func (o LookupJobResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.State }).(pulumi.StringOutput)
}

// Input parameters to create a training job.
func (o LookupJobResultOutput) TrainingInput() GoogleCloudMlV1__TrainingInputResponseOutput {
	return o.ApplyT(func(v LookupJobResult) GoogleCloudMlV1__TrainingInputResponse { return v.TrainingInput }).(GoogleCloudMlV1__TrainingInputResponseOutput)
}

// The current training job result.
func (o LookupJobResultOutput) TrainingOutput() GoogleCloudMlV1__TrainingOutputResponseOutput {
	return o.ApplyT(func(v LookupJobResult) GoogleCloudMlV1__TrainingOutputResponse { return v.TrainingOutput }).(GoogleCloudMlV1__TrainingOutputResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
