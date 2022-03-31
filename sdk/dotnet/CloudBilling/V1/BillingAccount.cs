// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBilling.V1
{
    /// <summary>
    /// This method creates [billing subaccounts](https://cloud.google.com/billing/docs/concepts#subaccounts). Google Cloud resellers should use the Channel Services APIs, [accounts.customers.create](https://cloud.google.com/channel/docs/reference/rest/v1/accounts.customers/create) and [accounts.customers.entitlements.create](https://cloud.google.com/channel/docs/reference/rest/v1/accounts.customers.entitlements/create). When creating a subaccount, the current authenticated user must have the `billing.accounts.update` IAM permission on the parent account, which is typically given to billing account [administrators](https://cloud.google.com/billing/docs/how-to/billing-access). This method will return an error if the parent account has not been provisioned as a reseller account.
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudbilling/v1:BillingAccount")]
    public partial class BillingAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
        /// </summary>
        [Output("masterBillingAccount")]
        public Output<string> MasterBillingAccount { get; private set; } = null!;

        /// <summary>
        /// The resource name of the billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF` would be the resource name for billing account `012345-567890-ABCDEF`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// True if the billing account is open, and will therefore be charged for any usage on associated projects. False if the billing account is closed, and therefore projects associated with it will be unable to use paid services.
        /// </summary>
        [Output("open")]
        public Output<bool> Open { get; private set; } = null!;


        /// <summary>
        /// Create a BillingAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BillingAccount(string name, BillingAccountArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:cloudbilling/v1:BillingAccount", name, args ?? new BillingAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BillingAccount(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudbilling/v1:BillingAccount", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BillingAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BillingAccount Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BillingAccount(name, id, options);
        }
    }

    public sealed class BillingAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
        /// </summary>
        [Input("masterBillingAccount")]
        public Input<string>? MasterBillingAccount { get; set; }

        public BillingAccountArgs()
        {
        }
    }
}
