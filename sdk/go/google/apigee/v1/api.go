// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an API proxy. The API proxy created will not be accessible at runtime until it is deployed to an environment. Create a new API proxy by setting the `name` query parameter to the name of the API proxy. Import an API proxy configuration bundle stored in zip format on your local machine to your organization by doing the following: * Set the `name` query parameter to the name of the API proxy. * Set the `action` query parameter to `import`. * Set the `Content-Type` header to `multipart/form-data`. * Pass as a file the name of API proxy configuration bundle stored in zip format on your local machine using the `file` form field. **Note**: To validate the API proxy configuration bundle only without importing it, set the `action` query parameter to `validate`. When importing an API proxy configuration bundle, if the API proxy does not exist, it will be created. If the API proxy exists, then a new revision is created. Invalid API proxy configurations are rejected, and a list of validation errors is returned to the client.
type Api struct {
	pulumi.CustomResourceState

	// User labels applied to this API Proxy.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The id of the most recently created revision for this api proxy.
	LatestRevisionId pulumi.StringOutput `pulumi:"latestRevisionId"`
	// Metadata describing the API proxy.
	MetaData GoogleCloudApigeeV1EntityMetadataResponseOutput `pulumi:"metaData"`
	// Name of the API proxy.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of revisons defined for the API proxy.
	Revision pulumi.StringArrayOutput `pulumi:"revision"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource Api
	err := ctx.RegisterResource("google-native:apigee/v1:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("google-native:apigee/v1:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
}

type ApiState struct {
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	Action *string `pulumi:"action"`
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType *string `pulumi:"contentType"`
	// The HTTP request/response body as raw binary.
	Data *string `pulumi:"data"`
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions     []map[string]string `pulumi:"extensions"`
	Name           *string             `pulumi:"name"`
	OrganizationId string              `pulumi:"organizationId"`
	Validate       *string             `pulumi:"validate"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	Action pulumi.StringPtrInput
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType pulumi.StringPtrInput
	// The HTTP request/response body as raw binary.
	Data pulumi.StringPtrInput
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions     pulumi.StringMapArrayInput
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	Validate       pulumi.StringPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiInput)(nil)).Elem(), &Api{})
	pulumi.RegisterOutputType(ApiOutput{})
}
