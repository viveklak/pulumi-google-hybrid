// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a transfer job.
func LookupTransferJob(ctx *pulumi.Context, args *LookupTransferJobArgs, opts ...pulumi.InvokeOption) (*LookupTransferJobResult, error) {
	var rv LookupTransferJobResult
	err := ctx.Invoke("google-native:storagetransfer/v1:getTransferJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransferJobArgs struct {
	ProjectId     string `pulumi:"projectId"`
	TransferJobId string `pulumi:"transferJobId"`
}

type LookupTransferJobResult struct {
	// The time that the transfer job was created.
	CreationTime string `pulumi:"creationTime"`
	// The time that the transfer job was deleted.
	DeletionTime string `pulumi:"deletionTime"`
	// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
	Description string `pulumi:"description"`
	// The time that the transfer job was last modified.
	LastModificationTime string `pulumi:"lastModificationTime"`
	// The name of the most recently started TransferOperation of this JobConfig. Present if a TransferOperation has been created for this JobConfig.
	LatestOperationName string `pulumi:"latestOperationName"`
	// Logging configuration.
	LoggingConfig LoggingConfigResponse `pulumi:"loggingConfig"`
	// A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service assigns a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. For transfers involving PosixFilesystem, this name must start with `transferJobs/OPI` specifically. For all other transfer types, this name must not start with `transferJobs/OPI`. Non-PosixFilesystem example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` PosixFilesystem example: `"transferJobs/OPI^[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Applications must not rely on the enforcement of naming requirements involving OPI. Invalid job names fail with an INVALID_ARGUMENT error.
	Name string `pulumi:"name"`
	// Notification configuration. This is not supported for transfers involving PosixFilesystem.
	NotificationConfig NotificationConfigResponse `pulumi:"notificationConfig"`
	// The ID of the Google Cloud project that owns the job.
	Project string `pulumi:"project"`
	// Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job never executes a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
	Schedule ScheduleResponse `pulumi:"schedule"`
	// Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
	Status string `pulumi:"status"`
	// Transfer specification.
	TransferSpec TransferSpecResponse `pulumi:"transferSpec"`
}

func LookupTransferJobOutput(ctx *pulumi.Context, args LookupTransferJobOutputArgs, opts ...pulumi.InvokeOption) LookupTransferJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTransferJobResult, error) {
			args := v.(LookupTransferJobArgs)
			r, err := LookupTransferJob(ctx, &args, opts...)
			return *r, err
		}).(LookupTransferJobResultOutput)
}

type LookupTransferJobOutputArgs struct {
	ProjectId     pulumi.StringInput `pulumi:"projectId"`
	TransferJobId pulumi.StringInput `pulumi:"transferJobId"`
}

func (LookupTransferJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransferJobArgs)(nil)).Elem()
}

type LookupTransferJobResultOutput struct{ *pulumi.OutputState }

func (LookupTransferJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransferJobResult)(nil)).Elem()
}

func (o LookupTransferJobResultOutput) ToLookupTransferJobResultOutput() LookupTransferJobResultOutput {
	return o
}

func (o LookupTransferJobResultOutput) ToLookupTransferJobResultOutputWithContext(ctx context.Context) LookupTransferJobResultOutput {
	return o
}

// The time that the transfer job was created.
func (o LookupTransferJobResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransferJobResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// The time that the transfer job was deleted.
func (o LookupTransferJobResultOutput) DeletionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransferJobResult) string { return v.DeletionTime }).(pulumi.StringOutput)
}

// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
func (o LookupTransferJobResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransferJobResult) string { return v.Description }).(pulumi.StringOutput)
}

// The time that the transfer job was last modified.
func (o LookupTransferJobResultOutput) LastModificationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransferJobResult) string { return v.LastModificationTime }).(pulumi.StringOutput)
}

// The name of the most recently started TransferOperation of this JobConfig. Present if a TransferOperation has been created for this JobConfig.
func (o LookupTransferJobResultOutput) LatestOperationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransferJobResult) string { return v.LatestOperationName }).(pulumi.StringOutput)
}

// Logging configuration.
func (o LookupTransferJobResultOutput) LoggingConfig() LoggingConfigResponseOutput {
	return o.ApplyT(func(v LookupTransferJobResult) LoggingConfigResponse { return v.LoggingConfig }).(LoggingConfigResponseOutput)
}

// A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service assigns a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. For transfers involving PosixFilesystem, this name must start with `transferJobs/OPI` specifically. For all other transfer types, this name must not start with `transferJobs/OPI`. Non-PosixFilesystem example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` PosixFilesystem example: `"transferJobs/OPI^[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Applications must not rely on the enforcement of naming requirements involving OPI. Invalid job names fail with an INVALID_ARGUMENT error.
func (o LookupTransferJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransferJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// Notification configuration. This is not supported for transfers involving PosixFilesystem.
func (o LookupTransferJobResultOutput) NotificationConfig() NotificationConfigResponseOutput {
	return o.ApplyT(func(v LookupTransferJobResult) NotificationConfigResponse { return v.NotificationConfig }).(NotificationConfigResponseOutput)
}

// The ID of the Google Cloud project that owns the job.
func (o LookupTransferJobResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransferJobResult) string { return v.Project }).(pulumi.StringOutput)
}

// Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job never executes a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
func (o LookupTransferJobResultOutput) Schedule() ScheduleResponseOutput {
	return o.ApplyT(func(v LookupTransferJobResult) ScheduleResponse { return v.Schedule }).(ScheduleResponseOutput)
}

// Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
func (o LookupTransferJobResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransferJobResult) string { return v.Status }).(pulumi.StringOutput)
}

// Transfer specification.
func (o LookupTransferJobResultOutput) TransferSpec() TransferSpecResponseOutput {
	return o.ApplyT(func(v LookupTransferJobResult) TransferSpecResponse { return v.TransferSpec }).(TransferSpecResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTransferJobResultOutput{})
}
