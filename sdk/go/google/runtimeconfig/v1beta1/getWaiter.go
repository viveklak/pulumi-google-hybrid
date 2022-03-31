// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a single waiter.
func LookupWaiter(ctx *pulumi.Context, args *LookupWaiterArgs, opts ...pulumi.InvokeOption) (*LookupWaiterResult, error) {
	var rv LookupWaiterResult
	err := ctx.Invoke("google-native:runtimeconfig/v1beta1:getWaiter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWaiterArgs struct {
	ConfigId string  `pulumi:"configId"`
	Project  *string `pulumi:"project"`
	WaiterId string  `pulumi:"waiterId"`
}

type LookupWaiterResult struct {
	// The instant at which this Waiter resource was created. Adding the value of `timeout` to this instant yields the timeout deadline for the waiter.
	CreateTime string `pulumi:"createTime"`
	// If the value is `false`, it means the waiter is still waiting for one of its conditions to be met. If true, the waiter has finished. If the waiter finished due to a timeout or failure, `error` will be set.
	Done bool `pulumi:"done"`
	// If the waiter ended due to a failure or timeout, this value will be set.
	Error StatusResponse `pulumi:"error"`
	// [Optional] The failure condition of this waiter. If this condition is met, `done` will be set to `true` and the `error` code will be set to `ABORTED`. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated. This value is optional; if no failure condition is set, the only failure scenario will be a timeout.
	Failure EndConditionResponse `pulumi:"failure"`
	// The name of the Waiter resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/waiters/[WAITER_NAME] The `[PROJECT_ID]` must be a valid Google Cloud project ID, the `[CONFIG_NAME]` must be a valid RuntimeConfig resource, the `[WAITER_NAME]` must match RFC 1035 segment specification, and the length of `[WAITER_NAME]` must be less than 64 bytes. After you create a Waiter resource, you cannot change the resource name.
	Name string `pulumi:"name"`
	// [Required] The success condition. If this condition is met, `done` will be set to `true` and the `error` value will remain unset. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated.
	Success EndConditionResponse `pulumi:"success"`
	// [Required] Specifies the timeout of the waiter in seconds, beginning from the instant that `waiters().create` method is called. If this time elapses before the success or failure conditions are met, the waiter fails and sets the `error` code to `DEADLINE_EXCEEDED`.
	Timeout string `pulumi:"timeout"`
}

func LookupWaiterOutput(ctx *pulumi.Context, args LookupWaiterOutputArgs, opts ...pulumi.InvokeOption) LookupWaiterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWaiterResult, error) {
			args := v.(LookupWaiterArgs)
			r, err := LookupWaiter(ctx, &args, opts...)
			return *r, err
		}).(LookupWaiterResultOutput)
}

type LookupWaiterOutputArgs struct {
	ConfigId pulumi.StringInput    `pulumi:"configId"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	WaiterId pulumi.StringInput    `pulumi:"waiterId"`
}

func (LookupWaiterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWaiterArgs)(nil)).Elem()
}

type LookupWaiterResultOutput struct{ *pulumi.OutputState }

func (LookupWaiterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWaiterResult)(nil)).Elem()
}

func (o LookupWaiterResultOutput) ToLookupWaiterResultOutput() LookupWaiterResultOutput {
	return o
}

func (o LookupWaiterResultOutput) ToLookupWaiterResultOutputWithContext(ctx context.Context) LookupWaiterResultOutput {
	return o
}

// The instant at which this Waiter resource was created. Adding the value of `timeout` to this instant yields the timeout deadline for the waiter.
func (o LookupWaiterResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWaiterResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// If the value is `false`, it means the waiter is still waiting for one of its conditions to be met. If true, the waiter has finished. If the waiter finished due to a timeout or failure, `error` will be set.
func (o LookupWaiterResultOutput) Done() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWaiterResult) bool { return v.Done }).(pulumi.BoolOutput)
}

// If the waiter ended due to a failure or timeout, this value will be set.
func (o LookupWaiterResultOutput) Error() StatusResponseOutput {
	return o.ApplyT(func(v LookupWaiterResult) StatusResponse { return v.Error }).(StatusResponseOutput)
}

// [Optional] The failure condition of this waiter. If this condition is met, `done` will be set to `true` and the `error` code will be set to `ABORTED`. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated. This value is optional; if no failure condition is set, the only failure scenario will be a timeout.
func (o LookupWaiterResultOutput) Failure() EndConditionResponseOutput {
	return o.ApplyT(func(v LookupWaiterResult) EndConditionResponse { return v.Failure }).(EndConditionResponseOutput)
}

// The name of the Waiter resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/waiters/[WAITER_NAME] The `[PROJECT_ID]` must be a valid Google Cloud project ID, the `[CONFIG_NAME]` must be a valid RuntimeConfig resource, the `[WAITER_NAME]` must match RFC 1035 segment specification, and the length of `[WAITER_NAME]` must be less than 64 bytes. After you create a Waiter resource, you cannot change the resource name.
func (o LookupWaiterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWaiterResult) string { return v.Name }).(pulumi.StringOutput)
}

// [Required] The success condition. If this condition is met, `done` will be set to `true` and the `error` value will remain unset. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated.
func (o LookupWaiterResultOutput) Success() EndConditionResponseOutput {
	return o.ApplyT(func(v LookupWaiterResult) EndConditionResponse { return v.Success }).(EndConditionResponseOutput)
}

// [Required] Specifies the timeout of the waiter in seconds, beginning from the instant that `waiters().create` method is called. If this time elapses before the success or failure conditions are met, the waiter fails and sets the `error` code to `DEADLINE_EXCEEDED`.
func (o LookupWaiterResultOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWaiterResult) string { return v.Timeout }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWaiterResultOutput{})
}
