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
    /// All date times begin at 12 AM US and Canadian Pacific Time (UTC-8).
    /// </summary>
    public sealed class GoogleCloudBillingBudgetsV1CustomPeriodArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The end date of the time period. Budgets with elapsed end date won't be processed. If unset, specifies to track all usage incurred since the start_date.
        /// </summary>
        [Input("endDate")]
        public Input<Inputs.GoogleTypeDateArgs>? EndDate { get; set; }

        /// <summary>
        /// The start date must be after January 1, 2017.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<Inputs.GoogleTypeDateArgs> StartDate { get; set; } = null!;

        public GoogleCloudBillingBudgetsV1CustomPeriodArgs()
        {
        }
    }
}
