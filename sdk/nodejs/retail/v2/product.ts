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
    public static readonly __pulumiType = 'google-native:retail/v2:Product';

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
     * Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 150 by default; 100 for Type.VARIANT. * The key must be a UTF-8 encoded string with a length limit of 128 characters.
     */
    public readonly attributes!: pulumi.Output<{[key: string]: string}>;
    /**
     * The online availability of the Product. Default to Availability.IN_STOCK. Google Merchant Center Property [availability](https://support.google.com/merchants/answer/6324448). Schema.org Property [Offer.availability](https://schema.org/availability).
     */
    public readonly availability!: pulumi.Output<string>;
    /**
     * The available quantity of the item.
     */
    public readonly availableQuantity!: pulumi.Output<number>;
    /**
     * The timestamp when this Product becomes available for recommendation.
     */
    public readonly availableTime!: pulumi.Output<string>;
    /**
     * Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, please replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
     */
    public readonly categories!: pulumi.Output<string[]>;
    /**
     * Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). schema.org property [Product.description](https://schema.org/description).
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Product images for the product.Highly recommended to put the main image to the first. A maximum of 300 images are allowed. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
     */
    public readonly images!: pulumi.Output<outputs.retail.v2.GoogleCloudRetailV2ImageResponse[]>;
    /**
     * Immutable. Full resource name of the product, such as `projects/*&#47;locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`. The branch ID must be "default_branch".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Product price and cost information. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
     */
    public readonly priceInfo!: pulumi.Output<outputs.retail.v2.GoogleCloudRetailV2PriceInfoResponse>;
    /**
     * Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center Property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org Property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID). This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
     */
    public readonly primaryProductId!: pulumi.Output<string>;
    /**
     * Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * Required. Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Immutable. The type of the product. This field is output-only. Default to Catalog.product_level_config.ingestion_product_type if unset.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
     */
    public readonly uri!: pulumi.Output<string>;

    /**
     * Create a Product resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.branchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'branchId'");
            }
            if ((!args || args.catalogId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalogId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.productId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["availability"] = args ? args.availability : undefined;
            inputs["availableQuantity"] = args ? args.availableQuantity : undefined;
            inputs["availableTime"] = args ? args.availableTime : undefined;
            inputs["branchId"] = args ? args.branchId : undefined;
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["categories"] = args ? args.categories : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["images"] = args ? args.images : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["priceInfo"] = args ? args.priceInfo : undefined;
            inputs["primaryProductId"] = args ? args.primaryProductId : undefined;
            inputs["productId"] = args ? args.productId : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["uri"] = args ? args.uri : undefined;
        } else {
            inputs["attributes"] = undefined /*out*/;
            inputs["availability"] = undefined /*out*/;
            inputs["availableQuantity"] = undefined /*out*/;
            inputs["availableTime"] = undefined /*out*/;
            inputs["categories"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["images"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["priceInfo"] = undefined /*out*/;
            inputs["primaryProductId"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["title"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Product.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Product resource.
 */
export interface ProductArgs {
    /**
     * Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 150 by default; 100 for Type.VARIANT. * The key must be a UTF-8 encoded string with a length limit of 128 characters.
     */
    attributes?: pulumi.Input<{[key: string]: string}>;
    /**
     * The online availability of the Product. Default to Availability.IN_STOCK. Google Merchant Center Property [availability](https://support.google.com/merchants/answer/6324448). Schema.org Property [Offer.availability](https://schema.org/availability).
     */
    availability?: pulumi.Input<enums.retail.v2.ProductAvailability>;
    /**
     * The available quantity of the item.
     */
    availableQuantity?: pulumi.Input<number>;
    /**
     * The timestamp when this Product becomes available for recommendation.
     */
    availableTime?: pulumi.Input<string>;
    branchId: pulumi.Input<string>;
    catalogId: pulumi.Input<string>;
    /**
     * Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, please replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
     */
    categories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). schema.org property [Product.description](https://schema.org/description).
     */
    description?: pulumi.Input<string>;
    /**
     * Immutable. Product identifier, which is the final component of name. For example, this field is "id_1", if name is `projects/*&#47;locations/global/catalogs/default_catalog/branches/default_branch/products/id_1`. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [id](https://support.google.com/merchants/answer/6324405). Schema.org Property [Product.sku](https://schema.org/sku).
     */
    id?: pulumi.Input<string>;
    /**
     * Product images for the product.Highly recommended to put the main image to the first. A maximum of 300 images are allowed. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
     */
    images?: pulumi.Input<pulumi.Input<inputs.retail.v2.GoogleCloudRetailV2ImageArgs>[]>;
    location: pulumi.Input<string>;
    /**
     * Immutable. Full resource name of the product, such as `projects/*&#47;locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`. The branch ID must be "default_branch".
     */
    name?: pulumi.Input<string>;
    /**
     * Product price and cost information. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
     */
    priceInfo?: pulumi.Input<inputs.retail.v2.GoogleCloudRetailV2PriceInfoArgs>;
    /**
     * Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center Property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org Property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID). This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
     */
    primaryProductId?: pulumi.Input<string>;
    productId: pulumi.Input<string>;
    project: pulumi.Input<string>;
    /**
     * Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
     */
    title?: pulumi.Input<string>;
    /**
     * Immutable. The type of the product. This field is output-only. Default to Catalog.product_level_config.ingestion_product_type if unset.
     */
    type?: pulumi.Input<enums.retail.v2.ProductType>;
    /**
     * Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
     */
    uri?: pulumi.Input<string>;
}
