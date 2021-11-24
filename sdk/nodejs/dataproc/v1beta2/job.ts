// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Submits a job to a cluster.
 * Auto-naming is currently not supported for this resource.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataproc/v1beta2:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * Indicates whether the job is completed. If the value is false, the job is still in progress. If true, the job is completed, and status.state field will indicate if it was successful, failed, or cancelled.
     */
    public /*out*/ readonly done!: pulumi.Output<boolean>;
    /**
     * If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
     */
    public /*out*/ readonly driverControlFilesUri!: pulumi.Output<string>;
    /**
     * A URI pointing to the location of the stdout of the job's driver program.
     */
    public /*out*/ readonly driverOutputResourceUri!: pulumi.Output<string>;
    /**
     * Optional. Job is a Hadoop job.
     */
    public readonly hadoopJob!: pulumi.Output<outputs.dataproc.v1beta2.HadoopJobResponse>;
    /**
     * Optional. Job is a Hive job.
     */
    public readonly hiveJob!: pulumi.Output<outputs.dataproc.v1beta2.HiveJobResponse>;
    /**
     * A UUID that uniquely identifies a job within the project over time. This is in contrast to a user-settable reference.job_id that may be reused over time.
     */
    public /*out*/ readonly jobUuid!: pulumi.Output<string>;
    /**
     * Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Optional. Job is a Pig job.
     */
    public readonly pigJob!: pulumi.Output<outputs.dataproc.v1beta2.PigJobResponse>;
    /**
     * Job information, including how, when, and where to run the job.
     */
    public readonly placement!: pulumi.Output<outputs.dataproc.v1beta2.JobPlacementResponse>;
    /**
     * Optional. Job is a Presto job.
     */
    public readonly prestoJob!: pulumi.Output<outputs.dataproc.v1beta2.PrestoJobResponse>;
    /**
     * Optional. Job is a PySpark job.
     */
    public readonly pysparkJob!: pulumi.Output<outputs.dataproc.v1beta2.PySparkJobResponse>;
    /**
     * Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
     */
    public readonly reference!: pulumi.Output<outputs.dataproc.v1beta2.JobReferenceResponse>;
    /**
     * Optional. Job scheduling configuration.
     */
    public readonly scheduling!: pulumi.Output<outputs.dataproc.v1beta2.JobSchedulingResponse>;
    /**
     * Optional. Job is a Spark job.
     */
    public readonly sparkJob!: pulumi.Output<outputs.dataproc.v1beta2.SparkJobResponse>;
    /**
     * Optional. Job is a SparkR job.
     */
    public readonly sparkRJob!: pulumi.Output<outputs.dataproc.v1beta2.SparkRJobResponse>;
    /**
     * Optional. Job is a SparkSql job.
     */
    public readonly sparkSqlJob!: pulumi.Output<outputs.dataproc.v1beta2.SparkSqlJobResponse>;
    /**
     * The job status. Additional application-specific status information may be contained in the type_job and yarn_applications fields.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.dataproc.v1beta2.JobStatusResponse>;
    /**
     * The previous job status.
     */
    public /*out*/ readonly statusHistory!: pulumi.Output<outputs.dataproc.v1beta2.JobStatusResponse[]>;
    /**
     * The email address of the user submitting the job. For jobs submitted on the cluster, the address is username@hostname.
     */
    public /*out*/ readonly submittedBy!: pulumi.Output<string>;
    /**
     * The collection of YARN applications spun up by this job.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
     */
    public /*out*/ readonly yarnApplications!: pulumi.Output<outputs.dataproc.v1beta2.YarnApplicationResponse[]>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.placement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'placement'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["hadoopJob"] = args ? args.hadoopJob : undefined;
            resourceInputs["hiveJob"] = args ? args.hiveJob : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["pigJob"] = args ? args.pigJob : undefined;
            resourceInputs["placement"] = args ? args.placement : undefined;
            resourceInputs["prestoJob"] = args ? args.prestoJob : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pysparkJob"] = args ? args.pysparkJob : undefined;
            resourceInputs["reference"] = args ? args.reference : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["scheduling"] = args ? args.scheduling : undefined;
            resourceInputs["sparkJob"] = args ? args.sparkJob : undefined;
            resourceInputs["sparkRJob"] = args ? args.sparkRJob : undefined;
            resourceInputs["sparkSqlJob"] = args ? args.sparkSqlJob : undefined;
            resourceInputs["done"] = undefined /*out*/;
            resourceInputs["driverControlFilesUri"] = undefined /*out*/;
            resourceInputs["driverOutputResourceUri"] = undefined /*out*/;
            resourceInputs["jobUuid"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusHistory"] = undefined /*out*/;
            resourceInputs["submittedBy"] = undefined /*out*/;
            resourceInputs["yarnApplications"] = undefined /*out*/;
        } else {
            resourceInputs["done"] = undefined /*out*/;
            resourceInputs["driverControlFilesUri"] = undefined /*out*/;
            resourceInputs["driverOutputResourceUri"] = undefined /*out*/;
            resourceInputs["hadoopJob"] = undefined /*out*/;
            resourceInputs["hiveJob"] = undefined /*out*/;
            resourceInputs["jobUuid"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["pigJob"] = undefined /*out*/;
            resourceInputs["placement"] = undefined /*out*/;
            resourceInputs["prestoJob"] = undefined /*out*/;
            resourceInputs["pysparkJob"] = undefined /*out*/;
            resourceInputs["reference"] = undefined /*out*/;
            resourceInputs["scheduling"] = undefined /*out*/;
            resourceInputs["sparkJob"] = undefined /*out*/;
            resourceInputs["sparkRJob"] = undefined /*out*/;
            resourceInputs["sparkSqlJob"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusHistory"] = undefined /*out*/;
            resourceInputs["submittedBy"] = undefined /*out*/;
            resourceInputs["yarnApplications"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Job.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * Optional. Job is a Hadoop job.
     */
    hadoopJob?: pulumi.Input<inputs.dataproc.v1beta2.HadoopJobArgs>;
    /**
     * Optional. Job is a Hive job.
     */
    hiveJob?: pulumi.Input<inputs.dataproc.v1beta2.HiveJobArgs>;
    /**
     * Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. Job is a Pig job.
     */
    pigJob?: pulumi.Input<inputs.dataproc.v1beta2.PigJobArgs>;
    /**
     * Job information, including how, when, and where to run the job.
     */
    placement: pulumi.Input<inputs.dataproc.v1beta2.JobPlacementArgs>;
    /**
     * Optional. Job is a Presto job.
     */
    prestoJob?: pulumi.Input<inputs.dataproc.v1beta2.PrestoJobArgs>;
    project?: pulumi.Input<string>;
    /**
     * Optional. Job is a PySpark job.
     */
    pysparkJob?: pulumi.Input<inputs.dataproc.v1beta2.PySparkJobArgs>;
    /**
     * Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
     */
    reference?: pulumi.Input<inputs.dataproc.v1beta2.JobReferenceArgs>;
    region: pulumi.Input<string>;
    /**
     * Optional. A unique id used to identify the request. If the server receives two SubmitJobRequest (https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1beta2#google.cloud.dataproc.v1.SubmitJobRequest)s with the same id, then the second request will be ignored and the first Job created and stored in the backend is returned.It is recommended to always set this value to a UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier).The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
     */
    requestId?: pulumi.Input<string>;
    /**
     * Optional. Job scheduling configuration.
     */
    scheduling?: pulumi.Input<inputs.dataproc.v1beta2.JobSchedulingArgs>;
    /**
     * Optional. Job is a Spark job.
     */
    sparkJob?: pulumi.Input<inputs.dataproc.v1beta2.SparkJobArgs>;
    /**
     * Optional. Job is a SparkR job.
     */
    sparkRJob?: pulumi.Input<inputs.dataproc.v1beta2.SparkRJobArgs>;
    /**
     * Optional. Job is a SparkSql job.
     */
    sparkSqlJob?: pulumi.Input<inputs.dataproc.v1beta2.SparkSqlJobArgs>;
}
