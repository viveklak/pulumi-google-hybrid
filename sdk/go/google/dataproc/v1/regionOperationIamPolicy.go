// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED errors.
type RegionOperationIamPolicy struct {
	pulumi.CustomResourceState

	// Associates a list of members to a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewRegionOperationIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewRegionOperationIamPolicy(ctx *pulumi.Context,
	name string, args *RegionOperationIamPolicyArgs, opts ...pulumi.ResourceOption) (*RegionOperationIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OperationId == nil {
		return nil, errors.New("invalid value for required argument 'OperationId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.RegionId == nil {
		return nil, errors.New("invalid value for required argument 'RegionId'")
	}
	var resource RegionOperationIamPolicy
	err := ctx.RegisterResource("google-native:dataproc/v1:RegionOperationIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionOperationIamPolicy gets an existing RegionOperationIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionOperationIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionOperationIamPolicyState, opts ...pulumi.ResourceOption) (*RegionOperationIamPolicy, error) {
	var resource RegionOperationIamPolicy
	err := ctx.ReadResource("google-native:dataproc/v1:RegionOperationIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionOperationIamPolicy resources.
type regionOperationIamPolicyState struct {
}

type RegionOperationIamPolicyState struct {
}

func (RegionOperationIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionOperationIamPolicyState)(nil)).Elem()
}

type regionOperationIamPolicyArgs struct {
	// Associates a list of members to a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one member.
	Bindings []Binding `pulumi:"bindings"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag        *string `pulumi:"etag"`
	OperationId string  `pulumi:"operationId"`
	Project     string  `pulumi:"project"`
	RegionId    string  `pulumi:"regionId"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a RegionOperationIamPolicy resource.
type RegionOperationIamPolicyArgs struct {
	// Associates a list of members to a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one member.
	Bindings BindingArrayInput
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag        pulumi.StringPtrInput
	OperationId pulumi.StringInput
	Project     pulumi.StringInput
	RegionId    pulumi.StringInput
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (RegionOperationIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionOperationIamPolicyArgs)(nil)).Elem()
}

type RegionOperationIamPolicyInput interface {
	pulumi.Input

	ToRegionOperationIamPolicyOutput() RegionOperationIamPolicyOutput
	ToRegionOperationIamPolicyOutputWithContext(ctx context.Context) RegionOperationIamPolicyOutput
}

func (*RegionOperationIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionOperationIamPolicy)(nil))
}

func (i *RegionOperationIamPolicy) ToRegionOperationIamPolicyOutput() RegionOperationIamPolicyOutput {
	return i.ToRegionOperationIamPolicyOutputWithContext(context.Background())
}

func (i *RegionOperationIamPolicy) ToRegionOperationIamPolicyOutputWithContext(ctx context.Context) RegionOperationIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionOperationIamPolicyOutput)
}

type RegionOperationIamPolicyOutput struct {
	*pulumi.OutputState
}

func (RegionOperationIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionOperationIamPolicy)(nil))
}

func (o RegionOperationIamPolicyOutput) ToRegionOperationIamPolicyOutput() RegionOperationIamPolicyOutput {
	return o
}

func (o RegionOperationIamPolicyOutput) ToRegionOperationIamPolicyOutputWithContext(ctx context.Context) RegionOperationIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionOperationIamPolicyOutput{})
}
