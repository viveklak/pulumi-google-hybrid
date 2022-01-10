// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Attribute definition in the parent consent store.
type AttributeDefinition struct {
	pulumi.CustomResourceState

	// Possible values for the attribute. The number of allowed values must not exceed 500. An empty list is invalid. The list can only be expanded after creation.
	AllowedValues pulumi.StringArrayOutput `pulumi:"allowedValues"`
	// The category of the attribute. The value of this field cannot be changed after creation.
	Category pulumi.StringOutput `pulumi:"category"`
	// Optional. Default values of the attribute in Consents. If no default values are specified, it defaults to an empty value.
	ConsentDefaultValues pulumi.StringArrayOutput `pulumi:"consentDefaultValues"`
	// Optional. Default value of the attribute in User data mappings. If no default value is specified, it defaults to an empty value. This field is only applicable to attributes of the category `RESOURCE`.
	DataMappingDefaultValue pulumi.StringOutput `pulumi:"dataMappingDefaultValue"`
	// Optional. A description of the attribute.
	Description pulumi.StringOutput `pulumi:"description"`
	// Resource name of the Attribute definition, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/attributeDefinitions/{attribute_definition_id}`. Cannot be changed after creation.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAttributeDefinition registers a new resource with the given unique name, arguments, and options.
func NewAttributeDefinition(ctx *pulumi.Context,
	name string, args *AttributeDefinitionArgs, opts ...pulumi.ResourceOption) (*AttributeDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowedValues == nil {
		return nil, errors.New("invalid value for required argument 'AllowedValues'")
	}
	if args.AttributeDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'AttributeDefinitionId'")
	}
	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.ConsentStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConsentStoreId'")
	}
	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	var resource AttributeDefinition
	err := ctx.RegisterResource("google-native:healthcare/v1beta1:AttributeDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttributeDefinition gets an existing AttributeDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttributeDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttributeDefinitionState, opts ...pulumi.ResourceOption) (*AttributeDefinition, error) {
	var resource AttributeDefinition
	err := ctx.ReadResource("google-native:healthcare/v1beta1:AttributeDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttributeDefinition resources.
type attributeDefinitionState struct {
}

type AttributeDefinitionState struct {
}

func (AttributeDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeDefinitionState)(nil)).Elem()
}

type attributeDefinitionArgs struct {
	// Possible values for the attribute. The number of allowed values must not exceed 500. An empty list is invalid. The list can only be expanded after creation.
	AllowedValues         []string `pulumi:"allowedValues"`
	AttributeDefinitionId string   `pulumi:"attributeDefinitionId"`
	// The category of the attribute. The value of this field cannot be changed after creation.
	Category AttributeDefinitionCategory `pulumi:"category"`
	// Optional. Default values of the attribute in Consents. If no default values are specified, it defaults to an empty value.
	ConsentDefaultValues []string `pulumi:"consentDefaultValues"`
	ConsentStoreId       string   `pulumi:"consentStoreId"`
	// Optional. Default value of the attribute in User data mappings. If no default value is specified, it defaults to an empty value. This field is only applicable to attributes of the category `RESOURCE`.
	DataMappingDefaultValue *string `pulumi:"dataMappingDefaultValue"`
	DatasetId               string  `pulumi:"datasetId"`
	// Optional. A description of the attribute.
	Description *string `pulumi:"description"`
	Location    *string `pulumi:"location"`
	// Resource name of the Attribute definition, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/attributeDefinitions/{attribute_definition_id}`. Cannot be changed after creation.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a AttributeDefinition resource.
type AttributeDefinitionArgs struct {
	// Possible values for the attribute. The number of allowed values must not exceed 500. An empty list is invalid. The list can only be expanded after creation.
	AllowedValues         pulumi.StringArrayInput
	AttributeDefinitionId pulumi.StringInput
	// The category of the attribute. The value of this field cannot be changed after creation.
	Category AttributeDefinitionCategoryInput
	// Optional. Default values of the attribute in Consents. If no default values are specified, it defaults to an empty value.
	ConsentDefaultValues pulumi.StringArrayInput
	ConsentStoreId       pulumi.StringInput
	// Optional. Default value of the attribute in User data mappings. If no default value is specified, it defaults to an empty value. This field is only applicable to attributes of the category `RESOURCE`.
	DataMappingDefaultValue pulumi.StringPtrInput
	DatasetId               pulumi.StringInput
	// Optional. A description of the attribute.
	Description pulumi.StringPtrInput
	Location    pulumi.StringPtrInput
	// Resource name of the Attribute definition, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/attributeDefinitions/{attribute_definition_id}`. Cannot be changed after creation.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (AttributeDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeDefinitionArgs)(nil)).Elem()
}

type AttributeDefinitionInput interface {
	pulumi.Input

	ToAttributeDefinitionOutput() AttributeDefinitionOutput
	ToAttributeDefinitionOutputWithContext(ctx context.Context) AttributeDefinitionOutput
}

func (*AttributeDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeDefinition)(nil)).Elem()
}

func (i *AttributeDefinition) ToAttributeDefinitionOutput() AttributeDefinitionOutput {
	return i.ToAttributeDefinitionOutputWithContext(context.Background())
}

func (i *AttributeDefinition) ToAttributeDefinitionOutputWithContext(ctx context.Context) AttributeDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeDefinitionOutput)
}

type AttributeDefinitionOutput struct{ *pulumi.OutputState }

func (AttributeDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeDefinition)(nil)).Elem()
}

func (o AttributeDefinitionOutput) ToAttributeDefinitionOutput() AttributeDefinitionOutput {
	return o
}

func (o AttributeDefinitionOutput) ToAttributeDefinitionOutputWithContext(ctx context.Context) AttributeDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeDefinitionInput)(nil)).Elem(), &AttributeDefinition{})
	pulumi.RegisterOutputType(AttributeDefinitionOutput{})
}
