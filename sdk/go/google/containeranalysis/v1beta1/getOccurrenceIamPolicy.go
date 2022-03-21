// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the access control policy for a note or an occurrence resource. Requires `containeranalysis.notes.setIamPolicy` or `containeranalysis.occurrences.setIamPolicy` permission if the resource is a note or occurrence, respectively. The resource takes the format `projects/[PROJECT_ID]/notes/[NOTE_ID]` for notes and `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]` for occurrences.
func LookupOccurrenceIamPolicy(ctx *pulumi.Context, args *LookupOccurrenceIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupOccurrenceIamPolicyResult, error) {
	var rv LookupOccurrenceIamPolicyResult
	err := ctx.Invoke("google-native:containeranalysis/v1beta1:getOccurrenceIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOccurrenceIamPolicyArgs struct {
	OccurrenceId string  `pulumi:"occurrenceId"`
	Project      *string `pulumi:"project"`
}

type LookupOccurrenceIamPolicyResult struct {
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []BindingResponse `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag string `pulumi:"etag"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version int `pulumi:"version"`
}

func LookupOccurrenceIamPolicyOutput(ctx *pulumi.Context, args LookupOccurrenceIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupOccurrenceIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOccurrenceIamPolicyResult, error) {
			args := v.(LookupOccurrenceIamPolicyArgs)
			r, err := LookupOccurrenceIamPolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupOccurrenceIamPolicyResultOutput)
}

type LookupOccurrenceIamPolicyOutputArgs struct {
	OccurrenceId pulumi.StringInput    `pulumi:"occurrenceId"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupOccurrenceIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOccurrenceIamPolicyArgs)(nil)).Elem()
}

type LookupOccurrenceIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupOccurrenceIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOccurrenceIamPolicyResult)(nil)).Elem()
}

func (o LookupOccurrenceIamPolicyResultOutput) ToLookupOccurrenceIamPolicyResultOutput() LookupOccurrenceIamPolicyResultOutput {
	return o
}

func (o LookupOccurrenceIamPolicyResultOutput) ToLookupOccurrenceIamPolicyResultOutputWithContext(ctx context.Context) LookupOccurrenceIamPolicyResultOutput {
	return o
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o LookupOccurrenceIamPolicyResultOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v LookupOccurrenceIamPolicyResult) []BindingResponse { return v.Bindings }).(BindingResponseArrayOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o LookupOccurrenceIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOccurrenceIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o LookupOccurrenceIamPolicyResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOccurrenceIamPolicyResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOccurrenceIamPolicyResultOutput{})
}
