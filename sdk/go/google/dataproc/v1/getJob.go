// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the resource representation for a job in a project.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("google-native:dataproc/v1:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	JobId   string `pulumi:"jobId"`
	Project string `pulumi:"project"`
	Region  string `pulumi:"region"`
}

type LookupJobResult struct {
	// Indicates whether the job is completed. If the value is false, the job is still in progress. If true, the job is completed, and status.state field will indicate if it was successful, failed, or cancelled.
	Done bool `pulumi:"done"`
	// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
	DriverControlFilesUri string `pulumi:"driverControlFilesUri"`
	// A URI pointing to the location of the stdout of the job's driver program.
	DriverOutputResourceUri string `pulumi:"driverOutputResourceUri"`
	// Optional. Job is a Hadoop job.
	HadoopJob HadoopJobResponse `pulumi:"hadoopJob"`
	// Optional. Job is a Hive job.
	HiveJob HiveJobResponse `pulumi:"hiveJob"`
	// A UUID that uniquely identifies a job within the project over time. This is in contrast to a user-settable reference.job_id that may be reused over time.
	JobUuid string `pulumi:"jobUuid"`
	// Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Job is a Pig job.
	PigJob PigJobResponse `pulumi:"pigJob"`
	// Job information, including how, when, and where to run the job.
	Placement JobPlacementResponse `pulumi:"placement"`
	// Optional. Job is a Presto job.
	PrestoJob PrestoJobResponse `pulumi:"prestoJob"`
	// Optional. Job is a PySpark job.
	PysparkJob PySparkJobResponse `pulumi:"pysparkJob"`
	// Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
	Reference JobReferenceResponse `pulumi:"reference"`
	// Optional. Job scheduling configuration.
	Scheduling JobSchedulingResponse `pulumi:"scheduling"`
	// Optional. Job is a Spark job.
	SparkJob SparkJobResponse `pulumi:"sparkJob"`
	// Optional. Job is a SparkR job.
	SparkRJob SparkRJobResponse `pulumi:"sparkRJob"`
	// Optional. Job is a SparkSql job.
	SparkSqlJob SparkSqlJobResponse `pulumi:"sparkSqlJob"`
	// The job status. Additional application-specific status information may be contained in the type_job and yarn_applications fields.
	Status JobStatusResponse `pulumi:"status"`
	// The previous job status.
	StatusHistory []JobStatusResponse `pulumi:"statusHistory"`
	// The collection of YARN applications spun up by this job.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
	YarnApplications []YarnApplicationResponse `pulumi:"yarnApplications"`
}
