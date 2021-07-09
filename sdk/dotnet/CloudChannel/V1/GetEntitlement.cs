// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudChannel.V1
{
    public static class GetEntitlement
    {
        /// <summary>
        /// Returns a requested Entitlement resource. Possible error codes: * PERMISSION_DENIED: The customer doesn't belong to the reseller. * INVALID_ARGUMENT: Required request parameters are missing or invalid. * NOT_FOUND: The customer entitlement was not found. Return value: The requested Entitlement resource.
        /// </summary>
        public static Task<GetEntitlementResult> InvokeAsync(GetEntitlementArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEntitlementResult>("google-native:cloudchannel/v1:getEntitlement", args ?? new GetEntitlementArgs(), options.WithVersion());
    }


    public sealed class GetEntitlementArgs : Pulumi.InvokeArgs
    {
        [Input("accountId", required: true)]
        public string AccountId { get; set; } = null!;

        [Input("customerId", required: true)]
        public string CustomerId { get; set; } = null!;

        [Input("entitlementId", required: true)]
        public string EntitlementId { get; set; } = null!;

        public GetEntitlementArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEntitlementResult
    {
        /// <summary>
        /// Association information to other entitlements.
        /// </summary>
        public readonly Outputs.GoogleCloudChannelV1AssociationInfoResponse AssociationInfo;
        /// <summary>
        /// Commitment settings for a commitment-based Offer. Required for commitment based offers.
        /// </summary>
        public readonly Outputs.GoogleCloudChannelV1CommitmentSettingsResponse CommitmentSettings;
        /// <summary>
        /// The time at which the entitlement is created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Resource name of an entitlement in the form: accounts/{account_id}/customers/{customer_id}/entitlements/{entitlement_id}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The offer resource name for which the entitlement is to be created. Takes the form: accounts/{account_id}/offers/{offer_id}.
        /// </summary>
        public readonly string Offer;
        /// <summary>
        /// Extended entitlement parameters. When creating an entitlement, valid parameters' names and values are defined in the offer's parameter definitions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudChannelV1ParameterResponse> Parameters;
        /// <summary>
        /// Service provisioning details for the entitlement.
        /// </summary>
        public readonly Outputs.GoogleCloudChannelV1ProvisionedServiceResponse ProvisionedService;
        /// <summary>
        /// Current provisioning state of the entitlement.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Optional. This purchase order (PO) information is for resellers to use for their company tracking usage. If a purchaseOrderId value is given, it appears in the API responses and shows up in the invoice. The property accepts up to 80 plain text characters.
        /// </summary>
        public readonly string PurchaseOrderId;
        /// <summary>
        /// Enumerable of all current suspension reasons for an entitlement.
        /// </summary>
        public readonly ImmutableArray<string> SuspensionReasons;
        /// <summary>
        /// Settings for trial offers.
        /// </summary>
        public readonly Outputs.GoogleCloudChannelV1TrialSettingsResponse TrialSettings;
        /// <summary>
        /// The time at which the entitlement is updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetEntitlementResult(
            Outputs.GoogleCloudChannelV1AssociationInfoResponse associationInfo,

            Outputs.GoogleCloudChannelV1CommitmentSettingsResponse commitmentSettings,

            string createTime,

            string name,

            string offer,

            ImmutableArray<Outputs.GoogleCloudChannelV1ParameterResponse> parameters,

            Outputs.GoogleCloudChannelV1ProvisionedServiceResponse provisionedService,

            string provisioningState,

            string purchaseOrderId,

            ImmutableArray<string> suspensionReasons,

            Outputs.GoogleCloudChannelV1TrialSettingsResponse trialSettings,

            string updateTime)
        {
            AssociationInfo = associationInfo;
            CommitmentSettings = commitmentSettings;
            CreateTime = createTime;
            Name = name;
            Offer = offer;
            Parameters = parameters;
            ProvisionedService = provisionedService;
            ProvisioningState = provisioningState;
            PurchaseOrderId = purchaseOrderId;
            SuspensionReasons = suspensionReasons;
            TrialSettings = trialSettings;
            UpdateTime = updateTime;
        }
    }
}
