// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a developer. Once created, the developer can register an app and obtain an API key. At creation time, a developer is set as `active`. To change the developer status, use the SetDeveloperStatus API.
type Developer struct {
	pulumi.CustomResourceState

	// Access type.
	AccessType pulumi.StringOutput `pulumi:"accessType"`
	// Developer app family.
	AppFamily pulumi.StringOutput `pulumi:"appFamily"`
	// List of apps associated with the developer.
	Apps pulumi.StringArrayOutput `pulumi:"apps"`
	// Optional. Developer attributes (name/value pairs). The custom attribute limit is 18.
	Attributes GoogleCloudApigeeV1AttributeResponseArrayOutput `pulumi:"attributes"`
	// List of companies associated with the developer.
	Companies pulumi.StringArrayOutput `pulumi:"companies"`
	// Time at which the developer was created in milliseconds since epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// ID of the developer. **Note**: IDs are generated internally by Apigee and are not guaranteed to stay the same over time.
	DeveloperId pulumi.StringOutput `pulumi:"developerId"`
	// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only.
	Email pulumi.StringOutput `pulumi:"email"`
	// First name of the developer.
	FirstName pulumi.StringOutput `pulumi:"firstName"`
	// Time at which the developer was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// Last name of the developer.
	LastName pulumi.StringOutput `pulumi:"lastName"`
	// Name of the Apigee organization in which the developer resides.
	OrganizationName pulumi.StringOutput `pulumi:"organizationName"`
	// Status of the developer. Valid values are `active` and `inactive`.
	Status pulumi.StringOutput `pulumi:"status"`
	// User name of the developer. Not used by Apigee hybrid.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewDeveloper registers a new resource with the given unique name, arguments, and options.
func NewDeveloper(ctx *pulumi.Context,
	name string, args *DeveloperArgs, opts ...pulumi.ResourceOption) (*Developer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.FirstName == nil {
		return nil, errors.New("invalid value for required argument 'FirstName'")
	}
	if args.LastName == nil {
		return nil, errors.New("invalid value for required argument 'LastName'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	var resource Developer
	err := ctx.RegisterResource("google-native:apigee/v1:Developer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeveloper gets an existing Developer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeveloper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeveloperState, opts ...pulumi.ResourceOption) (*Developer, error) {
	var resource Developer
	err := ctx.ReadResource("google-native:apigee/v1:Developer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Developer resources.
type developerState struct {
	// Access type.
	AccessType *string `pulumi:"accessType"`
	// Developer app family.
	AppFamily *string `pulumi:"appFamily"`
	// List of apps associated with the developer.
	Apps []string `pulumi:"apps"`
	// Optional. Developer attributes (name/value pairs). The custom attribute limit is 18.
	Attributes []GoogleCloudApigeeV1AttributeResponse `pulumi:"attributes"`
	// List of companies associated with the developer.
	Companies []string `pulumi:"companies"`
	// Time at which the developer was created in milliseconds since epoch.
	CreatedAt *string `pulumi:"createdAt"`
	// ID of the developer. **Note**: IDs are generated internally by Apigee and are not guaranteed to stay the same over time.
	DeveloperId *string `pulumi:"developerId"`
	// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only.
	Email *string `pulumi:"email"`
	// First name of the developer.
	FirstName *string `pulumi:"firstName"`
	// Time at which the developer was last modified in milliseconds since epoch.
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// Last name of the developer.
	LastName *string `pulumi:"lastName"`
	// Name of the Apigee organization in which the developer resides.
	OrganizationName *string `pulumi:"organizationName"`
	// Status of the developer. Valid values are `active` and `inactive`.
	Status *string `pulumi:"status"`
	// User name of the developer. Not used by Apigee hybrid.
	UserName *string `pulumi:"userName"`
}

type DeveloperState struct {
	// Access type.
	AccessType pulumi.StringPtrInput
	// Developer app family.
	AppFamily pulumi.StringPtrInput
	// List of apps associated with the developer.
	Apps pulumi.StringArrayInput
	// Optional. Developer attributes (name/value pairs). The custom attribute limit is 18.
	Attributes GoogleCloudApigeeV1AttributeResponseArrayInput
	// List of companies associated with the developer.
	Companies pulumi.StringArrayInput
	// Time at which the developer was created in milliseconds since epoch.
	CreatedAt pulumi.StringPtrInput
	// ID of the developer. **Note**: IDs are generated internally by Apigee and are not guaranteed to stay the same over time.
	DeveloperId pulumi.StringPtrInput
	// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only.
	Email pulumi.StringPtrInput
	// First name of the developer.
	FirstName pulumi.StringPtrInput
	// Time at which the developer was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringPtrInput
	// Last name of the developer.
	LastName pulumi.StringPtrInput
	// Name of the Apigee organization in which the developer resides.
	OrganizationName pulumi.StringPtrInput
	// Status of the developer. Valid values are `active` and `inactive`.
	Status pulumi.StringPtrInput
	// User name of the developer. Not used by Apigee hybrid.
	UserName pulumi.StringPtrInput
}

func (DeveloperState) ElementType() reflect.Type {
	return reflect.TypeOf((*developerState)(nil)).Elem()
}

type developerArgs struct {
	// Access type.
	AccessType *string `pulumi:"accessType"`
	// Developer app family.
	AppFamily *string `pulumi:"appFamily"`
	// List of apps associated with the developer.
	Apps []string `pulumi:"apps"`
	// Optional. Developer attributes (name/value pairs). The custom attribute limit is 18.
	Attributes []GoogleCloudApigeeV1Attribute `pulumi:"attributes"`
	// List of companies associated with the developer.
	Companies []string `pulumi:"companies"`
	// ID of the developer. **Note**: IDs are generated internally by Apigee and are not guaranteed to stay the same over time.
	DeveloperId *string `pulumi:"developerId"`
	// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only.
	Email string `pulumi:"email"`
	// First name of the developer.
	FirstName string `pulumi:"firstName"`
	// Last name of the developer.
	LastName       string `pulumi:"lastName"`
	OrganizationId string `pulumi:"organizationId"`
	// User name of the developer. Not used by Apigee hybrid.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a Developer resource.
type DeveloperArgs struct {
	// Access type.
	AccessType pulumi.StringPtrInput
	// Developer app family.
	AppFamily pulumi.StringPtrInput
	// List of apps associated with the developer.
	Apps pulumi.StringArrayInput
	// Optional. Developer attributes (name/value pairs). The custom attribute limit is 18.
	Attributes GoogleCloudApigeeV1AttributeArrayInput
	// List of companies associated with the developer.
	Companies pulumi.StringArrayInput
	// ID of the developer. **Note**: IDs are generated internally by Apigee and are not guaranteed to stay the same over time.
	DeveloperId pulumi.StringPtrInput
	// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only.
	Email pulumi.StringInput
	// First name of the developer.
	FirstName pulumi.StringInput
	// Last name of the developer.
	LastName       pulumi.StringInput
	OrganizationId pulumi.StringInput
	// User name of the developer. Not used by Apigee hybrid.
	UserName pulumi.StringInput
}

func (DeveloperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*developerArgs)(nil)).Elem()
}

type DeveloperInput interface {
	pulumi.Input

	ToDeveloperOutput() DeveloperOutput
	ToDeveloperOutputWithContext(ctx context.Context) DeveloperOutput
}

func (*Developer) ElementType() reflect.Type {
	return reflect.TypeOf((*Developer)(nil))
}

func (i *Developer) ToDeveloperOutput() DeveloperOutput {
	return i.ToDeveloperOutputWithContext(context.Background())
}

func (i *Developer) ToDeveloperOutputWithContext(ctx context.Context) DeveloperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeveloperOutput)
}

type DeveloperOutput struct {
	*pulumi.OutputState
}

func (DeveloperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Developer)(nil))
}

func (o DeveloperOutput) ToDeveloperOutput() DeveloperOutput {
	return o
}

func (o DeveloperOutput) ToDeveloperOutputWithContext(ctx context.Context) DeveloperOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DeveloperOutput{})
}
