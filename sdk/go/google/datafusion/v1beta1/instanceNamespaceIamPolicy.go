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
type InstanceNamespaceIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewInstanceNamespaceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewInstanceNamespaceIamPolicy(ctx *pulumi.Context,
	name string, args *InstanceNamespaceIamPolicyArgs, opts ...pulumi.ResourceOption) (*InstanceNamespaceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource InstanceNamespaceIamPolicy
	err := ctx.RegisterResource("google-native:datafusion/v1beta1:InstanceNamespaceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceNamespaceIamPolicy gets an existing InstanceNamespaceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceNamespaceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceNamespaceIamPolicyState, opts ...pulumi.ResourceOption) (*InstanceNamespaceIamPolicy, error) {
	var resource InstanceNamespaceIamPolicy
	err := ctx.ReadResource("google-native:datafusion/v1beta1:InstanceNamespaceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceNamespaceIamPolicy resources.
type instanceNamespaceIamPolicyState struct {
}

type InstanceNamespaceIamPolicyState struct {
}

func (InstanceNamespaceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceNamespaceIamPolicyState)(nil)).Elem()
}

type instanceNamespaceIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag        *string `pulumi:"etag"`
	InstanceId  string  `pulumi:"instanceId"`
	Location    string  `pulumi:"location"`
	NamespaceId string  `pulumi:"namespaceId"`
	Project     string  `pulumi:"project"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a InstanceNamespaceIamPolicy resource.
type InstanceNamespaceIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag        pulumi.StringPtrInput
	InstanceId  pulumi.StringInput
	Location    pulumi.StringInput
	NamespaceId pulumi.StringInput
	Project     pulumi.StringInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (InstanceNamespaceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceNamespaceIamPolicyArgs)(nil)).Elem()
}

type InstanceNamespaceIamPolicyInput interface {
	pulumi.Input

	ToInstanceNamespaceIamPolicyOutput() InstanceNamespaceIamPolicyOutput
	ToInstanceNamespaceIamPolicyOutputWithContext(ctx context.Context) InstanceNamespaceIamPolicyOutput
}

func (*InstanceNamespaceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNamespaceIamPolicy)(nil))
}

func (i *InstanceNamespaceIamPolicy) ToInstanceNamespaceIamPolicyOutput() InstanceNamespaceIamPolicyOutput {
	return i.ToInstanceNamespaceIamPolicyOutputWithContext(context.Background())
}

func (i *InstanceNamespaceIamPolicy) ToInstanceNamespaceIamPolicyOutputWithContext(ctx context.Context) InstanceNamespaceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNamespaceIamPolicyOutput)
}

type InstanceNamespaceIamPolicyOutput struct {
	*pulumi.OutputState
}

func (InstanceNamespaceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNamespaceIamPolicy)(nil))
}

func (o InstanceNamespaceIamPolicyOutput) ToInstanceNamespaceIamPolicyOutput() InstanceNamespaceIamPolicyOutput {
	return o
}

func (o InstanceNamespaceIamPolicyOutput) ToInstanceNamespaceIamPolicyOutputWithContext(ctx context.Context) InstanceNamespaceIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceNamespaceIamPolicyOutput{})
}
