// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AssuredWorkloads.V1
{
    /// <summary>
    /// Creates Assured Workload.
    /// </summary>
    [GoogleNativeResourceType("google-native:assuredworkloads/v1:Workload")]
    public partial class Workload : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF`.
        /// </summary>
        [Output("billingAccount")]
        public Output<string> BillingAccount { get; private set; } = null!;

        /// <summary>
        /// Immutable. Compliance Regime associated with this workload.
        /// </summary>
        [Output("complianceRegime")]
        public Output<string> ComplianceRegime { get; private set; } = null!;

        /// <summary>
        /// Immutable. The Workload creation timestamp.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.
        /// </summary>
        [Output("enableSovereignControls")]
        public Output<bool> EnableSovereignControls { get; private set; } = null!;

        /// <summary>
        /// Optional. ETag of the workload, it is calculated on the basis of the Workload contents. It will be used in Update &amp; Delete operations.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Represents the KAJ enrollment state of the given workload.
        /// </summary>
        [Output("kajEnrollmentState")]
        public Output<string> KajEnrollmentState { get; private set; } = null!;

        /// <summary>
        /// Input only. Settings used to create a CMEK crypto key. When set, a project with a KMS CMEK key is provisioned. This field is deprecated as of Feb 28, 2022. In order to create a Keyring, callers should specify, ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type field.
        /// </summary>
        [Output("kmsSettings")]
        public Output<Outputs.GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsResponse> KmsSettings { get; private set; } = null!;

        /// <summary>
        /// Optional. Labels applied to the workload.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Optional. The resource name of the workload. Format: organizations/{organization}/locations/{location}/workloads/{workload} Read-only.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Input only. The parent resource for the resources managed by this Assured Workload. May be either empty or a folder resource which is a child of the Workload parent. If not specified all resources are created under the parent organization. Format: folders/{folder_id}
        /// </summary>
        [Output("provisionedResourcesParent")]
        public Output<string> ProvisionedResourcesParent { get; private set; } = null!;

        /// <summary>
        /// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
        /// </summary>
        [Output("resourceSettings")]
        public Output<ImmutableArray<Outputs.GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsResponse>> ResourceSettings { get; private set; } = null!;

        /// <summary>
        /// The resources associated with this workload. These resources will be created when creating the workload. If any of the projects already exist, the workload creation will fail. Always read only.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.GoogleCloudAssuredworkloadsV1WorkloadResourceInfoResponse>> Resources { get; private set; } = null!;

        /// <summary>
        /// Represents the SAA enrollment response of the given workload. SAA enrollment response is queried during GetWorkload call. In failure cases, user friendly error message is shown in SAA details page.
        /// </summary>
        [Output("saaEnrollmentResponse")]
        public Output<Outputs.GoogleCloudAssuredworkloadsV1WorkloadSaaEnrollmentResponseResponse> SaaEnrollmentResponse { get; private set; } = null!;


        /// <summary>
        /// Create a Workload resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workload(string name, WorkloadArgs args, CustomResourceOptions? options = null)
            : base("google-native:assuredworkloads/v1:Workload", name, args ?? new WorkloadArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workload(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:assuredworkloads/v1:Workload", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Workload resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workload Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Workload(name, id, options);
        }
    }

    public sealed class WorkloadArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF`.
        /// </summary>
        [Input("billingAccount")]
        public Input<string>? BillingAccount { get; set; }

        /// <summary>
        /// Immutable. Compliance Regime associated with this workload.
        /// </summary>
        [Input("complianceRegime", required: true)]
        public Input<Pulumi.GoogleNative.AssuredWorkloads.V1.WorkloadComplianceRegime> ComplianceRegime { get; set; } = null!;

        /// <summary>
        /// The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.
        /// </summary>
        [Input("enableSovereignControls")]
        public Input<bool>? EnableSovereignControls { get; set; }

        /// <summary>
        /// Optional. ETag of the workload, it is calculated on the basis of the Workload contents. It will be used in Update &amp; Delete operations.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Optional. A identifier associated with the workload and underlying projects which allows for the break down of billing costs for a workload. The value provided for the identifier will add a label to the workload and contained projects with the identifier as the value.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Input only. Settings used to create a CMEK crypto key. When set, a project with a KMS CMEK key is provisioned. This field is deprecated as of Feb 28, 2022. In order to create a Keyring, callers should specify, ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type field.
        /// </summary>
        [Input("kmsSettings")]
        public Input<Inputs.GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsArgs>? KmsSettings { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Labels applied to the workload.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optional. The resource name of the workload. Format: organizations/{organization}/locations/{location}/workloads/{workload} Read-only.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// Input only. The parent resource for the resources managed by this Assured Workload. May be either empty or a folder resource which is a child of the Workload parent. If not specified all resources are created under the parent organization. Format: folders/{folder_id}
        /// </summary>
        [Input("provisionedResourcesParent")]
        public Input<string>? ProvisionedResourcesParent { get; set; }

        [Input("resourceSettings")]
        private InputList<Inputs.GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsArgs>? _resourceSettings;

        /// <summary>
        /// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
        /// </summary>
        public InputList<Inputs.GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsArgs> ResourceSettings
        {
            get => _resourceSettings ?? (_resourceSettings = new InputList<Inputs.GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsArgs>());
            set => _resourceSettings = value;
        }

        public WorkloadArgs()
        {
        }
    }
}
