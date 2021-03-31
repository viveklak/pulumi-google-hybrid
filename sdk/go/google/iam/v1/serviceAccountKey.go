// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a ServiceAccountKey.
type ServiceAccountKey struct {
	pulumi.CustomResourceState
}

// NewServiceAccountKey registers a new resource with the given unique name, arguments, and options.
func NewServiceAccountKey(ctx *pulumi.Context,
	name string, args *ServiceAccountKeyArgs, opts ...pulumi.ResourceOption) (*ServiceAccountKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeysId == nil {
		return nil, errors.New("invalid value for required argument 'KeysId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	if args.ServiceAccountsId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountsId'")
	}
	var resource ServiceAccountKey
	err := ctx.RegisterResource("google-cloud:iam/v1:ServiceAccountKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAccountKey gets an existing ServiceAccountKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAccountKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAccountKeyState, opts ...pulumi.ResourceOption) (*ServiceAccountKey, error) {
	var resource ServiceAccountKey
	err := ctx.ReadResource("google-cloud:iam/v1:ServiceAccountKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAccountKey resources.
type serviceAccountKeyState struct {
}

type ServiceAccountKeyState struct {
}

func (ServiceAccountKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountKeyState)(nil)).Elem()
}

type serviceAccountKeyArgs struct {
	// Which type of key and algorithm to use for the key. The default is currently a 2K RSA key. However this may change in the future.
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	KeysId       string  `pulumi:"keysId"`
	// The output format of the private key. The default value is `TYPE_GOOGLE_CREDENTIALS_FILE`, which is the Google Credentials File format.
	PrivateKeyType    *string `pulumi:"privateKeyType"`
	ProjectsId        string  `pulumi:"projectsId"`
	ServiceAccountsId string  `pulumi:"serviceAccountsId"`
}

// The set of arguments for constructing a ServiceAccountKey resource.
type ServiceAccountKeyArgs struct {
	// Which type of key and algorithm to use for the key. The default is currently a 2K RSA key. However this may change in the future.
	KeyAlgorithm pulumi.StringPtrInput
	KeysId       pulumi.StringInput
	// The output format of the private key. The default value is `TYPE_GOOGLE_CREDENTIALS_FILE`, which is the Google Credentials File format.
	PrivateKeyType    pulumi.StringPtrInput
	ProjectsId        pulumi.StringInput
	ServiceAccountsId pulumi.StringInput
}

func (ServiceAccountKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountKeyArgs)(nil)).Elem()
}

type ServiceAccountKeyInput interface {
	pulumi.Input

	ToServiceAccountKeyOutput() ServiceAccountKeyOutput
	ToServiceAccountKeyOutputWithContext(ctx context.Context) ServiceAccountKeyOutput
}

func (*ServiceAccountKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccountKey)(nil))
}

func (i *ServiceAccountKey) ToServiceAccountKeyOutput() ServiceAccountKeyOutput {
	return i.ToServiceAccountKeyOutputWithContext(context.Background())
}

func (i *ServiceAccountKey) ToServiceAccountKeyOutputWithContext(ctx context.Context) ServiceAccountKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountKeyOutput)
}

type ServiceAccountKeyOutput struct {
	*pulumi.OutputState
}

func (ServiceAccountKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccountKey)(nil))
}

func (o ServiceAccountKeyOutput) ToServiceAccountKeyOutput() ServiceAccountKeyOutput {
	return o
}

func (o ServiceAccountKeyOutput) ToServiceAccountKeyOutputWithContext(ctx context.Context) ServiceAccountKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceAccountKeyOutput{})
}
