// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the project identified by the specified `name` (for example, `projects/415104041262`). The caller must have `resourcemanager.projects.get` permission for this project.
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("google-native:cloudresourcemanager/v3:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	Project *string `pulumi:"project"`
}

type LookupProjectResult struct {
	// Creation time.
	CreateTime string `pulumi:"createTime"`
	// The time at which this resource was requested for deletion.
	DeleteTime string `pulumi:"deleteTime"`
	// Optional. A user-assigned display name of the project. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point. Example: `My Project`
	DisplayName string `pulumi:"displayName"`
	// A checksum computed by the server based on the current value of the Project resource. This may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag string `pulumi:"etag"`
	// Optional. The labels associated with this project. Label keys must be between 1 and 63 characters long and must conform to the following regular expression: \[a-z\](\[-a-z0-9\]*\[a-z0-9\])?. Label values must be between 0 and 63 characters long and must conform to the regular expression (\[a-z\](\[-a-z0-9\]*\[a-z0-9\])?)?. No more than 256 labels can be associated with a given resource. Clients should store labels in a representation such as JSON that does not depend on specific characters being disallowed. Example: `"myBusinessDimension" : "businessValue"`
	Labels map[string]string `pulumi:"labels"`
	// The unique resource name of the project. It is an int64 generated number prefixed by "projects/". Example: `projects/415104041262`
	Name string `pulumi:"name"`
	// Optional. A reference to a parent Resource. eg., `organizations/123` or `folders/876`.
	Parent string `pulumi:"parent"`
	// Immutable. The unique, user-assigned id of the project. It must be 6 to 30 lowercase ASCII letters, digits, or hyphens. It must start with a letter. Trailing hyphens are prohibited. Example: `tokyo-rain-123`
	ProjectId string `pulumi:"projectId"`
	// The project lifecycle state.
	State string `pulumi:"state"`
	// The most recent time this resource was modified.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			return *r, err
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// Creation time.
func (o LookupProjectResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The time at which this resource was requested for deletion.
func (o LookupProjectResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// Optional. A user-assigned display name of the project. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point. Example: `My Project`
func (o LookupProjectResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// A checksum computed by the server based on the current value of the Project resource. This may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o LookupProjectResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Optional. The labels associated with this project. Label keys must be between 1 and 63 characters long and must conform to the following regular expression: \[a-z\](\[-a-z0-9\]*\[a-z0-9\])?. Label values must be between 0 and 63 characters long and must conform to the regular expression (\[a-z\](\[-a-z0-9\]*\[a-z0-9\])?)?. No more than 256 labels can be associated with a given resource. Clients should store labels in a representation such as JSON that does not depend on specific characters being disallowed. Example: `"myBusinessDimension" : "businessValue"`
func (o LookupProjectResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProjectResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The unique resource name of the project. It is an int64 generated number prefixed by "projects/". Example: `projects/415104041262`
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. A reference to a parent Resource. eg., `organizations/123` or `folders/876`.
func (o LookupProjectResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Parent }).(pulumi.StringOutput)
}

// Immutable. The unique, user-assigned id of the project. It must be 6 to 30 lowercase ASCII letters, digits, or hyphens. It must start with a letter. Trailing hyphens are prohibited. Example: `tokyo-rain-123`
func (o LookupProjectResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The project lifecycle state.
func (o LookupProjectResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.State }).(pulumi.StringOutput)
}

// The most recent time this resource was modified.
func (o LookupProjectResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
