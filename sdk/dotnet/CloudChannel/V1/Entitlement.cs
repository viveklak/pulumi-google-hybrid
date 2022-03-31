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
    /// Creates an entitlement for a customer. Possible error codes: * PERMISSION_DENIED: The customer doesn't belong to the reseller. * INVALID_ARGUMENT: * Required request parameters are missing or invalid. * There is already a customer entitlement for a SKU from the same product family. * INVALID_VALUE: Make sure the OfferId is valid. If it is, contact Google Channel support for further troubleshooting. * NOT_FOUND: The customer or offer resource was not found. * ALREADY_EXISTS: * The SKU was already purchased for the customer. * The customer's primary email already exists. Retry after changing the customer's primary contact email. * CONDITION_NOT_MET or FAILED_PRECONDITION: * The domain required for purchasing a SKU has not been verified. * A pre-requisite SKU required to purchase an Add-On SKU is missing. For example, Google Workspace Business Starter is required to purchase Vault or Drive. * (Developer accounts only) Reseller and resold domain must meet the following naming requirements: * Domain names must start with goog-test. * Domain names must include the reseller domain. * INTERNAL: Any non-user error related to a technical issue in the backend. Contact Cloud Channel support. * UNKNOWN: Any non-user error related to a technical issue in the backend. Contact Cloud Channel support. Return value: The ID of a long-running operation. To get the results of the operation, call the GetOperation method of CloudChannelOperationsService. The Operation metadata will contain an instance of OperationMetadata.
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudchannel/v1:Entitlement")]
    public partial class Entitlement : Pulumi.CustomResource
    {
        /// <summary>
        /// Association information to other entitlements.
        /// </summary>
        [Output("associationInfo")]
        public Output<Outputs.GoogleCloudChannelV1AssociationInfoResponse> AssociationInfo { get; private set; } = null!;

        /// <summary>
        /// Commitment settings for a commitment-based Offer. Required for commitment based offers.
        /// </summary>
        [Output("commitmentSettings")]
        public Output<Outputs.GoogleCloudChannelV1CommitmentSettingsResponse> CommitmentSettings { get; private set; } = null!;

        /// <summary>
        /// The time at which the entitlement is created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Resource name of an entitlement in the form: accounts/{account_id}/customers/{customer_id}/entitlements/{entitlement_id}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The offer resource name for which the entitlement is to be created. Takes the form: accounts/{account_id}/offers/{offer_id}.
        /// </summary>
        [Output("offer")]
        public Output<string> Offer { get; private set; } = null!;

        /// <summary>
        /// Extended entitlement parameters. When creating an entitlement, valid parameter names and values are defined in the Offer.parameter_definitions. The response may include the following output-only Parameters: - assigned_units: The number of licenses assigned to users. - max_units: The maximum assignable units for a flexible offer. - num_units: The total commitment for commitment-based offers.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.GoogleCloudChannelV1ParameterResponse>> Parameters { get; private set; } = null!;

        /// <summary>
        /// Service provisioning details for the entitlement.
        /// </summary>
        [Output("provisionedService")]
        public Output<Outputs.GoogleCloudChannelV1ProvisionedServiceResponse> ProvisionedService { get; private set; } = null!;

        /// <summary>
        /// Current provisioning state of the entitlement.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Optional. This purchase order (PO) information is for resellers to use for their company tracking usage. If a purchaseOrderId value is given, it appears in the API responses and shows up in the invoice. The property accepts up to 80 plain text characters. This is only supported for Google Workspace entitlements.
        /// </summary>
        [Output("purchaseOrderId")]
        public Output<string> PurchaseOrderId { get; private set; } = null!;

        /// <summary>
        /// Enumerable of all current suspension reasons for an entitlement.
        /// </summary>
        [Output("suspensionReasons")]
        public Output<ImmutableArray<string>> SuspensionReasons { get; private set; } = null!;

        /// <summary>
        /// Settings for trial offers.
        /// </summary>
        [Output("trialSettings")]
        public Output<Outputs.GoogleCloudChannelV1TrialSettingsResponse> TrialSettings { get; private set; } = null!;

        /// <summary>
        /// The time at which the entitlement is updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Entitlement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Entitlement(string name, EntitlementArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudchannel/v1:Entitlement", name, args ?? new EntitlementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Entitlement(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudchannel/v1:Entitlement", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Entitlement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Entitlement Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Entitlement(name, id, options);
        }
    }

    public sealed class EntitlementArgs : Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// Association information to other entitlements.
        /// </summary>
        [Input("associationInfo")]
        public Input<Inputs.GoogleCloudChannelV1AssociationInfoArgs>? AssociationInfo { get; set; }

        /// <summary>
        /// Commitment settings for a commitment-based Offer. Required for commitment based offers.
        /// </summary>
        [Input("commitmentSettings")]
        public Input<Inputs.GoogleCloudChannelV1CommitmentSettingsArgs>? CommitmentSettings { get; set; }

        [Input("customerId", required: true)]
        public Input<string> CustomerId { get; set; } = null!;

        /// <summary>
        /// The offer resource name for which the entitlement is to be created. Takes the form: accounts/{account_id}/offers/{offer_id}.
        /// </summary>
        [Input("offer", required: true)]
        public Input<string> Offer { get; set; } = null!;

        [Input("parameters")]
        private InputList<Inputs.GoogleCloudChannelV1ParameterArgs>? _parameters;

        /// <summary>
        /// Extended entitlement parameters. When creating an entitlement, valid parameter names and values are defined in the Offer.parameter_definitions. The response may include the following output-only Parameters: - assigned_units: The number of licenses assigned to users. - max_units: The maximum assignable units for a flexible offer. - num_units: The total commitment for commitment-based offers.
        /// </summary>
        public InputList<Inputs.GoogleCloudChannelV1ParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.GoogleCloudChannelV1ParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Optional. This purchase order (PO) information is for resellers to use for their company tracking usage. If a purchaseOrderId value is given, it appears in the API responses and shows up in the invoice. The property accepts up to 80 plain text characters. This is only supported for Google Workspace entitlements.
        /// </summary>
        [Input("purchaseOrderId")]
        public Input<string>? PurchaseOrderId { get; set; }

        /// <summary>
        /// Optional. You can specify an optional unique request ID, and if you need to retry your request, the server will know to ignore the request if it's complete. For example, you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if it received the original operation with the same request ID. If it did, it will ignore the second request. The request ID must be a valid [UUID](https://tools.ietf.org/html/rfc4122) with the exception that zero UUID is not supported (`00000000-0000-0000-0000-000000000000`).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        public EntitlementArgs()
        {
        }
    }
}
