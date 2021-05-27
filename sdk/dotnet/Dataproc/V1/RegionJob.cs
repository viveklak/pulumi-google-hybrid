// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1
{
    /// <summary>
    /// Submits a job to a cluster.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataproc/v1:RegionJob")]
    public partial class RegionJob : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether the job is completed. If the value is false, the job is still in progress. If true, the job is completed, and status.state field will indicate if it was successful, failed, or cancelled.
        /// </summary>
        [Output("done")]
        public Output<bool> Done { get; private set; } = null!;

        /// <summary>
        /// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
        /// </summary>
        [Output("driverControlFilesUri")]
        public Output<string> DriverControlFilesUri { get; private set; } = null!;

        /// <summary>
        /// A URI pointing to the location of the stdout of the job's driver program.
        /// </summary>
        [Output("driverOutputResourceUri")]
        public Output<string> DriverOutputResourceUri { get; private set; } = null!;

        /// <summary>
        /// Optional. Job is a Hadoop job.
        /// </summary>
        [Output("hadoopJob")]
        public Output<Outputs.HadoopJobResponse> HadoopJob { get; private set; } = null!;

        /// <summary>
        /// Optional. Job is a Hive job.
        /// </summary>
        [Output("hiveJob")]
        public Output<Outputs.HiveJobResponse> HiveJob { get; private set; } = null!;

        /// <summary>
        /// A UUID that uniquely identifies a job within the project over time. This is in contrast to a user-settable reference.job_id that may be reused over time.
        /// </summary>
        [Output("jobUuid")]
        public Output<string> JobUuid { get; private set; } = null!;

        /// <summary>
        /// Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Optional. Job is a Pig job.
        /// </summary>
        [Output("pigJob")]
        public Output<Outputs.PigJobResponse> PigJob { get; private set; } = null!;

        /// <summary>
        /// Required. Job information, including how, when, and where to run the job.
        /// </summary>
        [Output("placement")]
        public Output<Outputs.JobPlacementResponse> Placement { get; private set; } = null!;

        /// <summary>
        /// Optional. Job is a Presto job.
        /// </summary>
        [Output("prestoJob")]
        public Output<Outputs.PrestoJobResponse> PrestoJob { get; private set; } = null!;

        /// <summary>
        /// Optional. Job is a PySpark job.
        /// </summary>
        [Output("pysparkJob")]
        public Output<Outputs.PySparkJobResponse> PysparkJob { get; private set; } = null!;

        /// <summary>
        /// Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
        /// </summary>
        [Output("reference")]
        public Output<Outputs.JobReferenceResponse> Reference { get; private set; } = null!;

        /// <summary>
        /// Optional. Job scheduling configuration.
        /// </summary>
        [Output("scheduling")]
        public Output<Outputs.JobSchedulingResponse> Scheduling { get; private set; } = null!;

        /// <summary>
        /// Optional. Job is a Spark job.
        /// </summary>
        [Output("sparkJob")]
        public Output<Outputs.SparkJobResponse> SparkJob { get; private set; } = null!;

        /// <summary>
        /// Optional. Job is a SparkR job.
        /// </summary>
        [Output("sparkRJob")]
        public Output<Outputs.SparkRJobResponse> SparkRJob { get; private set; } = null!;

        /// <summary>
        /// Optional. Job is a SparkSql job.
        /// </summary>
        [Output("sparkSqlJob")]
        public Output<Outputs.SparkSqlJobResponse> SparkSqlJob { get; private set; } = null!;

        /// <summary>
        /// The job status. Additional application-specific status information may be contained in the type_job and yarn_applications fields.
        /// </summary>
        [Output("status")]
        public Output<Outputs.JobStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// The previous job status.
        /// </summary>
        [Output("statusHistory")]
        public Output<ImmutableArray<Outputs.JobStatusResponse>> StatusHistory { get; private set; } = null!;

        /// <summary>
        /// The collection of YARN applications spun up by this job.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
        /// </summary>
        [Output("yarnApplications")]
        public Output<ImmutableArray<Outputs.YarnApplicationResponse>> YarnApplications { get; private set; } = null!;


        /// <summary>
        /// Create a RegionJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionJob(string name, RegionJobArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataproc/v1:RegionJob", name, args ?? new RegionJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataproc/v1:RegionJob", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RegionJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegionJob(name, id, options);
        }
    }

    public sealed class RegionJobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Job is a Hadoop job.
        /// </summary>
        [Input("hadoopJob")]
        public Input<Inputs.HadoopJobArgs>? HadoopJob { get; set; }

        /// <summary>
        /// Optional. Job is a Hive job.
        /// </summary>
        [Input("hiveJob")]
        public Input<Inputs.HiveJobArgs>? HiveJob { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Optional. Job is a Pig job.
        /// </summary>
        [Input("pigJob")]
        public Input<Inputs.PigJobArgs>? PigJob { get; set; }

        /// <summary>
        /// Required. Job information, including how, when, and where to run the job.
        /// </summary>
        [Input("placement")]
        public Input<Inputs.JobPlacementArgs>? Placement { get; set; }

        /// <summary>
        /// Optional. Job is a Presto job.
        /// </summary>
        [Input("prestoJob")]
        public Input<Inputs.PrestoJobArgs>? PrestoJob { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Optional. Job is a PySpark job.
        /// </summary>
        [Input("pysparkJob")]
        public Input<Inputs.PySparkJobArgs>? PysparkJob { get; set; }

        /// <summary>
        /// Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
        /// </summary>
        [Input("reference")]
        public Input<Inputs.JobReferenceArgs>? Reference { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// Optional. A unique id used to identify the request. If the server receives two SubmitJobRequest (https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1#google.cloud.dataproc.v1.SubmitJobRequest)s with the same id, then the second request will be ignored and the first Job created and stored in the backend is returned.It is recommended to always set this value to a UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier).The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Optional. Job scheduling configuration.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.JobSchedulingArgs>? Scheduling { get; set; }

        /// <summary>
        /// Optional. Job is a Spark job.
        /// </summary>
        [Input("sparkJob")]
        public Input<Inputs.SparkJobArgs>? SparkJob { get; set; }

        /// <summary>
        /// Optional. Job is a SparkR job.
        /// </summary>
        [Input("sparkRJob")]
        public Input<Inputs.SparkRJobArgs>? SparkRJob { get; set; }

        /// <summary>
        /// Optional. Job is a SparkSql job.
        /// </summary>
        [Input("sparkSqlJob")]
        public Input<Inputs.SparkSqlJobArgs>? SparkSqlJob { get; set; }

        public RegionJobArgs()
        {
        }
    }
}
