// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an API proxy. The API proxy created will not be accessible at runtime until it is deployed to an environment. Create a new API proxy by setting the `name` query parameter to the name of the API proxy. Import an API proxy configuration bundle stored in zip format on your local machine to your organization by doing the following: * Set the `name` query parameter to the name of the API proxy. * Set the `action` query parameter to `import`. * Set the `Content-Type` header to `multipart/form-data`. * Pass as a file the name of API proxy configuration bundle stored in zip format on your local machine using the `file` form field. **Note**: To validate the API proxy configuration bundle only without importing it, set the `action` query parameter to `validate`. When importing an API proxy configuration bundle, if the API proxy does not exist, it will be created. If the API proxy exists, then a new revision is created. Invalid API proxy configurations are rejected, and a list of validation errors is returned to the client.
type OrganizationApi struct {
	pulumi.CustomResourceState
}

// NewOrganizationApi registers a new resource with the given unique name, arguments, and options.
func NewOrganizationApi(ctx *pulumi.Context,
	name string, args *OrganizationApiArgs, opts ...pulumi.ResourceOption) (*OrganizationApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApisId == nil {
		return nil, errors.New("invalid value for required argument 'ApisId'")
	}
	if args.OrganizationsId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationsId'")
	}
	var resource OrganizationApi
	err := ctx.RegisterResource("google-cloud:apigee/v1:OrganizationApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationApi gets an existing OrganizationApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationApiState, opts ...pulumi.ResourceOption) (*OrganizationApi, error) {
	var resource OrganizationApi
	err := ctx.ReadResource("google-cloud:apigee/v1:OrganizationApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationApi resources.
type organizationApiState struct {
}

type OrganizationApiState struct {
}

func (OrganizationApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationApiState)(nil)).Elem()
}

type organizationApiArgs struct {
	ApisId string `pulumi:"apisId"`
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType *string `pulumi:"contentType"`
	// The HTTP request/response body as raw binary.
	Data *string `pulumi:"data"`
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions      []map[string]string `pulumi:"extensions"`
	OrganizationsId string              `pulumi:"organizationsId"`
}

// The set of arguments for constructing a OrganizationApi resource.
type OrganizationApiArgs struct {
	ApisId pulumi.StringInput
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType pulumi.StringPtrInput
	// The HTTP request/response body as raw binary.
	Data pulumi.StringPtrInput
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions      pulumi.StringMapArrayInput
	OrganizationsId pulumi.StringInput
}

func (OrganizationApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationApiArgs)(nil)).Elem()
}

type OrganizationApiInput interface {
	pulumi.Input

	ToOrganizationApiOutput() OrganizationApiOutput
	ToOrganizationApiOutputWithContext(ctx context.Context) OrganizationApiOutput
}

func (*OrganizationApi) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationApi)(nil))
}

func (i *OrganizationApi) ToOrganizationApiOutput() OrganizationApiOutput {
	return i.ToOrganizationApiOutputWithContext(context.Background())
}

func (i *OrganizationApi) ToOrganizationApiOutputWithContext(ctx context.Context) OrganizationApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationApiOutput)
}

type OrganizationApiOutput struct {
	*pulumi.OutputState
}

func (OrganizationApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationApi)(nil))
}

func (o OrganizationApiOutput) ToOrganizationApiOutput() OrganizationApiOutput {
	return o
}

func (o OrganizationApiOutput) ToOrganizationApiOutputWithContext(ctx context.Context) OrganizationApiOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationApiOutput{})
}
