// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new User data mapping in the parent consent store.
type UserDataMapping struct {
	pulumi.CustomResourceState

	// Indicates the time when this mapping was archived.
	ArchiveTime pulumi.StringOutput `pulumi:"archiveTime"`
	// Indicates whether this mapping is archived.
	Archived pulumi.BoolOutput `pulumi:"archived"`
	// A unique identifier for the mapped resource.
	DataId pulumi.StringOutput `pulumi:"dataId"`
	// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
	ResourceAttributes AttributeResponseArrayOutput `pulumi:"resourceAttributes"`
	// User's UUID provided by the client.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserDataMapping registers a new resource with the given unique name, arguments, and options.
func NewUserDataMapping(ctx *pulumi.Context,
	name string, args *UserDataMappingArgs, opts ...pulumi.ResourceOption) (*UserDataMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsentStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConsentStoreId'")
	}
	if args.DataId == nil {
		return nil, errors.New("invalid value for required argument 'DataId'")
	}
	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource UserDataMapping
	err := ctx.RegisterResource("google-native:healthcare/v1:UserDataMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserDataMapping gets an existing UserDataMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserDataMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserDataMappingState, opts ...pulumi.ResourceOption) (*UserDataMapping, error) {
	var resource UserDataMapping
	err := ctx.ReadResource("google-native:healthcare/v1:UserDataMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserDataMapping resources.
type userDataMappingState struct {
}

type UserDataMappingState struct {
}

func (UserDataMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*userDataMappingState)(nil)).Elem()
}

type userDataMappingArgs struct {
	ConsentStoreId string `pulumi:"consentStoreId"`
	// A unique identifier for the mapped resource.
	DataId    string  `pulumi:"dataId"`
	DatasetId string  `pulumi:"datasetId"`
	Location  *string `pulumi:"location"`
	// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
	ResourceAttributes []Attribute `pulumi:"resourceAttributes"`
	// User's UUID provided by the client.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a UserDataMapping resource.
type UserDataMappingArgs struct {
	ConsentStoreId pulumi.StringInput
	// A unique identifier for the mapped resource.
	DataId    pulumi.StringInput
	DatasetId pulumi.StringInput
	Location  pulumi.StringPtrInput
	// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
	ResourceAttributes AttributeArrayInput
	// User's UUID provided by the client.
	UserId pulumi.StringInput
}

func (UserDataMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userDataMappingArgs)(nil)).Elem()
}

type UserDataMappingInput interface {
	pulumi.Input

	ToUserDataMappingOutput() UserDataMappingOutput
	ToUserDataMappingOutputWithContext(ctx context.Context) UserDataMappingOutput
}

func (*UserDataMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDataMapping)(nil)).Elem()
}

func (i *UserDataMapping) ToUserDataMappingOutput() UserDataMappingOutput {
	return i.ToUserDataMappingOutputWithContext(context.Background())
}

func (i *UserDataMapping) ToUserDataMappingOutputWithContext(ctx context.Context) UserDataMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDataMappingOutput)
}

type UserDataMappingOutput struct{ *pulumi.OutputState }

func (UserDataMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDataMapping)(nil)).Elem()
}

func (o UserDataMappingOutput) ToUserDataMappingOutput() UserDataMappingOutput {
	return o
}

func (o UserDataMappingOutput) ToUserDataMappingOutputWithContext(ctx context.Context) UserDataMappingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserDataMappingInput)(nil)).Elem(), &UserDataMapping{})
	pulumi.RegisterOutputType(UserDataMappingOutput{})
}
