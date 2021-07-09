// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Creates an Apigee organization. See [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:Organization")]
    public partial class Organization : Pulumi.CustomResource
    {
        /// <summary>
        /// Addon configurations of the Apigee organization.
        /// </summary>
        [Output("addonsConfig")]
        public Output<Outputs.GoogleCloudApigeeV1AddonsConfigResponse> AddonsConfig { get; private set; } = null!;

        /// <summary>
        /// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
        /// </summary>
        [Output("analyticsRegion")]
        public Output<string> AnalyticsRegion { get; private set; } = null!;

        /// <summary>
        /// Not used by Apigee.
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableArray<string>> Attributes { get; private set; } = null!;

        /// <summary>
        /// Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started). Valid only when [RuntimeType](#RuntimeType) is set to `CLOUD`. The value must be set before the creation of a runtime instance and can be updated only when there are no runtime instances. For example: `default`. Apigee also supports shared VPC (that is, the host network project is not the same as the one that is peering with Apigee). See [Shared VPC overview](https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.
        /// </summary>
        [Output("authorizedNetwork")]
        public Output<string> AuthorizedNetwork { get; private set; } = null!;

        /// <summary>
        /// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
        /// </summary>
        [Output("billingType")]
        public Output<string> BillingType { get; private set; } = null!;

        /// <summary>
        /// Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when [RuntimeType](#RuntimeType) is `CLOUD`.
        /// </summary>
        [Output("caCertificate")]
        public Output<string> CaCertificate { get; private set; } = null!;

        /// <summary>
        /// Time that the Apigee organization was created in milliseconds since epoch.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Not used by Apigee.
        /// </summary>
        [Output("customerName")]
        public Output<string> CustomerName { get; private set; } = null!;

        /// <summary>
        /// Description of the Apigee organization.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// List of environments in the Apigee organization.
        /// </summary>
        [Output("environments")]
        public Output<ImmutableArray<string>> Environments { get; private set; } = null!;

        /// <summary>
        /// Time that the Apigee organization is scheduled for deletion.
        /// </summary>
        [Output("expiresAt")]
        public Output<string> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// Time that the Apigee organization was last modified in milliseconds since epoch.
        /// </summary>
        [Output("lastModifiedAt")]
        public Output<string> LastModifiedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the Apigee organization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Project ID associated with the Apigee organization.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Properties defined in the Apigee organization profile.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.GoogleCloudApigeeV1PropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. Required when [RuntimeType](#RuntimeType) is `CLOUD`. If not specified when [RuntimeType](#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.
        /// </summary>
        [Output("runtimeDatabaseEncryptionKeyName")]
        public Output<string> RuntimeDatabaseEncryptionKeyName { get; private set; } = null!;

        /// <summary>
        /// Runtime type of the Apigee organization based on the Apigee subscription purchased.
        /// </summary>
        [Output("runtimeType")]
        public Output<string> RuntimeType { get; private set; } = null!;

        /// <summary>
        /// State of the organization. Values other than ACTIVE means the resource is not ready to use.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Not used by Apigee.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Organization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Organization(string name, OrganizationArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:Organization", name, args ?? new OrganizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Organization(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:Organization", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Organization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Organization Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Organization(name, id, options);
        }
    }

    public sealed class OrganizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Addon configurations of the Apigee organization.
        /// </summary>
        [Input("addonsConfig")]
        public Input<Inputs.GoogleCloudApigeeV1AddonsConfigArgs>? AddonsConfig { get; set; }

        /// <summary>
        /// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
        /// </summary>
        [Input("analyticsRegion", required: true)]
        public Input<string> AnalyticsRegion { get; set; } = null!;

        [Input("attributes")]
        private InputList<string>? _attributes;

        /// <summary>
        /// Not used by Apigee.
        /// </summary>
        public InputList<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<string>());
            set => _attributes = value;
        }

        /// <summary>
        /// Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started). Valid only when [RuntimeType](#RuntimeType) is set to `CLOUD`. The value must be set before the creation of a runtime instance and can be updated only when there are no runtime instances. For example: `default`. Apigee also supports shared VPC (that is, the host network project is not the same as the one that is peering with Apigee). See [Shared VPC overview](https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.
        /// </summary>
        [Input("authorizedNetwork")]
        public Input<string>? AuthorizedNetwork { get; set; }

        /// <summary>
        /// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
        /// </summary>
        [Input("billingType")]
        public Input<Pulumi.GoogleNative.Apigee.V1.OrganizationBillingType>? BillingType { get; set; }

        /// <summary>
        /// Not used by Apigee.
        /// </summary>
        [Input("customerName")]
        public Input<string>? CustomerName { get; set; }

        /// <summary>
        /// Description of the Apigee organization.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        /// <summary>
        /// Properties defined in the Apigee organization profile.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.GoogleCloudApigeeV1PropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. Required when [RuntimeType](#RuntimeType) is `CLOUD`. If not specified when [RuntimeType](#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.
        /// </summary>
        [Input("runtimeDatabaseEncryptionKeyName")]
        public Input<string>? RuntimeDatabaseEncryptionKeyName { get; set; }

        /// <summary>
        /// Runtime type of the Apigee organization based on the Apigee subscription purchased.
        /// </summary>
        [Input("runtimeType", required: true)]
        public Input<Pulumi.GoogleNative.Apigee.V1.OrganizationRuntimeType> RuntimeType { get; set; } = null!;

        /// <summary>
        /// Not used by Apigee.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Apigee.V1.OrganizationType>? Type { get; set; }

        public OrganizationArgs()
        {
        }
    }
}
