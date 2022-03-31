// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.PolicySimulator.V1.Outputs
{

    /// <summary>
    /// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
    /// </summary>
    [OutputType]
    public sealed class GoogleTypeDateResponse
    {
        /// <summary>
        /// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
        /// </summary>
        public readonly int Day;
        /// <summary>
        /// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
        /// </summary>
        public readonly int Month;
        /// <summary>
        /// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
        /// </summary>
        public readonly int Year;

        [OutputConstructor]
        private GoogleTypeDateResponse(
            int day,

            int month,

            int year)
        {
            Day = day;
            Month = month;
            Year = year;
        }
    }
}
