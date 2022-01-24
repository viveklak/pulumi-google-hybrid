// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a Product.
 */
export function getProduct(args: GetProductArgs, opts?: pulumi.InvokeOptions): Promise<GetProductResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:retail/v2beta:getProduct", {
        "branchId": args.branchId,
        "catalogId": args.catalogId,
        "location": args.location,
        "productId": args.productId,
        "project": args.project,
    }, opts);
}

export interface GetProductArgs {
    branchId: string;
    catalogId: string;
    location: string;
    productId: string;
    project?: string;
}

export interface GetProductResult {
    /**
     * Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 200. * The key must be a UTF-8 encoded string with a length limit of 128 characters. * For indexable attribute, the key must match the pattern: `a-zA-Z0-9*`. For example, `key0LikeThis` or `KEY_1_LIKE_THIS`. * For text attributes, at most 400 values are allowed. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 256 characters. * For number attributes, at most 400 values are allowed.
     */
    readonly attributes: {[key: string]: string};
    /**
     * The target group associated with a given audience (e.g. male, veterans, car owners, musicians, etc.) of the product.
     */
    readonly audience: outputs.retail.v2beta.GoogleCloudRetailV2betaAudienceResponse;
    /**
     * The online availability of the Product. Default to Availability.IN_STOCK. Corresponding properties: Google Merchant Center property [availability](https://support.google.com/merchants/answer/6324448). Schema.org property [Offer.availability](https://schema.org/availability).
     */
    readonly availability: string;
    /**
     * The available quantity of the item.
     */
    readonly availableQuantity: number;
    /**
     * The timestamp when this Product becomes available for SearchService.Search.
     */
    readonly availableTime: string;
    /**
     * The brands of the product. A maximum of 30 brands are allowed. Each brand must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [brand](https://support.google.com/merchants/answer/6324351). Schema.org property [Product.brand](https://schema.org/brand).
     */
    readonly brands: string[];
    /**
     * Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, please replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
     */
    readonly categories: string[];
    /**
     * The id of the collection members when type is Type.COLLECTION. Should not set it for other types. A maximum of 1000 values are allowed. Otherwise, an INVALID_ARGUMENT error is return.
     */
    readonly collectionMemberIds: string[];
    /**
     * The color of the product. Corresponding properties: Google Merchant Center property [color](https://support.google.com/merchants/answer/6324487). Schema.org property [Product.color](https://schema.org/color).
     */
    readonly colorInfo: outputs.retail.v2beta.GoogleCloudRetailV2betaColorInfoResponse;
    /**
     * The condition of the product. Strongly encouraged to use the standard values: "new", "refurbished", "used". A maximum of 1 value is allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [condition](https://support.google.com/merchants/answer/6324469). Schema.org property [Offer.itemCondition](https://schema.org/itemCondition).
     */
    readonly conditions: string[];
    /**
     * Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). Schema.org property [Product.description](https://schema.org/description).
     */
    readonly description: string;
    /**
     * The timestamp when this product becomes unavailable for SearchService.Search. If it is set, the Product is not available for SearchService.Search after expire_time. However, the product can still be retrieved by ProductService.GetProduct and ProductService.ListProducts. expire_time must be later than available_time and publish_time, otherwise an INVALID_ARGUMENT error is thrown. Corresponding properties: Google Merchant Center property [expiration_date](https://support.google.com/merchants/answer/6324499).
     */
    readonly expireTime: string;
    /**
     * Fulfillment information, such as the store IDs for in-store pickup or region IDs for different shipping methods. All the elements must have distinct FulfillmentInfo.type. Otherwise, an INVALID_ARGUMENT error is returned.
     */
    readonly fulfillmentInfo: outputs.retail.v2beta.GoogleCloudRetailV2betaFulfillmentInfoResponse[];
    /**
     * The Global Trade Item Number (GTIN) of the product. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. This field must be a Unigram. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [gtin](https://support.google.com/merchants/answer/6324461). Schema.org property [Product.isbn](https://schema.org/isbn), [Product.gtin8](https://schema.org/gtin8), [Product.gtin12](https://schema.org/gtin12), [Product.gtin13](https://schema.org/gtin13), or [Product.gtin14](https://schema.org/gtin14). If the value is not a valid GTIN, an INVALID_ARGUMENT error is returned.
     */
    readonly gtin: string;
    /**
     * Product images for the product.Highly recommended to put the main image to the first. A maximum of 300 images are allowed. Corresponding properties: Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
     */
    readonly images: outputs.retail.v2beta.GoogleCloudRetailV2betaImageResponse[];
    /**
     * Language of the title/description and other string attributes. Use language tags defined by [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). For product prediction, this field is ignored and the model automatically detects the text language. The Product can include text in different languages, but duplicating Products to provide text in multiple languages can result in degraded model performance. For product search this field is in use. It defaults to "en-US" if unset.
     */
    readonly languageCode: string;
    /**
     * The material of the product. For example, "leather", "wooden". A maximum of 20 values are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [material](https://support.google.com/merchants/answer/6324410). Schema.org property [Product.material](https://schema.org/material).
     */
    readonly materials: string[];
    /**
     * Immutable. Full resource name of the product, such as `projects/*&#47;locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`.
     */
    readonly name: string;
    /**
     * The pattern or graphic print of the product. For example, "striped", "polka dot", "paisley". A maximum of 20 values are allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [pattern](https://support.google.com/merchants/answer/6324483). Schema.org property [Product.pattern](https://schema.org/pattern).
     */
    readonly patterns: string[];
    /**
     * Product price and cost information. Corresponding properties: Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
     */
    readonly priceInfo: outputs.retail.v2beta.GoogleCloudRetailV2betaPriceInfoResponse;
    /**
     * Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID).
     */
    readonly primaryProductId: string;
    /**
     * The promotions applied to the product. A maximum of 10 values are allowed per Product.
     */
    readonly promotions: outputs.retail.v2beta.GoogleCloudRetailV2betaPromotionResponse[];
    /**
     * The timestamp when the product is published by the retailer for the first time, which indicates the freshness of the products. Note that this field is different from available_time, given it purely describes product freshness regardless of when it is available on search and recommendation.
     */
    readonly publishTime: string;
    /**
     * The rating of this product.
     */
    readonly rating: outputs.retail.v2beta.GoogleCloudRetailV2betaRatingResponse;
    /**
     * Indicates which fields in the Products are returned in SearchResponse. Supported fields for all types: * audience * availability * brands * color_info * conditions * gtin * materials * name * patterns * price_info * rating * sizes * title * uri Supported fields only for Type.PRIMARY and Type.COLLECTION: * categories * description * images Supported fields only for Type.VARIANT: * Only the first image in images To mark attributes as retrievable, include paths of the form "attributes.key" where "key" is the key of a custom attribute, as specified in attributes. For Type.PRIMARY and Type.COLLECTION, the following fields are always returned in SearchResponse by default: * name For Type.VARIANT, the following fields are always returned in by default: * name * color_info Maximum number of paths is 30. Otherwise, an INVALID_ARGUMENT error is returned. Note: Returning more fields in SearchResponse may increase response payload size and serving latency.
     */
    readonly retrievableFields: string;
    /**
     * The size of the product. To represent different size systems or size types, consider using this format: [[[size_system:]size_type:]size_value]. For example, in "US:MENS:M", "US" represents size system; "MENS" represents size type; "M" represents size value. In "GIRLS:27", size system is empty; "GIRLS" represents size type; "27" represents size value. In "32 inches", both size system and size type are empty, while size value is "32 inches". A maximum of 20 values are allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [size](https://support.google.com/merchants/answer/6324492), [size_type](https://support.google.com/merchants/answer/6324497), and [size_system](https://support.google.com/merchants/answer/6324502). Schema.org property [Product.size](https://schema.org/size).
     */
    readonly sizes: string[];
    /**
     * Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Corresponding properties: Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
     */
    readonly tags: string[];
    /**
     * Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
     */
    readonly title: string;
    /**
     * Input only. The TTL (time to live) of the product. If it is set, it must be a non-negative value, and expire_time is set as current timestamp plus ttl. The derived expire_time is returned in the output and ttl is left blank when retrieving the Product. If it is set, the product is not available for SearchService.Search after current timestamp plus ttl. However, the product can still be retrieved by ProductService.GetProduct and ProductService.ListProducts.
     */
    readonly ttl: string;
    /**
     * Immutable. The type of the product. Default to Catalog.product_level_config.ingestion_product_type if unset.
     */
    readonly type: string;
    /**
     * Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
     */
    readonly uri: string;
    /**
     * Product variants grouped together on primary product which share similar product attributes. It's automatically grouped by primary_product_id for all the product variants. Only populated for Type.PRIMARY Products. Note: This field is OUTPUT_ONLY for ProductService.GetProduct. Do not set this field in API requests.
     */
    readonly variants: outputs.retail.v2beta.GoogleCloudRetailV2betaProductResponse[];
}

export function getProductOutput(args: GetProductOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProductResult> {
    return pulumi.output(args).apply(a => getProduct(a, opts))
}

export interface GetProductOutputArgs {
    branchId: pulumi.Input<string>;
    catalogId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    productId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
