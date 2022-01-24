// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Billingbudgets.V1Beta1
{
    public static class GetBudget
    {
        /// <summary>
        /// Returns a budget. WARNING: There are some fields exposed on the Google Cloud Console that aren't available on this API. When reading from the API, you will not see these fields in the return value, though they may have been set in the Cloud Console.
        /// </summary>
        public static Task<GetBudgetResult> InvokeAsync(GetBudgetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBudgetResult>("google-native:billingbudgets/v1beta1:getBudget", args ?? new GetBudgetArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a budget. WARNING: There are some fields exposed on the Google Cloud Console that aren't available on this API. When reading from the API, you will not see these fields in the return value, though they may have been set in the Cloud Console.
        /// </summary>
        public static Output<GetBudgetResult> Invoke(GetBudgetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBudgetResult>("google-native:billingbudgets/v1beta1:getBudget", args ?? new GetBudgetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBudgetArgs : Pulumi.InvokeArgs
    {
        [Input("billingAccountId", required: true)]
        public string BillingAccountId { get; set; } = null!;

        [Input("budgetId", required: true)]
        public string BudgetId { get; set; } = null!;

        public GetBudgetArgs()
        {
        }
    }

    public sealed class GetBudgetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("billingAccountId", required: true)]
        public Input<string> BillingAccountId { get; set; } = null!;

        [Input("budgetId", required: true)]
        public Input<string> BudgetId { get; set; } = null!;

        public GetBudgetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBudgetResult
    {
        /// <summary>
        /// Optional. Rules to apply to notifications sent based on budget spend and thresholds.
        /// </summary>
        public readonly Outputs.GoogleCloudBillingBudgetsV1beta1AllUpdatesRuleResponse AllUpdatesRule;
        /// <summary>
        /// Budgeted amount.
        /// </summary>
        public readonly Outputs.GoogleCloudBillingBudgetsV1beta1BudgetAmountResponse Amount;
        /// <summary>
        /// Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
        /// </summary>
        public readonly Outputs.GoogleCloudBillingBudgetsV1beta1FilterResponse BudgetFilter;
        /// <summary>
        /// User data for display name in UI. Validation: &lt;= 60 chars.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget. Optional for `pubsubTopic` notifications. Required if using email notifications.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudBillingBudgetsV1beta1ThresholdRuleResponse> ThresholdRules;

        [OutputConstructor]
        private GetBudgetResult(
            Outputs.GoogleCloudBillingBudgetsV1beta1AllUpdatesRuleResponse allUpdatesRule,

            Outputs.GoogleCloudBillingBudgetsV1beta1BudgetAmountResponse amount,

            Outputs.GoogleCloudBillingBudgetsV1beta1FilterResponse budgetFilter,

            string displayName,

            string etag,

            string name,

            ImmutableArray<Outputs.GoogleCloudBillingBudgetsV1beta1ThresholdRuleResponse> thresholdRules)
        {
            AllUpdatesRule = allUpdatesRule;
            Amount = amount;
            BudgetFilter = budgetFilter;
            DisplayName = displayName;
            Etag = etag;
            Name = name;
            ThresholdRules = thresholdRules;
        }
    }
}
