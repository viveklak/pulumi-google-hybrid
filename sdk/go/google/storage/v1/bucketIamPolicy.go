// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Updates an IAM policy for the specified bucket.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type BucketIamPolicy struct {
	pulumi.CustomResourceState

	// An association between a role, which comes with a set of permissions, and members who may assume that role.
	Bindings BucketIamPolicyBindingsItemResponseArrayOutput `pulumi:"bindings"`
	// HTTP 1.1  Entity tag for the policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The kind of item this is. For policies, this is always storage#policy. This field is ignored on input.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The ID of the resource to which this policy belongs. Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects. A specific generation may be specified by appending #generationNumber to the end of the object name, e.g. projects/_/buckets/my-bucket/objects/data.txt#17. The current generation can be denoted with #0. This field is ignored on input.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The IAM policy format version.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewBucketIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewBucketIamPolicy(ctx *pulumi.Context,
	name string, args *BucketIamPolicyArgs, opts ...pulumi.ResourceOption) (*BucketIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	var resource BucketIamPolicy
	err := ctx.RegisterResource("google-native:storage/v1:BucketIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketIamPolicy gets an existing BucketIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketIamPolicyState, opts ...pulumi.ResourceOption) (*BucketIamPolicy, error) {
	var resource BucketIamPolicy
	err := ctx.ReadResource("google-native:storage/v1:BucketIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketIamPolicy resources.
type bucketIamPolicyState struct {
}

type BucketIamPolicyState struct {
}

func (BucketIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIamPolicyState)(nil)).Elem()
}

type bucketIamPolicyArgs struct {
	// An association between a role, which comes with a set of permissions, and members who may assume that role.
	Bindings []BucketIamPolicyBindingsItem `pulumi:"bindings"`
	Bucket   string                        `pulumi:"bucket"`
	// HTTP 1.1  Entity tag for the policy.
	Etag *string `pulumi:"etag"`
	// The kind of item this is. For policies, this is always storage#policy. This field is ignored on input.
	Kind                   *string `pulumi:"kind"`
	ProvisionalUserProject *string `pulumi:"provisionalUserProject"`
	// The ID of the resource to which this policy belongs. Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects. A specific generation may be specified by appending #generationNumber to the end of the object name, e.g. projects/_/buckets/my-bucket/objects/data.txt#17. The current generation can be denoted with #0. This field is ignored on input.
	ResourceId  *string `pulumi:"resourceId"`
	UserProject *string `pulumi:"userProject"`
	// The IAM policy format version.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a BucketIamPolicy resource.
type BucketIamPolicyArgs struct {
	// An association between a role, which comes with a set of permissions, and members who may assume that role.
	Bindings BucketIamPolicyBindingsItemArrayInput
	Bucket   pulumi.StringInput
	// HTTP 1.1  Entity tag for the policy.
	Etag pulumi.StringPtrInput
	// The kind of item this is. For policies, this is always storage#policy. This field is ignored on input.
	Kind                   pulumi.StringPtrInput
	ProvisionalUserProject pulumi.StringPtrInput
	// The ID of the resource to which this policy belongs. Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects. A specific generation may be specified by appending #generationNumber to the end of the object name, e.g. projects/_/buckets/my-bucket/objects/data.txt#17. The current generation can be denoted with #0. This field is ignored on input.
	ResourceId  pulumi.StringPtrInput
	UserProject pulumi.StringPtrInput
	// The IAM policy format version.
	Version pulumi.IntPtrInput
}

func (BucketIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIamPolicyArgs)(nil)).Elem()
}

type BucketIamPolicyInput interface {
	pulumi.Input

	ToBucketIamPolicyOutput() BucketIamPolicyOutput
	ToBucketIamPolicyOutputWithContext(ctx context.Context) BucketIamPolicyOutput
}

func (*BucketIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketIamPolicy)(nil)).Elem()
}

func (i *BucketIamPolicy) ToBucketIamPolicyOutput() BucketIamPolicyOutput {
	return i.ToBucketIamPolicyOutputWithContext(context.Background())
}

func (i *BucketIamPolicy) ToBucketIamPolicyOutputWithContext(ctx context.Context) BucketIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIamPolicyOutput)
}

type BucketIamPolicyOutput struct{ *pulumi.OutputState }

func (BucketIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketIamPolicy)(nil)).Elem()
}

func (o BucketIamPolicyOutput) ToBucketIamPolicyOutput() BucketIamPolicyOutput {
	return o
}

func (o BucketIamPolicyOutput) ToBucketIamPolicyOutputWithContext(ctx context.Context) BucketIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketIamPolicyInput)(nil)).Elem(), &BucketIamPolicy{})
	pulumi.RegisterOutputType(BucketIamPolicyOutput{})
}
