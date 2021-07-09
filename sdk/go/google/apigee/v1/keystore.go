// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a keystore or truststore. - Keystore: Contains certificates and their associated keys. - Truststore: Contains trusted certificates used to validate a server's certificate. These certificates are typically self-signed certificates or certificates that are not signed by a trusted CA.
type Keystore struct {
	pulumi.CustomResourceState

	// Aliases in this keystore.
	Aliases pulumi.StringArrayOutput `pulumi:"aliases"`
	// Resource ID for this keystore. Values must match the regular expression `[\w[:space:]-.]{1,255}`.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewKeystore registers a new resource with the given unique name, arguments, and options.
func NewKeystore(ctx *pulumi.Context,
	name string, args *KeystoreArgs, opts ...pulumi.ResourceOption) (*Keystore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource Keystore
	err := ctx.RegisterResource("google-native:apigee/v1:Keystore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeystore gets an existing Keystore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeystore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeystoreState, opts ...pulumi.ResourceOption) (*Keystore, error) {
	var resource Keystore
	err := ctx.ReadResource("google-native:apigee/v1:Keystore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Keystore resources.
type keystoreState struct {
	// Aliases in this keystore.
	Aliases []string `pulumi:"aliases"`
	// Resource ID for this keystore. Values must match the regular expression `[\w[:space:]-.]{1,255}`.
	Name *string `pulumi:"name"`
}

type KeystoreState struct {
	// Aliases in this keystore.
	Aliases pulumi.StringArrayInput
	// Resource ID for this keystore. Values must match the regular expression `[\w[:space:]-.]{1,255}`.
	Name pulumi.StringPtrInput
}

func (KeystoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*keystoreState)(nil)).Elem()
}

type keystoreArgs struct {
	EnvironmentId string `pulumi:"environmentId"`
	// Resource ID for this keystore. Values must match the regular expression `[\w[:space:]-.]{1,255}`.
	Name           string `pulumi:"name"`
	OrganizationId string `pulumi:"organizationId"`
}

// The set of arguments for constructing a Keystore resource.
type KeystoreArgs struct {
	EnvironmentId pulumi.StringInput
	// Resource ID for this keystore. Values must match the regular expression `[\w[:space:]-.]{1,255}`.
	Name           pulumi.StringInput
	OrganizationId pulumi.StringInput
}

func (KeystoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keystoreArgs)(nil)).Elem()
}

type KeystoreInput interface {
	pulumi.Input

	ToKeystoreOutput() KeystoreOutput
	ToKeystoreOutputWithContext(ctx context.Context) KeystoreOutput
}

func (*Keystore) ElementType() reflect.Type {
	return reflect.TypeOf((*Keystore)(nil))
}

func (i *Keystore) ToKeystoreOutput() KeystoreOutput {
	return i.ToKeystoreOutputWithContext(context.Background())
}

func (i *Keystore) ToKeystoreOutputWithContext(ctx context.Context) KeystoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeystoreOutput)
}

type KeystoreOutput struct {
	*pulumi.OutputState
}

func (KeystoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Keystore)(nil))
}

func (o KeystoreOutput) ToKeystoreOutput() KeystoreOutput {
	return o
}

func (o KeystoreOutput) ToKeystoreOutputWithContext(ctx context.Context) KeystoreOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KeystoreOutput{})
}
