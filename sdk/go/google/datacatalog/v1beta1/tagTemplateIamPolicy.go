// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy for a resource. Replaces any existing policy. Supported resources are: - Tag templates. - Entries. - Entry groups. Note, this method cannot be used to manage policies for BigQuery, Pub/Sub and any external Google Cloud Platform resources synced to Data Catalog. Callers must have following Google IAM permission - `datacatalog.tagTemplates.setIamPolicy` to set policies on tag templates. - `datacatalog.entries.setIamPolicy` to set policies on entries. - `datacatalog.entryGroups.setIamPolicy` to set policies on entry groups.
type TagTemplateIamPolicy struct {
	pulumi.CustomResourceState

	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewTagTemplateIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewTagTemplateIamPolicy(ctx *pulumi.Context,
	name string, args *TagTemplateIamPolicyArgs, opts ...pulumi.ResourceOption) (*TagTemplateIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.TagTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TagTemplateId'")
	}
	var resource TagTemplateIamPolicy
	err := ctx.RegisterResource("google-native:datacatalog/v1beta1:TagTemplateIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagTemplateIamPolicy gets an existing TagTemplateIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagTemplateIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagTemplateIamPolicyState, opts ...pulumi.ResourceOption) (*TagTemplateIamPolicy, error) {
	var resource TagTemplateIamPolicy
	err := ctx.ReadResource("google-native:datacatalog/v1beta1:TagTemplateIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagTemplateIamPolicy resources.
type tagTemplateIamPolicyState struct {
}

type TagTemplateIamPolicyState struct {
}

func (TagTemplateIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateIamPolicyState)(nil)).Elem()
}

type tagTemplateIamPolicyArgs struct {
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag          *string `pulumi:"etag"`
	Location      string  `pulumi:"location"`
	Project       string  `pulumi:"project"`
	TagTemplateId string  `pulumi:"tagTemplateId"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a TagTemplateIamPolicy resource.
type TagTemplateIamPolicyArgs struct {
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag          pulumi.StringPtrInput
	Location      pulumi.StringInput
	Project       pulumi.StringInput
	TagTemplateId pulumi.StringInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (TagTemplateIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateIamPolicyArgs)(nil)).Elem()
}

type TagTemplateIamPolicyInput interface {
	pulumi.Input

	ToTagTemplateIamPolicyOutput() TagTemplateIamPolicyOutput
	ToTagTemplateIamPolicyOutputWithContext(ctx context.Context) TagTemplateIamPolicyOutput
}

func (*TagTemplateIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*TagTemplateIamPolicy)(nil))
}

func (i *TagTemplateIamPolicy) ToTagTemplateIamPolicyOutput() TagTemplateIamPolicyOutput {
	return i.ToTagTemplateIamPolicyOutputWithContext(context.Background())
}

func (i *TagTemplateIamPolicy) ToTagTemplateIamPolicyOutputWithContext(ctx context.Context) TagTemplateIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamPolicyOutput)
}

type TagTemplateIamPolicyOutput struct {
	*pulumi.OutputState
}

func (TagTemplateIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagTemplateIamPolicy)(nil))
}

func (o TagTemplateIamPolicyOutput) ToTagTemplateIamPolicyOutput() TagTemplateIamPolicyOutput {
	return o
}

func (o TagTemplateIamPolicyOutput) ToTagTemplateIamPolicyOutputWithContext(ctx context.Context) TagTemplateIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagTemplateIamPolicyOutput{})
}
