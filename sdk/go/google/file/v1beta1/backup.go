// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a backup.
type Backup struct {
	pulumi.CustomResourceState
}

// NewBackup registers a new resource with the given unique name, arguments, and options.
func NewBackup(ctx *pulumi.Context,
	name string, args *BackupArgs, opts ...pulumi.ResourceOption) (*Backup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupsId == nil {
		return nil, errors.New("invalid value for required argument 'BackupsId'")
	}
	if args.LocationsId == nil {
		return nil, errors.New("invalid value for required argument 'LocationsId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	var resource Backup
	err := ctx.RegisterResource("google-cloud:file/v1beta1:Backup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-cloud:file/v1beta1:Backup", name, id, state, &resource, opts...)
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
	BackupsId string `pulumi:"backupsId"`
	// Output only. Capacity of the source file share when the backup was created.
	CapacityGb *string `pulumi:"capacityGb"`
	// Output only. The time when the backup was created.
	CreateTime *string `pulumi:"createTime"`
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description *string `pulumi:"description"`
	// Output only. Amount of bytes that will be downloaded if the backup is restored
	DownloadBytes *string `pulumi:"downloadBytes"`
	// Resource labels to represent user provided metadata.
	Labels      map[string]string `pulumi:"labels"`
	LocationsId string            `pulumi:"locationsId"`
	// Output only. The resource name of the backup, in the format projects/{project_id}/locations/{location_id}/backups/{backup_id}.
	Name       *string `pulumi:"name"`
	ProjectsId string  `pulumi:"projectsId"`
	// Name of the file share in the source Cloud Filestore instance that the backup is created from.
	SourceFileShare *string `pulumi:"sourceFileShare"`
	// The resource name of the source Cloud Filestore instance, in the format projects/{project_id}/locations/{location_id}/instances/{instance_id}, used to create this backup.
	SourceInstance *string `pulumi:"sourceInstance"`
	// Output only. The service tier of the source Cloud Filestore instance that this backup is created from.
	SourceInstanceTier *string `pulumi:"sourceInstanceTier"`
	// Output only. The backup state.
	State *string `pulumi:"state"`
	// Output only. The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.
	StorageBytes *string `pulumi:"storageBytes"`
}

// The set of arguments for constructing a Backup resource.
type BackupArgs struct {
	BackupsId pulumi.StringInput
	// Output only. Capacity of the source file share when the backup was created.
	CapacityGb pulumi.StringPtrInput
	// Output only. The time when the backup was created.
	CreateTime pulumi.StringPtrInput
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description pulumi.StringPtrInput
	// Output only. Amount of bytes that will be downloaded if the backup is restored
	DownloadBytes pulumi.StringPtrInput
	// Resource labels to represent user provided metadata.
	Labels      pulumi.StringMapInput
	LocationsId pulumi.StringInput
	// Output only. The resource name of the backup, in the format projects/{project_id}/locations/{location_id}/backups/{backup_id}.
	Name       pulumi.StringPtrInput
	ProjectsId pulumi.StringInput
	// Name of the file share in the source Cloud Filestore instance that the backup is created from.
	SourceFileShare pulumi.StringPtrInput
	// The resource name of the source Cloud Filestore instance, in the format projects/{project_id}/locations/{location_id}/instances/{instance_id}, used to create this backup.
	SourceInstance pulumi.StringPtrInput
	// Output only. The service tier of the source Cloud Filestore instance that this backup is created from.
	SourceInstanceTier pulumi.StringPtrInput
	// Output only. The backup state.
	State pulumi.StringPtrInput
	// Output only. The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.
	StorageBytes pulumi.StringPtrInput
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
	return reflect.TypeOf((*Backup)(nil))
}

func (i *Backup) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i *Backup) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

type BackupOutput struct {
	*pulumi.OutputState
}

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Backup)(nil))
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BackupOutput{})
}
