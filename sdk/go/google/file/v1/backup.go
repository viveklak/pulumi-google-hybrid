// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a backup.
// Auto-naming is currently not supported for this resource.
type Backup struct {
	pulumi.CustomResourceState

	// Capacity of the source file share when the backup was created.
	CapacityGb pulumi.StringOutput `pulumi:"capacityGb"`
	// The time when the backup was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description pulumi.StringOutput `pulumi:"description"`
	// Amount of bytes that will be downloaded if the backup is restored. This may be different than storage bytes, since sequential backups of the same disk will share storage.
	DownloadBytes pulumi.StringOutput `pulumi:"downloadBytes"`
	// Resource labels to represent user provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name of the backup, in the format `projects/{project_number}/locations/{location_id}/backups/{backup_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// Name of the file share in the source Cloud Filestore instance that the backup is created from.
	SourceFileShare pulumi.StringOutput `pulumi:"sourceFileShare"`
	// The resource name of the source Cloud Filestore instance, in the format `projects/{project_number}/locations/{location_id}/instances/{instance_id}`, used to create this backup.
	SourceInstance pulumi.StringOutput `pulumi:"sourceInstance"`
	// The service tier of the source Cloud Filestore instance that this backup is created from.
	SourceInstanceTier pulumi.StringOutput `pulumi:"sourceInstanceTier"`
	// The backup state.
	State pulumi.StringOutput `pulumi:"state"`
	// The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.
	StorageBytes pulumi.StringOutput `pulumi:"storageBytes"`
}

// NewBackup registers a new resource with the given unique name, arguments, and options.
func NewBackup(ctx *pulumi.Context,
	name string, args *BackupArgs, opts ...pulumi.ResourceOption) (*Backup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupId == nil {
		return nil, errors.New("invalid value for required argument 'BackupId'")
	}
	var resource Backup
	err := ctx.RegisterResource("google-native:file/v1:Backup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackup gets an existing Backup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupState, opts ...pulumi.ResourceOption) (*Backup, error) {
	var resource Backup
	err := ctx.ReadResource("google-native:file/v1:Backup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backup resources.
type backupState struct {
}

type BackupState struct {
}

func (BackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupState)(nil)).Elem()
}

type backupArgs struct {
	// Required. The ID to use for the backup. The ID must be unique within the specified project and location. This value must start with a lowercase letter followed by up to 62 lowercase letters, numbers, or hyphens, and cannot end with a hyphen. Values that do not match this pattern will trigger an INVALID_ARGUMENT error.
	BackupId string `pulumi:"backupId"`
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description *string `pulumi:"description"`
	// Resource labels to represent user provided metadata.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
	// Name of the file share in the source Cloud Filestore instance that the backup is created from.
	SourceFileShare *string `pulumi:"sourceFileShare"`
	// The resource name of the source Cloud Filestore instance, in the format `projects/{project_number}/locations/{location_id}/instances/{instance_id}`, used to create this backup.
	SourceInstance *string `pulumi:"sourceInstance"`
}

// The set of arguments for constructing a Backup resource.
type BackupArgs struct {
	// Required. The ID to use for the backup. The ID must be unique within the specified project and location. This value must start with a lowercase letter followed by up to 62 lowercase letters, numbers, or hyphens, and cannot end with a hyphen. Values that do not match this pattern will trigger an INVALID_ARGUMENT error.
	BackupId pulumi.StringInput
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description pulumi.StringPtrInput
	// Resource labels to represent user provided metadata.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Name of the file share in the source Cloud Filestore instance that the backup is created from.
	SourceFileShare pulumi.StringPtrInput
	// The resource name of the source Cloud Filestore instance, in the format `projects/{project_number}/locations/{location_id}/instances/{instance_id}`, used to create this backup.
	SourceInstance pulumi.StringPtrInput
}

func (BackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupArgs)(nil)).Elem()
}

type BackupInput interface {
	pulumi.Input

	ToBackupOutput() BackupOutput
	ToBackupOutputWithContext(ctx context.Context) BackupOutput
}

func (*Backup) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (i *Backup) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i *Backup) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

type BackupOutput struct{ *pulumi.OutputState }

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupInput)(nil)).Elem(), &Backup{})
	pulumi.RegisterOutputType(BackupOutput{})
}
