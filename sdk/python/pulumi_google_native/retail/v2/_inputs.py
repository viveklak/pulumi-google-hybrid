# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'GoogleCloudRetailV2AudienceArgs',
    'GoogleCloudRetailV2ColorInfoArgs',
    'GoogleCloudRetailV2FulfillmentInfoArgs',
    'GoogleCloudRetailV2ImageArgs',
    'GoogleCloudRetailV2PriceInfoArgs',
    'GoogleCloudRetailV2PromotionArgs',
    'GoogleCloudRetailV2RatingArgs',
]

@pulumi.input_type
class GoogleCloudRetailV2AudienceArgs:
    def __init__(__self__, *,
                 age_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 genders: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        An intended audience of the Product for whom it's sold.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] age_groups: The age groups of the audience. Strongly encouraged to use the standard values: "newborn" (up to 3 months old), "infant" (3–12 months old), "toddler" (1–5 years old), "kids" (5–13 years old), "adult" (typically teens or older). At most 5 values are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [age_group](https://support.google.com/merchants/answer/6324463). Schema.org property [Product.audience.suggestedMinAge](https://schema.org/suggestedMinAge) and [Product.audience.suggestedMaxAge](https://schema.org/suggestedMaxAge).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] genders: The genders of the audience. Strongly encouraged to use the standard values: "male", "female", "unisex". At most 5 values are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [gender](https://support.google.com/merchants/answer/6324479). Schema.org property [Product.audience.suggestedGender](https://schema.org/suggestedGender).
        """
        if age_groups is not None:
            pulumi.set(__self__, "age_groups", age_groups)
        if genders is not None:
            pulumi.set(__self__, "genders", genders)

    @property
    @pulumi.getter(name="ageGroups")
    def age_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The age groups of the audience. Strongly encouraged to use the standard values: "newborn" (up to 3 months old), "infant" (3–12 months old), "toddler" (1–5 years old), "kids" (5–13 years old), "adult" (typically teens or older). At most 5 values are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [age_group](https://support.google.com/merchants/answer/6324463). Schema.org property [Product.audience.suggestedMinAge](https://schema.org/suggestedMinAge) and [Product.audience.suggestedMaxAge](https://schema.org/suggestedMaxAge).
        """
        return pulumi.get(self, "age_groups")

    @age_groups.setter
    def age_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "age_groups", value)

    @property
    @pulumi.getter
    def genders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The genders of the audience. Strongly encouraged to use the standard values: "male", "female", "unisex". At most 5 values are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [gender](https://support.google.com/merchants/answer/6324479). Schema.org property [Product.audience.suggestedGender](https://schema.org/suggestedGender).
        """
        return pulumi.get(self, "genders")

    @genders.setter
    def genders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "genders", value)


@pulumi.input_type
class GoogleCloudRetailV2ColorInfoArgs:
    def __init__(__self__, *,
                 color_families: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 colors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The color information of a Product.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] color_families: The standard color families. Strongly recommended to use the following standard color groups: "Red", "Pink", "Orange", "Yellow", "Purple", "Green", "Cyan", "Blue", "Brown", "White", "Gray", "Black" and "Mixed". Normally it is expected to have only 1 color family. May consider using single "Mixed" instead of multiple values. A maximum of 5 values are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [color](https://support.google.com/merchants/answer/6324487). Schema.org property [Product.color](https://schema.org/color).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] colors: The color display names, which may be different from standard color family names, such as the color aliases used in the website frontend. Normally it is expected to have only 1 color. May consider using single "Mixed" instead of multiple values. A maximum of 25 colors are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [color](https://support.google.com/merchants/answer/6324487). Schema.org property [Product.color](https://schema.org/color).
        """
        if color_families is not None:
            pulumi.set(__self__, "color_families", color_families)
        if colors is not None:
            pulumi.set(__self__, "colors", colors)

    @property
    @pulumi.getter(name="colorFamilies")
    def color_families(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The standard color families. Strongly recommended to use the following standard color groups: "Red", "Pink", "Orange", "Yellow", "Purple", "Green", "Cyan", "Blue", "Brown", "White", "Gray", "Black" and "Mixed". Normally it is expected to have only 1 color family. May consider using single "Mixed" instead of multiple values. A maximum of 5 values are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [color](https://support.google.com/merchants/answer/6324487). Schema.org property [Product.color](https://schema.org/color).
        """
        return pulumi.get(self, "color_families")

    @color_families.setter
    def color_families(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "color_families", value)

    @property
    @pulumi.getter
    def colors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The color display names, which may be different from standard color family names, such as the color aliases used in the website frontend. Normally it is expected to have only 1 color. May consider using single "Mixed" instead of multiple values. A maximum of 25 colors are allowed. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [color](https://support.google.com/merchants/answer/6324487). Schema.org property [Product.color](https://schema.org/color).
        """
        return pulumi.get(self, "colors")

    @colors.setter
    def colors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "colors", value)


@pulumi.input_type
class GoogleCloudRetailV2FulfillmentInfoArgs:
    def __init__(__self__, *,
                 place_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Fulfillment information, such as the store IDs for in-store pickup or region IDs for different shipping methods.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] place_ids: The IDs for this type, such as the store IDs for FulfillmentInfo.type.pickup-in-store or the region IDs for FulfillmentInfo.type.same-day-delivery. A maximum of 3000 values are allowed. Each value must be a string with a length limit of 30 characters, matching the pattern `[a-zA-Z0-9_-]+`, such as "store1" or "REGION-2". Otherwise, an INVALID_ARGUMENT error is returned.
        :param pulumi.Input[str] type: The fulfillment type, including commonly used types (such as pickup in store and same day delivery), and custom types. Customers have to map custom types to their display names before rendering UI. Supported values: * "pickup-in-store" * "ship-to-store" * "same-day-delivery" * "next-day-delivery" * "custom-type-1" * "custom-type-2" * "custom-type-3" * "custom-type-4" * "custom-type-5" If this field is set to an invalid value other than these, an INVALID_ARGUMENT error is returned.
        """
        if place_ids is not None:
            pulumi.set(__self__, "place_ids", place_ids)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="placeIds")
    def place_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs for this type, such as the store IDs for FulfillmentInfo.type.pickup-in-store or the region IDs for FulfillmentInfo.type.same-day-delivery. A maximum of 3000 values are allowed. Each value must be a string with a length limit of 30 characters, matching the pattern `[a-zA-Z0-9_-]+`, such as "store1" or "REGION-2". Otherwise, an INVALID_ARGUMENT error is returned.
        """
        return pulumi.get(self, "place_ids")

    @place_ids.setter
    def place_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "place_ids", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The fulfillment type, including commonly used types (such as pickup in store and same day delivery), and custom types. Customers have to map custom types to their display names before rendering UI. Supported values: * "pickup-in-store" * "ship-to-store" * "same-day-delivery" * "next-day-delivery" * "custom-type-1" * "custom-type-2" * "custom-type-3" * "custom-type-4" * "custom-type-5" If this field is set to an invalid value other than these, an INVALID_ARGUMENT error is returned.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class GoogleCloudRetailV2ImageArgs:
    def __init__(__self__, *,
                 uri: pulumi.Input[str],
                 height: Optional[pulumi.Input[int]] = None,
                 width: Optional[pulumi.Input[int]] = None):
        """
        Product image. Recommendations AI and Retail Search do not use product images to improve prediction and search results. However, product images can be returned in results, and are shown in prediction or search previews in the console.
        :param pulumi.Input[str] uri: URI of the image. This field must be a valid UTF-8 encoded URI with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
        :param pulumi.Input[int] height: Height of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
        :param pulumi.Input[int] width: Width of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
        """
        pulumi.set(__self__, "uri", uri)
        if height is not None:
            pulumi.set(__self__, "height", height)
        if width is not None:
            pulumi.set(__self__, "width", width)

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Input[str]:
        """
        URI of the image. This field must be a valid UTF-8 encoded URI with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "uri", value)

    @property
    @pulumi.getter
    def height(self) -> Optional[pulumi.Input[int]]:
        """
        Height of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
        """
        return pulumi.get(self, "height")

    @height.setter
    def height(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "height", value)

    @property
    @pulumi.getter
    def width(self) -> Optional[pulumi.Input[int]]:
        """
        Width of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
        """
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "width", value)


@pulumi.input_type
class GoogleCloudRetailV2PriceInfoArgs:
    def __init__(__self__, *,
                 cost: Optional[pulumi.Input[float]] = None,
                 currency_code: Optional[pulumi.Input[str]] = None,
                 original_price: Optional[pulumi.Input[float]] = None,
                 price: Optional[pulumi.Input[float]] = None,
                 price_effective_time: Optional[pulumi.Input[str]] = None,
                 price_expire_time: Optional[pulumi.Input[str]] = None):
        """
        The price information of a Product.
        :param pulumi.Input[float] cost: The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
        :param pulumi.Input[str] currency_code: The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned. The Product.Type.VARIANT Products with the same Product.primary_product_id must share the same currency_code. Otherwise, a FAILED_PRECONDITION error is returned.
        :param pulumi.Input[float] original_price: Price of the product without any discount. If zero, by default set to be the price.
        :param pulumi.Input[float] price: Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.price](https://schema.org/price).
        :param pulumi.Input[str] price_effective_time: The timestamp when the price starts to be effective. This can be set as a future timestamp, and the price is only used for search after price_effective_time. If so, the original_price must be set and original_price is used before price_effective_time. Do not set if price is always effective because it will cause additional latency during search.
        :param pulumi.Input[str] price_expire_time: The timestamp when the price stops to be effective. The price is used for search before price_expire_time. If this field is set, the original_price must be set and original_price is used after price_expire_time. Do not set if price is always effective because it will cause additional latency during search.
        """
        if cost is not None:
            pulumi.set(__self__, "cost", cost)
        if currency_code is not None:
            pulumi.set(__self__, "currency_code", currency_code)
        if original_price is not None:
            pulumi.set(__self__, "original_price", original_price)
        if price is not None:
            pulumi.set(__self__, "price", price)
        if price_effective_time is not None:
            pulumi.set(__self__, "price_effective_time", price_effective_time)
        if price_expire_time is not None:
            pulumi.set(__self__, "price_expire_time", price_expire_time)

    @property
    @pulumi.getter
    def cost(self) -> Optional[pulumi.Input[float]]:
        """
        The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
        """
        return pulumi.get(self, "cost")

    @cost.setter
    def cost(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "cost", value)

    @property
    @pulumi.getter(name="currencyCode")
    def currency_code(self) -> Optional[pulumi.Input[str]]:
        """
        The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned. The Product.Type.VARIANT Products with the same Product.primary_product_id must share the same currency_code. Otherwise, a FAILED_PRECONDITION error is returned.
        """
        return pulumi.get(self, "currency_code")

    @currency_code.setter
    def currency_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "currency_code", value)

    @property
    @pulumi.getter(name="originalPrice")
    def original_price(self) -> Optional[pulumi.Input[float]]:
        """
        Price of the product without any discount. If zero, by default set to be the price.
        """
        return pulumi.get(self, "original_price")

    @original_price.setter
    def original_price(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "original_price", value)

    @property
    @pulumi.getter
    def price(self) -> Optional[pulumi.Input[float]]:
        """
        Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.price](https://schema.org/price).
        """
        return pulumi.get(self, "price")

    @price.setter
    def price(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "price", value)

    @property
    @pulumi.getter(name="priceEffectiveTime")
    def price_effective_time(self) -> Optional[pulumi.Input[str]]:
        """
        The timestamp when the price starts to be effective. This can be set as a future timestamp, and the price is only used for search after price_effective_time. If so, the original_price must be set and original_price is used before price_effective_time. Do not set if price is always effective because it will cause additional latency during search.
        """
        return pulumi.get(self, "price_effective_time")

    @price_effective_time.setter
    def price_effective_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "price_effective_time", value)

    @property
    @pulumi.getter(name="priceExpireTime")
    def price_expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        The timestamp when the price stops to be effective. The price is used for search before price_expire_time. If this field is set, the original_price must be set and original_price is used after price_expire_time. Do not set if price is always effective because it will cause additional latency during search.
        """
        return pulumi.get(self, "price_expire_time")

    @price_expire_time.setter
    def price_expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "price_expire_time", value)


@pulumi.input_type
class GoogleCloudRetailV2PromotionArgs:
    def __init__(__self__, *,
                 promotion_id: Optional[pulumi.Input[str]] = None):
        """
        Promotion information.
        :param pulumi.Input[str] promotion_id: ID of the promotion. For example, "free gift". The value must be a UTF-8 encoded string with a length limit of 128 characters, and match the pattern: `a-zA-Z*`. For example, id0LikeThis or ID_1_LIKE_THIS. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [promotion](https://support.google.com/merchants/answer/7050148).
        """
        if promotion_id is not None:
            pulumi.set(__self__, "promotion_id", promotion_id)

    @property
    @pulumi.getter(name="promotionId")
    def promotion_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the promotion. For example, "free gift". The value must be a UTF-8 encoded string with a length limit of 128 characters, and match the pattern: `a-zA-Z*`. For example, id0LikeThis or ID_1_LIKE_THIS. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [promotion](https://support.google.com/merchants/answer/7050148).
        """
        return pulumi.get(self, "promotion_id")

    @promotion_id.setter
    def promotion_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "promotion_id", value)


@pulumi.input_type
class GoogleCloudRetailV2RatingArgs:
    def __init__(__self__, *,
                 average_rating: Optional[pulumi.Input[float]] = None,
                 rating_count: Optional[pulumi.Input[int]] = None,
                 rating_histogram: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None):
        """
        The rating of a Product.
        :param pulumi.Input[float] average_rating: The average rating of the Product. The rating is scaled at 1-5. Otherwise, an INVALID_ARGUMENT error is returned.
        :param pulumi.Input[int] rating_count: The total number of ratings. This value is independent of the value of rating_histogram. This value must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] rating_histogram: List of rating counts per rating value (index = rating - 1). The list is empty if there is no rating. If the list is non-empty, its size is always 5. Otherwise, an INVALID_ARGUMENT error is returned. For example, [41, 14, 13, 47, 303]. It means that the Product got 41 ratings with 1 star, 14 ratings with 2 star, and so on.
        """
        if average_rating is not None:
            pulumi.set(__self__, "average_rating", average_rating)
        if rating_count is not None:
            pulumi.set(__self__, "rating_count", rating_count)
        if rating_histogram is not None:
            pulumi.set(__self__, "rating_histogram", rating_histogram)

    @property
    @pulumi.getter(name="averageRating")
    def average_rating(self) -> Optional[pulumi.Input[float]]:
        """
        The average rating of the Product. The rating is scaled at 1-5. Otherwise, an INVALID_ARGUMENT error is returned.
        """
        return pulumi.get(self, "average_rating")

    @average_rating.setter
    def average_rating(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "average_rating", value)

    @property
    @pulumi.getter(name="ratingCount")
    def rating_count(self) -> Optional[pulumi.Input[int]]:
        """
        The total number of ratings. This value is independent of the value of rating_histogram. This value must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
        """
        return pulumi.get(self, "rating_count")

    @rating_count.setter
    def rating_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rating_count", value)

    @property
    @pulumi.getter(name="ratingHistogram")
    def rating_histogram(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        List of rating counts per rating value (index = rating - 1). The list is empty if there is no rating. If the list is non-empty, its size is always 5. Otherwise, an INVALID_ARGUMENT error is returned. For example, [41, 14, 13, 47, 303]. It means that the Product got 41 ratings with 1 star, 14 ratings with 2 star, and so on.
        """
        return pulumi.get(self, "rating_histogram")

    @rating_histogram.setter
    def rating_histogram(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "rating_histogram", value)


