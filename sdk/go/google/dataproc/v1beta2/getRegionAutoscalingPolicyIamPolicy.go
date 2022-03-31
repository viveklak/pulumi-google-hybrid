// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the access control policy for a resource. Returns an empty policy if the resource exists and does not have a policy set.
func LookupRegionAutoscalingPolicyIamPolicy(ctx *pulumi.Context, args *LookupRegionAutoscalingPolicyIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupRegionAutoscalingPolicyIamPolicyResult, error) {
	var rv LookupRegionAutoscalingPolicyIamPolicyResult
	err := ctx.Invoke("google-native:dataproc/v1beta2:getRegionAutoscalingPolicyIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionAutoscalingPolicyIamPolicyArgs struct {
	AutoscalingPolicyId           string  `pulumi:"autoscalingPolicyId"`
	OptionsRequestedPolicyVersion *string `pulumi:"optionsRequestedPolicyVersion"`
	Project                       *string `pulumi:"project"`
	RegionId                      string  `pulumi:"regionId"`
}

type LookupRegionAutoscalingPolicyIamPolicyResult struct {
	// Associates a list of members to a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one member.
	Bindings []BindingResponse `pulumi:"bindings"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag string `pulumi:"etag"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version int `pulumi:"version"`
}

func LookupRegionAutoscalingPolicyIamPolicyOutput(ctx *pulumi.Context, args LookupRegionAutoscalingPolicyIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupRegionAutoscalingPolicyIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegionAutoscalingPolicyIamPolicyResult, error) {
			args := v.(LookupRegionAutoscalingPolicyIamPolicyArgs)
			r, err := LookupRegionAutoscalingPolicyIamPolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupRegionAutoscalingPolicyIamPolicyResultOutput)
}

type LookupRegionAutoscalingPolicyIamPolicyOutputArgs struct {
	AutoscalingPolicyId           pulumi.StringInput    `pulumi:"autoscalingPolicyId"`
	OptionsRequestedPolicyVersion pulumi.StringPtrInput `pulumi:"optionsRequestedPolicyVersion"`
	Project                       pulumi.StringPtrInput `pulumi:"project"`
	RegionId                      pulumi.StringInput    `pulumi:"regionId"`
}

func (LookupRegionAutoscalingPolicyIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionAutoscalingPolicyIamPolicyArgs)(nil)).Elem()
}

type LookupRegionAutoscalingPolicyIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupRegionAutoscalingPolicyIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionAutoscalingPolicyIamPolicyResult)(nil)).Elem()
}

func (o LookupRegionAutoscalingPolicyIamPolicyResultOutput) ToLookupRegionAutoscalingPolicyIamPolicyResultOutput() LookupRegionAutoscalingPolicyIamPolicyResultOutput {
	return o
}

func (o LookupRegionAutoscalingPolicyIamPolicyResultOutput) ToLookupRegionAutoscalingPolicyIamPolicyResultOutputWithContext(ctx context.Context) LookupRegionAutoscalingPolicyIamPolicyResultOutput {
	return o
}

// Associates a list of members to a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one member.
func (o LookupRegionAutoscalingPolicyIamPolicyResultOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v LookupRegionAutoscalingPolicyIamPolicyResult) []BindingResponse { return v.Bindings }).(BindingResponseArrayOutput)
}

// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
func (o LookupRegionAutoscalingPolicyIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionAutoscalingPolicyIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
func (o LookupRegionAutoscalingPolicyIamPolicyResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRegionAutoscalingPolicyIamPolicyResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegionAutoscalingPolicyIamPolicyResultOutput{})
}
