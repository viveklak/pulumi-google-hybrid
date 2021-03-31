// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Datamigration.V1beta1
{
    /// <summary>
    /// Creates a new migration job in a given project and location.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:datamigration/v1beta1:MigrationJob")]
    public partial class MigrationJob : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a MigrationJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MigrationJob(string name, MigrationJobArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:datamigration/v1beta1:MigrationJob", name, args ?? new MigrationJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MigrationJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:datamigration/v1beta1:MigrationJob", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MigrationJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MigrationJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MigrationJob(name, id, options);
        }
    }

    public sealed class MigrationJobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output only. The timestamp when the migration job resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Required. The resource name (URI) of the destination connection profile.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// The database engine type and provider of the destination.
        /// </summary>
        [Input("destinationDatabase")]
        public Input<Inputs.DatabaseTypeArgs>? DestinationDatabase { get; set; }

        /// <summary>
        /// The migration job display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
        /// </summary>
        [Input("dumpPath")]
        public Input<string>? DumpPath { get; set; }

        /// <summary>
        /// Output only. The duration of the migration job (in seconds). A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// Output only. If the migration job is completed, the time when it was completed.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Output only. The error details in case of state FAILED.
        /// </summary>
        [Input("error")]
        public Input<Inputs.StatusArgs>? Error { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        [Input("migrationJobsId", required: true)]
        public Input<string> MigrationJobsId { get; set; } = null!;

        /// <summary>
        /// The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/instances/{instance}.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Output only. The current migration job phase.
        /// </summary>
        [Input("phase")]
        public Input<string>? Phase { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// The details needed to communicate to the source over Reverse SSH tunnel connectivity.
        /// </summary>
        [Input("reverseSshConnectivity")]
        public Input<Inputs.ReverseSshConnectivityArgs>? ReverseSshConnectivity { get; set; }

        /// <summary>
        /// Required. The resource name (URI) of the source connection profile.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// The database engine type and provider of the source.
        /// </summary>
        [Input("sourceDatabase")]
        public Input<Inputs.DatabaseTypeArgs>? SourceDatabase { get; set; }

        /// <summary>
        /// The current migration job state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// static ip connectivity data (default, no additional details needed).
        /// </summary>
        [Input("staticIpConnectivity")]
        public Input<Inputs.StaticIpConnectivityArgs>? StaticIpConnectivity { get; set; }

        /// <summary>
        /// Required. The migration job type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Output only. The timestamp when the migration job resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// The details of the VPC network that the source database is located in.
        /// </summary>
        [Input("vpcPeeringConnectivity")]
        public Input<Inputs.VpcPeeringConnectivityArgs>? VpcPeeringConnectivity { get; set; }

        public MigrationJobArgs()
        {
        }
    }
}
