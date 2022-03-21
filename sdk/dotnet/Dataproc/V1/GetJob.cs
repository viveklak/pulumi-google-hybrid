// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1
{
    public static class GetJob
    {
        /// <summary>
        /// Gets the resource representation for a job in a project.
        /// </summary>
        public static Task<GetJobResult> InvokeAsync(GetJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJobResult>("google-native:dataproc/v1:getJob", args ?? new GetJobArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the resource representation for a job in a project.
        /// </summary>
        public static Output<GetJobResult> Invoke(GetJobInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetJobResult>("google-native:dataproc/v1:getJob", args ?? new GetJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJobArgs : Pulumi.InvokeArgs
    {
        [Input("jobId", required: true)]
        public string JobId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetJobArgs()
        {
        }
    }

    public sealed class GetJobInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public GetJobInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetJobResult
    {
        /// <summary>
        /// Indicates whether the job is completed. If the value is false, the job is still in progress. If true, the job is completed, and status.state field will indicate if it was successful, failed, or cancelled.
        /// </summary>
        public readonly bool Done;
        /// <summary>
        /// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
        /// </summary>
        public readonly string DriverControlFilesUri;
        /// <summary>
        /// A URI pointing to the location of the stdout of the job's driver program.
        /// </summary>
        public readonly string DriverOutputResourceUri;
        /// <summary>
        /// Optional. Job is a Hadoop job.
        /// </summary>
        public readonly Outputs.HadoopJobResponse HadoopJob;
        /// <summary>
        /// Optional. Job is a Hive job.
        /// </summary>
        public readonly Outputs.HiveJobResponse HiveJob;
        /// <summary>
        /// A UUID that uniquely identifies a job within the project over time. This is in contrast to a user-settable reference.job_id that may be reused over time.
        /// </summary>
        public readonly string JobUuid;
        /// <summary>
        /// Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. Job is a Pig job.
        /// </summary>
        public readonly Outputs.PigJobResponse PigJob;
        /// <summary>
        /// Job information, including how, when, and where to run the job.
        /// </summary>
        public readonly Outputs.JobPlacementResponse Placement;
        /// <summary>
        /// Optional. Job is a Presto job.
        /// </summary>
        public readonly Outputs.PrestoJobResponse PrestoJob;
        /// <summary>
        /// Optional. Job is a PySpark job.
        /// </summary>
        public readonly Outputs.PySparkJobResponse PysparkJob;
        /// <summary>
        /// Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
        /// </summary>
        public readonly Outputs.JobReferenceResponse Reference;
        /// <summary>
        /// Optional. Job scheduling configuration.
        /// </summary>
        public readonly Outputs.JobSchedulingResponse Scheduling;
        /// <summary>
        /// Optional. Job is a Spark job.
        /// </summary>
        public readonly Outputs.SparkJobResponse SparkJob;
        /// <summary>
        /// Optional. Job is a SparkR job.
        /// </summary>
        public readonly Outputs.SparkRJobResponse SparkRJob;
        /// <summary>
        /// Optional. Job is a SparkSql job.
        /// </summary>
        public readonly Outputs.SparkSqlJobResponse SparkSqlJob;
        /// <summary>
        /// The job status. Additional application-specific status information may be contained in the type_job and yarn_applications fields.
        /// </summary>
        public readonly Outputs.JobStatusResponse Status;
        /// <summary>
        /// The previous job status.
        /// </summary>
        public readonly ImmutableArray<Outputs.JobStatusResponse> StatusHistory;
        /// <summary>
        /// The collection of YARN applications spun up by this job.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
        /// </summary>
        public readonly ImmutableArray<Outputs.YarnApplicationResponse> YarnApplications;

        [OutputConstructor]
        private GetJobResult(
            bool done,

            string driverControlFilesUri,

            string driverOutputResourceUri,

            Outputs.HadoopJobResponse hadoopJob,

            Outputs.HiveJobResponse hiveJob,

            string jobUuid,

            ImmutableDictionary<string, string> labels,

            Outputs.PigJobResponse pigJob,

            Outputs.JobPlacementResponse placement,

            Outputs.PrestoJobResponse prestoJob,

            Outputs.PySparkJobResponse pysparkJob,

            Outputs.JobReferenceResponse reference,

            Outputs.JobSchedulingResponse scheduling,

            Outputs.SparkJobResponse sparkJob,

            Outputs.SparkRJobResponse sparkRJob,

            Outputs.SparkSqlJobResponse sparkSqlJob,

            Outputs.JobStatusResponse status,

            ImmutableArray<Outputs.JobStatusResponse> statusHistory,

            ImmutableArray<Outputs.YarnApplicationResponse> yarnApplications)
        {
            Done = done;
            DriverControlFilesUri = driverControlFilesUri;
            DriverOutputResourceUri = driverOutputResourceUri;
            HadoopJob = hadoopJob;
            HiveJob = hiveJob;
            JobUuid = jobUuid;
            Labels = labels;
            PigJob = pigJob;
            Placement = placement;
            PrestoJob = prestoJob;
            PysparkJob = pysparkJob;
            Reference = reference;
            Scheduling = scheduling;
            SparkJob = sparkJob;
            SparkRJob = sparkRJob;
            SparkSqlJob = sparkSqlJob;
            Status = status;
            StatusHistory = statusHistory;
            YarnApplications = yarnApplications;
        }
    }
}
