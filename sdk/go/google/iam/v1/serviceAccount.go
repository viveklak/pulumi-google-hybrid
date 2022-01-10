// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a ServiceAccount.
// Auto-naming is currently not supported for this resource.
type ServiceAccount struct {
	pulumi.CustomResourceState

	// Optional. A user-specified, human-readable description of the service account. The maximum length is 256 UTF-8 bytes.
	Description pulumi.StringOutput `pulumi:"description"`
	// Whether the service account is disabled.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Optional. A user-specified, human-readable name for the service account. The maximum length is 100 UTF-8 bytes.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The email address of the service account.
	Email pulumi.StringOutput `pulumi:"email"`
	// The resource name of the service account. Use one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` As an alternative, you can use the `-` wildcard character instead of the project ID: * `projects/-/serviceAccounts/{EMAIL_ADDRESS}` * `projects/-/serviceAccounts/{UNIQUE_ID}` When possible, avoid using the `-` wildcard character, because it can cause response messages to contain misleading error codes. For example, if you try to get the service account `projects/-/serviceAccounts/fake@example.com`, which does not exist, the response contains an HTTP `403 Forbidden` error instead of a `404 Not Found` error.
	Name pulumi.StringOutput `pulumi:"name"`
	// The OAuth 2.0 client ID for the service account.
	Oauth2ClientId pulumi.StringOutput `pulumi:"oauth2ClientId"`
	// The ID of the project that owns the service account.
	Project pulumi.StringOutput `pulumi:"project"`
	// The unique, stable numeric ID for the service account. Each service account retains its unique ID even if you delete the service account. For example, if you delete a service account, then create a new service account with the same name, the new service account has a different unique ID than the deleted service account.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewServiceAccount(ctx *pulumi.Context,
	name string, args *ServiceAccountArgs, opts ...pulumi.ResourceOption) (*ServiceAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	var resource ServiceAccount
	err := ctx.RegisterResource("google-native:iam/v1:ServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAccount gets an existing ServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAccountState, opts ...pulumi.ResourceOption) (*ServiceAccount, error) {
	var resource ServiceAccount
	err := ctx.ReadResource("google-native:iam/v1:ServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAccount resources.
type serviceAccountState struct {
}

type ServiceAccountState struct {
}

func (ServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountState)(nil)).Elem()
}

type serviceAccountArgs struct {
	// The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])` to comply with RFC1035.
	AccountId string `pulumi:"accountId"`
	// Optional. A user-specified, human-readable description of the service account. The maximum length is 256 UTF-8 bytes.
	Description *string `pulumi:"description"`
	// Optional. A user-specified, human-readable name for the service account. The maximum length is 100 UTF-8 bytes.
	DisplayName *string `pulumi:"displayName"`
	// The resource name of the service account. Use one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` As an alternative, you can use the `-` wildcard character instead of the project ID: * `projects/-/serviceAccounts/{EMAIL_ADDRESS}` * `projects/-/serviceAccounts/{UNIQUE_ID}` When possible, avoid using the `-` wildcard character, because it can cause response messages to contain misleading error codes. For example, if you try to get the service account `projects/-/serviceAccounts/fake@example.com`, which does not exist, the response contains an HTTP `403 Forbidden` error instead of a `404 Not Found` error.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ServiceAccount resource.
type ServiceAccountArgs struct {
	// The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])` to comply with RFC1035.
	AccountId pulumi.StringInput
	// Optional. A user-specified, human-readable description of the service account. The maximum length is 256 UTF-8 bytes.
	Description pulumi.StringPtrInput
	// Optional. A user-specified, human-readable name for the service account. The maximum length is 100 UTF-8 bytes.
	DisplayName pulumi.StringPtrInput
	// The resource name of the service account. Use one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` As an alternative, you can use the `-` wildcard character instead of the project ID: * `projects/-/serviceAccounts/{EMAIL_ADDRESS}` * `projects/-/serviceAccounts/{UNIQUE_ID}` When possible, avoid using the `-` wildcard character, because it can cause response messages to contain misleading error codes. For example, if you try to get the service account `projects/-/serviceAccounts/fake@example.com`, which does not exist, the response contains an HTTP `403 Forbidden` error instead of a `404 Not Found` error.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (ServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountArgs)(nil)).Elem()
}

type ServiceAccountInput interface {
	pulumi.Input

	ToServiceAccountOutput() ServiceAccountOutput
	ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput
}

func (*ServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccount)(nil)).Elem()
}

func (i *ServiceAccount) ToServiceAccountOutput() ServiceAccountOutput {
	return i.ToServiceAccountOutputWithContext(context.Background())
}

func (i *ServiceAccount) ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountOutput)
}

type ServiceAccountOutput struct{ *pulumi.OutputState }

func (ServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccount)(nil)).Elem()
}

func (o ServiceAccountOutput) ToServiceAccountOutput() ServiceAccountOutput {
	return o
}

func (o ServiceAccountOutput) ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountInput)(nil)).Elem(), &ServiceAccount{})
	pulumi.RegisterOutputType(ServiceAccountOutput{})
}
