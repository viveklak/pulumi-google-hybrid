// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudChannel.V1
{
    /// <summary>
    /// Creates a new Customer resource under the reseller or distributor account. Possible error codes: * PERMISSION_DENIED: The reseller account making the request is different from the reseller account in the API request. * INVALID_ARGUMENT: * Required request parameters are missing or invalid. * Domain field value doesn't match the primary email domain. Return value: The newly created Customer resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudchannel/v1:AccountCustomer")]
    public partial class AccountCustomer : Pulumi.CustomResource
    {
        /// <summary>
        /// Secondary contact email. Alternate email and primary contact email are required to have different domains if primary contact email is present. When creating admin.google.com accounts, users get notified credentials at this email. This email address is also used as a recovery email.
        /// </summary>
        [Output("alternateEmail")]
        public Output<string> AlternateEmail { get; private set; } = null!;

        /// <summary>
        /// Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
        /// </summary>
        [Output("channelPartnerId")]
        public Output<string> ChannelPartnerId { get; private set; } = null!;

        /// <summary>
        /// Customer's cloud_identity_id. Populated only if a Cloud Identity resource exists for this customer.
        /// </summary>
        [Output("cloudIdentityId")]
        public Output<string> CloudIdentityId { get; private set; } = null!;

        /// <summary>
        /// Cloud Identity information for the customer. Populated only if a Cloud Identity account exists for this customer.
        /// </summary>
        [Output("cloudIdentityInfo")]
        public Output<Outputs.GoogleCloudChannelV1CloudIdentityInfoResponse> CloudIdentityInfo { get; private set; } = null!;

        /// <summary>
        /// The time at which the customer is created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Required. Primary domain used by the customer. Domain of primary contact email is required to be same as the provided domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
        /// </summary>
        [Output("languageCode")]
        public Output<string> LanguageCode { get; private set; } = null!;

        /// <summary>
        /// Resource name of the customer. Format: accounts/{account_id}/customers/{customer_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required. Name of the organization that the customer entity represents.
        /// </summary>
        [Output("orgDisplayName")]
        public Output<string> OrgDisplayName { get; private set; } = null!;

        /// <summary>
        /// Required. Address of the organization of the customer entity. Region and zip codes are required to enforce US laws and embargoes. Valid address lines are required for all customers. Language code is discarded. Use the Customer-level language code to set the customer's language.
        /// </summary>
        [Output("orgPostalAddress")]
        public Output<Outputs.GoogleTypePostalAddressResponse> OrgPostalAddress { get; private set; } = null!;

        /// <summary>
        /// Primary contact info.
        /// </summary>
        [Output("primaryContactInfo")]
        public Output<Outputs.GoogleCloudChannelV1ContactInfoResponse> PrimaryContactInfo { get; private set; } = null!;

        /// <summary>
        /// The time at which the customer is updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a AccountCustomer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountCustomer(string name, AccountCustomerArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudchannel/v1:AccountCustomer", name, args ?? new AccountCustomerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountCustomer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudchannel/v1:AccountCustomer", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AccountCustomer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountCustomer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccountCustomer(name, id, options);
        }
    }

    public sealed class AccountCustomerArgs : Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// Secondary contact email. Alternate email and primary contact email are required to have different domains if primary contact email is present. When creating admin.google.com accounts, users get notified credentials at this email. This email address is also used as a recovery email.
        /// </summary>
        [Input("alternateEmail")]
        public Input<string>? AlternateEmail { get; set; }

        /// <summary>
        /// Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
        /// </summary>
        [Input("channelPartnerId")]
        public Input<string>? ChannelPartnerId { get; set; }

        /// <summary>
        /// Required. Primary domain used by the customer. Domain of primary contact email is required to be same as the provided domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
        /// </summary>
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// Required. Name of the organization that the customer entity represents.
        /// </summary>
        [Input("orgDisplayName")]
        public Input<string>? OrgDisplayName { get; set; }

        /// <summary>
        /// Required. Address of the organization of the customer entity. Region and zip codes are required to enforce US laws and embargoes. Valid address lines are required for all customers. Language code is discarded. Use the Customer-level language code to set the customer's language.
        /// </summary>
        [Input("orgPostalAddress")]
        public Input<Inputs.GoogleTypePostalAddressArgs>? OrgPostalAddress { get; set; }

        /// <summary>
        /// Primary contact info.
        /// </summary>
        [Input("primaryContactInfo")]
        public Input<Inputs.GoogleCloudChannelV1ContactInfoArgs>? PrimaryContactInfo { get; set; }

        public AccountCustomerArgs()
        {
        }
    }
}
