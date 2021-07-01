// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type DatasetConsentStoreIamPolicy struct {
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

// NewDatasetConsentStoreIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatasetConsentStoreIamPolicy(ctx *pulumi.Context,
	name string, args *DatasetConsentStoreIamPolicyArgs, opts ...pulumi.ResourceOption) (*DatasetConsentStoreIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsentStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConsentStoreId'")
	}
	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource DatasetConsentStoreIamPolicy
	err := ctx.RegisterResource("google-native:healthcare/v1:DatasetConsentStoreIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetConsentStoreIamPolicy gets an existing DatasetConsentStoreIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetConsentStoreIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetConsentStoreIamPolicyState, opts ...pulumi.ResourceOption) (*DatasetConsentStoreIamPolicy, error) {
	var resource DatasetConsentStoreIamPolicy
	err := ctx.ReadResource("google-native:healthcare/v1:DatasetConsentStoreIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetConsentStoreIamPolicy resources.
type datasetConsentStoreIamPolicyState struct {
}

type DatasetConsentStoreIamPolicyState struct {
}

func (DatasetConsentStoreIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetConsentStoreIamPolicyState)(nil)).Elem()
}

type datasetConsentStoreIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings       []Binding `pulumi:"bindings"`
	ConsentStoreId string    `pulumi:"consentStoreId"`
	DatasetId      string    `pulumi:"datasetId"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     *string `pulumi:"etag"`
	Location string  `pulumi:"location"`
	Project  string  `pulumi:"project"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a DatasetConsentStoreIamPolicy resource.
type DatasetConsentStoreIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings       BindingArrayInput
	ConsentStoreId pulumi.StringInput
	DatasetId      pulumi.StringInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringPtrInput
	Location pulumi.StringInput
	Project  pulumi.StringInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (DatasetConsentStoreIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetConsentStoreIamPolicyArgs)(nil)).Elem()
}

type DatasetConsentStoreIamPolicyInput interface {
	pulumi.Input

	ToDatasetConsentStoreIamPolicyOutput() DatasetConsentStoreIamPolicyOutput
	ToDatasetConsentStoreIamPolicyOutputWithContext(ctx context.Context) DatasetConsentStoreIamPolicyOutput
}

func (*DatasetConsentStoreIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetConsentStoreIamPolicy)(nil))
}

func (i *DatasetConsentStoreIamPolicy) ToDatasetConsentStoreIamPolicyOutput() DatasetConsentStoreIamPolicyOutput {
	return i.ToDatasetConsentStoreIamPolicyOutputWithContext(context.Background())
}

func (i *DatasetConsentStoreIamPolicy) ToDatasetConsentStoreIamPolicyOutputWithContext(ctx context.Context) DatasetConsentStoreIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetConsentStoreIamPolicyOutput)
}

type DatasetConsentStoreIamPolicyOutput struct {
	*pulumi.OutputState
}

func (DatasetConsentStoreIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetConsentStoreIamPolicy)(nil))
}

func (o DatasetConsentStoreIamPolicyOutput) ToDatasetConsentStoreIamPolicyOutput() DatasetConsentStoreIamPolicyOutput {
	return o
}

func (o DatasetConsentStoreIamPolicyOutput) ToDatasetConsentStoreIamPolicyOutputWithContext(ctx context.Context) DatasetConsentStoreIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetConsentStoreIamPolicyOutput{})
}
