// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new `BitbucketServerConfig`. This API is experimental.
type BitbucketServerConfig struct {
	pulumi.CustomResourceState

	// Immutable. API Key that will be attached to webhook. Once this field has been set, it cannot be changed. If you need to change it, please create another BitbucketServerConfig.
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// Connected Bitbucket Server repositories for this config.
	ConnectedRepositories BitbucketServerRepositoryIdResponseArrayOutput `pulumi:"connectedRepositories"`
	// Time when the config was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Immutable. The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed. If you need to change it, please create another BitbucketServerConfig.
	HostUri pulumi.StringOutput `pulumi:"hostUri"`
	// The resource name for the config.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The network to be used when reaching out to the Bitbucket Server instance. The VPC network must be enabled for private service connection. This should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty, no network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format `projects/{project}/global/networks/{network}`, where {project} is a project number or id and {network} is the name of a VPC network in the project.
	PeeredNetwork pulumi.StringOutput `pulumi:"peeredNetwork"`
	// Secret Manager secrets needed by the config.
	Secrets BitbucketServerSecretsResponseOutput `pulumi:"secrets"`
	// Optional. SSL certificate to use for requests to Bitbucket Server. The format should be PEM format but the extension can be one of .pem, .cer, or .crt.
	SslCa pulumi.StringOutput `pulumi:"sslCa"`
	// Username of the account Cloud Build will use on Bitbucket Server.
	Username pulumi.StringOutput `pulumi:"username"`
	// UUID included in webhook requests. The UUID is used to look up the corresponding config.
	WebhookKey pulumi.StringOutput `pulumi:"webhookKey"`
}

// NewBitbucketServerConfig registers a new resource with the given unique name, arguments, and options.
func NewBitbucketServerConfig(ctx *pulumi.Context,
	name string, args *BitbucketServerConfigArgs, opts ...pulumi.ResourceOption) (*BitbucketServerConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.HostUri == nil {
		return nil, errors.New("invalid value for required argument 'HostUri'")
	}
	if args.Secrets == nil {
		return nil, errors.New("invalid value for required argument 'Secrets'")
	}
	var resource BitbucketServerConfig
	err := ctx.RegisterResource("google-native:cloudbuild/v1:BitbucketServerConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBitbucketServerConfig gets an existing BitbucketServerConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBitbucketServerConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BitbucketServerConfigState, opts ...pulumi.ResourceOption) (*BitbucketServerConfig, error) {
	var resource BitbucketServerConfig
	err := ctx.ReadResource("google-native:cloudbuild/v1:BitbucketServerConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BitbucketServerConfig resources.
type bitbucketServerConfigState struct {
}

type BitbucketServerConfigState struct {
}

func (BitbucketServerConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*bitbucketServerConfigState)(nil)).Elem()
}

type bitbucketServerConfigArgs struct {
	// Immutable. API Key that will be attached to webhook. Once this field has been set, it cannot be changed. If you need to change it, please create another BitbucketServerConfig.
	ApiKey                  string  `pulumi:"apiKey"`
	BitbucketServerConfigId *string `pulumi:"bitbucketServerConfigId"`
	// Time when the config was created.
	CreateTime *string `pulumi:"createTime"`
	// Immutable. The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed. If you need to change it, please create another BitbucketServerConfig.
	HostUri  string  `pulumi:"hostUri"`
	Location *string `pulumi:"location"`
	// The resource name for the config.
	Name *string `pulumi:"name"`
	// Optional. The network to be used when reaching out to the Bitbucket Server instance. The VPC network must be enabled for private service connection. This should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty, no network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format `projects/{project}/global/networks/{network}`, where {project} is a project number or id and {network} is the name of a VPC network in the project.
	PeeredNetwork *string `pulumi:"peeredNetwork"`
	Project       *string `pulumi:"project"`
	// Secret Manager secrets needed by the config.
	Secrets BitbucketServerSecrets `pulumi:"secrets"`
	// Optional. SSL certificate to use for requests to Bitbucket Server. The format should be PEM format but the extension can be one of .pem, .cer, or .crt.
	SslCa *string `pulumi:"sslCa"`
	// Username of the account Cloud Build will use on Bitbucket Server.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a BitbucketServerConfig resource.
type BitbucketServerConfigArgs struct {
	// Immutable. API Key that will be attached to webhook. Once this field has been set, it cannot be changed. If you need to change it, please create another BitbucketServerConfig.
	ApiKey                  pulumi.StringInput
	BitbucketServerConfigId pulumi.StringPtrInput
	// Time when the config was created.
	CreateTime pulumi.StringPtrInput
	// Immutable. The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed. If you need to change it, please create another BitbucketServerConfig.
	HostUri  pulumi.StringInput
	Location pulumi.StringPtrInput
	// The resource name for the config.
	Name pulumi.StringPtrInput
	// Optional. The network to be used when reaching out to the Bitbucket Server instance. The VPC network must be enabled for private service connection. This should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty, no network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format `projects/{project}/global/networks/{network}`, where {project} is a project number or id and {network} is the name of a VPC network in the project.
	PeeredNetwork pulumi.StringPtrInput
	Project       pulumi.StringPtrInput
	// Secret Manager secrets needed by the config.
	Secrets BitbucketServerSecretsInput
	// Optional. SSL certificate to use for requests to Bitbucket Server. The format should be PEM format but the extension can be one of .pem, .cer, or .crt.
	SslCa pulumi.StringPtrInput
	// Username of the account Cloud Build will use on Bitbucket Server.
	Username pulumi.StringPtrInput
}

func (BitbucketServerConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bitbucketServerConfigArgs)(nil)).Elem()
}

type BitbucketServerConfigInput interface {
	pulumi.Input

	ToBitbucketServerConfigOutput() BitbucketServerConfigOutput
	ToBitbucketServerConfigOutputWithContext(ctx context.Context) BitbucketServerConfigOutput
}

func (*BitbucketServerConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**BitbucketServerConfig)(nil)).Elem()
}

func (i *BitbucketServerConfig) ToBitbucketServerConfigOutput() BitbucketServerConfigOutput {
	return i.ToBitbucketServerConfigOutputWithContext(context.Background())
}

func (i *BitbucketServerConfig) ToBitbucketServerConfigOutputWithContext(ctx context.Context) BitbucketServerConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BitbucketServerConfigOutput)
}

type BitbucketServerConfigOutput struct{ *pulumi.OutputState }

func (BitbucketServerConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BitbucketServerConfig)(nil)).Elem()
}

func (o BitbucketServerConfigOutput) ToBitbucketServerConfigOutput() BitbucketServerConfigOutput {
	return o
}

func (o BitbucketServerConfigOutput) ToBitbucketServerConfigOutputWithContext(ctx context.Context) BitbucketServerConfigOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BitbucketServerConfigInput)(nil)).Elem(), &BitbucketServerConfig{})
	pulumi.RegisterOutputType(BitbucketServerConfigOutput{})
}
