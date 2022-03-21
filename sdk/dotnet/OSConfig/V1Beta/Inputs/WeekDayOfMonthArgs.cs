// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Inputs
{

    /// <summary>
    /// Represents one week day in a month. An example is "the 4th Sunday".
    /// </summary>
    public sealed class WeekDayOfMonthArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A day of the week.
        /// </summary>
        [Input("dayOfWeek", required: true)]
        public Input<Pulumi.GoogleNative.OSConfig.V1Beta.WeekDayOfMonthDayOfWeek> DayOfWeek { get; set; } = null!;

        /// <summary>
        /// Optional. Represents the number of days before or after the given week day of month that the patch deployment is scheduled for. For example if `week_ordinal` and `day_of_week` values point to the second day of the month and this `day_offset` value is set to `3`, the patch deployment takes place three days after the second Tuesday of the month. If this value is negative, for example -5, the patches are deployed five days before before the second Tuesday of the month. Allowed values are in range [-30, 30].
        /// </summary>
        [Input("dayOffset")]
        public Input<int>? DayOffset { get; set; }

        /// <summary>
        /// Week number in a month. 1-4 indicates the 1st to 4th week of the month. -1 indicates the last week of the month.
        /// </summary>
        [Input("weekOrdinal", required: true)]
        public Input<int> WeekOrdinal { get; set; } = null!;

        public WeekDayOfMonthArgs()
        {
        }
    }
}
