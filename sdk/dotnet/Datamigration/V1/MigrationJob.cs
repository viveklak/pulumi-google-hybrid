// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1
{
    /// <summary>
    /// Creates a new migration job in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:datamigration/v1:MigrationJob")]
    public partial class MigrationJob : Pulumi.CustomResource
    {
        /// <summary>
        /// The timestamp when the migration job resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The resource name (URI) of the destination connection profile.
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// The database engine type and provider of the destination.
        /// </summary>
        [Output("destinationDatabase")]
        public Output<Outputs.DatabaseTypeResponse> DestinationDatabase { get; private set; } = null!;

        /// <summary>
        /// The migration job display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The initial dump flags. This field and the "dump_path" field are mutually exclusive.
        /// </summary>
        [Output("dumpFlags")]
        public Output<Outputs.DumpFlagsResponse> DumpFlags { get; private set; } = null!;

        /// <summary>
        /// The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]). This field and the "dump_flags" field are mutually exclusive.
        /// </summary>
        [Output("dumpPath")]
        public Output<string> DumpPath { get; private set; } = null!;

        /// <summary>
        /// The duration of the migration job (in seconds). A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        [Output("duration")]
        public Output<string> Duration { get; private set; } = null!;

        /// <summary>
        /// If the migration job is completed, the time when it was completed.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// The error details in case of state FAILED.
        /// </summary>
        [Output("error")]
        public Output<Outputs.StatusResponse> Error { get; private set; } = null!;

        /// <summary>
        /// The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current migration job phase.
        /// </summary>
        [Output("phase")]
        public Output<string> Phase { get; private set; } = null!;

        /// <summary>
        /// The details needed to communicate to the source over Reverse SSH tunnel connectivity.
        /// </summary>
        [Output("reverseSshConnectivity")]
        public Output<Outputs.ReverseSshConnectivityResponse> ReverseSshConnectivity { get; private set; } = null!;

        /// <summary>
        /// The resource name (URI) of the source connection profile.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// The database engine type and provider of the source.
        /// </summary>
        [Output("sourceDatabase")]
        public Output<Outputs.DatabaseTypeResponse> SourceDatabase { get; private set; } = null!;

        /// <summary>
        /// The current migration job state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// static ip connectivity data (default, no additional details needed).
        /// </summary>
        [Output("staticIpConnectivity")]
        public Output<Outputs.StaticIpConnectivityResponse> StaticIpConnectivity { get; private set; } = null!;

        /// <summary>
        /// The migration job type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the migration job resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The details of the VPC network that the source database is located in.
        /// </summary>
        [Output("vpcPeeringConnectivity")]
        public Output<Outputs.VpcPeeringConnectivityResponse> VpcPeeringConnectivity { get; private set; } = null!;


        /// <summary>
        /// Create a MigrationJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MigrationJob(string name, MigrationJobArgs args, CustomResourceOptions? options = null)
            : base("google-native:datamigration/v1:MigrationJob", name, args ?? new MigrationJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MigrationJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:datamigration/v1:MigrationJob", name, null, MakeResourceOptions(options, id))
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
        /// The resource name (URI) of the destination connection profile.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

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
        /// The initial dump flags. This field and the "dump_path" field are mutually exclusive.
        /// </summary>
        [Input("dumpFlags")]
        public Input<Inputs.DumpFlagsArgs>? DumpFlags { get; set; }

        /// <summary>
        /// The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]). This field and the "dump_flags" field are mutually exclusive.
        /// </summary>
        [Input("dumpPath")]
        public Input<string>? DumpPath { get; set; }

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

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Required. The ID of the instance to create.
        /// </summary>
        [Input("migrationJobId", required: true)]
        public Input<string> MigrationJobId { get; set; } = null!;

        /// <summary>
        /// The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A unique id used to identify the request. If the server receives two requests with the same id, then the second request will be ignored. It is recommended to always set this value to a UUID. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The details needed to communicate to the source over Reverse SSH tunnel connectivity.
        /// </summary>
        [Input("reverseSshConnectivity")]
        public Input<Inputs.ReverseSshConnectivityArgs>? ReverseSshConnectivity { get; set; }

        /// <summary>
        /// The resource name (URI) of the source connection profile.
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        /// <summary>
        /// The database engine type and provider of the source.
        /// </summary>
        [Input("sourceDatabase")]
        public Input<Inputs.DatabaseTypeArgs>? SourceDatabase { get; set; }

        /// <summary>
        /// The current migration job state.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.Datamigration.V1.MigrationJobState>? State { get; set; }

        /// <summary>
        /// static ip connectivity data (default, no additional details needed).
        /// </summary>
        [Input("staticIpConnectivity")]
        public Input<Inputs.StaticIpConnectivityArgs>? StaticIpConnectivity { get; set; }

        /// <summary>
        /// The migration job type.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.Datamigration.V1.MigrationJobType> Type { get; set; } = null!;

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
