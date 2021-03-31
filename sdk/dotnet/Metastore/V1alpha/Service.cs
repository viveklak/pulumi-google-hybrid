// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Metastore.V1alpha
{
    /// <summary>
    /// Creates a metastore service in a project and location.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:metastore/v1alpha:Service")]
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:metastore/v1alpha:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:metastore/v1alpha:Service", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Service(name, id, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output only. A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
        /// </summary>
        [Input("artifactGcsUri")]
        public Input<string>? ArtifactGcsUri { get; set; }

        /// <summary>
        /// Output only. The time when the metastore service was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Output only. The URI of the endpoint used to access the metastore service.
        /// </summary>
        [Input("endpointUri")]
        public Input<string>? EndpointUri { get; set; }

        /// <summary>
        /// Configuration information specific to running Hive metastore software as the metastore service.
        /// </summary>
        [Input("hiveMetastoreConfig")]
        public Input<Inputs.HiveMetastoreConfigArgs>? HiveMetastoreConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels for the metastore service.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        /// <summary>
        /// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The setting that defines how metastore metadata should be integrated with external services and systems.
        /// </summary>
        [Input("metadataIntegration")]
        public Input<Inputs.MetadataIntegrationArgs>? MetadataIntegration { get; set; }

        /// <summary>
        /// Output only. The metadata management activities of the metastore service.
        /// </summary>
        [Input("metadataManagementActivity")]
        public Input<Inputs.MetadataManagementActivityArgs>? MetadataManagementActivity { get; set; }

        /// <summary>
        /// Immutable. The relative resource name of the metastore service, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The TCP port at which the metastore service is reached. Default: 9083.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
        /// </summary>
        [Input("releaseChannel")]
        public Input<string>? ReleaseChannel { get; set; }

        [Input("servicesId", required: true)]
        public Input<string> ServicesId { get; set; } = null!;

        /// <summary>
        /// Output only. The current state of the metastore service.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Output only. Additional information about the current state of the metastore service, if available.
        /// </summary>
        [Input("stateMessage")]
        public Input<string>? StateMessage { get; set; }

        /// <summary>
        /// The tier of the service.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        /// <summary>
        /// Output only. The globally unique resource identifier of the metastore service.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Output only. The time when the metastore service was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public ServiceArgs()
        {
        }
    }
}
