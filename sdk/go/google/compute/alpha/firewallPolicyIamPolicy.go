// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.
type FirewallPolicyIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy.
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringOutput `pulumi:"etag"`
	IamOwned pulumi.BoolOutput   `pulumi:"iamOwned"`
	// If more than one rule is specified, the rules are applied in the following manner: - All matching LOG rules are always applied. - If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging will be applied if one or more matching rule requires logging. - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is granted. Logging will be applied if one or more matching rule requires logging. - Otherwise, if no rule applies, permission is denied.
	Rules RuleResponseArrayOutput `pulumi:"rules"`
	// Specifies the format of the policy.
	//
	// Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected.
	//
	// Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations:
	//
	// * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	//
	// If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.
	//
	// To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewFirewallPolicyIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicyIamPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyIamPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicyIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	var resource FirewallPolicyIamPolicy
	err := ctx.RegisterResource("google-native:compute/alpha:FirewallPolicyIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicyIamPolicy gets an existing FirewallPolicyIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicyIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyIamPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicyIamPolicy, error) {
	var resource FirewallPolicyIamPolicy
	err := ctx.ReadResource("google-native:compute/alpha:FirewallPolicyIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicyIamPolicy resources.
type firewallPolicyIamPolicyState struct {
}

type FirewallPolicyIamPolicyState struct {
}

func (FirewallPolicyIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyIamPolicyState)(nil)).Elem()
}

type firewallPolicyIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy.
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     *string `pulumi:"etag"`
	IamOwned *bool   `pulumi:"iamOwned"`
	Resource string  `pulumi:"resource"`
	// If more than one rule is specified, the rules are applied in the following manner: - All matching LOG rules are always applied. - If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging will be applied if one or more matching rule requires logging. - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is granted. Logging will be applied if one or more matching rule requires logging. - Otherwise, if no rule applies, permission is denied.
	Rules []Rule `pulumi:"rules"`
	// Specifies the format of the policy.
	//
	// Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected.
	//
	// Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations:
	//
	// * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	//
	// If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.
	//
	// To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a FirewallPolicyIamPolicy resource.
type FirewallPolicyIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy.
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringPtrInput
	IamOwned pulumi.BoolPtrInput
	Resource pulumi.StringInput
	// If more than one rule is specified, the rules are applied in the following manner: - All matching LOG rules are always applied. - If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging will be applied if one or more matching rule requires logging. - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is granted. Logging will be applied if one or more matching rule requires logging. - Otherwise, if no rule applies, permission is denied.
	Rules RuleArrayInput
	// Specifies the format of the policy.
	//
	// Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected.
	//
	// Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations:
	//
	// * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	//
	// If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.
	//
	// To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (FirewallPolicyIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyIamPolicyArgs)(nil)).Elem()
}

type FirewallPolicyIamPolicyInput interface {
	pulumi.Input

	ToFirewallPolicyIamPolicyOutput() FirewallPolicyIamPolicyOutput
	ToFirewallPolicyIamPolicyOutputWithContext(ctx context.Context) FirewallPolicyIamPolicyOutput
}

func (*FirewallPolicyIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyIamPolicy)(nil))
}

func (i *FirewallPolicyIamPolicy) ToFirewallPolicyIamPolicyOutput() FirewallPolicyIamPolicyOutput {
	return i.ToFirewallPolicyIamPolicyOutputWithContext(context.Background())
}

func (i *FirewallPolicyIamPolicy) ToFirewallPolicyIamPolicyOutputWithContext(ctx context.Context) FirewallPolicyIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyIamPolicyOutput)
}

type FirewallPolicyIamPolicyOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyIamPolicy)(nil))
}

func (o FirewallPolicyIamPolicyOutput) ToFirewallPolicyIamPolicyOutput() FirewallPolicyIamPolicyOutput {
	return o
}

func (o FirewallPolicyIamPolicyOutput) ToFirewallPolicyIamPolicyOutputWithContext(ctx context.Context) FirewallPolicyIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyIamPolicyOutput{})
}
