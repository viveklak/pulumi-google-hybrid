// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1
{
    /// <summary>
    /// Creates a new DatacenterConnector in a given Source.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:vmmigration/v1alpha1:DatacenterConnector")]
    public partial class DatacenterConnector : Pulumi.CustomResource
    {
        /// <summary>
        /// Appliance OVA version. This is the OVA which is manually installed by the user and contains the infrastructure for the automatically updatable components on the appliance.
        /// </summary>
        [Output("applianceInfrastructureVersion")]
        public Output<string> ApplianceInfrastructureVersion { get; private set; } = null!;

        /// <summary>
        /// Appliance last installed update bundle version. This is the version of the automatically updatable components on the appliance.
        /// </summary>
        [Output("applianceSoftwareVersion")]
        public Output<string> ApplianceSoftwareVersion { get; private set; } = null!;

        /// <summary>
        /// The available versions for updating this appliance.
        /// </summary>
        [Output("availableVersions")]
        public Output<Outputs.AvailableUpdatesResponse> AvailableVersions { get; private set; } = null!;

        /// <summary>
        /// The communication channel between the datacenter connector and GCP.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The time the connector was created (as an API call, not when it was actually installed).
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Provides details on the state of the Datacenter Connector in case of an error.
        /// </summary>
        [Output("error")]
        public Output<Outputs.StatusResponse> Error { get; private set; } = null!;

        /// <summary>
        /// The connector's name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
        /// </summary>
        [Output("registrationId")]
        public Output<string> RegistrationId { get; private set; } = null!;

        /// <summary>
        /// The service account to use in the connector when communicating with the cloud.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// State of the DatacenterConnector, as determined by the health checks.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The time the state was last set.
        /// </summary>
        [Output("stateTime")]
        public Output<string> StateTime { get; private set; } = null!;

        /// <summary>
        /// The last time the connector was updated with an API call.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The status of the current / last upgradeAppliance operation.
        /// </summary>
        [Output("upgradeStatus")]
        public Output<Outputs.UpgradeStatusResponse> UpgradeStatus { get; private set; } = null!;

        /// <summary>
        /// The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a DatacenterConnector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatacenterConnector(string name, DatacenterConnectorArgs args, CustomResourceOptions? options = null)
            : base("google-native:vmmigration/v1alpha1:DatacenterConnector", name, args ?? new DatacenterConnectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatacenterConnector(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:vmmigration/v1alpha1:DatacenterConnector", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DatacenterConnector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatacenterConnector Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatacenterConnector(name, id, options);
        }
    }

    public sealed class DatacenterConnectorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The datacenterConnector identifier.
        /// </summary>
        [Input("datacenterConnectorId", required: true)]
        public Input<string> DatacenterConnectorId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
        /// </summary>
        [Input("registrationId")]
        public Input<string>? RegistrationId { get; set; }

        /// <summary>
        /// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The service account to use in the connector when communicating with the cloud.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        /// <summary>
        /// The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public DatacenterConnectorArgs()
        {
        }
    }
}
