// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an instance.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("google-native:bigtableadmin/v2:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	InstanceId string  `pulumi:"instanceId"`
	Project    *string `pulumi:"project"`
}

type LookupInstanceResult struct {
	// A server-assigned timestamp representing when this Instance was created. For instances created before this field was added (August 2021), this value is `seconds: 0, nanos: 1`.
	CreateTime string `pulumi:"createTime"`
	// The descriptive name for this instance as it appears in UIs. Can be changed at any time, but should be kept globally unique to avoid confusion.
	DisplayName string `pulumi:"displayName"`
	// Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies. They can be used to filter resources and aggregate metrics. * Label keys must be between 1 and 63 characters long and must conform to the regular expression: `\p{Ll}\p{Lo}{0,62}`. * Label values must be between 0 and 63 characters long and must conform to the regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`. * No more than 64 labels can be associated with a given resource. * Keys and values must both be under 128 bytes.
	Labels map[string]string `pulumi:"labels"`
	// The unique name of the instance. Values are of the form `projects/{project}/instances/a-z+[a-z0-9]`.
	Name string `pulumi:"name"`
	// The current state of the instance.
	State string `pulumi:"state"`
	// The type of the instance. Defaults to `PRODUCTION`.
	Type string `pulumi:"type"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

// A server-assigned timestamp representing when this Instance was created. For instances created before this field was added (August 2021), this value is `seconds: 0, nanos: 1`.
func (o LookupInstanceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The descriptive name for this instance as it appears in UIs. Can be changed at any time, but should be kept globally unique to avoid confusion.
func (o LookupInstanceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies. They can be used to filter resources and aggregate metrics. * Label keys must be between 1 and 63 characters long and must conform to the regular expression: `\p{Ll}\p{Lo}{0,62}`. * Label values must be between 0 and 63 characters long and must conform to the regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`. * No more than 64 labels can be associated with a given resource. * Keys and values must both be under 128 bytes.
func (o LookupInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The unique name of the instance. Values are of the form `projects/{project}/instances/a-z+[a-z0-9]`.
func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current state of the instance.
func (o LookupInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

// The type of the instance. Defaults to `PRODUCTION`.
func (o LookupInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
