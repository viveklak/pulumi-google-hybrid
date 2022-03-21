// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type RegionBackendServiceIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// This is deprecated and has no effect. Do not use.
	Rules RuleResponseArrayOutput `pulumi:"rules"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewRegionBackendServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewRegionBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, args *RegionBackendServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*RegionBackendServiceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	var resource RegionBackendServiceIamPolicy
	err := ctx.RegisterResource("google-native:compute/beta:RegionBackendServiceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionBackendServiceIamPolicy gets an existing RegionBackendServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionBackendServiceIamPolicyState, opts ...pulumi.ResourceOption) (*RegionBackendServiceIamPolicy, error) {
	var resource RegionBackendServiceIamPolicy
	err := ctx.ReadResource("google-native:compute/beta:RegionBackendServiceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionBackendServiceIamPolicy resources.
type regionBackendServiceIamPolicyState struct {
}

type RegionBackendServiceIamPolicyState struct {
}

func (RegionBackendServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionBackendServiceIamPolicyState)(nil)).Elem()
}

type regionBackendServiceIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     *string `pulumi:"etag"`
	Project  *string `pulumi:"project"`
	Region   string  `pulumi:"region"`
	Resource string  `pulumi:"resource"`
	// This is deprecated and has no effect. Do not use.
	Rules []Rule `pulumi:"rules"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a RegionBackendServiceIamPolicy resource.
type RegionBackendServiceIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Region   pulumi.StringInput
	Resource pulumi.StringInput
	// This is deprecated and has no effect. Do not use.
	Rules RuleArrayInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (RegionBackendServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionBackendServiceIamPolicyArgs)(nil)).Elem()
}

type RegionBackendServiceIamPolicyInput interface {
	pulumi.Input

	ToRegionBackendServiceIamPolicyOutput() RegionBackendServiceIamPolicyOutput
	ToRegionBackendServiceIamPolicyOutputWithContext(ctx context.Context) RegionBackendServiceIamPolicyOutput
}

func (*RegionBackendServiceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionBackendServiceIamPolicy)(nil)).Elem()
}

func (i *RegionBackendServiceIamPolicy) ToRegionBackendServiceIamPolicyOutput() RegionBackendServiceIamPolicyOutput {
	return i.ToRegionBackendServiceIamPolicyOutputWithContext(context.Background())
}

func (i *RegionBackendServiceIamPolicy) ToRegionBackendServiceIamPolicyOutputWithContext(ctx context.Context) RegionBackendServiceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionBackendServiceIamPolicyOutput)
}

type RegionBackendServiceIamPolicyOutput struct{ *pulumi.OutputState }

func (RegionBackendServiceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionBackendServiceIamPolicy)(nil)).Elem()
}

func (o RegionBackendServiceIamPolicyOutput) ToRegionBackendServiceIamPolicyOutput() RegionBackendServiceIamPolicyOutput {
	return o
}

func (o RegionBackendServiceIamPolicyOutput) ToRegionBackendServiceIamPolicyOutputWithContext(ctx context.Context) RegionBackendServiceIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionBackendServiceIamPolicyInput)(nil)).Elem(), &RegionBackendServiceIamPolicy{})
	pulumi.RegisterOutputType(RegionBackendServiceIamPolicyOutput{})
}
