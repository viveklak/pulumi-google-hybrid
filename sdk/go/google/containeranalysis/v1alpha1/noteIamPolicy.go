// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified `Note` or `Occurrence`. Requires `containeranalysis.notes.setIamPolicy` or `containeranalysis.occurrences.setIamPolicy` permission if the resource is a `Note` or an `Occurrence`, respectively. Attempting to call this method without these permissions will result in a ` `PERMISSION_DENIED`error. Attempting to call this method on a non-existent resource will result in a`NOT_FOUND`error if the user has`containeranalysis.notes.list`permission on a`Note`or`containeranalysis.occurrences.list`on an`Occurrence` , or a  `PERMISSION_DENIED`error otherwise. The resource takes the following formats:`projects/{projectid}/occurrences/{occurrenceid}` for occurrences and projects/{projectid}/notes/{noteid} for notes
type NoteIamPolicy struct {
	pulumi.CustomResourceState

	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewNoteIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewNoteIamPolicy(ctx *pulumi.Context,
	name string, args *NoteIamPolicyArgs, opts ...pulumi.ResourceOption) (*NoteIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NoteId == nil {
		return nil, errors.New("invalid value for required argument 'NoteId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource NoteIamPolicy
	err := ctx.RegisterResource("google-native:containeranalysis/v1alpha1:NoteIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNoteIamPolicy gets an existing NoteIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNoteIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NoteIamPolicyState, opts ...pulumi.ResourceOption) (*NoteIamPolicy, error) {
	var resource NoteIamPolicy
	err := ctx.ReadResource("google-native:containeranalysis/v1alpha1:NoteIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NoteIamPolicy resources.
type noteIamPolicyState struct {
}

type NoteIamPolicyState struct {
}

func (NoteIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*noteIamPolicyState)(nil)).Elem()
}

type noteIamPolicyArgs struct {
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag    *string `pulumi:"etag"`
	NoteId  string  `pulumi:"noteId"`
	Project string  `pulumi:"project"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a NoteIamPolicy resource.
type NoteIamPolicyArgs struct {
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag    pulumi.StringPtrInput
	NoteId  pulumi.StringInput
	Project pulumi.StringInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (NoteIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*noteIamPolicyArgs)(nil)).Elem()
}

type NoteIamPolicyInput interface {
	pulumi.Input

	ToNoteIamPolicyOutput() NoteIamPolicyOutput
	ToNoteIamPolicyOutputWithContext(ctx context.Context) NoteIamPolicyOutput
}

func (*NoteIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*NoteIamPolicy)(nil))
}

func (i *NoteIamPolicy) ToNoteIamPolicyOutput() NoteIamPolicyOutput {
	return i.ToNoteIamPolicyOutputWithContext(context.Background())
}

func (i *NoteIamPolicy) ToNoteIamPolicyOutputWithContext(ctx context.Context) NoteIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteIamPolicyOutput)
}

type NoteIamPolicyOutput struct {
	*pulumi.OutputState
}

func (NoteIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NoteIamPolicy)(nil))
}

func (o NoteIamPolicyOutput) ToNoteIamPolicyOutput() NoteIamPolicyOutput {
	return o
}

func (o NoteIamPolicyOutput) ToNoteIamPolicyOutputWithContext(ctx context.Context) NoteIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NoteIamPolicyOutput{})
}
