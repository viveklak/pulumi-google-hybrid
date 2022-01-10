// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type ClientTlsPolicyIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs GoogleIamV1AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings GoogleIamV1BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewClientTlsPolicyIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewClientTlsPolicyIamPolicy(ctx *pulumi.Context,
	name string, args *ClientTlsPolicyIamPolicyArgs, opts ...pulumi.ResourceOption) (*ClientTlsPolicyIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientTlsPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'ClientTlsPolicyId'")
	}
	var resource ClientTlsPolicyIamPolicy
	err := ctx.RegisterResource("google-native:networksecurity/v1beta1:ClientTlsPolicyIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientTlsPolicyIamPolicy gets an existing ClientTlsPolicyIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientTlsPolicyIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientTlsPolicyIamPolicyState, opts ...pulumi.ResourceOption) (*ClientTlsPolicyIamPolicy, error) {
	var resource ClientTlsPolicyIamPolicy
	err := ctx.ReadResource("google-native:networksecurity/v1beta1:ClientTlsPolicyIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientTlsPolicyIamPolicy resources.
type clientTlsPolicyIamPolicyState struct {
}

type ClientTlsPolicyIamPolicyState struct {
}

func (ClientTlsPolicyIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientTlsPolicyIamPolicyState)(nil)).Elem()
}

type clientTlsPolicyIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []GoogleIamV1AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings          []GoogleIamV1Binding `pulumi:"bindings"`
	ClientTlsPolicyId string               `pulumi:"clientTlsPolicyId"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     *string `pulumi:"etag"`
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a ClientTlsPolicyIamPolicy resource.
type ClientTlsPolicyIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs GoogleIamV1AuditConfigArrayInput
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings          GoogleIamV1BindingArrayInput
	ClientTlsPolicyId pulumi.StringInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (ClientTlsPolicyIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientTlsPolicyIamPolicyArgs)(nil)).Elem()
}

type ClientTlsPolicyIamPolicyInput interface {
	pulumi.Input

	ToClientTlsPolicyIamPolicyOutput() ClientTlsPolicyIamPolicyOutput
	ToClientTlsPolicyIamPolicyOutputWithContext(ctx context.Context) ClientTlsPolicyIamPolicyOutput
}

func (*ClientTlsPolicyIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientTlsPolicyIamPolicy)(nil)).Elem()
}

func (i *ClientTlsPolicyIamPolicy) ToClientTlsPolicyIamPolicyOutput() ClientTlsPolicyIamPolicyOutput {
	return i.ToClientTlsPolicyIamPolicyOutputWithContext(context.Background())
}

func (i *ClientTlsPolicyIamPolicy) ToClientTlsPolicyIamPolicyOutputWithContext(ctx context.Context) ClientTlsPolicyIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientTlsPolicyIamPolicyOutput)
}

type ClientTlsPolicyIamPolicyOutput struct{ *pulumi.OutputState }

func (ClientTlsPolicyIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientTlsPolicyIamPolicy)(nil)).Elem()
}

func (o ClientTlsPolicyIamPolicyOutput) ToClientTlsPolicyIamPolicyOutput() ClientTlsPolicyIamPolicyOutput {
	return o
}

func (o ClientTlsPolicyIamPolicyOutput) ToClientTlsPolicyIamPolicyOutputWithContext(ctx context.Context) ClientTlsPolicyIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientTlsPolicyIamPolicyInput)(nil)).Elem(), &ClientTlsPolicyIamPolicy{})
	pulumi.RegisterOutputType(ClientTlsPolicyIamPolicyOutput{})
}
