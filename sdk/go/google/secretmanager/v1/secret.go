// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Secret containing no SecretVersions.
// Auto-naming is currently not supported for this resource.
type Secret struct {
	pulumi.CustomResourceState

	// The time at which the Secret was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Etag of the currently stored Secret.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Optional. Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// The labels assigned to this Secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `\p{Ll}\p{Lo}{0,62}` Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}` No more than 64 labels can be assigned to a given resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name of the Secret in the format `projects/*/secrets/*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The replication policy of the secret data attached to the Secret. The replication policy cannot be changed after the Secret has been created.
	Replication ReplicationResponseOutput `pulumi:"replication"`
	// Optional. Rotation policy attached to the Secret. May be excluded if there is no rotation policy.
	Rotation RotationResponseOutput `pulumi:"rotation"`
	// Optional. A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
	Topics TopicResponseArrayOutput `pulumi:"topics"`
	// Input only. The TTL for the Secret.
	Ttl pulumi.StringOutput `pulumi:"ttl"`
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Replication == nil {
		return nil, errors.New("invalid value for required argument 'Replication'")
	}
	if args.SecretId == nil {
		return nil, errors.New("invalid value for required argument 'SecretId'")
	}
	var resource Secret
	err := ctx.RegisterResource("google-native:secretmanager/v1:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("google-native:secretmanager/v1:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secret resources.
type secretState struct {
}

type SecretState struct {
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	// Optional. Etag of the currently stored Secret.
	Etag *string `pulumi:"etag"`
	// Optional. Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
	ExpireTime *string `pulumi:"expireTime"`
	// The labels assigned to this Secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `\p{Ll}\p{Lo}{0,62}` Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}` No more than 64 labels can be assigned to a given resource.
	Labels  map[string]string `pulumi:"labels"`
	Project *string           `pulumi:"project"`
	// Immutable. The replication policy of the secret data attached to the Secret. The replication policy cannot be changed after the Secret has been created.
	Replication Replication `pulumi:"replication"`
	// Optional. Rotation policy attached to the Secret. May be excluded if there is no rotation policy.
	Rotation *Rotation `pulumi:"rotation"`
	SecretId string    `pulumi:"secretId"`
	// Optional. A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
	Topics []Topic `pulumi:"topics"`
	// Input only. The TTL for the Secret.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	// Optional. Etag of the currently stored Secret.
	Etag pulumi.StringPtrInput
	// Optional. Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
	ExpireTime pulumi.StringPtrInput
	// The labels assigned to this Secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `\p{Ll}\p{Lo}{0,62}` Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}` No more than 64 labels can be assigned to a given resource.
	Labels  pulumi.StringMapInput
	Project pulumi.StringPtrInput
	// Immutable. The replication policy of the secret data attached to the Secret. The replication policy cannot be changed after the Secret has been created.
	Replication ReplicationInput
	// Optional. Rotation policy attached to the Secret. May be excluded if there is no rotation policy.
	Rotation RotationPtrInput
	SecretId pulumi.StringInput
	// Optional. A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
	Topics TopicArrayInput
	// Input only. The TTL for the Secret.
	Ttl pulumi.StringPtrInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}

type SecretInput interface {
	pulumi.Input

	ToSecretOutput() SecretOutput
	ToSecretOutputWithContext(ctx context.Context) SecretOutput
}

func (*Secret) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (i *Secret) ToSecretOutput() SecretOutput {
	return i.ToSecretOutputWithContext(context.Background())
}

func (i *Secret) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretOutput)
}

type SecretOutput struct{ *pulumi.OutputState }

func (SecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (o SecretOutput) ToSecretOutput() SecretOutput {
	return o
}

func (o SecretOutput) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretInput)(nil)).Elem(), &Secret{})
	pulumi.RegisterOutputType(SecretOutput{})
}
