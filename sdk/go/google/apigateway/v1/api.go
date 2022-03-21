// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Api in a given project and location.
// Auto-naming is currently not supported for this resource.
type Api struct {
	pulumi.CustomResourceState

	// Created time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Optional. Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed). If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService pulumi.StringOutput `pulumi:"managedService"`
	// Resource name of the API. Format: projects/{project}/locations/global/apis/{api}
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the API.
	State pulumi.StringOutput `pulumi:"state"`
	// Updated time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	var resource Api
	err := ctx.RegisterResource("google-native:apigateway/v1:Api", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:apigateway/v1:Api", name, id, state, &resource, opts...)
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
	// Required. Identifier to assign to the API. Must be unique within scope of the parent resource.
	ApiId string `pulumi:"apiId"`
	// Optional. Display name.
	DisplayName *string `pulumi:"displayName"`
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed). If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService *string `pulumi:"managedService"`
	Project        *string `pulumi:"project"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// Required. Identifier to assign to the API. Must be unique within scope of the parent resource.
	ApiId pulumi.StringInput
	// Optional. Display name.
	DisplayName pulumi.StringPtrInput
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed). If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService pulumi.StringPtrInput
	Project        pulumi.StringPtrInput
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
