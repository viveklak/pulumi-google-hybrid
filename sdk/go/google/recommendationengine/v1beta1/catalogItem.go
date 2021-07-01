// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a catalog item.
type CatalogItem struct {
	pulumi.CustomResourceState

	// Required. Catalog item categories. This field is repeated for supporting one catalog item belonging to several parallel category hierarchies. For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categoryHierarchies": [ { "categories": ["Shoes & Accessories", "Shoes"]}, { "categories": ["Sports & Fitness", "Athletic Clothing", "Shoes"] } ]
	CategoryHierarchies GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyResponseArrayOutput `pulumi:"categoryHierarchies"`
	// Optional. Catalog item description. UTF-8 encoded string with a length limit of 5 KiB.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. Highly encouraged. Extra catalog item attributes to be included in the recommendation model. For example, for retail products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the item attributes here.
	ItemAttributes GoogleCloudRecommendationengineV1beta1FeatureMapResponseOutput `pulumi:"itemAttributes"`
	// Optional. Variant group identifier for prediction results. UTF-8 encoded string with a length limit of 128 bytes. This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
	ItemGroupId pulumi.StringOutput `pulumi:"itemGroupId"`
	// Optional. Metadata specific to retail products.
	ProductMetadata GoogleCloudRecommendationengineV1beta1ProductCatalogItemResponseOutput `pulumi:"productMetadata"`
	// Optional. Filtering tags associated with the catalog item. Each tag should be a UTF-8 encoded string with a length limit of 1 KiB. This tag can be used for filtering recommendation results by passing the tag as part of the predict request filter.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Required. Catalog item title. UTF-8 encoded string with a length limit of 1 KiB.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewCatalogItem registers a new resource with the given unique name, arguments, and options.
func NewCatalogItem(ctx *pulumi.Context,
	name string, args *CatalogItemArgs, opts ...pulumi.ResourceOption) (*CatalogItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource CatalogItem
	err := ctx.RegisterResource("google-native:recommendationengine/v1beta1:CatalogItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCatalogItem gets an existing CatalogItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCatalogItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CatalogItemState, opts ...pulumi.ResourceOption) (*CatalogItem, error) {
	var resource CatalogItem
	err := ctx.ReadResource("google-native:recommendationengine/v1beta1:CatalogItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CatalogItem resources.
type catalogItemState struct {
}

type CatalogItemState struct {
}

func (CatalogItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogItemState)(nil)).Elem()
}

type catalogItemArgs struct {
	CatalogId string `pulumi:"catalogId"`
	// Required. Catalog item categories. This field is repeated for supporting one catalog item belonging to several parallel category hierarchies. For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categoryHierarchies": [ { "categories": ["Shoes & Accessories", "Shoes"]}, { "categories": ["Sports & Fitness", "Athletic Clothing", "Shoes"] } ]
	CategoryHierarchies []GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchy `pulumi:"categoryHierarchies"`
	// Optional. Catalog item description. UTF-8 encoded string with a length limit of 5 KiB.
	Description *string `pulumi:"description"`
	// Required. Catalog item identifier. UTF-8 encoded string with a length limit of 128 bytes. This id must be unique among all catalog items within the same catalog. It should also be used when logging user events in order for the user events to be joined with the Catalog.
	Id *string `pulumi:"id"`
	// Optional. Highly encouraged. Extra catalog item attributes to be included in the recommendation model. For example, for retail products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the item attributes here.
	ItemAttributes *GoogleCloudRecommendationengineV1beta1FeatureMap `pulumi:"itemAttributes"`
	// Optional. Variant group identifier for prediction results. UTF-8 encoded string with a length limit of 128 bytes. This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
	ItemGroupId *string `pulumi:"itemGroupId"`
	Location    string  `pulumi:"location"`
	// Optional. Metadata specific to retail products.
	ProductMetadata *GoogleCloudRecommendationengineV1beta1ProductCatalogItem `pulumi:"productMetadata"`
	Project         string                                                    `pulumi:"project"`
	// Optional. Filtering tags associated with the catalog item. Each tag should be a UTF-8 encoded string with a length limit of 1 KiB. This tag can be used for filtering recommendation results by passing the tag as part of the predict request filter.
	Tags []string `pulumi:"tags"`
	// Required. Catalog item title. UTF-8 encoded string with a length limit of 1 KiB.
	Title *string `pulumi:"title"`
}

// The set of arguments for constructing a CatalogItem resource.
type CatalogItemArgs struct {
	CatalogId pulumi.StringInput
	// Required. Catalog item categories. This field is repeated for supporting one catalog item belonging to several parallel category hierarchies. For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categoryHierarchies": [ { "categories": ["Shoes & Accessories", "Shoes"]}, { "categories": ["Sports & Fitness", "Athletic Clothing", "Shoes"] } ]
	CategoryHierarchies GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyArrayInput
	// Optional. Catalog item description. UTF-8 encoded string with a length limit of 5 KiB.
	Description pulumi.StringPtrInput
	// Required. Catalog item identifier. UTF-8 encoded string with a length limit of 128 bytes. This id must be unique among all catalog items within the same catalog. It should also be used when logging user events in order for the user events to be joined with the Catalog.
	Id pulumi.StringPtrInput
	// Optional. Highly encouraged. Extra catalog item attributes to be included in the recommendation model. For example, for retail products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the item attributes here.
	ItemAttributes GoogleCloudRecommendationengineV1beta1FeatureMapPtrInput
	// Optional. Variant group identifier for prediction results. UTF-8 encoded string with a length limit of 128 bytes. This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
	ItemGroupId pulumi.StringPtrInput
	Location    pulumi.StringInput
	// Optional. Metadata specific to retail products.
	ProductMetadata GoogleCloudRecommendationengineV1beta1ProductCatalogItemPtrInput
	Project         pulumi.StringInput
	// Optional. Filtering tags associated with the catalog item. Each tag should be a UTF-8 encoded string with a length limit of 1 KiB. This tag can be used for filtering recommendation results by passing the tag as part of the predict request filter.
	Tags pulumi.StringArrayInput
	// Required. Catalog item title. UTF-8 encoded string with a length limit of 1 KiB.
	Title pulumi.StringPtrInput
}

func (CatalogItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogItemArgs)(nil)).Elem()
}

type CatalogItemInput interface {
	pulumi.Input

	ToCatalogItemOutput() CatalogItemOutput
	ToCatalogItemOutputWithContext(ctx context.Context) CatalogItemOutput
}

func (*CatalogItem) ElementType() reflect.Type {
	return reflect.TypeOf((*CatalogItem)(nil))
}

func (i *CatalogItem) ToCatalogItemOutput() CatalogItemOutput {
	return i.ToCatalogItemOutputWithContext(context.Background())
}

func (i *CatalogItem) ToCatalogItemOutputWithContext(ctx context.Context) CatalogItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogItemOutput)
}

type CatalogItemOutput struct {
	*pulumi.OutputState
}

func (CatalogItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CatalogItem)(nil))
}

func (o CatalogItemOutput) ToCatalogItemOutput() CatalogItemOutput {
	return o
}

func (o CatalogItemOutput) ToCatalogItemOutputWithContext(ctx context.Context) CatalogItemOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CatalogItemOutput{})
}
