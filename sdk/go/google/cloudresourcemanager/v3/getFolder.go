// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a folder identified by the supplied resource name. Valid folder resource names have the format `folders/{folder_id}` (for example, `folders/1234`). The caller must have `resourcemanager.folders.get` permission on the identified folder.
func LookupFolder(ctx *pulumi.Context, args *LookupFolderArgs, opts ...pulumi.InvokeOption) (*LookupFolderResult, error) {
	var rv LookupFolderResult
	err := ctx.Invoke("google-native:cloudresourcemanager/v3:getFolder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFolderArgs struct {
	FolderId string `pulumi:"folderId"`
}

type LookupFolderResult struct {
	// Timestamp when the folder was created.
	CreateTime string `pulumi:"createTime"`
	// Timestamp when the folder was requested to be deleted.
	DeleteTime string `pulumi:"deleteTime"`
	// The folder's display name. A folder's display name must be unique amongst its siblings. For example, no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters. This is captured by the regular expression: `[\p{L}\p{N}]([\p{L}\p{N}_- ]{0,28}[\p{L}\p{N}])?`.
	DisplayName string `pulumi:"displayName"`
	// A checksum computed by the server based on the current value of the folder resource. This may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag string `pulumi:"etag"`
	// The resource name of the folder. Its format is `folders/{folder_id}`, for example: "folders/1234".
	Name string `pulumi:"name"`
	// The folder's parent's resource name. Updates to the folder's parent must be performed using MoveFolder.
	Parent string `pulumi:"parent"`
	// The lifecycle state of the folder. Updates to the state must be performed using DeleteFolder and UndeleteFolder.
	State string `pulumi:"state"`
	// Timestamp when the folder was last modified.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupFolderOutput(ctx *pulumi.Context, args LookupFolderOutputArgs, opts ...pulumi.InvokeOption) LookupFolderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFolderResult, error) {
			args := v.(LookupFolderArgs)
			r, err := LookupFolder(ctx, &args, opts...)
			return *r, err
		}).(LookupFolderResultOutput)
}

type LookupFolderOutputArgs struct {
	FolderId pulumi.StringInput `pulumi:"folderId"`
}

func (LookupFolderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderArgs)(nil)).Elem()
}

type LookupFolderResultOutput struct{ *pulumi.OutputState }

func (LookupFolderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderResult)(nil)).Elem()
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutput() LookupFolderResultOutput {
	return o
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutputWithContext(ctx context.Context) LookupFolderResultOutput {
	return o
}

// Timestamp when the folder was created.
func (o LookupFolderResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Timestamp when the folder was requested to be deleted.
func (o LookupFolderResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// The folder's display name. A folder's display name must be unique amongst its siblings. For example, no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters. This is captured by the regular expression: `[\p{L}\p{N}]([\p{L}\p{N}_- ]{0,28}[\p{L}\p{N}])?`.
func (o LookupFolderResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// A checksum computed by the server based on the current value of the folder resource. This may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o LookupFolderResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource name of the folder. Its format is `folders/{folder_id}`, for example: "folders/1234".
func (o LookupFolderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Name }).(pulumi.StringOutput)
}

// The folder's parent's resource name. Updates to the folder's parent must be performed using MoveFolder.
func (o LookupFolderResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Parent }).(pulumi.StringOutput)
}

// The lifecycle state of the folder. Updates to the state must be performed using DeleteFolder and UndeleteFolder.
func (o LookupFolderResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.State }).(pulumi.StringOutput)
}

// Timestamp when the folder was last modified.
func (o LookupFolderResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderResultOutput{})
}
