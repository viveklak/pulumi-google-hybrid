// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a transfer job that runs periodically.
type TransferJob struct {
	pulumi.CustomResourceState

	// The time that the transfer job was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The time that the transfer job was deleted.
	DeletionTime pulumi.StringOutput `pulumi:"deletionTime"`
	// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
	Description pulumi.StringOutput `pulumi:"description"`
	// The time that the transfer job was last modified.
	LastModificationTime pulumi.StringOutput `pulumi:"lastModificationTime"`
	// The name of the most recently started TransferOperation of this JobConfig. Present if and only if at least one TransferOperation has been created for this JobConfig.
	LatestOperationName pulumi.StringOutput `pulumi:"latestOperationName"`
	// A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service will assign a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. This name must not start with 'transferJobs/OPI'. 'transferJobs/OPI' is a reserved prefix. Example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Invalid job names will fail with an INVALID_ARGUMENT error.
	Name pulumi.StringOutput `pulumi:"name"`
	// Notification configuration.
	NotificationConfig NotificationConfigResponseOutput `pulumi:"notificationConfig"`
	// The ID of the Google Cloud Platform Project that owns the job.
	Project pulumi.StringOutput `pulumi:"project"`
	// Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job will never execute a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
	Schedule ScheduleResponseOutput `pulumi:"schedule"`
	// Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
	Status pulumi.StringOutput `pulumi:"status"`
	// Transfer specification.
	TransferSpec TransferSpecResponseOutput `pulumi:"transferSpec"`
}

// NewTransferJob registers a new resource with the given unique name, arguments, and options.
func NewTransferJob(ctx *pulumi.Context,
	name string, args *TransferJobArgs, opts ...pulumi.ResourceOption) (*TransferJob, error) {
	if args == nil {
		args = &TransferJobArgs{}
	}

	var resource TransferJob
	err := ctx.RegisterResource("google-native:storagetransfer/v1:TransferJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransferJob gets an existing TransferJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransferJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransferJobState, opts ...pulumi.ResourceOption) (*TransferJob, error) {
	var resource TransferJob
	err := ctx.ReadResource("google-native:storagetransfer/v1:TransferJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransferJob resources.
type transferJobState struct {
	// The time that the transfer job was created.
	CreationTime *string `pulumi:"creationTime"`
	// The time that the transfer job was deleted.
	DeletionTime *string `pulumi:"deletionTime"`
	// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
	Description *string `pulumi:"description"`
	// The time that the transfer job was last modified.
	LastModificationTime *string `pulumi:"lastModificationTime"`
	// The name of the most recently started TransferOperation of this JobConfig. Present if and only if at least one TransferOperation has been created for this JobConfig.
	LatestOperationName *string `pulumi:"latestOperationName"`
	// A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service will assign a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. This name must not start with 'transferJobs/OPI'. 'transferJobs/OPI' is a reserved prefix. Example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Invalid job names will fail with an INVALID_ARGUMENT error.
	Name *string `pulumi:"name"`
	// Notification configuration.
	NotificationConfig *NotificationConfigResponse `pulumi:"notificationConfig"`
	// The ID of the Google Cloud Platform Project that owns the job.
	Project *string `pulumi:"project"`
	// Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job will never execute a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
	Schedule *ScheduleResponse `pulumi:"schedule"`
	// Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
	Status *string `pulumi:"status"`
	// Transfer specification.
	TransferSpec *TransferSpecResponse `pulumi:"transferSpec"`
}

type TransferJobState struct {
	// The time that the transfer job was created.
	CreationTime pulumi.StringPtrInput
	// The time that the transfer job was deleted.
	DeletionTime pulumi.StringPtrInput
	// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
	Description pulumi.StringPtrInput
	// The time that the transfer job was last modified.
	LastModificationTime pulumi.StringPtrInput
	// The name of the most recently started TransferOperation of this JobConfig. Present if and only if at least one TransferOperation has been created for this JobConfig.
	LatestOperationName pulumi.StringPtrInput
	// A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service will assign a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. This name must not start with 'transferJobs/OPI'. 'transferJobs/OPI' is a reserved prefix. Example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Invalid job names will fail with an INVALID_ARGUMENT error.
	Name pulumi.StringPtrInput
	// Notification configuration.
	NotificationConfig NotificationConfigResponsePtrInput
	// The ID of the Google Cloud Platform Project that owns the job.
	Project pulumi.StringPtrInput
	// Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job will never execute a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
	Schedule ScheduleResponsePtrInput
	// Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
	Status pulumi.StringPtrInput
	// Transfer specification.
	TransferSpec TransferSpecResponsePtrInput
}

func (TransferJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*transferJobState)(nil)).Elem()
}

type transferJobArgs struct {
	// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
	Description *string `pulumi:"description"`
	// The name of the most recently started TransferOperation of this JobConfig. Present if and only if at least one TransferOperation has been created for this JobConfig.
	LatestOperationName *string `pulumi:"latestOperationName"`
	// A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service will assign a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. This name must not start with 'transferJobs/OPI'. 'transferJobs/OPI' is a reserved prefix. Example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Invalid job names will fail with an INVALID_ARGUMENT error.
	Name *string `pulumi:"name"`
	// Notification configuration.
	NotificationConfig *NotificationConfig `pulumi:"notificationConfig"`
	// The ID of the Google Cloud Platform Project that owns the job.
	Project *string `pulumi:"project"`
	// Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job will never execute a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
	Schedule *Schedule `pulumi:"schedule"`
	// Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
	Status *string `pulumi:"status"`
	// Transfer specification.
	TransferSpec *TransferSpec `pulumi:"transferSpec"`
}

// The set of arguments for constructing a TransferJob resource.
type TransferJobArgs struct {
	// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
	Description pulumi.StringPtrInput
	// The name of the most recently started TransferOperation of this JobConfig. Present if and only if at least one TransferOperation has been created for this JobConfig.
	LatestOperationName pulumi.StringPtrInput
	// A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service will assign a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. This name must not start with 'transferJobs/OPI'. 'transferJobs/OPI' is a reserved prefix. Example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Invalid job names will fail with an INVALID_ARGUMENT error.
	Name pulumi.StringPtrInput
	// Notification configuration.
	NotificationConfig NotificationConfigPtrInput
	// The ID of the Google Cloud Platform Project that owns the job.
	Project pulumi.StringPtrInput
	// Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job will never execute a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
	Schedule SchedulePtrInput
	// Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
	Status pulumi.StringPtrInput
	// Transfer specification.
	TransferSpec TransferSpecPtrInput
}

func (TransferJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transferJobArgs)(nil)).Elem()
}

type TransferJobInput interface {
	pulumi.Input

	ToTransferJobOutput() TransferJobOutput
	ToTransferJobOutputWithContext(ctx context.Context) TransferJobOutput
}

func (*TransferJob) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferJob)(nil))
}

func (i *TransferJob) ToTransferJobOutput() TransferJobOutput {
	return i.ToTransferJobOutputWithContext(context.Background())
}

func (i *TransferJob) ToTransferJobOutputWithContext(ctx context.Context) TransferJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferJobOutput)
}

type TransferJobOutput struct {
	*pulumi.OutputState
}

func (TransferJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferJob)(nil))
}

func (o TransferJobOutput) ToTransferJobOutput() TransferJobOutput {
	return o
}

func (o TransferJobOutput) ToTransferJobOutputWithContext(ctx context.Context) TransferJobOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TransferJobOutput{})
}
