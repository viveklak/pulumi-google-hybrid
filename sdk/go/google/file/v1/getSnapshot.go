// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a specific snapshot.
func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	var rv LookupSnapshotResult
	err := ctx.Invoke("google-native:file/v1:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSnapshotArgs struct {
	InstanceId string  `pulumi:"instanceId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
	SnapshotId string  `pulumi:"snapshotId"`
}

type LookupSnapshotResult struct {
	// The time when the snapshot was created.
	CreateTime string `pulumi:"createTime"`
	// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description string `pulumi:"description"`
	// The amount of bytes needed to allocate a full copy of the snapshot content
	FilesystemUsedBytes string `pulumi:"filesystemUsedBytes"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the snapshot, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}/snapshots/{snapshot_id}`.
	Name string `pulumi:"name"`
	// The snapshot state.
	State string `pulumi:"state"`
}

func LookupSnapshotOutput(ctx *pulumi.Context, args LookupSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotResult, error) {
			args := v.(LookupSnapshotArgs)
			r, err := LookupSnapshot(ctx, &args, opts...)
			return *r, err
		}).(LookupSnapshotResultOutput)
}

type LookupSnapshotOutputArgs struct {
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
	SnapshotId pulumi.StringInput    `pulumi:"snapshotId"`
}

func (LookupSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotArgs)(nil)).Elem()
}

type LookupSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotResult)(nil)).Elem()
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutput() LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutputWithContext(ctx context.Context) LookupSnapshotResultOutput {
	return o
}

// The time when the snapshot was created.
func (o LookupSnapshotResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
func (o LookupSnapshotResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Description }).(pulumi.StringOutput)
}

// The amount of bytes needed to allocate a full copy of the snapshot content
func (o LookupSnapshotResultOutput) FilesystemUsedBytes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.FilesystemUsedBytes }).(pulumi.StringOutput)
}

// Resource labels to represent user provided metadata.
func (o LookupSnapshotResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSnapshotResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name of the snapshot, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}/snapshots/{snapshot_id}`.
func (o LookupSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

// The snapshot state.
func (o LookupSnapshotResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotResultOutput{})
}
