// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Beta.Inputs
{

    /// <summary>
    /// The price information of a Product.
    /// </summary>
    public sealed class GoogleCloudRetailV2betaPriceInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
        /// </summary>
        [Input("cost")]
        public Input<double>? Cost { get; set; }

        /// <summary>
        /// The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned. The Product.Type.VARIANT Products with the same Product.primary_product_id must share the same currency_code. Otherwise, a FAILED_PRECONDITION error is returned.
        /// </summary>
        [Input("currencyCode")]
        public Input<string>? CurrencyCode { get; set; }

        /// <summary>
        /// Price of the product without any discount. If zero, by default set to be the price.
        /// </summary>
        [Input("originalPrice")]
        public Input<double>? OriginalPrice { get; set; }

        /// <summary>
        /// Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.price](https://schema.org/price).
        /// </summary>
        [Input("price")]
        public Input<double>? Price { get; set; }

        /// <summary>
        /// The timestamp when the price starts to be effective. This can be set as a future timestamp, and the price is only used for search after price_effective_time. If so, the original_price must be set and original_price is used before price_effective_time. Do not set if price is always effective because it will cause additional latency during search.
        /// </summary>
        [Input("priceEffectiveTime")]
        public Input<string>? PriceEffectiveTime { get; set; }

        /// <summary>
        /// The timestamp when the price stops to be effective. The price is used for search before price_expire_time. If this field is set, the original_price must be set and original_price is used after price_expire_time. Do not set if price is always effective because it will cause additional latency during search.
        /// </summary>
        [Input("priceExpireTime")]
        public Input<string>? PriceExpireTime { get; set; }

        public GoogleCloudRetailV2betaPriceInfoArgs()
        {
        }
    }
}
