// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Billingbudgets.V1.Outputs
{

    /// <summary>
    /// The budgeted amount for each usage period.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudBillingBudgetsV1BudgetAmountResponse
    {
        /// <summary>
        /// Use the last period's actual spend as the budget for the present period. LastPeriodAmount can only be set when the budget's time period is a Filter.calendar_period. It cannot be set in combination with Filter.custom_period.
        /// </summary>
        public readonly Outputs.GoogleCloudBillingBudgetsV1LastPeriodAmountResponse LastPeriodAmount;
        /// <summary>
        /// A specified amount to use as the budget. `currency_code` is optional. If specified when creating a budget, it must match the currency of the billing account. If specified when updating a budget, it must match the currency_code of the existing budget. The `currency_code` is provided on output.
        /// </summary>
        public readonly Outputs.GoogleTypeMoneyResponse SpecifiedAmount;

        [OutputConstructor]
        private GoogleCloudBillingBudgetsV1BudgetAmountResponse(
            Outputs.GoogleCloudBillingBudgetsV1LastPeriodAmountResponse lastPeriodAmount,

            Outputs.GoogleTypeMoneyResponse specifiedAmount)
        {
            LastPeriodAmount = lastPeriodAmount;
            SpecifiedAmount = specifiedAmount;
        }
    }
}
