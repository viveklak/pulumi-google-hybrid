// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single TargetProject. NOTE: TargetProject is a global resource; hence the only supported value for location is `global`.
func LookupTargetProject(ctx *pulumi.Context, args *LookupTargetProjectArgs, opts ...pulumi.InvokeOption) (*LookupTargetProjectResult, error) {
	var rv LookupTargetProjectResult
	err := ctx.Invoke("google-native:vmmigration/v1alpha1:getTargetProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetProjectArgs struct {
	Location        string  `pulumi:"location"`
	Project         *string `pulumi:"project"`
	TargetProjectId string  `pulumi:"targetProjectId"`
}

type LookupTargetProjectResult struct {
	// The time this target project resource was created (not related to when the Compute Engine project it points to was created).
	CreateTime string `pulumi:"createTime"`
	// The target project's description.
	Description string `pulumi:"description"`
	// The name of the target project.
	Name string `pulumi:"name"`
	// The target project ID (number) or project name.
	Project string `pulumi:"project"`
	// The last time the target project resource was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupTargetProjectOutput(ctx *pulumi.Context, args LookupTargetProjectOutputArgs, opts ...pulumi.InvokeOption) LookupTargetProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTargetProjectResult, error) {
			args := v.(LookupTargetProjectArgs)
			r, err := LookupTargetProject(ctx, &args, opts...)
			return *r, err
		}).(LookupTargetProjectResultOutput)
}

type LookupTargetProjectOutputArgs struct {
	Location        pulumi.StringInput    `pulumi:"location"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
	TargetProjectId pulumi.StringInput    `pulumi:"targetProjectId"`
}

func (LookupTargetProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetProjectArgs)(nil)).Elem()
}

type LookupTargetProjectResultOutput struct{ *pulumi.OutputState }

func (LookupTargetProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetProjectResult)(nil)).Elem()
}

func (o LookupTargetProjectResultOutput) ToLookupTargetProjectResultOutput() LookupTargetProjectResultOutput {
	return o
}

func (o LookupTargetProjectResultOutput) ToLookupTargetProjectResultOutputWithContext(ctx context.Context) LookupTargetProjectResultOutput {
	return o
}

// The time this target project resource was created (not related to when the Compute Engine project it points to was created).
func (o LookupTargetProjectResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetProjectResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The target project's description.
func (o LookupTargetProjectResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetProjectResult) string { return v.Description }).(pulumi.StringOutput)
}

// The name of the target project.
func (o LookupTargetProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// The target project ID (number) or project name.
func (o LookupTargetProjectResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetProjectResult) string { return v.Project }).(pulumi.StringOutput)
}

// The last time the target project resource was updated.
func (o LookupTargetProjectResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetProjectResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTargetProjectResultOutput{})
}
