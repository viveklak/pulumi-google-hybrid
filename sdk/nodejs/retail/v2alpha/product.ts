// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Product.
 */
export class Product extends pulumi.CustomResource {
    /**
     * Get an existing Product resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Product {
        return new Product(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:retail/v2alpha:Product';

    /**
     * Returns true if the given object is an instance of Product.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Product {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Product.__pulumiType;
    }

    /**
     * Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 200. * The key must be a UTF-8 encoded string with a length limit of 128 characters. * For indexable attribute, the key must match the pattern: `a-zA-Z0-9*`. For example, key0LikeThis or KEY_1_LIKE_THIS.
     */
    public readonly attributes!: pulumi.Output<{[key: string]: string}>;
    /**
     * The target group associated with a given audience (e.g. male, veterans, car owners, musicians, etc.) of the product.
     */
    public readonly audience!: pulumi.Output<outputs.retail.v2alpha.GoogleCloudRetailV2alphaAudienceResponse>;
    /**
     * The online availability of the Product. Default to Availability.IN_STOCK. Google Merchant Center Property [availability](https://support.google.com/merchants/answer/6324448). Schema.org Property [Offer.availability](https://schema.org/availability).
     */
    public readonly availability!: pulumi.Output<string>;
    /**
     * The available quantity of the item.
     */
    public readonly availableQuantity!: pulumi.Output<number>;
    /**
     * The timestamp when this Product becomes available for SearchService.Search.
     */
    public readonly availableTime!: pulumi.Output<string>;
    /**
     * The brands of the product. A maximum of 30 brands are allowed. Each brand must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [brand](https://support.google.com/merchants/answer/6324351). Schema.org property [Product.brand](https://schema.org/brand).
     */
    public readonly brands!: pulumi.Output<string[]>;
    /**
     * Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, please replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
     */
    public readonly categories!: pulumi.Output<string[]>;
    /**
     * The id of the collection members when type is Type.COLLECTION. Should not set it for other types. A maximum of 1000 values are allowed. Otherwise, an INVALID_ARGUMENT error is return.
     */
    public readonly collectionMemberIds!: pulumi.Output<string[]>;
    /**
     * The color of the product. Google Merchant Center property [color](https://support.google.com/merchants/answer/6324487). Schema.org property [Product.color](https://schema.org/color).
     */
    public readonly colorInfo!: pulumi.Output<outputs.retail.v2alpha.GoogleCloudRetailV2alphaColorInfoResponse>;
    /**
     * The condition of the product. Strongly encouraged to use the standard values: "new", "refurbished", "used". A maximum of 5 values are allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [condition](https://support.google.com/merchants/answer/6324469). Schema.org property [Offer.itemCondition](https://schema.org/itemCondition).
     */
    public readonly conditions!: pulumi.Output<string[]>;
    /**
     * Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). schema.org property [Product.description](https://schema.org/description).
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The timestamp when this product becomes unavailable for SearchService.Search. If it is set, the Product is not available for SearchService.Search after expire_time. However, the product can still be retrieved by ProductService.GetProduct and ProductService.ListProducts. expire_time must be later than available_time and publish_time, otherwise an INVALID_ARGUMENT error is thrown. Google Merchant Center property [expiration_date](https://support.google.com/merchants/answer/6324499).
     */
    public readonly expireTime!: pulumi.Output<string>;
    /**
     * Fulfillment information, such as the store IDs for in-store pickup or region IDs for different shipping methods. All the elements must have distinct FulfillmentInfo.type. Otherwise, an INVALID_ARGUMENT error is returned.
     */
    public readonly fulfillmentInfo!: pulumi.Output<outputs.retail.v2alpha.GoogleCloudRetailV2alphaFulfillmentInfoResponse[]>;
    /**
     * The Global Trade Item Number (GTIN) of the product. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. This field must be a Unigram. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [gtin](https://support.google.com/merchants/answer/6324461). Schema.org property [Product.isbn](https://schema.org/isbn) or [Product.gtin8](https://schema.org/gtin8) or [Product.gtin12](https://schema.org/gtin12) or [Product.gtin13](https://schema.org/gtin13) or [Product.gtin14](https://schema.org/gtin14). If the value is not a valid GTIN, an INVALID_ARGUMENT error is returned.
     */
    public readonly gtin!: pulumi.Output<string>;
    /**
     * Product images for the product.Highly recommended to put the main image to the first. A maximum of 300 images are allowed. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
     */
    public readonly images!: pulumi.Output<outputs.retail.v2alpha.GoogleCloudRetailV2alphaImageResponse[]>;
    /**
     * Language of the title/description and other string attributes. Use language tags defined by BCP 47. For product prediction, this field is ignored and the model automatically detects the text language. The Product can include text in different languages, but duplicating Products to provide text in multiple languages can result in degraded model performance. For product search this field is in use. It defaults to "en-US" if unset.
     */
    public readonly languageCode!: pulumi.Output<string>;
    /**
     * The material of the product. For example, "leather", "wooden". A maximum of 20 values are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [material](https://support.google.com/merchants/answer/6324410). Schema.org property [Product.material](https://schema.org/material).
     */
    public readonly materials!: pulumi.Output<string[]>;
    /**
     * Immutable. Full resource name of the product, such as `projects/*&#47;locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The pattern or graphic print of the product. For example, "striped", "polka dot", "paisley". A maximum of 20 values are allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [pattern](https://support.google.com/merchants/answer/6324483). Schema.org property [Product.pattern](https://schema.org/pattern).
     */
    public readonly patterns!: pulumi.Output<string[]>;
    /**
     * Product price and cost information. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
     */
    public readonly priceInfo!: pulumi.Output<outputs.retail.v2alpha.GoogleCloudRetailV2alphaPriceInfoResponse>;
    /**
     * Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center Property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org Property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID). This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
     */
    public readonly primaryProductId!: pulumi.Output<string>;
    /**
     * The promotions applied to the product. A maximum of 10 values are allowed per Product.
     */
    public readonly promotions!: pulumi.Output<outputs.retail.v2alpha.GoogleCloudRetailV2alphaPromotionResponse[]>;
    /**
     * The timestamp when the product is published by the retailer for the first time, which indicates the freshness of the products. Note that this field is different from available_time, given it purely describes product freshness regardless of when it is available on search and recommendation.
     */
    public readonly publishTime!: pulumi.Output<string>;
    /**
     * The rating of this product.
     */
    public readonly rating!: pulumi.Output<outputs.retail.v2alpha.GoogleCloudRetailV2alphaRatingResponse>;
    /**
     * Indicates which fields in the Products are returned in SearchResponse. Supported fields for all types: * audience * availability * brands * color_info * conditions * gtin * materials * name * patterns * price_info * rating * sizes * title * uri Supported fields only for Type.PRIMARY and Type.COLLECTION: * categories * description * images Supported fields only for Type.VARIANT: * Only the first image in images To mark attributes as retrievable, include paths of the form "attributes.key" where "key" is the key of a custom attribute, as specified in attributes. For Type.PRIMARY and Type.COLLECTION, the following fields are always returned in SearchResponse by default: * name For Type.VARIANT, the following fields are always returned in by default: * name * color_info Maximum number of paths is 30. Otherwise, an INVALID_ARGUMENT error is returned. Note: Returning more fields in SearchResponse may increase response payload size and serving latency.
     */
    public readonly retrievableFields!: pulumi.Output<string>;
    /**
     * The size of the product. To represent different size systems or size types, consider using this format: [[[size_system:]size_type:]size_value]. For example, in "US:MENS:M", "US" represents size system; "MENS" represents size type; "M" represents size value. In "GIRLS:27", size system is empty; "GIRLS" represents size type; "27" represents size value. In "32 inches", both size system and size type are empty, while size value is "32 inches". A maximum of 20 values are allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [size](https://support.google.com/merchants/answer/6324492), [size_type](https://support.google.com/merchants/answer/6324497) and [size_system](https://support.google.com/merchants/answer/6324502). Schema.org property [Product.size](https://schema.org/size).
     */
    public readonly sizes!: pulumi.Output<string[]>;
    /**
     * Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Input only. The TTL (time to live) of the product. If it is set, it must be a non-negative value, and expire_time is set as current timestamp plus ttl. The derived expire_time is returned in the output and ttl is left blank when retrieving the Product. If it is set, the product is not available for SearchService.Search after current timestamp plus ttl. However, the product can still be retrieved by ProductService.GetProduct and ProductService.ListProducts.
     */
    public readonly ttl!: pulumi.Output<string>;
    /**
     * Immutable. The type of the product. Default to Catalog.product_level_config.ingestion_product_type if unset.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
     */
    public readonly uri!: pulumi.Output<string>;
    /**
     * Product variants grouped together on primary product which share similar product attributes. It's automatically grouped by primary_product_id for all the product variants. Only populated for Type.PRIMARY Products. Note: This field is OUTPUT_ONLY for ProductService.GetProduct. Do not set this field in API requests.
     */
    public /*out*/ readonly variants!: pulumi.Output<outputs.retail.v2alpha.GoogleCloudRetailV2alphaProductResponse[]>;

    /**
     * Create a Product resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.branchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'branchId'");
            }
            if ((!args || args.catalogId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalogId'");
            }
            if ((!args || args.productId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productId'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["audience"] = args ? args.audience : undefined;
            resourceInputs["availability"] = args ? args.availability : undefined;
            resourceInputs["availableQuantity"] = args ? args.availableQuantity : undefined;
            resourceInputs["availableTime"] = args ? args.availableTime : undefined;
            resourceInputs["branchId"] = args ? args.branchId : undefined;
            resourceInputs["brands"] = args ? args.brands : undefined;
            resourceInputs["catalogId"] = args ? args.catalogId : undefined;
            resourceInputs["categories"] = args ? args.categories : undefined;
            resourceInputs["collectionMemberIds"] = args ? args.collectionMemberIds : undefined;
            resourceInputs["colorInfo"] = args ? args.colorInfo : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["expireTime"] = args ? args.expireTime : undefined;
            resourceInputs["fulfillmentInfo"] = args ? args.fulfillmentInfo : undefined;
            resourceInputs["gtin"] = args ? args.gtin : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["images"] = args ? args.images : undefined;
            resourceInputs["languageCode"] = args ? args.languageCode : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["materials"] = args ? args.materials : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["patterns"] = args ? args.patterns : undefined;
            resourceInputs["priceInfo"] = args ? args.priceInfo : undefined;
            resourceInputs["primaryProductId"] = args ? args.primaryProductId : undefined;
            resourceInputs["productId"] = args ? args.productId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["promotions"] = args ? args.promotions : undefined;
            resourceInputs["publishTime"] = args ? args.publishTime : undefined;
            resourceInputs["rating"] = args ? args.rating : undefined;
            resourceInputs["retrievableFields"] = args ? args.retrievableFields : undefined;
            resourceInputs["sizes"] = args ? args.sizes : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["uri"] = args ? args.uri : undefined;
            resourceInputs["variants"] = undefined /*out*/;
        } else {
            resourceInputs["attributes"] = undefined /*out*/;
            resourceInputs["audience"] = undefined /*out*/;
            resourceInputs["availability"] = undefined /*out*/;
            resourceInputs["availableQuantity"] = undefined /*out*/;
            resourceInputs["availableTime"] = undefined /*out*/;
            resourceInputs["brands"] = undefined /*out*/;
            resourceInputs["categories"] = undefined /*out*/;
            resourceInputs["collectionMemberIds"] = undefined /*out*/;
            resourceInputs["colorInfo"] = undefined /*out*/;
            resourceInputs["conditions"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["fulfillmentInfo"] = undefined /*out*/;
            resourceInputs["gtin"] = undefined /*out*/;
            resourceInputs["images"] = undefined /*out*/;
            resourceInputs["languageCode"] = undefined /*out*/;
            resourceInputs["materials"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["patterns"] = undefined /*out*/;
            resourceInputs["priceInfo"] = undefined /*out*/;
            resourceInputs["primaryProductId"] = undefined /*out*/;
            resourceInputs["promotions"] = undefined /*out*/;
            resourceInputs["publishTime"] = undefined /*out*/;
            resourceInputs["rating"] = undefined /*out*/;
            resourceInputs["retrievableFields"] = undefined /*out*/;
            resourceInputs["sizes"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["ttl"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uri"] = undefined /*out*/;
            resourceInputs["variants"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Product.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Product resource.
 */
export interface ProductArgs {
    /**
     * Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 200. * The key must be a UTF-8 encoded string with a length limit of 128 characters. * For indexable attribute, the key must match the pattern: `a-zA-Z0-9*`. For example, key0LikeThis or KEY_1_LIKE_THIS.
     */
    attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The target group associated with a given audience (e.g. male, veterans, car owners, musicians, etc.) of the product.
     */
    audience?: pulumi.Input<inputs.retail.v2alpha.GoogleCloudRetailV2alphaAudienceArgs>;
    /**
     * The online availability of the Product. Default to Availability.IN_STOCK. Google Merchant Center Property [availability](https://support.google.com/merchants/answer/6324448). Schema.org Property [Offer.availability](https://schema.org/availability).
     */
    availability?: pulumi.Input<enums.retail.v2alpha.ProductAvailability>;
    /**
     * The available quantity of the item.
     */
    availableQuantity?: pulumi.Input<number>;
    /**
     * The timestamp when this Product becomes available for SearchService.Search.
     */
    availableTime?: pulumi.Input<string>;
    branchId: pulumi.Input<string>;
    /**
     * The brands of the product. A maximum of 30 brands are allowed. Each brand must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [brand](https://support.google.com/merchants/answer/6324351). Schema.org property [Product.brand](https://schema.org/brand).
     */
    brands?: pulumi.Input<pulumi.Input<string>[]>;
    catalogId: pulumi.Input<string>;
    /**
     * Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, please replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
     */
    categories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the collection members when type is Type.COLLECTION. Should not set it for other types. A maximum of 1000 values are allowed. Otherwise, an INVALID_ARGUMENT error is return.
     */
    collectionMemberIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The color of the product. Google Merchant Center property [color](https://support.google.com/merchants/answer/6324487). Schema.org property [Product.color](https://schema.org/color).
     */
    colorInfo?: pulumi.Input<inputs.retail.v2alpha.GoogleCloudRetailV2alphaColorInfoArgs>;
    /**
     * The condition of the product. Strongly encouraged to use the standard values: "new", "refurbished", "used". A maximum of 5 values are allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [condition](https://support.google.com/merchants/answer/6324469). Schema.org property [Offer.itemCondition](https://schema.org/itemCondition).
     */
    conditions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). schema.org property [Product.description](https://schema.org/description).
     */
    description?: pulumi.Input<string>;
    /**
     * The timestamp when this product becomes unavailable for SearchService.Search. If it is set, the Product is not available for SearchService.Search after expire_time. However, the product can still be retrieved by ProductService.GetProduct and ProductService.ListProducts. expire_time must be later than available_time and publish_time, otherwise an INVALID_ARGUMENT error is thrown. Google Merchant Center property [expiration_date](https://support.google.com/merchants/answer/6324499).
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Fulfillment information, such as the store IDs for in-store pickup or region IDs for different shipping methods. All the elements must have distinct FulfillmentInfo.type. Otherwise, an INVALID_ARGUMENT error is returned.
     */
    fulfillmentInfo?: pulumi.Input<pulumi.Input<inputs.retail.v2alpha.GoogleCloudRetailV2alphaFulfillmentInfoArgs>[]>;
    /**
     * The Global Trade Item Number (GTIN) of the product. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. This field must be a Unigram. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [gtin](https://support.google.com/merchants/answer/6324461). Schema.org property [Product.isbn](https://schema.org/isbn) or [Product.gtin8](https://schema.org/gtin8) or [Product.gtin12](https://schema.org/gtin12) or [Product.gtin13](https://schema.org/gtin13) or [Product.gtin14](https://schema.org/gtin14). If the value is not a valid GTIN, an INVALID_ARGUMENT error is returned.
     */
    gtin?: pulumi.Input<string>;
    /**
     * Immutable. Product identifier, which is the final component of name. For example, this field is "id_1", if name is `projects/*&#47;locations/global/catalogs/default_catalog/branches/default_branch/products/id_1`. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [id](https://support.google.com/merchants/answer/6324405). Schema.org Property [Product.sku](https://schema.org/sku).
     */
    id?: pulumi.Input<string>;
    /**
     * Product images for the product.Highly recommended to put the main image to the first. A maximum of 300 images are allowed. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
     */
    images?: pulumi.Input<pulumi.Input<inputs.retail.v2alpha.GoogleCloudRetailV2alphaImageArgs>[]>;
    /**
     * Language of the title/description and other string attributes. Use language tags defined by BCP 47. For product prediction, this field is ignored and the model automatically detects the text language. The Product can include text in different languages, but duplicating Products to provide text in multiple languages can result in degraded model performance. For product search this field is in use. It defaults to "en-US" if unset.
     */
    languageCode?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The material of the product. For example, "leather", "wooden". A maximum of 20 values are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [material](https://support.google.com/merchants/answer/6324410). Schema.org property [Product.material](https://schema.org/material).
     */
    materials?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Immutable. Full resource name of the product, such as `projects/*&#47;locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`.
     */
    name?: pulumi.Input<string>;
    /**
     * The pattern or graphic print of the product. For example, "striped", "polka dot", "paisley". A maximum of 20 values are allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [pattern](https://support.google.com/merchants/answer/6324483). Schema.org property [Product.pattern](https://schema.org/pattern).
     */
    patterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Product price and cost information. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
     */
    priceInfo?: pulumi.Input<inputs.retail.v2alpha.GoogleCloudRetailV2alphaPriceInfoArgs>;
    /**
     * Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center Property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org Property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID). This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
     */
    primaryProductId?: pulumi.Input<string>;
    productId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The promotions applied to the product. A maximum of 10 values are allowed per Product.
     */
    promotions?: pulumi.Input<pulumi.Input<inputs.retail.v2alpha.GoogleCloudRetailV2alphaPromotionArgs>[]>;
    /**
     * The timestamp when the product is published by the retailer for the first time, which indicates the freshness of the products. Note that this field is different from available_time, given it purely describes product freshness regardless of when it is available on search and recommendation.
     */
    publishTime?: pulumi.Input<string>;
    /**
     * The rating of this product.
     */
    rating?: pulumi.Input<inputs.retail.v2alpha.GoogleCloudRetailV2alphaRatingArgs>;
    /**
     * Indicates which fields in the Products are returned in SearchResponse. Supported fields for all types: * audience * availability * brands * color_info * conditions * gtin * materials * name * patterns * price_info * rating * sizes * title * uri Supported fields only for Type.PRIMARY and Type.COLLECTION: * categories * description * images Supported fields only for Type.VARIANT: * Only the first image in images To mark attributes as retrievable, include paths of the form "attributes.key" where "key" is the key of a custom attribute, as specified in attributes. For Type.PRIMARY and Type.COLLECTION, the following fields are always returned in SearchResponse by default: * name For Type.VARIANT, the following fields are always returned in by default: * name * color_info Maximum number of paths is 30. Otherwise, an INVALID_ARGUMENT error is returned. Note: Returning more fields in SearchResponse may increase response payload size and serving latency.
     */
    retrievableFields?: pulumi.Input<string>;
    /**
     * The size of the product. To represent different size systems or size types, consider using this format: [[[size_system:]size_type:]size_value]. For example, in "US:MENS:M", "US" represents size system; "MENS" represents size type; "M" represents size value. In "GIRLS:27", size system is empty; "GIRLS" represents size type; "27" represents size value. In "32 inches", both size system and size type are empty, while size value is "32 inches". A maximum of 20 values are allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [size](https://support.google.com/merchants/answer/6324492), [size_type](https://support.google.com/merchants/answer/6324497) and [size_system](https://support.google.com/merchants/answer/6324502). Schema.org property [Product.size](https://schema.org/size).
     */
    sizes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
     */
    title: pulumi.Input<string>;
    /**
     * Input only. The TTL (time to live) of the product. If it is set, it must be a non-negative value, and expire_time is set as current timestamp plus ttl. The derived expire_time is returned in the output and ttl is left blank when retrieving the Product. If it is set, the product is not available for SearchService.Search after current timestamp plus ttl. However, the product can still be retrieved by ProductService.GetProduct and ProductService.ListProducts.
     */
    ttl?: pulumi.Input<string>;
    /**
     * Immutable. The type of the product. Default to Catalog.product_level_config.ingestion_product_type if unset.
     */
    type?: pulumi.Input<enums.retail.v2alpha.ProductType>;
    /**
     * Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
     */
    uri?: pulumi.Input<string>;
}
