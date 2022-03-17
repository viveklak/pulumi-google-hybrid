// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Billingbudgets.V1.Inputs
{

    /// <summary>
    /// Represents an amount of money with its currency type.
    /// </summary>
    public sealed class GoogleTypeMoneyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The three-letter currency code defined in ISO 4217.
        /// </summary>
        [Input("currencyCode")]
        public Input<string>? CurrencyCode { get; set; }

        /// <summary>
        /// Number of nano (10^-9) units of the amount. The value must be between -999,999,999 and +999,999,999 inclusive. If `units` is positive, `nanos` must be positive or zero. If `units` is zero, `nanos` can be positive, zero, or negative. If `units` is negative, `nanos` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
        /// </summary>
        [Input("nanos")]
        public Input<int>? Nanos { get; set; }

        /// <summary>
        /// The whole units of the amount. For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
        /// </summary>
        [Input("units")]
        public Input<string>? Units { get; set; }

        public GoogleTypeMoneyArgs()
        {
        }
    }
}
