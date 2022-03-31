// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    /// <summary>
    /// Represents an amount of money with its currency type.
    /// </summary>
    [OutputType]
    public sealed class GoogleTypeMoneyResponse
    {
        /// <summary>
        /// The three-letter currency code defined in ISO 4217.
        /// </summary>
        public readonly string CurrencyCode;
        /// <summary>
        /// Number of nano (10^-9) units of the amount. The value must be between -999,999,999 and +999,999,999 inclusive. If `units` is positive, `nanos` must be positive or zero. If `units` is zero, `nanos` can be positive, zero, or negative. If `units` is negative, `nanos` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
        /// </summary>
        public readonly int Nanos;
        /// <summary>
        /// The whole units of the amount. For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
        /// </summary>
        public readonly string Units;

        [OutputConstructor]
        private GoogleTypeMoneyResponse(
            string currencyCode,

            int nanos,

            string units)
        {
            CurrencyCode = currencyCode;
            Nanos = nanos;
            Units = units;
        }
    }
}
